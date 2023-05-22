package depstorage

import (
	"context"
	"fmt"

	"github.com/amidgo/amiddocs/internal/models/depmodel"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
)

var (
	insertDepQuery = fmt.Sprintf(
		`
		INSERT INTO %s 
		(%s, %s)
		VALUES
			($1,$2)
		RETURNING %s`,
		depmodel.DepartmentTable,

		depmodel.SQL.Name,
		depmodel.SQL.ShortName,
		depmodel.SQL.ID,
	)
)

func (s *departmentStorage) InsertDepartment(ctx context.Context, dep *depmodel.DepartmentDTO) (*depmodel.DepartmentDTO, error) {
	err := s.p.Pool.QueryRow(ctx,
		insertDepQuery,
		dep.Name, dep.ShortName,
	).Scan(&dep.ID)
	if err != nil {
		return nil, departmentError(err, amiderrors.NewCause("insert department query", "InsertDepartment", _PROVIDER))
	}
	return dep, nil
}
