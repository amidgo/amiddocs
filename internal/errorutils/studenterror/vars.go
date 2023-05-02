package studenterror

import (
	"net/http"

	"github.com/amidgo/amiddocs/pkg/amiderrors"
)

var (
	STUDENT_NOT_FOUND = amiderrors.NewErrorResponse("Студент не найден", http.StatusNotFound, "student_not_found")
)
