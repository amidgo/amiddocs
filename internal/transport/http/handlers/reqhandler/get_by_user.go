package reqhandler

import (
	"net/http"

	"github.com/amidgo/amiddocs/pkg/amiderrors"
	"github.com/gofiber/fiber/v2"
)

// MyRequests godoc
//
//	@Summary		UserRequests
//	@Description	get user requests by id from jwt token
//	@Tags			requests
//	@Accept			json
//	@Produce		json
//
//	@Success		200	{array}		reqmodel.RequestDTO
//	@Failure		400	{object}	amiderrors.ErrorResponse
//	@Failure		401	{object}	amiderrors.ErrorResponse
//	@Failure		403	{object}	amiderrors.ErrorResponse
//	@Failure		404	{object}	amiderrors.ErrorResponse
//	@Failure		500	{object}	amiderrors.ErrorResponse
//	@Security		Bearer
//	@Security		Token
//	@Router			/requests/my-requests [get]
func (h *requestHandler) MyRequests(c *fiber.Ctx) error {
	id, err := h.jwtser.UserID(c)
	if err != nil {
		return amiderrors.NewInternalErrorResponse(err, amiderrors.NewCause("user id from token", "MyRequests", _PROVIDER))
	}
	reqList, err := h.reqprov.RequestListByUser(c.UserContext(), id)
	if err != nil {
		return err
	}
	return c.Status(http.StatusOK).JSON(reqList)
}
