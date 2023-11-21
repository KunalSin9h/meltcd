package meltcd

import (
	"fmt"
	"os"
	"strings"
)

func getHost() string {
	server := "http://127.0.0.1:11771"
	if os.Getenv("MELTCD_SERVER") != "" {
		server, _ = strings.CutSuffix(os.Getenv("MELTCD_SERVER"), "/")
	}

	return server
}

func info(text string, args ...any) {
	fmt.Printf(text+"\n", args...)
}

func error_msg(text string, args ...any) {
	fmt.Printf(text+"\n", args...)
}
