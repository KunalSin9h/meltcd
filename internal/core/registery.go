package core

import (
	"fmt"
	"meltred/meltcd/internal/core/application"

	"github.com/charmbracelet/log"
)

var Applications []*application.Application

func Register(app *application.Application) error {
	log.Info("Registering application", "name", app.Name)

	for _, regApp := range Applications {
		if regApp.Name == app.Name {
			return fmt.Errorf("app already exists with name: %s", app.Name)
		}
	}

	go app.Run()
	Applications = append(Applications, app)

	return nil
}
