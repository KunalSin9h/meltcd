package server

import (
	"fmt"
	"meltred/meltcd/version"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/charmbracelet/log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

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

	app.Use(cors.New())

	if verboseOutput {
		app.Use(logger.New())
	}

	app.Static("/", "./ui/dist/")

	api := app.Group("api")

	api.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(app.Stack())
	})
	api.Get("/health_check", func(c *fiber.Ctx) error {
		return c.Status(200).SendString("MeltCD is running\n")
	})

	log.Infof("Listening on %s", ln.Addr())

	signals := make(chan os.Signal, 1)
	go signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-signals
		log.Info("Shutting down server...")
		os.Exit(0)
	}()

	return app.Listener(ln)
}
