package meltcd

import (
	"fmt"
	"os"
	"strings"
)

func getServer() string {
	server := "http://127.0.0.1:11771"
	serverEnvVar := os.Getenv("MELTCD_SERVER")
	if serverEnvVar != "" {
		server, _ = strings.CutSuffix(serverEnvVar, "/")
	}

	return server
}

func info(text string, args ...any) {
	fmt.Printf(text+"\n", args...)
}

func errorMsg(text string, args ...any) {
	fmt.Printf(text+"\n", args...)
}
