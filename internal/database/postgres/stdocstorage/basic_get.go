package stdocstorage

import (
	"context"
	"fmt"

	"github.com/amidgo/amiddocs/internal/models/stdocmodel"
	"github.com/amidgo/amiddocs/internal/models/stdocmodel/stdocfields"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
	"github.com/jackc/pgx/v5"
)

var (
	getStudentDocQuery = fmt.Sprintf(
		`SELECT %s,%s,%s,%s,%s FROM %s `,
		stdocmodel.SQL.ID,
		stdocmodel.SQL.DocNumber,
		stdocmodel.SQL.OrderNumber,
		stdocmodel.SQL.OrderDate,
		stdocmodel.SQL.EducationStartDate,
		stdocmodel.StudentDocumentTable,
	)
)

func (st *studentDocStorage) getDocumentByQuery(
	ctx context.Context,
	query string,
	args ...interface{},
) (*stdocmodel.StudentDocumentDTO, error) {
	doc := new(stdocmodel.StudentDocumentDTO)
	row := st.p.Pool.QueryRow(ctx, query, args...)
	err := scanStudentDoc(row, doc)
	if err != nil {
		return nil, err
	}
	return doc, nil
}

func scanStudentDoc(row pgx.Row, doc *stdocmodel.StudentDocumentDTO) error {
	return row.Scan(
		&doc.ID,
		&doc.DocNumber,
		&doc.OrderNumber,
		&doc.OrderDate,
		&doc.EducationStartDate,
	)
}

func (st *studentDocStorage) DocumentById(ctx context.Context, id uint64) (*stdocmodel.StudentDocumentDTO, error) {
	doc, err := st.getDocumentByQuery(ctx, getStudentDocQuery+`WHERE `+stdocmodel.SQL.ID.String()+` = $1`, id)
	if err != nil {
		return nil, studentDocumentError(err, amiderrors.NewCause("document by id query", "DocumentById", _PROVIDER))
	}
	return doc, nil
}

func (st *studentDocStorage) DocumentByDocNumber(ctx context.Context, docNumber stdocfields.DocNumber) (*stdocmodel.StudentDocumentDTO, error) {
	doc, err := st.getDocumentByQuery(ctx, getStudentDocQuery+`WHERE `+stdocmodel.SQL.DocNumber.String()+` = $1`, docNumber)
	if err != nil {
		return nil, studentDocumentError(err, amiderrors.NewCause("document by doc number", "DocumentByDocNumber", _PROVIDER))
	}
	return doc, nil
}

func (st *studentDocStorage) DocumentByOrderNumber(ctx context.Context, orderNumber stdocfields.OrderNumber) (*stdocmodel.StudentDocumentDTO, error) {
	doc, err := st.getDocumentByQuery(ctx, getStudentDocQuery+`WHERE `+stdocmodel.SQL.OrderNumber.String()+` = $1`, orderNumber)
	if err != nil {
		return nil, studentDocumentError(err, amiderrors.NewCause("document by order number", "DocumentByOrderNumber", _PROVIDER))
	}
	return doc, nil
}
