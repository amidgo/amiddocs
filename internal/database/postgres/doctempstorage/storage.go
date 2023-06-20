package doctempstorage

import (
	"database/sql"
	"errors"

	"github.com/amidgo/amiddocs/internal/errorutils/departmenterror"
	"github.com/amidgo/amiddocs/internal/errorutils/doctemperror"
	"github.com/amidgo/amiddocs/internal/errorutils/doctypeerror"
	"github.com/amidgo/amiddocs/internal/models/doctempmodel"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
	"github.com/amidgo/amiddocs/pkg/postgres"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

const _PROVIDER = "internal/database/postgres/doctempstorage"

type doctempStorage struct {
	p *postgres.Postgres
}

func New(p *postgres.Postgres) *doctempStorage {
	return &doctempStorage{p: p}
}

func templateError(err error, cause *amiderrors.Cause) error {
	pgerr := new(pgconn.PgError)
	if errors.As(err, &pgerr) {
		switch pgerr.ConstraintName {
		case doctempmodel.ForeignKey_DocumentTemplates__Departments:
			return departmenterror.DEPARTMENT_NOT_FOUND
		case doctempmodel.ForeignKey_DocumentTemplates__DocumentTypes:
			return doctypeerror.DOC_TYPE_NOT_FOUND
		}
	}
	switch {
	case amiderrors.Is(err, sql.ErrNoRows):
		return doctemperror.DOC_TEMP_NOT_FOUND
	case amiderrors.Is(err, pgx.ErrNoRows):
		return doctemperror.DOC_TEMP_NOT_FOUND
	default:
		return amiderrors.Wrap(err, cause)
	}
}
