package departmentservice

import (
	"context"

	"github.com/amidgo/amiddocs/internal/filestorage"
	"github.com/amidgo/amiddocs/internal/models/depmodel"
	"github.com/amidgo/amiddocs/internal/models/depmodel/depfields"
)

type departmentRepository interface {
	InsertDepartment(ctx context.Context, dep *depmodel.DepartmentDTO) (*depmodel.DepartmentDTO, error)
	UpdateDepartmentPhoto(ctx context.Context, id uint64, imageUrl depfields.ImageUrl) error
}

type departmentProvider interface {
	DepartmentById(ctx context.Context, id uint64) (*depmodel.DepartmentDTO, error)
	DepartmentList(ctx context.Context) ([]*depmodel.DepartmentDTO, error)
	DepartmentByName(ctx context.Context, name depfields.Name) (*depmodel.DepartmentDTO, error)
	DepartmentByShortName(ctx context.Context, shortName depfields.ShortName) (*depmodel.DepartmentDTO, error)
}

type departmentService struct {
	depRep      departmentRepository
	depProvider departmentProvider
	depFS       filestorage.FileStorage
}

func New(depRep departmentRepository, depProvider departmentProvider, depFS filestorage.FileStorage) *departmentService {
	return &departmentService{depRep: depRep, depProvider: depProvider, depFS: depFS}
}
