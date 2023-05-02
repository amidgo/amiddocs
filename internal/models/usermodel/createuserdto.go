package usermodel

import (
	"fmt"
	"math/rand"

	"github.com/amidgo/amiddocs/internal/errorutils/usererror"
	"github.com/amidgo/amiddocs/internal/models/usermodel/userfields"

	e "github.com/amidgo/amiddocs/pkg/amiderrors"
	"github.com/amidgo/amiddocs/pkg/validate"
)

type CreateUserDTO struct {
	Name       userfields.Name       `json:"name" db:"name"`
	Surname    userfields.Surname    `json:"surname" db:"surname"`
	FatherName userfields.FatherName `json:"fatherName" db:"father_name"`
	Email      userfields.Email      `json:"email" db:"email"`
	Roles      []userfields.Role     `json:"roles" db:"roles"`
}

func NewCreateUserDTO(
	name userfields.Name,
	surname userfields.Surname,
	fatherName userfields.FatherName,
	email userfields.Email,
	role []userfields.Role,
) *CreateUserDTO {
	return &CreateUserDTO{name, surname, fatherName, email, role}
}

func (dto *CreateUserDTO) GenerateLoginAndPassword() (userfields.Login, userfields.Password, error) {
	if len(dto.Name) == 0 || len(dto.Surname) == 0 {
		return "", "", e.EmptyValues(userfields.NAME_FIELD_NAME)
	}
	randInt := rand.Intn(1000)
	fsym := ""
	if dto.FatherName != "" {
		fsym = string([]rune(dto.FatherName)[0])
	}
	login := string(dto.Surname) + string([]rune(dto.Name)[0]) + fsym + fmt.Sprint(randInt)
	return userfields.Login(login), userfields.Password(login), nil
}

func (dto *CreateUserDTO) ValidatableVariables() []validate.Validatable {
	list := []validate.Validatable{dto.Name, dto.Surname, dto.FatherName, dto.Email}
	for _, r := range dto.Roles {
		list = append(list, r)
	}
	return list
}

func (dto *CreateUserDTO) Validate() error {
	if len(dto.Roles) == 0 {
		return usererror.EMPTY_ROLES
	}
	err := validate.ValidateFields(dto)
	if err != nil {
		return err
	}
	return nil
}
