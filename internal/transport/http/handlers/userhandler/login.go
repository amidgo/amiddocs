package userhandler

import (
	"net/http"

	"github.com/amidgo/amiddocs/internal/models/usermodel"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
	"github.com/gofiber/fiber/v2"
)

// Login godoc
//
//	@Summary		Login
//	@Description	login by login and password
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			loginForm	body		usermodel.LoginForm	true	"login form"
//	@Success		200			{object}	tokenmodel.TokenResponse
//	@Failure		400			{object}	amiderrors.ErrorResponse
//	@Failure		500			{object}	amiderrors.ErrorResponse
//	@Security		Token
//	@Router			/users/login [post]
func (h *userHandler) Login(c *fiber.Ctx) error {
	loginForm := new(usermodel.LoginForm)
	err := c.BodyParser(loginForm)
	if err != nil {
		return amiderrors.Wrap(err, amiderrors.NewCause("parse body", "Login", _PROVIDER))
	}
	err = loginForm.Validate()
	if err != nil {
		return err
	}
	tokenBody, err := h.userS.Login(c.UserContext(), loginForm)
	if err != nil {
		return err
	}
	return c.Status(http.StatusOK).JSON(tokenBody)
}
