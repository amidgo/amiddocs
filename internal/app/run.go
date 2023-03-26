package app

import (
	"fmt"
	"log"
	"os"

	"github.com/amidgo/amiddocs/internal/database/postgres/userstorage"
	"github.com/amidgo/amiddocs/internal/domain/jwttoken"
	"github.com/amidgo/amiddocs/internal/domain/userservice"
	"github.com/amidgo/amiddocs/internal/jwtgen"
	"github.com/amidgo/amiddocs/internal/models/usermodel/userfields"
	"github.com/amidgo/amiddocs/internal/transport/http/handlers/userhandlers"
	"github.com/amidgo/amiddocs/internal/transport/http/routing/userrouting"
	"github.com/amidgo/amiddocs/pkg/middleware"
	"github.com/amidgo/amiddocs/pkg/postgres"
	"github.com/gofiber/fiber/v2"
)

const (
	PORTENV     = "PORT"
	HOSTENV     = "HOST"
	DATABASEURL = "DATABASEURL"
)

func Run() {

	app := fiber.New()
	middleware.SetUpMiddleWare(app)
	pg, err := postgres.New(os.Getenv(DATABASEURL))
	if err != nil {
		log.Fatal(err)
	}
	jwtGen := new(jwtgen.RsJWT)
	tokenFabric := jwttoken.NewTokenFabric(jwtGen)
	ms, _ := tokenFabric.CreateUserAccessToken(1, []userfields.UserRole{userfields.ADMIN})
	fmt.Println(ms)
	userRepo := userstorage.New(pg)

	userService := userservice.New(userRepo, tokenFabric)

	userHandler := userhandlers.New(userService, tokenFabric)
	userrouting.SetUp(app, userHandler)

	err = app.Listen(os.Getenv(HOSTENV) + ":" + os.Getenv(PORTENV))
	log.Fatal(err)
}
