package doctempstorage

import (
	"context"
	"fmt"

	"github.com/amidgo/amiddocs/internal/models/doctempmodel"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
)

var (
	doctempInsertQuery = fmt.Sprintf(
		`
		INSERT INTO 
			%s
		(%s, %s,%s)
			VALUES (
				$1,
				$2,
				$3
			)
		`,
		doctempmodel.DocTempTable,

		// insert values
		doctempmodel.SQL.DepartmentId,
		doctempmodel.SQL.DocumentTypeId,
		doctempmodel.SQL.Data,
	)
)

func (s *doctempStorage) InsertDocTemp(
	ctx context.Context,
	template *doctempmodel.DocumentTemplateDTO,
) error {
	_, err := s.p.Pool.Exec(
		ctx,
		doctempInsertQuery,
		template.DepartmentID, template.DocumentTypeID, template.Document,
	)
	if err != nil {
		return templateError(err, amiderrors.NewCause("insert doc temp into storage", "InsertDocTemp", _PROVIDER))
	}
	return nil
}
