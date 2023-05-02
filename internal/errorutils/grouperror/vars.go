package grouperror

import (
	"net/http"

	"github.com/amidgo/amiddocs/pkg/amiderrors"
)

var (
	GROUP_NOT_FOUND          = amiderrors.NewErrorResponse("Группа не найдена", http.StatusNotFound, "group_not_found")
	EDUCATION_FORM_NOT_EXIST = amiderrors.NewErrorResponse("Формы обучения не существует", http.StatusBadRequest, "education_form_not_exist")
	EDUCATION_YEAR_WRONG     = amiderrors.NewErrorResponse("Неправильно задан год обучения", http.StatusBadRequest, "education_year_wrong")
	GROUP_NAME_ALREADY_EXIST = amiderrors.NewErrorResponse("Данное название группы уже используется", http.StatusBadRequest, "group_name_already_exist")
	INVALID_EDUCATION_DATE   = amiderrors.NewErrorResponse("Начало не может быть позже конца обучения", http.StatusBadRequest, "invalid_education_date")
)
