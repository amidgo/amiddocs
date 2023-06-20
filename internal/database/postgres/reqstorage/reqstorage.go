package reqstorage

import (
	"database/sql"
	"errors"

	"github.com/amidgo/amiddocs/internal/errorutils/departmenterror"
	"github.com/amidgo/amiddocs/internal/errorutils/doctypeerror"
	"github.com/amidgo/amiddocs/internal/errorutils/reqerror"
	"github.com/amidgo/amiddocs/internal/errorutils/usererror"
	"github.com/amidgo/amiddocs/internal/models/reqmodel"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
	"github.com/amidgo/amiddocs/pkg/postgres"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

const _PROVIDER = "internal/database/postgres/reqstorage"

type requestStorage struct {
	p *postgres.Postgres
}

func New(p *postgres.Postgres) *requestStorage {
	return &requestStorage{p: p}
}

func requestError(err error, cause *amiderrors.Cause) error {
	pgerr := new(pgconn.PgError)
	if errors.As(err, &pgerr) {
		switch pgerr.ConstraintName {
		case reqmodel.ForeignKey_Requests__Departments:
			return departmenterror.DEPARTMENT_NOT_FOUND
		case reqmodel.ForeignKey_Requests__DocumentTypes:
			return doctypeerror.DOC_TYPE_NOT_FOUND
		case reqmodel.ForeignKey_Requests__Users:
			return usererror.NOT_FOUND
		}
	}
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return reqerror.REQ_NOT_FOUND
	case errors.Is(err, pgx.ErrNoRows):
		return reqerror.REQ_NOT_FOUND
	default:
		return amiderrors.Wrap(err, cause)
	}
}
