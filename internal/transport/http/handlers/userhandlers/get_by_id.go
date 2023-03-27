package userhandlers

import (
	"net/http"

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
//	@Param			id	query		int	true	"get user by id"
//	@Success		200	{object}	usermodel.UserDTO
//	@Failure		400	{object}	amiderrors.ErrorResponse
//	@Failure		500	{object}	amiderrors.ErrorResponse
//	@Router			/users/get-by-id [get]
func (h *UserHandler) GetUserById(c *fiber.Ctx) error {
	id := uint64(c.QueryInt(_ID_Q, 0))
	usr, err := h.userS.GetUserById(c.Context(), id)
	if err != nil {

		return err.SendWithFiber(c)
	}
	return c.Status(http.StatusOK).JSON(usr)
}
