package depmodel

const DepartmentTable = "departments"

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
}

var SQL = dep_table{
	ID:        "id",
	Name:      "name",
	ShortName: "short_name",
}

const StudyDepartmentTable = "study_departments"

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
