package groupservice

import (
	"context"

	"github.com/amidgo/amiddocs/internal/errorutils/grouperror"
	"github.com/amidgo/amiddocs/internal/models/groupmodel"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
)

func (s *groupService) CreateGroup(
	ctx context.Context,
	group *groupmodel.GroupDTO,
) (*groupmodel.GroupDTO, error) {

	_, err := s.depRep.DepartmentById(ctx, group.DepartmentId)
	if err != nil {
		return nil, err
	}
	_, err = s.groupProv.GroupByName(ctx, group.Name)
	if !amiderrors.Is(err, grouperror.GROUP_NOT_FOUND) {
		return nil, grouperror.GROUP_NAME_ALREADY_EXIST
	}

	gr, err := s.groupRep.InsertGroup(ctx, group)
	if err != nil {
		return nil, err
	}
	return gr, nil
}
