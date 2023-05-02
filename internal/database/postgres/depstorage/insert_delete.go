package depstorage

import (
	"context"

	"github.com/amidgo/amiddocs/internal/models/depmodel"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
)

func (s *departmentStorage) InsertDepartment(ctx context.Context, dep *depmodel.DepartmentDTO) (*depmodel.DepartmentDTO, error) {
	row, err := s.p.DB.NamedQueryContext(ctx,
		`INSERT INTO departments (name, short_name) VALUES (:name, :short_name) RETURNING id`,
		dep,
	)
	if err != nil {
		return nil, departmentError(err, amiderrors.NewCause("insert department query", "InsertDepartment", _PROVIDER))
	}
	for row.Next() {
		err = row.Scan(&dep.ID)
		if err != nil {
			return nil, departmentError(err, amiderrors.NewCause("scan department id", "InsertDepartment", _PROVIDER))
		}
	}
	return dep, nil
}
