package doctemperror

import (
	"net/http"

	"github.com/amidgo/amiddocs/pkg/amiderrors"
)

var (
	DOC_TEMP_NOT_FOUND = amiderrors.NewErrorResponse("Шаблон документа не найден", http.StatusNotFound, "doc_temp_not_found")
)
