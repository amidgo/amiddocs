package depstorage

import (
	"context"
	"fmt"

	"github.com/amidgo/amiddocs/internal/models/depmodel"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
	"github.com/amidgo/amiddocs/pkg/sqlutils"
)

var (
	studyDepQuery = fmt.Sprintf(
		`
		SELECT %s, %s, %s
			FROM %s
		INNER JOIN %s ON %s = %s

		WHERE %s = $1
		
		GROUP BY %s
		`,
		sqlutils.Full(depmodel.SQL.ID),
		sqlutils.Full(depmodel.SQL.Name),
		sqlutils.Full(depmodel.SQL.ShortName),

		depmodel.DepartmentTable,

		// inner join study departments
		depmodel.StudyDepartmentTable,
		sqlutils.Full(depmodel.SQL_STUDY_DEP.DepartmentId),
		sqlutils.Full(depmodel.SQL.ID),

		// where study dep id = $1
		sqlutils.Full(depmodel.SQL_STUDY_DEP.DepartmentId),

		// group by dep id
		sqlutils.Full(depmodel.SQL.ID),
	)
)

func (s departmentStorage) StudyDepartment(ctx context.Context, studyDepId uint64) (*depmodel.DepartmentDTO, error) {
	row := s.p.Pool.QueryRow(ctx, studyDepQuery, studyDepId)
	dep := new(depmodel.DepartmentDTO)
	err := scanDep(row, dep)
	if err != nil {
		return nil, departmentError(err, amiderrors.NewCause("get dep by study dep id query", "StudyDepartment", _PROVIDER))
	}
	return dep, nil
}
