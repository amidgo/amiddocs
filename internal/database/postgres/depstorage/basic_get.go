package depstorage

import (
	"context"

	"github.com/amidgo/amiddocs/internal/models/depmodel"
	"github.com/amidgo/amiddocs/internal/models/depmodel/depfields"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
)

const (
	_DEP_BASIC_GET = `SELECT id, name, short_name FROM departments `
)

func (s *departmentStorage) getDepartmentByQuery(ctx context.Context, query string, args ...interface{}) (*depmodel.DepartmentDTO, error) {
	dep := new(depmodel.DepartmentDTO)
	err := s.p.DB.GetContext(ctx, dep,
		query,
		args...,
	)
	if err != nil {
		return nil, err
	}
	return dep, nil
}

func (s *departmentStorage) DepartmentById(ctx context.Context, depId uint64) (*depmodel.DepartmentDTO, error) {
	dep, err := s.getDepartmentByQuery(ctx, _DEP_BASIC_GET+`WHERE id = $1`, depId)
	if err != nil {
		return nil, departmentError(err, amiderrors.NewCause("department by id query", "DepartmentById", _PROVIDER))
	}
	return dep, nil
}

func (s *departmentStorage) DepartmentByName(ctx context.Context, name depfields.Name) (*depmodel.DepartmentDTO, error) {
	dep, err := s.getDepartmentByQuery(ctx, _DEP_BASIC_GET+`WHERE name = $1`, name)
	if err != nil {
		return nil, departmentError(err, amiderrors.NewCause("deparment by name query", "DepartmentByName", _PROVIDER))
	}
	return dep, nil
}

func (s *departmentStorage) DepartmentByShortName(ctx context.Context, shortName depfields.ShortName) (*depmodel.DepartmentDTO, error) {
	dep, err := s.getDepartmentByQuery(ctx, _DEP_BASIC_GET+`WHERE short_name = $1`, shortName)
	if err != nil {
		return nil, departmentError(err, amiderrors.NewCause("department by short name query", "DeparmentByShortName", _PROVIDER))
	}
	return dep, nil
}

func (s *departmentStorage) AllDepartments(ctx context.Context) ([]*depmodel.DepartmentDTO, error) {
	departmentModelList := make([]*depmodel.DepartmentDTO, 0)
	err := s.p.DB.SelectContext(ctx, &departmentModelList, _DEP_BASIC_GET)
	if err != nil {
		return nil, departmentError(err, amiderrors.NewCause("all departments query", "AllDepartments", _PROVIDER))
	}
	return departmentModelList, nil
}
