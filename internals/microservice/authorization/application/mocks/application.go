// Code generated by MockGen. DO NOT EDIT.
// Source: 2021_2_GORYACHIE_MEKSIKANSI/internals/microservice/authorization/orm (interfaces: WrapperAuthorizationInterface)

// Package mocks is a generated GoMock package.
package mocks

import (
	authorization "2021_2_GORYACHIE_MEKSIKANSI/internals/microservice/authorization"
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

// AddCookie mocks base method.
func (m *MockWrapperAuthorizationInterface) AddCookie(arg0 *authorization.Defense, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddCookie", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddCookie indicates an expected call of AddCookie.
func (mr *MockWrapperAuthorizationInterfaceMockRecorder) AddCookie(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddCookie", reflect.TypeOf((*MockWrapperAuthorizationInterface)(nil).AddCookie), arg0, arg1)
}

// CheckAccess mocks base method.
func (m *MockWrapperAuthorizationInterface) CheckAccess(arg0 *authorization.Defense) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckAccess", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckAccess indicates an expected call of CheckAccess.
func (mr *MockWrapperAuthorizationInterfaceMockRecorder) CheckAccess(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckAccess", reflect.TypeOf((*MockWrapperAuthorizationInterface)(nil).CheckAccess), arg0)
}

// DeleteCookie mocks base method.
func (m *MockWrapperAuthorizationInterface) DeleteCookie(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCookie", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteCookie indicates an expected call of DeleteCookie.
func (mr *MockWrapperAuthorizationInterfaceMockRecorder) DeleteCookie(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCookie", reflect.TypeOf((*MockWrapperAuthorizationInterface)(nil).DeleteCookie), arg0)
}

// GetIdByCookie mocks base method.
func (m *MockWrapperAuthorizationInterface) GetIdByCookie(arg0 *authorization.Defense) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIdByCookie", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIdByCookie indicates an expected call of GetIdByCookie.
func (mr *MockWrapperAuthorizationInterfaceMockRecorder) GetIdByCookie(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIdByCookie", reflect.TypeOf((*MockWrapperAuthorizationInterface)(nil).GetIdByCookie), arg0)
}

// LoginByEmail mocks base method.
func (m *MockWrapperAuthorizationInterface) LoginByEmail(arg0, arg1 string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoginByEmail", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoginByEmail indicates an expected call of LoginByEmail.
func (mr *MockWrapperAuthorizationInterfaceMockRecorder) LoginByEmail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoginByEmail", reflect.TypeOf((*MockWrapperAuthorizationInterface)(nil).LoginByEmail), arg0, arg1)
}

// LoginByPhone mocks base method.
func (m *MockWrapperAuthorizationInterface) LoginByPhone(arg0, arg1 string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoginByPhone", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoginByPhone indicates an expected call of LoginByPhone.
func (mr *MockWrapperAuthorizationInterfaceMockRecorder) LoginByPhone(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoginByPhone", reflect.TypeOf((*MockWrapperAuthorizationInterface)(nil).LoginByPhone), arg0, arg1)
}

// NewCSRF mocks base method.
func (m *MockWrapperAuthorizationInterface) NewCSRF(arg0 *authorization.Defense) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewCSRF", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewCSRF indicates an expected call of NewCSRF.
func (mr *MockWrapperAuthorizationInterfaceMockRecorder) NewCSRF(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewCSRF", reflect.TypeOf((*MockWrapperAuthorizationInterface)(nil).NewCSRF), arg0)
}

// NewDefense mocks base method.
func (m *MockWrapperAuthorizationInterface) NewDefense() *authorization.Defense {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewDefense")
	ret0, _ := ret[0].(*authorization.Defense)
	return ret0
}

// NewDefense indicates an expected call of NewDefense.
func (mr *MockWrapperAuthorizationInterfaceMockRecorder) NewDefense() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewDefense", reflect.TypeOf((*MockWrapperAuthorizationInterface)(nil).NewDefense))
}

// SignupClient mocks base method.
func (m *MockWrapperAuthorizationInterface) SignupClient(arg0 *authorization.RegistrationRequest, arg1 *authorization.Defense) (*authorization.Defense, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignupClient", arg0, arg1)
	ret0, _ := ret[0].(*authorization.Defense)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignupClient indicates an expected call of SignupClient.
func (mr *MockWrapperAuthorizationInterfaceMockRecorder) SignupClient(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignupClient", reflect.TypeOf((*MockWrapperAuthorizationInterface)(nil).SignupClient), arg0, arg1)
}

// SignupCourier mocks base method.
func (m *MockWrapperAuthorizationInterface) SignupCourier(arg0 *authorization.RegistrationRequest, arg1 *authorization.Defense) (*authorization.Defense, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignupCourier", arg0, arg1)
	ret0, _ := ret[0].(*authorization.Defense)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignupCourier indicates an expected call of SignupCourier.
func (mr *MockWrapperAuthorizationInterfaceMockRecorder) SignupCourier(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignupCourier", reflect.TypeOf((*MockWrapperAuthorizationInterface)(nil).SignupCourier), arg0, arg1)
}

// SignupHost mocks base method.
func (m *MockWrapperAuthorizationInterface) SignupHost(arg0 *authorization.RegistrationRequest, arg1 *authorization.Defense) (*authorization.Defense, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignupHost", arg0, arg1)
	ret0, _ := ret[0].(*authorization.Defense)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignupHost indicates an expected call of SignupHost.
func (mr *MockWrapperAuthorizationInterfaceMockRecorder) SignupHost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignupHost", reflect.TypeOf((*MockWrapperAuthorizationInterface)(nil).SignupHost), arg0, arg1)
}
