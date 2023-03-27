package userstorage

import (
	"context"

	"github.com/amidgo/amiddocs/internal/models/usermodel/userfields"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
)

func (u *UserStorage) UpdateName(ctx context.Context, userId uint64, name userfields.Name) *amiderrors.ErrorResponse {
	_, err := u.p.DB.NamedExecContext(ctx,
		`UPDATE users SET name = :name WHERE id = :id `,
		map[string]interface{}{"name": name, "id": userId})
	return amiderrors.NewInternalErrorResponse(err)
}

func (u *UserStorage) UpdateSurname(ctx context.Context, userId uint64, surname userfields.Surname) *amiderrors.ErrorResponse {
	_, err := u.p.DB.NamedExecContext(ctx,
		`UPDATE users SET surname = :surname WHERE id = :id `,
		map[string]interface{}{
			"surname": surname,
			"id":      userId,
		},
	)
	return amiderrors.NewInternalErrorResponse(err)
}

func (u *UserStorage) UpdateFatherName(ctx context.Context, userId uint64, fatherName userfields.FatherName) *amiderrors.ErrorResponse {
	_, err := u.p.DB.NamedExecContext(ctx,
		`UPDATE users SET father_name = :father_name WHERE id = :id `,
		map[string]interface{}{
			"father_name": fatherName,
			"id":          userId,
		},
	)
	return amiderrors.NewInternalErrorResponse(err)
}

func (u *UserStorage) UpdateLogin(ctx context.Context, userId uint64, login userfields.Login) *amiderrors.ErrorResponse {
	_, err := u.p.DB.NamedExecContext(ctx,
		`UPDATE users SET login = :login WHERE id = :id `,
		map[string]interface{}{
			"login": login,
			"id":    userId,
		},
	)
	return amiderrors.NewInternalErrorResponse(err)
}

func (u *UserStorage) UpdatePassword(ctx context.Context, userId uint64, hashPassword string) *amiderrors.ErrorResponse {
	_, err := u.p.DB.NamedExecContext(ctx,
		`UPDATE users SET password = :password WHERE id = :id `,
		map[string]interface{}{
			"password": hashPassword,
			"id":       userId,
		},
	)
	return amiderrors.NewInternalErrorResponse(err)
}

func (u *UserStorage) UpdateEmail(ctx context.Context, userId uint64, email userfields.Email) *amiderrors.ErrorResponse {
	_, err := u.p.DB.NamedExecContext(ctx,
		`UPDATE users SET email = :email WHERE id = :id `,
		map[string]interface{}{
			"email": email,
			"id":    userId,
		},
	)
	return amiderrors.NewInternalErrorResponse(err)
}

func (u *UserStorage) RemoveRole(ctx context.Context, userId uint64, role userfields.UserRole) *amiderrors.ErrorResponse {
	_, err := u.p.DB.NamedExecContext(ctx,
		`DELETE FROM roles where user_id = :user_id AND role = :role`,
		map[string]interface{}{
			"user_id": userId,
			"role":    role,
		},
	)
	return amiderrors.NewInternalErrorResponse(err)
}

func (u *UserStorage) AddRole(ctx context.Context, userId uint64, role userfields.UserRole) *amiderrors.ErrorResponse {
	_, err := u.p.DB.NamedExecContext(ctx,
		`INSERT INTO roles (user_id,role) VALUES :user_id, :role`,
		map[string]interface{}{
			"user_id": userId,
			"role":    role,
		},
	)
	return amiderrors.NewInternalErrorResponse(err)
}
