package groupstorage

import (
	"context"
	"fmt"

	"github.com/amidgo/amiddocs/internal/models/groupmodel"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
)

var (
	insertGroupQuery = fmt.Sprintf(
		`INSERT INTO %s 
		(%s,%s,%s,%s,%s,%s,%s) 
		VALUES ($1,$2,$3,$4,$5,$6,$7)
		RETURNING %s`,
		groupmodel.GroupTable,

		groupmodel.SQL.Name,
		groupmodel.SQL.IsBudget,
		groupmodel.SQL.EducationForm,
		groupmodel.SQL.EducationStartDate,
		groupmodel.SQL.EducationYear,
		groupmodel.SQL.EducationFinishDate,
		groupmodel.SQL.StudyDepartmentId,

		groupmodel.SQL.ID,
	)

	deleteGroupQuery = fmt.Sprintf(
		`DELETE FROM %s WHERE %s = $1`,
		groupmodel.GroupTable,
		groupmodel.SQL.ID,
	)
)

func (g *groupStorage) InsertGroup(ctx context.Context, group *groupmodel.GroupDTO) (*groupmodel.GroupDTO, error) {
	err := g.p.Pool.QueryRow(ctx,
		insertGroupQuery,
		group.Name,
		group.IsBudget,
		group.EducationForm,
		group.EducationStartDate,
		group.EducationYear,
		group.EducationFinishDate,
		group.StudyDepartmentId,
	).Scan(&group.ID)
	if err != nil {
		return nil, groupError(err, amiderrors.NewCause("insert group query", "InsertGroup", _PROVIDER))
	}
	return group, nil
}

func (g *groupStorage) DeleteGroup(ctx context.Context, id uint64) error {
	_, err := g.p.Pool.Exec(ctx, deleteGroupQuery, id)
	if err != nil {
		return groupError(err, amiderrors.NewCause("delete group exec", "DeleteGroup", _PROVIDER))
	}
	return nil
}
