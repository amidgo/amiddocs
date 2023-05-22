package groupstorage

import (
	"context"
	"fmt"

	"github.com/amidgo/amiddocs/internal/models/groupmodel"
	"github.com/amidgo/amiddocs/internal/models/groupmodel/groupfields"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
	"github.com/jackc/pgx/v5"
)

var (
	groupGetQuery = fmt.Sprintf(
		`SELECT 
		%s,%s,%s,%s,%s,%s,%s,%s
		FROM %s
		`,
		groupmodel.SQL.ID,
		groupmodel.SQL.Name,
		groupmodel.SQL.IsBudget,
		groupmodel.SQL.EducationForm,
		groupmodel.SQL.EducationStartDate,
		groupmodel.SQL.EducationYear,
		groupmodel.SQL.EducationFinishDate,
		groupmodel.SQL.StudyDepartmentId,

		groupmodel.GroupTable,
	)
)

func scanGroup(ctx context.Context, row pgx.Row, group *groupmodel.GroupDTO) error {
	return row.Scan(
		&group.ID,
		&group.Name,
		&group.IsBudget,
		&group.EducationForm,
		&group.EducationStartDate,
		&group.EducationYear,
		&group.EducationFinishDate,
		&group.StudyDepartmentId,
	)
}

func (g *groupStorage) getGroupByQuery(ctx context.Context, query string, args ...interface{}) (*groupmodel.GroupDTO, error) {
	group := new(groupmodel.GroupDTO)
	row := g.p.Pool.QueryRow(ctx, query, args...)
	err := scanGroup(ctx, row, group)
	if err != nil {
		return nil, err
	}
	return group, nil
}

func (g *groupStorage) GroupById(ctx context.Context, id uint64) (*groupmodel.GroupDTO, error) {
	group, err := g.getGroupByQuery(ctx, groupGetQuery+` WHERE `+groupmodel.SQL.ID.String()+` = $1`, id)
	if err != nil {
		return nil, groupError(err, amiderrors.NewCause("group by id query", "GroupById", _PROVIDER))
	}
	return group, nil
}

func (g *groupStorage) GroupByName(ctx context.Context, name groupfields.Name) (*groupmodel.GroupDTO, error) {
	group, err := g.getGroupByQuery(ctx, groupGetQuery+` WHERE `+groupmodel.SQL.Name.String()+` = $1`, name)
	if err != nil {
		return nil, groupError(err, amiderrors.NewCause("group by name query", "GroupByName", _PROVIDER))
	}
	return group, nil
}
