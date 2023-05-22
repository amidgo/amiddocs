package logger

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetUp(app *fiber.App) {
	config := logger.ConfigDefault
	app.Use(logger.New(config))
}
