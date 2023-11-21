package core

import (
	"meltred/meltcd/internal/core/application"

	"github.com/charmbracelet/log"
)

func Register(app *application.Application) error {
	log.Info("Registering application", "name", app.Name)

	//TODO: check if application with name already exists

	return nil
}
