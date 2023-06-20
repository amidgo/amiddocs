package userstorage

import (
	"database/sql"
	"errors"

	"github.com/amidgo/amiddocs/internal/errorutils/usererror"
	"github.com/amidgo/amiddocs/internal/models/usermodel"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
	"github.com/amidgo/amiddocs/pkg/postgres"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

const _PROVIDER = "/internal/database/postgres/userstorage"

type userStorage struct {
	p *postgres.Postgres
}

func New(p *postgres.Postgres) *userStorage {
	return &userStorage{p: p}
}

func UserError(err error, cause *amiderrors.Cause) error {
	pgerr := new(pgconn.PgError)
	if errors.As(err, &pgerr) {
		switch pgerr.ConstraintName {
		case usermodel.UsersLoginUniqueConstraint:
			return usererror.LOGIN_ALREADY_EXIST
		case usermodel.UsersEmailUniqueConstraint:
			return usererror.EMAIL_ALREADY_EXIST
		}
	}
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return usererror.NOT_FOUND
	case errors.Is(err, pgx.ErrNoRows):
		return usererror.NOT_FOUND
	default:
		return amiderrors.Wrap(err, cause)
	}
}
