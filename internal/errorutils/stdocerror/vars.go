package stdocerror

import (
	"net/http"

	"github.com/amidgo/amiddocs/pkg/amiderrors"
)

var (
	DOC_NOT_FOUND      = amiderrors.NewErrorResponse("Студенческий билет не найден", http.StatusNotFound, "st_doc_not_found")
	DOC_NUMBER_EXIST   = amiderrors.NewErrorResponse("Номер студенческого билета уже существует", http.StatusBadRequest, "doc_number_exist")
	ORDER_NUMBER_EXIST = amiderrors.NewErrorResponse("Данный номер приказа уже используется", http.StatusBadRequest, "order_number_exist")
)
