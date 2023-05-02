package doctempstorage

import (
	"context"

	"github.com/amidgo/amiddocs/internal/models/doctempmodel"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
)

func (s *doctempStorage) InsertDocTemp(
	ctx context.Context,
	template *doctempmodel.DocumentTemplateDTO,
) (*doctempmodel.DocumentTemplateDTO, error) {
	_, err := s.p.Pool.Exec(
		ctx,
		`INSERT INTO 
			document_templates 
		(department_id, document_type_id,data)
			VALUES ($1,(SELECT id FROM document_types WHERE type = $2),$3) `,
		template.DepartmentID, template.DocumentType, template.Document,
	)
	if err != nil {
		return nil, tempalateError(err, amiderrors.NewCause("insert doc temp into storage", "InsertDocTemp", _PROVIDER))
	}
	return template, nil
}
