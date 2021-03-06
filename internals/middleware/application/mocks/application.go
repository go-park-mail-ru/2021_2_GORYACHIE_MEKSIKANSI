// Code generated by MockGen. DO NOT EDIT.
// Source: 2021_2_GORYACHIE_MEKSIKANSI/internals/middleware/orm (interfaces: WrapperMiddlewareInterface)

// Package mocks is a generated GoMock package.
package mocks

import (
	util "2021_2_GORYACHIE_MEKSIKANSI/internals/util"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockWrapperMiddlewareInterface is a mock of WrapperMiddlewareInterface interface.
type MockWrapperMiddlewareInterface struct {
	ctrl     *gomock.Controller
	recorder *MockWrapperMiddlewareInterfaceMockRecorder
}

// MockWrapperMiddlewareInterfaceMockRecorder is the mock recorder for MockWrapperMiddlewareInterface.
type MockWrapperMiddlewareInterfaceMockRecorder struct {
	mock *MockWrapperMiddlewareInterface
}

// NewMockWrapperMiddlewareInterface creates a new mock instance.
func NewMockWrapperMiddlewareInterface(ctrl *gomock.Controller) *MockWrapperMiddlewareInterface {
	mock := &MockWrapperMiddlewareInterface{ctrl: ctrl}
	mock.recorder = &MockWrapperMiddlewareInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWrapperMiddlewareInterface) EXPECT() *MockWrapperMiddlewareInterfaceMockRecorder {
	return m.recorder
}

// CheckAccess mocks base method.
func (m *MockWrapperMiddlewareInterface) CheckAccess(arg0 *util.Defense) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckAccess", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckAccess indicates an expected call of CheckAccess.
func (mr *MockWrapperMiddlewareInterfaceMockRecorder) CheckAccess(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckAccess", reflect.TypeOf((*MockWrapperMiddlewareInterface)(nil).CheckAccess), arg0)
}

// CheckAccessWebsocket mocks base method.
func (m *MockWrapperMiddlewareInterface) CheckAccessWebsocket(arg0 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckAccessWebsocket", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckAccessWebsocket indicates an expected call of CheckAccessWebsocket.
func (mr *MockWrapperMiddlewareInterfaceMockRecorder) CheckAccessWebsocket(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckAccessWebsocket", reflect.TypeOf((*MockWrapperMiddlewareInterface)(nil).CheckAccessWebsocket), arg0)
}

// GetIdByCookie mocks base method.
func (m *MockWrapperMiddlewareInterface) GetIdByCookie(arg0 *util.Defense) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIdByCookie", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIdByCookie indicates an expected call of GetIdByCookie.
func (mr *MockWrapperMiddlewareInterfaceMockRecorder) GetIdByCookie(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIdByCookie", reflect.TypeOf((*MockWrapperMiddlewareInterface)(nil).GetIdByCookie), arg0)
}

// NewCSRF mocks base method.
func (m *MockWrapperMiddlewareInterface) NewCSRF(arg0 *util.Defense) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewCSRF", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewCSRF indicates an expected call of NewCSRF.
func (mr *MockWrapperMiddlewareInterfaceMockRecorder) NewCSRF(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewCSRF", reflect.TypeOf((*MockWrapperMiddlewareInterface)(nil).NewCSRF), arg0)
}
