package reqhandler

import (
	"net/http"
	"strconv"

	"github.com/amidgo/amiddocs/pkg/amiderrors"
	"github.com/gofiber/fiber/v2"
)

// CancelRequest godoc
//
//	@Summary		Cancel request by id
//	@Description	"Delete request if status == SEND"
//	@Tags			requests
//	@Accept			json
//	@Param			requestId	query	uint64	true	"request id"
//	@Success		204
//	@Failure		400	{object}	amiderrors.ErrorResponse
//	@Failure		401	{object}	amiderrors.ErrorResponse
//	@Failure		403	{object}	amiderrors.ErrorResponse
//	@Failure		404	{object}	amiderrors.ErrorResponse
//	@Failure		500	{object}	amiderrors.ErrorResponse
//	@Security		Bearer
//	@Security		Token
//	@Router			/requests/cancel [delete]
func (h *requestHandler) CancelRequest(c *fiber.Ctx) error {
	userId, err := h.jwtser.UserID(c)
	if err != nil {
		return amiderrors.Wrap(err, amiderrors.NewCause("get userId from token", "CancelRequest", _PROVIDER))
	}
	reqId, err := strconv.ParseUint(c.Query("requestId"), 10, 64)
	if err != nil {
		return amiderrors.Wrap(err, amiderrors.NewCause("get requestId from query", "CancelRequest", _PROVIDER))
	}
	err = h.reqser.DeleteRequest(c.UserContext(), userId, reqId)
	if err != nil {
		return err
	}
	return c.SendStatus(http.StatusNoContent)
}
