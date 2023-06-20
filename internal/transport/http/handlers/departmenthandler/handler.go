package departmenthandler

import (
	"context"

	"github.com/amidgo/amiddocs/internal/models/depmodel"
	"github.com/amidgo/amiddocs/internal/transport/http/handlers"
	"github.com/gofiber/fiber/v2"
)

const _PROVIDER = "internal/transport/http/handlers/departmenthandler"

type departmentService interface {
	CreateDepartment(
		ctx context.Context,
		dep *depmodel.CreateDepartmentDTO,
	) (*depmodel.DepartmentDTO, error)
}

type departmentProvider interface {
	DepartmentById(ctx context.Context, id uint64) (*depmodel.DepartmentDTO, error)
	DepartmentList(ctx context.Context) ([]*depmodel.DepartmentDTO, error)
	DepartmentListWithTypes(ctx context.Context) ([]*depmodel.DepartmentTypes, error)
}

type departmentHandler struct {
	depS departmentService
	depP departmentProvider
}

const (
	_DEPARTMENT_ROUTE   = "/departments"
	_GET_BY_ID          = "/get-by-id"
	_GET_ALL            = "/get-all"
	_CREATE             = "/create"
	_STUDENT_DEPARTMENT = "/student-department"
	_GET_ALL_TYPES      = "/get-all-types"
)

func SetUp(
	app *fiber.App,
	jwt handlers.JwtManager,
	depS departmentService,
	depP departmentProvider,
) {
	handler := &departmentHandler{depS: depS, depP: depP}
	route := app.Group(_DEPARTMENT_ROUTE)

	route.Get(_GET_BY_ID, handler.GetDepartmentById)
	route.Get(_GET_ALL, handler.GetAllDepartments)
	route.Get(_GET_ALL_TYPES, handler.GetAllWithTypes)
	route.Post(_CREATE, jwt.Ware(), jwt.AdminAccess, handler.CreateDepartment)
}
