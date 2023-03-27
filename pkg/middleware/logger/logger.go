package logger

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetUp(app *fiber.App) {
	app.Use(logger.New(logger.ConfigDefault))
}
