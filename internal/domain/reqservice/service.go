package reqservice

import (
	"context"

	"github.com/amidgo/amiddocs/internal/models/depmodel"
	"github.com/amidgo/amiddocs/internal/models/doctypemodel"
	"github.com/amidgo/amiddocs/internal/models/doctypemodel/doctypefields"
	"github.com/amidgo/amiddocs/internal/models/reqmodel"
	"github.com/amidgo/amiddocs/internal/models/reqmodel/reqfields"
)

const _PROVIDER = "internal/domain/reqservice"

type departmentProvider interface {
	DepartmentById(ctx context.Context, id uint64) (*depmodel.DepartmentDTO, error)
}

type requestProvider interface {
	RequestById(ctx context.Context, reqId uint64) (*reqmodel.RequestDTO, error)
	LastRequestByUserId(ctx context.Context, userId uint64, dtype doctypefields.DocumentType) (*reqmodel.RequestDTO, error)
}
type requestRepository interface {
	InsertRequest(ctx context.Context, req *reqmodel.RequestDTO) (*reqmodel.RequestDTO, error)
	UpdateRequestStatus(ctx context.Context, reqId uint64, status reqfields.Status) error
	DeleteRequest(ctx context.Context, requestId uint64) error
}
type docTypeProvider interface {
	DocTypeByType(ctx context.Context, dtype doctypefields.DocumentType) (*doctypemodel.DocumentTypeDTO, error)
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
