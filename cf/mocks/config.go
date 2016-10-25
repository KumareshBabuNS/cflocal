// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/sclevine/cflocal/cf (interfaces: Config)

package mocks

import (
	gomock "github.com/golang/mock/gomock"
	local "github.com/sclevine/cflocal/local"
)

// Mock of Config interface
type MockConfig struct {
	ctrl     *gomock.Controller
	recorder *_MockConfigRecorder
}

// Recorder for MockConfig (not exported)
type _MockConfigRecorder struct {
	mock *MockConfig
}

func NewMockConfig(ctrl *gomock.Controller) *MockConfig {
	mock := &MockConfig{ctrl: ctrl}
	mock.recorder = &_MockConfigRecorder{mock}
	return mock
}

func (_m *MockConfig) EXPECT() *_MockConfigRecorder {
	return _m.recorder
}

func (_m *MockConfig) Load() (*local.LocalYML, error) {
	ret := _m.ctrl.Call(_m, "Load")
	ret0, _ := ret[0].(*local.LocalYML)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockConfigRecorder) Load() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Load")
}

func (_m *MockConfig) Save(_param0 *local.LocalYML) error {
	ret := _m.ctrl.Call(_m, "Save", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockConfigRecorder) Save(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Save", arg0)
}
