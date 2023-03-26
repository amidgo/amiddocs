package jwtgen

import (
	"crypto/rsa"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
)

var (
	rsaPrivateKey *rsa.PrivateKey = loadKey()
	jwtSecretKey  []byte          = []byte(os.Getenv("JWTSECRET"))
)

const (
	TESTRSAPATH = "./config/private.pem"
	PRODRSAPATH = "./config/private.pem"
)

const (
	JWTNAME_ENV = "JWTNAME"
)

func loadKey() *rsa.PrivateKey {
	b, err := os.ReadFile(TESTRSAPATH)
	if err != nil {
		log.Fatal(err)
	}
	key, err := jwt.ParseRSAPrivateKeyFromPEM(b)
	if err != nil {
		log.Fatal(err)
	}
	return key
}

func RsJwtWare() func(*fiber.Ctx) error {
	return jwtware.New(
		jwtware.Config{
			SigningMethod: "RS256",
			SigningKey:    rsaPrivateKey.Public(),
		},
	)
}

func HsJwtWare() func(*fiber.Ctx) error {
	return jwtware.New(
		jwtware.Config{
			SigningMethod: "HS256",
			SigningKey:    jwtSecretKey,
		},
	)
}

type HsJWT struct{}

func (h *HsJWT) CreateToken(claims jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString(jwtSecretKey)
	if err != nil {
		return "", err
	}
	return t, nil
}

type RsJWT struct{}

func (r *RsJWT) CreateToken(claims jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	t, err := token.SignedString(rsaPrivateKey)
	if err != nil {
		return "", err
	}
	return t, nil
}
