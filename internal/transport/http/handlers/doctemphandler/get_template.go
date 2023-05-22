package doctemphandler

import (
	"net/http"
	"strconv"

	"github.com/amidgo/amiddocs/internal/models/doctypemodel/doctypefields"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
	"github.com/gofiber/fiber/v2"
)

// GetTemplate godoc
//
//	@Summary		get template document
//	@Description	returns raw file
//	@Tags			document-templates
//	@Accept			json
//	@Produce		octet-stream
//	@Param			type			query		doctypefields.DocumentType	true	"document type"
//	@Param			departmentId	query		uint64						true	"department id"
//	@Success		200				{file}		file
//	@Failure		400				{object}	amiderrors.ErrorResponse
//	@Failure		500				{object}	amiderrors.ErrorResponse
//	@Security		Token
//	@Router			/document-templates/get [get]
func (h *docTempHandler) GetTemplate(c *fiber.Ctx) error {
	docType := doctypefields.DocumentType(c.Query("type"))
	depId, err := strconv.ParseUint(c.Query("departmentId"), 10, 64)
	if err != nil {
		return amiderrors.NewInternalErrorResponse(err, amiderrors.NewCause("parse dep id", "GetTemplate", _PROVIDER))
	}
	err = h.tempProv.DocumentTemplate(c.UserContext(), c, depId, docType)
	if err != nil {
		return err
	}
	c.Attachment("template.docx")
	c.Status(http.StatusOK)
	return nil
}
