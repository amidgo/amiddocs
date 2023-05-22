package studentmodel

import (
	"github.com/amidgo/amiddocs/internal/models/depmodel"
	"github.com/amidgo/amiddocs/internal/models/groupmodel"
	"github.com/amidgo/amiddocs/internal/models/groupmodel/groupfields"
	"github.com/amidgo/amiddocs/internal/models/stdocmodel"
	"github.com/amidgo/amiddocs/internal/models/stdocmodel/stdocfields"
	"github.com/amidgo/amiddocs/internal/models/usermodel"
	"github.com/amidgo/amiddocs/internal/models/usermodel/userfields"
	"github.com/amidgo/amiddocs/pkg/amidtime"
	"github.com/amidgo/amiddocs/pkg/validate"
)

type CreateStudentDTO struct {
	User *struct {
		Name       userfields.Name       `json:"name" db:"name"`
		Surname    userfields.Surname    `json:"surname" db:"surname"`
		FatherName userfields.FatherName `json:"fatherName" db:"father_name"`
		Email      userfields.Email      `json:"email" db:"email"`
	} `json:"user"`
	Document *struct {
		DocNumber          stdocfields.DocNumber   `json:"docNumber" db:"doc_number"`
		OrderNumber        stdocfields.OrderNumber `json:"orderNumber" db:"order_number"`
		OrderDate          amidtime.Date           `json:"orderDate" db:"order_date"`
		EducationStartDate amidtime.Date           `json:"studyStartDate" db:"education_start_date"`
	} `json:"document"`
	GroupName groupfields.Name `json:"groupName"`
}

func (c *CreateStudentDTO) ValidatableVariables() []validate.Validatable {
	vars := make([]validate.Validatable, 0)
	vars = append(vars, c.User.Name, c.User.Surname, c.User.FatherName, c.User.Email)
	vars = append(vars, c.Document.DocNumber, c.Document.OrderNumber)
	vars = append(vars, c.GroupName)
	return vars
}

func (student *CreateStudentDTO) StudentDTO(
	login userfields.Login,
	password userfields.Password,
	group *groupmodel.GroupDTO,
	department *depmodel.DepartmentDTO,
) *StudentDTO {
	return NewStudentDTO(
		0,
		usermodel.NewUserDTO(
			0,
			login,
			password,
			student.User.Name,
			student.User.Surname,
			student.User.FatherName,
			student.User.Email,
			[]userfields.Role{userfields.STUDENT},
		),
		stdocmodel.NewStudentDocumentDTO(0, 0,
			student.Document.DocNumber,
			student.Document.OrderNumber,
			student.Document.OrderDate,
			student.Document.EducationStartDate,
		),
		group,
		department,
	)
}
