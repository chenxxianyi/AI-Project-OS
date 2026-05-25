package middleware

import (
	"github.com/ai-project-os/api/internal/utils"
	"github.com/gofiber/fiber/v2"
)

func AdminRequired() fiber.Handler {
	return func(c *fiber.Ctx) error {
		role, _ := c.Locals("user_role").(string)
		if role != "admin" {
			return utils.Forbidden(c)
		}
		return c.Next()
	}
}
