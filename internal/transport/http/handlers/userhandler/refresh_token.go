package userhandler

import (
	"net/http"

	"github.com/amidgo/amiddocs/internal/models/rtokenmodel"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
	"github.com/gofiber/fiber/v2"
)

// RefreshToken godoc
//
//	@Summary		refresh token
//	@Description	refresh token by old token and user id
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			refreshdto	body		rtokenmodel.RefreshDTO	true	"refresh dto with user id and old token"
//	@Success		200			{object}	tokenmodel.TokenResponse
//	@Failure		400			{object}	amiderrors.ErrorResponse
//	@Failure		404			{object}	amiderrors.ErrorResponse
//	@Failure		500			{object}	amiderrors.ErrorResponse
//	@Security		Token
//	@Router			/users/refresh-token [post]
func (h *userHandler) RefreshToken(c *fiber.Ctx) error {
	token := new(rtokenmodel.RefreshDTO)
	err := c.BodyParser(token)
	if err != nil {
		return amiderrors.Wrap(err, amiderrors.NewCause("parse body", "RefreshToken", _PROVIDER))
	}
	rtokenResponse, err := h.userS.RefreshToken(c.UserContext(), token.Token.String(), token.UserId)
	if err != nil {
		return err
	}
	return c.Status(http.StatusOK).JSON(rtokenResponse)
}
