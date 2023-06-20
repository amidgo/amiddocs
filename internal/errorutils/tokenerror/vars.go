package tokenerror

import (
	"net/http"

	"github.com/amidgo/amiddocs/pkg/amiderrors"
	"github.com/amidgo/amiddocs/pkg/jwtrs"
)

const TOKEN_TYPE = jwtrs.TOKEN_TYPE

var (
	TOKEN_NOT_FOUND       = amiderrors.NewException(http.StatusNotFound, TOKEN_TYPE, "not_found")
	STUDENTID_UNDEFINED   = amiderrors.NewException(http.StatusUnauthorized, TOKEN_TYPE, "student_id_undefined")
	ROLE_UNDEFINED        = amiderrors.NewException(http.StatusUnauthorized, TOKEN_TYPE, "role_undefined")
	USERID_UNDEFINED      = amiderrors.NewException(http.StatusUnauthorized, TOKEN_TYPE, "user_id_undefined")
	REFRESH_TOKEN_EXPIRED = amiderrors.NewException(http.StatusUnauthorized, TOKEN_TYPE, "refresh_expired")
	FORBIDDEN             = amiderrors.NewException(http.StatusForbidden, TOKEN_TYPE, "forbidden")
)
