package fiberconfig

import (
	"github.com/amidgo/amiddocs/internal/errorutils/restapierror"
	"github.com/gofiber/fiber/v2"
)

func ClientTokenHandler(clientKey string) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		if c.GetReqHeaders()["Token"] != clientKey {
			return restapierror.WRONG_CLIENT_KEY
		}
		return c.Next()
	}
}
