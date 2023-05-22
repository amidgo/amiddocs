package studentmodel

const StudentTable = "students"

type student_column string

func (s student_column) String() string {
	return string(s)
}

func (s student_column) TableName() string {
	return StudentTable
}

var SQL = struct {
	ID      student_column
	UserID  student_column
	GroupId student_column
}{
	ID:      "id",
	UserID:  "user_id",
	GroupId: "group_id",
}
