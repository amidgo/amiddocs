package userfields

import (
	"github.com/amidgo/amiddocs/pkg/amiderrors"
	"github.com/amidgo/amiddocs/pkg/validate"
	"golang.org/x/crypto/bcrypt"
)

type Password string

const (
	PASSWORD_NAME       = "Пароль"
	PASSWORD_MIN_LENGTH = 6
	PASSWORD_MAX_LENGTH = 32
	PASSWORD_HASH_COST  = 10
)

func (p Password) Verify(password string) bool {
	hash := []byte(p)
	err := bcrypt.CompareHashAndPassword(hash, []byte(password))
	return err == nil
}

func (p Password) Hash() (string, *amiderrors.ErrorResponse) {
	b, err := bcrypt.GenerateFromPassword([]byte(p), PASSWORD_HASH_COST)
	return string(b), amiderrors.NewInternalErrorResponse(err)
}

func (p Password) Validate() *amiderrors.ErrorResponse {
	return validate.StringValidate(string(p), PASSWORD_NAME, PASSWORD_MIN_LENGTH, PASSWORD_MAX_LENGTH)
}
