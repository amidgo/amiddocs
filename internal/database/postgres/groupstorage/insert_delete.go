package groupstorage

import (
	"context"

	"github.com/amidgo/amiddocs/internal/models/groupmodel"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
)

func (g *groupStorage) InsertGroup(ctx context.Context, group *groupmodel.GroupDTO) (*groupmodel.GroupDTO, error) {
	rows, err := g.p.DB.NamedQueryContext(ctx,
		`INSERT INTO groups (name,is_budget,education_form,education_start_date,education_year,education_finish_date,department_id) 
		 VALUES (:name,:is_budget,:education_form,:education_start_date,:education_year,:education_finish_date,:department_id)
		 RETURNING id`,
		group,
	)
	if err != nil {
		return nil, groupError(err, amiderrors.NewCause("insert group query", "InsertGroup", _PROVIDER))
	}
	for rows.Next() {
		err = rows.Scan(&group.ID)
	}
	if err != nil {
		return nil, groupError(err, amiderrors.NewCause("scan group id", "InsertGroup", _PROVIDER))
	}
	return group, nil
}

func (g *groupStorage) DeleteGroup(ctx context.Context, id uint64) error {
	_, err := g.p.DB.ExecContext(ctx, "DELETE FROM groups where id = $1", id)
	if err != nil {
		return groupError(err, amiderrors.NewCause("delete group exec", "DeleteGroup", _PROVIDER))
	}
	return nil
}
