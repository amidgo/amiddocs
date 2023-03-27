package userhandlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// GetAllUsers godoc
//	@Summary		Get All Users
//	@Description	get all users from database
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		usermodel.UserDTO
//	@Failure		500	{object}	amiderrors.ErrorResponse
//	@Router			/users/all [get]
func (h *UserHandler) GetAllUsers(c *fiber.Ctx) error {
	userList, err := h.userS.GetAllUsers(c.Context())
	if err != nil {
		return err.SendWithFiber(c)
	}
	return c.Status(http.StatusOK).JSON(userList)
}
