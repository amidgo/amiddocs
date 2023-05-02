package depstorage

import (
	"database/sql"
	"errors"

	"github.com/amidgo/amiddocs/internal/errorutils/departmenterror"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
	"github.com/amidgo/amiddocs/pkg/postgres"
	"github.com/jackc/pgx/v5"
)

const _PROVIDER = "internal/database/postgres/depstorage"

type departmentStorage struct {
	p *postgres.Postgres
}

func New(p *postgres.Postgres) *departmentStorage {
	return &departmentStorage{p: p}
}

func departmentError(err error, cause *amiderrors.Cause) error {
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return departmenterror.DEPARMENT_NOT_FOUND
	case errors.Is(err, pgx.ErrNoRows):
		return departmenterror.DEPARMENT_NOT_FOUND
	default:
		return amiderrors.NewInternalErrorResponse(err, cause)
	}
}
