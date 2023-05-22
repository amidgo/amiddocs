package stdocstorage

import (
	"context"
	"fmt"

	"github.com/amidgo/amiddocs/internal/models/stdocmodel"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
)

var (
	insertDocQuery = fmt.Sprintf(
		`INSERT INTO %s (%s,%s,%s,%s)
		VALUES ($1,$2,$3,$4)
		RETURNING %s`,
		stdocmodel.StudentDocumentTable,

		stdocmodel.SQL.DocNumber,
		stdocmodel.SQL.OrderNumber,
		stdocmodel.SQL.OrderDate,
		stdocmodel.SQL.EducationStartDate,

		stdocmodel.SQL.ID,
	)
)

func (st *studentDocStorage) InsertDocument(ctx context.Context, doc *stdocmodel.StudentDocumentDTO) (*stdocmodel.StudentDocumentDTO, error) {
	err := st.p.Pool.QueryRow(ctx,
		insertDocQuery,
		doc.DocNumber, doc.OrderNumber, doc.OrderDate, doc.EducationStartDate,
	).Scan(&doc.ID)
	if err != nil {
		return nil, studentDocumentError(err, amiderrors.NewCause("insert student query", "InsertDocument", _PROVIDER))
	}
	return doc, nil
}
