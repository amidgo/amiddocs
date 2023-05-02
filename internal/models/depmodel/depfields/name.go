package depfields

import (
	"github.com/amidgo/amiddocs/pkg/validate"
)

type Name string

const (
	NAME_FIELD_NAME = "Название Отделения"
	NAME_MAX_LENGTH = 200
	NAME_MIN_LENGTH = 1
)

func (n Name) Validate() error {
	return validate.StringValidate(string(n), NAME_FIELD_NAME, NAME_MIN_LENGTH, NAME_MAX_LENGTH)
}
