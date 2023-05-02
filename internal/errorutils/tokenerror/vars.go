package tokenerror

import (
	"net/http"

	"github.com/amidgo/amiddocs/pkg/amiderrors"
)

var (
	TOKEN_EXPIRED       = amiderrors.NewErrorResponse("Время действия токена вышло", http.StatusUnauthorized, "token_expired")
	STUDENTID_UNDEFINED = amiderrors.NewErrorResponse("He удалось получить studentid при авторизации", http.StatusUnauthorized, "student_id_undefined")
	ROLE_UNDEFINED      = amiderrors.NewErrorResponse("He удалось получить role при авторизации", http.StatusUnauthorized, "role_undefined")
	USERID_UNDEFINED    = amiderrors.NewErrorResponse("He удалось получить userid при авторизации", http.StatusUnauthorized, "user_id_undefined")
	FORBIDDEN           = amiderrors.NewErrorResponse("Нет доступа", http.StatusForbidden, "forbidden")
)
