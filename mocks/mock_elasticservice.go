// Code generated by MockGen. DO NOT EDIT.
// Source: elastic/elasticservice.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIElasticService is a mock of IElasticService interface.
type MockIElasticService struct {
	ctrl     *gomock.Controller
	recorder *MockIElasticServiceMockRecorder
}

// MockIElasticServiceMockRecorder is the mock recorder for MockIElasticService.
type MockIElasticServiceMockRecorder struct {
	mock *MockIElasticService
}

// NewMockIElasticService creates a new mock instance.
func NewMockIElasticService(ctrl *gomock.Controller) *MockIElasticService {
	mock := &MockIElasticService{ctrl: ctrl}
	mock.recorder = &MockIElasticServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIElasticService) EXPECT() *MockIElasticServiceMockRecorder {
	return m.recorder
}

// ESscroll mocks base method.
func (m *MockIElasticService) ESscroll(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ESscroll", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// ESscroll indicates an expected call of ESscroll.
func (mr *MockIElasticServiceMockRecorder) ESscroll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ESscroll", reflect.TypeOf((*MockIElasticService)(nil).ESscroll), arg0)
}

// Method2 mocks base method.
func (m *MockIElasticService) Method2(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Method2", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// Method2 indicates an expected call of Method2.
func (mr *MockIElasticServiceMockRecorder) Method2(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Method2", reflect.TypeOf((*MockIElasticService)(nil).Method2), arg0)
}
