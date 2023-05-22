package doctypestorage

import (
	"context"
	"fmt"

	"github.com/amidgo/amiddocs/internal/models/doctypemodel"
	"github.com/amidgo/amiddocs/internal/models/doctypemodel/doctypefields"
	"github.com/amidgo/amiddocs/internal/models/usermodel"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
	"github.com/amidgo/amiddocs/pkg/sqlutils"
)

var (
	refreshTimeQuery = fmt.Sprintf(
		`SELECT %s FROM %s WHERE %s = $1`,
		doctypemodel.SQL.RefreshTime,
		doctypemodel.DocTypeTable,
		doctypemodel.SQL.Type,
	)

	docTypeExistQuery = fmt.Sprintf(
		`SELECT %s FROM %s WHERE %s = $1`,
		doctypemodel.SQL.Type,
		doctypemodel.DocTypeTable,
		doctypemodel.SQL.Type,
	)

	docTypeByTypeQuery = fmt.Sprintf(
		`SELECT 
			%s, %s, %s, array_agg(%s)
		FROM 
			%s
		INNER JOIN %s ON %s = %s
		INNER JOIN %s ON %s = %s
		WHERE %s = $1
		GROUP BY %s
		`,

		// selectable variables
		sqlutils.Full(doctypemodel.SQL.ID),
		sqlutils.Full(doctypemodel.SQL.Type),
		sqlutils.Full(doctypemodel.SQL.RefreshTime),
		sqlutils.Full(usermodel.SQL_ROLES.Role),

		// from doc type table
		doctypemodel.DocTypeTable,

		// inner join doc type roles on document type id
		doctypemodel.DocTypeRoleTable,
		sqlutils.Full(doctypemodel.SQL_ROLES.DocumentTypeId),
		sqlutils.Full(doctypemodel.SQL.ID),

		// inner join on roles by role id
		usermodel.RolesTable,
		sqlutils.Full(usermodel.SQL_ROLES.ID),
		sqlutils.Full(doctypemodel.SQL_ROLES.RoleId),

		// where doc type = $1
		sqlutils.Full(doctypemodel.SQL.Type),

		// group by columns
		sqlutils.Full(doctypemodel.SQL.ID),
	)
)

func (s *docTypeStorage) DocTypeRefreshTime(ctx context.Context, docType doctypefields.DocumentType) (uint8, error) {
	rTime := uint8(0)
	err := s.p.Pool.QueryRow(ctx,
		refreshTimeQuery,
		docType,
	).Scan(&rTime)
	if err != nil {
		return 0, docTypeError(err, amiderrors.NewCause("get refresh type query with scan", "DocTypeRefreshTime", _PROVIDER))
	}
	return rTime, nil
}

func (s *docTypeStorage) DocTypeExists(ctx context.Context, docType doctypefields.DocumentType) error {
	t := doctypefields.DocumentType("")
	err := s.p.Pool.QueryRow(ctx, docTypeExistQuery, docType).Scan(&t)
	if err != nil {
		return docTypeError(err, amiderrors.NewCause("get doc type query", "DocTypeExists", _PROVIDER))
	}
	return nil
}

func (s *docTypeStorage) DocTypeByType(ctx context.Context, dtype doctypefields.DocumentType) (*doctypemodel.DocumentTypeDTO, error) {
	docType := new(doctypemodel.DocumentTypeDTO)
	err := s.p.Pool.QueryRow(ctx,
		docTypeByTypeQuery,
		dtype,
	).Scan(&docType.ID, &docType.Type, &docType.RefreshTime, &docType.Roles)
	if err != nil {
		return nil, docTypeError(err, amiderrors.NewCause("get doc type by type query", "DocTypeByType", _PROVIDER))
	}
	return docType, nil
}

// func (s *docTypeStorage) AllDocTypes(ctx context.Context) ([]*doctypemodel.DocumentTypeDTO, error) {

// }
