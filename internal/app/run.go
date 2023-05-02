package app

import (
	"fmt"
	"log"
	"os"

	"github.com/amidgo/amiddocs/internal/database/postgres/depstorage"
	"github.com/amidgo/amiddocs/internal/database/postgres/doctempstorage"
	"github.com/amidgo/amiddocs/internal/database/postgres/doctypestorage"
	"github.com/amidgo/amiddocs/internal/database/postgres/groupstorage"
	"github.com/amidgo/amiddocs/internal/database/postgres/reqstorage"
	"github.com/amidgo/amiddocs/internal/database/postgres/stdocstorage"
	"github.com/amidgo/amiddocs/internal/database/postgres/studentstorage"
	"github.com/amidgo/amiddocs/internal/database/postgres/userstorage"
	"github.com/amidgo/amiddocs/internal/domain/departmentservice"
	"github.com/amidgo/amiddocs/internal/domain/doctempservice"
	"github.com/amidgo/amiddocs/internal/domain/groupservice"
	"github.com/amidgo/amiddocs/internal/domain/reqservice"
	"github.com/amidgo/amiddocs/internal/domain/stdocservice"
	"github.com/amidgo/amiddocs/internal/domain/studentservice"
	"github.com/amidgo/amiddocs/internal/domain/userservice"
	"github.com/amidgo/amiddocs/internal/encrypt"
	"github.com/amidgo/amiddocs/internal/fiberconfig"
	"github.com/amidgo/amiddocs/internal/jwttoken"
	"github.com/amidgo/amiddocs/internal/swagger"
	"github.com/amidgo/amiddocs/internal/transport/http/handlers/departmenthandler"
	"github.com/amidgo/amiddocs/internal/transport/http/handlers/doctemphandler"
	"github.com/amidgo/amiddocs/internal/transport/http/handlers/grouphandler"
	"github.com/amidgo/amiddocs/internal/transport/http/handlers/reqhandler"
	"github.com/amidgo/amiddocs/internal/transport/http/handlers/studenthandler"
	"github.com/amidgo/amiddocs/internal/transport/http/handlers/userhandler"
	"github.com/amidgo/amiddocs/pkg/jwtrs"
	"github.com/amidgo/amiddocs/pkg/middleware"
	"github.com/amidgo/amiddocs/pkg/postgres"
	"github.com/gofiber/fiber/v2"
)

const (
	PORTENV     = "PORT"
	HOSTENV     = "HOST"
	DATABASEURL = "DATABASEURL"
	PEMPATH     = "PEMPATH"
	FILESPATH   = "FILESPATH"
)

func Run() {

	// create base app fiber instance
	app := fiber.New(
		fiberconfig.Config(),
	)

	// set up config, fiber logger, cors
	middleware.SetUpMiddleWare(app)

	// create postgres instance
	pg, err := postgres.New(os.Getenv(DATABASEURL))
	if err != nil {
		log.Fatal(err)
	}
	// create bearer token middleware
	jwtGen := jwtrs.New(os.Getenv(PEMPATH))
	tokenFabric := jwttoken.NewTokenFabric(jwtGen)

	// create encrypter
	encrypter := encrypt.New(10)
	// docxReplacer := docxreplacer.New()

	// initialize the repos
	userRepo := userstorage.New(pg)
	groupRepo := groupstorage.New(pg)
	depRepo := depstorage.New(pg)
	stDocRepo := stdocstorage.New(pg)
	studentRepo := studentstorage.New(pg)
	requestRepo := reqstorage.New(pg)
	docTypeRepo := doctypestorage.New(pg)
	docTempRepo := doctempstorage.New(pg)

	// initialize services
	groupService := groupservice.New(groupRepo, depRepo, groupRepo)
	userService := userservice.New(userRepo, tokenFabric, userRepo, encrypter)
	depService := departmentservice.New(depRepo, depRepo)
	stDocService := stdocservice.New(stDocRepo)
	studentService := studentservice.New(groupRepo, stDocRepo, userRepo, studentRepo, encrypter)
	reqService := reqservice.New(depRepo, requestRepo, requestRepo, docTypeRepo)
	doctempService := doctempservice.New(depRepo, docTypeRepo, docTempRepo, docTempRepo)
	_ = stDocService

	fmt.Println(tokenFabric.TokenWithWrongExp())

	ware := jwtGen.Ware()
	// setUp handlers with routing
	grouphandler.SetUp(app, ware, groupService, groupRepo)
	userhandler.SetUp(app, ware, userService, userRepo)
	departmenthandler.SetUp(app, ware, depService, depRepo)
	studenthandler.SetUp(app, ware, studentService, studentRepo)
	reqhandler.SetUp(app, ware, reqService, tokenFabric, requestRepo)
	doctemphandler.SetUp(app, ware, doctempService)
	//set up swagger
	swagger.SetUp(app)

	// start the server application
	err = app.Listen(os.Getenv(HOSTENV) + ":" + os.Getenv(PORTENV))
	log.Fatal(err)
}
