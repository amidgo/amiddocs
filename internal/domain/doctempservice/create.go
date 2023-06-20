package doctempservice

import (
	"bytes"
	"context"

	"github.com/amidgo/amiddocs/internal/errorutils/doctemperror"
	"github.com/amidgo/amiddocs/internal/models/doctempmodel"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
)

func (s *doctempService) SaveTemplate(
	ctx context.Context,
	template *doctempmodel.CreateTemplateDTO,
) error {
	err := s.docTempProv.DocumentTemplate(ctx, &bytes.Buffer{}, template.DepartmentID, template.DocumentType)
	if amiderrors.Is(err, doctemperror.DOC_TEMP_NOT_FOUND) {
		err = s.InsertDocTemp(ctx, template)
		if err != nil {
			return err
		}
		return nil
	}
	if err != nil {
		return err
	}
	err = s.docTempServ.UpdateDocTemp(ctx, template)
	if err != nil {
		return err
	}
	return nil
}

func (s *doctempService) InsertDocTemp(ctx context.Context, template *doctempmodel.CreateTemplateDTO) error {
	docType, err := s.docTypeProv.DocTypeByType(ctx, template.DocumentType)
	if err != nil {
		return err
	}
	docTempDTO := doctempmodel.NewDocumentTemplateDTO(template.DepartmentID, docType.ID, template.Document)
	err = s.docTempServ.InsertDocTemp(ctx, docTempDTO)
	if err != nil {
		return err
	}
	return nil
}
