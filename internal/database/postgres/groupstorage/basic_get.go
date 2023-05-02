package groupstorage

import (
	"context"

	"github.com/amidgo/amiddocs/internal/models/groupmodel"
	"github.com/amidgo/amiddocs/internal/models/groupmodel/groupfields"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
)

const (
	_BASIC_GROUP_GET = `SELECT id,name,is_budget,education_form,education_start_date,education_year,education_finish_date,department_id FROM groups `
)

func (g *groupStorage) getGroupByQuery(ctx context.Context, query string, args ...interface{}) (*groupmodel.GroupDTO, error) {
	group := new(groupmodel.GroupDTO)
	err := g.p.DB.GetContext(ctx, group, query, args...)
	if err != nil {
		return nil, err
	}
	return group, nil
}

func (g *groupStorage) GroupById(ctx context.Context, id uint64) (*groupmodel.GroupDTO, error) {
	group, err := g.getGroupByQuery(ctx, _BASIC_GROUP_GET+`WHERE id = $1`, id)
	if err != nil {
		return nil, groupError(err, amiderrors.NewCause("group by id query", "GroupById", _PROVIDER))
	}
	return group, nil
}

func (g *groupStorage) GroupByName(ctx context.Context, name groupfields.Name) (*groupmodel.GroupDTO, error) {
	group, err := g.getGroupByQuery(ctx, _BASIC_GROUP_GET+`WHERE name = $1`, name)
	if err != nil {
		return nil, groupError(err, amiderrors.NewCause("group by name query", "GroupByName", _PROVIDER))
	}
	return group, nil
}
