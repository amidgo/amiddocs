package departmentservice

import (
	"context"

	"github.com/amidgo/amiddocs/internal/errorutils/departmenterror"
	"github.com/amidgo/amiddocs/internal/models/depmodel"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
)

func (s *departmentService) CreateDepartment(
	ctx context.Context,
	dep *depmodel.DepartmentDTO,
) (*depmodel.DepartmentDTO, error) {

	_, err := s.depProvider.DepartmentByName(ctx, dep.Name)
	if !amiderrors.Is(err, departmenterror.DEPARMENT_NOT_FOUND) {
		return nil, departmenterror.NAME_EXIST
	}
	_, err = s.depProvider.DepartmentByShortName(ctx, dep.ShortName)
	if !amiderrors.Is(err, departmenterror.DEPARMENT_NOT_FOUND) {
		return nil, departmenterror.SHORT_NAME_EXIST
	}
	department, err := s.depRep.InsertDepartment(ctx, dep)
	if err != nil {
		return nil, err
	}
	return department, nil
}
