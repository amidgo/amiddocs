package doctempstorage

import (
	"context"

	"github.com/amidgo/amiddocs/internal/models/doctempmodel"
	"github.com/amidgo/amiddocs/internal/models/reqmodel/reqfields"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
)

func (s *doctempStorage) DocTemp(
	ctx context.Context,
	depID uint64,
	docType reqfields.DocumentType,
) (*doctempmodel.DocumentTemplateDTO, error) {
	docTemplateDTO := new(doctempmodel.DocumentTemplateDTO)
	docTemplateDTO.DepartmentID = depID
	docTemplateDTO.DocumentType = docType
	err := s.p.Pool.QueryRow(
		ctx,
		`SELECT document_templates.data 
			FROM document_templates 
			INNER JOIN document_types
			ON document_templates.document_type = document_types.id
			WHERE document_templates.department_id = $1 AND document_types.type = $2
			GROUP BY document_types.id
			`,
		depID, docType,
	).Scan(&docTemplateDTO.Document)
	if err != nil {
		return nil, tempalateError(err, amiderrors.NewCause("get doc temp query", "DocTemplate", _PROVIDER))
	}
	return docTemplateDTO, nil
}
