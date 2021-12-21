// Code generated by MockGen. DO NOT EDIT.
// Source: 2021_2_GORYACHIE_MEKSIKANSI/internals/order/application (interfaces: OrderApplicationInterface)

// Package mocks is a generated GoMock package.
package mocks

import (
	order "2021_2_GORYACHIE_MEKSIKANSI/internals/order"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockOrderApplicationInterface is a mock of OrderApplicationInterface interface.
type MockOrderApplicationInterface struct {
	ctrl     *gomock.Controller
	recorder *MockOrderApplicationInterfaceMockRecorder
}

// MockOrderApplicationInterfaceMockRecorder is the mock recorder for MockOrderApplicationInterface.
type MockOrderApplicationInterfaceMockRecorder struct {
	mock *MockOrderApplicationInterface
}

// NewMockOrderApplicationInterface creates a new mock instance.
func NewMockOrderApplicationInterface(ctrl *gomock.Controller) *MockOrderApplicationInterface {
	mock := &MockOrderApplicationInterface{ctrl: ctrl}
	mock.recorder = &MockOrderApplicationInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrderApplicationInterface) EXPECT() *MockOrderApplicationInterfaceMockRecorder {
	return m.recorder
}

// CreateOrder mocks base method.
func (m *MockOrderApplicationInterface) CreateOrder(arg0 int, arg1 order.CreateOrder) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrder", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrder indicates an expected call of CreateOrder.
func (mr *MockOrderApplicationInterfaceMockRecorder) CreateOrder(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrder", reflect.TypeOf((*MockOrderApplicationInterface)(nil).CreateOrder), arg0, arg1)
}

// GetActiveOrder mocks base method.
func (m *MockOrderApplicationInterface) GetActiveOrder(arg0, arg1 int) (*order.ActiveOrder, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetActiveOrder", arg0, arg1)
	ret0, _ := ret[0].(*order.ActiveOrder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetActiveOrder indicates an expected call of GetActiveOrder.
func (mr *MockOrderApplicationInterfaceMockRecorder) GetActiveOrder(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActiveOrder", reflect.TypeOf((*MockOrderApplicationInterface)(nil).GetActiveOrder), arg0, arg1)
}

// GetOrders mocks base method.
func (m *MockOrderApplicationInterface) GetOrders(arg0 int) (*order.HistoryOrderArray, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrders", arg0)
	ret0, _ := ret[0].(*order.HistoryOrderArray)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrders indicates an expected call of GetOrders.
func (mr *MockOrderApplicationInterfaceMockRecorder) GetOrders(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrders", reflect.TypeOf((*MockOrderApplicationInterface)(nil).GetOrders), arg0)
}

// UpdateStatusOrder mocks base method.
func (m *MockOrderApplicationInterface) UpdateStatusOrder(arg0, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStatusOrder", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateStatusOrder indicates an expected call of UpdateStatusOrder.
func (mr *MockOrderApplicationInterfaceMockRecorder) UpdateStatusOrder(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStatusOrder", reflect.TypeOf((*MockOrderApplicationInterface)(nil).UpdateStatusOrder), arg0, arg1)
}
