package groupservice

import (
	"context"

	"github.com/amidgo/amiddocs/internal/models/depmodel"
	"github.com/amidgo/amiddocs/internal/models/depmodel/depfields"
	"github.com/amidgo/amiddocs/internal/models/groupmodel"
	"github.com/amidgo/amiddocs/internal/models/groupmodel/groupfields"
)

type GroupRepository interface {
	InsertGroup(ctx context.Context, group *groupmodel.GroupDTO) (*groupmodel.GroupDTO, error)
	InsertManyGroups(ctx context.Context, groups []*groupmodel.GroupDTO) error
}

type GroupProvider interface {
	GroupById(ctx context.Context, id uint64) (*groupmodel.GroupDTO, error)
	GroupByName(ctx context.Context, name groupfields.Name) (*groupmodel.GroupDTO, error)
}

type DepartmentProvider interface {
	StudyDepartmentById(ctx context.Context, studyDepId uint64) (*depmodel.StudyDepartmentDTO, error)
	StudyDepartmentByShortName(ctx context.Context, shortName depfields.ShortName) (*depmodel.StudyDepartmentDTO, error)
}

type GroupService struct {
	depProv   DepartmentProvider
	groupProv GroupProvider
	groupRep  GroupRepository
}

func New(groupRep GroupRepository, depRep DepartmentProvider, groupProv GroupProvider) *GroupService {
	return &GroupService{groupProv: groupProv, depProv: depRep, groupRep: groupRep}
}
