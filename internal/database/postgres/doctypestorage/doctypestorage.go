package doctypestorage

import (
	"database/sql"
	"errors"

	"github.com/amidgo/amiddocs/internal/errorutils/doctypeerror"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
	"github.com/amidgo/amiddocs/pkg/postgres"
	"github.com/jackc/pgx/v5"
)

const _PROVIDER = "internal/database/postgres/doctypestorage"

type docTypeStorage struct {
	p *postgres.Postgres
}

func New(p *postgres.Postgres) *docTypeStorage {
	return &docTypeStorage{p: p}
}

func DocTypeError(err error, cause *amiderrors.Cause) error {
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return doctypeerror.DOC_TYPE_NOT_FOUND
	case errors.Is(err, pgx.ErrNoRows):
		return doctypeerror.DOC_TYPE_NOT_FOUND
	default:
		return amiderrors.Wrap(err, cause)
	}
}
