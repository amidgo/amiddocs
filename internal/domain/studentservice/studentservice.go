package studentservice

import (
	"context"

	"github.com/amidgo/amiddocs/internal/models/groupmodel"
	"github.com/amidgo/amiddocs/internal/models/groupmodel/groupfields"
	"github.com/amidgo/amiddocs/internal/models/stdocmodel"
	"github.com/amidgo/amiddocs/internal/models/stdocmodel/stdocfields"
	"github.com/amidgo/amiddocs/internal/models/studentmodel"
	"github.com/amidgo/amiddocs/internal/models/usermodel"
	"github.com/amidgo/amiddocs/internal/models/usermodel/userfields"
)

type groupRep interface {
	GroupByName(ctx context.Context, name groupfields.Name) (*groupmodel.GroupDTO, error)
}
type studentDocRep interface {
	DocumentByOrderNumber(ctx context.Context, orderNumber stdocfields.OrderNumber) (*stdocmodel.StudentDocumentDTO, error)
	DocumentByDocNumber(ctx context.Context, docNumber stdocfields.DocNumber) (*stdocmodel.StudentDocumentDTO, error)
}
type userRep interface {
	UserByEmail(ctx context.Context, email userfields.Email) (*usermodel.UserDTO, error)
}
type studentRep interface {
	StudentById(ctx context.Context, id uint64) (*studentmodel.StudentDTO, error)
	AllStudents(ctx context.Context) ([]*studentmodel.StudentDTO, error)
	InsertStudent(ctx context.Context, student *studentmodel.StudentDTO) (*studentmodel.StudentDTO, error)
}

type encrypter interface {
	Hash(input string) (string, error)
	Verify(hashPassword string, password string) bool
}

type studentService struct {
	groupRep      groupRep
	studentDocRep studentDocRep
	userRep       userRep
	studentRep    studentRep
	encrypter     encrypter
}

func New(groupRep groupRep, studentDocRep studentDocRep, userRep userRep, studentRep studentRep, enctypter encrypter) *studentService {
	return &studentService{groupRep: groupRep, studentDocRep: studentDocRep, userRep: userRep, studentRep: studentRep, encrypter: enctypter}
}
