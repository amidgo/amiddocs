package userfields

import (
	"github.com/amidgo/amiddocs/pkg/amiderrors"
	"github.com/amidgo/amiddocs/pkg/validate"
)

type Name string

const (
	NAME_MIN_LENGTH = 1
	NAME_MAX_LENGTH = 40
	NAME_FIELD_NAME = "Имя"
)

func (n Name) Validate() *amiderrors.ErrorResponse {
	return validate.StringValidate(string(n), NAME_FIELD_NAME, NAME_MIN_LENGTH, NAME_MAX_LENGTH)
}
