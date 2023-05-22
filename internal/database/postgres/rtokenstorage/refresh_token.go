package rtokenstorage

import (
	"database/sql"
	"errors"
	"time"

	"github.com/amidgo/amiddocs/internal/errorutils/tokenerror"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
	"github.com/amidgo/amiddocs/pkg/postgres"
	"github.com/jackc/pgx/v5"
)

const _PROVIDER = "internal/database/postgres/rtokenstorage"

type refreshTokenStorage struct {
	rtoken_time time.Duration
	p           *postgres.Postgres
}

func New(rtoken_time time.Duration, p *postgres.Postgres) *refreshTokenStorage {
	return &refreshTokenStorage{rtoken_time: rtoken_time, p: p}
}

func rtokenError(err error, cause *amiderrors.Cause) error {
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return tokenerror.TOKEN_NOT_FOUND
	case errors.Is(err, pgx.ErrNoRows):
		return tokenerror.TOKEN_NOT_FOUND
	default:
		return amiderrors.NewInternalErrorResponse(err, cause)
	}
}
