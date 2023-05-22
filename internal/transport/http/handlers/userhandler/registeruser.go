package userhandler

import (
	"net/http"

	"github.com/amidgo/amiddocs/internal/models/usermodel"
	"github.com/gofiber/fiber/v2"
)

// RegisterUser godoc
//
//	@Summary		RegisterUser
//	@Description	register user, require createUserModel, email should be unique
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			user	body		usermodel.CreateUserDTO	true	"create user dto"
//	@Success		201		{object}	usermodel.UserDTO
//	@Failure		400		{object}	amiderrors.ErrorResponse
//	@Failure		500		{object}	amiderrors.ErrorResponse
//	@Security		Bearer
//	@Security		Token
//	@Router			/users/register [post]
func (h *userHandler) RegisterUser(c *fiber.Ctx) error {

	createUserDTO := new(usermodel.CreateUserDTO)
	err := c.BodyParser(createUserDTO)
	if err != nil {
		return err
	}

	err = createUserDTO.Validate()
	if err != nil {
		return err
	}

	userDTO, err := h.userS.CreateUser(c.UserContext(), createUserDTO)
	if err != nil {
		return err
	}
	return c.Status(http.StatusCreated).JSON(userDTO)
}
