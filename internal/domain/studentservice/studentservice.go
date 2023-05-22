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

type groupProvider interface {
	GroupByName(ctx context.Context, name groupfields.Name) (*groupmodel.GroupDTO, error)
}
type studentDocProvider interface {
	DocumentByOrderNumber(ctx context.Context, orderNumber stdocfields.OrderNumber) (*stdocmodel.StudentDocumentDTO, error)
	DocumentByDocNumber(ctx context.Context, docNumber stdocfields.DocNumber) (*stdocmodel.StudentDocumentDTO, error)
}
type userProvider interface {
	UserByEmail(ctx context.Context, email userfields.Email) (*usermodel.UserDTO, error)
}
type studentProvider interface {
	StudentById(ctx context.Context, id uint64) (*studentmodel.StudentDTO, error)
	AllStudents(ctx context.Context) ([]*studentmodel.StudentDTO, error)
	InsertStudent(ctx context.Context, student *studentmodel.StudentDTO) (*studentmodel.StudentDTO, error)
}
type depProvider interface {
	StudyDepartment(ctx context.Context, studyDepId uint64) (*depmodel.DepartmentDTO, error)
}

type encrypter interface {
	Hash(input string) (string, error)
	Verify(hashPassword string, password string) bool
}

type studentService struct {
	groupProv      groupProvider
	studentDocProv studentDocProvider
	userProv       userProvider
	studentRep     studentProvider
	depProv        depProvider
	encrypter      encrypter
}

func New(
	groupRep groupProvider,
	studentDocRep studentDocProvider,
	userRep userProvider,
	studentRep studentProvider,
	depRep depProvider,
	enctypter encrypter,
) *studentService {
	return &studentService{
		groupProv:      groupRep,
		studentDocProv: studentDocRep,
		userProv:       userRep,
		studentRep:     studentRep,
		depProv:        depRep,
		encrypter:      enctypter,
	}
}
