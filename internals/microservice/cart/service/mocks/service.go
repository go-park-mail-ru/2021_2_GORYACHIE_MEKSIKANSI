// Code generated by MockGen. DO NOT EDIT.
// Source: 2021_2_GORYACHIE_MEKSIKANSI/internals/microservice/cart/application (interfaces: CartInterface)

// Package mocks is a generated GoMock package.
package mocks

import (
	cart "2021_2_GORYACHIE_MEKSIKANSI/internals/microservice/cart"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockCartInterface is a mock of CartInterface interface.
type MockCartInterface struct {
	ctrl     *gomock.Controller
	recorder *MockCartInterfaceMockRecorder
}

// MockCartInterfaceMockRecorder is the mock recorder for MockCartInterface.
type MockCartInterfaceMockRecorder struct {
	mock *MockCartInterface
}

// NewMockCartInterface creates a new mock instance.
func NewMockCartInterface(ctrl *gomock.Controller) *MockCartInterface {
	mock := &MockCartInterface{ctrl: ctrl}
	mock.recorder = &MockCartInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCartInterface) EXPECT() *MockCartInterfaceMockRecorder {
	return m.recorder
}

// CalculateCost mocks base method.
func (m *MockCartInterface) CalculateCost(arg0 *cart.ResponseCartErrors, arg1 *cart.RestaurantId) (*cart.CostCartResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CalculateCost", arg0, arg1)
	ret0, _ := ret[0].(*cart.CostCartResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CalculateCost indicates an expected call of CalculateCost.
func (mr *MockCartInterfaceMockRecorder) CalculateCost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CalculateCost", reflect.TypeOf((*MockCartInterface)(nil).CalculateCost), arg0, arg1)
}

// CalculatePriceDelivery mocks base method.
func (m *MockCartInterface) CalculatePriceDelivery(arg0 int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CalculatePriceDelivery", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CalculatePriceDelivery indicates an expected call of CalculatePriceDelivery.
func (mr *MockCartInterfaceMockRecorder) CalculatePriceDelivery(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CalculatePriceDelivery", reflect.TypeOf((*MockCartInterface)(nil).CalculatePriceDelivery), arg0)
}

// GetCart mocks base method.
func (m *MockCartInterface) GetCart(arg0 int) (*cart.ResponseCartErrors, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCart", arg0)
	ret0, _ := ret[0].(*cart.ResponseCartErrors)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCart indicates an expected call of GetCart.
func (mr *MockCartInterfaceMockRecorder) GetCart(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCart", reflect.TypeOf((*MockCartInterface)(nil).GetCart), arg0)
}

// UpdateCart mocks base method.
func (m *MockCartInterface) UpdateCart(arg0 cart.RequestCartDefault, arg1 int) (*cart.ResponseCartErrors, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCart", arg0, arg1)
	ret0, _ := ret[0].(*cart.ResponseCartErrors)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCart indicates an expected call of UpdateCart.
func (mr *MockCartInterfaceMockRecorder) UpdateCart(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCart", reflect.TypeOf((*MockCartInterface)(nil).UpdateCart), arg0, arg1)
}
