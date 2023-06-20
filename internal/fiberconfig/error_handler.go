package fiberconfig

import (
	"log"

	"github.com/amidgo/amiddocs/pkg/amiderrors"
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	log.Printf("Url %s %s", c.OriginalURL(), err)
	var httpCode int
	switch err := err.(type) {
	case *amiderrors.ErrorResponse:
		httpCode = 400
	case *amiderrors.Exception:
		httpCode = err.HttpCode
	default:
		httpCode = 500
	}
	return c.Status(httpCode).JSON(amiderrors.ErrorToResponse(err))
}
