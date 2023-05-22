package userhandler

import (
	"context"

	"github.com/amidgo/amiddocs/internal/models/tokenmodel"
	"github.com/amidgo/amiddocs/internal/models/usermodel"
	"github.com/amidgo/amiddocs/internal/transport/http/handlers"
	"github.com/gofiber/fiber/v2"
)

const _PROVIDER = "internal/transport/http/handlers/userhandler"

type userService interface {
	CreateUser(context.Context, *usermodel.CreateUserDTO) (*usermodel.UserDTO, error)
	RefreshToken(ctx context.Context, oldRefreshToken string, userId uint64) (*tokenmodel.TokenResponse, error)
	Login(ctx context.Context, loginform *usermodel.LoginForm) (*tokenmodel.TokenResponse, error)
}

type userProvider interface {
	AllUsers(ctx context.Context) ([]*usermodel.UserDTO, error)
	UserById(ctx context.Context, id uint64) (*usermodel.UserDTO, error)
}

type userHandler struct {
	userS userService
	userP userProvider
	jwt   handlers.JwtManager
}

const _USER_PATH = "/users"

const (
	_REGISTER_USER = "/register"
	_LOGIN         = "/login"
	_GET_ALL       = "/all"
	_GET_BY_ID     = "/get-by-id"
	_GET_INFO      = "/info"
	_REFRESH_TOKEN = "/refresh-token"
)

func SetUp(
	app *fiber.App,
	jwt handlers.JwtManager,
	userS userService,
	userP userProvider,
) {

	handler := &userHandler{userS: userS, userP: userP, jwt: jwt}
	route := app.Group(_USER_PATH)

	route.Get(_GET_BY_ID, handler.GetUserById)
	route.Get(_GET_INFO, jwt.Ware(), handler.UserInfo)
	route.Get(_GET_ALL, handler.GetAllUsers)
	route.Post(_LOGIN, handler.Login)
	route.Post(_REGISTER_USER, jwt.Ware(), jwt.AdminAccess, handler.RegisterUser)
	route.Post(_REFRESH_TOKEN, handler.RefreshToken)
}
