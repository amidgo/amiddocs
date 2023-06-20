package userhandler

import (
	"net/http"

	"github.com/amidgo/amiddocs/pkg/amiderrors"
	"github.com/gofiber/fiber/v2"
)

// GetUserById godoc
//
//	@Summary		Return User
//	@Description	return user by id from path
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	usermodel.UserDTO
//	@Failure		400	{object}	amiderrors.ErrorResponse
//	@Failure		404	{object}	amiderrors.ErrorResponse
//	@Failure		500	{object}	amiderrors.ErrorResponse
//	@Security		Bearer
//	@Security		Token
//	@Router			/users/info [get]
func (s *userHandler) UserInfo(c *fiber.Ctx) error {
	id, err := s.jwt.UserID(c)
	if err != nil {
		return amiderrors.Wrap(err, amiderrors.NewCause("get user id from claims", "UserInfo", _PROVIDER))
	}
	u, err := s.userP.UserById(c.UserContext(), id)
	if err != nil {
		return err
	}
	return c.Status(http.StatusOK).JSON(u)
}
