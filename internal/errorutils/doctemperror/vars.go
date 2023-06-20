package doctemperror

import (
	"net/http"

	"github.com/amidgo/amiddocs/pkg/amiderrors"
)

const DOC_TEMPLATE_TYPE = "document_template"

var (
	DOC_TEMP_NOT_FOUND = amiderrors.NewException(http.StatusNotFound, DOC_TEMPLATE_TYPE, "not_found")
)
