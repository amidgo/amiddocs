package fiberconfig

import (
	"log"
	"net/http"

	"github.com/amidgo/amiddocs/pkg/amiderrors"
	"github.com/gofiber/fiber/v2"
)

func Config() fiber.Config {
	return fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			switch err.(type) {
			case *amiderrors.ErrorResponse:
				err := err.(*amiderrors.ErrorResponse)
				log.Printf("Err of %s is %v", c.OriginalURL(), err.Err)
				return c.Status(err.HttpCode).JSON(err)
			default:
				c.Status(http.StatusInternalServerError).Request().SetBodyString(err.Error())
				return nil
			}
		},
	}
}
