package doctempservice

import (
	"context"

	"github.com/amidgo/amiddocs/internal/errorutils/doctemperror"
	"github.com/amidgo/amiddocs/internal/models/doctempmodel"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
)

func (s *doctempService) UploadTemplate(
	ctx context.Context,
	template *doctempmodel.DocumentTemplateDTO,
) (*doctempmodel.DocumentTemplateDTO, error) {
	_, err := s.depProv.DepartmentById(ctx, template.DepartmentID)
	if err != nil {
		return nil, err
	}
	err = s.docTypeProv.DocTypeExists(ctx, template.DocumentType)
	if err != nil {
		return nil, err
	}
	temp, err := s.docTempProv.DocTemp(ctx, template.DepartmentID, template.DocumentType)
	if amiderrors.Is(err, doctemperror.DOC_TEMP_NOT_FOUND) {
		temp, err = s.docTempServ.InsertDocTemp(ctx, template)
		if err != nil {
			return nil, err
		}
		return temp, nil
	}
	if err != nil {
		return nil, err
	}
	err = s.docTempServ.UpdateDocTemp(ctx, temp)
	if err != nil {
		return nil, err
	}
	return temp, nil
}
