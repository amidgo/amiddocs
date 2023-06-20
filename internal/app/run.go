package app

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/amidgo/amiddocs/internal/config"
	"github.com/amidgo/amiddocs/internal/database/postgres/depstorage"
	"github.com/amidgo/amiddocs/internal/database/postgres/doctempstorage"
	"github.com/amidgo/amiddocs/internal/database/postgres/doctypestorage"
	"github.com/amidgo/amiddocs/internal/database/postgres/groupstorage"
	"github.com/amidgo/amiddocs/internal/database/postgres/reqstorage"
	"github.com/amidgo/amiddocs/internal/database/postgres/rtokenstorage"
	"github.com/amidgo/amiddocs/internal/database/postgres/stdocstorage"
	"github.com/amidgo/amiddocs/internal/database/postgres/studentstorage"
	"github.com/amidgo/amiddocs/internal/database/postgres/userstorage"
	"github.com/amidgo/amiddocs/internal/docxreplacer"
	"github.com/amidgo/amiddocs/internal/domain/departmentservice"
	"github.com/amidgo/amiddocs/internal/domain/docgenerator"
	"github.com/amidgo/amiddocs/internal/domain/doctempservice"
	"github.com/amidgo/amiddocs/internal/domain/groupservice"
	"github.com/amidgo/amiddocs/internal/domain/reqservice"
	"github.com/amidgo/amiddocs/internal/domain/stdocservice"
	"github.com/amidgo/amiddocs/internal/domain/studentservice"
	"github.com/amidgo/amiddocs/internal/domain/userservice"
	"github.com/amidgo/amiddocs/internal/encrypt"
	"github.com/amidgo/amiddocs/internal/fiberconfig"
	"github.com/amidgo/amiddocs/internal/filestorage"
	"github.com/amidgo/amiddocs/internal/jwttoken"
	"github.com/amidgo/amiddocs/internal/swagger"
	"github.com/amidgo/amiddocs/internal/transport/http/handlers/departmenthandler"
	"github.com/amidgo/amiddocs/internal/transport/http/handlers/doctemphandler"
	"github.com/amidgo/amiddocs/internal/transport/http/handlers/grouphandler"
	"github.com/amidgo/amiddocs/internal/transport/http/handlers/reqhandler"
	"github.com/amidgo/amiddocs/internal/transport/http/handlers/studenthandler"
	"github.com/amidgo/amiddocs/internal/transport/http/handlers/userhandler"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
	"github.com/amidgo/amiddocs/pkg/jwtrs"
	"github.com/amidgo/amiddocs/pkg/middleware"
	"github.com/amidgo/amiddocs/pkg/postgres"
	"github.com/gofiber/fiber/v2"
)

func Run() {
	// set unix time
	os.Setenv("TZ", time.UTC.String())
	// load config from config/config.yaml
	config := config.LoadConfig()
	// load errors from config/error/ru.yaml file
	amiderrors.Init("config/error/ru.yaml")
	// create base app fiber instance
	app := fiber.New(
		fiberconfig.Config(),
	)
	// set up config, fiber logger, cors
	middleware.SetUpMiddleWare(app)

	// create postgres instance
	pg, err := postgres.New(config.DatabaseURL())
	if err != nil {
		log.Fatal(err)
	}

	// filestorage module
	os.Mkdir(config.FileStorage.Root, os.ModePerm)
	depFileStorage := filestorage.NewDepFileStorage(app, config.FileStorage.DepFileStorage)

	// create bearer token middleware
	jwtGen := jwtrs.New(config.Jwt.Pempath)
	ware := jwtGen.Ware(jwtrs.ContextKeyOption(config.Jwt.Name))
	tokenMaster := jwttoken.NewTokenMaster(jwtGen, config, ware)

	// create encrypter
	encrypter := encrypt.New(10)
	docxReplacer := docxreplacer.New()

	// initialize the repos
	userRepo := userstorage.New(pg)
	groupRepo := groupstorage.New(pg)
	depRepo := depstorage.New(pg)
	stDocRepo := stdocstorage.New(pg)
	studentRepo := studentstorage.New(pg)
	requestRepo := reqstorage.New(pg)
	docTypeRepo := doctypestorage.New(pg)
	docTempRepo := doctempstorage.New(pg)
	rTokenRepo := rtokenstorage.New(time.Second*time.Duration(config.Jwt.RefreshTokenTime), pg)

	// initialize services
	groupService := groupservice.New(groupRepo, depRepo, groupRepo)
	userService := userservice.New(userRepo, tokenMaster, userRepo, encrypter, rTokenRepo)
	depService := departmentservice.New(depRepo, depRepo, depFileStorage)
	stDocService := stdocservice.New(stDocRepo)
	studentService := studentservice.New(groupRepo, stDocRepo, userRepo, studentRepo, studentRepo, depRepo, encrypter)
	reqService := reqservice.New(depRepo, requestRepo, requestRepo, docTypeRepo)
	doctempService := doctempservice.New(depRepo, docTypeRepo, docTempRepo, docTempRepo)
	docGenService := docgenerator.New(docxReplacer, studentRepo, userRepo, requestRepo, docTempRepo)
	_ = stDocService

	// set up swagger route
	swagger.SetUp(app)

	// add client token check
	app.Use(fiberconfig.ClientTokenHandler(config.Server.Token))

	// setUp handlers with routing
	grouphandler.SetUp(app, tokenMaster, groupService, groupRepo)
	userhandler.SetUp(app, tokenMaster, userService, userRepo)
	departmenthandler.SetUp(app, tokenMaster, depService, depRepo)
	studenthandler.SetUp(app, tokenMaster, studentService, studentRepo)
	reqhandler.SetUp(app, tokenMaster, reqService, tokenMaster, requestRepo, docGenService)
	doctemphandler.SetUp(app, tokenMaster, doctempService, docTempRepo)
	//set up swagger

	// start the server application
	addr := fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)

	// app start
	app.Listen(addr)
}
