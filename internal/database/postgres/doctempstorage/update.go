package doctempstorage

import (
	"context"

	"github.com/amidgo/amiddocs/internal/models/doctempmodel"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
)

func (s *doctempStorage) UpdateDocTemp(ctx context.Context, doctemp *doctempmodel.DocumentTemplateDTO) error {
	_, err := s.p.Pool.Exec(ctx,
		`
		UPDATE document_templates 
		INNER JOIN document_types
		ON document_templates.document_type = document_types.id
		SET document_templates.data = $1
		WHERE document_templates.department_id = $2 AND document_types.type = $3
		`,
		doctemp.Document, doctemp.DepartmentID, doctemp.DocumentType,
	)
	if err != nil {
		return tempalateError(err, amiderrors.NewCause("udpate doc tepmlate", "UpdateDocTemp", _PROVIDER))
	}
	return nil
}
