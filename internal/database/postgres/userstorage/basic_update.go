package userstorage

import (
	"context"

	"github.com/amidgo/amiddocs/internal/models/usermodel/userfields"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
)

func (u *userStorage) UpdateName(ctx context.Context, userId uint64, name userfields.Name) error {
	_, err := u.p.DB.NamedExecContext(ctx,
		`UPDATE users SET name = :name WHERE id = :id `,
		map[string]interface{}{"name": name, "id": userId})
	if err != nil {
		return userError(err, amiderrors.NewCause("update name query", "UpdateName", _PROVIDER))
	}
	return nil
}

func (u *userStorage) UpdateSurname(ctx context.Context, userId uint64, surname userfields.Surname) error {
	_, err := u.p.DB.NamedExecContext(ctx,
		`UPDATE users SET surname = :surname WHERE id = :id `,
		map[string]interface{}{
			"surname": surname,
			"id":      userId,
		},
	)
	if err != nil {
		return userError(err, amiderrors.NewCause("update surname query", "UpdateSurname", _PROVIDER))
	}
	return nil
}

func (u *userStorage) UpdateFatherName(ctx context.Context, userId uint64, fatherName userfields.FatherName) error {
	_, err := u.p.DB.NamedExecContext(ctx,
		`UPDATE users SET father_name = :father_name WHERE id = :id `,
		map[string]interface{}{
			"father_name": fatherName,
			"id":          userId,
		},
	)
	if err != nil {
		return userError(err, amiderrors.NewCause("update father name query", "UpdateFatherName", _PROVIDER))
	}
	return nil
}

func (u *userStorage) UpdateLogin(ctx context.Context, userId uint64, login userfields.Login) error {
	_, err := u.p.DB.NamedExecContext(ctx,
		`UPDATE users SET login = :login WHERE id = :id `,
		map[string]interface{}{
			"login": login,
			"id":    userId,
		},
	)
	if err != nil {
		return userError(err, amiderrors.NewCause("update login query", "UpdateLogin", _PROVIDER))
	}
	return nil
}

func (u *userStorage) UpdatePassword(ctx context.Context, userId uint64, hashPassword string) error {
	_, err := u.p.DB.NamedExecContext(ctx,
		`UPDATE users SET password = :password WHERE id = :id `,
		map[string]interface{}{
			"password": hashPassword,
			"id":       userId,
		},
	)
	if err != nil {
		return userError(err, amiderrors.NewCause("update password query", "UpdatePassword", _PROVIDER))
	}
	return nil
}

func (u *userStorage) UpdateEmail(ctx context.Context, userId uint64, email userfields.Email) error {
	_, err := u.p.DB.NamedExecContext(ctx,
		`UPDATE users SET email = :email WHERE id = :id `,
		map[string]interface{}{
			"email": email,
			"id":    userId,
		},
	)
	if err != nil {
		return userError(err, amiderrors.NewCause("update email query", "UpdateEmail", _PROVIDER))
	}
	return nil
}

func (u *userStorage) RemoveRole(ctx context.Context, userId uint64, role userfields.Role) error {
	role_id := 0
	err := u.p.Pool.QueryRow(ctx, `SELECT id FROM roles WHERE role = $1`, role).Scan(&role_id)
	if err != nil {
		userError(err, amiderrors.NewCause("scan role_id query", "RemoveRole", _PROVIDER))
	}
	_, err = u.p.DB.NamedExecContext(ctx,
		`DELETE FROM roles where user_id = :user_id AND role_id = :role`,
		map[string]interface{}{
			"user_id": userId,
			"role_id": role_id,
		},
	)
	if err != nil {
		userError(err, amiderrors.NewCause("remove rol query", "RemoveRole", _PROVIDER))
	}
	return nil
}

func (u *userStorage) AddRole(ctx context.Context, userId uint64, role userfields.Role) error {
	_, err := u.p.DB.NamedExecContext(ctx,
		`INSERT INTO user_roles (user_id,role_id) VALUES :user_id, (SELECT roles.role FROM roles WHERE roles.role = :role)`,
		map[string]interface{}{
			"user_id": userId,
			"role_id": role,
		},
	)
	if err != nil {
		return userError(err, amiderrors.NewCause("add role query", "AddRole", _PROVIDER))
	}
	return nil
}
