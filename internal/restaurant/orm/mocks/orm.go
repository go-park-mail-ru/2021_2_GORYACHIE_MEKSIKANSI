// Code generated by MockGen. DO NOT EDIT.
// Source: 2021_2_GORYACHIE_MEKSIKANSI/internal/restaurant/orm (interfaces: WrapperRestaurantServerInterface,ConnectRestaurantServiceInterface)

// Package mocks is a generated GoMock package.
package mocks

import (
	resProto "2021_2_GORYACHIE_MEKSIKANSI/internal/microservice/restaurant/proto"
	restaurant "2021_2_GORYACHIE_MEKSIKANSI/internal/restaurant"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockWrapperRestaurantServerInterface is a mock of WrapperRestaurantServerInterface interface.
type MockWrapperRestaurantServerInterface struct {
	ctrl     *gomock.Controller
	recorder *MockWrapperRestaurantServerInterfaceMockRecorder
}

// MockWrapperRestaurantServerInterfaceMockRecorder is the mock recorder for MockWrapperRestaurantServerInterface.
type MockWrapperRestaurantServerInterfaceMockRecorder struct {
	mock *MockWrapperRestaurantServerInterface
}

// NewMockWrapperRestaurantServerInterface creates a new mock instance.
func NewMockWrapperRestaurantServerInterface(ctrl *gomock.Controller) *MockWrapperRestaurantServerInterface {
	mock := &MockWrapperRestaurantServerInterface{ctrl: ctrl}
	mock.recorder = &MockWrapperRestaurantServerInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWrapperRestaurantServerInterface) EXPECT() *MockWrapperRestaurantServerInterfaceMockRecorder {
	return m.recorder
}

// AllRestaurants mocks base method.
func (m *MockWrapperRestaurantServerInterface) AllRestaurants() (*restaurant.AllRestaurants, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllRestaurants")
	ret0, _ := ret[0].(*restaurant.AllRestaurants)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllRestaurants indicates an expected call of AllRestaurants.
func (mr *MockWrapperRestaurantServerInterfaceMockRecorder) AllRestaurants() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllRestaurants", reflect.TypeOf((*MockWrapperRestaurantServerInterface)(nil).AllRestaurants))
}

// CreateReview mocks base method.
func (m *MockWrapperRestaurantServerInterface) CreateReview(arg0 int, arg1 restaurant.NewReview) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateReview", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateReview indicates an expected call of CreateReview.
func (mr *MockWrapperRestaurantServerInterfaceMockRecorder) CreateReview(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateReview", reflect.TypeOf((*MockWrapperRestaurantServerInterface)(nil).CreateReview), arg0, arg1)
}

// EditRestaurantInFavorite mocks base method.
func (m *MockWrapperRestaurantServerInterface) EditRestaurantInFavorite(arg0, arg1 int) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EditRestaurantInFavorite", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EditRestaurantInFavorite indicates an expected call of EditRestaurantInFavorite.
func (mr *MockWrapperRestaurantServerInterfaceMockRecorder) EditRestaurantInFavorite(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditRestaurantInFavorite", reflect.TypeOf((*MockWrapperRestaurantServerInterface)(nil).EditRestaurantInFavorite), arg0, arg1)
}

// GetFavoriteRestaurants mocks base method.
func (m *MockWrapperRestaurantServerInterface) GetFavoriteRestaurants(arg0 int) ([]restaurant.Restaurants, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFavoriteRestaurants", arg0)
	ret0, _ := ret[0].([]restaurant.Restaurants)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFavoriteRestaurants indicates an expected call of GetFavoriteRestaurants.
func (mr *MockWrapperRestaurantServerInterfaceMockRecorder) GetFavoriteRestaurants(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFavoriteRestaurants", reflect.TypeOf((*MockWrapperRestaurantServerInterface)(nil).GetFavoriteRestaurants), arg0)
}

// GetRestaurant mocks base method.
func (m *MockWrapperRestaurantServerInterface) GetRestaurant(arg0, arg1 int) (*restaurant.RestaurantId, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRestaurant", arg0, arg1)
	ret0, _ := ret[0].(*restaurant.RestaurantId)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRestaurant indicates an expected call of GetRestaurant.
func (mr *MockWrapperRestaurantServerInterfaceMockRecorder) GetRestaurant(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRestaurant", reflect.TypeOf((*MockWrapperRestaurantServerInterface)(nil).GetRestaurant), arg0, arg1)
}

// GetReview mocks base method.
func (m *MockWrapperRestaurantServerInterface) GetReview(arg0, arg1 int) (*restaurant.ResReview, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReview", arg0, arg1)
	ret0, _ := ret[0].(*restaurant.ResReview)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReview indicates an expected call of GetReview.
func (mr *MockWrapperRestaurantServerInterfaceMockRecorder) GetReview(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReview", reflect.TypeOf((*MockWrapperRestaurantServerInterface)(nil).GetReview), arg0, arg1)
}

// RestaurantDishes mocks base method.
func (m *MockWrapperRestaurantServerInterface) RestaurantDishes(arg0, arg1 int) (*restaurant.Dishes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RestaurantDishes", arg0, arg1)
	ret0, _ := ret[0].(*restaurant.Dishes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RestaurantDishes indicates an expected call of RestaurantDishes.
func (mr *MockWrapperRestaurantServerInterfaceMockRecorder) RestaurantDishes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RestaurantDishes", reflect.TypeOf((*MockWrapperRestaurantServerInterface)(nil).RestaurantDishes), arg0, arg1)
}

// SearchRestaurant mocks base method.
func (m *MockWrapperRestaurantServerInterface) SearchRestaurant(arg0 string) ([]restaurant.Restaurants, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchRestaurant", arg0)
	ret0, _ := ret[0].([]restaurant.Restaurants)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchRestaurant indicates an expected call of SearchRestaurant.
func (mr *MockWrapperRestaurantServerInterfaceMockRecorder) SearchRestaurant(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchRestaurant", reflect.TypeOf((*MockWrapperRestaurantServerInterface)(nil).SearchRestaurant), arg0)
}

// MockConnectRestaurantServiceInterface is a mock of ConnectRestaurantServiceInterface interface.
type MockConnectRestaurantServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockConnectRestaurantServiceInterfaceMockRecorder
}

// MockConnectRestaurantServiceInterfaceMockRecorder is the mock recorder for MockConnectRestaurantServiceInterface.
type MockConnectRestaurantServiceInterfaceMockRecorder struct {
	mock *MockConnectRestaurantServiceInterface
}

// NewMockConnectRestaurantServiceInterface creates a new mock instance.
func NewMockConnectRestaurantServiceInterface(ctrl *gomock.Controller) *MockConnectRestaurantServiceInterface {
	mock := &MockConnectRestaurantServiceInterface{ctrl: ctrl}
	mock.recorder = &MockConnectRestaurantServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConnectRestaurantServiceInterface) EXPECT() *MockConnectRestaurantServiceInterfaceMockRecorder {
	return m.recorder
}

// AllRestaurants mocks base method.
func (m *MockConnectRestaurantServiceInterface) AllRestaurants(arg0 context.Context, arg1 *resProto.Empty, arg2 ...grpc.CallOption) (*resProto.RestaurantsTagsPromo, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AllRestaurants", varargs...)
	ret0, _ := ret[0].(*resProto.RestaurantsTagsPromo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllRestaurants indicates an expected call of AllRestaurants.
func (mr *MockConnectRestaurantServiceInterfaceMockRecorder) AllRestaurants(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllRestaurants", reflect.TypeOf((*MockConnectRestaurantServiceInterface)(nil).AllRestaurants), varargs...)
}

// CreateReview mocks base method.
func (m *MockConnectRestaurantServiceInterface) CreateReview(arg0 context.Context, arg1 *resProto.NewReview, arg2 ...grpc.CallOption) (*resProto.Error, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateReview", varargs...)
	ret0, _ := ret[0].(*resProto.Error)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateReview indicates an expected call of CreateReview.
func (mr *MockConnectRestaurantServiceInterfaceMockRecorder) CreateReview(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateReview", reflect.TypeOf((*MockConnectRestaurantServiceInterface)(nil).CreateReview), varargs...)
}

// EditRestaurantInFavorite mocks base method.
func (m *MockConnectRestaurantServiceInterface) EditRestaurantInFavorite(arg0 context.Context, arg1 *resProto.EditRestaurantInFavoriteRequest, arg2 ...grpc.CallOption) (*resProto.ResponseEditRestaurantInFavorite, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "EditRestaurantInFavorite", varargs...)
	ret0, _ := ret[0].(*resProto.ResponseEditRestaurantInFavorite)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EditRestaurantInFavorite indicates an expected call of EditRestaurantInFavorite.
func (mr *MockConnectRestaurantServiceInterfaceMockRecorder) EditRestaurantInFavorite(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditRestaurantInFavorite", reflect.TypeOf((*MockConnectRestaurantServiceInterface)(nil).EditRestaurantInFavorite), varargs...)
}

// GetFavoriteRestaurants mocks base method.
func (m *MockConnectRestaurantServiceInterface) GetFavoriteRestaurants(arg0 context.Context, arg1 *resProto.UserId, arg2 ...grpc.CallOption) (*resProto.Restaurants, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetFavoriteRestaurants", varargs...)
	ret0, _ := ret[0].(*resProto.Restaurants)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFavoriteRestaurants indicates an expected call of GetFavoriteRestaurants.
func (mr *MockConnectRestaurantServiceInterfaceMockRecorder) GetFavoriteRestaurants(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFavoriteRestaurants", reflect.TypeOf((*MockConnectRestaurantServiceInterface)(nil).GetFavoriteRestaurants), varargs...)
}

// GetRestaurant mocks base method.
func (m *MockConnectRestaurantServiceInterface) GetRestaurant(arg0 context.Context, arg1 *resProto.RestaurantId, arg2 ...grpc.CallOption) (*resProto.RestaurantInfo, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRestaurant", varargs...)
	ret0, _ := ret[0].(*resProto.RestaurantInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRestaurant indicates an expected call of GetRestaurant.
func (mr *MockConnectRestaurantServiceInterfaceMockRecorder) GetRestaurant(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRestaurant", reflect.TypeOf((*MockConnectRestaurantServiceInterface)(nil).GetRestaurant), varargs...)
}

// GetReview mocks base method.
func (m *MockConnectRestaurantServiceInterface) GetReview(arg0 context.Context, arg1 *resProto.RestaurantClientId, arg2 ...grpc.CallOption) (*resProto.ResReview, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetReview", varargs...)
	ret0, _ := ret[0].(*resProto.ResReview)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReview indicates an expected call of GetReview.
func (mr *MockConnectRestaurantServiceInterfaceMockRecorder) GetReview(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReview", reflect.TypeOf((*MockConnectRestaurantServiceInterface)(nil).GetReview), varargs...)
}

// RestaurantDishes mocks base method.
func (m *MockConnectRestaurantServiceInterface) RestaurantDishes(arg0 context.Context, arg1 *resProto.DishInfo, arg2 ...grpc.CallOption) (*resProto.Dishes, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RestaurantDishes", varargs...)
	ret0, _ := ret[0].(*resProto.Dishes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RestaurantDishes indicates an expected call of RestaurantDishes.
func (mr *MockConnectRestaurantServiceInterfaceMockRecorder) RestaurantDishes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RestaurantDishes", reflect.TypeOf((*MockConnectRestaurantServiceInterface)(nil).RestaurantDishes), varargs...)
}

// SearchRestaurant mocks base method.
func (m *MockConnectRestaurantServiceInterface) SearchRestaurant(arg0 context.Context, arg1 *resProto.SearchRestaurantText, arg2 ...grpc.CallOption) (*resProto.Restaurants, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SearchRestaurant", varargs...)
	ret0, _ := ret[0].(*resProto.Restaurants)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchRestaurant indicates an expected call of SearchRestaurant.
func (mr *MockConnectRestaurantServiceInterfaceMockRecorder) SearchRestaurant(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchRestaurant", reflect.TypeOf((*MockConnectRestaurantServiceInterface)(nil).SearchRestaurant), varargs...)
}
