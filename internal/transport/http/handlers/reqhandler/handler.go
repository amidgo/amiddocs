package reqhandler

import (
	"context"
	"io"

	"github.com/amidgo/amiddocs/internal/models/reqmodel"
	"github.com/amidgo/amiddocs/internal/models/reqmodel/reqfields"
	"github.com/amidgo/amiddocs/internal/models/usermodel/userfields"
	"github.com/amidgo/amiddocs/internal/transport/http/handlers"
	"github.com/gofiber/fiber/v2"
)

const _PROVIDER = "internal/transport/http/handlers/streqhandler"

type requestProvider interface {
	RequestListByUser(ctx context.Context, userId uint64) ([]*reqmodel.RequestDTO, error)
	RequestListByDepartmentId(ctx context.Context, depId uint64) ([]*reqmodel.RequestViewDTO, error)
}

type docGenerator interface {
	GenerateDocument(ctx context.Context, wr io.Writer, reqId uint64) error
}

type requestService interface {
	UpdateRequestStatus(ctx context.Context, reqId uint64, status reqfields.Status) error
	SendRequest(ctx context.Context, roles []userfields.Role, streq *reqmodel.CreateRequestDTO) (*reqmodel.RequestDTO, error)
	DeleteRequest(ctx context.Context, userId, requestId uint64) error
}

type requestHandler struct {
	docgen  docGenerator
	reqprov requestProvider
	reqser  requestService
	jwtser  handlers.JwtManager
}

const (
	_ROUTE_PATH             = "/requests"
	_SEND                   = "/send"
	_REQUEST_INFO           = "/my-requests"
	_GENERATE_DOCUMENT      = "/generate-document"
	_REQUESTS_BY_DEPARTMENT = "/by-department-id"
	_SET_DONE               = "/set-done"
	_CANCEL                 = "/cancel"
)

func SetUp(app *fiber.App, jwt handlers.JwtManager, streqser requestService, jwtser handlers.JwtManager, reqprov requestProvider, docgen docGenerator) {

	route := app.Group(_ROUTE_PATH)
	handler := &requestHandler{reqser: streqser, jwtser: jwtser, reqprov: reqprov, docgen: docgen}

	route.Delete(_CANCEL, jwt.Ware(), handler.CancelRequest)

	route.Post(_SEND, jwt.Ware(), handler.Send)
	route.Get(_GENERATE_DOCUMENT, jwt.Ware(), jwt.SecretaryAccess, handler.GenerateDocumentFromRequest)

	route.Patch(_SET_DONE, jwt.Ware(), jwt.SecretaryAccess, handler.SetRequestStatusDone)

	route.Get(_REQUESTS_BY_DEPARTMENT, jwt.Ware(), jwt.SecretaryAccess, handler.DepartmentRequests)
	route.Get(_REQUEST_INFO, jwt.Ware(), handler.MyRequests)
}
