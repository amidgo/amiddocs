package usermodel

import (
	"fmt"

	"github.com/amidgo/amiddocs/internal/models/usermodel/userfields"
	"github.com/amidgo/amiddocs/pkg/validate"
)

type UserDTO struct {
	ID         uint64                `json:"id" db:"id"`
	Login      userfields.Login      `json:"login" db:"login"`
	Password   userfields.Password   `json:"-" db:"password"`
	Name       userfields.Name       `json:"name" db:"name"`
	Surname    userfields.Surname    `json:"surname" db:"surname"`
	FatherName userfields.FatherName `json:"fatherName" db:"father_name"`
	Email      userfields.Email      `json:"email" db:"email"`
	Roles      []userfields.Role     `json:"roles" db:"roles"`
}

func (u *UserDTO) ValidatableVariables() []validate.Validatable {
	list := []validate.Validatable{u.Login, u.Password, u.Name, u.Surname, u.FatherName, u.Email}
	for _, r := range u.Roles {
		list = append(list, r)
	}
	return list
}

func NewUserDTO(id uint64, login userfields.Login, password userfields.Password, name userfields.Name, surname userfields.Surname, fatherName userfields.FatherName, email userfields.Email, role []userfields.Role) *UserDTO {
	return &UserDTO{id, login, password, name, surname, fatherName, email, role}
}

func (u *UserDTO) Fio() string {
	if len(u.FatherName) == 0 {
		return fmt.Sprintf("%s %s", u.Surname, u.Name)
	}
	return fmt.Sprintf("%s %s %s", u.Surname, u.Name, u.FatherName)
}
