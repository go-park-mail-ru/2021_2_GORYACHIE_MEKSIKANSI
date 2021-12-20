// Code generated by MockGen. DO NOT EDIT.
// Source: 2021_2_GORYACHIE_MEKSIKANSI/internals/microservice/authorization/application (interfaces: AuthorizationInterface)

// Package mocks is a generated GoMock package.
package mocks

import (
	authorization "2021_2_GORYACHIE_MEKSIKANSI/internals/microservice/authorization"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockAuthorizationInterface is a mock of AuthorizationInterface interface.
type MockAuthorizationInterface struct {
	ctrl     *gomock.Controller
	recorder *MockAuthorizationInterfaceMockRecorder
}

// MockAuthorizationInterfaceMockRecorder is the mock recorder for MockAuthorizationInterface.
type MockAuthorizationInterfaceMockRecorder struct {
	mock *MockAuthorizationInterface
}

// NewMockAuthorizationInterface creates a new mock instance.
func NewMockAuthorizationInterface(ctrl *gomock.Controller) *MockAuthorizationInterface {
	mock := &MockAuthorizationInterface{ctrl: ctrl}
	mock.recorder = &MockAuthorizationInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthorizationInterface) EXPECT() *MockAuthorizationInterfaceMockRecorder {
	return m.recorder
}

// CheckAccess mocks base method.
func (m *MockAuthorizationInterface) CheckAccess(arg0 *authorization.Defense) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckAccess", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckAccess indicates an expected call of CheckAccess.
func (mr *MockAuthorizationInterfaceMockRecorder) CheckAccess(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckAccess", reflect.TypeOf((*MockAuthorizationInterface)(nil).CheckAccess), arg0)
}

// GetIdByCookie mocks base method.
func (m *MockAuthorizationInterface) GetIdByCookie(arg0 *authorization.Defense) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIdByCookie", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIdByCookie indicates an expected call of GetIdByCookie.
func (mr *MockAuthorizationInterfaceMockRecorder) GetIdByCookie(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIdByCookie", reflect.TypeOf((*MockAuthorizationInterface)(nil).GetIdByCookie), arg0)
}

// Login mocks base method.
func (m *MockAuthorizationInterface) Login(arg0 *authorization.Authorization) (*authorization.Defense, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", arg0)
	ret0, _ := ret[0].(*authorization.Defense)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockAuthorizationInterfaceMockRecorder) Login(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockAuthorizationInterface)(nil).Login), arg0)
}

// Logout mocks base method.
func (m *MockAuthorizationInterface) Logout(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Logout", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Logout indicates an expected call of Logout.
func (mr *MockAuthorizationInterfaceMockRecorder) Logout(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logout", reflect.TypeOf((*MockAuthorizationInterface)(nil).Logout), arg0)
}

// NewCSRF mocks base method.
func (m *MockAuthorizationInterface) NewCSRF(arg0 *authorization.Defense) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewCSRF", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewCSRF indicates an expected call of NewCSRF.
func (mr *MockAuthorizationInterfaceMockRecorder) NewCSRF(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewCSRF", reflect.TypeOf((*MockAuthorizationInterface)(nil).NewCSRF), arg0)
}

// NewCSRFWebsocket mocks base method.
func (m *MockAuthorizationInterface) NewCSRFWebsocket(arg0 int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewCSRFWebsocket", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewCSRFWebsocket indicates an expected call of NewCSRFWebsocket.
func (mr *MockAuthorizationInterfaceMockRecorder) NewCSRFWebsocket(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewCSRFWebsocket", reflect.TypeOf((*MockAuthorizationInterface)(nil).NewCSRFWebsocket), arg0)
}

// SignUp mocks base method.
func (m *MockAuthorizationInterface) SignUp(arg0 *authorization.RegistrationRequest) (*authorization.Defense, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignUp", arg0)
	ret0, _ := ret[0].(*authorization.Defense)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignUp indicates an expected call of SignUp.
func (mr *MockAuthorizationInterfaceMockRecorder) SignUp(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignUp", reflect.TypeOf((*MockAuthorizationInterface)(nil).SignUp), arg0)
}
