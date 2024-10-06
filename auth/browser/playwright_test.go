// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/playwright-community/playwright-go (interfaces: Request)
//
// Generated by this command:
//
//	mockgen -package browser -destination playwright_test.go github.com/playwright-community/playwright-go Request
//

// Package browser is a generated GoMock package.
package browser

import (
	reflect "reflect"

	playwright "github.com/playwright-community/playwright-go"
	gomock "go.uber.org/mock/gomock"
)

// MockRequest is a mock of Request interface.
type MockRequest struct {
	ctrl     *gomock.Controller
	recorder *MockRequestMockRecorder
}

// MockRequestMockRecorder is the mock recorder for MockRequest.
type MockRequestMockRecorder struct {
	mock *MockRequest
}

// NewMockRequest creates a new mock instance.
func NewMockRequest(ctrl *gomock.Controller) *MockRequest {
	mock := &MockRequest{ctrl: ctrl}
	mock.recorder = &MockRequestMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRequest) EXPECT() *MockRequestMockRecorder {
	return m.recorder
}

// AllHeaders mocks base method.
func (m *MockRequest) AllHeaders() (map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllHeaders")
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllHeaders indicates an expected call of AllHeaders.
func (mr *MockRequestMockRecorder) AllHeaders() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllHeaders", reflect.TypeOf((*MockRequest)(nil).AllHeaders))
}

// Failure mocks base method.
func (m *MockRequest) Failure() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Failure")
	ret0, _ := ret[0].(error)
	return ret0
}

// Failure indicates an expected call of Failure.
func (mr *MockRequestMockRecorder) Failure() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Failure", reflect.TypeOf((*MockRequest)(nil).Failure))
}

// Frame mocks base method.
func (m *MockRequest) Frame() playwright.Frame {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Frame")
	ret0, _ := ret[0].(playwright.Frame)
	return ret0
}

// Frame indicates an expected call of Frame.
func (mr *MockRequestMockRecorder) Frame() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Frame", reflect.TypeOf((*MockRequest)(nil).Frame))
}

// HeaderValue mocks base method.
func (m *MockRequest) HeaderValue(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HeaderValue", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HeaderValue indicates an expected call of HeaderValue.
func (mr *MockRequestMockRecorder) HeaderValue(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HeaderValue", reflect.TypeOf((*MockRequest)(nil).HeaderValue), arg0)
}

// Headers mocks base method.
func (m *MockRequest) Headers() map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Headers")
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// Headers indicates an expected call of Headers.
func (mr *MockRequestMockRecorder) Headers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Headers", reflect.TypeOf((*MockRequest)(nil).Headers))
}

// HeadersArray mocks base method.
func (m *MockRequest) HeadersArray() ([]playwright.NameValue, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HeadersArray")
	ret0, _ := ret[0].([]playwright.NameValue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HeadersArray indicates an expected call of HeadersArray.
func (mr *MockRequestMockRecorder) HeadersArray() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HeadersArray", reflect.TypeOf((*MockRequest)(nil).HeadersArray))
}

// IsNavigationRequest mocks base method.
func (m *MockRequest) IsNavigationRequest() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsNavigationRequest")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsNavigationRequest indicates an expected call of IsNavigationRequest.
func (mr *MockRequestMockRecorder) IsNavigationRequest() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsNavigationRequest", reflect.TypeOf((*MockRequest)(nil).IsNavigationRequest))
}

// Method mocks base method.
func (m *MockRequest) Method() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Method")
	ret0, _ := ret[0].(string)
	return ret0
}

// Method indicates an expected call of Method.
func (mr *MockRequestMockRecorder) Method() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Method", reflect.TypeOf((*MockRequest)(nil).Method))
}

// PostData mocks base method.
func (m *MockRequest) PostData() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostData")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostData indicates an expected call of PostData.
func (mr *MockRequestMockRecorder) PostData() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostData", reflect.TypeOf((*MockRequest)(nil).PostData))
}

// PostDataBuffer mocks base method.
func (m *MockRequest) PostDataBuffer() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostDataBuffer")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostDataBuffer indicates an expected call of PostDataBuffer.
func (mr *MockRequestMockRecorder) PostDataBuffer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostDataBuffer", reflect.TypeOf((*MockRequest)(nil).PostDataBuffer))
}

// PostDataJSON mocks base method.
func (m *MockRequest) PostDataJSON(arg0 any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostDataJSON", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// PostDataJSON indicates an expected call of PostDataJSON.
func (mr *MockRequestMockRecorder) PostDataJSON(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostDataJSON", reflect.TypeOf((*MockRequest)(nil).PostDataJSON), arg0)
}

// RedirectedFrom mocks base method.
func (m *MockRequest) RedirectedFrom() playwright.Request {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RedirectedFrom")
	ret0, _ := ret[0].(playwright.Request)
	return ret0
}

// RedirectedFrom indicates an expected call of RedirectedFrom.
func (mr *MockRequestMockRecorder) RedirectedFrom() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RedirectedFrom", reflect.TypeOf((*MockRequest)(nil).RedirectedFrom))
}

// RedirectedTo mocks base method.
func (m *MockRequest) RedirectedTo() playwright.Request {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RedirectedTo")
	ret0, _ := ret[0].(playwright.Request)
	return ret0
}

// RedirectedTo indicates an expected call of RedirectedTo.
func (mr *MockRequestMockRecorder) RedirectedTo() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RedirectedTo", reflect.TypeOf((*MockRequest)(nil).RedirectedTo))
}

// ResourceType mocks base method.
func (m *MockRequest) ResourceType() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResourceType")
	ret0, _ := ret[0].(string)
	return ret0
}

// ResourceType indicates an expected call of ResourceType.
func (mr *MockRequestMockRecorder) ResourceType() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResourceType", reflect.TypeOf((*MockRequest)(nil).ResourceType))
}

// Response mocks base method.
func (m *MockRequest) Response() (playwright.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Response")
	ret0, _ := ret[0].(playwright.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Response indicates an expected call of Response.
func (mr *MockRequestMockRecorder) Response() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Response", reflect.TypeOf((*MockRequest)(nil).Response))
}

// Sizes mocks base method.
func (m *MockRequest) Sizes() (*playwright.RequestSizesResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sizes")
	ret0, _ := ret[0].(*playwright.RequestSizesResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Sizes indicates an expected call of Sizes.
func (mr *MockRequestMockRecorder) Sizes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sizes", reflect.TypeOf((*MockRequest)(nil).Sizes))
}

// Timing mocks base method.
func (m *MockRequest) Timing() *playwright.RequestTiming {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Timing")
	ret0, _ := ret[0].(*playwright.RequestTiming)
	return ret0
}

// Timing indicates an expected call of Timing.
func (mr *MockRequestMockRecorder) Timing() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Timing", reflect.TypeOf((*MockRequest)(nil).Timing))
}

// URL mocks base method.
func (m *MockRequest) URL() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "URL")
	ret0, _ := ret[0].(string)
	return ret0
}

// URL indicates an expected call of URL.
func (mr *MockRequestMockRecorder) URL() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "URL", reflect.TypeOf((*MockRequest)(nil).URL))
}
