package userstorage

import (
	"context"

	"github.com/amidgo/amiddocs/internal/models/usermodel"
	"github.com/amidgo/amiddocs/internal/models/usermodel/userfields"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
	"github.com/jmoiron/sqlx"
)

type insertUserAmid struct {
	usrSt *UserStorage
	err   error
}

func (a *insertUserAmid) insertUserDTO(ctx context.Context, usr *usermodel.UserDTO, id *uint64) *sqlx.Rows {
	if a.err != nil {
		return nil
	}
	rows, err := a.usrSt.p.DB.
		NamedQueryContext(ctx,
			`INSERT INTO users (login,password,surname,name,father_name,email) 
		VALUES (:login, :password, :surname, :name, :father_name, :email) RETURNING id`,
			usr,
		)
	a.err = err
	return rows
}

func (a *insertUserAmid) scan(ctx context.Context, rows *sqlx.Rows, id *uint64) {
	if a.err != nil {
		return
	}
	for rows.Next() {
		err := rows.Scan(id)
		if err != nil {
			a.err = err
		}
	}
}

func (a *insertUserAmid) insertUserRoles(ctx context.Context, id uint64, roles []userfields.UserRole) {
	if a.err != nil {
		return
	}
	for _, r := range roles {
		_, err := a.usrSt.p.DB.NamedExecContext(ctx,
			`INSERT INTO roles (user_id,role) VALUES (:id, :role)`,
			map[string]interface{}{"id": id, "role": r})
		if err != nil {
			a.err = err
			return
		}
	}
}

func (u *UserStorage) InsertUser(ctx context.Context, usr *usermodel.UserDTO) (uint64, *amiderrors.ErrorResponse) {
	var id uint64
	amid := insertUserAmid{usrSt: u, err: nil}
	rows := amid.insertUserDTO(ctx, usr, &id)
	amid.scan(ctx, rows, &id)
	amid.insertUserRoles(ctx, id, usr.Roles)
	if amid.err != nil {
		return 0, amiderrors.NewInternalErrorResponse(amid.err)
	}
	return id, nil
}

func (u *UserStorage) DeleteUser(ctx context.Context, userId uint64) *amiderrors.ErrorResponse {
	_, err := u.p.Pool.
		Exec(ctx, "DELETE FROM users WHERE id = $1", userId)
	return amiderrors.NewInternalErrorResponse(err)
}
