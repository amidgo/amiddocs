package userhandlers

import (
	"context"
	"net/http"

	"github.com/amidgo/amiddocs/internal/models/tokenmodel"
	"github.com/amidgo/amiddocs/internal/models/usermodel"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
	"github.com/gofiber/fiber/v2"
)

type loginAmid struct {
	h   *UserHandler
	err *amiderrors.ErrorResponse
}

func (a *loginAmid) parseBody(c *fiber.Ctx) *usermodel.LoginForm {
	if a.err != nil {
		return nil
	}
	loginForm := new(usermodel.LoginForm)
	err := c.BodyParser(loginForm)
	a.err = amiderrors.NewInternalErrorResponse(err)
	return loginForm
}

func (a *loginAmid) validate(loginForm *usermodel.LoginForm) {
	if a.err != nil {
		return
	}
	a.err = loginForm.Validate()
}

func (a *loginAmid) login(ctx context.Context, loginForm usermodel.LoginForm) *tokenmodel.TokenResponse {
	if a.err != nil {
		return nil
	}
	tokenR, err := a.h.userS.Login(ctx, loginForm)
	a.err = err
	return tokenR
}

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
//	@Router			/users/login [post]
func (h *UserHandler) Login(c *fiber.Ctx) error {
	amid := loginAmid{h: h, err: nil}
	loginForm := amid.parseBody(c)
	amid.validate(loginForm)
	tokenR := amid.login(c.Context(), *loginForm)
	if amid.err != nil {
		return amid.err.SendWithFiber(c)
	}
	return c.Status(http.StatusOK).JSON(tokenR)
}
