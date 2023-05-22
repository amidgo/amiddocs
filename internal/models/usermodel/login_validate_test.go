package usermodel_test

import (
	"testing"

	"github.com/amidgo/amiddocs/internal/models/usermodel"
	"github.com/amidgo/amiddocs/internal/models/usermodel/userfields"
	"github.com/amidgo/amiddocs/pkg/amidstr"
)

var (
	login    = userfields.Login("amidman")
	password = userfields.Password("amidman")

	maxLogin    = userfields.Login(amidstr.MakeString(userfields.LOGIN_MAX_LENGTH))
	maxPassword = userfields.Password(amidstr.MakeString(userfields.PASSWORD_MAX_LENGTH))

	minLogin    = userfields.Login(amidstr.MakeString(userfields.LOGIN_MIN_LENGTH))
	minPassword = userfields.Password(amidstr.MakeString(userfields.PASSWORD_MIN_LENGTH))

	maxLoginDTO    = usermodel.NewLoginForm(maxLogin, password)
	maxPasswordDTO = usermodel.NewLoginForm(login, maxPassword)
)

func ValidateLoginFormMinLengthValues(t *testing.T) {

}

func ValidateLoginFormMaxLengthValues(t *testing.T) {

}
