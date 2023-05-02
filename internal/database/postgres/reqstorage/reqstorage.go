package reqstorage

import (
	"database/sql"
	"errors"

	"github.com/amidgo/amiddocs/internal/errorutils/reqerror"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
	"github.com/amidgo/amiddocs/pkg/postgres"
	"github.com/jackc/pgx/v5"
)

const _PROVIDER = "internal/database/postgres/reqstorage"

type requestStorage struct {
	p *postgres.Postgres
}

func New(p *postgres.Postgres) *requestStorage {
	return &requestStorage{p: p}
}

func requestError(err error, cause *amiderrors.Cause) error {
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return reqerror.REQ_NOT_FOUND
	case errors.Is(err, pgx.ErrNoRows):
		return reqerror.REQ_NOT_FOUND
	default:
		return amiderrors.NewInternalErrorResponse(err, cause)
	}
}
