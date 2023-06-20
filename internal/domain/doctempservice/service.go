package doctempservice

import (
	"context"
	"io"

	"github.com/amidgo/amiddocs/internal/models/depmodel"
	"github.com/amidgo/amiddocs/internal/models/doctempmodel"
	"github.com/amidgo/amiddocs/internal/models/doctypemodel"
	"github.com/amidgo/amiddocs/internal/models/doctypemodel/doctypefields"
)

const _PROVIDER = "internal/domain/doctempservice"

type departmentProvider interface {
	DepartmentById(ctx context.Context, id uint64) (*depmodel.DepartmentDTO, error)
}

type docTypeProvider interface {
	DocTypeByType(ctx context.Context, docType doctypefields.DocumentType) (*doctypemodel.DocumentTypeDTO, error)
}

type docTempProvider interface {
	DocumentTemplate(ctx context.Context, wr io.Writer, depId uint64, docType doctypefields.DocumentType) error
}

type docTempService interface {
	InsertDocTemp(ctx context.Context, template *doctempmodel.DocumentTemplateDTO) error
	UpdateDocTemp(ctx context.Context, template *doctempmodel.CreateTemplateDTO) error
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
