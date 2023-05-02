package departmentservice

import (
	"context"

	"github.com/amidgo/amiddocs/internal/models/depmodel"
	"github.com/amidgo/amiddocs/internal/models/depmodel/depfields"
)

type departmentRepository interface {
	InsertDepartment(ctx context.Context, dep *depmodel.DepartmentDTO) (*depmodel.DepartmentDTO, error)
}

type departmentProvider interface {
	DepartmentById(ctx context.Context, id uint64) (*depmodel.DepartmentDTO, error)
	AllDepartments(ctx context.Context) ([]*depmodel.DepartmentDTO, error)
	DepartmentByName(ctx context.Context, name depfields.Name) (*depmodel.DepartmentDTO, error)
	DepartmentByShortName(ctx context.Context, shortName depfields.ShortName) (*depmodel.DepartmentDTO, error)
}

type departmentService struct {
	depRep      departmentRepository
	depProvider departmentProvider
}

func New(depRep departmentRepository, depProvider departmentProvider) *departmentService {
	return &departmentService{depRep: depRep, depProvider: depProvider}
}
