package groupfields

import (
	"github.com/amidgo/amiddocs/internal/errorutils/grouperror"
)

type EducationForm string

const EDUCATION_FORM_FIELD_NAME = "Тип Формы Обучения"

const (
	FULL_TIME  EducationForm = "FULL_TIME"
	EXTRAMURAL EducationForm = "EXTRAMURAL"
)

func (edf EducationForm) Validate() error {
	for _, r := range []EducationForm{FULL_TIME, EXTRAMURAL} {
		if edf == r {
			return nil
		}
	}
	return grouperror.EDUCATION_FORM_NOT_EXIST
}
