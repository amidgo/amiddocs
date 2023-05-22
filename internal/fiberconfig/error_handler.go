package fiberconfig

import (
	"log"
	"net/http"

	"github.com/amidgo/amiddocs/pkg/amiderrors"
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	switch err.(type) {
	case *amiderrors.ErrorResponse:
		err := err.(*amiderrors.ErrorResponse)
		log.Printf("Url %s %s", c.OriginalURL(), err.String())
		return c.Status(err.HttpCode).JSON(err)
	default:
		c.Status(http.StatusInternalServerError).Request().SetBodyString(err.Error())
		return nil
	}
}
