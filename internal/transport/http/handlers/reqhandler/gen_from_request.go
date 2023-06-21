package reqhandler

import (
	"net/http"
	"strconv"

	"github.com/amidgo/amiddocs/pkg/amiderrors"
	"github.com/gofiber/fiber/v2"
)

// GenerateDocumentFromRequest godoc
//
//	@Summary		GenerateDocumentFromRequest
//	@Description	returns document from request
//	@Tags			requests
//	@Accept			json
//	@Produce		json
//	@Param			reqId	query		uint64	true	"request id"
//	@Failure		400		{object}	amiderrors.ErrorResponse
//	@Failure		401		{object}	amiderrors.ErrorResponse
//	@Failure		403		{object}	amiderrors.ErrorResponse
//	@Failure		404		{object}	amiderrors.ErrorResponse
//	@Failure		500		{object}	amiderrors.ErrorResponse
//	@Security		Bearer
//	@Security		Token
//	@Router			/requests/generate-document [get]
func (s *requestHandler) GenerateDocumentFromRequest(c *fiber.Ctx) error {
	reqId, err := strconv.ParseUint(c.Query("reqId"), 10, 64)
	if err != nil {
		return amiderrors.Wrap(err, amiderrors.NewCause("get request id from request", "GenerateFromRequest", _PROVIDER))
	}
	err = s.docgen.GenerateDocument(c.UserContext(), c, reqId)
	if err != nil {
		return err
	}
	c.Status(http.StatusOK)
	c.Attachment("document.docx")
	return nil
}
