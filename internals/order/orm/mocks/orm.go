// Code generated by MockGen. DO NOT EDIT.
// Source: 2021_2_GORYACHIE_MEKSIKANSI/internals/order/orm (interfaces: WrapperOrderInterface,ConnectionInterface,TransactionInterface)

// Package mocks is a generated GoMock package.
package mocks

import (
	cart "2021_2_GORYACHIE_MEKSIKANSI/internals/cart"
	order "2021_2_GORYACHIE_MEKSIKANSI/internals/order"
	restaurant "2021_2_GORYACHIE_MEKSIKANSI/internals/restaurant"
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

// MockTransactionInterface is a mock of TransactionInterface interface.
type MockTransactionInterface struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionInterfaceMockRecorder
}

// MockTransactionInterfaceMockRecorder is the mock recorder for MockTransactionInterface.
type MockTransactionInterfaceMockRecorder struct {
	mock *MockTransactionInterface
}

// NewMockTransactionInterface creates a new mock instance.
func NewMockTransactionInterface(ctrl *gomock.Controller) *MockTransactionInterface {
	mock := &MockTransactionInterface{ctrl: ctrl}
	mock.recorder = &MockTransactionInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransactionInterface) EXPECT() *MockTransactionInterfaceMockRecorder {
	return m.recorder
}

// Begin mocks base method.
func (m *MockTransactionInterface) Begin(arg0 context.Context) (pgx.Tx, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Begin", arg0)
	ret0, _ := ret[0].(pgx.Tx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Begin indicates an expected call of Begin.
func (mr *MockTransactionInterfaceMockRecorder) Begin(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Begin", reflect.TypeOf((*MockTransactionInterface)(nil).Begin), arg0)
}

// BeginFunc mocks base method.
func (m *MockTransactionInterface) BeginFunc(arg0 context.Context, arg1 func(pgx.Tx) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BeginFunc", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// BeginFunc indicates an expected call of BeginFunc.
func (mr *MockTransactionInterfaceMockRecorder) BeginFunc(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeginFunc", reflect.TypeOf((*MockTransactionInterface)(nil).BeginFunc), arg0, arg1)
}

// Commit mocks base method.
func (m *MockTransactionInterface) Commit(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Commit", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Commit indicates an expected call of Commit.
func (mr *MockTransactionInterfaceMockRecorder) Commit(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockTransactionInterface)(nil).Commit), arg0)
}

// Conn mocks base method.
func (m *MockTransactionInterface) Conn() *pgx.Conn {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Conn")
	ret0, _ := ret[0].(*pgx.Conn)
	return ret0
}

// Conn indicates an expected call of Conn.
func (mr *MockTransactionInterfaceMockRecorder) Conn() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Conn", reflect.TypeOf((*MockTransactionInterface)(nil).Conn))
}

// CopyFrom mocks base method.
func (m *MockTransactionInterface) CopyFrom(arg0 context.Context, arg1 pgx.Identifier, arg2 []string, arg3 pgx.CopyFromSource) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CopyFrom", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CopyFrom indicates an expected call of CopyFrom.
func (mr *MockTransactionInterfaceMockRecorder) CopyFrom(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CopyFrom", reflect.TypeOf((*MockTransactionInterface)(nil).CopyFrom), arg0, arg1, arg2, arg3)
}

// Exec mocks base method.
func (m *MockTransactionInterface) Exec(arg0 context.Context, arg1 string, arg2 ...interface{}) (pgconn.CommandTag, error) {
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
func (mr *MockTransactionInterfaceMockRecorder) Exec(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exec", reflect.TypeOf((*MockTransactionInterface)(nil).Exec), varargs...)
}

// LargeObjects mocks base method.
func (m *MockTransactionInterface) LargeObjects() pgx.LargeObjects {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LargeObjects")
	ret0, _ := ret[0].(pgx.LargeObjects)
	return ret0
}

// LargeObjects indicates an expected call of LargeObjects.
func (mr *MockTransactionInterfaceMockRecorder) LargeObjects() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LargeObjects", reflect.TypeOf((*MockTransactionInterface)(nil).LargeObjects))
}

// Prepare mocks base method.
func (m *MockTransactionInterface) Prepare(arg0 context.Context, arg1, arg2 string) (*pgconn.StatementDescription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Prepare", arg0, arg1, arg2)
	ret0, _ := ret[0].(*pgconn.StatementDescription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Prepare indicates an expected call of Prepare.
func (mr *MockTransactionInterfaceMockRecorder) Prepare(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Prepare", reflect.TypeOf((*MockTransactionInterface)(nil).Prepare), arg0, arg1, arg2)
}

// Query mocks base method.
func (m *MockTransactionInterface) Query(arg0 context.Context, arg1 string, arg2 ...interface{}) (pgx.Rows, error) {
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
func (mr *MockTransactionInterfaceMockRecorder) Query(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockTransactionInterface)(nil).Query), varargs...)
}

// QueryFunc mocks base method.
func (m *MockTransactionInterface) QueryFunc(arg0 context.Context, arg1 string, arg2, arg3 []interface{}, arg4 func(pgx.QueryFuncRow) error) (pgconn.CommandTag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryFunc", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(pgconn.CommandTag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryFunc indicates an expected call of QueryFunc.
func (mr *MockTransactionInterfaceMockRecorder) QueryFunc(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryFunc", reflect.TypeOf((*MockTransactionInterface)(nil).QueryFunc), arg0, arg1, arg2, arg3, arg4)
}

// QueryRow mocks base method.
func (m *MockTransactionInterface) QueryRow(arg0 context.Context, arg1 string, arg2 ...interface{}) pgx.Row {
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
func (mr *MockTransactionInterfaceMockRecorder) QueryRow(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryRow", reflect.TypeOf((*MockTransactionInterface)(nil).QueryRow), varargs...)
}

// Rollback mocks base method.
func (m *MockTransactionInterface) Rollback(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Rollback", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Rollback indicates an expected call of Rollback.
func (mr *MockTransactionInterfaceMockRecorder) Rollback(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rollback", reflect.TypeOf((*MockTransactionInterface)(nil).Rollback), arg0)
}

// SendBatch mocks base method.
func (m *MockTransactionInterface) SendBatch(arg0 context.Context, arg1 *pgx.Batch) pgx.BatchResults {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendBatch", arg0, arg1)
	ret0, _ := ret[0].(pgx.BatchResults)
	return ret0
}

// SendBatch indicates an expected call of SendBatch.
func (mr *MockTransactionInterfaceMockRecorder) SendBatch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendBatch", reflect.TypeOf((*MockTransactionInterface)(nil).SendBatch), arg0, arg1)
}
