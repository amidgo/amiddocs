package usererror

import (
	"net/http"

	"github.com/amidgo/amiddocs/pkg/amiderrors"
)

var (
	NOT_FOUND           = amiderrors.NewErrorResponse("Пользователь не найден", http.StatusNotFound, "user_not_found")
	ROLE_NOT_EXIST      = amiderrors.NewErrorResponse("Данной роли не существует", http.StatusBadRequest, "role_not_exist")
	WRONG_EMAIL         = amiderrors.NewErrorResponse("Введённая почта недействительна", http.StatusBadRequest, "wrong_email")
	LOGIN_ALREADY_EXIST = amiderrors.NewErrorResponse("Данный логин уже занят", http.StatusBadRequest, "login_already_exist")
	EMAIL_ALREADY_EXIST = amiderrors.NewErrorResponse("Данная почта уже используется", http.StatusBadRequest, "email_already_exist")
	WRONG_PASSWORD      = amiderrors.NewErrorResponse("Неверный пароль", http.StatusBadRequest, "wrong_password")
	EMPTY_ROLES         = amiderrors.NewErrorResponse("У пользователя должна быть хотя бы одна роль", http.StatusBadRequest, "empty_roles")
)
