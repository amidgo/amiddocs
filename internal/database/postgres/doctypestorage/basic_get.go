package doctypestorage

import (
	"context"

	"github.com/amidgo/amiddocs/internal/models/doctypemodel"
	"github.com/amidgo/amiddocs/internal/models/reqmodel/reqfields"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
)

func (s *docTypeStorage) DocTypeRefreshTime(ctx context.Context, docType reqfields.DocumentType) (uint8, error) {
	rTime := uint8(0)
	err := s.p.Pool.QueryRow(ctx,
		`SELECT refresh_time FROM document_types WHERE type = $1`, docType,
	).Scan(&rTime)
	if err != nil {
		return 0, docTypeError(err, amiderrors.NewCause("get refresh type query with scan", "DocTypeRefreshTime", _PROVIDER))
	}
	return rTime, nil
}

func (s *docTypeStorage) DocTypeExists(ctx context.Context, docType reqfields.DocumentType) error {
	t := new(reqfields.DocumentType)
	err := s.p.Pool.QueryRow(ctx, `SELECT type FROM document_types WHERE type = $1`, docType).Scan(t)
	if err != nil {
		return docTypeError(err, amiderrors.NewCause("get doc type query", "DocTypeExists", _PROVIDER))
	}
	return nil
}

func (s *docTypeStorage) DocTypeByType(ctx context.Context, dtype reqfields.DocumentType) (*doctypemodel.DocumentTypeDTO, error) {
	docType := new(doctypemodel.DocumentTypeDTO)
	err := s.p.Pool.QueryRow(ctx,
		`SELECT 
			document_types.id, document_types.type, document_types.refresh_time, roles.role
		FROM 
			document_types
		INNER JOIN roles ON roles.id = document_types.role_id
		WHERE document_types.type = $1`,
		dtype,
	).Scan(&docType.ID, &docType.Type, &docType.RefreshTime, &docType.Role)
	if err != nil {
		return nil, docTypeError(err, amiderrors.NewCause("get doc type by type query", "DocTypeByType", _PROVIDER))
	}
	return docType, nil
}
