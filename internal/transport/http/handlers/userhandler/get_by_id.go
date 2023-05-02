package userhandler

import (
	"net/http"
	"strconv"

	"github.com/amidgo/amiddocs/pkg/amiderrors"
	"github.com/gofiber/fiber/v2"
)

const _ID_Q = "id"

// GetUserById godoc
//
//	@Summary		Return User
//	@Description	return user by id from path
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			id	query		uint64	true	"get user by id"
//	@Success		200	{object}	usermodel.UserDTO
//	@Failure		400	{object}	amiderrors.ErrorResponse
//	@Failure		404	{object}	amiderrors.ErrorResponse
//	@Failure		500	{object}	amiderrors.ErrorResponse
//	@Router			/users/get-by-id [get]
func (h *UserHandler) GetUserById(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Query(_ID_Q, "word"), 10, 64)
	if err != nil {
		return amiderrors.NewInternalErrorResponse(err, amiderrors.NewCause("parse id", "GetUserById", _PROVIDER))
	}
	usr, err := h.userP.UserById(c.UserContext(), id)
	if err != nil {
		return err
	}
	return c.Status(http.StatusOK).JSON(usr)
}
