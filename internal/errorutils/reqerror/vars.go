package reqerror

import (
	"net/http"

	"github.com/amidgo/amiddocs/pkg/amiderrors"
)

var (
	REQ_NOT_FOUND        = amiderrors.NewErrorResponse("Заявка не найдена", http.StatusNotFound, "st_req_not_found")
	WRONG_DOCUMENT_TYPE  = amiderrors.NewErrorResponse("Данного типа документа не существует", http.StatusBadRequest, "wrong_document_type")
	WRONG_DOCUMENT_COUNT = amiderrors.NewErrorResponse("Вы можете заказать от 1 до 3 экземпляров максимум", http.StatusBadRequest, "wrong_document_count")
	REQ_REFRESH_DATE     = amiderrors.NewErrorResponse(
		"Вы уже заказывали этот документ в последнее время, "+
			"вам требуется подождать некоторое время перед тем как вы сможете сделать это снова",
		http.StatusBadRequest,
		"req_refresh_date",
	)
)
