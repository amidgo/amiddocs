package jwttoken

import (
	"errors"
	"time"

	"github.com/amidgo/amiddocs/internal/config"
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

type TokenMaster struct {
	jwtware func(c *fiber.Ctx) error
	config  *config.Config
	j       jwtGeneratorInterface
}

func NewTokenMaster(j jwtGeneratorInterface, config *config.Config, jwtware func(c *fiber.Ctx) error) *TokenMaster {
	return &TokenMaster{j: j, config: config, jwtware: jwtware}
}

func (tm *TokenMaster) Ware() func(c *fiber.Ctx) error {
	return tm.jwtware
}

func (tf *TokenMaster) CreateAccessToken(userid uint64, roles []userfields.Role) (string, error) {
	claims := amidClaims{
		jwt.MapClaims{
			_USER_ID:    userid,
			_ROLES:      roles,
			_EXPIRATION: time.Now().Add(time.Second * time.Duration(tf.config.Jwt.AccessTokenTime)).Unix(),
		},
	}
	token, err := tf.j.CreateToken(claims.MapClaims)
	if err != nil {
		return "", amiderrors.NewInternalErrorResponse(err, amiderrors.NewCause("create token", "CreateAccessToken", _PROVIDER))
	}
	return token, nil
}

func (tf *TokenMaster) TokenWithWrongExp() (string, error) {
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

func (tf *TokenMaster) Claims(c *fiber.Ctx) *amidClaims {
	token := c.Locals(tf.config.Jwt.Name).(*jwt.Token)
	mpCl := token.Claims.(jwt.MapClaims)
	return &amidClaims{mpCl}
}

func (tf *TokenMaster) ValidateRole(c *fiber.Ctx, role userfields.Role) error {
	claims := tf.Claims(c)
	clRole := claims.UserRoles()
	for _, r := range clRole {
		if r == role {
			return nil
		}
	}
	return tokenerrorutils.FORBIDDEN
}

func (tf *TokenMaster) UserID(c *fiber.Ctx) (uint64, error) {
	claims := tf.Claims(c)
	return claims.UserID()
}

func (tf *TokenMaster) UserRoles(c *fiber.Ctx) ([]userfields.Role, error) {
	claims := tf.Claims(c)
	roles := claims.UserRoles()
	if len(roles) == 0 {
		return nil, amiderrors.NewInternalErrorResponse(errors.New("empty roles"), amiderrors.NewCause("get user roles", "UserRoles", _PROVIDER))
	}
	return roles, nil
}
