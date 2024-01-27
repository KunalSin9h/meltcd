package api

import (
	"bufio"
	"fmt"
	"net/http"
	"strings"
	"time"

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
    c.Set("Access-Control-Allow-Origin", "*")
    c.Set("Access-Control-Allow-Headers", "Content-Type")
	c.Status(http.StatusOK)

	logsStream := make(chan []byte)
	core.CurrentSession.AddSession(&logsStream)

	notifyConnClose := c.Context().Done()

	c.Context().SetBodyStreamWriter(fasthttp.StreamWriter(func(w *bufio.Writer) {
		keepAliveTickler := time.NewTicker(15 * time.Second)

		go func() {
			<-notifyConnClose
			core.CurrentSession.RemoveSession(&logsStream)
		}()

		for {
			select {
			case l := <-logsStream:
				_, err := fmt.Fprint(w, formatSSEMessage("log", string(l)))
				if err != nil {
					continue
				}

				err = w.Flush()
				// Connection is closed now
				if err != nil {
					core.CurrentSession.RemoveSession(&logsStream)
					return
				}
			case <-keepAliveTickler.C:
				fmt.Fprint(w, formatSSEMessage("message", "keepalive"))
				err := w.Flush()
				// Connection is closed now
				if err != nil {
					core.CurrentSession.RemoveSession(&logsStream)
					return
				}
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
