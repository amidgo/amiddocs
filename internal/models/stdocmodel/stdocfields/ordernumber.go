package stdocfields

import (
	"github.com/amidgo/amiddocs/pkg/validate"
)

type OrderNumber string

const (
	ORDER_NUMBER_FIELD_NAME       = "Номер приказа"
	ORDER_NUMBER_MAX_FIELD_LENGTH = 60
	ORDER_NUMBER_MIN_FIELD_LENGTH = 1
)

func (d OrderNumber) Validate() error {
	return validate.StringValidate(string(d), ORDER_NUMBER_FIELD_NAME, ORDER_NUMBER_MIN_FIELD_LENGTH, ORDER_NUMBER_MAX_FIELD_LENGTH)
}
