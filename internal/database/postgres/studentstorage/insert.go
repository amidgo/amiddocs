package studentstorage

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/amidgo/amiddocs/internal/database/postgres/stdocstorage"
	"github.com/amidgo/amiddocs/internal/database/postgres/userstorage"
	"github.com/amidgo/amiddocs/internal/errorutils/csverror"
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
	INSERT INTO %s (%s,%s) VALUES ($1,$2) RETURNING %s
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
		return nil, userstorage.UserError(err, amiderrors.NewCause("insert user", "InsertStudent", _PROVIDER))
	}
	err = addStudentRoleToUser(ctx, tx, student.User.ID)
	if err != nil {
		return nil, userstorage.UserError(err, amiderrors.NewCause("insert student role query", "InsertStudent", _PROVIDER))
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
		return nil, stdocstorage.StudentDocumentError(err, amiderrors.NewCause("insert student document", "InsertStudent", _PROVIDER))
	}
	err = tx.Commit(ctx)
	if err != nil {
		return nil, studentError(err, amiderrors.NewCause("commit tx", "InsertStudent", _PROVIDER))
	}
	return student, nil
}

var manyStudentInsertQuery = fmt.Sprintf(
	`
	WITH 
	user_insert as (
		INSERT INTO %s (%s,%s,%s,%s,%s,%s) VALUES ($1,$2,$3,$4,$5,$6) RETURNING %s as user_id
	),
	role_insert as (
		INSERT INTO %s (%s,%s) VALUES ((SELECT user_id FROM user_insert),$7)
	),
	student_insert as (
		INSERT INTO %s (%s,%s) VALUES ((SELECT user_id FROM user_insert),$8) RETURNING %s as student_id
	)
	INSERT INTO %s (%s,%s,%s,%s,%s) VALUES ((SELECT student_id FROM student_insert),$9,$10,$11,$12)
	`,
	// user insert
	usermodel.UserTable,

	// insertable values
	usermodel.SQL.Name,
	usermodel.SQL.Surname,
	usermodel.SQL.FatherName,
	usermodel.SQL.Login,
	usermodel.SQL.Email,
	usermodel.SQL.Password,
	// returning id
	usermodel.SQL.ID,

	// insert student role
	usermodel.UserRolesTable,
	usermodel.SQL_USER_ROLES.UserId,
	usermodel.SQL_USER_ROLES.RoleId,

	// insert into student table
	studentmodel.StudentTable,
	// insert values
	studentmodel.SQL.UserID,
	studentmodel.SQL.GroupId,
	// returning id
	studentmodel.SQL.ID,

	// insert into student document table
	stdocmodel.StudentDocumentTable,
	// insert values
	stdocmodel.SQL.StudentId,
	stdocmodel.SQL.DocNumber,
	stdocmodel.SQL.OrderNumber,
	stdocmodel.SQL.OrderDate,
	stdocmodel.SQL.EducationStartDate,
)

var selectStudentRole = fmt.Sprintf(
	"SELECT %s FROM %s WHERE %s = %s",
	usermodel.SQL_ROLES.Role,
	usermodel.RolesTable,
	usermodel.SQL_ROLES.ID,
	userfields.STUDENT,
)

func (s *studentStorage) InsertManyStudents(ctx context.Context, students []*studentmodel.StudentDTO) error {

	errCh := make(chan error)
	quit := make(chan struct{})

	tx, err := s.p.Pool.Begin(ctx)
	if err != nil {
		return studentError(err, amiderrors.NewCause("start tx", "InsertManyStudents", _PROVIDER))
	}
	defer tx.Rollback(ctx)
	roleId := uint64(0)
	err = tx.QueryRow(ctx, selectStudentRole).Scan(&roleId)
	if err != nil {
		return studentError(err, amiderrors.NewCause("get id of STUDENT role", "InsertManyStudents", _PROVIDER))
	}
	wg := new(sync.WaitGroup)
	wg.Add(len(students))
	for index, student := range students {
		go func(index int, student studentmodel.StudentDTO, tx pgx.Tx, roleId uint64) {
			defer wg.Done()
			err := insertSingleStudent(ctx, tx, &student, roleId)
			if err != nil {
				errCh <- &csverror.ErrorWithIndex{Err: errors.New("error while insert student"), Index: index}
				return
			}
		}(index, *student, tx, roleId)
	}
	go func() {
		wg.Wait()
		select {
		case <-quit:
		default:
			quit <- struct{}{}
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return errors.New("timeout")
		case err := <-errCh:
			if err != nil {
				return err
			}
		case <-quit:
			return nil
		}
	}
}

func insertSingleStudent(ctx context.Context, tx pgx.Tx, student *studentmodel.StudentDTO, roleId uint64) error {
	_, err := tx.Exec(ctx,
		manyStudentInsertQuery,
		// user insert values
		student.User.Name,
		student.User.Surname,
		student.User.FatherName,
		student.User.Login,
		student.User.Email,
		student.User.Password,
		// role insert values
		roleId,
		// student insert values
		student.Group.ID,
		// student document insert values
		student.Document.DocNumber,
		student.Document.OrderNumber,
		student.Document.OrderDate,
		student.Document.EducationStartDate,
	)
	if err != nil {
		return studentError(err, amiderrors.NewCause("insert single student", "insertSingleStudent", _PROVIDER))
	}
	return nil
}
