package jwths

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
)

type HsJWT struct {
	jwtsecret string
}

func (h *HsJWT) CreateToken(claims jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString(h.jwtsecret)
	if err != nil {
		return "", err
	}
	return t, nil
}

type Option func(w *jwtware.Config)

func (h *HsJWT) Ware(opts ...Option) func(*fiber.Ctx) error {
	c := &jwtware.Config{
		SigningMethod: "HS256",
		SigningKey:    h.jwtsecret,
	}
	for _, op := range opts {
		op(c)
	}
	return jwtware.New(*c)
}
