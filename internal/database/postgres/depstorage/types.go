package depstorage

import (
	"context"
	"fmt"

	"github.com/amidgo/amiddocs/internal/models/depmodel"
	"github.com/amidgo/amiddocs/internal/models/doctempmodel"
	"github.com/amidgo/amiddocs/internal/models/doctypemodel"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
	"github.com/amidgo/amiddocs/pkg/sqlutils"
	"github.com/jackc/pgx/v5"
)

var (
	depTypesQuery = fmt.Sprintf(
		`
		SELECT %s,%s,%s,array_agg(%s)
			FROM %s
		INNER JOIN %s ON %s = %s
		INNER JOIN %s ON %s = %s
		
		GROUP BY %s, %s
		`,
		// selectable values
		sqlutils.Full(depmodel.SQL.ID),
		sqlutils.Full(depmodel.SQL.Name),
		sqlutils.Full(depmodel.SQL.ShortName),
		sqlutils.Full(doctypemodel.SQL.Type),
		// from departments
		depmodel.DepartmentTable,

		// inner join document templates
		doctempmodel.DocTempTable,
		sqlutils.Full(doctempmodel.SQL.DepartmentId),
		sqlutils.Full(depmodel.SQL.ID),

		// inner join doc types
		doctypemodel.DocTypeTable,
		sqlutils.Full(doctypemodel.SQL.ID),
		sqlutils.Full(doctempmodel.SQL.DocumentTypeId),

		// group by columns
		sqlutils.Full(depmodel.SQL.ID),
		sqlutils.Full(doctypemodel.SQL.ID),
	)
)

func scanDepTypes(row pgx.Row, depTypes *depmodel.DepartmentTypes) error {
	dep := new(depmodel.DepartmentDTO)
	depTypes.Department = dep
	return row.Scan(
		&dep.ID,
		&dep.Name,
		&dep.ShortName,
		&depTypes.Types,
	)
}

func scanDepTypesList(rows pgx.Rows) ([]*depmodel.DepartmentTypes, error) {
	depTypesList := make([]*depmodel.DepartmentTypes, 0)
	for rows.Next() {
		dep := new(depmodel.DepartmentTypes)
		err := scanDepTypes(rows, dep)
		if err != nil {
			return nil, departmentError(err, amiderrors.NewCause("scan dep from row", "DepartmentListWithTypes", _PROVIDER))
		}
		depTypesList = append(depTypesList, dep)
	}
	return depTypesList, nil
}

func (s *departmentStorage) DepartmentListWithTypes(ctx context.Context) ([]*depmodel.DepartmentTypes, error) {
	fmt.Println(depTypesQuery)
	rows, err := s.p.Pool.Query(ctx, depTypesQuery)
	if err != nil {
		return nil, departmentError(err, amiderrors.NewCause("get dep types rows by query", "DepartmentListWithTypes", _PROVIDER))
	}
	depList, err := scanDepTypesList(rows)
	if err != nil {
		return nil, departmentError(err, amiderrors.NewCause("scan dep types list", "DepartmentListWithTypes", _PROVIDER))
	}
	return depList, nil
}
