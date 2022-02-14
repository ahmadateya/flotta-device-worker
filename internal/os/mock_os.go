// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/project-flotta/flotta-device-worker/internal/os (interfaces: osExecCommands)

// Package os is a generated GoMock package.
package os

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockOsExecCommands is a mock of osExecCommands interface.
type MockOsExecCommands struct {
	ctrl     *gomock.Controller
	recorder *MockOsExecCommandsMockRecorder
}

// MockOsExecCommandsMockRecorder is the mock recorder for MockOsExecCommands.
type MockOsExecCommandsMockRecorder struct {
	mock *MockOsExecCommands
}

// NewMockOsExecCommands creates a new mock instance.
func NewMockOsExecCommands(ctrl *gomock.Controller) *MockOsExecCommands {
	mock := &MockOsExecCommands{ctrl: ctrl}
	mock.recorder = &MockOsExecCommandsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOsExecCommands) EXPECT() *MockOsExecCommandsMockRecorder {
	return m.recorder
}

// EnsureScriptExists mocks base method.
func (m *MockOsExecCommands) EnsureScriptExists(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureScriptExists", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureScriptExists indicates an expected call of EnsureScriptExists.
func (mr *MockOsExecCommandsMockRecorder) EnsureScriptExists(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureScriptExists", reflect.TypeOf((*MockOsExecCommands)(nil).EnsureScriptExists), arg0, arg1)
}

// RpmOstreeStatus mocks base method.
func (m *MockOsExecCommands) RpmOstreeStatus() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RpmOstreeStatus")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RpmOstreeStatus indicates an expected call of RpmOstreeStatus.
func (mr *MockOsExecCommandsMockRecorder) RpmOstreeStatus() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RpmOstreeStatus", reflect.TypeOf((*MockOsExecCommands)(nil).RpmOstreeStatus))
}

// RpmOstreeUpdatePreview mocks base method.
func (m *MockOsExecCommands) RpmOstreeUpdatePreview() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RpmOstreeUpdatePreview")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RpmOstreeUpdatePreview indicates an expected call of RpmOstreeUpdatePreview.
func (mr *MockOsExecCommandsMockRecorder) RpmOstreeUpdatePreview() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RpmOstreeUpdatePreview", reflect.TypeOf((*MockOsExecCommands)(nil).RpmOstreeUpdatePreview))
}

// RpmOstreeUpgrade mocks base method.
func (m *MockOsExecCommands) RpmOstreeUpgrade() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RpmOstreeUpgrade")
	ret0, _ := ret[0].(error)
	return ret0
}

// RpmOstreeUpgrade indicates an expected call of RpmOstreeUpgrade.
func (mr *MockOsExecCommandsMockRecorder) RpmOstreeUpgrade() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RpmOstreeUpgrade", reflect.TypeOf((*MockOsExecCommands)(nil).RpmOstreeUpgrade))
}

// SystemReboot mocks base method.
func (m *MockOsExecCommands) SystemReboot() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SystemReboot")
	ret0, _ := ret[0].(error)
	return ret0
}

// SystemReboot indicates an expected call of SystemReboot.
func (mr *MockOsExecCommandsMockRecorder) SystemReboot() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SystemReboot", reflect.TypeOf((*MockOsExecCommands)(nil).SystemReboot))
}

// UpdateUrlInEdgeRemote mocks base method.
func (m *MockOsExecCommands) UpdateUrlInEdgeRemote(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUrlInEdgeRemote", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUrlInEdgeRemote indicates an expected call of UpdateUrlInEdgeRemote.
func (mr *MockOsExecCommandsMockRecorder) UpdateUrlInEdgeRemote(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUrlInEdgeRemote", reflect.TypeOf((*MockOsExecCommands)(nil).UpdateUrlInEdgeRemote), arg0, arg1)
}
