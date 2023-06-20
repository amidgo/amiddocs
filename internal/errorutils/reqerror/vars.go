package reqerror

import (
	"net/http"

	"github.com/amidgo/amiddocs/pkg/amiderrors"
)

const REQUEST_TYPE = "request"

var (
	REQ_NOT_FOUND        = amiderrors.NewException(http.StatusNotFound, REQUEST_TYPE, "not_found")
	WRONG_DOCUMENT_TYPE  = amiderrors.NewException(http.StatusBadRequest, REQUEST_TYPE, "wrong_document_type")
	WRONG_DOCUMENT_COUNT = amiderrors.NewException(http.StatusBadRequest, REQUEST_TYPE, "wrong_document_count")
	REQ_REFRESH_DATE     = amiderrors.NewException(http.StatusBadRequest, REQUEST_TYPE, "req_refresh_date")
)
