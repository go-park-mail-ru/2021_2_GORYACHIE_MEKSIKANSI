// Code generated by MockGen. DO NOT EDIT.
// Source: 2021_2_GORYACHIE_MEKSIKANSI/internals/microservice/promocode/orm (interfaces: WrapperPromocodeInterface,ConnectionInterface,TransactionInterface)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	pgconn "github.com/jackc/pgconn"
	pgx "github.com/jackc/pgx/v4"
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
