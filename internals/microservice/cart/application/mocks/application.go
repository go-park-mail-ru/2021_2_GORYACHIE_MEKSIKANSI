// Code generated by MockGen. DO NOT EDIT.
// Source: 2021_2_GORYACHIE_MEKSIKANSI/internals/microservice/cart/orm (interfaces: WrapperCartInterface)

// Package mocks is a generated GoMock package.
package mocks

import (
	cart "2021_2_GORYACHIE_MEKSIKANSI/internals/microservice/cart"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockWrapperCartInterface is a mock of WrapperCartInterface interface.
type MockWrapperCartInterface struct {
	ctrl     *gomock.Controller
	recorder *MockWrapperCartInterfaceMockRecorder
}

// MockWrapperCartInterfaceMockRecorder is the mock recorder for MockWrapperCartInterface.
type MockWrapperCartInterfaceMockRecorder struct {
	mock *MockWrapperCartInterface
}

// NewMockWrapperCartInterface creates a new mock instance.
func NewMockWrapperCartInterface(ctrl *gomock.Controller) *MockWrapperCartInterface {
	mock := &MockWrapperCartInterface{ctrl: ctrl}
	mock.recorder = &MockWrapperCartInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWrapperCartInterface) EXPECT() *MockWrapperCartInterfaceMockRecorder {
	return m.recorder
}

// AddPromoCode mocks base method.
func (m *MockWrapperCartInterface) AddPromoCode(arg0 string, arg1, arg2 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddPromoCode", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddPromoCode indicates an expected call of AddPromoCode.
func (mr *MockWrapperCartInterfaceMockRecorder) AddPromoCode(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPromoCode", reflect.TypeOf((*MockWrapperCartInterface)(nil).AddPromoCode), arg0, arg1, arg2)
}

// DeleteCart mocks base method.
func (m *MockWrapperCartInterface) DeleteCart(arg0 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCart", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCart indicates an expected call of DeleteCart.
func (mr *MockWrapperCartInterfaceMockRecorder) DeleteCart(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCart", reflect.TypeOf((*MockWrapperCartInterface)(nil).DeleteCart), arg0)
}

// DoPromoCode mocks base method.
func (m *MockWrapperCartInterface) DoPromoCode(arg0 string, arg1 int, arg2 *cart.ResponseCartErrors) (*cart.ResponseCartErrors, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DoPromoCode", arg0, arg1, arg2)
	ret0, _ := ret[0].(*cart.ResponseCartErrors)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DoPromoCode indicates an expected call of DoPromoCode.
func (mr *MockWrapperCartInterfaceMockRecorder) DoPromoCode(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoPromoCode", reflect.TypeOf((*MockWrapperCartInterface)(nil).DoPromoCode), arg0, arg1, arg2)
}

// GetCart mocks base method.
func (m *MockWrapperCartInterface) GetCart(arg0 int) (*cart.ResponseCartErrors, []cart.CastDishesErrs, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCart", arg0)
	ret0, _ := ret[0].(*cart.ResponseCartErrors)
	ret1, _ := ret[1].([]cart.CastDishesErrs)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetCart indicates an expected call of GetCart.
func (mr *MockWrapperCartInterfaceMockRecorder) GetCart(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCart", reflect.TypeOf((*MockWrapperCartInterface)(nil).GetCart), arg0)
}

// GetPriceDelivery mocks base method.
func (m *MockWrapperCartInterface) GetPriceDelivery(arg0 int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPriceDelivery", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPriceDelivery indicates an expected call of GetPriceDelivery.
func (mr *MockWrapperCartInterfaceMockRecorder) GetPriceDelivery(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPriceDelivery", reflect.TypeOf((*MockWrapperCartInterface)(nil).GetPriceDelivery), arg0)
}

// GetPromoCode mocks base method.
func (m *MockWrapperCartInterface) GetPromoCode(arg0 int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPromoCode", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPromoCode indicates an expected call of GetPromoCode.
func (mr *MockWrapperCartInterfaceMockRecorder) GetPromoCode(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPromoCode", reflect.TypeOf((*MockWrapperCartInterface)(nil).GetPromoCode), arg0)
}

// GetRestaurant mocks base method.
func (m *MockWrapperCartInterface) GetRestaurant(arg0 int) (*cart.RestaurantId, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRestaurant", arg0)
	ret0, _ := ret[0].(*cart.RestaurantId)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRestaurant indicates an expected call of GetRestaurant.
func (mr *MockWrapperCartInterfaceMockRecorder) GetRestaurant(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRestaurant", reflect.TypeOf((*MockWrapperCartInterface)(nil).GetRestaurant), arg0)
}

// UpdateCart mocks base method.
func (m *MockWrapperCartInterface) UpdateCart(arg0 cart.RequestCartDefault, arg1 int) (*cart.ResponseCartErrors, []cart.CastDishesErrs, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCart", arg0, arg1)
	ret0, _ := ret[0].(*cart.ResponseCartErrors)
	ret1, _ := ret[1].([]cart.CastDishesErrs)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// UpdateCart indicates an expected call of UpdateCart.
func (mr *MockWrapperCartInterfaceMockRecorder) UpdateCart(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCart", reflect.TypeOf((*MockWrapperCartInterface)(nil).UpdateCart), arg0, arg1)
}
