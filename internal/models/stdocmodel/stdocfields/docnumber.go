package stdocfields

import (
	"github.com/amidgo/amiddocs/pkg/validate"
)

type DocNumber string

const (
	DOC_NUMBER_FIELD_NAME       = "Номер студенческого билета"
	DOC_NUMBER_MAX_FIELD_LENGTH = 60
	DOC_NUMBER_MIN_FIELD_LENGTH = 1
)

func (d DocNumber) Validate() error {
	return validate.StringValidate(string(d), DOC_NUMBER_FIELD_NAME, DOC_NUMBER_MIN_FIELD_LENGTH, DOC_NUMBER_MAX_FIELD_LENGTH)
}
