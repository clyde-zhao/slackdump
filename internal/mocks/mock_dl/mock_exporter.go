// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/rusq/slackdump/v3/internal/structures/files/dl (interfaces: Exporter)
//
// Generated by this command:
//
//	mockgen -destination ../../../../internal/mocks/mock_dl/mock_exporter.go github.com/rusq/slackdump/v3/internal/structures/files/dl Exporter
//
// Package mock_dl is a generated GoMock package.
package mock_dl

import (
	context "context"
	reflect "reflect"

	slackdump "github.com/rusq/slackdump/v3"
	gomock "go.uber.org/mock/gomock"
)

// MockExporter is a mock of Exporter interface.
type MockExporter struct {
	ctrl     *gomock.Controller
	recorder *MockExporterMockRecorder
}

// MockExporterMockRecorder is the mock recorder for MockExporter.
type MockExporterMockRecorder struct {
	mock *MockExporter
}

// NewMockExporter creates a new mock instance.
func NewMockExporter(ctrl *gomock.Controller) *MockExporter {
	mock := &MockExporter{ctrl: ctrl}
	mock.recorder = &MockExporterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExporter) EXPECT() *MockExporterMockRecorder {
	return m.recorder
}

// ProcessFunc mocks base method.
func (m *MockExporter) ProcessFunc(arg0 string) slackdump.ProcessFunc {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessFunc", arg0)
	ret0, _ := ret[0].(slackdump.ProcessFunc)
	return ret0
}

// ProcessFunc indicates an expected call of ProcessFunc.
func (mr *MockExporterMockRecorder) ProcessFunc(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessFunc", reflect.TypeOf((*MockExporter)(nil).ProcessFunc), arg0)
}

// Start mocks base method.
func (m *MockExporter) Start(arg0 context.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Start", arg0)
}

// Start indicates an expected call of Start.
func (mr *MockExporterMockRecorder) Start(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockExporter)(nil).Start), arg0)
}

// Stop mocks base method.
func (m *MockExporter) Stop() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop.
func (mr *MockExporterMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockExporter)(nil).Stop))
}
