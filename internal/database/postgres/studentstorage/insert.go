package studentstorage

import (
	"context"
	"fmt"

	"github.com/amidgo/amiddocs/internal/models/stdocmodel"
	"github.com/amidgo/amiddocs/internal/models/studentmodel"
	"github.com/amidgo/amiddocs/internal/models/usermodel"
	"github.com/amidgo/amiddocs/internal/models/usermodel/userfields"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

var (
	insertUserQuery = fmt.Sprintf(
		`INSERT INTO %s (%s,%s,%s,%s,%s,%s) VALUES ($1,$2,$3,$4,$5,$6) RETURNING %s`,
		usermodel.UserTable,
		usermodel.SQL.Name,
		usermodel.SQL.Surname,
		usermodel.SQL.FatherName,
		usermodel.SQL.Login,
		usermodel.SQL.Email,
		usermodel.SQL.Password,
		usermodel.SQL.ID,
	)
	insertStudentDocumentQuery = fmt.Sprintf(
		`INSERT INTO %s (%s,%s,%s,%s,%s) VALUES ($1,$2,$3,$4,$5) RETURNING %s`,
		// insert into student document table
		stdocmodel.StudentDocumentTable,
		// insert values
		stdocmodel.SQL.StudentId,
		stdocmodel.SQL.DocNumber,
		stdocmodel.SQL.OrderNumber,
		stdocmodel.SQL.OrderDate,
		stdocmodel.SQL.EducationStartDate,
		// returning id
		stdocmodel.SQL.ID,
	)

	addRoleToUserQuery = fmt.Sprintf(
		`INSERT INTO %s (%s,%s) VALUES ($1,(SELECT %s FROM %s WHERE %s = $2))`,
		usermodel.UserRolesTable,
		usermodel.SQL_USER_ROLES.UserId,
		usermodel.SQL_USER_ROLES.RoleId,
		usermodel.SQL_ROLES.ID,
		usermodel.RolesTable,
		usermodel.SQL_ROLES.Role,
	)

	insertStudentQuery = fmt.Sprintf(`
	INSERT INTO %s
		(%s,%s)
		VALUES ($1,$2)
	RETURNING %s
	`,
		// insert into student table
		studentmodel.StudentTable,
		// insert values
		studentmodel.SQL.UserID,
		studentmodel.SQL.GroupId,
		// returning id
		studentmodel.SQL.ID,
	)
)

func insertUser(ctx context.Context, tx pgx.Tx, user *usermodel.UserDTO) error {
	err := tx.QueryRow(
		ctx,
		insertUserQuery,
		user.Name,
		user.Surname,
		pgtype.Text{String: string(user.FatherName), Valid: user.FatherName != ""},
		user.Login,
		pgtype.Text{String: string(user.Email), Valid: user.Email != ""},
		user.Password,
	).Scan(&user.ID)
	return err
}

func insertStDoc(ctx context.Context, tx pgx.Tx, document *stdocmodel.StudentDocumentDTO) error {
	err := tx.QueryRow(
		ctx,
		insertStudentDocumentQuery,
		document.StudentId,
		document.DocNumber,
		document.OrderNumber,
		document.OrderDate,
		document.EducationStartDate,
	).Scan(&document.ID)
	return err
}

func addStudentRoleToUser(ctx context.Context, tx pgx.Tx, userId uint64) error {
	_, err := tx.Exec(
		ctx,
		addRoleToUserQuery,
		userId, userfields.STUDENT,
	)
	return err
}

func (s *studentStorage) InsertStudent(ctx context.Context, student *studentmodel.StudentDTO) (*studentmodel.StudentDTO, error) {
	tx, err := s.p.Pool.Begin(ctx)
	if err != nil {
		return nil, studentError(err, amiderrors.NewCause("begin transaction", "InsertStudent", _PROVIDER))
	}
	defer tx.Rollback(ctx)
	err = insertUser(ctx, tx, student.User)
	if err != nil {
		return nil, studentError(err, amiderrors.NewCause("insert user", "InsertStudent", _PROVIDER))
	}

	err = addStudentRoleToUser(ctx, tx, student.User.ID)
	if err != nil {
		return nil, studentError(err, amiderrors.NewCause("insert student role query", "InsertStudent", _PROVIDER))
	}
	err = tx.QueryRow(ctx,
		insertStudentQuery,
		student.User.ID,
		student.Group.ID,
	).Scan(&student.ID)
	if err != nil {
		return nil, studentError(err, amiderrors.NewCause("student insert query", "InsertStudent", _PROVIDER))
	}
	student.Document.StudentId = student.ID
	err = insertStDoc(ctx, tx, student.Document)
	if err != nil {
		return nil, studentError(err, amiderrors.NewCause("insert student document", "InsertStudent", _PROVIDER))
	}
	err = tx.Commit(ctx)
	if err != nil {
		return nil, studentError(err, amiderrors.NewCause("commit tx", "InsertStudent", _PROVIDER))
	}
	return student, nil
}
