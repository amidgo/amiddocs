package depstorage

import (
	"database/sql"
	"errors"

	"github.com/amidgo/amiddocs/internal/errorutils/departmenterror"
	"github.com/amidgo/amiddocs/internal/models/depmodel"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
	"github.com/amidgo/amiddocs/pkg/postgres"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

const _PROVIDER = "internal/database/postgres/depstorage"

type departmentStorage struct {
	p *postgres.Postgres
}

func New(p *postgres.Postgres) *departmentStorage {
	return &departmentStorage{p: p}
}

func departmentError(err error, cause *amiderrors.Cause) error {
	var pgerr *pgconn.PgError
	if errors.As(err, &pgerr) {
		switch pgerr.ConstraintName {
		case depmodel.DepartmentNameUniqueConstraint:
			return departmenterror.NAME_EXIST
		case depmodel.DepartmentShortNameUniqueConstraint:
			return departmenterror.SHORT_NAME_EXIST
		}
	}
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return departmenterror.DEPARTMENT_NOT_FOUND
	case errors.Is(err, pgx.ErrNoRows):
		return departmenterror.DEPARTMENT_NOT_FOUND
	default:
		return amiderrors.Wrap(err, cause)
	}
}

func studyDepartmentError(err error, cause *amiderrors.Cause) error {
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return departmenterror.STUDY_DEPARTMENT_NOT_FOUND
	case errors.Is(err, pgx.ErrNoRows):
		return departmenterror.STUDY_DEPARTMENT_NOT_FOUND
	default:
		return amiderrors.Wrap(err, cause)
	}
}
