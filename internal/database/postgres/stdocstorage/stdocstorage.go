package stdocstorage

import (
	"database/sql"
	"errors"

	"github.com/amidgo/amiddocs/internal/errorutils/stdocerror"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
	"github.com/amidgo/amiddocs/pkg/postgres"
	"github.com/jackc/pgx/v5"
)

const _PROVIDER = "internal/database/postgres/stdocstorage"

type studentDocStorage struct {
	p *postgres.Postgres
}

func New(p *postgres.Postgres) *studentDocStorage {
	return &studentDocStorage{p: p}
}

func studentDocumentError(err error, cause *amiderrors.Cause) error {
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return stdocerror.DOC_NOT_FOUND
	case errors.Is(err, pgx.ErrNoRows):
		return stdocerror.DOC_NOT_FOUND
	default:
		return amiderrors.NewInternalErrorResponse(err, cause)
	}
}
