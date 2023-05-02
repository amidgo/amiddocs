package userfields

import (
	"github.com/amidgo/amiddocs/pkg/validate"
)

type Password string

const (
	PASSWORD_NAME       = "Пароль"
	PASSWORD_MIN_LENGTH = 6
	PASSWORD_MAX_LENGTH = 32
)

func (p Password) Validate() error {
	return validate.StringValidate(string(p), PASSWORD_NAME, PASSWORD_MIN_LENGTH, PASSWORD_MAX_LENGTH)
}
