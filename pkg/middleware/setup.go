package middleware

import (
	"github.com/amidgo/amiddocs/pkg/middleware/config"
	"github.com/amidgo/amiddocs/pkg/middleware/cors"
	"github.com/gofiber/fiber/v2"
)

func SetUpMiddleWare(a *fiber.App) {
	cors.DefaultCors(a)
	config.SetUpConfig()
}
