package middleware

import (
	"github.com/amidgo/amiddocs/pkg/middleware/cors"
	"github.com/amidgo/amiddocs/pkg/middleware/logger"
	"github.com/amidgo/amiddocs/pkg/middleware/recover"
	"github.com/gofiber/fiber/v2"
)

func SetUpMiddleWare(a *fiber.App) {
	logger.SetUp(a)
	cors.SetUp(a)
	recover.SetUp(a)
}
