package userfields

import (
	"github.com/amidgo/amiddocs/pkg/validate"
)

type Surname string

const (
	SURNAME_MIN_LENGTH = 1
	SURNAME_MAX_LENGTH = 60
	SURNAME_FIELD_NAME = "Фамилия"
)

func (sn Surname) Validate() error {
	return validate.StringValidate(string(sn), SURNAME_FIELD_NAME, SURNAME_MIN_LENGTH, SURNAME_MAX_LENGTH)
}
