package depmodel

const DepartmentTable = "departments"

const (
	DepartmentNameUniqueConstraint      = "departments_name_unique"
	DepartmentShortNameUniqueConstraint = "departments_short_name_unique"
	DepartmentPhotoUniqueConstraint     = "departments_photo_unique"
)

type dep_column string

func (d dep_column) String() string {
	return string(d)
}

func (d dep_column) TableName() string {
	return DepartmentTable
}

type dep_table struct {
	ID        dep_column
	Name      dep_column
	ShortName dep_column
	ImageUrl  dep_column
}

var SQL = dep_table{
	ID:        "id",
	Name:      "name",
	ShortName: "short_name",
	ImageUrl:  "photo",
}

const StudyDepartmentTable = "study_departments"

const (
	ForeignKey_StudyDepartments__Departments = "fk_study_departments__departments"
)

type study_department_column string

func (s study_department_column) String() string {
	return string(s)
}

func (s study_department_column) TableName() string {
	return StudyDepartmentTable
}

type study_department_table struct {
	ID           study_department_column
	DepartmentId study_department_column
}

var SQL_STUDY_DEP = study_department_table{
	ID:           "id",
	DepartmentId: "department_id",
}
