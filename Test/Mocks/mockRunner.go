// Code generated by MockGen. DO NOT EDIT.
// Source: 2021_2_GORYACHIE_MEKSIKANSI/Utils (interfaces: ConnectionInterface,WrapperRestaurant,WrapperProfile,WrapperAuthorization)

// Package mocks is a generated GoMock package.
package mocks

import (
	Utils "2021_2_GORYACHIE_MEKSIKANSI/Utils"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	pgconn "github.com/jackc/pgconn"
	pgx "github.com/jackc/pgx/v4"
)

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

// MockWrapperRestaurant is a mock of WrapperRestaurant interface.
type MockWrapperRestaurant struct {
	ctrl     *gomock.Controller
	recorder *MockWrapperRestaurantMockRecorder
}

// MockWrapperRestaurantMockRecorder is the mock recorder for MockWrapperRestaurant.
type MockWrapperRestaurantMockRecorder struct {
	mock *MockWrapperRestaurant
}

// NewMockWrapperRestaurant creates a new mock instance.
func NewMockWrapperRestaurant(ctrl *gomock.Controller) *MockWrapperRestaurant {
	mock := &MockWrapperRestaurant{ctrl: ctrl}
	mock.recorder = &MockWrapperRestaurantMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWrapperRestaurant) EXPECT() *MockWrapperRestaurantMockRecorder {
	return m.recorder
}

// GetRestaurants mocks base method.
func (m *MockWrapperRestaurant) GetRestaurants() ([]Utils.Restaurant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRestaurants")
	ret0, _ := ret[0].([]Utils.Restaurant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRestaurants indicates an expected call of GetRestaurants.
func (mr *MockWrapperRestaurantMockRecorder) GetRestaurants() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRestaurants", reflect.TypeOf((*MockWrapperRestaurant)(nil).GetRestaurants))
}

// MockWrapperProfile is a mock of WrapperProfile interface.
type MockWrapperProfile struct {
	ctrl     *gomock.Controller
	recorder *MockWrapperProfileMockRecorder
}

// MockWrapperProfileMockRecorder is the mock recorder for MockWrapperProfile.
type MockWrapperProfileMockRecorder struct {
	mock *MockWrapperProfile
}

// NewMockWrapperProfile creates a new mock instance.
func NewMockWrapperProfile(ctrl *gomock.Controller) *MockWrapperProfile {
	mock := &MockWrapperProfile{ctrl: ctrl}
	mock.recorder = &MockWrapperProfileMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWrapperProfile) EXPECT() *MockWrapperProfileMockRecorder {
	return m.recorder
}

// GetProfileClient mocks base method.
func (m *MockWrapperProfile) GetProfileClient(arg0 int) (*Utils.Profile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProfileClient", arg0)
	ret0, _ := ret[0].(*Utils.Profile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProfileClient indicates an expected call of GetProfileClient.
func (mr *MockWrapperProfileMockRecorder) GetProfileClient(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProfileClient", reflect.TypeOf((*MockWrapperProfile)(nil).GetProfileClient), arg0)
}

// GetProfileCourier mocks base method.
func (m *MockWrapperProfile) GetProfileCourier(arg0 int) (*Utils.Profile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProfileCourier", arg0)
	ret0, _ := ret[0].(*Utils.Profile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProfileCourier indicates an expected call of GetProfileCourier.
func (mr *MockWrapperProfileMockRecorder) GetProfileCourier(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProfileCourier", reflect.TypeOf((*MockWrapperProfile)(nil).GetProfileCourier), arg0)
}

// GetProfileHost mocks base method.
func (m *MockWrapperProfile) GetProfileHost(arg0 int) (*Utils.Profile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProfileHost", arg0)
	ret0, _ := ret[0].(*Utils.Profile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProfileHost indicates an expected call of GetProfileHost.
func (mr *MockWrapperProfileMockRecorder) GetProfileHost(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProfileHost", reflect.TypeOf((*MockWrapperProfile)(nil).GetProfileHost), arg0)
}

// GetRoleById mocks base method.
func (m *MockWrapperProfile) GetRoleById(arg0 int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRoleById", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRoleById indicates an expected call of GetRoleById.
func (mr *MockWrapperProfileMockRecorder) GetRoleById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoleById", reflect.TypeOf((*MockWrapperProfile)(nil).GetRoleById), arg0)
}

// MockWrapperAuthorization is a mock of WrapperAuthorization interface.
type MockWrapperAuthorization struct {
	ctrl     *gomock.Controller
	recorder *MockWrapperAuthorizationMockRecorder
}

// MockWrapperAuthorizationMockRecorder is the mock recorder for MockWrapperAuthorization.
type MockWrapperAuthorizationMockRecorder struct {
	mock *MockWrapperAuthorization
}

// NewMockWrapperAuthorization creates a new mock instance.
func NewMockWrapperAuthorization(ctrl *gomock.Controller) *MockWrapperAuthorization {
	mock := &MockWrapperAuthorization{ctrl: ctrl}
	mock.recorder = &MockWrapperAuthorizationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWrapperAuthorization) EXPECT() *MockWrapperAuthorizationMockRecorder {
	return m.recorder
}

// AddCookie mocks base method.
func (m *MockWrapperAuthorization) AddCookie(arg0 *Utils.Defense, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddCookie", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddCookie indicates an expected call of AddCookie.
func (mr *MockWrapperAuthorizationMockRecorder) AddCookie(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddCookie", reflect.TypeOf((*MockWrapperAuthorization)(nil).AddCookie), arg0, arg1)
}

// DeleteCookie mocks base method.
func (m *MockWrapperAuthorization) DeleteCookie(arg0 *Utils.Defense) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCookie", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCookie indicates an expected call of DeleteCookie.
func (mr *MockWrapperAuthorizationMockRecorder) DeleteCookie(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCookie", reflect.TypeOf((*MockWrapperAuthorization)(nil).DeleteCookie), arg0)
}

// GenerateNew mocks base method.
func (m *MockWrapperAuthorization) GenerateNew() *Utils.Defense {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateNew")
	ret0, _ := ret[0].(*Utils.Defense)
	return ret0
}

// GenerateNew indicates an expected call of GenerateNew.
func (mr *MockWrapperAuthorizationMockRecorder) GenerateNew() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateNew", reflect.TypeOf((*MockWrapperAuthorization)(nil).GenerateNew))
}

// LoginByEmail mocks base method.
func (m *MockWrapperAuthorization) LoginByEmail(arg0, arg1 string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoginByEmail", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoginByEmail indicates an expected call of LoginByEmail.
func (mr *MockWrapperAuthorizationMockRecorder) LoginByEmail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoginByEmail", reflect.TypeOf((*MockWrapperAuthorization)(nil).LoginByEmail), arg0, arg1)
}

// LoginByPhone mocks base method.
func (m *MockWrapperAuthorization) LoginByPhone(arg0, arg1 string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoginByPhone", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoginByPhone indicates an expected call of LoginByPhone.
func (mr *MockWrapperAuthorizationMockRecorder) LoginByPhone(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoginByPhone", reflect.TypeOf((*MockWrapperAuthorization)(nil).LoginByPhone), arg0, arg1)
}

// SignupClient mocks base method.
func (m *MockWrapperAuthorization) SignupClient(arg0 *Utils.RegistrationRequest) (*Utils.Defense, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignupClient", arg0)
	ret0, _ := ret[0].(*Utils.Defense)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignupClient indicates an expected call of SignupClient.
func (mr *MockWrapperAuthorizationMockRecorder) SignupClient(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignupClient", reflect.TypeOf((*MockWrapperAuthorization)(nil).SignupClient), arg0)
}

// SignupCourier mocks base method.
func (m *MockWrapperAuthorization) SignupCourier(arg0 *Utils.RegistrationRequest) (*Utils.Defense, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignupCourier", arg0)
	ret0, _ := ret[0].(*Utils.Defense)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignupCourier indicates an expected call of SignupCourier.
func (mr *MockWrapperAuthorizationMockRecorder) SignupCourier(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignupCourier", reflect.TypeOf((*MockWrapperAuthorization)(nil).SignupCourier), arg0)
}

// SignupHost mocks base method.
func (m *MockWrapperAuthorization) SignupHost(arg0 *Utils.RegistrationRequest) (*Utils.Defense, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignupHost", arg0)
	ret0, _ := ret[0].(*Utils.Defense)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignupHost indicates an expected call of SignupHost.
func (mr *MockWrapperAuthorizationMockRecorder) SignupHost(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignupHost", reflect.TypeOf((*MockWrapperAuthorization)(nil).SignupHost), arg0)
}