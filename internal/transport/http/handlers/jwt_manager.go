package handlers

import (
	"github.com/amidgo/amiddocs/internal/models/usermodel/userfields"
	"github.com/gofiber/fiber/v2"
)

type JwtManager interface {
	UserRoles(c *fiber.Ctx) ([]userfields.Role, error)
	UserID(c *fiber.Ctx) (uint64, error)
	Ware() func(c *fiber.Ctx) error
	AdminAccess(c *fiber.Ctx) error
	StudentAccess(c *fiber.Ctx) error
	SecretaryAccess(c *fiber.Ctx) error
}
