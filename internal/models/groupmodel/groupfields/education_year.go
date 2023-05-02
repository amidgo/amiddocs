package groupfields

import (
	"github.com/amidgo/amiddocs/internal/errorutils/grouperror"
)

type EducationYear uint8

const (
	MIN_EDUCATION_YEAR uint8 = 1
	MAX_EDUCATION_YEAR uint8 = 4
)

func (ey EducationYear) Validate() error {
	if ey < EducationYear(MIN_EDUCATION_YEAR) || ey > EducationYear(MAX_EDUCATION_YEAR) {
		return grouperror.EDUCATION_YEAR_WRONG
	}
	return nil
}
