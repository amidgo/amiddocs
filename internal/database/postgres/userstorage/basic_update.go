package userstorage

import (
	"context"
	"fmt"

	"github.com/amidgo/amiddocs/internal/models/usermodel"
	"github.com/amidgo/amiddocs/internal/models/usermodel/userfields"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
	"github.com/amidgo/amiddocs/pkg/sqlutils"
)

var (
	updateNameQuery     = updateQuery(usermodel.SQL.Name)
	updateSurnameQuery  = updateQuery(usermodel.SQL.Surname)
	updateFatherName    = updateQuery(usermodel.SQL.FatherName)
	updateEmailQuery    = updateQuery(usermodel.SQL.Email)
	updateLoginQuery    = updateQuery(usermodel.SQL.Login)
	updatePasswordQuery = updateQuery(usermodel.SQL.Password)
)

var (
	addRoleQuery = fmt.Sprintf(
		`INSERT INTO %s
		(%s,%s)
		VALUES (
			$1,
			(SELECT %s FROM %s WHERE %s = $2)
		)`,
		// selectable variables
		usermodel.UserRolesTable,
		usermodel.SQL_USER_ROLES.UserId,
		usermodel.SQL_USER_ROLES.RoleId,

		// select id from roles where role = $1
		usermodel.SQL_ROLES.ID,
		usermodel.RolesTable,
		usermodel.SQL_ROLES.Role,
	)

	removeRoleQuery = fmt.Sprintf(
		`
		DELETE FROM %s
		INNER JOIN %s on %s = %s
		WHERE %s = $1 %s = $2
		`,
		// delete from user roles
		usermodel.UserRolesTable,
		// inner join on roles by role id
		usermodel.RolesTable,
		sqlutils.Full(usermodel.SQL_ROLES.ID),
		sqlutils.Full(usermodel.SQL_USER_ROLES.RoleId),
		// where user id = $2
		sqlutils.Full(usermodel.SQL_USER_ROLES.UserId),
		sqlutils.Full(usermodel.SQL_ROLES.Role),
	)
)

func updateQuery(column sqlutils.Column) string {
	return fmt.Sprintf(
		"UPDATE %s SET %s = $1 WHERE %s = $2",
		usermodel.UserTable,
		column,
		usermodel.SQL.ID,
	)
}

func (u *userStorage) UpdateName(ctx context.Context, userId uint64, name userfields.Name) error {
	_, err := u.p.Pool.Exec(ctx,
		updateNameQuery,
		name, userId,
	)
	if err != nil {
		return UserError(err, amiderrors.NewCause("update name query", "UpdateName", _PROVIDER))
	}
	return nil
}

func (u *userStorage) UpdateSurname(ctx context.Context, userId uint64, surname userfields.Surname) error {
	_, err := u.p.Pool.Exec(ctx,
		updateSurnameQuery,
		surname, userId,
	)
	if err != nil {
		return UserError(err, amiderrors.NewCause("update surname query", "UpdateSurname", _PROVIDER))
	}
	return nil
}

func (u *userStorage) UpdateFatherName(ctx context.Context, userId uint64, fatherName userfields.FatherName) error {
	_, err := u.p.Pool.Exec(ctx,
		updateFatherName,
		fatherName, userId,
	)
	if err != nil {
		return UserError(err, amiderrors.NewCause("update father name query", "UpdateFatherName", _PROVIDER))
	}
	return nil
}

func (u *userStorage) UpdateLogin(ctx context.Context, userId uint64, login userfields.Login) error {
	_, err := u.p.Pool.Exec(ctx,
		updateLoginQuery,
		login, userId,
	)
	if err != nil {
		return UserError(err, amiderrors.NewCause("update login query", "UpdateLogin", _PROVIDER))
	}
	return nil
}

func (u *userStorage) UpdatePassword(ctx context.Context, userId uint64, hashPassword string) error {
	_, err := u.p.Pool.Exec(ctx,
		updatePasswordQuery,
		hashPassword, userId,
	)
	if err != nil {
		return UserError(err, amiderrors.NewCause("update password query", "UpdatePassword", _PROVIDER))
	}
	return nil
}

func (u *userStorage) UpdateEmail(ctx context.Context, userId uint64, email userfields.Email) error {
	_, err := u.p.Pool.Exec(ctx,
		updateEmailQuery,
		email, userId,
	)
	if err != nil {
		return UserError(err, amiderrors.NewCause("update email query", "UpdateEmail", _PROVIDER))
	}
	return nil
}

func (u *userStorage) RemoveRole(ctx context.Context, userId uint64, role userfields.Role) error {
	_, err := u.p.Pool.Exec(ctx,
		removeRoleQuery,
		userId, role,
	)
	if err != nil {
		UserError(err, amiderrors.NewCause("remove rol query", "RemoveRole", _PROVIDER))
	}
	return nil
}

func (u *userStorage) AddRole(ctx context.Context, userId uint64, role userfields.Role) error {
	_, err := u.p.Pool.Exec(ctx,
		addRoleQuery,
		userId, role,
	)
	if err != nil {
		return UserError(err, amiderrors.NewCause("add role query", "AddRole", _PROVIDER))
	}
	return nil
}
