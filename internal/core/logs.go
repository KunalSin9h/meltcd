package core

import (
	"os"
)

// LogsStream Live Logs sharing
// shared between api and log aggregator
var LogsStream chan []byte

func StoreLog(f *os.File, d *[]byte) error {
	_, err := f.Write(*d)
	return err
}

func CreateLogFile() (*os.File, error) {
	meltcdDir := getMeltcdDir()

	_, err := os.Stat(meltcdDir)
	if err != nil {
		err = os.Mkdir(meltcdDir, os.ModePerm)
		if err != nil {
			return nil, err
		}
	}

	logFile := getLogFile()

	_, err = os.Stat(logFile)
	if err != nil {
		_, err = os.Create(logFile)
		if err != nil {
			return nil, err
		}
	}

	f, err := os.OpenFile(logFile, os.O_RDWR|os.O_APPEND, os.ModePerm)
	if err != nil {
		return nil, err
	}

	return f, nil
}
