package reqhandler

import (
	"net/http"
	"strconv"

	"github.com/amidgo/amiddocs/internal/models/reqmodel/reqfields"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
	"github.com/gofiber/fiber/v2"
)

// SetRequestStatusDone godoc
//
//	@Summary		SetRequestStatusDone
//	@Description	set request status done by req id
//	@Tags			requests
//	@Accept			json
//	@Produce		json
//	@Param			reqId	query		uint64	true	"request id"
//	@Success		200		{array}		reqmodel.RequestViewDTO
//	@Failure		400		{object}	amiderrors.ErrorResponse
//	@Failure		401		{object}	amiderrors.ErrorResponse
//	@Failure		403		{object}	amiderrors.ErrorResponse
//	@Failure		404		{object}	amiderrors.ErrorResponse
//	@Failure		500		{object}	amiderrors.ErrorResponse
//	@Security		Bearer
//	@Security		Token
//	@Router			/requests/set-done [patch]
func (h *requestHandler) SetRequestStatusDone(c *fiber.Ctx) error {
	reqId, err := strconv.ParseUint(c.Query("reqId"), 10, 64)
	if err != nil {
		return amiderrors.NewInternalErrorResponse(err, amiderrors.NewCause("parse reqId from query", "SetRequestStatusDone", _PROVIDER))
	}
	err = h.reqser.UpdateRequestStatus(c.UserContext(), reqId, reqfields.DONE)
	if err != nil {
		return err
	}
	return c.SendStatus(http.StatusOK)
}
