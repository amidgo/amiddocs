package doctemphandler

import (
	"bytes"
	"errors"
	"net/http"
	"strconv"

	"github.com/amidgo/amiddocs/internal/models/doctempmodel"
	"github.com/amidgo/amiddocs/internal/models/reqmodel/reqfields"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
	"github.com/gofiber/fiber/v2"
)

// LoadTemplate godoc
//
//	@Summary		load document template
//	@Description	load or update template in files
//	@Tags			document-templates
//	@Accept			octet-stream
//	@Produce		json
//	@Param			depId		query		uint64					true	"department id"
//	@Param			type		query		reqfields.DocumentType	true	"document type"
//	@Param			document	formData	file					true	"document binary (.docx file)"
//
//	@Success		200			{object}	doctempmodel.DocumentTemplateDTO
//	@Failure		400			{object}	amiderrors.ErrorResponse
//	@Failure		403			{object}	amiderrors.ErrorResponse
//	@Failure		500			{object}	amiderrors.ErrorResponse
//
//	@Security		Bearer
//	@Router			/document-templates/upload [post]
func (h *docTempHandler) LoadTemplate(c *fiber.Ctx) error {
	depId, err := strconv.ParseUint(c.Query("depId"), 10, 64)
	if err != nil {
		return amiderrors.NewInternalErrorResponse(err, amiderrors.NewCause("get department id from query", "LoadTemplate", _PROVIDER))
	}
	docType := c.Query("type")
	if len(docType) == 0 {
		return amiderrors.NewInternalErrorResponse(
			errors.New("missing doc type query"),
			amiderrors.NewCause("get doctype from query", "LoadTemplate", _PROVIDER),
		)
	}
	doc, err := c.FormFile("document")
	if err != nil {
		return amiderrors.NewInternalErrorResponse(
			errors.New("missing body"),
			amiderrors.NewCause("get body", "LoadTemplate", _PROVIDER),
		)
	}
	document := &bytes.Buffer{}
	file, err := doc.Open()
	defer file.Close()
	document.ReadFrom(file)
	template := doctempmodel.NewCreateDocTemplate(depId, reqfields.DocumentType(docType), document.Bytes())
	tempDTO, err := h.tempSer.UploadTemplate(c.UserContext(), template)
	if err != nil {
		return err
	}
	return c.Status(http.StatusCreated).JSON(tempDTO)
}
