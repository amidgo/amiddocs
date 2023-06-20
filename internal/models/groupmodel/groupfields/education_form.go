package groupfields

import (
	"reflect"
	"strings"

	"github.com/amidgo/amiddocs/internal/errorutils/grouperror"
)

type EducationForm string

const EDUCATION_FORM_FIELD_NAME = "Тип Формы Обучения"

const (
	CSV_FULL_TIME                = "очная"
	CSV_EXTRAMURAL               = "заочная"
	FULL_TIME      EducationForm = "FULL_TIME"
	EXTRAMURAL     EducationForm = "EXTRAMURAL"
)

func (e EducationForm) Parse(s string, r *reflect.Value) {
	s = strings.ToLower(s)
	switch s {
	case CSV_FULL_TIME:
		r.SetString(string(FULL_TIME))
	case CSV_EXTRAMURAL:
		r.SetString(string(EXTRAMURAL))
	}
}

func (edf EducationForm) Validate() error {
	for _, r := range []EducationForm{FULL_TIME, EXTRAMURAL} {
		if edf == r {
			return nil
		}
	}
	return grouperror.EDUCATION_FORM_NOT_EXIST
}
