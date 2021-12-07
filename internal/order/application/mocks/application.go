// Code generated by MockGen. DO NOT EDIT.
// Source: 2021_2_GORYACHIE_MEKSIKANSI/internal/order/orm (interfaces: WrapperOrderInterface)

// Package mocks is a generated GoMock package.
package mocks

import (
	cart "2021_2_GORYACHIE_MEKSIKANSI/internal/cart"
	order "2021_2_GORYACHIE_MEKSIKANSI/internal/order"
	restaurant "2021_2_GORYACHIE_MEKSIKANSI/internal/restaurant"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockWrapperOrderInterface is a mock of WrapperOrderInterface interface.
type MockWrapperOrderInterface struct {
	ctrl     *gomock.Controller
	recorder *MockWrapperOrderInterfaceMockRecorder
}

// MockWrapperOrderInterfaceMockRecorder is the mock recorder for MockWrapperOrderInterface.
type MockWrapperOrderInterfaceMockRecorder struct {
	mock *MockWrapperOrderInterface
}

// NewMockWrapperOrderInterface creates a new mock instance.
func NewMockWrapperOrderInterface(ctrl *gomock.Controller) *MockWrapperOrderInterface {
	mock := &MockWrapperOrderInterface{ctrl: ctrl}
	mock.recorder = &MockWrapperOrderInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWrapperOrderInterface) EXPECT() *MockWrapperOrderInterfaceMockRecorder {
	return m.recorder
}

// CheckRun mocks base method.
func (m *MockWrapperOrderInterface) CheckRun(arg0 int) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckRun", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckRun indicates an expected call of CheckRun.
func (mr *MockWrapperOrderInterfaceMockRecorder) CheckRun(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckRun", reflect.TypeOf((*MockWrapperOrderInterface)(nil).CheckRun), arg0)
}

// CreateOrder mocks base method.
func (m *MockWrapperOrderInterface) CreateOrder(arg0 int, arg1 order.CreateOrder, arg2 int, arg3 cart.ResponseCartErrors, arg4 int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrder", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrder indicates an expected call of CreateOrder.
func (mr *MockWrapperOrderInterfaceMockRecorder) CreateOrder(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrder", reflect.TypeOf((*MockWrapperOrderInterface)(nil).CreateOrder), arg0, arg1, arg2, arg3, arg4)
}

// DeleteCart mocks base method.
func (m *MockWrapperOrderInterface) DeleteCart(arg0 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCart", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCart indicates an expected call of DeleteCart.
func (mr *MockWrapperOrderInterfaceMockRecorder) DeleteCart(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCart", reflect.TypeOf((*MockWrapperOrderInterface)(nil).DeleteCart), arg0)
}

// GetCart mocks base method.
func (m *MockWrapperOrderInterface) GetCart(arg0 int) (*cart.ResponseCartErrors, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCart", arg0)
	ret0, _ := ret[0].(*cart.ResponseCartErrors)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCart indicates an expected call of GetCart.
func (mr *MockWrapperOrderInterfaceMockRecorder) GetCart(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCart", reflect.TypeOf((*MockWrapperOrderInterface)(nil).GetCart), arg0)
}

// GetOrder mocks base method.
func (m *MockWrapperOrderInterface) GetOrder(arg0, arg1 int) (*order.ActiveOrder, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrder", arg0, arg1)
	ret0, _ := ret[0].(*order.ActiveOrder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrder indicates an expected call of GetOrder.
func (mr *MockWrapperOrderInterfaceMockRecorder) GetOrder(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrder", reflect.TypeOf((*MockWrapperOrderInterface)(nil).GetOrder), arg0, arg1)
}

// GetOrders mocks base method.
func (m *MockWrapperOrderInterface) GetOrders(arg0 int) (*order.HistoryOrderArray, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrders", arg0)
	ret0, _ := ret[0].(*order.HistoryOrderArray)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrders indicates an expected call of GetOrders.
func (mr *MockWrapperOrderInterfaceMockRecorder) GetOrders(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrders", reflect.TypeOf((*MockWrapperOrderInterface)(nil).GetOrders), arg0)
}

// GetRestaurant mocks base method.
func (m *MockWrapperOrderInterface) GetRestaurant(arg0 int) (*restaurant.RestaurantId, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRestaurant", arg0)
	ret0, _ := ret[0].(*restaurant.RestaurantId)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRestaurant indicates an expected call of GetRestaurant.
func (mr *MockWrapperOrderInterfaceMockRecorder) GetRestaurant(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRestaurant", reflect.TypeOf((*MockWrapperOrderInterface)(nil).GetRestaurant), arg0)
}

// UpdateStatusOrder mocks base method.
func (m *MockWrapperOrderInterface) UpdateStatusOrder(arg0, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStatusOrder", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateStatusOrder indicates an expected call of UpdateStatusOrder.
func (mr *MockWrapperOrderInterfaceMockRecorder) UpdateStatusOrder(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStatusOrder", reflect.TypeOf((*MockWrapperOrderInterface)(nil).UpdateStatusOrder), arg0, arg1)
}