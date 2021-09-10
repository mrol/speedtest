// Code generated by MockGen. DO NOT EDIT.
// Source: deps.go

// Package ookla is a generated GoMock package.
package ookla

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockserverFinder is a mock of serverFinder interface.
type MockserverFinder struct {
	ctrl     *gomock.Controller
	recorder *MockserverFinderMockRecorder
}

// MockserverFinderMockRecorder is the mock recorder for MockserverFinder.
type MockserverFinderMockRecorder struct {
	mock *MockserverFinder
}

// NewMockserverFinder creates a new mock instance.
func NewMockserverFinder(ctrl *gomock.Controller) *MockserverFinder {
	mock := &MockserverFinder{ctrl: ctrl}
	mock.recorder = &MockserverFinderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockserverFinder) EXPECT() *MockserverFinderMockRecorder {
	return m.recorder
}

// FindServers mocks base method.
func (m *MockserverFinder) FindServers() ([]server, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindServers")
	ret0, _ := ret[0].([]server)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindServers indicates an expected call of FindServers.
func (mr *MockserverFinderMockRecorder) FindServers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindServers", reflect.TypeOf((*MockserverFinder)(nil).FindServers))
}

// Mockserver is a mock of server interface.
type Mockserver struct {
	ctrl     *gomock.Controller
	recorder *MockserverMockRecorder
}

// MockserverMockRecorder is the mock recorder for Mockserver.
type MockserverMockRecorder struct {
	mock *Mockserver
}

// NewMockserver creates a new mock instance.
func NewMockserver(ctrl *gomock.Controller) *Mockserver {
	mock := &Mockserver{ctrl: ctrl}
	mock.recorder = &MockserverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockserver) EXPECT() *MockserverMockRecorder {
	return m.recorder
}

// DownloadTest mocks base method.
func (m *Mockserver) DownloadTest() (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DownloadTest")
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DownloadTest indicates an expected call of DownloadTest.
func (mr *MockserverMockRecorder) DownloadTest() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DownloadTest", reflect.TypeOf((*Mockserver)(nil).DownloadTest))
}

// UploadTest mocks base method.
func (m *Mockserver) UploadTest() (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadTest")
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UploadTest indicates an expected call of UploadTest.
func (mr *MockserverMockRecorder) UploadTest() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadTest", reflect.TypeOf((*Mockserver)(nil).UploadTest))
}