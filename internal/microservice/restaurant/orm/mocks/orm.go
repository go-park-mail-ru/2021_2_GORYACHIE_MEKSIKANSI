// Code generated by MockGen. DO NOT EDIT.
// Source: 2021_2_GORYACHIE_MEKSIKANSI/internal/microservice/restaurant/orm (interfaces: WrapperRestaurantInterface,ConnectionInterface)

// Package mocks is a generated GoMock package.
package mocks

import (
	restaurantlf "2021_2_GORYACHIE_MEKSIKANSI/internal/microservice/restaurant"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	pgconn "github.com/jackc/pgconn"
	pgx "github.com/jackc/pgx/v4"
)

// MockWrapperRestaurantInterface is a mock of WrapperRestaurantInterface interface.
type MockWrapperRestaurantInterface struct {
	ctrl     *gomock.Controller
	recorder *MockWrapperRestaurantInterfaceMockRecorder
}

// MockWrapperRestaurantInterfaceMockRecorder is the mock recorder for MockWrapperRestaurantInterface.
type MockWrapperRestaurantInterfaceMockRecorder struct {
	mock *MockWrapperRestaurantInterface
}

// NewMockWrapperRestaurantInterface creates a new mock instance.
func NewMockWrapperRestaurantInterface(ctrl *gomock.Controller) *MockWrapperRestaurantInterface {
	mock := &MockWrapperRestaurantInterface{ctrl: ctrl}
	mock.recorder = &MockWrapperRestaurantInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWrapperRestaurantInterface) EXPECT() *MockWrapperRestaurantInterfaceMockRecorder {
	return m.recorder
}

// CreateReview mocks base method.
func (m *MockWrapperRestaurantInterface) CreateReview(arg0 int, arg1 restaurantlf.NewReview) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateReview", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateReview indicates an expected call of CreateReview.
func (mr *MockWrapperRestaurantInterfaceMockRecorder) CreateReview(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateReview", reflect.TypeOf((*MockWrapperRestaurantInterface)(nil).CreateReview), arg0, arg1)
}

// EditRestaurantInFavorite mocks base method.
func (m *MockWrapperRestaurantInterface) EditRestaurantInFavorite(arg0, arg1 int) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EditRestaurantInFavorite", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EditRestaurantInFavorite indicates an expected call of EditRestaurantInFavorite.
func (mr *MockWrapperRestaurantInterfaceMockRecorder) EditRestaurantInFavorite(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditRestaurantInFavorite", reflect.TypeOf((*MockWrapperRestaurantInterface)(nil).EditRestaurantInFavorite), arg0, arg1)
}

// GetDishes mocks base method.
func (m *MockWrapperRestaurantInterface) GetDishes(arg0, arg1 int) (*restaurantlf.Dishes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDishes", arg0, arg1)
	ret0, _ := ret[0].(*restaurantlf.Dishes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDishes indicates an expected call of GetDishes.
func (mr *MockWrapperRestaurantInterfaceMockRecorder) GetDishes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDishes", reflect.TypeOf((*MockWrapperRestaurantInterface)(nil).GetDishes), arg0, arg1)
}

// GetFavoriteRestaurants mocks base method.
func (m *MockWrapperRestaurantInterface) GetFavoriteRestaurants(arg0 int) ([]restaurantlf.Restaurants, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFavoriteRestaurants", arg0)
	ret0, _ := ret[0].([]restaurantlf.Restaurants)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFavoriteRestaurants indicates an expected call of GetFavoriteRestaurants.
func (mr *MockWrapperRestaurantInterfaceMockRecorder) GetFavoriteRestaurants(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFavoriteRestaurants", reflect.TypeOf((*MockWrapperRestaurantInterface)(nil).GetFavoriteRestaurants), arg0)
}

// GetGeneralInfoRestaurant mocks base method.
func (m *MockWrapperRestaurantInterface) GetGeneralInfoRestaurant(arg0 int) (*restaurantlf.Restaurants, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGeneralInfoRestaurant", arg0)
	ret0, _ := ret[0].(*restaurantlf.Restaurants)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGeneralInfoRestaurant indicates an expected call of GetGeneralInfoRestaurant.
func (mr *MockWrapperRestaurantInterfaceMockRecorder) GetGeneralInfoRestaurant(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGeneralInfoRestaurant", reflect.TypeOf((*MockWrapperRestaurantInterface)(nil).GetGeneralInfoRestaurant), arg0)
}

// GetMenu mocks base method.
func (m *MockWrapperRestaurantInterface) GetMenu(arg0 int) ([]restaurantlf.Menu, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMenu", arg0)
	ret0, _ := ret[0].([]restaurantlf.Menu)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMenu indicates an expected call of GetMenu.
func (mr *MockWrapperRestaurantInterfaceMockRecorder) GetMenu(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMenu", reflect.TypeOf((*MockWrapperRestaurantInterface)(nil).GetMenu), arg0)
}

// GetRestaurant mocks base method.
func (m *MockWrapperRestaurantInterface) GetRestaurant(arg0, arg1 int) (*restaurantlf.RestaurantId, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRestaurant", arg0, arg1)
	ret0, _ := ret[0].(*restaurantlf.RestaurantId)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRestaurant indicates an expected call of GetRestaurant.
func (mr *MockWrapperRestaurantInterfaceMockRecorder) GetRestaurant(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRestaurant", reflect.TypeOf((*MockWrapperRestaurantInterface)(nil).GetRestaurant), arg0, arg1)
}

// GetRestaurants mocks base method.
func (m *MockWrapperRestaurantInterface) GetRestaurants() (*restaurantlf.AllRestaurants, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRestaurants")
	ret0, _ := ret[0].(*restaurantlf.AllRestaurants)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRestaurants indicates an expected call of GetRestaurants.
func (mr *MockWrapperRestaurantInterfaceMockRecorder) GetRestaurants() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRestaurants", reflect.TypeOf((*MockWrapperRestaurantInterface)(nil).GetRestaurants))
}

// GetReview mocks base method.
func (m *MockWrapperRestaurantInterface) GetReview(arg0 int) ([]restaurantlf.Review, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReview", arg0)
	ret0, _ := ret[0].([]restaurantlf.Review)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReview indicates an expected call of GetReview.
func (mr *MockWrapperRestaurantInterfaceMockRecorder) GetReview(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReview", reflect.TypeOf((*MockWrapperRestaurantInterface)(nil).GetReview), arg0)
}

// GetTagsRestaurant mocks base method.
func (m *MockWrapperRestaurantInterface) GetTagsRestaurant(arg0 int) ([]restaurantlf.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTagsRestaurant", arg0)
	ret0, _ := ret[0].([]restaurantlf.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTagsRestaurant indicates an expected call of GetTagsRestaurant.
func (mr *MockWrapperRestaurantInterfaceMockRecorder) GetTagsRestaurant(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTagsRestaurant", reflect.TypeOf((*MockWrapperRestaurantInterface)(nil).GetTagsRestaurant), arg0)
}

// IsFavoriteRestaurant mocks base method.
func (m *MockWrapperRestaurantInterface) IsFavoriteRestaurant(arg0, arg1 int) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsFavoriteRestaurant", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsFavoriteRestaurant indicates an expected call of IsFavoriteRestaurant.
func (mr *MockWrapperRestaurantInterfaceMockRecorder) IsFavoriteRestaurant(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsFavoriteRestaurant", reflect.TypeOf((*MockWrapperRestaurantInterface)(nil).IsFavoriteRestaurant), arg0, arg1)
}

// SearchCategory mocks base method.
func (m *MockWrapperRestaurantInterface) SearchCategory(arg0 string) ([]int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchCategory", arg0)
	ret0, _ := ret[0].([]int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchCategory indicates an expected call of SearchCategory.
func (mr *MockWrapperRestaurantInterfaceMockRecorder) SearchCategory(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchCategory", reflect.TypeOf((*MockWrapperRestaurantInterface)(nil).SearchCategory), arg0)
}

// SearchRestaurant mocks base method.
func (m *MockWrapperRestaurantInterface) SearchRestaurant(arg0 string) ([]int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchRestaurant", arg0)
	ret0, _ := ret[0].([]int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchRestaurant indicates an expected call of SearchRestaurant.
func (mr *MockWrapperRestaurantInterfaceMockRecorder) SearchRestaurant(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchRestaurant", reflect.TypeOf((*MockWrapperRestaurantInterface)(nil).SearchRestaurant), arg0)
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
