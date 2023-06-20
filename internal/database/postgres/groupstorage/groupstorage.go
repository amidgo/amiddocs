package groupstorage

import (
	"database/sql"
	"errors"

	"github.com/amidgo/amiddocs/internal/errorutils/departmenterror"
	"github.com/amidgo/amiddocs/internal/errorutils/grouperror"
	"github.com/amidgo/amiddocs/internal/models/groupmodel"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
	"github.com/amidgo/amiddocs/pkg/postgres"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

const _PROVIDER = "internal/database/postgres/groupstorage"

type groupStorage struct {
	p *postgres.Postgres
}

func New(p *postgres.Postgres) *groupStorage {
	return &groupStorage{p: p}
}

func groupError(err error, cause *amiderrors.Cause) error {
	pgerr := new(pgconn.PgError)
	if errors.As(err, &pgerr) {
		switch pgerr.ConstraintName {
		case groupmodel.GroupNameUniqueConstraint:
			return grouperror.GROUP_NAME_ALREADY_EXIST
		case groupmodel.ForeignKey_Groups__StudyDepartments:
			return departmenterror.STUDY_DEPARTMENT_NOT_FOUND
		}
	}
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return grouperror.GROUP_NOT_FOUND
	case errors.Is(err, pgx.ErrNoRows):
		return grouperror.GROUP_NOT_FOUND
	default:
		return amiderrors.Wrap(err, cause)
	}
}
