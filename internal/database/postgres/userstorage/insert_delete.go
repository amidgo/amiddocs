package userstorage

import (
	"context"
	"fmt"

	"github.com/amidgo/amiddocs/internal/models/usermodel"
	"github.com/amidgo/amiddocs/internal/models/usermodel/userfields"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

var (
	deleteUserQuery = fmt.Sprintf(
		`DELETE FROM %s WHERE %s = $1`,
		usermodel.UserTable,
		usermodel.SQL.ID,
	)
	insertUserQuery = fmt.Sprintf(
		`INSERT INTO %s 
			(%s,%s,%s,%s,%s,%s)
		VALUES ($1,$2,$3,$4,$5,$6)
		RETURNING %s`,
		usermodel.UserTable,

		usermodel.SQL.Name,
		usermodel.SQL.Surname,
		usermodel.SQL.FatherName,
		usermodel.SQL.Login,
		usermodel.SQL.Email,
		usermodel.SQL.Password,

		usermodel.SQL.ID,
	)
)

func insertUserDTO(ctx context.Context, tx pgx.Tx, usr *usermodel.UserDTO, id *uint64) error {
	err := tx.QueryRow(
		ctx,
		insertUserQuery,
		usr.Name,
		usr.Surname,
		pgtype.Text{String: string(usr.FatherName), Valid: usr.FatherName != ""},
		usr.Login,
		pgtype.Text{String: string(usr.Email), Valid: usr.Email != ""},
		usr.Password,
	).Scan(id)
	if err != nil {
		return UserError(err, amiderrors.NewCause("insert user dto", "InsertUser", _PROVIDER))
	}
	return nil
}

func insertUserRoles(ctx context.Context, tx pgx.Tx, id uint64, roles []userfields.Role) error {
	rows := make([][]interface{}, 0)
	role_id := 0
	query := fmt.Sprintf(`SELECT %s FROM %s WHERE %s = $1`, usermodel.SQL_ROLES.ID, usermodel.RolesTable, usermodel.SQL_ROLES.Role)
	for _, r := range roles {
		err := tx.QueryRow(ctx, query, r).Scan(&role_id)
		if err != nil {
			return UserError(err, amiderrors.NewCause("scan role id query", "insertUserRoles", _PROVIDER))
		}
		rows = append(rows, []interface{}{id, role_id})
	}
	_, err := tx.CopyFrom(
		ctx,
		pgx.Identifier{usermodel.UserRolesTable},
		[]string{usermodel.SQL_USER_ROLES.UserId.String(), usermodel.SQL_USER_ROLES.RoleId.String()},
		pgx.CopyFromRows(rows),
	)
	if err != nil {
		return UserError(err, amiderrors.NewCause("insert user roles", "InsertUser", _PROVIDER))
	}
	return nil
}

func (u *userStorage) InsertUser(ctx context.Context, usr *usermodel.UserDTO) (*usermodel.UserDTO, error) {
	tx, err := u.p.Pool.Begin(ctx)
	if err != nil {
		return nil, UserError(err, amiderrors.NewCause("begin tx", "InsertUser", _PROVIDER))
	}
	defer tx.Rollback(ctx)
	err = insertUserDTO(ctx, tx, usr, &usr.ID)
	if err != nil {
		return nil, err
	}
	err = insertUserRoles(ctx, tx, usr.ID, usr.Roles)
	if err != nil {
		return nil, err
	}
	err = tx.Commit(ctx)
	if err != nil {
		return nil, UserError(err, amiderrors.NewCause("commit transaction", "InsertUser", _PROVIDER))
	}
	return usr, nil
}

func (u *userStorage) DeleteUser(ctx context.Context, userId uint64) error {
	_, err := u.p.DB.
		ExecContext(ctx, deleteUserQuery, userId)
	return UserError(err, amiderrors.NewCause("delete user query", "DeleteUser", _PROVIDER))
}
