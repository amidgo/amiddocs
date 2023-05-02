package userfields

import (
	"github.com/amidgo/amiddocs/pkg/amidstr"
	"github.com/amidgo/amiddocs/pkg/validate"
)

type FatherName string

const (
	FATHERNAME_MIN_LENGTH = 0
	FATHERNAME_MAX_LENGTH = 40
	FATHERNAME_FIELD_NAME = "Отчество"
)

func (fn FatherName) Validate() error {
	return validate.StringValidate(string(fn), FATHERNAME_FIELD_NAME, FATHERNAME_MIN_LENGTH, FATHERNAME_MAX_LENGTH)
}

func (fn *FatherName) Scan(src interface{}) error {
	s, err := amidstr.ScanNullString(src)
	*fn = FatherName(s)
	return err
}

func (fn *FatherName) UnmarshalJSON(b []byte) error {
	s, err := amidstr.UnmarshalNullString(b)
	*fn = FatherName(s)
	return err
}

func (fn FatherName) MarshalJSON() ([]byte, error) {
	return amidstr.MarshalNullString(string(fn))
}
