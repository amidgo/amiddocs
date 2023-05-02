package departmenterror

import (
	"net/http"

	"github.com/amidgo/amiddocs/pkg/amiderrors"
)

var (
	DEPARMENT_NOT_FOUND = amiderrors.NewErrorResponse("Отделение не найдено", http.StatusNotFound, "department_not_found")
	NAME_EXIST          = amiderrors.NewErrorResponse("Данное название уже используется", http.StatusBadRequest, "dep_name_exist")
	SHORT_NAME_EXIST    = amiderrors.NewErrorResponse("Короткое имя уже занято", http.StatusBadRequest, "dep_short_name_exist")
)
