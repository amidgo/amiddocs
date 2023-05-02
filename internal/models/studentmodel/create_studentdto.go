package studentmodel

import (
	"github.com/amidgo/amiddocs/internal/models/groupmodel/groupfields"
	"github.com/amidgo/amiddocs/internal/models/stdocmodel"
	"github.com/amidgo/amiddocs/internal/models/usermodel"
	"github.com/amidgo/amiddocs/pkg/validate"
)

type CreateStudentDTO struct {
	User      *usermodel.CreateUserDTO       `json:"user"`
	Document  *stdocmodel.StudentDocumentDTO `json:"document"`
	GroupName groupfields.Name               `json:"groupName"`
}

func (c *CreateStudentDTO) ValidatableVariables() []validate.Validatable {
	vars := make([]validate.Validatable, 0)
	vars = append(vars, c.User.ValidatableVariables()...)
	vars = append(vars, c.Document.ValidatableVariables()...)
	vars = append(vars, c.GroupName)
	return vars
}
