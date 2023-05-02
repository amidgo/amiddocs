package usermodel

import (
	"github.com/amidgo/amiddocs/internal/models/usermodel/userfields"
	"github.com/amidgo/amiddocs/pkg/validate"
)

type LoginForm struct {
	Login    userfields.Login    `json:"login"`
	Password userfields.Password `json:"password"`
}

func NewLoginForm(login userfields.Login, password userfields.Password) *LoginForm {
	return &LoginForm{Login: login, Password: password}
}

func (l *LoginForm) ValidatableVariables() []validate.Validatable {
	return []validate.Validatable{l.Login, l.Password}
}
