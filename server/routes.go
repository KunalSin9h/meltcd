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
	"meltred/meltcd/internal/core"
	"meltred/meltcd/version"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/charmbracelet/log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
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

	if verboseOutput {
		app.Use(logger.New())
	}

	// Server frontend on `/`
	app.Use("/", filesystem.New(filesystem.Config{
		Root:       http.FS(frontendSource),
		Browse:     true,
		Index:      "index.html",
		PathPrefix: "static", // the name of the folder because the files will be as static/index.html
	}))

	api := app.Group("api")
	api.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(app.Stack())
	})
	api.Get("/health_check", func(c *fiber.Ctx) error {
		return c.Status(200).SendString(fmt.Sprintf("MeltCD is running (version: %s)\n", version.Version))
	})

	application := api.Group("application")
	application.Post("/register", Register)
	application.Post("/update", Update)
	application.Post("/refresh/:app_name", Refresh)
	application.Get("/get", AllApplications)
	application.Get("/get/:app_name", Details)

	err := core.Setup()
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	log.Infof("Listening on %s (version: %s)", ln.Addr(), version.Version)

	signals := make(chan os.Signal, 1)
	go signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-signals
		log.Info("Shutting down server...")
		os.Exit(0)
	}()

	return app.Listener(ln)
}
