// Code generated by MockGen. DO NOT EDIT.
// Source: 2021_2_GORYACHIE_MEKSIKANSI/internal/middleware/orm (interfaces: WrapperMiddlewareInterface,ConnectionMiddlewareInterface,ConnectionInterface,TransactionInterface)

// Package mocks is a generated GoMock package.
package mocks

import (
	proto "2021_2_GORYACHIE_MEKSIKANSI/internal/microservice/authorization/proto"
	util "2021_2_GORYACHIE_MEKSIKANSI/internal/util"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	pgconn "github.com/jackc/pgconn"
	pgx "github.com/jackc/pgx/v4"
	grpc "google.golang.org/grpc"
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

// MockConnectionMiddlewareInterface is a mock of ConnectionMiddlewareInterface interface.
type MockConnectionMiddlewareInterface struct {
	ctrl     *gomock.Controller
	recorder *MockConnectionMiddlewareInterfaceMockRecorder
}

// MockConnectionMiddlewareInterfaceMockRecorder is the mock recorder for MockConnectionMiddlewareInterface.
type MockConnectionMiddlewareInterfaceMockRecorder struct {
	mock *MockConnectionMiddlewareInterface
}

// NewMockConnectionMiddlewareInterface creates a new mock instance.
func NewMockConnectionMiddlewareInterface(ctrl *gomock.Controller) *MockConnectionMiddlewareInterface {
	mock := &MockConnectionMiddlewareInterface{ctrl: ctrl}
	mock.recorder = &MockConnectionMiddlewareInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConnectionMiddlewareInterface) EXPECT() *MockConnectionMiddlewareInterfaceMockRecorder {
	return m.recorder
}

// CheckAccessUser mocks base method.
func (m *MockConnectionMiddlewareInterface) CheckAccessUser(arg0 context.Context, arg1 *proto.Defense, arg2 ...grpc.CallOption) (*proto.CheckAccess, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CheckAccessUser", varargs...)
	ret0, _ := ret[0].(*proto.CheckAccess)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckAccessUser indicates an expected call of CheckAccessUser.
func (mr *MockConnectionMiddlewareInterfaceMockRecorder) CheckAccessUser(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckAccessUser", reflect.TypeOf((*MockConnectionMiddlewareInterface)(nil).CheckAccessUser), varargs...)
}

// GetIdByCookie mocks base method.
func (m *MockConnectionMiddlewareInterface) GetIdByCookie(arg0 context.Context, arg1 *proto.Defense, arg2 ...grpc.CallOption) (*proto.IdClientResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetIdByCookie", varargs...)
	ret0, _ := ret[0].(*proto.IdClientResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIdByCookie indicates an expected call of GetIdByCookie.
func (mr *MockConnectionMiddlewareInterfaceMockRecorder) GetIdByCookie(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIdByCookie", reflect.TypeOf((*MockConnectionMiddlewareInterface)(nil).GetIdByCookie), varargs...)
}

// NewCSRFUser mocks base method.
func (m *MockConnectionMiddlewareInterface) NewCSRFUser(arg0 context.Context, arg1 *proto.Defense, arg2 ...grpc.CallOption) (*proto.CSRFResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "NewCSRFUser", varargs...)
	ret0, _ := ret[0].(*proto.CSRFResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewCSRFUser indicates an expected call of NewCSRFUser.
func (mr *MockConnectionMiddlewareInterfaceMockRecorder) NewCSRFUser(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewCSRFUser", reflect.TypeOf((*MockConnectionMiddlewareInterface)(nil).NewCSRFUser), varargs...)
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
