package meltcd

import (
	"fmt"
	"os"
)

func getHost() string {
	// TODO: remove edge case like `/` in ending of env var
	// but we are add `/` so this will make the url invalid
	// MAKE IT MORE NEAT
	server := "http://127.0.0.1:11771"
	if os.Getenv("MELTCD_SERVER") != "" {
		server = os.Getenv("MELTCD_SERVER")
	}

	return server
}

func info(text string, args ...any) {
	fmt.Printf(text+"\n", args...)
}

func error_msg(text string, args ...any) {
	fmt.Printf(text+"\n", args...)
}
