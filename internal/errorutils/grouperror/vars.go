package grouperror

import (
	"net/http"

	"github.com/amidgo/amiddocs/pkg/amiderrors"
)

const GROUP_TYPE = "group"

var (
	GROUP_NOT_FOUND          = amiderrors.NewException(http.StatusNotFound, GROUP_TYPE, "not_found")
	EDUCATION_FORM_NOT_EXIST = amiderrors.NewException(http.StatusBadRequest, GROUP_TYPE, "education_form_not_exist")
	EDUCATION_YEAR_WRONG     = amiderrors.NewException(http.StatusBadRequest, GROUP_TYPE, "education_year_wrong")
	GROUP_NAME_ALREADY_EXIST = amiderrors.NewException(http.StatusBadRequest, GROUP_TYPE, "name_exist")
	INVALID_EDUCATION_DATE   = amiderrors.NewException(http.StatusBadRequest, GROUP_TYPE, "invalid_education_date")
)
