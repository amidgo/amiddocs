package doctypeerror

import (
	"net/http"

	"github.com/amidgo/amiddocs/pkg/amiderrors"
)

const DOCUMENT_TYPE_TYPE = "document_type"

var (
	DOC_TYPE_NOT_FOUND = amiderrors.NewException(http.StatusNotFound, DOCUMENT_TYPE_TYPE, "not_found")
	WRONG_USER_ROLE    = amiderrors.NewException(http.StatusBadRequest, DOCUMENT_TYPE_TYPE, "wrong_user_role")
)
