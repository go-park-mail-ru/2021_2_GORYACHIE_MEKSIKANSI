// Code generated by MockGen. DO NOT EDIT.
// Source: 2021_2_GORYACHIE_MEKSIKANSI/internals/profile/orm (interfaces: WrapperProfileInterface)

// Package mocks is a generated GoMock package.
package mocks

import (
	profile "2021_2_GORYACHIE_MEKSIKANSI/internals/profile"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockWrapperProfileInterface is a mock of WrapperProfileInterface interface.
type MockWrapperProfileInterface struct {
	ctrl     *gomock.Controller
	recorder *MockWrapperProfileInterfaceMockRecorder
}

// MockWrapperProfileInterfaceMockRecorder is the mock recorder for MockWrapperProfileInterface.
type MockWrapperProfileInterfaceMockRecorder struct {
	mock *MockWrapperProfileInterface
}

// NewMockWrapperProfileInterface creates a new mock instance.
func NewMockWrapperProfileInterface(ctrl *gomock.Controller) *MockWrapperProfileInterface {
	mock := &MockWrapperProfileInterface{ctrl: ctrl}
	mock.recorder = &MockWrapperProfileInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWrapperProfileInterface) EXPECT() *MockWrapperProfileInterfaceMockRecorder {
	return m.recorder
}

// AddAddress mocks base method.
func (m *MockWrapperProfileInterface) AddAddress(arg0 int, arg1 profile.AddressCoordinates) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAddress", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddAddress indicates an expected call of AddAddress.
func (mr *MockWrapperProfileInterfaceMockRecorder) AddAddress(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAddress", reflect.TypeOf((*MockWrapperProfileInterface)(nil).AddAddress), arg0, arg1)
}

// DeleteAddress mocks base method.
func (m *MockWrapperProfileInterface) DeleteAddress(arg0, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAddress", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAddress indicates an expected call of DeleteAddress.
func (mr *MockWrapperProfileInterfaceMockRecorder) DeleteAddress(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAddress", reflect.TypeOf((*MockWrapperProfileInterface)(nil).DeleteAddress), arg0, arg1)
}

// GetProfileClient mocks base method.
func (m *MockWrapperProfileInterface) GetProfileClient(arg0 int) (*profile.Profile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProfileClient", arg0)
	ret0, _ := ret[0].(*profile.Profile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProfileClient indicates an expected call of GetProfileClient.
func (mr *MockWrapperProfileInterfaceMockRecorder) GetProfileClient(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProfileClient", reflect.TypeOf((*MockWrapperProfileInterface)(nil).GetProfileClient), arg0)
}

// GetProfileCourier mocks base method.
func (m *MockWrapperProfileInterface) GetProfileCourier(arg0 int) (*profile.Profile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProfileCourier", arg0)
	ret0, _ := ret[0].(*profile.Profile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProfileCourier indicates an expected call of GetProfileCourier.
func (mr *MockWrapperProfileInterfaceMockRecorder) GetProfileCourier(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProfileCourier", reflect.TypeOf((*MockWrapperProfileInterface)(nil).GetProfileCourier), arg0)
}

// GetProfileHost mocks base method.
func (m *MockWrapperProfileInterface) GetProfileHost(arg0 int) (*profile.Profile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProfileHost", arg0)
	ret0, _ := ret[0].(*profile.Profile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProfileHost indicates an expected call of GetProfileHost.
func (mr *MockWrapperProfileInterfaceMockRecorder) GetProfileHost(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProfileHost", reflect.TypeOf((*MockWrapperProfileInterface)(nil).GetProfileHost), arg0)
}

// GetRoleById mocks base method.
func (m *MockWrapperProfileInterface) GetRoleById(arg0 int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRoleById", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRoleById indicates an expected call of GetRoleById.
func (mr *MockWrapperProfileInterfaceMockRecorder) GetRoleById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoleById", reflect.TypeOf((*MockWrapperProfileInterface)(nil).GetRoleById), arg0)
}

// UpdateAddress mocks base method.
func (m *MockWrapperProfileInterface) UpdateAddress(arg0 int, arg1 profile.AddressCoordinates) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAddress", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAddress indicates an expected call of UpdateAddress.
func (mr *MockWrapperProfileInterfaceMockRecorder) UpdateAddress(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAddress", reflect.TypeOf((*MockWrapperProfileInterface)(nil).UpdateAddress), arg0, arg1)
}

// UpdateAvatar mocks base method.
func (m *MockWrapperProfileInterface) UpdateAvatar(arg0 int, arg1 *profile.UpdateAvatar, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAvatar", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAvatar indicates an expected call of UpdateAvatar.
func (mr *MockWrapperProfileInterfaceMockRecorder) UpdateAvatar(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAvatar", reflect.TypeOf((*MockWrapperProfileInterface)(nil).UpdateAvatar), arg0, arg1, arg2)
}

// UpdateBirthday mocks base method.
func (m *MockWrapperProfileInterface) UpdateBirthday(arg0 int, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBirthday", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateBirthday indicates an expected call of UpdateBirthday.
func (mr *MockWrapperProfileInterfaceMockRecorder) UpdateBirthday(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBirthday", reflect.TypeOf((*MockWrapperProfileInterface)(nil).UpdateBirthday), arg0, arg1)
}

// UpdateEmail mocks base method.
func (m *MockWrapperProfileInterface) UpdateEmail(arg0 int, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateEmail", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateEmail indicates an expected call of UpdateEmail.
func (mr *MockWrapperProfileInterfaceMockRecorder) UpdateEmail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateEmail", reflect.TypeOf((*MockWrapperProfileInterface)(nil).UpdateEmail), arg0, arg1)
}

// UpdateName mocks base method.
func (m *MockWrapperProfileInterface) UpdateName(arg0 int, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateName", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateName indicates an expected call of UpdateName.
func (mr *MockWrapperProfileInterfaceMockRecorder) UpdateName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateName", reflect.TypeOf((*MockWrapperProfileInterface)(nil).UpdateName), arg0, arg1)
}

// UpdatePassword mocks base method.
func (m *MockWrapperProfileInterface) UpdatePassword(arg0 int, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePassword", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePassword indicates an expected call of UpdatePassword.
func (mr *MockWrapperProfileInterfaceMockRecorder) UpdatePassword(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePassword", reflect.TypeOf((*MockWrapperProfileInterface)(nil).UpdatePassword), arg0, arg1)
}

// UpdatePhone mocks base method.
func (m *MockWrapperProfileInterface) UpdatePhone(arg0 int, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePhone", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePhone indicates an expected call of UpdatePhone.
func (mr *MockWrapperProfileInterfaceMockRecorder) UpdatePhone(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePhone", reflect.TypeOf((*MockWrapperProfileInterface)(nil).UpdatePhone), arg0, arg1)
}
