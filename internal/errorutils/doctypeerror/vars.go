package doctypeerror

import (
	"net/http"

	"github.com/amidgo/amiddocs/pkg/amiderrors"
)

var (
	DOC_TYPE_NOT_FOUND = amiderrors.NewErrorResponse("Тип документа не найден", http.StatusNotFound, "doc_type_not_found")
	WRONG_USER_ROLE    = amiderrors.NewErrorResponse("Вы не можете запросить данный тип документа", http.StatusBadRequest, "doc_type_wrong_user_role")
)
