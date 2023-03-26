package userstorage

import (
	"context"

	"github.com/amidgo/amiddocs/internal/models/usermodel/userfields"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
)

func (u *UserStorage) UpdateName(ctx context.Context, userId uint64, userName userfields.Name) *amiderrors.ErrorResponse {
	_, err := u.p.Pool.
		Exec(ctx, `UPDATE users SET name = $1 WHERE id = $2 `, userName, userId)
	return amiderrors.NewInternalErrorResponse(err)
}

func (u *UserStorage) UpdateSurname(ctx context.Context, userId uint64, userSurname userfields.Surname) *amiderrors.ErrorResponse {
	_, err := u.p.Pool.
		Exec(ctx, `UPDATE users SET surname = $1 WHERE id = $2 `, userSurname, userId)
	return amiderrors.NewInternalErrorResponse(err)
}

func (u *UserStorage) UpdateFatherName(ctx context.Context, userId uint64, userFatherName userfields.FatherName) *amiderrors.ErrorResponse {
	_, err := u.p.Pool.
		Exec(ctx, `UPDATE users SET "fatherName" = $1 WHERE id = $2 `, userFatherName, userId)
	return amiderrors.NewInternalErrorResponse(err)
}

func (u *UserStorage) UpdateLogin(ctx context.Context, userId uint64, login userfields.Login) *amiderrors.ErrorResponse {
	_, err := u.p.Pool.
		Exec(ctx, `UPDATE users SET login = $1 WHERE id = $2 `, login, userId)
	return amiderrors.NewInternalErrorResponse(err)
}

func (u *UserStorage) UpdatePassword(ctx context.Context, userId uint64, hashPassword string) *amiderrors.ErrorResponse {
	_, err := u.p.Pool.
		Exec(ctx, `UPDATE users SET password = $1 WHERE id = $2 `, hashPassword, userId)
	return amiderrors.NewInternalErrorResponse(err)
}

func (u *UserStorage) UpdateEmail(ctx context.Context, userId uint64, email userfields.Email) *amiderrors.ErrorResponse {
	_, err := u.p.Pool.
		Exec(ctx, `UPDATE users SET email = $1 WHERE id = $2 `, email, userId)
	return amiderrors.NewInternalErrorResponse(err)
}

func (u *UserStorage) RemoveRole(ctx context.Context, userId uint64, role userfields.UserRole) *amiderrors.ErrorResponse {
	_, err := u.p.Pool.
		Exec(ctx, `DELETE FROM roles where "userId" = $1 AND role = $2`, userId, role)
	return amiderrors.NewInternalErrorResponse(err)
}

func (u *UserStorage) AddRole(ctx context.Context, userId uint64, role userfields.UserRole) *amiderrors.ErrorResponse {
	_, err := u.p.Pool.
		Exec(ctx, `INSERT INTO roles ("userId",role) VALUES $1, $2`, userId, role)
	return amiderrors.NewInternalErrorResponse(err)
}
