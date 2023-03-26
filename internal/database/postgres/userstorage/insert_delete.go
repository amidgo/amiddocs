package userstorage

import (
	"context"

	"github.com/amidgo/amiddocs/internal/models/usermodel"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
)

func (u *UserStorage) InsertUser(ctx context.Context, usr *usermodel.UserDTO) (uint64, *amiderrors.ErrorResponse) {
	var id uint64
	err := u.p.Pool.
		QueryRow(ctx,
			`INSERT INTO users (login,password,surname,name,"fatherName",email) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`,
			usr.Login, usr.Password, usr.Surname, usr.Surname, usr.Name, usr.Email,
		).Scan(&id)
	if err != nil {
		return 0, amiderrors.NewInternalErrorResponse(err)
	}
	for _, r := range usr.Roles {
		_, err = u.p.Pool.Exec(ctx,
			`INSERT INTO roles ("userId",role) VALUES ($1, $2)`, id, r)
		if err != nil {
			return 0, amiderrors.NewInternalErrorResponse(err)
		}
	}

	return id, nil
}

func (u *UserStorage) DeleteUser(ctx context.Context, userId uint64) *amiderrors.ErrorResponse {
	_, err := u.p.Pool.
		Exec(ctx, "DELETE FROM users WHERE id = $1", userId)
	if err != nil {
		return amiderrors.NewInternalErrorResponse(err)
	}
	return nil
}
