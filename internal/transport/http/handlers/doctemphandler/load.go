package doctemphandler

import (
	"bytes"
	"errors"
	"net/http"
	"strconv"

	"github.com/amidgo/amiddocs/internal/models/doctempmodel"
	"github.com/amidgo/amiddocs/internal/models/doctypemodel/doctypefields"
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
//	@Param			depId		query		uint64						true	"department id"
//	@Param			type		query		doctypefields.DocumentType	true	"document type"
//	@Param			document	formData	file						true	"document binary (.docx file)"
//	@Failure		400			{object}	amiderrors.ErrorResponse
//	@Failure		403			{object}	amiderrors.ErrorResponse
//	@Failure		500			{object}	amiderrors.ErrorResponse
//	@Security		Bearer
//	@Security		Token
//	@Router			/document-templates/upload [post]
func (h *docTempHandler) LoadTemplate(c *fiber.Ctx) error {
	depId, err := strconv.ParseUint(c.Query("depId"), 10, 64)
	if err != nil {
		return amiderrors.Wrap(err, amiderrors.NewCause("get department id from query", "LoadTemplate", _PROVIDER))
	}
	docType := c.Query("type")
	if len(docType) == 0 {
		return amiderrors.Wrap(
			errors.New("missing doc type query"),
			amiderrors.NewCause("get doctype from query", "LoadTemplate", _PROVIDER),
		)
	}
	doc, err := c.FormFile("document")
	if err != nil {
		return amiderrors.Wrap(
			errors.New("missing body"),
			amiderrors.NewCause("get body", "LoadTemplate", _PROVIDER),
		)
	}
	document := &bytes.Buffer{}
	file, err := doc.Open()
	if err != nil {
		return amiderrors.Wrap(
			err,
			amiderrors.NewCause("open doc", "LoadTemplate", _PROVIDER),
		)
	}
	defer file.Close()
	document.ReadFrom(file)
	template := doctempmodel.NewCreateDocTemplate(depId, doctypefields.DocumentType(docType), document.Bytes())
	err = h.tempSer.SaveTemplate(c.UserContext(), template)
	if err != nil {
		return err
	}
	c.Status(http.StatusCreated)
	return nil
}
