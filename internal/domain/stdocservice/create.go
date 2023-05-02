package stdocservice

import (
	"context"

	"github.com/amidgo/amiddocs/internal/errorutils/stdocerror"
	"github.com/amidgo/amiddocs/internal/models/stdocmodel"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
)

func (s *studentDocumentService) CreateDocument(
	ctx context.Context,
	doc *stdocmodel.StudentDocumentDTO,
) (*stdocmodel.StudentDocumentDTO, error) {

	_, err := s.stDocRep.DocumentByDocNumber(ctx, doc.DocNumber)
	if !amiderrors.Is(err, stdocerror.DOC_NOT_FOUND) {
		return nil, stdocerror.DOC_NUMBER_EXIST
	}

	_, err = s.stDocRep.DocumentByOrderNumber(ctx, doc.OrderNumber)
	if !amiderrors.Is(err, stdocerror.DOC_NOT_FOUND) {
		return nil, stdocerror.ORDER_NUMBER_EXIST
	}

	stdoc, err := s.stDocRep.InsertDocument(ctx, doc)
	if err != nil {
		return nil, err
	}

	return stdoc, nil
}
