package doctempstorage

import (
	"context"
	"fmt"

	"github.com/amidgo/amiddocs/internal/models/doctempmodel"
	"github.com/amidgo/amiddocs/internal/models/doctypemodel"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
	"github.com/amidgo/amiddocs/pkg/sqlutils"
)

var (
	updateDocTempQuery = fmt.Sprintf(
		`
		UPDATE %s 
			SET %s = $1
		FROM %s
		WHERE 
			%s = $2 
					AND
			%s = $3
					AND
			%s = %s
		`,

		// update table
		doctempmodel.DocTempTable,

		doctempmodel.SQL.Data,

		// from doc type table
		doctypemodel.DocTypeTable,

		// where doctemp.department_id = $2
		sqlutils.Full(doctempmodel.SQL.DepartmentId),
		// and doctypes.type = $3
		sqlutils.Full(doctypemodel.SQL.Type),

		sqlutils.Full(doctempmodel.SQL.DocumentTypeId),
		sqlutils.Full(doctypemodel.SQL.ID),
	)
)

func (s *doctempStorage) UpdateDocTemp(ctx context.Context, doctemp *doctempmodel.CreateTemplateDTO) error {
	fmt.Println(updateDocTempQuery)
	_, err := s.p.Pool.Exec(ctx,
		updateDocTempQuery,
		doctemp.Document, doctemp.DepartmentID, doctemp.DocumentType,
	)
	if err != nil {
		return templateError(err, amiderrors.NewCause("udpate doc tepmlate", "UpdateDocTemp", _PROVIDER))
	}
	return nil
}
