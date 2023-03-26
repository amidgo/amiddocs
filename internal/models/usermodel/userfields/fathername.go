package userfields

import (
	"github.com/amidgo/amiddocs/pkg/amiderrors"
	"github.com/amidgo/amiddocs/pkg/validate"
)

type FatherName string

const (
	FATHERNAME_MIN_LENGTH = 0
	FATHERNAME_MAX_LENGTH = 40
	FATHERNAME_FIELD_NAME = "Отчество"
)

func (fn FatherName) Validate() *amiderrors.ErrorResponse {
	return validate.StringValidate(string(fn), FATHERNAME_FIELD_NAME, FATHERNAME_MIN_LENGTH, FATHERNAME_MAX_LENGTH)
}
