package log

import (
	"os"
	"path"

	"github.com/charmbracelet/log"
	"github.com/meltred/meltcd/internal/core"
)

// We can have a an application with mutux to log file file
func GetGeneralLogger() (*log.Logger, error) {
	meltcdDir := core.GetMeltcdDir()
	logDir := path.Join(meltcdDir, core.LOG_DIR)
	generalLogFile := path.Join(logDir, core.GENERAL_LOG_FILE)

	file, err := os.Open(generalLogFile)
	if err != nil {
		return nil, err
	}

	logger := log.New(file)

	return logger, nil
}
