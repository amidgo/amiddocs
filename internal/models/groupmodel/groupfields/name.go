package groupfields

import (
	"github.com/amidgo/amiddocs/pkg/validate"
)

type Name string

const (
	NAME_FIELD_NAME       = "Название Группы"
	NAME_FIELD_MIN_LENGTH = 1
	NAME_FIELD_MAX_LENGTH = 20
)

func (n Name) Validate() error {
	return validate.StringValidate(string(n), NAME_FIELD_NAME, NAME_FIELD_MIN_LENGTH, NAME_FIELD_MAX_LENGTH)
}
