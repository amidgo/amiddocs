package doctemphandler

import (
	"context"

	"github.com/amidgo/amiddocs/internal/jwttoken"
	"github.com/amidgo/amiddocs/internal/models/doctempmodel"
	"github.com/gofiber/fiber/v2"
)

type docTempService interface {
	UploadTemplate(ctx context.Context, template *doctempmodel.DocumentTemplateDTO) (*doctempmodel.DocumentTemplateDTO, error)
}

type docTempHandler struct {
	tempSer docTempService
}

const _PROVIDER = "/internal/transport/http/handlers/doctemphandler"

const (
	_ROUTE         = "/document-templates"
	_LOAD_TEMPLATE = "/upload"
)

func SetUp(app *fiber.App, jwt func(*fiber.Ctx) error, templateService docTempService) {
	route := app.Group(_ROUTE)
	handler := docTempHandler{tempSer: templateService}
	route.Post(_LOAD_TEMPLATE, jwt, jwttoken.AdminAccess, handler.LoadTemplate)
}
