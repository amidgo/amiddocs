package jwttoken

import (
	"errors"
	"time"

	tokenerrorutils "github.com/amidgo/amiddocs/internal/errorutils/tokenerror"
	"github.com/amidgo/amiddocs/internal/models/usermodel/userfields"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

const _PROVIDER = "internal/domain/jwttoken/TokenFabric"

type jwtGeneratorInterface interface {
	CreateToken(jwt.MapClaims) (string, error)
}

type TokenFabric struct {
	j jwtGeneratorInterface
}

func NewTokenFabric(j jwtGeneratorInterface) *TokenFabric {
	return &TokenFabric{j: j}
}

func (tf *TokenFabric) CreateAccessToken(userid uint64, roles []userfields.Role) (string, error) {
	claims := amidClaims{
		jwt.MapClaims{
			_USER_ID:    userid,
			_ROLES:      roles,
			_EXPIRATION: time.Now().Add(time.Hour * 24 * 30).Unix(),
		},
	}
	token, err := tf.j.CreateToken(claims.MapClaims)
	if err != nil {
		return "", amiderrors.NewInternalErrorResponse(err, amiderrors.NewCause("create token", "CreateAccessToken", _PROVIDER))
	}
	return token, nil
}

func (tf *TokenFabric) TokenWithWrongExp() (string, error) {
	claims := amidClaims{
		jwt.MapClaims{
			_USER_ID:    0,
			_ROLES:      []userfields.Role{userfields.ADMIN},
			_EXPIRATION: time.Now().Unix(),
		},
	}
	token, err := tf.j.CreateToken(claims.MapClaims)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (tf *TokenFabric) Claims(c *fiber.Ctx) *amidClaims {
	return claims(c)
}

func (tf *TokenFabric) ValidateRole(c *fiber.Ctx, role userfields.Role) error {
	return validateRole(c, role)
}

func (tf *TokenFabric) UserID(c *fiber.Ctx) (uint64, error) {
	claims := tf.Claims(c)
	return claims.UserID()
}

func (tf *TokenFabric) UserRoles(c *fiber.Ctx) ([]userfields.Role, error) {
	claims := tf.Claims(c)
	roles := claims.UserRoles()
	if len(roles) == 0 {
		return nil, amiderrors.NewInternalErrorResponse(errors.New("empty roles"), amiderrors.NewCause("get user roles", "UserRoles", _PROVIDER))
	}
	return roles, nil
}

func claims(c *fiber.Ctx) *amidClaims {
	token := c.Locals("user").(*jwt.Token)
	mpCl := token.Claims.(jwt.MapClaims)
	return &amidClaims{mpCl}
}

func validateRole(c *fiber.Ctx, role userfields.Role) error {
	claims := claims(c)
	clRole := claims.UserRoles()
	for _, r := range clRole {
		if r == role {
			return nil
		}
	}
	return tokenerrorutils.FORBIDDEN
}
