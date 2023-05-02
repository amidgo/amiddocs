package studentstorage

import (
	"database/sql"
	"errors"

	"github.com/amidgo/amiddocs/internal/errorutils/studenterror"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
	"github.com/amidgo/amiddocs/pkg/postgres"
	"github.com/jackc/pgx/v5"
)

const _PROVIDER = "/internal/database/postgres/studentstorage"

type studentStorage struct {
	p *postgres.Postgres
}

func New(p *postgres.Postgres) *studentStorage {
	return &studentStorage{p: p}
}

func studentError(err error, cause *amiderrors.Cause) error {
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return studenterror.STUDENT_NOT_FOUND
	case errors.Is(err, pgx.ErrNoRows):
		return studenterror.STUDENT_NOT_FOUND
	case err == nil:
		return nil
	default:
		return amiderrors.NewInternalErrorResponse(err, cause)
	}
}
