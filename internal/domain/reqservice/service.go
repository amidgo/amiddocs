package reqservice

import (
	"context"

	"github.com/amidgo/amiddocs/internal/models/depmodel"
	"github.com/amidgo/amiddocs/internal/models/doctypemodel"
	"github.com/amidgo/amiddocs/internal/models/reqmodel"
	"github.com/amidgo/amiddocs/internal/models/reqmodel/reqfields"
)

const _PROVIDER = "internal/domain/reqservice"

type departmentProvider interface {
	DepartmentById(ctx context.Context, id uint64) (*depmodel.DepartmentDTO, error)
}

type requestProvider interface {
	LastRequestByUserId(ctx context.Context, userId uint64, dtype reqfields.DocumentType) (*reqmodel.RequestDTO, error)
}
type requestRepository interface {
	InsertRequest(ctx context.Context, req *reqmodel.RequestDTO) (*reqmodel.RequestDTO, error)
}
type docTypeProvider interface {
	DocTypeByType(ctx context.Context, dtype reqfields.DocumentType) (*doctypemodel.DocumentTypeDTO, error)
}

type requestService struct {
	depProv     departmentProvider
	reqProv     requestProvider
	reqRepo     requestRepository
	docTypeProv docTypeProvider
}

func New(
	depProv departmentProvider,
	reqProv requestProvider,
	reqRepo requestRepository,
	docTypeProv docTypeProvider,
) *requestService {
	return &requestService{
		depProv:     depProv,
		reqProv:     reqProv,
		reqRepo:     reqRepo,
		docTypeProv: docTypeProv,
	}
}