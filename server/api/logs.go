package api

import (
	"bufio"
	"fmt"
	"net/http"
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
	c.Status(http.StatusOK)

	c.Context().SetBodyStreamWriter(fasthttp.StreamWriter(func(w *bufio.Writer) {
		core.LogsStream = make(chan []byte)

		for l := range core.LogsStream {
			d := formatSSEMessage("log", string(l))

			_, err := fmt.Fprint(w, d)
			if err != nil {
				continue
			}

			err = w.Flush()

			// Connection is closed now
			if err != nil {
				close(core.LogsStream)
				core.LogsStream = nil
				break
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
	c.Status(http.StatusOK)
	return nil
}

func formatSSEMessage(eventType, data string) string {
	sb := strings.Builder{}

	sb.WriteString(fmt.Sprintf("event: %s\n", eventType))
	sb.WriteString(fmt.Sprintf("retry: %d\n", 15000))
	sb.WriteString(fmt.Sprintf("data: %v\n\n", data))

	return sb.String()
}
