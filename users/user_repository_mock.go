// Code generated by MockGen. DO NOT EDIT.
// Source: users/user_repository.go

// Package users is a generated GoMock package.
package users

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockuserRepositoryInterface is a mock of userRepositoryInterface interface
type MockuserRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockuserRepositoryInterfaceMockRecorder
}

// MockuserRepositoryInterfaceMockRecorder is the mock recorder for MockuserRepositoryInterface
type MockuserRepositoryInterfaceMockRecorder struct {
	mock *MockuserRepositoryInterface
}

// NewMockuserRepositoryInterface creates a new mock instance
func NewMockuserRepositoryInterface(ctrl *gomock.Controller) *MockuserRepositoryInterface {
	mock := &MockuserRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockuserRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockuserRepositoryInterface) EXPECT() *MockuserRepositoryInterfaceMockRecorder {
	return m.recorder
}

// createUser mocks base method
func (m *MockuserRepositoryInterface) createUser(userstore user) error {
	ret := m.ctrl.Call(m, "createUser", userstore)
	ret0, _ := ret[0].(error)
	return ret0
}

// createUser indicates an expected call of createUser
func (mr *MockuserRepositoryInterfaceMockRecorder) createUser(userstore interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "createUser", reflect.TypeOf((*MockuserRepositoryInterface)(nil).createUser), userstore)
}

// loginUser mocks base method
func (m *MockuserRepositoryInterface) loginUser(loginData userv) ([]byte, []byte, error) {
	ret := m.ctrl.Call(m, "loginUser", loginData)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].([]byte)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// loginUser indicates an expected call of loginUser
func (mr *MockuserRepositoryInterfaceMockRecorder) loginUser(loginData interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "loginUser", reflect.TypeOf((*MockuserRepositoryInterface)(nil).loginUser), loginData)
}

// getAllUser mocks base method
func (m *MockuserRepositoryInterface) getAllUser() ([]user, error) {
	ret := m.ctrl.Call(m, "getAllUser")
	ret0, _ := ret[0].([]user)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// getAllUser indicates an expected call of getAllUser
func (mr *MockuserRepositoryInterfaceMockRecorder) getAllUser() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getAllUser", reflect.TypeOf((*MockuserRepositoryInterface)(nil).getAllUser))
}

// getUserData mocks base method
func (m *MockuserRepositoryInterface) getUserData(username string) ([]byte, error) {
	ret := m.ctrl.Call(m, "getUserData", username)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// getUserData indicates an expected call of getUserData
func (mr *MockuserRepositoryInterfaceMockRecorder) getUserData(username interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getUserData", reflect.TypeOf((*MockuserRepositoryInterface)(nil).getUserData), username)
}
