package jwttoken

import (
	"time"

	tokenerrorutils "github.com/amidgo/amiddocs/internal/errorutils/token_error_utils"
	"github.com/amidgo/amiddocs/internal/models/usermodel/userfields"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type jwtGeneratorInterface interface {
	CreateToken(jwt.MapClaims) (string, error)
}

type TokenFabric struct {
	j jwtGeneratorInterface
}

func NewTokenFabric(j jwtGeneratorInterface) *TokenFabric {
	return &TokenFabric{j: j}
}

func (tf *TokenFabric) CreateUserAccessToken(userid uint64, roles []userfields.UserRole) (string, *amiderrors.ErrorResponse) {
	claims := amidClaims{
		jwt.MapClaims{
			_USER_ID:    userid,
			_ROLES:      roles,
			_EXPIRATION: time.Now().Add(time.Hour * 24 * 30).Unix(),
		},
	}
	token, err := tf.j.CreateToken(claims.MapClaims)
	if err != nil {
		return "", amiderrors.NewInternalErrorResponse(err)
	}
	return token, nil
}

func (tf *TokenFabric) GetClaims(c *fiber.Ctx) *amidClaims {
	token := c.Locals("user").(*jwt.Token)
	mpCl := token.Claims.(jwt.MapClaims)
	return &amidClaims{mpCl}
}

func (tf *TokenFabric) ValidateRole(c *fiber.Ctx, role userfields.UserRole) *amiderrors.ErrorResponse {
	claims := tf.GetClaims(c)
	clRole := claims.GetUserRoles()
	for _, r := range clRole {
		if r == role {
			return nil
		}
	}
	return tokenerrorutils.FORBIDDEN
}

func (tf *TokenFabric) GetUserId(c *fiber.Ctx) (uint64, *amiderrors.ErrorResponse) {
	claims := tf.GetClaims(c)
	return claims.GetUserId()
}

func (tf *TokenFabric) GetUserRoles(c *fiber.Ctx) []userfields.UserRole {
	claims := tf.GetClaims(c)
	return claims.GetUserRoles()
}
