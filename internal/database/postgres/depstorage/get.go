package depstorage

import (
	"context"
	"fmt"

	"github.com/amidgo/amiddocs/internal/models/depmodel"
	"github.com/amidgo/amiddocs/internal/models/depmodel/depfields"
	"github.com/amidgo/amiddocs/internal/models/groupmodel"
	"github.com/amidgo/amiddocs/internal/models/studentmodel"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
	"github.com/amidgo/amiddocs/pkg/sqlutils"
	"github.com/jackc/pgx/v5"
)

var (
	getDepQuery = fmt.Sprintf(
		`SELECT %s, %s, %s, %s FROM %s`,
		depmodel.SQL.ID,
		depmodel.SQL.Name,
		depmodel.SQL.ShortName,
		depmodel.SQL.ImageUrl,

		depmodel.DepartmentTable,
	)

	getDepByUserId = fmt.Sprintf(
		`
		SELECT %s, %s, %s, %s
			FROM %s
		INNER JOIN %s ON %s = %s
		INNER JOIN %s ON %s = %s
			WHERE %s = $1
		GROUP BY %s, %s, %s
		`,
		sqlutils.Full(depmodel.SQL.ID),
		sqlutils.Full(depmodel.SQL.Name),
		sqlutils.Full(depmodel.SQL.ShortName),
		sqlutils.Full(depmodel.SQL.ImageUrl),

		depmodel.DepartmentTable,

		// inner join on groups
		groupmodel.GroupTable,
		sqlutils.Full(groupmodel.SQL.StudyDepartmentId),
		sqlutils.Full(depmodel.SQL.ID),

		// inner join on students
		studentmodel.StudentTable,
		sqlutils.Full(studentmodel.SQL.GroupId),
		sqlutils.Full(groupmodel.SQL.ID),

		// where students.user_id = $1
		sqlutils.Full(studentmodel.SQL.UserID),

		// group by columns
		sqlutils.Full(groupmodel.SQL.ID),
		sqlutils.Full(depmodel.SQL.ID),
		sqlutils.Full(studentmodel.SQL.ID),
	)
)

func getByColumnQuery(column sqlutils.Column) string {
	return getDepQuery + fmt.Sprintf(` WHERE %s = $1`, column)
}

func scanDep(row pgx.Row, dep *depmodel.DepartmentDTO) error {
	return row.Scan(
		&dep.ID,
		&dep.Name,
		&dep.ShortName,
		&dep.ImageUrl,
	)
}

func (s *departmentStorage) getDepartmentByQuery(ctx context.Context, query string, args ...interface{}) (*depmodel.DepartmentDTO, error) {
	dep := new(depmodel.DepartmentDTO)
	row := s.p.Pool.QueryRow(
		ctx,
		query,
		args...,
	)
	err := scanDep(row, dep)
	if err != nil {
		return nil, err
	}
	return dep, nil
}

func (s *departmentStorage) DepartmentById(ctx context.Context, depId uint64) (*depmodel.DepartmentDTO, error) {
	dep, err := s.getDepartmentByQuery(ctx, getByColumnQuery(depmodel.SQL.ID), depId)
	if err != nil {
		return nil, departmentError(err, amiderrors.NewCause("department by id query", "DepartmentById", _PROVIDER))
	}
	return dep, nil
}

func (s *departmentStorage) DepartmentByName(ctx context.Context, name depfields.Name) (*depmodel.DepartmentDTO, error) {
	dep, err := s.getDepartmentByQuery(ctx, getByColumnQuery(depmodel.SQL.Name), name)
	if err != nil {
		return nil, departmentError(err, amiderrors.NewCause("deparment by name query", "DepartmentByName", _PROVIDER))
	}
	return dep, nil
}

func (s *departmentStorage) DepartmentByShortName(ctx context.Context, shortName depfields.ShortName) (*depmodel.DepartmentDTO, error) {
	dep, err := s.getDepartmentByQuery(ctx, getByColumnQuery(depmodel.SQL.ShortName), shortName)
	if err != nil {
		return nil, departmentError(err, amiderrors.NewCause("department by short name query", "DeparmentByShortName", _PROVIDER))
	}
	return dep, nil
}

func (s *departmentStorage) DepartmentList(ctx context.Context) ([]*depmodel.DepartmentDTO, error) {
	departmentModelList := make([]*depmodel.DepartmentDTO, 0)
	rows, err := s.p.Pool.Query(ctx, getDepQuery)
	if err != nil {
		return nil, departmentError(err, amiderrors.NewCause("all departments query", "AllDepartments", _PROVIDER))
	}
	for rows.Next() {
		dep := new(depmodel.DepartmentDTO)
		err := scanDep(rows, dep)
		if err != nil {
			return nil, departmentError(err, amiderrors.NewCause("scan department from row", "AllDepartments", _PROVIDER))
		}
		departmentModelList = append(departmentModelList, dep)
	}
	return departmentModelList, nil
}

func (s *departmentStorage) DepartmentByUserId(ctx context.Context, userId uint64) (*depmodel.DepartmentDTO, error) {
	dep := new(depmodel.DepartmentDTO)
	row := s.p.Pool.QueryRow(ctx, getDepByUserId, userId)
	err := scanDep(row, dep)
	if err != nil {
		return nil, departmentError(err, amiderrors.NewCause("get dep by user id query", "DepartmentByUserId", _PROVIDER))
	}
	return dep, nil
}
