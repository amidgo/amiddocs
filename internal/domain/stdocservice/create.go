package stdocservice

import (
	"context"

	"github.com/amidgo/amiddocs/internal/models/stdocmodel"
)

func (s *studentDocumentService) CreateDocument(
	ctx context.Context,
	doc *stdocmodel.StudentDocumentDTO,
) (*stdocmodel.StudentDocumentDTO, error) {
	stdoc, err := s.stDocRep.InsertDocument(ctx, doc)
	if err != nil {
		return nil, err
	}

	return stdoc, nil
}
