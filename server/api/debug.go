package api

import (
	"net/http"
	"runtime"

	"github.com/gofiber/fiber/v2"
	"github.com/meltred/meltcd/internal/core"
)

// LiveLogs godoc
//
//	@summary	Get Session and Open Connections
//	@tags		Debug
//	@security	ApiKeyAuth
//	@success	200	object any
//	@router		/connections [get]
func Connections(c *fiber.Ctx) error {
	m := map[string]any{
		"open-connections": c.App().Server().GetOpenConnectionsCount(),
		"sessions":         len(core.CurrentSession.Sessions),
	}
	return c.Status(http.StatusOK).JSON(m)
}

// LiveLogs godoc
//
//	@summary	Get System memory, allocation, Go Routines and GC Count
//	@tags		Debug
//	@security	ApiKeyAuth
//	@success	200	object any
//	@router		/infos [get]
func SystemInfo(c *fiber.Ctx) error {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	res := map[string]any{
		"Alloc":      bToMb(m.Alloc),
		"TotalAlloc": bToMb(m.TotalAlloc),
		"tSys":       bToMb(m.Sys),
		"tNumGC":     m.NumGC,
		"goroutines": runtime.NumGoroutine(),
	}

	return c.JSON(res)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
