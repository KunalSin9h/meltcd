package server

import (
	"errors"
	"io"
	"net/http"

	"github.com/fatih/color"
	"github.com/meltred/meltcd/internal/core"
)

func HTTPRequestWithBearerToken(method, url string, body io.Reader, json bool) (*http.Request, *http.Client, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, nil, err
	}
	req.Header.Set("X-API-Key", core.GetAccessToken())

	if json {
		req.Header.Add("Content-Type", "application/json")
	}

	return req, &http.Client{}, nil
}

func ReadAuthError(body io.Reader) error {
	msg, err := io.ReadAll(body)
	if err != nil {
		return err
	}
	color.HiCyan("Login first, do \n\t$ meltcd login\n")
	return errors.New(string(msg))
}
