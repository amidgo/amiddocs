package depfields

import (
	"github.com/amidgo/amiddocs/pkg/validate"
)

type ShortName string

const (
	SHORT_NAME_FIELD_NAME = "Название Отделения"
	SHORT_NAME_MAX_LENGTH = 200
	SHORT_NAME_MIN_LENGTH = 1
)

func (n ShortName) Validate() error {
	return validate.StringValidate(string(n), SHORT_NAME_FIELD_NAME, SHORT_NAME_MIN_LENGTH, SHORT_NAME_MAX_LENGTH)
}
