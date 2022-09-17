// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/PereRohit/test-service/internal/handler (interfaces: TestServiceHandler)

// Package mock is a generated GoMock package.
package mock

import (
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockTestServiceHandler is a mock of TestServiceHandler interface.
type MockTestServiceHandler struct {
	ctrl     *gomock.Controller
	recorder *MockTestServiceHandlerMockRecorder
}

// MockTestServiceHandlerMockRecorder is the mock recorder for MockTestServiceHandler.
type MockTestServiceHandlerMockRecorder struct {
	mock *MockTestServiceHandler
}

// NewMockTestServiceHandler creates a new mock instance.
func NewMockTestServiceHandler(ctrl *gomock.Controller) *MockTestServiceHandler {
	mock := &MockTestServiceHandler{ctrl: ctrl}
	mock.recorder = &MockTestServiceHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTestServiceHandler) EXPECT() *MockTestServiceHandlerMockRecorder {
	return m.recorder
}

// HealthCheck mocks base method.
func (m *MockTestServiceHandler) HealthCheck() (string, string, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HealthCheck")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(bool)
	return ret0, ret1, ret2
}

// HealthCheck indicates an expected call of HealthCheck.
func (mr *MockTestServiceHandlerMockRecorder) HealthCheck() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HealthCheck", reflect.TypeOf((*MockTestServiceHandler)(nil).HealthCheck))
}

// Ping mocks base method.
func (m *MockTestServiceHandler) Ping(arg0 http.ResponseWriter, arg1 *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Ping", arg0, arg1)
}

// Ping indicates an expected call of Ping.
func (mr *MockTestServiceHandlerMockRecorder) Ping(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockTestServiceHandler)(nil).Ping), arg0, arg1)
}
