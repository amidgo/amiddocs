package studenterror

import (
	"net/http"

	"github.com/amidgo/amiddocs/pkg/amiderrors"
)

const STUDENT_TYPE = "student"

var (
	STUDENT_NOT_FOUND = amiderrors.NewException(http.StatusNotFound, STUDENT_TYPE, "not_found")
)
