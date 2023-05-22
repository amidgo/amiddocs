package jwttoken

import (
	"github.com/amidgo/amiddocs/internal/errorutils/tokenerror"
	"github.com/amidgo/amiddocs/internal/models/usermodel/userfields"
	"github.com/gofiber/fiber/v2"
)

func (tf *TokenMaster) roleAccess(c *fiber.Ctx, role userfields.Role) error {
	if err := tf.ValidateRole(c, userfields.ADMIN); err != nil {
		return tokenerror.FORBIDDEN
	}
	return c.Next()
}

func (tf *TokenMaster) AdminAccess(c *fiber.Ctx) error {
	return tf.roleAccess(c, userfields.ADMIN)
}

func (tf *TokenMaster) StudentAccess(c *fiber.Ctx) error {
	return tf.roleAccess(c, userfields.STUDENT)
}

func (tf *TokenMaster) SecretaryAccess(c *fiber.Ctx) error {
	return tf.roleAccess(c, userfields.SECRETARY)
}
