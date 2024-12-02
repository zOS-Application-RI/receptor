// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/services/udp_proxy.go
//
// Generated by this command:
//
//	mockgen -source=pkg/services/udp_proxy.go -destination=pkg/services/mock_services/udp_proxy.go
//

// Package mock_services is a generated GoMock package.
package mock_services

import (
	reflect "reflect"

	logger "github.com/ansible/receptor/pkg/logger"
	netceptor "github.com/ansible/receptor/pkg/netceptor"
	gomock "go.uber.org/mock/gomock"
)

// MockNetcForUDPProxy is a mock of NetcForUDPProxy interface.
type MockNetcForUDPProxy struct {
	ctrl     *gomock.Controller
	recorder *MockNetcForUDPProxyMockRecorder
	isgomock struct{}
}

// MockNetcForUDPProxyMockRecorder is the mock recorder for MockNetcForUDPProxy.
type MockNetcForUDPProxyMockRecorder struct {
	mock *MockNetcForUDPProxy
}

// NewMockNetcForUDPProxy creates a new mock instance.
func NewMockNetcForUDPProxy(ctrl *gomock.Controller) *MockNetcForUDPProxy {
	mock := &MockNetcForUDPProxy{ctrl: ctrl}
	mock.recorder = &MockNetcForUDPProxyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNetcForUDPProxy) EXPECT() *MockNetcForUDPProxyMockRecorder {
	return m.recorder
}

// GetLogger mocks base method.
func (m *MockNetcForUDPProxy) GetLogger() *logger.ReceptorLogger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLogger")
	ret0, _ := ret[0].(*logger.ReceptorLogger)
	return ret0
}

// GetLogger indicates an expected call of GetLogger.
func (mr *MockNetcForUDPProxyMockRecorder) GetLogger() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLogger", reflect.TypeOf((*MockNetcForUDPProxy)(nil).GetLogger))
}

// ListenPacket mocks base method.
func (m *MockNetcForUDPProxy) ListenPacket(service string) (netceptor.PacketConner, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListenPacket", service)
	ret0, _ := ret[0].(netceptor.PacketConner)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListenPacket indicates an expected call of ListenPacket.
func (mr *MockNetcForUDPProxyMockRecorder) ListenPacket(service any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListenPacket", reflect.TypeOf((*MockNetcForUDPProxy)(nil).ListenPacket), service)
}

// ListenPacketAndAdvertise mocks base method.
func (m *MockNetcForUDPProxy) ListenPacketAndAdvertise(service string, tags map[string]string) (netceptor.PacketConner, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListenPacketAndAdvertise", service, tags)
	ret0, _ := ret[0].(netceptor.PacketConner)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListenPacketAndAdvertise indicates an expected call of ListenPacketAndAdvertise.
func (mr *MockNetcForUDPProxyMockRecorder) ListenPacketAndAdvertise(service, tags any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListenPacketAndAdvertise", reflect.TypeOf((*MockNetcForUDPProxy)(nil).ListenPacketAndAdvertise), service, tags)
}

// NewAddr mocks base method.
func (m *MockNetcForUDPProxy) NewAddr(node, service string) netceptor.Addr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewAddr", node, service)
	ret0, _ := ret[0].(netceptor.Addr)
	return ret0
}

// NewAddr indicates an expected call of NewAddr.
func (mr *MockNetcForUDPProxyMockRecorder) NewAddr(node, service any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewAddr", reflect.TypeOf((*MockNetcForUDPProxy)(nil).NewAddr), node, service)
}
