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
	Name               userfields.Name         `json:"name" csv:"Имя"`
	Surname            userfields.Surname      `json:"surname" csv:"Фамилия"`
	FatherName         userfields.FatherName   `json:"fatherName" csv:"Отчество"`
	Email              userfields.Email        `json:"email" csv:"-"`
	DocNumber          stdocfields.DocNumber   `json:"docNumber" csv:"Номер Студ Билета"`
	OrderNumber        stdocfields.OrderNumber `json:"orderNumber" csv:"Номер Приказа о зачислении"`
	OrderDate          amidtime.Date           `json:"orderDate" csv:"Дата Приказа о Зачислении"`
	EducationStartDate amidtime.Date           `json:"studyStartDate" csv:"Время Начала Обучения"`
	GroupName          groupfields.Name        `json:"groupName" csv:"Группа"`
}

func NewCreateStudentDTO(
	name userfields.Name,
	surname userfields.Surname,
	fatherName userfields.FatherName,
	email userfields.Email,
	docNumber stdocfields.DocNumber,
	orderNumber stdocfields.OrderNumber,
	orderDate amidtime.Date,
	educationStartDate amidtime.Date,
	groupName groupfields.Name,
) *CreateStudentDTO {
	return &CreateStudentDTO{
		Name:               name,
		Surname:            surname,
		FatherName:         fatherName,
		Email:              email,
		DocNumber:          docNumber,
		OrderNumber:        orderNumber,
		OrderDate:          orderDate,
		EducationStartDate: educationStartDate,
		GroupName:          groupName,
	}
}

func (c *CreateStudentDTO) ValidatableVariables() []validate.Validatable {
	vars := make([]validate.Validatable, 0)
	vars = append(vars, c.Name, c.Surname, c.FatherName, c.Email)
	vars = append(vars, c.DocNumber, c.OrderNumber)
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
			student.Name,
			student.Surname,
			student.FatherName,
			student.Email,
			[]userfields.Role{userfields.STUDENT},
		),
		stdocmodel.NewStudentDocumentDTO(0, 0,
			student.DocNumber,
			student.OrderNumber,
			student.OrderDate,
			student.EducationStartDate,
		),
		group,
		department,
	)
}
