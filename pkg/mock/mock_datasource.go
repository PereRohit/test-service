// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/PereRohit/test-service/internal/repo/datasource (interfaces: DataSource)

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	model "github.com/PereRohit/test-service/internal/model"
	gomock "github.com/golang/mock/gomock"
)

// MockDataSource is a mock of DataSource interface.
type MockDataSource struct {
	ctrl     *gomock.Controller
	recorder *MockDataSourceMockRecorder
}

// MockDataSourceMockRecorder is the mock recorder for MockDataSource.
type MockDataSourceMockRecorder struct {
	mock *MockDataSource
}

// NewMockDataSource creates a new mock instance.
func NewMockDataSource(ctrl *gomock.Controller) *MockDataSource {
	mock := &MockDataSource{ctrl: ctrl}
	mock.recorder = &MockDataSourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDataSource) EXPECT() *MockDataSourceMockRecorder {
	return m.recorder
}

// HealthCheck mocks base method.
func (m *MockDataSource) HealthCheck() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HealthCheck")
	ret0, _ := ret[0].(bool)
	return ret0
}

// HealthCheck indicates an expected call of HealthCheck.
func (mr *MockDataSourceMockRecorder) HealthCheck() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HealthCheck", reflect.TypeOf((*MockDataSource)(nil).HealthCheck))
}

// Ping mocks base method.
func (m *MockDataSource) Ping(arg0 *model.PingDs) (*model.DsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ping", arg0)
	ret0, _ := ret[0].(*model.DsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Ping indicates an expected call of Ping.
func (mr *MockDataSourceMockRecorder) Ping(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockDataSource)(nil).Ping), arg0)
}
