// Code generated by MockGen. DO NOT EDIT.
// Source: 2021_2_GORYACHIE_MEKSIKANSI/internals/authorization/orm (interfaces: WrapperAuthorizationInterface)

// Package mocks is a generated GoMock package.
package mocks

import (
	authorization "2021_2_GORYACHIE_MEKSIKANSI/internals/authorization"
	util "2021_2_GORYACHIE_MEKSIKANSI/internals/util"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockWrapperAuthorizationInterface is a mock of WrapperAuthorizationInterface interface.
type MockWrapperAuthorizationInterface struct {
	ctrl     *gomock.Controller
	recorder *MockWrapperAuthorizationInterfaceMockRecorder
}

// MockWrapperAuthorizationInterfaceMockRecorder is the mock recorder for MockWrapperAuthorizationInterface.
type MockWrapperAuthorizationInterfaceMockRecorder struct {
	mock *MockWrapperAuthorizationInterface
}

// NewMockWrapperAuthorizationInterface creates a new mock instance.
func NewMockWrapperAuthorizationInterface(ctrl *gomock.Controller) *MockWrapperAuthorizationInterface {
	mock := &MockWrapperAuthorizationInterface{ctrl: ctrl}
	mock.recorder = &MockWrapperAuthorizationInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWrapperAuthorizationInterface) EXPECT() *MockWrapperAuthorizationInterfaceMockRecorder {
	return m.recorder
}

// Login mocks base method.
func (m *MockWrapperAuthorizationInterface) Login(arg0 *authorization.Authorization) (*util.Defense, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", arg0)
	ret0, _ := ret[0].(*util.Defense)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockWrapperAuthorizationInterfaceMockRecorder) Login(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockWrapperAuthorizationInterface)(nil).Login), arg0)
}

// Logout mocks base method.
func (m *MockWrapperAuthorizationInterface) Logout(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Logout", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Logout indicates an expected call of Logout.
func (mr *MockWrapperAuthorizationInterfaceMockRecorder) Logout(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logout", reflect.TypeOf((*MockWrapperAuthorizationInterface)(nil).Logout), arg0)
}

// NewCSRFWebsocket mocks base method.
func (m *MockWrapperAuthorizationInterface) NewCSRFWebsocket(arg0 int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewCSRFWebsocket", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewCSRFWebsocket indicates an expected call of NewCSRFWebsocket.
func (mr *MockWrapperAuthorizationInterfaceMockRecorder) NewCSRFWebsocket(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewCSRFWebsocket", reflect.TypeOf((*MockWrapperAuthorizationInterface)(nil).NewCSRFWebsocket), arg0)
}

// SignUp mocks base method.
func (m *MockWrapperAuthorizationInterface) SignUp(arg0 *authorization.RegistrationRequest) (*util.Defense, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignUp", arg0)
	ret0, _ := ret[0].(*util.Defense)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignUp indicates an expected call of SignUp.
func (mr *MockWrapperAuthorizationInterfaceMockRecorder) SignUp(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignUp", reflect.TypeOf((*MockWrapperAuthorizationInterface)(nil).SignUp), arg0)
}
