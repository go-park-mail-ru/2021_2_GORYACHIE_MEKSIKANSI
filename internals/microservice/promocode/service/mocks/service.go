// Code generated by MockGen. DO NOT EDIT.
// Source: 2021_2_GORYACHIE_MEKSIKANSI/internals/microservice/promocode/application (interfaces: PromocodeApplicationInterface)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockPromocodeApplicationInterface is a mock of PromocodeApplicationInterface interface.
type MockPromocodeApplicationInterface struct {
	ctrl     *gomock.Controller
	recorder *MockPromocodeApplicationInterfaceMockRecorder
}

// MockPromocodeApplicationInterfaceMockRecorder is the mock recorder for MockPromocodeApplicationInterface.
type MockPromocodeApplicationInterfaceMockRecorder struct {
	mock *MockPromocodeApplicationInterface
}

// NewMockPromocodeApplicationInterface creates a new mock instance.
func NewMockPromocodeApplicationInterface(ctrl *gomock.Controller) *MockPromocodeApplicationInterface {
	mock := &MockPromocodeApplicationInterface{ctrl: ctrl}
	mock.recorder = &MockPromocodeApplicationInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPromocodeApplicationInterface) EXPECT() *MockPromocodeApplicationInterfaceMockRecorder {
	return m.recorder
}

// ActiveCostForFreeDish mocks base method.
func (m *MockPromocodeApplicationInterface) ActiveCostForFreeDish(arg0 string, arg1 int) (int, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ActiveCostForFreeDish", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ActiveCostForFreeDish indicates an expected call of ActiveCostForFreeDish.
func (mr *MockPromocodeApplicationInterfaceMockRecorder) ActiveCostForFreeDish(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActiveCostForFreeDish", reflect.TypeOf((*MockPromocodeApplicationInterface)(nil).ActiveCostForFreeDish), arg0, arg1)
}

// ActiveCostForSale mocks base method.
func (m *MockPromocodeApplicationInterface) ActiveCostForSale(arg0 string, arg1, arg2 int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ActiveCostForSale", arg0, arg1, arg2)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ActiveCostForSale indicates an expected call of ActiveCostForSale.
func (mr *MockPromocodeApplicationInterfaceMockRecorder) ActiveCostForSale(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActiveCostForSale", reflect.TypeOf((*MockPromocodeApplicationInterface)(nil).ActiveCostForSale), arg0, arg1, arg2)
}

// ActiveFreeDelivery mocks base method.
func (m *MockPromocodeApplicationInterface) ActiveFreeDelivery(arg0 string, arg1 int) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ActiveFreeDelivery", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ActiveFreeDelivery indicates an expected call of ActiveFreeDelivery.
func (mr *MockPromocodeApplicationInterfaceMockRecorder) ActiveFreeDelivery(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActiveFreeDelivery", reflect.TypeOf((*MockPromocodeApplicationInterface)(nil).ActiveFreeDelivery), arg0, arg1)
}

// ActiveTimeForSale mocks base method.
func (m *MockPromocodeApplicationInterface) ActiveTimeForSale(arg0 string, arg1, arg2 int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ActiveTimeForSale", arg0, arg1, arg2)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ActiveTimeForSale indicates an expected call of ActiveTimeForSale.
func (mr *MockPromocodeApplicationInterfaceMockRecorder) ActiveTimeForSale(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActiveTimeForSale", reflect.TypeOf((*MockPromocodeApplicationInterface)(nil).ActiveTimeForSale), arg0, arg1, arg2)
}

// AddPromoCode mocks base method.
func (m *MockPromocodeApplicationInterface) AddPromoCode(arg0 string, arg1, arg2 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddPromoCode", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddPromoCode indicates an expected call of AddPromoCode.
func (mr *MockPromocodeApplicationInterfaceMockRecorder) AddPromoCode(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPromoCode", reflect.TypeOf((*MockPromocodeApplicationInterface)(nil).AddPromoCode), arg0, arg1, arg2)
}

// GetPromoCode mocks base method.
func (m *MockPromocodeApplicationInterface) GetPromoCode(arg0 int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPromoCode", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPromoCode indicates an expected call of GetPromoCode.
func (mr *MockPromocodeApplicationInterfaceMockRecorder) GetPromoCode(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPromoCode", reflect.TypeOf((*MockPromocodeApplicationInterface)(nil).GetPromoCode), arg0)
}

// GetTypePromoCode mocks base method.
func (m *MockPromocodeApplicationInterface) GetTypePromoCode(arg0 string, arg1 int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTypePromoCode", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTypePromoCode indicates an expected call of GetTypePromoCode.
func (mr *MockPromocodeApplicationInterfaceMockRecorder) GetTypePromoCode(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTypePromoCode", reflect.TypeOf((*MockPromocodeApplicationInterface)(nil).GetTypePromoCode), arg0, arg1)
}
