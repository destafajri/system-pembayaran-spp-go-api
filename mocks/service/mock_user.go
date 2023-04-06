// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/destafajri/system-pembayaran-spp-go-api/internal/service (interfaces: UserService)

// Package mock_service is a generated GoMock package.
package mock_service

import (
	reflect "reflect"
	time "time"

	model "github.com/destafajri/system-pembayaran-spp-go-api/internal/domain/model"
	meta "github.com/destafajri/system-pembayaran-spp-go-api/meta"
	gomock "github.com/golang/mock/gomock"
)

// MockUserService is a mock of UserService interface.
type MockUserService struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceMockRecorder
}

// MockUserServiceMockRecorder is the mock recorder for MockUserService.
type MockUserServiceMockRecorder struct {
	mock *MockUserService
}

// NewMockUserService creates a new mock instance.
func NewMockUserService(ctrl *gomock.Controller) *MockUserService {
	mock := &MockUserService{ctrl: ctrl}
	mock.recorder = &MockUserServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserService) EXPECT() *MockUserServiceMockRecorder {
	return m.recorder
}

// ActivateUser mocks base method.
func (m *MockUserService) ActivateUser(arg0 string, arg1 time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ActivateUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ActivateUser indicates an expected call of ActivateUser.
func (mr *MockUserServiceMockRecorder) ActivateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActivateUser", reflect.TypeOf((*MockUserService)(nil).ActivateUser), arg0, arg1)
}

// CreateAdmin mocks base method.
func (m *MockUserService) CreateAdmin(arg0 *model.CreateAdminRequest, arg1 time.Time) (*model.CreateAdminResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAdmin", arg0, arg1)
	ret0, _ := ret[0].(*model.CreateAdminResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAdmin indicates an expected call of CreateAdmin.
func (mr *MockUserServiceMockRecorder) CreateAdmin(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAdmin", reflect.TypeOf((*MockUserService)(nil).CreateAdmin), arg0, arg1)
}

// DeactivateUser mocks base method.
func (m *MockUserService) DeactivateUser(arg0 string, arg1 time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeactivateUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeactivateUser indicates an expected call of DeactivateUser.
func (mr *MockUserServiceMockRecorder) DeactivateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeactivateUser", reflect.TypeOf((*MockUserService)(nil).DeactivateUser), arg0, arg1)
}

// GetDetailUser mocks base method.
func (m *MockUserService) GetDetailUser(arg0 string) (*model.GetDetailUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDetailUser", arg0)
	ret0, _ := ret[0].(*model.GetDetailUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDetailUser indicates an expected call of GetDetailUser.
func (mr *MockUserServiceMockRecorder) GetDetailUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDetailUser", reflect.TypeOf((*MockUserService)(nil).GetDetailUser), arg0)
}

// GetListUser mocks base method.
func (m *MockUserService) GetListUser(arg0 *meta.Metadata) ([]model.GetListUserResponse, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetListUser", arg0)
	ret0, _ := ret[0].([]model.GetListUserResponse)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetListUser indicates an expected call of GetListUser.
func (mr *MockUserServiceMockRecorder) GetListUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetListUser", reflect.TypeOf((*MockUserService)(nil).GetListUser), arg0)
}

// Login mocks base method.
func (m *MockUserService) Login(arg0 *model.LoginRequest) (*model.LoginResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", arg0)
	ret0, _ := ret[0].(*model.LoginResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockUserServiceMockRecorder) Login(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockUserService)(nil).Login), arg0)
}