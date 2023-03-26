package userfields

import (
	"github.com/amidgo/amiddocs/pkg/amiderrors"
	"github.com/amidgo/amiddocs/pkg/validate"
)

type Login string

const (
	LOGIN_MAX_LENGTH = 100
	LOGIN_MIN_LENGTH = 0
	LOGIN_FIELD_NAME = "Логин"
)

func (l Login) Validate() *amiderrors.ErrorResponse {
	return validate.StringValidate(string(l), LOGIN_FIELD_NAME, LOGIN_MIN_LENGTH, LOGIN_MAX_LENGTH)
}
