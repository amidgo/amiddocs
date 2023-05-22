package doctemphandler

import (
	"context"
	"io"

	"github.com/amidgo/amiddocs/internal/models/doctempmodel"
	"github.com/amidgo/amiddocs/internal/models/doctypemodel/doctypefields"
	"github.com/amidgo/amiddocs/internal/transport/http/handlers"
	"github.com/gofiber/fiber/v2"
)

type docTempService interface {
	SaveTemplate(ctx context.Context, template *doctempmodel.DocumentTemplateDTO) error
}

type docTempProvider interface {
	DocumentTemplate(ctx context.Context, wr io.Writer, depId uint64, docType doctypefields.DocumentType) error
}

type docTempHandler struct {
	tempSer  docTempService
	tempProv docTempProvider
}

const _PROVIDER = "/internal/transport/http/handlers/doctemphandler"

const (
	_ROUTE         = "/document-templates"
	_LOAD_TEMPLATE = "/upload"
	_GET_TEMPLATE  = "/get"
)

func SetUp(app *fiber.App, jwt handlers.JwtManager, templateService docTempService, templateProvider docTempProvider) {
	route := app.Group(_ROUTE)
	handler := docTempHandler{tempSer: templateService, tempProv: templateProvider}
	route.Post(_LOAD_TEMPLATE, jwt.Ware(), jwt.AdminAccess, handler.LoadTemplate)
	route.Get(_GET_TEMPLATE, handler.GetTemplate)
}
