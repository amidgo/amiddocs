package usererror

import (
	"net/http"

	"github.com/amidgo/amiddocs/pkg/amiderrors"
)

const USER_TYPE = "user"

var (
	NOT_FOUND           = amiderrors.NewException(http.StatusNotFound, USER_TYPE, "not_found")
	ROLE_NOT_EXIST      = amiderrors.NewException(http.StatusBadRequest, USER_TYPE, "role_not_exist")
	WRONG_EMAIL         = amiderrors.NewException(http.StatusBadRequest, USER_TYPE, "wrong_email")
	LOGIN_ALREADY_EXIST = amiderrors.NewException(http.StatusBadRequest, USER_TYPE, "login_exist")
	EMAIL_ALREADY_EXIST = amiderrors.NewException(http.StatusBadRequest, USER_TYPE, "email_exist")
	WRONG_PASSWORD      = amiderrors.NewException(http.StatusBadRequest, USER_TYPE, "wrong_password")
	EMPTY_ROLES         = amiderrors.NewException(http.StatusBadRequest, USER_TYPE, "empty_roles")
)
