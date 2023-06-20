package stdocstorage

import (
	"database/sql"
	"errors"

	"github.com/amidgo/amiddocs/internal/errorutils/stdocerror"
	"github.com/amidgo/amiddocs/internal/errorutils/studenterror"
	"github.com/amidgo/amiddocs/internal/models/stdocmodel"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
	"github.com/amidgo/amiddocs/pkg/postgres"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

const _PROVIDER = "internal/database/postgres/stdocstorage"

type studentDocStorage struct {
	p *postgres.Postgres
}

func New(p *postgres.Postgres) *studentDocStorage {
	return &studentDocStorage{p: p}
}

func StudentDocumentError(err error, cause *amiderrors.Cause) error {
	pgerr := new(pgconn.PgError)
	if errors.As(err, &pgerr) {
		switch pgerr.ConstraintName {
		case stdocmodel.ForeignKey_StudentDocuments__Students:
			return studenterror.STUDENT_NOT_FOUND
		case stdocmodel.StudentDocumentsNumberUniqueConstraint:
			return stdocerror.DOC_NUMBER_EXIST
		case stdocmodel.StudentDocumentsStudentIdUniqueConstraint:
			return stdocerror.STUDENT_ALREADY_HAVE_DOCUMENT
		}
	}
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return stdocerror.DOC_NOT_FOUND
	case errors.Is(err, pgx.ErrNoRows):
		return stdocerror.DOC_NOT_FOUND
	default:
		return amiderrors.Wrap(err, cause)
	}
}
