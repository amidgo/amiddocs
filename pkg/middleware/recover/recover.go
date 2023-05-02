package recover

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func SetUp(app *fiber.App) {
	app.Use(recover.New())
}
