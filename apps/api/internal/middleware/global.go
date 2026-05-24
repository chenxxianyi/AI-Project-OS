package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	requestid "github.com/gofiber/fiber/v2/middleware/requestid"
	"go.uber.org/zap"
)

func RegisterGlobal(app *fiber.App) {
	app.Use(requestid.New())
	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
		StackTraceHandler: func(c *fiber.Ctx, e interface{}) {
			zap.L().Error("panic recovered",
				zap.Any("error", e),
				zap.String("path", c.Path()),
			)
		},
	}))
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
	}))
	app.Use(func(c *fiber.Ctx) error {
		zap.L().Debug("request",
			zap.String("method", c.Method()),
			zap.String("path", c.Path()),
		)
		return c.Next()
	})
}
