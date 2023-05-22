package groupmodel

import (
	"github.com/amidgo/amiddocs/internal/models/groupmodel/groupfields"
	"github.com/amidgo/amiddocs/pkg/amidtime"
	"github.com/amidgo/amiddocs/pkg/validate"
)

type GroupDTO struct {
	ID                  uint64                    `json:"id" db:"id"`
	Name                groupfields.Name          `json:"name" db:"name"`
	IsBudget            bool                      `json:"isBudget" db:"is_budget"`
	EducationForm       groupfields.EducationForm `json:"educationForm" db:"education_form"`
	EducationStartDate  amidtime.Date             `json:"educationStartDate" db:"education_start_date"`
	EducationYear       groupfields.EducationYear `json:"educationYear" db:"education_year"`
	EducationFinishDate amidtime.Date             `json:"educationFinishDate" db:"education_finish_date"`
	StudyDepartmentId   uint64                    `json:"studyDepartmentId" db:"department_id"`
}

func (g *GroupDTO) ValidatableVariables() []validate.Validatable {
	return []validate.Validatable{g.Name, g.EducationForm, g.EducationYear}
}

func NewGroupDTO(
	id uint64,
	name groupfields.Name,
	isBudget bool,
	educationForm groupfields.EducationForm,
	educationStartDate, educationFinishDate amidtime.Date,
	educationYear groupfields.EducationYear,
	departmentId uint64,
) *GroupDTO {
	return &GroupDTO{
		ID:                  id,
		Name:                name,
		IsBudget:            isBudget,
		EducationForm:       educationForm,
		EducationStartDate:  educationStartDate,
		EducationFinishDate: educationFinishDate,
		EducationYear:       educationYear,
		StudyDepartmentId:   departmentId,
	}
}
