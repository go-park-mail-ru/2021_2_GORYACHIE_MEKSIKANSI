// Code generated by MockGen. DO NOT EDIT.
// Source: 2021_2_GORYACHIE_MEKSIKANSI/internal/authorization/orm (interfaces: WrapperAuthorizationInterface,ConnectionInterface,ConnectAuthServiceInterface)

// Package mocks is a generated GoMock package.
package mocks

import (
	authorization "2021_2_GORYACHIE_MEKSIKANSI/internal/authorization"
	proto "2021_2_GORYACHIE_MEKSIKANSI/internal/microservice/authorization/proto"
	util "2021_2_GORYACHIE_MEKSIKANSI/internal/util"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	pgconn "github.com/jackc/pgconn"
	pgx "github.com/jackc/pgx/v4"
	grpc "google.golang.org/grpc"
)

// MockWrapperAuthorizationInterface is a mock of WrapperAuthorizationInterface interface.
type MockWrapperAuthorizationInterface struct {
	ctrl     *gomock.Controller
	recorder *MockWrapperAuthorizationInterfaceMockRecorder
}

// MockWrapperAuthorizationInterfaceMockRecorder is the mock recorder for MockWrapperAuthorizationInterface.
type MockWrapperAuthorizationInterfaceMockRecorder struct {
	mock *MockWrapperAuthorizationInterface
}

// NewMockWrapperAuthorizationInterface creates a new mock instance.
func NewMockWrapperAuthorizationInterface(ctrl *gomock.Controller) *MockWrapperAuthorizationInterface {
	mock := &MockWrapperAuthorizationInterface{ctrl: ctrl}
	mock.recorder = &MockWrapperAuthorizationInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWrapperAuthorizationInterface) EXPECT() *MockWrapperAuthorizationInterfaceMockRecorder {
	return m.recorder
}

// Login mocks base method.
func (m *MockWrapperAuthorizationInterface) Login(arg0 *authorization.Authorization) (*util.Defense, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", arg0)
	ret0, _ := ret[0].(*util.Defense)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockWrapperAuthorizationInterfaceMockRecorder) Login(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockWrapperAuthorizationInterface)(nil).Login), arg0)
}

// Logout mocks base method.
func (m *MockWrapperAuthorizationInterface) Logout(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Logout", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Logout indicates an expected call of Logout.
func (mr *MockWrapperAuthorizationInterfaceMockRecorder) Logout(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logout", reflect.TypeOf((*MockWrapperAuthorizationInterface)(nil).Logout), arg0)
}

// NewCSRFWebsocket mocks base method.
func (m *MockWrapperAuthorizationInterface) NewCSRFWebsocket(arg0 int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewCSRFWebsocket", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewCSRFWebsocket indicates an expected call of NewCSRFWebsocket.
func (mr *MockWrapperAuthorizationInterfaceMockRecorder) NewCSRFWebsocket(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewCSRFWebsocket", reflect.TypeOf((*MockWrapperAuthorizationInterface)(nil).NewCSRFWebsocket), arg0)
}

// SignUp mocks base method.
func (m *MockWrapperAuthorizationInterface) SignUp(arg0 *authorization.RegistrationRequest) (*util.Defense, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignUp", arg0)
	ret0, _ := ret[0].(*util.Defense)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignUp indicates an expected call of SignUp.
func (mr *MockWrapperAuthorizationInterfaceMockRecorder) SignUp(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignUp", reflect.TypeOf((*MockWrapperAuthorizationInterface)(nil).SignUp), arg0)
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

// MockConnectAuthServiceInterface is a mock of ConnectAuthServiceInterface interface.
type MockConnectAuthServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockConnectAuthServiceInterfaceMockRecorder
}

// MockConnectAuthServiceInterfaceMockRecorder is the mock recorder for MockConnectAuthServiceInterface.
type MockConnectAuthServiceInterfaceMockRecorder struct {
	mock *MockConnectAuthServiceInterface
}

// NewMockConnectAuthServiceInterface creates a new mock instance.
func NewMockConnectAuthServiceInterface(ctrl *gomock.Controller) *MockConnectAuthServiceInterface {
	mock := &MockConnectAuthServiceInterface{ctrl: ctrl}
	mock.recorder = &MockConnectAuthServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConnectAuthServiceInterface) EXPECT() *MockConnectAuthServiceInterfaceMockRecorder {
	return m.recorder
}

// Login mocks base method.
func (m *MockConnectAuthServiceInterface) Login(arg0 context.Context, arg1 *proto.Authorization, arg2 ...grpc.CallOption) (*proto.DefenseResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Login", varargs...)
	ret0, _ := ret[0].(*proto.DefenseResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockConnectAuthServiceInterfaceMockRecorder) Login(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockConnectAuthServiceInterface)(nil).Login), varargs...)
}

// Logout mocks base method.
func (m *MockConnectAuthServiceInterface) Logout(arg0 context.Context, arg1 *proto.CSRF, arg2 ...grpc.CallOption) (*proto.CSRFResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Logout", varargs...)
	ret0, _ := ret[0].(*proto.CSRFResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Logout indicates an expected call of Logout.
func (mr *MockConnectAuthServiceInterfaceMockRecorder) Logout(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logout", reflect.TypeOf((*MockConnectAuthServiceInterface)(nil).Logout), varargs...)
}

// SignUp mocks base method.
func (m *MockConnectAuthServiceInterface) SignUp(arg0 context.Context, arg1 *proto.RegistrationRequest, arg2 ...grpc.CallOption) (*proto.DefenseResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SignUp", varargs...)
	ret0, _ := ret[0].(*proto.DefenseResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignUp indicates an expected call of SignUp.
func (mr *MockConnectAuthServiceInterfaceMockRecorder) SignUp(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignUp", reflect.TypeOf((*MockConnectAuthServiceInterface)(nil).SignUp), varargs...)
}
