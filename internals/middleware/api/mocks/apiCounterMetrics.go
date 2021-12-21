// Code generated by MockGen. DO NOT EDIT.
// Source: 2021_2_GORYACHIE_MEKSIKANSI/internals/middleware/api (interfaces: CounterMetricInterface)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockCounterMetricInterface is a mock of CounterMetricInterface interface.
type MockCounterMetricInterface struct {
	ctrl     *gomock.Controller
	recorder *MockCounterMetricInterfaceMockRecorder
}

// MockCounterMetricInterfaceMockRecorder is the mock recorder for MockCounterMetricInterface.
type MockCounterMetricInterfaceMockRecorder struct {
	mock *MockCounterMetricInterface
}

// NewMockCounterMetricInterface creates a new mock instance.
func NewMockCounterMetricInterface(ctrl *gomock.Controller) *MockCounterMetricInterface {
	mock := &MockCounterMetricInterface{ctrl: ctrl}
	mock.recorder = &MockCounterMetricInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCounterMetricInterface) EXPECT() *MockCounterMetricInterfaceMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockCounterMetricInterface) Add(arg0 float64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Add", arg0)
}

// Add indicates an expected call of Add.
func (mr *MockCounterMetricInterfaceMockRecorder) Add(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockCounterMetricInterface)(nil).Add), arg0)
}
