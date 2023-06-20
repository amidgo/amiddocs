package departmenterror

import (
	"net/http"

	"github.com/amidgo/amiddocs/pkg/amiderrors"
)

const DEPARTMENT_TYPE = "department"

var (
	STUDY_DEPARTMENT_NOT_FOUND = amiderrors.NewException(http.StatusNotFound, DEPARTMENT_TYPE, "study_dep_not_found")
	DEPARTMENT_NOT_FOUND       = amiderrors.NewException(http.StatusNotFound, DEPARTMENT_TYPE, "not_found")
	NAME_EXIST                 = amiderrors.NewException(http.StatusBadRequest, DEPARTMENT_TYPE, "name_exist")
	SHORT_NAME_EXIST           = amiderrors.NewException(http.StatusBadRequest, DEPARTMENT_TYPE, "short_name_exist")
)
