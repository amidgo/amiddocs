package userfields

import (
	"net/mail"

	usererrorutils "github.com/amidgo/amiddocs/internal/errorutils/user_error_utils"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
	"github.com/amidgo/amiddocs/pkg/validate"
)

const (
	EMAIL_MAX_LENGTH = 100
	EMAIL_MIN_LENGTH = 0
	EMAIL_FIELD_NAME = "Почта"
)

type Email string

func (e Email) Validate() *amiderrors.ErrorResponse {
	if !e.isValid() {
		return usererrorutils.WRONG_EMAIL
	}
	return validate.StringValidate(string(e), EMAIL_FIELD_NAME, EMAIL_MIN_LENGTH, EMAIL_MAX_LENGTH)
}

func (e Email) isValid() bool {
	_, err := mail.ParseAddress(string(e))
	return err == nil
}
