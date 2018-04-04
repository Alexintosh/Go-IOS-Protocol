// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/iost-official/PrototypeWorks/protocol (interfaces: component)

// Package protocol_mock is a generated GoMock package.
package protocol_mock

import (
	gomock "github.com/golang/mock/gomock"
	iosbase "github.com/iost-official/PrototypeWorks/iosbase"
	protocol "github.com/iost-official/PrototypeWorks/protocol"
	reflect "reflect"
)

// MockComponent is a mock of component interface
type MockComponent struct {
	ctrl     *gomock.Controller
	recorder *MockComponentMockRecorder
}

// MockComponentMockRecorder is the mock recorder for MockComponent
type MockComponentMockRecorder struct {
	mock *MockComponent
}

// NewMockComponent creates a new mock instance
func NewMockComponent(ctrl *gomock.Controller) *MockComponent {
	mock := &MockComponent{ctrl: ctrl}
	mock.recorder = &MockComponentMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockComponent) EXPECT() *MockComponentMockRecorder {
	return m.recorder
}

// Init mocks base method
func (m *MockComponent) Init(arg0 iosbase.Member, arg1 protocol.Database, arg2 protocol.Router) error {
	ret := m.ctrl.Call(m, "Init", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init
func (mr *MockComponentMockRecorder) Init(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockComponent)(nil).Init), arg0, arg1, arg2)
}

// Run mocks base method
func (m *MockComponent) Run() {
	m.ctrl.Call(m, "Run")
}

// Run indicates an expected call of Run
func (mr *MockComponentMockRecorder) Run() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockComponent)(nil).Run))
}

// Stop mocks base method
func (m *MockComponent) Stop() {
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop
func (mr *MockComponentMockRecorder) Stop() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockComponent)(nil).Stop))
}
