package core

import (
	"errors"
	"fmt"
	"meltred/meltcd/internal/core/application"

	"github.com/charmbracelet/log"
)

var Applications []*application.Application

func Register(app *application.Application) error {
	log.Info("Registering application", "name", app.Name)

	for _, regApp := range Applications {
		if regApp.Name == app.Name {
			return errors.New(fmt.Sprintf("App already exists with name: %s\n", app.Name))
		}
	}

	go app.Run()
	Applications = append(Applications, app)

	return nil
}
