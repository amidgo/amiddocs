package stdocservice

import (
	"context"

	"github.com/amidgo/amiddocs/internal/models/stdocmodel"
	"github.com/amidgo/amiddocs/internal/models/stdocmodel/stdocfields"
)

type studentDocumentRepository interface {
	InsertDocument(ctx context.Context, doc *stdocmodel.StudentDocumentDTO) (*stdocmodel.StudentDocumentDTO, error)
	DocumentById(ctx context.Context, id uint64) (*stdocmodel.StudentDocumentDTO, error)
	DocumentByDocNumber(ctx context.Context, docNumber stdocfields.DocNumber) (*stdocmodel.StudentDocumentDTO, error)
	DocumentByOrderNumber(ctx context.Context, orderNumber stdocfields.OrderNumber) (*stdocmodel.StudentDocumentDTO, error)
}

type studentDocumentService struct {
	stDocRep studentDocumentRepository
}

func New(stDocRep studentDocumentRepository) *studentDocumentService {
	return &studentDocumentService{stDocRep: stDocRep}
}
