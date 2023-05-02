package reqhandler

import (
	"context"

	"github.com/amidgo/amiddocs/internal/models/reqmodel"
	"github.com/amidgo/amiddocs/internal/models/usermodel/userfields"
	"github.com/gofiber/fiber/v2"
)

const _PROVIDER = "internal/transport/http/handlers/streqhandler"

type requestProvider interface {
	RequestListByDepartmentLimit(ctx context.Context, departmentId uint64, limit uint64, offset uint64) ([]*reqmodel.RequestDTO, error)
	RequestListByDepartmentIdHistoryLimit(ctx context.Context, departmentId uint64, limit uint64, offset uint64) ([]*reqmodel.RequestDTO, error)
}

type requestService interface {
	SendRequest(ctx context.Context, roles []userfields.Role, streq *reqmodel.CreateRequestDTO) (*reqmodel.RequestDTO, error)
}

type jwtService interface {
	UserRoles(c *fiber.Ctx) ([]userfields.Role, error)
	UserID(c *fiber.Ctx) (uint64, error)
}

type requestHandler struct {
	reqprov requestProvider
	reqser  requestService
	jwtser  jwtService
}

const (
	_ROUTE_PATH = "/requests"
	_SEND       = "/send"
)

func SetUp(app *fiber.App, jwt func(c *fiber.Ctx) error, streqser requestService, jwtser jwtService, reqprov requestProvider) {

	route := app.Group(_ROUTE_PATH)
	handler := &requestHandler{reqser: streqser, jwtser: jwtser, reqprov: reqprov}

	route.Post(_SEND, jwt, handler.Send)
}
