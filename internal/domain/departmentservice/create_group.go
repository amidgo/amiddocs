package departmentservice

import (
	"context"

	"github.com/amidgo/amiddocs/internal/models/depmodel"
)

func (s *departmentService) CreateDepartment(
	ctx context.Context,
	dep *depmodel.CreateDepartmentDTO,
) (*depmodel.DepartmentDTO, error) {
	department := depmodel.NewDepartmentDTO(0, dep.Name, dep.ShortName)
	department, err := s.depRep.InsertDepartment(ctx, department)
	if err != nil {
		return nil, err
	}
	return department, nil
}
