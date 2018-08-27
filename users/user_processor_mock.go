// Code generated by MockGen. DO NOT EDIT.
// Source: users/user_processor.go

// Package users is a generated GoMock package.
package users

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockuserProcessorInterface is a mock of userProcessorInterface interface
type MockuserProcessorInterface struct {
	ctrl     *gomock.Controller
	recorder *MockuserProcessorInterfaceMockRecorder
}

// MockuserProcessorInterfaceMockRecorder is the mock recorder for MockuserProcessorInterface
type MockuserProcessorInterfaceMockRecorder struct {
	mock *MockuserProcessorInterface
}

// NewMockuserProcessorInterface creates a new mock instance
func NewMockuserProcessorInterface(ctrl *gomock.Controller) *MockuserProcessorInterface {
	mock := &MockuserProcessorInterface{ctrl: ctrl}
	mock.recorder = &MockuserProcessorInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockuserProcessorInterface) EXPECT() *MockuserProcessorInterfaceMockRecorder {
	return m.recorder
}

// getRepository mocks base method
func (m *MockuserProcessorInterface) getRepository() userRepositoryInterface {
	ret := m.ctrl.Call(m, "getRepository")
	ret0, _ := ret[0].(userRepositoryInterface)
	return ret0
}

// getRepository indicates an expected call of getRepository
func (mr *MockuserProcessorInterfaceMockRecorder) getRepository() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getRepository", reflect.TypeOf((*MockuserProcessorInterface)(nil).getRepository))
}

// createUserProcessor mocks base method
func (m *MockuserProcessorInterface) createUserProcessor(userdata userv) error {
	ret := m.ctrl.Call(m, "createUserProcessor", userdata)
	ret0, _ := ret[0].(error)
	return ret0
}

// createUserProcessor indicates an expected call of createUserProcessor
func (mr *MockuserProcessorInterfaceMockRecorder) createUserProcessor(userdata interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "createUserProcessor", reflect.TypeOf((*MockuserProcessorInterface)(nil).createUserProcessor), userdata)
}

// createUserEncryptPassword mocks base method
func (m *MockuserProcessorInterface) createUserEncryptPassword(salt, password string) string {
	ret := m.ctrl.Call(m, "createUserEncryptPassword", salt, password)
	ret0, _ := ret[0].(string)
	return ret0
}

// createUserEncryptPassword indicates an expected call of createUserEncryptPassword
func (mr *MockuserProcessorInterfaceMockRecorder) createUserEncryptPassword(salt, password interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "createUserEncryptPassword", reflect.TypeOf((*MockuserProcessorInterface)(nil).createUserEncryptPassword), salt, password)
}

// loginUserProcessor mocks base method
func (m *MockuserProcessorInterface) loginUserProcessor(loginData userv) ([]byte, []byte, error) {
	ret := m.ctrl.Call(m, "loginUserProcessor", loginData)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].([]byte)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// loginUserProcessor indicates an expected call of loginUserProcessor
func (mr *MockuserProcessorInterfaceMockRecorder) loginUserProcessor(loginData interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "loginUserProcessor", reflect.TypeOf((*MockuserProcessorInterface)(nil).loginUserProcessor), loginData)
}

// getAllUserProcessor mocks base method
func (m *MockuserProcessorInterface) getAllUserProcessor() (string, error) {
	ret := m.ctrl.Call(m, "getAllUserProcessor")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// getAllUserProcessor indicates an expected call of getAllUserProcessor
func (mr *MockuserProcessorInterfaceMockRecorder) getAllUserProcessor() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getAllUserProcessor", reflect.TypeOf((*MockuserProcessorInterface)(nil).getAllUserProcessor))
}

// getUserDataProcessor mocks base method
func (m *MockuserProcessorInterface) getUserDataProcessor(username string) ([]byte, error) {
	ret := m.ctrl.Call(m, "getUserDataProcessor", username)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// getUserDataProcessor indicates an expected call of getUserDataProcessor
func (mr *MockuserProcessorInterfaceMockRecorder) getUserDataProcessor(username interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getUserDataProcessor", reflect.TypeOf((*MockuserProcessorInterface)(nil).getUserDataProcessor), username)
}
