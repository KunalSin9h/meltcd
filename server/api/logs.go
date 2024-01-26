package api

import (
	"bufio"
	"fmt"
	"strings"

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
		}()

		for l := range core.LogsStream {
			d, err := formatSSEMessage("log", string(l))
			if err != nil {
				return
			}

			_, err = fmt.Fprint(w, d)
			if err != nil {
				return
			}

			err = w.Flush()

			// Connection is closed now
			if err != nil {
				return
			}
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

func formatSSEMessage(eventType, data string) (string, error) {
	sb := strings.Builder{}

	sb.WriteString(fmt.Sprintf("event: %s\n", eventType))
	sb.WriteString(fmt.Sprintf("retry: %d\n", 15000))
	sb.WriteString(fmt.Sprintf("data: %v\n\n", data))

	return sb.String(), nil
}
