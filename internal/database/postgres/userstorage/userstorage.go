package userstorage

import (
	"database/sql"
	"errors"
	"fmt"

	usererrorutils "github.com/amidgo/amiddocs/internal/errorutils/usererror"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
	"github.com/amidgo/amiddocs/pkg/postgres"
	"github.com/jackc/pgx/v5"
)

const _PROVIDER = "/internal/database/postgres/userstorage"

type userStorage struct {
	p *postgres.Postgres
}

func New(p *postgres.Postgres) *userStorage {
	return &userStorage{p: p}
}

func userError(err error, cause *amiderrors.Cause) error {
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return usererrorutils.NOT_FOUND
	case errors.Is(err, pgx.ErrNoRows):
		return usererrorutils.NOT_FOUND
	default:
		fmt.Printf("Type error is %T", err)
		return amiderrors.NewInternalErrorResponse(err, cause)
	}
}
