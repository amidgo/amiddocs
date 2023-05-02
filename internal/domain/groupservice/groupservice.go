package groupservice

import (
	"context"

	"github.com/amidgo/amiddocs/internal/models/depmodel"
	"github.com/amidgo/amiddocs/internal/models/groupmodel"
	"github.com/amidgo/amiddocs/internal/models/groupmodel/groupfields"
)

type groupRepository interface {
	InsertGroup(ctx context.Context, group *groupmodel.GroupDTO) (*groupmodel.GroupDTO, error)
}

type groupProvider interface {
	GroupById(ctx context.Context, id uint64) (*groupmodel.GroupDTO, error)
	GroupByName(ctx context.Context, name groupfields.Name) (*groupmodel.GroupDTO, error)
}

type departmentRepository interface {
	DepartmentById(ctx context.Context, id uint64) (*depmodel.DepartmentDTO, error)
}

type groupService struct {
	depRep    departmentRepository
	groupProv groupProvider
	groupRep  groupRepository
}

func New(groupRep groupRepository, depRep departmentRepository, groupProv groupProvider) *groupService {
	return &groupService{groupProv: groupProv, depRep: depRep, groupRep: groupRep}
}
