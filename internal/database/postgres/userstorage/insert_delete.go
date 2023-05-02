package userstorage

import (
	"context"

	"github.com/amidgo/amiddocs/internal/models/usermodel"
	"github.com/amidgo/amiddocs/internal/models/usermodel/userfields"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

// create table if not exists users (
//     id bigserial primary key,
//     name varchar(40) not null,
//     surname varchar(60) not null,
//     father_name varchar(40) not null,
//     login varchar(100) not null unique,
//     email varchar(100),
//     password varchar(200) not null
// );

func insertUserDTO(ctx context.Context, tx pgx.Tx, usr *usermodel.UserDTO, id *uint64) error {
	err := tx.QueryRow(
		ctx,
		`INSERT INTO users 
			(name,surname,father_name,login,email,password)
		VALUES ($1,$2,$3,$4,$5,$6)
		RETURNING id`,
		usr.Name,
		usr.Surname,
		pgtype.Text{String: string(usr.FatherName), Valid: usr.FatherName != ""},
		usr.Login,
		pgtype.Text{String: string(usr.Email), Valid: usr.Email != ""},
		usr.Password,
	).Scan(id)
	if err != nil {
		return userError(err, amiderrors.NewCause("insert user dto", "InsertUser", _PROVIDER))
	}
	return nil
}

func insertUserRoles(ctx context.Context, tx pgx.Tx, id uint64, roles []userfields.Role) error {
	rows := make([][]interface{}, 0)
	role_id := 0
	for _, r := range roles {
		err := tx.QueryRow(ctx, `SELECT id FROM roles WHERE role = $1`, r).Scan(&role_id)
		if err != nil {
			return userError(err, amiderrors.NewCause("scan role id query", "insertUserRoles", _PROVIDER))
		}
		rows = append(rows, []interface{}{id, role_id})
	}
	_, err := tx.CopyFrom(
		ctx,
		pgx.Identifier{"user_roles"},
		[]string{"user_id", "role_id"},
		pgx.CopyFromRows(rows),
	)
	if err != nil {
		return userError(err, amiderrors.NewCause("insert user roles", "InsertUser", _PROVIDER))
	}
	return nil
}

func (u *userStorage) InsertUser(ctx context.Context, usr *usermodel.UserDTO) (*usermodel.UserDTO, error) {
	tx, err := u.p.Pool.Begin(ctx)
	if err != nil {
		return nil, userError(err, amiderrors.NewCause("begin tx", "InsertUser", _PROVIDER))
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
		return nil, userError(err, amiderrors.NewCause("commit transaction", "InsertUser", _PROVIDER))
	}
	return usr, nil
}

func (u *userStorage) DeleteUser(ctx context.Context, userId uint64) error {
	_, err := u.p.DB.
		ExecContext(ctx, "DELETE FROM users WHERE id = $1", userId)
	return userError(err, amiderrors.NewCause("delete user query", "DeleteUser", _PROVIDER))
}
