package studentmodel

import (
	"github.com/amidgo/amiddocs/internal/models/depmodel"
	"github.com/amidgo/amiddocs/internal/models/groupmodel"
	"github.com/amidgo/amiddocs/internal/models/stdocmodel"
	"github.com/amidgo/amiddocs/internal/models/usermodel"
	"github.com/amidgo/amiddocs/pkg/validate"
)

type StudentDTO struct {
	ID         uint64                         `json:"id" db:"id"`
	User       *usermodel.UserDTO             `json:"user"`
	Document   *stdocmodel.StudentDocumentDTO `json:"document"`
	Group      *groupmodel.GroupDTO           `json:"group"`
	Department *depmodel.DepartmentDTO        `json:"department"`
}

func (s *StudentDTO) ValidatableVariables() []validate.Validatable {
	vars := make([]validate.Validatable, 0)
	vars = append(vars, s.User.ValidatableVariables()...)
	vars = append(vars, s.Group.ValidatableVariables()...)
	vars = append(vars, s.Document.ValidatableVariables()...)
	return vars
}

func NewStudentDTO(
	id uint64,
	user *usermodel.UserDTO,
	document *stdocmodel.StudentDocumentDTO,
	group *groupmodel.GroupDTO,
	department *depmodel.DepartmentDTO,
) *StudentDTO {
	return &StudentDTO{
		ID:         id,
		User:       user,
		Document:   document,
		Group:      group,
		Department: department,
	}
}
