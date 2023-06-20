package depstorage

import (
	"context"
	"fmt"

	"github.com/amidgo/amiddocs/internal/models/depmodel"
	"github.com/amidgo/amiddocs/internal/models/depmodel/depfields"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
	"github.com/amidgo/amiddocs/pkg/sqlutils"
	"github.com/jackc/pgx/v5"
)

func studyDepQuery(query string) string {
	return fmt.Sprintf(
		`
		SELECT %s,%s, %s, %s, %s
			FROM %s
		INNER JOIN %s ON %s = %s

		%s
		
		GROUP BY %s, %s
		`,
		sqlutils.Full(depmodel.SQL_STUDY_DEP.ID),
		sqlutils.Full(depmodel.SQL.ID),
		sqlutils.Full(depmodel.SQL.Name),
		sqlutils.Full(depmodel.SQL.ShortName),
		sqlutils.Full(depmodel.SQL.ImageUrl),

		depmodel.DepartmentTable,

		// inner join study departments
		depmodel.StudyDepartmentTable,
		sqlutils.Full(depmodel.SQL_STUDY_DEP.DepartmentId),
		sqlutils.Full(depmodel.SQL.ID),

		query,

		// group by dep id
		sqlutils.Full(depmodel.SQL.ID),
		sqlutils.Full(depmodel.SQL_STUDY_DEP.ID),
	)
}

func scanStudyDep(row pgx.Row, dep *depmodel.StudyDepartmentDTO) error {
	return row.Scan(
		&dep.StudyDepartmentID,
		&dep.ID,
		&dep.Name,
		&dep.ShortName,
		&dep.ImageUrl,
	)
}

func (s *departmentStorage) StudyDepartmentById(ctx context.Context, studyDepId uint64) (*depmodel.StudyDepartmentDTO, error) {
	row := s.p.Pool.QueryRow(ctx,
		studyDepQuery(
			fmt.Sprintf("WHERE %s = $1", sqlutils.Full(depmodel.SQL_STUDY_DEP.ID)),
		),
		studyDepId,
	)
	dep := new(depmodel.StudyDepartmentDTO)
	err := scanStudyDep(row, dep)
	if err != nil {
		return nil, studyDepartmentError(err, amiderrors.NewCause("get dep by study dep id query", "StudyDepartmentById", _PROVIDER))
	}
	return dep, nil
}

func (s *departmentStorage) StudyDepartmentByShortName(
	ctx context.Context,
	shortName depfields.ShortName,
) (*depmodel.StudyDepartmentDTO, error) {
	row := s.p.Pool.QueryRow(ctx, studyDepQuery(
		fmt.Sprintf("WHERE %s = $1", sqlutils.Full(depmodel.SQL.ShortName)),
	))
	dep := new(depmodel.StudyDepartmentDTO)
	err := scanStudyDep(row, dep)
	if err != nil {
		return nil, studyDepartmentError(err, amiderrors.NewCause("get dep by short name query", "StudyDepartmentByShortName", _PROVIDER))
	}
	return dep, nil
}
