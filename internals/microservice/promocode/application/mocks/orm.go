// Code generated by MockGen. DO NOT EDIT.
// Source: 2021_2_GORYACHIE_MEKSIKANSI/internals/microservice/promocode/orm (interfaces: WrapperPromocodeInterface)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockWrapperPromocodeInterface is a mock of WrapperPromocodeInterface interface.
type MockWrapperPromocodeInterface struct {
	ctrl     *gomock.Controller
	recorder *MockWrapperPromocodeInterfaceMockRecorder
}

// MockWrapperPromocodeInterfaceMockRecorder is the mock recorder for MockWrapperPromocodeInterface.
type MockWrapperPromocodeInterfaceMockRecorder struct {
	mock *MockWrapperPromocodeInterface
}

// NewMockWrapperPromocodeInterface creates a new mock instance.
func NewMockWrapperPromocodeInterface(ctrl *gomock.Controller) *MockWrapperPromocodeInterface {
	mock := &MockWrapperPromocodeInterface{ctrl: ctrl}
	mock.recorder = &MockWrapperPromocodeInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWrapperPromocodeInterface) EXPECT() *MockWrapperPromocodeInterfaceMockRecorder {
	return m.recorder
}

// ActiveCostForFreeDish mocks base method.
func (m *MockWrapperPromocodeInterface) ActiveCostForFreeDish(arg0 string, arg1 int) (int, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ActiveCostForFreeDish", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ActiveCostForFreeDish indicates an expected call of ActiveCostForFreeDish.
func (mr *MockWrapperPromocodeInterfaceMockRecorder) ActiveCostForFreeDish(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActiveCostForFreeDish", reflect.TypeOf((*MockWrapperPromocodeInterface)(nil).ActiveCostForFreeDish), arg0, arg1)
}

// ActiveCostForSale mocks base method.
func (m *MockWrapperPromocodeInterface) ActiveCostForSale(arg0 string, arg1, arg2 int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ActiveCostForSale", arg0, arg1, arg2)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ActiveCostForSale indicates an expected call of ActiveCostForSale.
func (mr *MockWrapperPromocodeInterfaceMockRecorder) ActiveCostForSale(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActiveCostForSale", reflect.TypeOf((*MockWrapperPromocodeInterface)(nil).ActiveCostForSale), arg0, arg1, arg2)
}

// ActiveFreeDelivery mocks base method.
func (m *MockWrapperPromocodeInterface) ActiveFreeDelivery(arg0 string, arg1 int) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ActiveFreeDelivery", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ActiveFreeDelivery indicates an expected call of ActiveFreeDelivery.
func (mr *MockWrapperPromocodeInterfaceMockRecorder) ActiveFreeDelivery(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActiveFreeDelivery", reflect.TypeOf((*MockWrapperPromocodeInterface)(nil).ActiveFreeDelivery), arg0, arg1)
}

// ActiveTimeForSale mocks base method.
func (m *MockWrapperPromocodeInterface) ActiveTimeForSale(arg0 string, arg1, arg2 int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ActiveTimeForSale", arg0, arg1, arg2)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ActiveTimeForSale indicates an expected call of ActiveTimeForSale.
func (mr *MockWrapperPromocodeInterfaceMockRecorder) ActiveTimeForSale(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActiveTimeForSale", reflect.TypeOf((*MockWrapperPromocodeInterface)(nil).ActiveTimeForSale), arg0, arg1, arg2)
}

// GetTypePromoCode mocks base method.
func (m *MockWrapperPromocodeInterface) GetTypePromoCode(arg0 string, arg1 int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTypePromoCode", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTypePromoCode indicates an expected call of GetTypePromoCode.
func (mr *MockWrapperPromocodeInterfaceMockRecorder) GetTypePromoCode(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTypePromoCode", reflect.TypeOf((*MockWrapperPromocodeInterface)(nil).GetTypePromoCode), arg0, arg1)
}
