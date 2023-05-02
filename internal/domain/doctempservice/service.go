package doctempservice

import (
	"context"

	"github.com/amidgo/amiddocs/internal/models/depmodel"
	"github.com/amidgo/amiddocs/internal/models/doctempmodel"
	"github.com/amidgo/amiddocs/internal/models/reqmodel/reqfields"
)

const _PROVIDER = "internal/domain/doctempservice"

type departmentProvider interface {
	DepartmentById(ctx context.Context, id uint64) (*depmodel.DepartmentDTO, error)
}

type docTypeProvider interface {
	DocTypeExists(ctx context.Context, docType reqfields.DocumentType) error
}

type docTempProvider interface {
	DocTemp(ctx context.Context, depId uint64, docType reqfields.DocumentType) (*doctempmodel.DocumentTemplateDTO, error)
}

type docTempService interface {
	InsertDocTemp(ctx context.Context, template *doctempmodel.DocumentTemplateDTO) (*doctempmodel.DocumentTemplateDTO, error)
	UpdateDocTemp(ctx context.Context, template *doctempmodel.DocumentTemplateDTO) error
}

type doctempService struct {
	depProv     departmentProvider
	docTypeProv docTypeProvider
	docTempServ docTempService
	docTempProv docTempProvider
}

func New(
	depProv departmentProvider,
	docTypeProv docTypeProvider,
	docTempProv docTempProvider,
	docTempServ docTempService,
) *doctempService {
	return &doctempService{
		depProv:     depProv,
		docTypeProv: docTypeProv,
		docTempServ: docTempServ,
		docTempProv: docTempProv,
	}
}
