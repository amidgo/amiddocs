package groupservice_test

import (
	"context"
	"testing"

	"github.com/amidgo/amiddocs/internal/domain/groupservice"
	"github.com/amidgo/amiddocs/internal/errorutils/departmenterror"
	"github.com/amidgo/amiddocs/internal/errorutils/grouperror"
	"github.com/amidgo/amiddocs/internal/models/depmodel"
	"github.com/amidgo/amiddocs/internal/models/groupmodel"
	"github.com/amidgo/amiddocs/internal/models/groupmodel/groupfields"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
	"github.com/amidgo/amiddocs/pkg/amidtime"
	"github.com/amidgo/amiddocs/pkg/assert"
	"github.com/golang/mock/gomock"
)

var (
	groupDTO = groupmodel.NewGroupDTO(
		1,
		"isp-335",
		false,
		groupfields.EXTRAMURAL,
		amidtime.NewDate(2020, 1, 1),
		amidtime.NewDate(2020, 1, 1),
		1,
		1,
	)
	departmentDTO = depmodel.NewStudyDepartmentDTO(1,
		*depmodel.NewDepartmentDTO(
			1,
			"it",
			"it",
		),
	)
)

func TestCreateGroup(t *testing.T) {
	ctrl := gomock.NewController(t)

	var act, exp error

	groupRepository := groupservice.NewMockGroupRepository(ctrl)
	groupProvider := groupservice.NewMockGroupProvider(ctrl)
	departmentProvider := groupservice.NewMockDepartmentProvider(ctrl)

	groupService := groupservice.New(groupRepository, departmentProvider, groupProvider)
	ctx := context.Background()
	group := groupDTO
	department := departmentDTO

	// case when we dont find study department
	departmentProvider.EXPECT().
		StudyDepartmentById(gomock.Any(), groupDTO.StudyDepartmentId).
		Return(nil, departmenterror.DEPARTMENT_NOT_FOUND).
		Times(1)
	exp = departmenterror.DEPARTMENT_NOT_FOUND
	_, act = groupService.CreateGroup(ctx, group)
	assert.ErrorEqual(t, exp, act, "error from department repo test failed")

	departmentProvider.EXPECT().
		StudyDepartmentById(gomock.Any(), groupDTO.StudyDepartmentId).
		Return(department, nil).
		Times(1)
	groupProvider.EXPECT().
		GroupByName(gomock.Any(), group.Name).
		Return(group, nil).
		Times(1)
	exp = grouperror.GROUP_NAME_ALREADY_EXIST
	_, act = groupService.CreateGroup(ctx, group)
	assert.ErrorEqual(t, exp, act, "error from group repo test failed")

	// case when groupProvider return non-nil error not equal grouperror.GROUP_NOT_FOUND
	departmentProvider.EXPECT().
		StudyDepartmentById(gomock.Any(), groupDTO.StudyDepartmentId).
		Return(department, nil).
		Times(1)
	groupProvider.EXPECT().
		GroupByName(gomock.Any(), group.Name).
		Return(nil, amiderrors.Internal()).
		Times(1)
	exp = amiderrors.Internal()
	_, act = groupService.CreateGroup(ctx, group)
	assert.ErrorEqual(t, exp, act, "error in insert group test failed")

	// case insert group method return error
	// group studyDepartment exist
	// group with the same name doesnt exist in db
	departmentProvider.EXPECT().
		StudyDepartmentById(gomock.Any(), groupDTO.StudyDepartmentId).
		Return(department, nil).
		Times(1)
	groupProvider.EXPECT().
		GroupByName(gomock.Any(), group.Name).
		Return(nil, grouperror.GROUP_NOT_FOUND).
		Times(1)
	groupRepository.EXPECT().
		InsertGroup(gomock.Any(), gomock.Any()).
		Return(nil, amiderrors.Internal())
	exp = amiderrors.Internal()
	_, act = groupService.CreateGroup(ctx, group)
	assert.ErrorEqual(t, exp, act, "error in insert group test failed")

	// case insert group method no return error
	// group studyDepartment exist
	// group with the same name doesnt exist in db
	departmentProvider.EXPECT().
		StudyDepartmentById(gomock.Any(), groupDTO.StudyDepartmentId).
		Return(department, nil).
		Times(1)
	groupProvider.EXPECT().
		GroupByName(gomock.Any(), group.Name).
		Return(nil, grouperror.GROUP_NOT_FOUND).
		Times(1)
	groupRepository.EXPECT().
		InsertGroup(gomock.Any(), gomock.Any()).
		Return(nil, nil)
	exp = nil
	_, act = groupService.CreateGroup(ctx, group)
	assert.ErrorEqual(t, exp, act, "error in insert group test failed")
}
