package cli

import (
	"flag"
	"log"
	"os"
	"path"

	"github.com/amidgo/amiddocs/internal/config"
	"github.com/amidgo/amiddocs/internal/csvimport"
	"github.com/amidgo/amiddocs/internal/database/postgres/depstorage"
	"github.com/amidgo/amiddocs/internal/database/postgres/groupstorage"
	"github.com/amidgo/amiddocs/internal/database/postgres/stdocstorage"
	"github.com/amidgo/amiddocs/internal/database/postgres/studentstorage"
	"github.com/amidgo/amiddocs/internal/database/postgres/userstorage"
	"github.com/amidgo/amiddocs/internal/domain/departmentservice"
	"github.com/amidgo/amiddocs/internal/domain/groupservice"
	"github.com/amidgo/amiddocs/internal/domain/studentservice"
	"github.com/amidgo/amiddocs/internal/encrypt"
	"github.com/amidgo/amiddocs/internal/filestorage"
	"github.com/amidgo/amiddocs/internal/models/depmodel"
	"github.com/amidgo/amiddocs/pkg/postgres"
)

func Run() {

	config := config.LoadConfig()

	p, err := postgres.New(config.DatabaseURL())
	if err != nil {
		log.Fatal(err)
	}

	depFileStorage := filestorage.New(path.Join(os.Getenv("PWD"), config.FileStorage.DepFileStorage))

	depRep := depstorage.New(p)
	groupRep := groupstorage.New(p)
	userRep := userstorage.New(p)
	studentRep := studentstorage.New(p)
	stDocRep := stdocstorage.New(p)

	encrypter := encrypt.New(10)

	depSer := departmentservice.New(depRep, depRep, depFileStorage)
	groupSer := groupservice.New(groupRep, depRep, groupRep)
	studentSer := studentservice.New(groupRep, stDocRep, userRep, studentRep, studentRep, depRep, encrypter)
	_ = depSer

	var depFile, groupFile, studentFile string
	flag.StringVar(&depFile, "dep", "", "path to csv file which contains department info")
	flag.StringVar(&groupFile, "group", "", "path to csv file which contains group info")
	flag.StringVar(&studentFile, "student", "", "path to csv file which contains student info")

	csvimport.CreateEntityFromCsv[depmodel.CreateDepartmentDTO](depFile, nil)
	csvimport.CreateEntityFromCsv(groupFile, groupSer.CreateManyGroups)
	csvimport.CreateEntityFromCsv(studentFile, studentSer.CreateManyStudents)
}
