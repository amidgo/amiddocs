package doctempstorage

import (
	"context"
	"fmt"
	"io"

	"github.com/amidgo/amiddocs/internal/models/doctempmodel"
	"github.com/amidgo/amiddocs/internal/models/doctypemodel"
	"github.com/amidgo/amiddocs/internal/models/doctypemodel/doctypefields"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
	"github.com/amidgo/amiddocs/pkg/sqlutils"
)

var (
	documentTemplateQuery = fmt.Sprintf(
		`
		SELECT %s 
		FROM %s 
		INNER JOIN %s
			ON %s = %s
		WHERE %s = $1 AND %s = $2
		GROUP BY %s,%s, %s`,

		// select data
		sqlutils.Full(doctempmodel.SQL.Data),

		// from doc temp table
		doctempmodel.DocTempTable,

		// inner join document types by document type id
		doctypemodel.DocTypeTable,
		sqlutils.Full(doctempmodel.SQL.DocumentTypeId),
		sqlutils.Full(doctypemodel.SQL.ID),

		// where departmentid = %1 and doc type = $2
		sqlutils.Full(doctempmodel.SQL.DepartmentId),
		sqlutils.Full(doctypemodel.SQL.Type),

		// group by columns
		sqlutils.Full(doctypemodel.SQL.ID),
		sqlutils.Full(doctempmodel.SQL.DocumentTypeId),
		sqlutils.Full(doctempmodel.SQL.DepartmentId),
	)
)

func (s *doctempStorage) DocumentTemplate(
	ctx context.Context,
	wr io.Writer,
	depID uint64,
	docType doctypefields.DocumentType,
) error {
	docTemplateDTO := new(doctempmodel.DocumentTemplateDTO)
	docTemplateDTO.DepartmentID = depID
	docTemplateDTO.DocumentType = docType
	err := s.p.Pool.QueryRow(
		ctx,
		documentTemplateQuery,
		depID, docType,
	).Scan(&docTemplateDTO.Document)
	if err != nil {
		return tempalateError(err, amiderrors.NewCause("get doc temp query", "DocTemplate", _PROVIDER))
	}
	_, err = wr.Write(docTemplateDTO.Document)
	if err != nil {
		return tempalateError(err, amiderrors.NewCause("write in io.Writer", "DocumentTemplate", _PROVIDER))
	}
	return nil
}
