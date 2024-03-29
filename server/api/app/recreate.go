package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kunalsin9h/meltcd/internal/core"
)

// Recreate godoc
//
//	@summary	Recreate application
//	@tags		Apps
//	@Security	ApiKeyAuth || cookies
//	@param		app_name	path	string	true	"Application name"
//	@success	200
//	@failure	500	{object}	GlobalResponse
//	@router		/apps/{app_name}/recreate [post]
func Recreate(c *fiber.Ctx) error {
	appName := c.Params("app_name")

	if err := core.Recreate(appName); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(GlobalResponse{
			Message: err.Error(),
		})
	}

	return c.SendStatus(200)
}
