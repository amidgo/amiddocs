package fiberconfig

import (
	"github.com/gofiber/fiber/v2"
)

func Config() fiber.Config {
	return fiber.Config{
		ErrorHandler: ErrorHandler,
	}
}
