package userhandler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// GetAllUsers godoc
//
//	@Summary		Get All Users
//	@Description	get all users from database
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		usermodel.UserDTO
//	@Failure		500	{object}	amiderrors.ErrorResponse
//	@Security		Token
//	@Router			/users/all [get]
func (h *userHandler) GetAllUsers(c *fiber.Ctx) error {
	userList, err := h.userP.AllUsers(c.UserContext())
	if err != nil {
		return err
	}
	return c.Status(http.StatusOK).JSON(userList)
}
