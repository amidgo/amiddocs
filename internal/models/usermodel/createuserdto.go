package usermodel

import (
	"fmt"
	"math/rand"

	"github.com/amidgo/amiddocs/internal/models/usermodel/userfields"

	e "github.com/amidgo/amiddocs/pkg/amiderrors"
	"github.com/amidgo/amiddocs/pkg/validate"
)

type CreateUserDTO struct {
	Name       userfields.Name       `json:"name" db:"name"`
	Surname    userfields.Surname    `json:"surname" db:"surname"`
	FatherName userfields.FatherName `json:"fatherName" db:"father_name"`
	Email      userfields.Email      `json:"email" db:"email"`
	Roles      []userfields.UserRole `json:"roles" db:"roles"`
}

func NewCreateUserDTO(
	name userfields.Name,
	surname userfields.Surname,
	fatherName userfields.FatherName,
	email userfields.Email,
	role []userfields.UserRole,
) *CreateUserDTO {
	return &CreateUserDTO{name, surname, fatherName, email, role}
}

func (dto *CreateUserDTO) GenerateLoginAndPassword() (userfields.Login, userfields.Password, *e.ErrorResponse) {
	if len(dto.Name) == 0 || len(dto.Surname) == 0 {
		return "", "", e.EmptyValues(userfields.NAME_FIELD_NAME)
	}
	randInt := rand.Intn(1000)
	login := string(dto.Surname) + string(dto.Name[0]) + string(dto.FatherName[0]) + fmt.Sprint(randInt)
	password, err := userfields.Password(login).Hash()
	if err != nil {
		return "", "", err
	}
	return userfields.Login(login), userfields.Password(password), nil
}

func (dto *CreateUserDTO) ValidateVariables() []validate.Validatable {
	list := []validate.Validatable{dto.Name, dto.Surname, dto.FatherName, dto.Email}
	for _, r := range dto.Roles {
		list = append(list, r)
	}
	return list
}

func (dto *CreateUserDTO) Validate() *e.ErrorResponse {
	return validate.ValidateStructVariables(dto.ValidateVariables()...)
}
