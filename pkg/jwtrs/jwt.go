package jwtrs

import (
	"crypto/rsa"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/amidgo/amiddocs/pkg/amiderrors"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
)

func loadKey(pempath string) *rsa.PrivateKey {
	b, err := os.ReadFile(pempath)
	if err != nil {
		log.Fatal(err)
	}
	key, err := jwt.ParseRSAPrivateKeyFromPEM(b)
	if err != nil {
		log.Fatal(err)
	}
	return key
}

type RsJWT struct {
	pempath string
	key     *rsa.PrivateKey
}

func New(pempath string) *RsJWT {
	r := RsJWT{pempath: pempath}
	r.key = loadKey(pempath)
	return &r
}

func (r *RsJWT) CreateToken(claims jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	t, err := token.SignedString(r.key)
	if err != nil {
		return "", err
	}
	return t, nil
}

var TOKEN_EXPIRED = amiderrors.NewErrorResponse("Время действия токена вышло", http.StatusUnauthorized, "token_expired")

func errorHandler(c *fiber.Ctx, err error) error {
	switch err.(type) {
	case *amiderrors.ErrorResponse:
		return err.(*amiderrors.ErrorResponse)
	case *jwt.ValidationError:
		jwtErr := err.(*jwt.ValidationError)
		if amidErr, ok := jwtErr.Inner.(*amiderrors.ErrorResponse); ok {
			return amidErr
		}
		return amiderrors.NewErrorResponse(err.Error(), http.StatusBadRequest, "token_error")
	default:
		return amiderrors.NewErrorResponse(err.Error(), http.StatusBadRequest, "token_error")
	}
}

func (r *RsJWT) keyFunc(t *jwt.Token) (interface{}, error) {
	if t.Method.Alg() != jwtware.RS256 {
		return nil, amiderrors.NewErrorResponse("wrong token", http.StatusUnauthorized, "wrong_token_sign_method")
	}
	mclaims, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		return nil, amiderrors.NewErrorResponse("token_error", http.StatusUnauthorized, "token_error")
	}
	unix := mclaims["exp"].(float64)
	tm := time.Unix(int64(unix), 0)
	if tm.Before(time.Now()) {
		return nil, TOKEN_EXPIRED
	}
	return r.key.Public(), nil
}

func (r *RsJWT) Ware(opts ...Option) func(c *fiber.Ctx) error {
	c := &jwtware.Config{
		SigningMethod: jwtware.RS256,
		SigningKey:    r.key.Public(),
		ErrorHandler:  errorHandler,
		KeyFunc:       r.keyFunc,
	}
	for _, op := range opts {
		op(c)
	}
	return jwtware.New(*c)
}

type Option func(*jwtware.Config)

type JwtWare func(opts ...Option) func(*fiber.Ctx) error

func KeyFuncOption(kf jwt.Keyfunc) Option {
	return func(c *jwtware.Config) {
		c.KeyFunc = kf
	}
}

func ContextKeyOption(key string) Option {
	return func(c *jwtware.Config) {
		c.ContextKey = key
	}
}

func ErrorHandlerOption(errHandler func(c *fiber.Ctx, err error)) Option {
	return func(c *jwtware.Config) {
		c.ErrorHandler = errorHandler
	}
}
