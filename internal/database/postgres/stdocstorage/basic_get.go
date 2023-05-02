package stdocstorage

import (
	"context"

	"github.com/amidgo/amiddocs/internal/models/stdocmodel"
	"github.com/amidgo/amiddocs/internal/models/stdocmodel/stdocfields"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
)

const _BASIC_GET_QUERY = `SELECT id,doc_number,order_number,order_date,study_start_date FROM student_documents `

func (st *studentDocStorage) getDocumentByQuery(
	ctx context.Context,
	query string,
	args ...interface{},
) (*stdocmodel.StudentDocumentDTO, error) {
	doc := new(stdocmodel.StudentDocumentDTO)
	err := st.p.DB.GetContext(ctx, doc, query, args...)
	if err != nil {
		return nil, err
	}
	return doc, nil
}

func (st *studentDocStorage) DocumentById(ctx context.Context, id uint64) (*stdocmodel.StudentDocumentDTO, error) {
	doc, err := st.getDocumentByQuery(ctx, _BASIC_GET_QUERY+`WHERE id = $1`, id)
	if err != nil {
		return nil, studentDocumentError(err, amiderrors.NewCause("document by id query", "DocumentById", _PROVIDER))
	}
	return doc, nil
}

func (st *studentDocStorage) DocumentByDocNumber(ctx context.Context, docNumber stdocfields.DocNumber) (*stdocmodel.StudentDocumentDTO, error) {
	doc, err := st.getDocumentByQuery(ctx, _BASIC_GET_QUERY+`WHERE doc_number = $1`, docNumber)
	if err != nil {
		return nil, studentDocumentError(err, amiderrors.NewCause("document by doc number", "DocumentByDocNumber", _PROVIDER))
	}
	return doc, nil
}

func (st *studentDocStorage) DocumentByOrderNumber(ctx context.Context, orderNumber stdocfields.OrderNumber) (*stdocmodel.StudentDocumentDTO, error) {
	doc, err := st.getDocumentByQuery(ctx, _BASIC_GET_QUERY+`WHERE order_number = $1`, orderNumber)
	if err != nil {
		return nil, studentDocumentError(err, amiderrors.NewCause("document by order number", "DocumentByOrderNumber", _PROVIDER))
	}
	return doc, nil
}
