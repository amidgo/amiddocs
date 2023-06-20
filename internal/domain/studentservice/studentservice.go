package studentservice

import (
	"context"

	"github.com/amidgo/amiddocs/internal/models/depmodel"
	"github.com/amidgo/amiddocs/internal/models/groupmodel"
	"github.com/amidgo/amiddocs/internal/models/groupmodel/groupfields"
	"github.com/amidgo/amiddocs/internal/models/stdocmodel"
	"github.com/amidgo/amiddocs/internal/models/stdocmodel/stdocfields"
	"github.com/amidgo/amiddocs/internal/models/studentmodel"
	"github.com/amidgo/amiddocs/internal/models/usermodel"
	"github.com/amidgo/amiddocs/internal/models/usermodel/userfields"
)

type GroupProvider interface {
	GroupByName(ctx context.Context, name groupfields.Name) (*groupmodel.GroupDTO, error)
}
type StudentDocProvider interface {
	DocumentByOrderNumber(ctx context.Context, orderNumber stdocfields.OrderNumber) (*stdocmodel.StudentDocumentDTO, error)
	DocumentByDocNumber(ctx context.Context, docNumber stdocfields.DocNumber) (*stdocmodel.StudentDocumentDTO, error)
}
type UserProvider interface {
	UserByEmail(ctx context.Context, email userfields.Email) (*usermodel.UserDTO, error)
}
type StudentProvider interface {
	StudentById(ctx context.Context, id uint64) (*studentmodel.StudentDTO, error)
	AllStudents(ctx context.Context) ([]*studentmodel.StudentDTO, error)
}
type StudentRepository interface {
	InsertStudent(ctx context.Context, student *studentmodel.StudentDTO) (*studentmodel.StudentDTO, error)
	InsertManyStudents(ctx context.Context, students []*studentmodel.StudentDTO) error
}
type DepProvider interface {
	StudyDepartmentById(ctx context.Context, studyDepId uint64) (*depmodel.StudyDepartmentDTO, error)
}

type Encrypter interface {
	Hash(input string) (string, error)
	Verify(hashPassword string, password string) bool
}

const _PROVIDER = "internal/domain/studentservice"

type StudentService struct {
	groupProv      GroupProvider
	studentDocProv StudentDocProvider
	userProv       UserProvider
	studentProv    StudentProvider
	studentRep     StudentRepository
	depProv        DepProvider
	encrypter      Encrypter
}

func New(
	groupRep GroupProvider,
	studentDocRep StudentDocProvider,
	userRep UserProvider,
	studentProv StudentProvider,
	studentRepo StudentRepository,
	depRep DepProvider,
	enctypter Encrypter,
) *StudentService {
	return &StudentService{
		groupProv:      groupRep,
		studentDocProv: studentDocRep,
		userProv:       userRep,
		studentProv:    studentProv,
		studentRep:     studentRepo,
		depProv:        depRep,
		encrypter:      enctypter,
	}
}
