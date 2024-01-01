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
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/meltred/meltcd/internal/core"
	Api "github.com/meltred/meltcd/server/api"
	appApi "github.com/meltred/meltcd/server/api/app"
	repoApi "github.com/meltred/meltcd/server/api/repo"
	"github.com/meltred/meltcd/version"

	"github.com/charmbracelet/log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
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

func Serve(ln net.Listener, origins string, verboseOutput bool) error {
	config := cors.ConfigDefault

	for _, o := range defaultAllowOrigins {
		origins = fmt.Sprintf("%s, http://%s, https://%s, http://%s:*, https://%s:*", origins, o, o, o, o)
	}

	config.AllowOrigins = origins

	app := fiber.New(fiber.Config{
		AppName: fmt.Sprintf("MeltCD Server v%s", version.Version),
	})

	app.Use(cors.New(config))
	app.Use(recover.New())

	encryptionKey, err := Api.GenerateToken(64)
	if err != nil {
		return err
	}

	app.Use(encryptcookie.New(encryptcookie.Config{
		Key: encryptionKey[:32],
	}))

	if verboseOutput {
		app.Use(logger.New())
	}

	app.Get("/swagger/*", swagger.HandlerDefault)

	// Server frontend on `/`
	app.Use("/", filesystem.New(filesystem.Config{
		Root:       http.FS(frontendSource),
		Browse:     true,
		Index:      "index.html",
		PathPrefix: "static", // the name of the folder because the files will be as static/index.html
	}))

	api := app.Group("api")

	api.Get("/", CheckAPIStatus)
	api.Post("/login", Api.Login)

	apps := api.Group("apps")
	apps.Get("/", appApi.AllApplications)
	apps.Post("/", appApi.Register)
	apps.Get("/:app_name", appApi.Details)
	apps.Delete("/:app_name", appApi.Remove)
	apps.Put("/", appApi.Update)
	apps.Post("/:app_name/refresh", appApi.Refresh)
	apps.Post("/:app_name/recreate", appApi.Recreate)

	repo := api.Group("repo")
	repo.Get("/", repoApi.List)
	repo.Post("/", repoApi.Add) // url, username and password will be send in body
	repo.Delete("/", repoApi.Remove)
	repo.Put("/", repoApi.Update)

	err = core.Setup()
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	log.Infof("Listening on %s (version: %s)", ln.Addr(), version.Version)

	signals := make(chan os.Signal, 1)
	go signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM, syscall.SIGABRT, syscall.SIGILL)

	go func() {
		<-signals
		log.Info("Shutting down server...")

		if err := core.ShutDown(); err != nil {
			log.Error(err.Error())
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
