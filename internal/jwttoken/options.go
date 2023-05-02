package jwttoken

import (
	"github.com/amidgo/amiddocs/internal/errorutils/tokenerror"
	"github.com/amidgo/amiddocs/internal/models/usermodel/userfields"
	"github.com/gofiber/fiber/v2"
)

func roleAccess(c *fiber.Ctx, role userfields.Role) error {
	if err := validateRole(c, userfields.ADMIN); err != nil {
		return tokenerror.FORBIDDEN
	}
	return c.Next()
}

func AdminAccess(c *fiber.Ctx) error {
	return roleAccess(c, userfields.ADMIN)
}

func StudentAccess(c *fiber.Ctx) error {
	return roleAccess(c, userfields.STUDENT)
}

func SecretaryAccess(c *fiber.Ctx) error {
	return roleAccess(c, userfields.SECRETARY)
}
