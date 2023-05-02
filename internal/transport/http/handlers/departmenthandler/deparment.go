package departmenthandler

import (
	"context"

	"github.com/amidgo/amiddocs/internal/jwttoken"
	"github.com/amidgo/amiddocs/internal/models/depmodel"
	"github.com/amidgo/amiddocs/internal/models/usermodel/userfields"
	"github.com/gofiber/fiber/v2"
)

const _PROVIDER = "internal/transport/http/handlers/departmenthandler"

type departmentService interface {
	CreateDepartment(
		ctx context.Context,
		dep *depmodel.DepartmentDTO,
	) (*depmodel.DepartmentDTO, error)
}

type departmentProvider interface {
	DepartmentById(ctx context.Context, id uint64) (*depmodel.DepartmentDTO, error)
	AllDepartments(ctx context.Context) ([]*depmodel.DepartmentDTO, error)
}

type roleValidator interface {
	ValidateRole(c *fiber.Ctx, role userfields.Role) error
}
type DepartmentHandler struct {
	depS departmentService
	depP departmentProvider
}

const (
	_DEPARTMENT_ROUTE = "/departments"
	_GET_BY_ID        = "/get-by-id"
	_GET_ALL          = "/get-all"
	_CREATE           = "/create"
)

func SetUp(
	app *fiber.App,
	jwt func(c *fiber.Ctx) error,
	depS departmentService,
	depP departmentProvider) {
	handler := &DepartmentHandler{depS: depS, depP: depP}
	route := app.Group(_DEPARTMENT_ROUTE)

	route.Get(_GET_BY_ID, handler.GetDepartmentById)
	route.Get(_GET_ALL, handler.GetAllDepartments)
	route.Post(_CREATE, jwt, jwttoken.AdminAccess, handler.CreateDepartment)
}
