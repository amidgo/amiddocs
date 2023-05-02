package userfields

import (
	"net/mail"

	usererrorutils "github.com/amidgo/amiddocs/internal/errorutils/usererror"
	"github.com/amidgo/amiddocs/pkg/amidstr"
	"github.com/amidgo/amiddocs/pkg/validate"
)

const (
	EMAIL_MAX_LENGTH = 100
	EMAIL_MIN_LENGTH = 0
	EMAIL_FIELD_NAME = "Почта"
)

type Email string

func (e Email) Validate() error {
	if !e.isValid() {
		return usererrorutils.WRONG_EMAIL
	}
	return validate.StringValidate(string(e), EMAIL_FIELD_NAME, EMAIL_MIN_LENGTH, EMAIL_MAX_LENGTH)
}

func (e Email) isValid() bool {
	if len(e) == 0 {
		return true
	}
	_, err := mail.ParseAddress(string(e))
	return err == nil
}

func (e *Email) Scan(src interface{}) error {
	s, err := amidstr.ScanNullString(src)
	*e = Email(s)
	return err
}

func (e *Email) UnmarshalJSON(b []byte) error {
	s, err := amidstr.UnmarshalNullString(b)
	*e = Email(s)
	return err
}

func (e Email) MarshalJSON() ([]byte, error) {
	return amidstr.MarshalNullString(string(e))
}
