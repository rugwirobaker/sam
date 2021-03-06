// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/rugwirobaker/helmes (interfaces: SendService,Pubsub)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	helmes "github.com/rugwirobaker/helmes"
	reflect "reflect"
)

// MockSendService is a mock of SendService interface
type MockSendService struct {
	ctrl     *gomock.Controller
	recorder *MockSendServiceMockRecorder
}

// MockSendServiceMockRecorder is the mock recorder for MockSendService
type MockSendServiceMockRecorder struct {
	mock *MockSendService
}

// NewMockSendService creates a new mock instance
func NewMockSendService(ctrl *gomock.Controller) *MockSendService {
	mock := &MockSendService{ctrl: ctrl}
	mock.recorder = &MockSendServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSendService) EXPECT() *MockSendServiceMockRecorder {
	return m.recorder
}

// Send mocks base method
func (m *MockSendService) Send(arg0 context.Context, arg1 *helmes.SMS) (*helmes.Report, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0, arg1)
	ret0, _ := ret[0].(*helmes.Report)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Send indicates an expected call of Send
func (mr *MockSendServiceMockRecorder) Send(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockSendService)(nil).Send), arg0, arg1)
}

// MockPubsub is a mock of Pubsub interface
type MockPubsub struct {
	ctrl     *gomock.Controller
	recorder *MockPubsubMockRecorder
}

// MockPubsubMockRecorder is the mock recorder for MockPubsub
type MockPubsubMockRecorder struct {
	mock *MockPubsub
}

// NewMockPubsub creates a new mock instance
func NewMockPubsub(ctrl *gomock.Controller) *MockPubsub {
	mock := &MockPubsub{ctrl: ctrl}
	mock.recorder = &MockPubsubMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPubsub) EXPECT() *MockPubsubMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockPubsub) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close
func (mr *MockPubsubMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockPubsub)(nil).Close))
}

// Done mocks base method
func (m *MockPubsub) Done(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Done", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Done indicates an expected call of Done
func (mr *MockPubsubMockRecorder) Done(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Done", reflect.TypeOf((*MockPubsub)(nil).Done), arg0, arg1)
}

// Publish mocks base method
func (m *MockPubsub) Publish(arg0 context.Context, arg1 helmes.Event) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Publish", arg0, arg1)
}

// Publish indicates an expected call of Publish
func (mr *MockPubsubMockRecorder) Publish(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Publish", reflect.TypeOf((*MockPubsub)(nil).Publish), arg0, arg1)
}

// Subscribe mocks base method
func (m *MockPubsub) Subscribe(arg0 context.Context, arg1 string) (<-chan helmes.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subscribe", arg0, arg1)
	ret0, _ := ret[0].(<-chan helmes.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Subscribe indicates an expected call of Subscribe
func (mr *MockPubsubMockRecorder) Subscribe(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockPubsub)(nil).Subscribe), arg0, arg1)
}
