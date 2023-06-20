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

// load key from pempath attribute
// return rsa key or log.Fatal(err)
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

// RS jwt struct
type RsJWT struct {
	pempath string
	key     *rsa.PrivateKey
}

// return *RsJWT
// load key by pem path
func New(pempath string) *RsJWT {
	r := RsJWT{pempath: pempath}
	r.key = loadKey(pempath)
	return &r
}

// create token with jwt
func (r *RsJWT) CreateToken(claims jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	t, err := token.SignedString(r.key)
	if err != nil {
		return "", err
	}
	return t, nil
}

const TOKEN_TYPE = "token"

// token expired error
var TOKEN_EXPIRED = amiderrors.NewException(http.StatusUnauthorized, TOKEN_TYPE, "expired")

// error handler with amiderrors.ErrorResponse support
// if type err is *jwt.ValidationError try convert err.Inner to amiderrors.ErrorResponse
// default amiderrors.NewErrorResponse(err.Error(), http.StatusBadRequest, "token_error")
func errorHandler(c *fiber.Ctx, err error) error {
	switch err := err.(type) {
	case *amiderrors.Exception:
		return err
	case *amiderrors.ErrorResponse:
		return err
	case *jwt.ValidationError:
		jwtErr := err
		if amidErr, ok := jwtErr.Inner.(*amiderrors.Exception); ok {
			return amidErr
		}
		return amiderrors.NewException(http.StatusBadRequest, TOKEN_TYPE, amiderrors.INTERNAL)
	default:
		return amiderrors.NewException(http.StatusBadRequest, TOKEN_TYPE, amiderrors.INTERNAL)
	}
}

// default keyfunc
// check alg, exp properties
func (r *RsJWT) keyFunc(t *jwt.Token) (interface{}, error) {
	if t.Method.Alg() != jwtware.RS256 {
		return nil, amiderrors.NewException(http.StatusUnauthorized, TOKEN_TYPE, "wrong_sign")
	}
	mclaims, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		return nil, amiderrors.NewException(http.StatusUnauthorized, TOKEN_TYPE, amiderrors.INTERNAL)
	}
	unix := mclaims["exp"].(float64)
	tm := time.Unix(int64(unix), 0)
	if tm.Before(time.Now()) {
		return nil, TOKEN_EXPIRED
	}
	return r.key.Public(), nil
}

// default ware with option support
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

// jwt ware option
type Option func(*jwtware.Config)

// jwt ware type
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
