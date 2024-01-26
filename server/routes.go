/*
Copyright 2023 - PRESENT Meltred

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package server

import (
	"embed"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"log/slog"

	"github.com/meltred/meltcd/internal/core"
	Api "github.com/meltred/meltcd/server/api"
	appApi "github.com/meltred/meltcd/server/api/app"
	repoApi "github.com/meltred/meltcd/server/api/repo"
	"github.com/meltred/meltcd/server/middleware"
	"github.com/meltred/meltcd/version"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
)

//go:embed static/*
var frontendSource embed.FS

var defaultAllowOrigins = []string{
	"localhost",
	"127.0.0.1",
	"0.0.0.0",
}

type LogWriter struct {
	LogFile *os.File
	Stream  *chan []byte
}

func (lw LogWriter) Write(p []byte) (n int, err error) {
	go func() {
		data := make([]byte, len(p))
		copy(data, p)

		err = core.StoreLog(lw.LogFile, &data)

		if err != nil {
			fmt.Println("Failed to store logs in file")
			fmt.Println(err.Error())
		}

		if lw.Stream != nil && *lw.Stream != nil {
			*lw.Stream <- data
		}
	}()

	return len(p), nil
}

// Verify is LogWriter implements io.Writer Interface
var _ io.Writer = (*LogWriter)(nil)

func Serve(ln net.Listener, origins string, verboseOutput bool) error {
	logFile, err := core.CreateLogFile()
	if err != nil {
		return err
	}
	defer logFile.Close()

	lw := LogWriter{
		LogFile: logFile,
		Stream:  &core.LogsStream,
	}

	// Setting default slog logger
	cLogger := slog.New(slog.NewJSONHandler(lw, nil))
	slog.SetDefault(cLogger)

	err = core.Setup()

	if err != nil {
		slog.Error(err.Error())
		return err
	}

	config := cors.ConfigDefault

	for _, o := range defaultAllowOrigins {
		origins = fmt.Sprintf("%s, http://%s, https://%s, http://%s:*, https://%s:*", origins, o, o, o, o)
	}

	config.AllowOrigins = origins

	app := fiber.New(fiber.Config{
		AppName: fmt.Sprintf("MeltCD Server v%s", version.Version),

		ErrorHandler: func(c *fiber.Ctx, err error) error {
			if err != nil {
				slog.Error(err.Error())
				return err
			}
			return nil
		},
	})

	app.Use(cors.New(config))
	app.Use(recover.New())

	if verboseOutput {
		app.Use(logger.New())
	}

	app.Get("/swagger/*", swagger.HandlerDefault)

	// FRONTEND INSTRUMENTATIONS
	allFrontendRoutes := []string{
		"/",
		"/login",
		"/apps",
		"/repos",
		"/secrets",
		"/users",
		"/settings",
		"/logs",
	}

	for _, route := range allFrontendRoutes {
		// Server frontend
		app.Use(route, filesystem.New(filesystem.Config{
			Root:       http.FS(frontendSource),
			Browse:     true,
			Index:      "index.html",
			PathPrefix: "static", // the name of the folder because the files will be as static/index.html
		}))
	}

	// This does some crepy work
	// it captures the request which is dynamic and app.Use (the above) cant capture
	// then it redirects to /apps telling to future navigate to /apps/x
	// it looks bad but only choice
	app.Get("/apps/:app_name", func(c *fiber.Ctx) error {
		appName := c.Params("app_name")
		return c.Redirect(fmt.Sprintf("/apps?redirect=%s", appName), http.StatusTemporaryRedirect)
	})

	// ---------------------------------------------------------------------

	// API
	// PROTECTED BY Rate Limiting
	// And Encrypted Cookies
	api := app.Group("api")

	if strings.TrimSpace(os.Getenv("RL_DISABLE")) != "true" {
		slog.Warn("Rate Limiting is enabled by default, to disable set RL_DISABLE=true")
		api.Use(limiter.New(*rateLimiterConfig()))
	}

	api.Use(encryptcookie.New(encryptcookie.Config{
		Key: encryptcookie.GenerateKey(),
	}))

	api.Get("/", CheckAPIStatus)
	api.Post("/login", Api.Login)

	// Logs
	api.Get("/logs", Api.Logs)
	// Live Logs using SSE
	api.Get("/logs/live", Api.LiveLogs)

	users := api.Group("users", middleware.VerifyUser)
	users.Get("/", Api.GetUsers)
	users.Get("/current", Api.GetUsername)
	users.Patch("/:username/password", Api.ChangePassword)
	users.Patch("/:username/username", Api.ChangeUsername)

	apps := api.Group("apps", middleware.VerifyUser)
	apps.Get("/", appApi.AllApplications)
	apps.Post("/", appApi.Register)
	apps.Get("/:app_name", appApi.Details)
	apps.Delete("/:app_name", appApi.Remove)
	apps.Put("/", appApi.Update)
	apps.Post("/:app_name/refresh", appApi.Refresh)
	apps.Post("/:app_name/recreate", appApi.Recreate)

	repo := api.Group("repo", middleware.VerifyUser)
	repo.Get("/", repoApi.List)
	repo.Post("/", repoApi.Add) // url, username and password will be send in body
	repo.Delete("/", repoApi.Remove)
	repo.Put("/", repoApi.Update)

	info := fmt.Sprintf("Listening on %s (version: %s)\n", ln.Addr(), version.Version)

	slog.Info(info)
	fmt.Println(info)

	signals := make(chan os.Signal, 1)
	go signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM, syscall.SIGABRT, syscall.SIGILL)

	go func() {
		<-signals
		slog.Info("Shutting down server...")

		if err := core.ShutDown(); err != nil {
			slog.Error(err.Error())
			os.Exit(1)
		}

		os.Exit(0)
	}()

	return app.Listener(ln)
}

// @summary	Check server status
// @tags		General
// @produce	plain
// @router		/ [get]
func CheckAPIStatus(c *fiber.Ctx) error {
	return c.Status(200).SendString(fmt.Sprintf("Meltcd API is running (version: %s)\n", version.Version))
}
