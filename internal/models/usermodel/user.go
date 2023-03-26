package usermodel

import (
	"github.com/amidgo/amiddocs/internal/models/usermodel/userfields"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
	"github.com/amidgo/amiddocs/pkg/validate"
)

type UserDTO struct {
	ID         uint64                `json:"id"`
	Login      userfields.Login      `json:"login"`
	Password   userfields.Password   `json:"password"`
	Name       userfields.Name       `json:"name"`
	Surname    userfields.Surname    `json:"surname"`
	FatherName userfields.FatherName `json:"fatherName"`
	Email      userfields.Email      `json:"email"`
	Roles      []userfields.UserRole `json:"roles"`
}

func (u *UserDTO) ValidatableVariables() []validate.Validatable {
	list := []validate.Validatable{u.Login, u.Password, u.Name, u.Surname, u.FatherName, u.Email}
	for _, r := range u.Roles {
		list = append(list, r)
	}
	return list
}

func (u *UserDTO) Validate() *amiderrors.ErrorResponse {
	return validate.ValidateStructVariables(u.ValidatableVariables()...)
}

func NewUserDTO(id uint64, login userfields.Login, password userfields.Password, name userfields.Name, surname userfields.Surname, fatherName userfields.FatherName, email userfields.Email, role []userfields.UserRole) *UserDTO {
	return &UserDTO{id, login, password, name, surname, fatherName, email, role}
}
