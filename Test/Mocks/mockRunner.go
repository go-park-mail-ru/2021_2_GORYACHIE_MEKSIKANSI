// Code generated by MockGen. DO NOT EDIT.
// Source: 2021_2_GORYACHIE_MEKSIKANSI/Utils (interfaces: ConnectionInterface,WrapperRestaurant,WrapperProfile,WrapperAuthorization,WrapperCart)

// Package mocks is a generated GoMock package.
package mocks

import (
	Utils "2021_2_GORYACHIE_MEKSIKANSI/Utils"
	context "context"
	reflect "reflect"
	time "time"

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

// GetDishes mocks base method.
func (m *MockWrapperRestaurant) GetDishes(arg0, arg1 int) (*Utils.Dishes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDishes", arg0, arg1)
	ret0, _ := ret[0].(*Utils.Dishes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDishes indicates an expected call of GetDishes.
func (mr *MockWrapperRestaurantMockRecorder) GetDishes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDishes", reflect.TypeOf((*MockWrapperRestaurant)(nil).GetDishes), arg0, arg1)
}

// GetGeneralInfoRestaurant mocks base method.
func (m *MockWrapperRestaurant) GetGeneralInfoRestaurant(arg0 int) (*Utils.RestaurantId, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGeneralInfoRestaurant", arg0)
	ret0, _ := ret[0].(*Utils.RestaurantId)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGeneralInfoRestaurant indicates an expected call of GetGeneralInfoRestaurant.
func (mr *MockWrapperRestaurantMockRecorder) GetGeneralInfoRestaurant(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGeneralInfoRestaurant", reflect.TypeOf((*MockWrapperRestaurant)(nil).GetGeneralInfoRestaurant), arg0)
}

// GetMenu mocks base method.
func (m *MockWrapperRestaurant) GetMenu(arg0 int) ([]Utils.Menu, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMenu", arg0)
	ret0, _ := ret[0].([]Utils.Menu)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMenu indicates an expected call of GetMenu.
func (mr *MockWrapperRestaurantMockRecorder) GetMenu(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMenu", reflect.TypeOf((*MockWrapperRestaurant)(nil).GetMenu), arg0)
}

// GetRadios mocks base method.
func (m *MockWrapperRestaurant) GetRadios(arg0 int) ([]Utils.Radios, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRadios", arg0)
	ret0, _ := ret[0].([]Utils.Radios)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRadios indicates an expected call of GetRadios.
func (mr *MockWrapperRestaurantMockRecorder) GetRadios(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRadios", reflect.TypeOf((*MockWrapperRestaurant)(nil).GetRadios), arg0)
}

// GetRestaurants mocks base method.
func (m *MockWrapperRestaurant) GetRestaurants() ([]Utils.Restaurants, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRestaurants")
	ret0, _ := ret[0].([]Utils.Restaurants)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRestaurants indicates an expected call of GetRestaurants.
func (mr *MockWrapperRestaurantMockRecorder) GetRestaurants() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRestaurants", reflect.TypeOf((*MockWrapperRestaurant)(nil).GetRestaurants))
}

// GetStructDishes mocks base method.
func (m *MockWrapperRestaurant) GetStructDishes(arg0 int) ([]Utils.Ingredients, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStructDishes", arg0)
	ret0, _ := ret[0].([]Utils.Ingredients)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStructDishes indicates an expected call of GetStructDishes.
func (mr *MockWrapperRestaurantMockRecorder) GetStructDishes(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStructDishes", reflect.TypeOf((*MockWrapperRestaurant)(nil).GetStructDishes), arg0)
}

// GetTagsRestaurant mocks base method.
func (m *MockWrapperRestaurant) GetTagsRestaurant(arg0 int) ([]Utils.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTagsRestaurant", arg0)
	ret0, _ := ret[0].([]Utils.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTagsRestaurant indicates an expected call of GetTagsRestaurant.
func (mr *MockWrapperRestaurantMockRecorder) GetTagsRestaurant(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTagsRestaurant", reflect.TypeOf((*MockWrapperRestaurant)(nil).GetTagsRestaurant), arg0)
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

// UpdateAddress mocks base method.
func (m *MockWrapperProfile) UpdateAddress(arg0 int, arg1 Utils.AddressCoordinates) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAddress", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAddress indicates an expected call of UpdateAddress.
func (mr *MockWrapperProfileMockRecorder) UpdateAddress(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAddress", reflect.TypeOf((*MockWrapperProfile)(nil).UpdateAddress), arg0, arg1)
}

// UpdateAvatar mocks base method.
func (m *MockWrapperProfile) UpdateAvatar(arg0 int, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAvatar", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAvatar indicates an expected call of UpdateAvatar.
func (mr *MockWrapperProfileMockRecorder) UpdateAvatar(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAvatar", reflect.TypeOf((*MockWrapperProfile)(nil).UpdateAvatar), arg0, arg1)
}

// UpdateBirthday mocks base method.
func (m *MockWrapperProfile) UpdateBirthday(arg0 int, arg1 time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBirthday", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateBirthday indicates an expected call of UpdateBirthday.
func (mr *MockWrapperProfileMockRecorder) UpdateBirthday(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBirthday", reflect.TypeOf((*MockWrapperProfile)(nil).UpdateBirthday), arg0, arg1)
}

// UpdateEmail mocks base method.
func (m *MockWrapperProfile) UpdateEmail(arg0 int, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateEmail", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateEmail indicates an expected call of UpdateEmail.
func (mr *MockWrapperProfileMockRecorder) UpdateEmail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateEmail", reflect.TypeOf((*MockWrapperProfile)(nil).UpdateEmail), arg0, arg1)
}

// UpdateName mocks base method.
func (m *MockWrapperProfile) UpdateName(arg0 int, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateName", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateName indicates an expected call of UpdateName.
func (mr *MockWrapperProfileMockRecorder) UpdateName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateName", reflect.TypeOf((*MockWrapperProfile)(nil).UpdateName), arg0, arg1)
}

// UpdatePassword mocks base method.
func (m *MockWrapperProfile) UpdatePassword(arg0 int, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePassword", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePassword indicates an expected call of UpdatePassword.
func (mr *MockWrapperProfileMockRecorder) UpdatePassword(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePassword", reflect.TypeOf((*MockWrapperProfile)(nil).UpdatePassword), arg0, arg1)
}

// UpdatePhone mocks base method.
func (m *MockWrapperProfile) UpdatePhone(arg0 int, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePhone", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePhone indicates an expected call of UpdatePhone.
func (mr *MockWrapperProfileMockRecorder) UpdatePhone(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePhone", reflect.TypeOf((*MockWrapperProfile)(nil).UpdatePhone), arg0, arg1)
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
func (m *MockWrapperAuthorization) SignupClient(arg0 *Utils.RegistrationRequest, arg1 *Utils.Defense) (*Utils.Defense, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignupClient", arg0, arg1)
	ret0, _ := ret[0].(*Utils.Defense)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignupClient indicates an expected call of SignupClient.
func (mr *MockWrapperAuthorizationMockRecorder) SignupClient(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignupClient", reflect.TypeOf((*MockWrapperAuthorization)(nil).SignupClient), arg0, arg1)
}

// SignupCourier mocks base method.
func (m *MockWrapperAuthorization) SignupCourier(arg0 *Utils.RegistrationRequest, arg1 *Utils.Defense) (*Utils.Defense, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignupCourier", arg0, arg1)
	ret0, _ := ret[0].(*Utils.Defense)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignupCourier indicates an expected call of SignupCourier.
func (mr *MockWrapperAuthorizationMockRecorder) SignupCourier(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignupCourier", reflect.TypeOf((*MockWrapperAuthorization)(nil).SignupCourier), arg0, arg1)
}

// SignupHost mocks base method.
func (m *MockWrapperAuthorization) SignupHost(arg0 *Utils.RegistrationRequest, arg1 *Utils.Defense) (*Utils.Defense, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignupHost", arg0, arg1)
	ret0, _ := ret[0].(*Utils.Defense)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignupHost indicates an expected call of SignupHost.
func (mr *MockWrapperAuthorizationMockRecorder) SignupHost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignupHost", reflect.TypeOf((*MockWrapperAuthorization)(nil).SignupHost), arg0, arg1)
}

// MockWrapperCart is a mock of WrapperCart interface.
type MockWrapperCart struct {
	ctrl     *gomock.Controller
	recorder *MockWrapperCartMockRecorder
}

// MockWrapperCartMockRecorder is the mock recorder for MockWrapperCart.
type MockWrapperCartMockRecorder struct {
	mock *MockWrapperCart
}

// NewMockWrapperCart creates a new mock instance.
func NewMockWrapperCart(ctrl *gomock.Controller) *MockWrapperCart {
	mock := &MockWrapperCart{ctrl: ctrl}
	mock.recorder = &MockWrapperCartMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWrapperCart) EXPECT() *MockWrapperCartMockRecorder {
	return m.recorder
}

// DeleteCart mocks base method.
func (m *MockWrapperCart) DeleteCart(arg0 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCart", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCart indicates an expected call of DeleteCart.
func (mr *MockWrapperCartMockRecorder) DeleteCart(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCart", reflect.TypeOf((*MockWrapperCart)(nil).DeleteCart), arg0)
}

// GetCart mocks base method.
func (m *MockWrapperCart) GetCart(arg0 int) (*Utils.ResponseCartErrors, []Utils.CastDishesErrs, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCart", arg0)
	ret0, _ := ret[0].(*Utils.ResponseCartErrors)
	ret1, _ := ret[1].([]Utils.CastDishesErrs)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetCart indicates an expected call of GetCart.
func (mr *MockWrapperCartMockRecorder) GetCart(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCart", reflect.TypeOf((*MockWrapperCart)(nil).GetCart), arg0)
}

// GetConn mocks base method.
func (m *MockWrapperCart) GetConn() Utils.ConnectionInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConn")
	ret0, _ := ret[0].(Utils.ConnectionInterface)
	return ret0
}

// GetConn indicates an expected call of GetConn.
func (mr *MockWrapperCartMockRecorder) GetConn() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConn", reflect.TypeOf((*MockWrapperCart)(nil).GetConn))
}

// GetPriceDelivery mocks base method.
func (m *MockWrapperCart) GetPriceDelivery(arg0 int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPriceDelivery", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPriceDelivery indicates an expected call of GetPriceDelivery.
func (mr *MockWrapperCartMockRecorder) GetPriceDelivery(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPriceDelivery", reflect.TypeOf((*MockWrapperCart)(nil).GetPriceDelivery), arg0)
}

// GetStructFood mocks base method.
func (m *MockWrapperCart) GetStructFood(arg0 int) ([]Utils.IngredientCartResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStructFood", arg0)
	ret0, _ := ret[0].([]Utils.IngredientCartResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStructFood indicates an expected call of GetStructFood.
func (mr *MockWrapperCartMockRecorder) GetStructFood(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStructFood", reflect.TypeOf((*MockWrapperCart)(nil).GetStructFood), arg0)
}

// GetStructRadios mocks base method.
func (m *MockWrapperCart) GetStructRadios(arg0 int) ([]Utils.RadiosCartResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStructRadios", arg0)
	ret0, _ := ret[0].([]Utils.RadiosCartResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStructRadios indicates an expected call of GetStructRadios.
func (mr *MockWrapperCartMockRecorder) GetStructRadios(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStructRadios", reflect.TypeOf((*MockWrapperCart)(nil).GetStructRadios), arg0)
}

// UpdateCart mocks base method.
func (m *MockWrapperCart) UpdateCart(arg0 Utils.RequestCartDefault, arg1 int) (*Utils.ResponseCartErrors, []Utils.CastDishesErrs, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCart", arg0, arg1)
	ret0, _ := ret[0].(*Utils.ResponseCartErrors)
	ret1, _ := ret[1].([]Utils.CastDishesErrs)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// UpdateCart indicates an expected call of UpdateCart.
func (mr *MockWrapperCartMockRecorder) UpdateCart(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCart", reflect.TypeOf((*MockWrapperCart)(nil).UpdateCart), arg0, arg1)
}

// UpdateCartRadios mocks base method.
func (m *MockWrapperCart) UpdateCartRadios(arg0 []Utils.RadiosCartRequest, arg1 int, arg2 pgx.Tx) ([]Utils.RadiosCartResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCartRadios", arg0, arg1, arg2)
	ret0, _ := ret[0].([]Utils.RadiosCartResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCartRadios indicates an expected call of UpdateCartRadios.
func (mr *MockWrapperCartMockRecorder) UpdateCartRadios(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCartRadios", reflect.TypeOf((*MockWrapperCart)(nil).UpdateCartRadios), arg0, arg1, arg2)
}

// UpdateCartStructureFood mocks base method.
func (m *MockWrapperCart) UpdateCartStructureFood(arg0 []Utils.IngredientsCartRequest, arg1 int, arg2 pgx.Tx) ([]Utils.IngredientCartResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCartStructureFood", arg0, arg1, arg2)
	ret0, _ := ret[0].([]Utils.IngredientCartResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCartStructureFood indicates an expected call of UpdateCartStructureFood.
func (mr *MockWrapperCartMockRecorder) UpdateCartStructureFood(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCartStructureFood", reflect.TypeOf((*MockWrapperCart)(nil).UpdateCartStructureFood), arg0, arg1, arg2)
}
