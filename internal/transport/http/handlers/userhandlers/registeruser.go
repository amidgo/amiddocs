package userhandlers

import (
	"net/http"

	"github.com/amidgo/amiddocs/internal/models/usermodel"
	"github.com/amidgo/amiddocs/internal/models/usermodel/userfields"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
	"github.com/gofiber/fiber/v2"
)

type registerUserAmid struct {
	h   *UserHandler
	err *amiderrors.ErrorResponse
}

func (a *registerUserAmid) checkAccess(c *fiber.Ctx) {
	if a.err != nil {
		return
	}
	err := a.h.tokenS.ValidateRole(c, userfields.ADMIN)
	a.err = err
}
func (a *registerUserAmid) parseBody(c *fiber.Ctx) *usermodel.CreateUserDTO {
	if a.err != nil {
		return nil
	}
	user := new(usermodel.CreateUserDTO)
	err := c.BodyParser(user)
	a.err = amiderrors.NewInternalErrorResponse(err)
	return user
}
func (a *registerUserAmid) validate(user *usermodel.CreateUserDTO) {
	if a.err != nil {
		return
	}
	a.err = user.Validate()
}
func (a *registerUserAmid) insertUser(c *fiber.Ctx, user *usermodel.CreateUserDTO) *usermodel.UserDTO {
	if a.err != nil {
		return nil
	}
	usr, err := a.h.userS.CreateUser(c.Context(), user)
	a.err = err
	return usr
}

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
//	@Router			/users/register [post]
func (h *UserHandler) RegisterUser(c *fiber.Ctx) error {
	a := registerUserAmid{h: h}
	a.checkAccess(c)
	user := a.parseBody(c)
	a.validate(user)
	userDto := a.insertUser(c, user)
	if a.err != nil {
		return a.err.SendWithFiber(c)
	}
	return c.Status(http.StatusCreated).JSON(userDto)
}
