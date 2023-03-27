package userhandlers

import (
	"context"

	"github.com/amidgo/amiddocs/internal/models/tokenmodel"
	"github.com/amidgo/amiddocs/internal/models/usermodel"
	"github.com/amidgo/amiddocs/internal/models/usermodel/userfields"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
	"github.com/gofiber/fiber/v2"
)

type userService interface {
	CreateUser(context.Context, *usermodel.CreateUserDTO) (*usermodel.UserDTO, *amiderrors.ErrorResponse)
	Login(ctx context.Context, loginform usermodel.LoginForm) (*tokenmodel.TokenResponse, *amiderrors.ErrorResponse)
	GetAllUsers(ctx context.Context) ([]*usermodel.UserDTO, *amiderrors.ErrorResponse)
	GetUserById(ctx context.Context, id uint64) (*usermodel.UserDTO, *amiderrors.ErrorResponse)
}

type tokenService interface {
	ValidateRole(c *fiber.Ctx, role userfields.UserRole) *amiderrors.ErrorResponse
}

type UserHandler struct {
	userS  userService
	tokenS tokenService
}

func New(userS userService, tokenS tokenService) *UserHandler {
	return &UserHandler{userS: userS, tokenS: tokenS}
}
