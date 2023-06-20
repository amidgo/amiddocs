package groupservice

import (
	"context"

	"github.com/amidgo/amiddocs/internal/models/groupmodel"
)

func (s *GroupService) CreateGroup(
	ctx context.Context,
	group *groupmodel.GroupDTO,
) (*groupmodel.GroupDTO, error) {
	group, err := s.groupRep.InsertGroup(ctx, group)
	if err != nil {
		return nil, err
	}
	return group, nil
}
