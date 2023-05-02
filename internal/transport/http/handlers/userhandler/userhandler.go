package userhandler

import (
	"context"

	"github.com/amidgo/amiddocs/internal/jwttoken"
	"github.com/amidgo/amiddocs/internal/models/tokenmodel"
	"github.com/amidgo/amiddocs/internal/models/usermodel"
	"github.com/amidgo/amiddocs/internal/models/usermodel/userfields"
	"github.com/gofiber/fiber/v2"
)

const _PROVIDER = "internal/transport/http/handlers/userhandler"

type userService interface {
	CreateUser(context.Context, *usermodel.CreateUserDTO) (*usermodel.UserDTO, error)
	Login(ctx context.Context, loginform *usermodel.LoginForm) (*tokenmodel.TokenResponse, error)
}

type userProvider interface {
	AllUsers(ctx context.Context) ([]*usermodel.UserDTO, error)
	UserById(ctx context.Context, id uint64) (*usermodel.UserDTO, error)
}

type roleValidator interface {
	ValidateRole(c *fiber.Ctx, role userfields.Role) error
}

type UserHandler struct {
	userS userService
	roleV roleValidator
	userP userProvider
}

const _USER_PATH = "/users"

const (
	_REGISTER_USER = "/register"
	_LOGIN         = "/login"
	_GET_ALL       = "/all"
	_GET_BY_ID     = "/get-by-id"
)

func SetUp(
	app *fiber.App,
	jwt func(c *fiber.Ctx) error,
	userS userService,
	userP userProvider,
) {

	handler := &UserHandler{userS: userS, userP: userP}
	route := app.Group(_USER_PATH)

	route.Get(_GET_BY_ID, handler.GetUserById)
	route.Get(_GET_ALL, handler.GetAllUsers)
	route.Post(_LOGIN, handler.Login)
	route.Post(_REGISTER_USER, jwt, jwttoken.AdminAccess, handler.RegisterUser)
}
