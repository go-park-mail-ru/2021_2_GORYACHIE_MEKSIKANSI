// Code generated by MockGen. DO NOT EDIT.
// Source: 2021_2_GORYACHIE_MEKSIKANSI/internal/order/orm (interfaces: WrapperOrderInterface,ConnectionInterface)

// Package mocks is a generated GoMock package.
package mocks

import (
	cart "2021_2_GORYACHIE_MEKSIKANSI/internal/cart"
	order "2021_2_GORYACHIE_MEKSIKANSI/internal/order"
	restaurant "2021_2_GORYACHIE_MEKSIKANSI/internal/restaurant"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	pgconn "github.com/jackc/pgconn"
	pgx "github.com/jackc/pgx/v4"
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

// MockConnectionInterface is a mock of ConnectionInterface interface.
type MockConnectionInterface struct {
	ctrl     *gomock.Controller
	recorder *MockConnectionInterfaceMockRecorder
}

// MockConnectionInterfaceMockRecorder is the mock recorder for MockConnectionInterface.
type MockConnectionInterfaceMockRecorder struct {
	mock *MockConnectionInterface
}

// NewMockConnectionInterface creates a new mock instance.
func NewMockConnectionInterface(ctrl *gomock.Controller) *MockConnectionInterface {
	mock := &MockConnectionInterface{ctrl: ctrl}
	mock.recorder = &MockConnectionInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConnectionInterface) EXPECT() *MockConnectionInterfaceMockRecorder {
	return m.recorder
}

// Begin mocks base method.
func (m *MockConnectionInterface) Begin(arg0 context.Context) (pgx.Tx, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Begin", arg0)
	ret0, _ := ret[0].(pgx.Tx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Begin indicates an expected call of Begin.
func (mr *MockConnectionInterfaceMockRecorder) Begin(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Begin", reflect.TypeOf((*MockConnectionInterface)(nil).Begin), arg0)
}

// Exec mocks base method.
func (m *MockConnectionInterface) Exec(arg0 context.Context, arg1 string, arg2 ...interface{}) (pgconn.CommandTag, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Exec", varargs...)
	ret0, _ := ret[0].(pgconn.CommandTag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exec indicates an expected call of Exec.
func (mr *MockConnectionInterfaceMockRecorder) Exec(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exec", reflect.TypeOf((*MockConnectionInterface)(nil).Exec), varargs...)
}

// Query mocks base method.
func (m *MockConnectionInterface) Query(arg0 context.Context, arg1 string, arg2 ...interface{}) (pgx.Rows, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Query", varargs...)
	ret0, _ := ret[0].(pgx.Rows)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Query indicates an expected call of Query.
func (mr *MockConnectionInterfaceMockRecorder) Query(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockConnectionInterface)(nil).Query), varargs...)
}

// QueryRow mocks base method.
func (m *MockConnectionInterface) QueryRow(arg0 context.Context, arg1 string, arg2 ...interface{}) pgx.Row {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "QueryRow", varargs...)
	ret0, _ := ret[0].(pgx.Row)
	return ret0
}

// QueryRow indicates an expected call of QueryRow.
func (mr *MockConnectionInterfaceMockRecorder) QueryRow(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryRow", reflect.TypeOf((*MockConnectionInterface)(nil).QueryRow), varargs...)
}
