package api

import (
	"bufio"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/meltred/meltcd/internal/core"
	"github.com/valyala/fasthttp"
)

// LiveLogs godoc
//
//	@summary	Get Live Logs using SSE
//	@tags		General
//	@security	ApiKeyAuth
//	@success	200	string	string
//	@router		/logs/live [get]
func LiveLogs(c *fiber.Ctx) error {
	// Server Sent Events
	c.Set("Content-Type", "text/event-stream")
	c.Set("Cache-Control", "no-cache")
	c.Set("Connection", "keep-alive")
	c.Set("Transfer-Encoding", "chunked")

	c.Context().SetBodyStreamWriter(fasthttp.StreamWriter(func(w *bufio.Writer) {
		core.LogsStream = make(chan []byte)

		defer func() {
			close(core.LogsStream)
			core.LogsStream = nil
			fmt.Println("Defer called")
			// TODO: make it call
		}()

		for l := range core.LogsStream {
			fmt.Fprintf(w, "data: %s\n\n", string(l))
			w.Flush()
		}
	}))

	return nil
}

// Logs godoc
//
//	@summary	Get Logs
//	@tags		General
//	@security	ApiKeyAuth
//	@success	200	string	string
//	@router		/logs [get]
func Logs(c *fiber.Ctx) error {
	return nil
}
