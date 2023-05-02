package doctempstorage

import (
	"database/sql"

	"github.com/amidgo/amiddocs/internal/errorutils/doctemperror"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
	"github.com/amidgo/amiddocs/pkg/postgres"
	"github.com/jackc/pgx/v5"
)

const _PROVIDER = "internal/database/postgres/doctempstorage"

type doctempStorage struct {
	p *postgres.Postgres
}

func New(p *postgres.Postgres) *doctempStorage {
	return &doctempStorage{p: p}
}

func tempalateError(err error, cause *amiderrors.Cause) error {
	switch {
	case amiderrors.Is(err, sql.ErrNoRows):
		return doctemperror.DOC_TEMP_NOT_FOUND
	case amiderrors.Is(err, pgx.ErrNoRows):
		return doctemperror.DOC_TEMP_NOT_FOUND
	default:
		return amiderrors.NewInternalErrorResponse(err, cause)
	}
}
