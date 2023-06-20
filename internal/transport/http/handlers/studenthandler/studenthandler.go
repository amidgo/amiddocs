package studenthandler

import (
	"context"

	"github.com/amidgo/amiddocs/internal/models/studentmodel"
	"github.com/amidgo/amiddocs/internal/transport/http/handlers"
	"github.com/gofiber/fiber/v2"
)

const _PROVIDER = "internal/transport/http/handlers/studenthandler"

const (
	_ROUTE_PATH = "/students"
	_CREATE     = "/create"
	_GET_BY_ID  = "/get-by-id"
	_GET_INFO   = "/info"
)

type studentService interface {
	CreateStudent(ctx context.Context, student *studentmodel.CreateStudentDTO) (*studentmodel.StudentDTO, error)
}

type studentProvider interface {
	StudentByUserId(ctx context.Context, id uint64) (*studentmodel.StudentDTO, error)
	StudentById(ctx context.Context, id uint64) (*studentmodel.StudentDTO, error)
}

type studentHandler struct {
	studentP studentProvider
	studentS studentService
	jwt      handlers.JwtManager
}

func SetUp(
	app *fiber.App,
	jwt handlers.JwtManager,
	studentS studentService,
	studentP studentProvider,
) {
	handler := studentHandler{studentS: studentS, studentP: studentP, jwt: jwt}
	route := app.Group(_ROUTE_PATH)

	route.Get(_GET_BY_ID, handler.GetStudentById)
	route.Post(_CREATE, jwt.Ware(), jwt.AdminAccess, handler.CreateStudent)
	route.Get(_GET_INFO, jwt.Ware(), jwt.StudentAccess, handler.StudentInfo)
}
