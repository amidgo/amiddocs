package groupmodel

const GroupTable = "groups"

type group_column string

func (g group_column) String() string {
	return string(g)
}

func (g group_column) TableName() string {
	return GroupTable
}

var SQL = struct {
	ID                  group_column
	Name                group_column
	IsBudget            group_column
	EducationForm       group_column
	EducationStartDate  group_column
	EducationYear       group_column
	EducationFinishDate group_column
	StudyDepartmentId   group_column
}{
	ID:                  "id",
	Name:                "name",
	IsBudget:            "is_budget",
	EducationForm:       "education_form",
	EducationStartDate:  "education_start_date",
	EducationYear:       "education_year",
	EducationFinishDate: "education_finish_date",
	StudyDepartmentId:   "study_department_id",
}
