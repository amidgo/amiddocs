package middleware

import (
	"github.com/amidgo/amiddocs/pkg/middleware/config"
	"github.com/amidgo/amiddocs/pkg/middleware/cors"
	"github.com/amidgo/amiddocs/pkg/middleware/logger"
	"github.com/gofiber/fiber/v2"
)

func SetUpMiddleWare(a *fiber.App) {
	config.SetUpConfig()
	logger.SetUp(a)
	cors.DefaultCors(a)
}
