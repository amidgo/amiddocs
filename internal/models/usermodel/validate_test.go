package usermodel_test

import (
	"testing"

	"github.com/amidgo/amiddocs/internal/errorutils/usererror"
	"github.com/amidgo/amiddocs/internal/models/usermodel"
	"github.com/amidgo/amiddocs/internal/models/usermodel/userfields"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
	"github.com/amidgo/amiddocs/pkg/amidstr"
	"github.com/amidgo/amiddocs/pkg/assert"
)

var (
	name       userfields.Name       = "name"
	surname    userfields.Surname    = "surname"
	fatherName userfields.FatherName = "fatherName"
	email      userfields.Email      = "example@mail.ru"
	roles                            = []userfields.Role{
		userfields.ADMIN,
	}

	maxName       userfields.Name       = userfields.Name(amidstr.MakeString(userfields.NAME_MAX_LENGTH))
	maxSurname    userfields.Surname    = userfields.Surname(amidstr.MakeString(userfields.SURNAME_MAX_LENGTH))
	maxFatherName userfields.FatherName = userfields.FatherName(amidstr.MakeString(userfields.FATHERNAME_MAX_LENGTH))
	maxEmail      userfields.Email      = userfields.Email(amidstr.MakeString(userfields.EMAIL_MAX_LENGTH-8) + "@mail.ru")

	minName    userfields.Name    = userfields.Name(amidstr.MakeString(userfields.NAME_MIN_LENGTH))
	minSurname userfields.Surname = userfields.Surname(amidstr.MakeString(userfields.SURNAME_MIN_LENGTH))

	wrongRoles = []userfields.Role{
		"INDUS",
	}

	minUserNameDTO = usermodel.NewCreateUserDTO(minName, surname, fatherName, email, roles)
	minSurnameDTO  = usermodel.NewCreateUserDTO(name, minSurname, fatherName, email, roles)
	minRolesDTO    = usermodel.NewCreateUserDTO(name, surname, fatherName, email, make([]userfields.Role, 0))

	minFatherNameDTO = usermodel.NewCreateUserDTO(name, surname, "", email, roles)
	minEmailDTO      = usermodel.NewCreateUserDTO(name, surname, fatherName, "", roles)

	wrongRolesDTO = usermodel.NewCreateUserDTO(name, surname, fatherName, email, wrongRoles)

	maxNameDTO       = usermodel.NewCreateUserDTO(maxName, surname, fatherName, email, roles)
	maxSurnameDTO    = usermodel.NewCreateUserDTO(name, maxSurname, fatherName, email, roles)
	maxFatherNameDTO = usermodel.NewCreateUserDTO(name, surname, maxFatherName, email, roles)
	maxEmailDTO      = usermodel.NewCreateUserDTO(name, surname, fatherName, maxEmail, roles)
)

func TestValidateCreateUserDTOWithMinLengthValues(t *testing.T) {
	exp := error(nil)
	act := minUserNameDTO.Validate()
	assert.ErrorEqual(t, exp, act, "validate user name failed with rigth input")

	minUserNameDTO.Name = userfields.Name(amidstr.MakeString(userfields.NAME_MIN_LENGTH - 1))
	exp = amiderrors.WrongLength(userfields.NAME_FIELD_NAME, userfields.NAME_MIN_LENGTH, userfields.NAME_MAX_LENGTH)
	act = minUserNameDTO.Validate()
	assert.ErrorEqual(t, exp, act, "validate user name failed with wrong input")

	exp = nil
	act = minSurnameDTO.Validate()
	assert.ErrorEqual(t, exp, act, "validate surname failed with rigth input")

	minSurnameDTO.Surname = userfields.Surname(amidstr.MakeString(userfields.SURNAME_MIN_LENGTH - 1))
	exp = amiderrors.WrongLength(userfields.SURNAME_FIELD_NAME, userfields.SURNAME_MIN_LENGTH, userfields.SURNAME_MAX_LENGTH)
	act = minSurnameDTO.Validate()
	assert.ErrorEqual(t, exp, act, "validate surname failed with wrong input")

	act = minFatherNameDTO.Validate()
	assert.ErrorEqual(t, nil, act, "empty father name failed")

	act = minEmailDTO.Validate()
	assert.ErrorEqual(t, nil, act, "empty email failed")

	exp = usererror.EMPTY_ROLES
	act = minRolesDTO.Validate()
	assert.ErrorEqual(t, exp, act, "empty roles failed")
}

func TestValidateCreateUserDTOWithMaxLengthValues(t *testing.T) {
	exp := error(nil)
	act := maxNameDTO.Validate()
	assert.ErrorEqual(t, exp, act, "validate name failed with right input")

	maxNameDTO.Name += userfields.Name(amidstr.MakeString(1))
	exp = amiderrors.WrongLength(userfields.NAME_FIELD_NAME, userfields.NAME_MIN_LENGTH, userfields.NAME_MAX_LENGTH)
	act = maxNameDTO.Validate()
	assert.ErrorEqual(t, exp, act, "validate name failed with wrong input")

	exp = nil
	act = maxSurnameDTO.Validate()
	assert.ErrorEqual(t, exp, act, "validate surname failed with right input")

	maxSurnameDTO.Surname += userfields.Surname(amidstr.MakeString(1))
	exp = amiderrors.WrongLength(userfields.SURNAME_FIELD_NAME, userfields.SURNAME_MIN_LENGTH, userfields.SURNAME_MAX_LENGTH)
	act = maxSurnameDTO.Validate()
	assert.ErrorEqual(t, exp, act, "validate surname failed with wrong input")

	exp = nil
	act = maxFatherNameDTO.Validate()
	assert.ErrorEqual(t, exp, act, "validate father name failed with right input")

	maxFatherNameDTO.FatherName += userfields.FatherName(amidstr.MakeString(1))
	exp = amiderrors.WrongLength(userfields.FATHERNAME_FIELD_NAME, userfields.FATHERNAME_MIN_LENGTH, userfields.FATHERNAME_MAX_LENGTH)
	act = maxFatherNameDTO.Validate()
	assert.ErrorEqual(t, exp, act, "validte father name failed with wrong input")

	exp = nil
	act = maxEmailDTO.Validate()
	assert.ErrorEqual(t, exp, act, "validate email failed with right input")

	maxEmailDTO.Email += userfields.Email(amidstr.MakeString(1))
	exp = amiderrors.WrongLength(userfields.EMAIL_FIELD_NAME, userfields.EMAIL_MIN_LENGTH, userfields.EMAIL_MAX_LENGTH)
	act = maxEmailDTO.Validate()
	assert.ErrorEqual(t, exp, act, "validate email failed with wrong input")
}

func TestValidateRoles(t *testing.T) {
	exp := error(usererror.ROLE_NOT_EXIST)
	act := wrongRolesDTO.Validate()
	assert.ErrorEqual(t, exp, act, "validate roles failed with wrong input")

	wrongRolesDTO.Roles = roles
	exp = nil
	act = wrongRolesDTO.Validate()
	assert.ErrorEqual(t, exp, act, "validate roles failed with right input")
}
