// Code generated by MockGen. DO NOT EDIT.
// Source: internal/domain/studentservice/studentservice.go

// Package studentservice is a generated GoMock package.
package studentservice

import (
	context "context"
	reflect "reflect"

	depmodel "github.com/amidgo/amiddocs/internal/models/depmodel"
	groupmodel "github.com/amidgo/amiddocs/internal/models/groupmodel"
	groupfields "github.com/amidgo/amiddocs/internal/models/groupmodel/groupfields"
	stdocmodel "github.com/amidgo/amiddocs/internal/models/stdocmodel"
	stdocfields "github.com/amidgo/amiddocs/internal/models/stdocmodel/stdocfields"
	studentmodel "github.com/amidgo/amiddocs/internal/models/studentmodel"
	usermodel "github.com/amidgo/amiddocs/internal/models/usermodel"
	userfields "github.com/amidgo/amiddocs/internal/models/usermodel/userfields"
	gomock "github.com/golang/mock/gomock"
)

// MockGroupProvider is a mock of GroupProvider interface.
type MockGroupProvider struct {
	ctrl     *gomock.Controller
	recorder *MockGroupProviderMockRecorder
}

// MockGroupProviderMockRecorder is the mock recorder for MockGroupProvider.
type MockGroupProviderMockRecorder struct {
	mock *MockGroupProvider
}

// NewMockGroupProvider creates a new mock instance.
func NewMockGroupProvider(ctrl *gomock.Controller) *MockGroupProvider {
	mock := &MockGroupProvider{ctrl: ctrl}
	mock.recorder = &MockGroupProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGroupProvider) EXPECT() *MockGroupProviderMockRecorder {
	return m.recorder
}

// GroupByName mocks base method.
func (m *MockGroupProvider) GroupByName(ctx context.Context, name groupfields.Name) (*groupmodel.GroupDTO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GroupByName", ctx, name)
	ret0, _ := ret[0].(*groupmodel.GroupDTO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GroupByName indicates an expected call of GroupByName.
func (mr *MockGroupProviderMockRecorder) GroupByName(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GroupByName", reflect.TypeOf((*MockGroupProvider)(nil).GroupByName), ctx, name)
}

// MockStudentDocProvider is a mock of StudentDocProvider interface.
type MockStudentDocProvider struct {
	ctrl     *gomock.Controller
	recorder *MockStudentDocProviderMockRecorder
}

// MockStudentDocProviderMockRecorder is the mock recorder for MockStudentDocProvider.
type MockStudentDocProviderMockRecorder struct {
	mock *MockStudentDocProvider
}

// NewMockStudentDocProvider creates a new mock instance.
func NewMockStudentDocProvider(ctrl *gomock.Controller) *MockStudentDocProvider {
	mock := &MockStudentDocProvider{ctrl: ctrl}
	mock.recorder = &MockStudentDocProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStudentDocProvider) EXPECT() *MockStudentDocProviderMockRecorder {
	return m.recorder
}

// DocumentByDocNumber mocks base method.
func (m *MockStudentDocProvider) DocumentByDocNumber(ctx context.Context, docNumber stdocfields.DocNumber) (*stdocmodel.StudentDocumentDTO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DocumentByDocNumber", ctx, docNumber)
	ret0, _ := ret[0].(*stdocmodel.StudentDocumentDTO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DocumentByDocNumber indicates an expected call of DocumentByDocNumber.
func (mr *MockStudentDocProviderMockRecorder) DocumentByDocNumber(ctx, docNumber interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DocumentByDocNumber", reflect.TypeOf((*MockStudentDocProvider)(nil).DocumentByDocNumber), ctx, docNumber)
}

// DocumentByOrderNumber mocks base method.
func (m *MockStudentDocProvider) DocumentByOrderNumber(ctx context.Context, orderNumber stdocfields.OrderNumber) (*stdocmodel.StudentDocumentDTO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DocumentByOrderNumber", ctx, orderNumber)
	ret0, _ := ret[0].(*stdocmodel.StudentDocumentDTO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DocumentByOrderNumber indicates an expected call of DocumentByOrderNumber.
func (mr *MockStudentDocProviderMockRecorder) DocumentByOrderNumber(ctx, orderNumber interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DocumentByOrderNumber", reflect.TypeOf((*MockStudentDocProvider)(nil).DocumentByOrderNumber), ctx, orderNumber)
}

// MockUserProvider is a mock of UserProvider interface.
type MockUserProvider struct {
	ctrl     *gomock.Controller
	recorder *MockUserProviderMockRecorder
}

// MockUserProviderMockRecorder is the mock recorder for MockUserProvider.
type MockUserProviderMockRecorder struct {
	mock *MockUserProvider
}

// NewMockUserProvider creates a new mock instance.
func NewMockUserProvider(ctrl *gomock.Controller) *MockUserProvider {
	mock := &MockUserProvider{ctrl: ctrl}
	mock.recorder = &MockUserProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserProvider) EXPECT() *MockUserProviderMockRecorder {
	return m.recorder
}

// UserByEmail mocks base method.
func (m *MockUserProvider) UserByEmail(ctx context.Context, email userfields.Email) (*usermodel.UserDTO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserByEmail", ctx, email)
	ret0, _ := ret[0].(*usermodel.UserDTO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserByEmail indicates an expected call of UserByEmail.
func (mr *MockUserProviderMockRecorder) UserByEmail(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserByEmail", reflect.TypeOf((*MockUserProvider)(nil).UserByEmail), ctx, email)
}

// MockStudentProvider is a mock of StudentProvider interface.
type MockStudentProvider struct {
	ctrl     *gomock.Controller
	recorder *MockStudentProviderMockRecorder
}

// MockStudentProviderMockRecorder is the mock recorder for MockStudentProvider.
type MockStudentProviderMockRecorder struct {
	mock *MockStudentProvider
}

// NewMockStudentProvider creates a new mock instance.
func NewMockStudentProvider(ctrl *gomock.Controller) *MockStudentProvider {
	mock := &MockStudentProvider{ctrl: ctrl}
	mock.recorder = &MockStudentProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStudentProvider) EXPECT() *MockStudentProviderMockRecorder {
	return m.recorder
}

// AllStudents mocks base method.
func (m *MockStudentProvider) AllStudents(ctx context.Context) ([]*studentmodel.StudentDTO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllStudents", ctx)
	ret0, _ := ret[0].([]*studentmodel.StudentDTO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllStudents indicates an expected call of AllStudents.
func (mr *MockStudentProviderMockRecorder) AllStudents(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllStudents", reflect.TypeOf((*MockStudentProvider)(nil).AllStudents), ctx)
}

// StudentById mocks base method.
func (m *MockStudentProvider) StudentById(ctx context.Context, id uint64) (*studentmodel.StudentDTO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StudentById", ctx, id)
	ret0, _ := ret[0].(*studentmodel.StudentDTO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StudentById indicates an expected call of StudentById.
func (mr *MockStudentProviderMockRecorder) StudentById(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StudentById", reflect.TypeOf((*MockStudentProvider)(nil).StudentById), ctx, id)
}

// MockStudentRepository is a mock of StudentRepository interface.
type MockStudentRepository struct {
	ctrl     *gomock.Controller
	recorder *MockStudentRepositoryMockRecorder
}

// MockStudentRepositoryMockRecorder is the mock recorder for MockStudentRepository.
type MockStudentRepositoryMockRecorder struct {
	mock *MockStudentRepository
}

// NewMockStudentRepository creates a new mock instance.
func NewMockStudentRepository(ctrl *gomock.Controller) *MockStudentRepository {
	mock := &MockStudentRepository{ctrl: ctrl}
	mock.recorder = &MockStudentRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStudentRepository) EXPECT() *MockStudentRepositoryMockRecorder {
	return m.recorder
}

// InsertManyStudents mocks base method.
func (m *MockStudentRepository) InsertManyStudents(ctx context.Context, students []*studentmodel.StudentDTO) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertManyStudents", ctx, students)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertManyStudents indicates an expected call of InsertManyStudents.
func (mr *MockStudentRepositoryMockRecorder) InsertManyStudents(ctx, students interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertManyStudents", reflect.TypeOf((*MockStudentRepository)(nil).InsertManyStudents), ctx, students)
}

// InsertStudent mocks base method.
func (m *MockStudentRepository) InsertStudent(ctx context.Context, student *studentmodel.StudentDTO) (*studentmodel.StudentDTO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertStudent", ctx, student)
	ret0, _ := ret[0].(*studentmodel.StudentDTO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertStudent indicates an expected call of InsertStudent.
func (mr *MockStudentRepositoryMockRecorder) InsertStudent(ctx, student interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertStudent", reflect.TypeOf((*MockStudentRepository)(nil).InsertStudent), ctx, student)
}

// MockDepProvider is a mock of DepProvider interface.
type MockDepProvider struct {
	ctrl     *gomock.Controller
	recorder *MockDepProviderMockRecorder
}

// MockDepProviderMockRecorder is the mock recorder for MockDepProvider.
type MockDepProviderMockRecorder struct {
	mock *MockDepProvider
}

// NewMockDepProvider creates a new mock instance.
func NewMockDepProvider(ctrl *gomock.Controller) *MockDepProvider {
	mock := &MockDepProvider{ctrl: ctrl}
	mock.recorder = &MockDepProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDepProvider) EXPECT() *MockDepProviderMockRecorder {
	return m.recorder
}

// StudyDepartmentById mocks base method.
func (m *MockDepProvider) StudyDepartmentById(ctx context.Context, studyDepId uint64) (*depmodel.StudyDepartmentDTO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StudyDepartmentById", ctx, studyDepId)
	ret0, _ := ret[0].(*depmodel.StudyDepartmentDTO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StudyDepartmentById indicates an expected call of StudyDepartmentById.
func (mr *MockDepProviderMockRecorder) StudyDepartmentById(ctx, studyDepId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StudyDepartmentById", reflect.TypeOf((*MockDepProvider)(nil).StudyDepartmentById), ctx, studyDepId)
}

// MockEncrypter is a mock of Encrypter interface.
type MockEncrypter struct {
	ctrl     *gomock.Controller
	recorder *MockEncrypterMockRecorder
}

// MockEncrypterMockRecorder is the mock recorder for MockEncrypter.
type MockEncrypterMockRecorder struct {
	mock *MockEncrypter
}

// NewMockEncrypter creates a new mock instance.
func NewMockEncrypter(ctrl *gomock.Controller) *MockEncrypter {
	mock := &MockEncrypter{ctrl: ctrl}
	mock.recorder = &MockEncrypterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEncrypter) EXPECT() *MockEncrypterMockRecorder {
	return m.recorder
}

// Hash mocks base method.
func (m *MockEncrypter) Hash(input string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Hash", input)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Hash indicates an expected call of Hash.
func (mr *MockEncrypterMockRecorder) Hash(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Hash", reflect.TypeOf((*MockEncrypter)(nil).Hash), input)
}

// Verify mocks base method.
func (m *MockEncrypter) Verify(hashPassword, password string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Verify", hashPassword, password)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Verify indicates an expected call of Verify.
func (mr *MockEncrypterMockRecorder) Verify(hashPassword, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Verify", reflect.TypeOf((*MockEncrypter)(nil).Verify), hashPassword, password)
}
