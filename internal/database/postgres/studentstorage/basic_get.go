package studentstorage

import (
	"context"
	"fmt"

	"github.com/amidgo/amiddocs/internal/models/depmodel"
	"github.com/amidgo/amiddocs/internal/models/groupmodel"
	"github.com/amidgo/amiddocs/internal/models/stdocmodel"
	"github.com/amidgo/amiddocs/internal/models/studentmodel"
	"github.com/amidgo/amiddocs/internal/models/usermodel"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
	"github.com/amidgo/amiddocs/pkg/sqlutils"
	"github.com/jackc/pgx/v5"
)

func getStudentQuery(query string) string {
	return fmt.Sprintf(
		`
		SELECT 
			%s,%s, %s, %s, %s, %s, %s, %s, array_agg(%s) as roles,
			%s, %s, %s,%s, %s, %s,%s,%s,%s,%s,%s,%s,%s,
			%s, %s, %s, %s
		FROM students 
			INNER JOIN %s ON %s = %s 
			INNER JOIN %s ON %s = %s 
			INNER JOIN %s ON %s = %s 
			INNER JOIN %s ON %s = %s
			INNER JOIN %s ON %s = %s
			INNER JOIN %s ON %s = %s
			INNER JOIN %s ON %s = %s
		%s 
		GROUP BY %s, %s, %s, %s, %s`,

		// selectable variables
		sqlutils.Full(studentmodel.SQL.ID),
		sqlutils.Full(usermodel.SQL.ID),
		sqlutils.Full(usermodel.SQL.Login),
		sqlutils.Full(usermodel.SQL.Password),
		sqlutils.Full(usermodel.SQL.Name),
		sqlutils.Full(usermodel.SQL.Surname),
		sqlutils.Full(usermodel.SQL.FatherName),
		sqlutils.Full(usermodel.SQL.Email),
		sqlutils.Full(usermodel.SQL_ROLES.Role),
		sqlutils.Full(stdocmodel.SQL.ID),
		sqlutils.Full(stdocmodel.SQL.StudentId),
		sqlutils.Full(stdocmodel.SQL.DocNumber),
		sqlutils.Full(stdocmodel.SQL.OrderNumber),
		sqlutils.Full(stdocmodel.SQL.OrderDate),
		sqlutils.Full(stdocmodel.SQL.EducationStartDate),
		sqlutils.Full(groupmodel.SQL.ID),
		sqlutils.Full(groupmodel.SQL.Name),
		sqlutils.Full(groupmodel.SQL.IsBudget),
		sqlutils.Full(groupmodel.SQL.EducationForm),
		sqlutils.Full(groupmodel.SQL.EducationStartDate),
		sqlutils.Full(groupmodel.SQL.EducationYear),
		sqlutils.Full(groupmodel.SQL.EducationFinishDate),
		sqlutils.Full(groupmodel.SQL.StudyDepartmentId),
		sqlutils.Full(depmodel.SQL.ID),
		sqlutils.Full(depmodel.SQL.Name),
		sqlutils.Full(depmodel.SQL.ShortName),

		// inner join user table on student table by user id
		usermodel.UserTable,
		sqlutils.Full(studentmodel.SQL.UserID),
		sqlutils.Full(usermodel.SQL.ID),

		// inner join group table on student table by group id
		groupmodel.GroupTable,
		sqlutils.Full(studentmodel.SQL.GroupId),
		sqlutils.Full(groupmodel.SQL.ID),

		// inner join student document table on student table by student document id
		stdocmodel.StudentDocumentTable,
		sqlutils.Full(studentmodel.SQL.ID),
		sqlutils.Full(stdocmodel.SQL.StudentId),

		// inner join user roles on user table by user id
		usermodel.UserRolesTable,
		sqlutils.Full(usermodel.SQL_USER_ROLES.UserId),
		sqlutils.Full(usermodel.SQL.ID),

		// inner join roles table on user roles by role id
		usermodel.RolesTable,
		sqlutils.Full(usermodel.SQL_USER_ROLES.RoleId),
		sqlutils.Full(usermodel.SQL_ROLES.ID),

		// inner join study departments
		depmodel.StudyDepartmentTable,
		sqlutils.Full(depmodel.SQL_STUDY_DEP.ID),
		sqlutils.Full(groupmodel.SQL.StudyDepartmentId),

		// inner join departments
		depmodel.DepartmentTable,
		sqlutils.Full(depmodel.SQL.ID),
		sqlutils.Full(depmodel.SQL_STUDY_DEP.DepartmentId),

		// inserted in func query
		query,

		// group by tables
		sqlutils.Full(usermodel.SQL.ID),
		sqlutils.Full(studentmodel.SQL.ID),
		sqlutils.Full(stdocmodel.SQL.ID),
		sqlutils.Full(groupmodel.SQL.ID),
		sqlutils.Full(depmodel.SQL.ID),
	)

}

func scanStudent(row pgx.Row) (*studentmodel.StudentDTO, error) {
	user := new(usermodel.UserDTO)
	document := new(stdocmodel.StudentDocumentDTO)
	group := new(groupmodel.GroupDTO)
	deparment := new(depmodel.DepartmentDTO)
	student := studentmodel.NewStudentDTO(0, user, document, group, deparment)
	err := row.Scan(
		&student.ID,
		&user.ID,
		&user.Login, &user.Password, &user.Name, &user.Surname, &user.FatherName, &user.Email, &user.Roles,
		&document.ID, &document.StudentId,
		&document.DocNumber, &document.OrderNumber, &document.OrderDate, &document.EducationStartDate,
		&group.ID,
		&group.Name, &group.IsBudget, &group.EducationForm, &group.EducationStartDate,
		&group.EducationYear, &group.EducationFinishDate, &group.StudyDepartmentId,
		&deparment.ID, &deparment.Name, &deparment.ShortName,
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
	query := getStudentQuery("WHERE " + sqlutils.Full(studentmodel.SQL.ID) + " = $1")
	st, err := s.getStudentByQuery(ctx, query, id)
	if err != nil {
		return nil, studentError(err, amiderrors.NewCause("student by id query", "StudentById", _PROVIDER))
	}
	return st, nil
}

func (s *studentStorage) StudentByUserId(ctx context.Context, id uint64) (*studentmodel.StudentDTO, error) {
	st, err := s.getStudentByQuery(ctx, getStudentQuery("WHERE "+sqlutils.Full(usermodel.SQL.ID)+" = $1"), id)
	if err != nil {
		return nil, studentError(err, amiderrors.NewCause("student by id query", "StudentByUserId", _PROVIDER))
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
