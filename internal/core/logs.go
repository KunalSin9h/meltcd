package core

import (
	"os"
	"slices"
	"sync"
)

// CurrentSession Value to be used by both API and CORE
var CurrentSession LogsStreamSessions

// LogsStreamSessions
type LogsStreamSessions struct {
	MU       sync.Mutex
	Sessions []*chan []byte
}

func (l *LogsStreamSessions) AddSession(s *chan []byte) {
	l.MU.Lock()
	l.Sessions = append(l.Sessions, s)
	l.MU.Unlock()
}

func (l *LogsStreamSessions) RemoveSession(s *chan []byte) {
	idx := slices.Index(l.Sessions, s)

	if idx != -1 {
		l.MU.Lock()

		l.Sessions[idx] = nil // To Be Garbage Collected
		l.Sessions = slices.Delete(l.Sessions, idx, idx+1)

		l.MU.Unlock()
	}
}

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
