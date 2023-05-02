package studentstorage

import (
	"context"

	"github.com/amidgo/amiddocs/internal/models/stdocmodel"
	"github.com/amidgo/amiddocs/internal/models/studentmodel"
	"github.com/amidgo/amiddocs/internal/models/usermodel"
	"github.com/amidgo/amiddocs/internal/models/usermodel/userfields"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

func insertUser(ctx context.Context, tx pgx.Tx, user *usermodel.UserDTO) error {
	err := tx.QueryRow(
		ctx,
		`INSERT INTO users (name,surname,father_name,login,email,password)
	 	VALUES ($1,$2,$3,$4,$5,$6) RETURNING id`,
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
		`INSERT INTO student_documents (doc_number,order_number,order_date,study_start_date)
		VALUES ($1,$2,$3,$4) RETURNING id`,
		document.DocNumber,
		document.OrderNumber,
		document.OrderDate,
		document.StudyStartDate,
	).Scan(&document.ID)
	return err
}

func addStudentRoleToUser(ctx context.Context, tx pgx.Tx, userId uint64) error {
	_, err := tx.Exec(
		ctx,
		`INSERT INTO user_roles (user_id, role_id) VALUES ($1,(SELECT roles.id FROM roles WHERE role = $2))`,
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
	err = insertStDoc(ctx, tx, student.Document)
	if err != nil {
		return nil, studentError(err, amiderrors.NewCause("insert student document", "InsertStudent", _PROVIDER))
	}
	err = addStudentRoleToUser(ctx, tx, student.User.ID)
	if err != nil {
		return nil, studentError(err, amiderrors.NewCause("insert student role query", "InsertStudent", _PROVIDER))
	}
	err = tx.QueryRow(ctx,
		`INSERT INTO students (user_id,group_id,student_document_id) VALUES ($1,$2,$3) RETURNING id`,
		student.User.ID, student.Group.ID, student.Document.ID).Scan(&student.ID)
	if err != nil {
		return nil, studentError(err, amiderrors.NewCause("student insert query", "InsertStudent", _PROVIDER))
	}
	err = tx.Commit(ctx)
	if err != nil {
		return nil, studentError(err, amiderrors.NewCause("commit tx", "InsertStudent", _PROVIDER))
	}
	return student, nil
}
