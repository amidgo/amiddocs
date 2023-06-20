package groupmodel

import (
	"reflect"
	"strings"

	"github.com/amidgo/amiddocs/internal/models/depmodel/depfields"
	"github.com/amidgo/amiddocs/internal/models/groupmodel/groupfields"
	"github.com/amidgo/amiddocs/pkg/amidtime"
	"github.com/amidgo/amiddocs/pkg/validate"
)

type BudgetStatus bool

func (b BudgetStatus) Parse(s string, r *reflect.Value) {
	s = strings.ToLower(s)
	r.SetBool(s == "да")
}

type GroupCSV struct {
	Name                groupfields.Name          `csv:"Название"`
	IsBudget            BudgetStatus              `csv:"Бюджетная"`
	DepartmentShortName depfields.ShortName       `csv:"Инициалы Отделения"`
	EducationForm       groupfields.EducationForm `csv:"Форма Обучения"`
	EducationStartDate  amidtime.Date             `csv:"Дата Начала Обучения"`
	EducationFinishDate amidtime.Date             `csv:"Дата Окончания Обучения"`
	EducationYear       groupfields.EducationYear `csv:"Курс"`
}

func NewGroupCSV(
	name groupfields.Name,
	isBudget bool,
	departmentShortName depfields.ShortName,
	educationForm groupfields.EducationForm,
	educationStartYear amidtime.Date,
	educationFinishDate amidtime.Date,
	educationYear groupfields.EducationYear,
) *GroupCSV {
	return &GroupCSV{
		Name:                name,
		IsBudget:            BudgetStatus(isBudget),
		DepartmentShortName: departmentShortName,
		EducationForm:       educationForm,
		EducationStartDate:  educationStartYear,
		EducationFinishDate: educationFinishDate,
		EducationYear:       educationYear,
	}
}

func (g *GroupCSV) ValidatableVariables() []validate.Validatable {
	return []validate.Validatable{g.Name, g.EducationForm, g.EducationYear}
}

func (g *GroupCSV) GroupDTO(groupId, departmentId uint64) *GroupDTO {
	return NewGroupDTO(
		groupId,
		g.Name,
		bool(g.IsBudget),
		g.EducationForm,
		g.EducationStartDate,
		g.EducationFinishDate,
		g.EducationYear,
		departmentId,
	)
}
