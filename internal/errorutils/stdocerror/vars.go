package stdocerror

import (
	"net/http"

	"github.com/amidgo/amiddocs/pkg/amiderrors"
)

const STUDENT_DOCUMENT_TYPE = "student_document"

var (
	STUDENT_ALREADY_HAVE_DOCUMENT = amiderrors.NewException(http.StatusBadRequest, STUDENT_DOCUMENT_TYPE, "student_already_have_document")
	DOC_NOT_FOUND                 = amiderrors.NewException(http.StatusNotFound, STUDENT_DOCUMENT_TYPE, "not_found")
	DOC_NUMBER_EXIST              = amiderrors.NewException(http.StatusBadRequest, STUDENT_DOCUMENT_TYPE, "doc_number_exist")
)
