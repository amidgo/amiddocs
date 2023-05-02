package studentstorage

import (
	"context"
	"fmt"

	"github.com/amidgo/amiddocs/internal/models/groupmodel"
	"github.com/amidgo/amiddocs/internal/models/stdocmodel"
	"github.com/amidgo/amiddocs/internal/models/studentmodel"
	"github.com/amidgo/amiddocs/internal/models/usermodel"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
	"github.com/jackc/pgx/v5"
)

func getStudentQuery(query string) string {
	return fmt.Sprintf(
		`
	SELECT students.id,
		users.id,
		users.login, users.password, users.name, users.surname, users.father_name, users.email, 
		array_agg(roles.role) as roles,
		student_documents.id, 
		student_documents.doc_number, student_documents.order_number,
		student_documents.order_date,student_documents.study_start_date,
		groups.id,
		groups.name,groups.is_budget,groups.education_form,groups.education_start_date,
		groups.education_year,groups.education_finish_date,groups.department_id 
	FROM students 
		INNER JOIN users ON students.user_id = users.id 
		INNER JOIN groups ON students.group_id = groups.id 
		INNER JOIN student_documents ON students.student_document_id = student_documents.id 
		INNER JOIN user_roles ON user_roles.user_id = users.id
		INNER JOIN roles ON user_roles.role_id = roles.id
	%s 
	GROUP BY users.id, students.id, student_documents.id, groups.id `,
		query,
	)
}

func scanStudent(row pgx.Row) (*studentmodel.StudentDTO, error) {
	user := new(usermodel.UserDTO)
	document := new(stdocmodel.StudentDocumentDTO)
	group := new(groupmodel.GroupDTO)
	student := studentmodel.NewStudentDTO(0, user, document, group)
	err := row.Scan(
		&student.ID,
		&user.ID,
		&user.Login, &user.Password, &user.Name, &user.Surname, &user.FatherName, &user.Email, &user.Roles,
		&document.ID,
		&document.DocNumber, &document.OrderNumber, &document.OrderDate, &document.StudyStartDate,
		&group.ID,
		&group.Name, &group.IsBudget, &group.EducationForm, &group.EducationStartDate,
		&group.EducationYear, &group.EducationFinishDate, &group.DepartmentId,
	)
	if err != nil {
		return nil, err
	}
	return student, nil
}

func (s *studentStorage) getStudentByQuery(
	ctx context.Context,
	query string,
	args ...interface{},
) (*studentmodel.StudentDTO, error) {
	row := s.p.Pool.QueryRow(ctx, query, args...)
	return scanStudent(row)
}

func (s *studentStorage) StudentById(ctx context.Context, id uint64) (*studentmodel.StudentDTO, error) {
	st, err := s.getStudentByQuery(ctx, getStudentQuery("WHERE students.id = $1"), id)
	if err != nil {
		return nil, studentError(err, amiderrors.NewCause("student by id query", "StudentById", _PROVIDER))
	}
	return st, nil
}

func (c *studentStorage) AllStudents(ctx context.Context) ([]*studentmodel.StudentDTO, error) {
	rows, err := c.p.Pool.Query(ctx, getStudentQuery(""))
	if err != nil {
		return nil, studentError(err, amiderrors.NewCause("all students query", "AllStudents", _PROVIDER))
	}
	studentList := make([]*studentmodel.StudentDTO, 0)
	for rows.Next() {
		st, err := scanStudent(rows)
		if err != nil {
			return nil, studentError(err, amiderrors.NewCause("scan student", "AllStudents", _PROVIDER))
		}
		studentList = append(studentList, st)
	}
	return studentList, nil
}
