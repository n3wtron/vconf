// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/grpcapi/vecosy.pb.go

// Package grpcapi is a generated GoMock package.
package grpcapi

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	metadata "google.golang.org/grpc/metadata"
	reflect "reflect"
)

// MockSmartConfigClient is a mock of SmartConfigClient interface
type MockSmartConfigClient struct {
	ctrl     *gomock.Controller
	recorder *MockSmartConfigClientMockRecorder
}

// MockSmartConfigClientMockRecorder is the mock recorder for MockSmartConfigClient
type MockSmartConfigClientMockRecorder struct {
	mock *MockSmartConfigClient
}

// NewMockSmartConfigClient creates a new mock instance
func NewMockSmartConfigClient(ctrl *gomock.Controller) *MockSmartConfigClient {
	mock := &MockSmartConfigClient{ctrl: ctrl}
	mock.recorder = &MockSmartConfigClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSmartConfigClient) EXPECT() *MockSmartConfigClientMockRecorder {
	return m.recorder
}

// GetConfig mocks base method
func (m *MockSmartConfigClient) GetConfig(ctx context.Context, in *GetConfigRequest, opts ...grpc.CallOption) (*GetConfigResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetConfig", varargs...)
	ret0, _ := ret[0].(*GetConfigResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConfig indicates an expected call of GetConfig
func (mr *MockSmartConfigClientMockRecorder) GetConfig(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfig", reflect.TypeOf((*MockSmartConfigClient)(nil).GetConfig), varargs...)
}

// MockSmartConfigServer is a mock of SmartConfigServer interface
type MockSmartConfigServer struct {
	ctrl     *gomock.Controller
	recorder *MockSmartConfigServerMockRecorder
}

// MockSmartConfigServerMockRecorder is the mock recorder for MockSmartConfigServer
type MockSmartConfigServerMockRecorder struct {
	mock *MockSmartConfigServer
}

// NewMockSmartConfigServer creates a new mock instance
func NewMockSmartConfigServer(ctrl *gomock.Controller) *MockSmartConfigServer {
	mock := &MockSmartConfigServer{ctrl: ctrl}
	mock.recorder = &MockSmartConfigServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSmartConfigServer) EXPECT() *MockSmartConfigServerMockRecorder {
	return m.recorder
}

// GetConfig mocks base method
func (m *MockSmartConfigServer) GetConfig(arg0 context.Context, arg1 *GetConfigRequest) (*GetConfigResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConfig", arg0, arg1)
	ret0, _ := ret[0].(*GetConfigResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConfig indicates an expected call of GetConfig
func (mr *MockSmartConfigServerMockRecorder) GetConfig(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfig", reflect.TypeOf((*MockSmartConfigServer)(nil).GetConfig), arg0, arg1)
}

// MockRawClient is a mock of RawClient interface
type MockRawClient struct {
	ctrl     *gomock.Controller
	recorder *MockRawClientMockRecorder
}

// MockRawClientMockRecorder is the mock recorder for MockRawClient
type MockRawClientMockRecorder struct {
	mock *MockRawClient
}

// NewMockRawClient creates a new mock instance
func NewMockRawClient(ctrl *gomock.Controller) *MockRawClient {
	mock := &MockRawClient{ctrl: ctrl}
	mock.recorder = &MockRawClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRawClient) EXPECT() *MockRawClientMockRecorder {
	return m.recorder
}

// GetFile mocks base method
func (m *MockRawClient) GetFile(ctx context.Context, in *GetFileRequest, opts ...grpc.CallOption) (*GetFileResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "getFile", varargs...)
	ret0, _ := ret[0].(*GetFileResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFile indicates an expected call of GetFile
func (mr *MockRawClientMockRecorder) GetFile(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getFile", reflect.TypeOf((*MockRawClient)(nil).GetFile), varargs...)
}

// MockRawServer is a mock of RawServer interface
type MockRawServer struct {
	ctrl     *gomock.Controller
	recorder *MockRawServerMockRecorder
}

// MockRawServerMockRecorder is the mock recorder for MockRawServer
type MockRawServerMockRecorder struct {
	mock *MockRawServer
}

// NewMockRawServer creates a new mock instance
func NewMockRawServer(ctrl *gomock.Controller) *MockRawServer {
	mock := &MockRawServer{ctrl: ctrl}
	mock.recorder = &MockRawServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRawServer) EXPECT() *MockRawServerMockRecorder {
	return m.recorder
}

// GetFile mocks base method
func (m *MockRawServer) GetFile(arg0 context.Context, arg1 *GetFileRequest) (*GetFileResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "getFile", arg0, arg1)
	ret0, _ := ret[0].(*GetFileResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFile indicates an expected call of GetFile
func (mr *MockRawServerMockRecorder) GetFile(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getFile", reflect.TypeOf((*MockRawServer)(nil).GetFile), arg0, arg1)
}

// MockWatchServiceClient is a mock of WatchServiceClient interface
type MockWatchServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockWatchServiceClientMockRecorder
}

// MockWatchServiceClientMockRecorder is the mock recorder for MockWatchServiceClient
type MockWatchServiceClientMockRecorder struct {
	mock *MockWatchServiceClient
}

// NewMockWatchServiceClient creates a new mock instance
func NewMockWatchServiceClient(ctrl *gomock.Controller) *MockWatchServiceClient {
	mock := &MockWatchServiceClient{ctrl: ctrl}
	mock.recorder = &MockWatchServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockWatchServiceClient) EXPECT() *MockWatchServiceClientMockRecorder {
	return m.recorder
}

// Watch mocks base method
func (m *MockWatchServiceClient) Watch(ctx context.Context, in *WatchRequest, opts ...grpc.CallOption) (WatchService_WatchClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Watch", varargs...)
	ret0, _ := ret[0].(WatchService_WatchClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Watch indicates an expected call of Watch
func (mr *MockWatchServiceClientMockRecorder) Watch(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockWatchServiceClient)(nil).Watch), varargs...)
}

// MockWatchService_WatchClient is a mock of WatchService_WatchClient interface
type MockWatchService_WatchClient struct {
	ctrl     *gomock.Controller
	recorder *MockWatchService_WatchClientMockRecorder
}

// MockWatchService_WatchClientMockRecorder is the mock recorder for MockWatchService_WatchClient
type MockWatchService_WatchClientMockRecorder struct {
	mock *MockWatchService_WatchClient
}

// NewMockWatchService_WatchClient creates a new mock instance
func NewMockWatchService_WatchClient(ctrl *gomock.Controller) *MockWatchService_WatchClient {
	mock := &MockWatchService_WatchClient{ctrl: ctrl}
	mock.recorder = &MockWatchService_WatchClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockWatchService_WatchClient) EXPECT() *MockWatchService_WatchClientMockRecorder {
	return m.recorder
}

// Recv mocks base method
func (m *MockWatchService_WatchClient) Recv() (*WatchResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*WatchResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv
func (mr *MockWatchService_WatchClientMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockWatchService_WatchClient)(nil).Recv))
}

// Header mocks base method
func (m *MockWatchService_WatchClient) Header() (metadata.MD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header
func (mr *MockWatchService_WatchClientMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockWatchService_WatchClient)(nil).Header))
}

// Trailer mocks base method
func (m *MockWatchService_WatchClient) Trailer() metadata.MD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer
func (mr *MockWatchService_WatchClientMockRecorder) Trailer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockWatchService_WatchClient)(nil).Trailer))
}

// CloseSend mocks base method
func (m *MockWatchService_WatchClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend
func (mr *MockWatchService_WatchClientMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockWatchService_WatchClient)(nil).CloseSend))
}

// Context mocks base method
func (m *MockWatchService_WatchClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (mr *MockWatchService_WatchClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockWatchService_WatchClient)(nil).Context))
}

// SendMsg mocks base method
func (m_2 *MockWatchService_WatchClient) SendMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "SendMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg
func (mr *MockWatchService_WatchClientMockRecorder) SendMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockWatchService_WatchClient)(nil).SendMsg), m)
}

// RecvMsg mocks base method
func (m_2 *MockWatchService_WatchClient) RecvMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "RecvMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg
func (mr *MockWatchService_WatchClientMockRecorder) RecvMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockWatchService_WatchClient)(nil).RecvMsg), m)
}

// MockWatchServiceServer is a mock of WatchServiceServer interface
type MockWatchServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockWatchServiceServerMockRecorder
}

// MockWatchServiceServerMockRecorder is the mock recorder for MockWatchServiceServer
type MockWatchServiceServerMockRecorder struct {
	mock *MockWatchServiceServer
}

// NewMockWatchServiceServer creates a new mock instance
func NewMockWatchServiceServer(ctrl *gomock.Controller) *MockWatchServiceServer {
	mock := &MockWatchServiceServer{ctrl: ctrl}
	mock.recorder = &MockWatchServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockWatchServiceServer) EXPECT() *MockWatchServiceServerMockRecorder {
	return m.recorder
}

// Watch mocks base method
func (m *MockWatchServiceServer) Watch(arg0 *WatchRequest, arg1 WatchService_WatchServer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Watch", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Watch indicates an expected call of Watch
func (mr *MockWatchServiceServerMockRecorder) Watch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockWatchServiceServer)(nil).Watch), arg0, arg1)
}

// MockWatchService_WatchServer is a mock of WatchService_WatchServer interface
type MockWatchService_WatchServer struct {
	ctrl     *gomock.Controller
	recorder *MockWatchService_WatchServerMockRecorder
}

// MockWatchService_WatchServerMockRecorder is the mock recorder for MockWatchService_WatchServer
type MockWatchService_WatchServerMockRecorder struct {
	mock *MockWatchService_WatchServer
}

// NewMockWatchService_WatchServer creates a new mock instance
func NewMockWatchService_WatchServer(ctrl *gomock.Controller) *MockWatchService_WatchServer {
	mock := &MockWatchService_WatchServer{ctrl: ctrl}
	mock.recorder = &MockWatchService_WatchServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockWatchService_WatchServer) EXPECT() *MockWatchService_WatchServerMockRecorder {
	return m.recorder
}

// Send mocks base method
func (m *MockWatchService_WatchServer) Send(arg0 *WatchResponse) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send
func (mr *MockWatchService_WatchServerMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockWatchService_WatchServer)(nil).Send), arg0)
}

// SetHeader mocks base method
func (m *MockWatchService_WatchServer) SetHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeader indicates an expected call of SetHeader
func (mr *MockWatchService_WatchServerMockRecorder) SetHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockWatchService_WatchServer)(nil).SetHeader), arg0)
}

// SendHeader mocks base method
func (m *MockWatchService_WatchServer) SendHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHeader indicates an expected call of SendHeader
func (mr *MockWatchService_WatchServerMockRecorder) SendHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeader", reflect.TypeOf((*MockWatchService_WatchServer)(nil).SendHeader), arg0)
}

// SetTrailer mocks base method
func (m *MockWatchService_WatchServer) SetTrailer(arg0 metadata.MD) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTrailer", arg0)
}

// SetTrailer indicates an expected call of SetTrailer
func (mr *MockWatchService_WatchServerMockRecorder) SetTrailer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrailer", reflect.TypeOf((*MockWatchService_WatchServer)(nil).SetTrailer), arg0)
}

// Context mocks base method
func (m *MockWatchService_WatchServer) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (mr *MockWatchService_WatchServerMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockWatchService_WatchServer)(nil).Context))
}

// SendMsg mocks base method
func (m_2 *MockWatchService_WatchServer) SendMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "SendMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg
func (mr *MockWatchService_WatchServerMockRecorder) SendMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockWatchService_WatchServer)(nil).SendMsg), m)
}

// RecvMsg mocks base method
func (m_2 *MockWatchService_WatchServer) RecvMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "RecvMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg
func (mr *MockWatchService_WatchServerMockRecorder) RecvMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockWatchService_WatchServer)(nil).RecvMsg), m)
}
