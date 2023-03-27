package userrouting

import (
	"github.com/amidgo/amiddocs/internal/jwtgen"
	"github.com/amidgo/amiddocs/internal/transport/http/handlers/userhandlers"
	"github.com/gofiber/fiber/v2"
)

const _USER_PATH = "/users"

const (
	_REGISTER_USER = "/register"
	_LOGIN         = "/login"
	_GET_ALL       = "/all"
	_GET_BY_ID     = "/get-by-id"
)

func SetUp(app *fiber.App, uHandler *userhandlers.UserHandler) {

	route := app.Group(_USER_PATH)

	route.Get(_GET_BY_ID, uHandler.GetUserById)
	route.Get(_GET_ALL, uHandler.GetAllUsers)
	route.Post(_LOGIN, uHandler.Login)
	route.Post(_REGISTER_USER, jwtgen.RsJwtWare(), uHandler.RegisterUser)
}
