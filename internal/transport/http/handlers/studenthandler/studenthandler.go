package studenthandler

import (
	"context"

	"github.com/amidgo/amiddocs/internal/jwttoken"
	"github.com/amidgo/amiddocs/internal/models/studentmodel"
	"github.com/amidgo/amiddocs/internal/models/usermodel/userfields"
	"github.com/gofiber/fiber/v2"
)

const _PROVIDER = "internal/transport/http/handlers/studenthandler"

const (
	_ROUTE_PATH = "/students"
	_CREATE     = "/create"
	_GET_BY_ID  = "/get-by-id"
)

type studentService interface {
	CreateStudent(ctx context.Context, student *studentmodel.CreateStudentDTO) (*studentmodel.StudentDTO, error)
}

type studentProvider interface {
	StudentById(ctx context.Context, id uint64) (*studentmodel.StudentDTO, error)
}

type roleValidator interface {
	ValidateRole(c *fiber.Ctx, role userfields.Role) error
}

type studentHandler struct {
	studentP studentProvider
	studentS studentService
}

func SetUp(
	app *fiber.App,
	jwt func(c *fiber.Ctx) error,
	studentS studentService,
	studentP studentProvider,
) {
	handler := studentHandler{studentS: studentS, studentP: studentP}
	route := app.Group(_ROUTE_PATH)

	route.Get(_GET_BY_ID, handler.GetStudentById)
	route.Post(_CREATE, jwt, jwttoken.AdminAccess, handler.CreateStudent)
}
