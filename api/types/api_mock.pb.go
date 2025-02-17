// Code generated by MockGen. DO NOT EDIT.
// Source: api_grpc.pb.go
//
// Generated by this command:
//
//	mockgen -source=api_grpc.pb.go -destination=api_mock.pb.go -package=types
//

// Package types is a generated GoMock package.
package types

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockTakeHomeServiceClient is a mock of TakeHomeServiceClient interface.
type MockTakeHomeServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockTakeHomeServiceClientMockRecorder
	isgomock struct{}
}

// MockTakeHomeServiceClientMockRecorder is the mock recorder for MockTakeHomeServiceClient.
type MockTakeHomeServiceClientMockRecorder struct {
	mock *MockTakeHomeServiceClient
}

// NewMockTakeHomeServiceClient creates a new mock instance.
func NewMockTakeHomeServiceClient(ctrl *gomock.Controller) *MockTakeHomeServiceClient {
	mock := &MockTakeHomeServiceClient{ctrl: ctrl}
	mock.recorder = &MockTakeHomeServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTakeHomeServiceClient) EXPECT() *MockTakeHomeServiceClientMockRecorder {
	return m.recorder
}

// CreateItem mocks base method.
func (m *MockTakeHomeServiceClient) CreateItem(ctx context.Context, in *CreateItemRequest, opts ...grpc.CallOption) (*CreateItemResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateItem", varargs...)
	ret0, _ := ret[0].(*CreateItemResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateItem indicates an expected call of CreateItem.
func (mr *MockTakeHomeServiceClientMockRecorder) CreateItem(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateItem", reflect.TypeOf((*MockTakeHomeServiceClient)(nil).CreateItem), varargs...)
}

// GetItem mocks base method.
func (m *MockTakeHomeServiceClient) GetItem(ctx context.Context, in *GetItemRequest, opts ...grpc.CallOption) (*GetItemResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetItem", varargs...)
	ret0, _ := ret[0].(*GetItemResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetItem indicates an expected call of GetItem.
func (mr *MockTakeHomeServiceClientMockRecorder) GetItem(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetItem", reflect.TypeOf((*MockTakeHomeServiceClient)(nil).GetItem), varargs...)
}

// GetItems mocks base method.
func (m *MockTakeHomeServiceClient) GetItems(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*GetItemsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetItems", varargs...)
	ret0, _ := ret[0].(*GetItemsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetItems indicates an expected call of GetItems.
func (mr *MockTakeHomeServiceClientMockRecorder) GetItems(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetItems", reflect.TypeOf((*MockTakeHomeServiceClient)(nil).GetItems), varargs...)
}

// MockTakeHomeServiceServer is a mock of TakeHomeServiceServer interface.
type MockTakeHomeServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockTakeHomeServiceServerMockRecorder
	isgomock struct{}
}

// MockTakeHomeServiceServerMockRecorder is the mock recorder for MockTakeHomeServiceServer.
type MockTakeHomeServiceServerMockRecorder struct {
	mock *MockTakeHomeServiceServer
}

// NewMockTakeHomeServiceServer creates a new mock instance.
func NewMockTakeHomeServiceServer(ctrl *gomock.Controller) *MockTakeHomeServiceServer {
	mock := &MockTakeHomeServiceServer{ctrl: ctrl}
	mock.recorder = &MockTakeHomeServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTakeHomeServiceServer) EXPECT() *MockTakeHomeServiceServerMockRecorder {
	return m.recorder
}

// CreateItem mocks base method.
func (m *MockTakeHomeServiceServer) CreateItem(arg0 context.Context, arg1 *CreateItemRequest) (*CreateItemResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateItem", arg0, arg1)
	ret0, _ := ret[0].(*CreateItemResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateItem indicates an expected call of CreateItem.
func (mr *MockTakeHomeServiceServerMockRecorder) CreateItem(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateItem", reflect.TypeOf((*MockTakeHomeServiceServer)(nil).CreateItem), arg0, arg1)
}

// GetItem mocks base method.
func (m *MockTakeHomeServiceServer) GetItem(arg0 context.Context, arg1 *GetItemRequest) (*GetItemResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetItem", arg0, arg1)
	ret0, _ := ret[0].(*GetItemResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetItem indicates an expected call of GetItem.
func (mr *MockTakeHomeServiceServerMockRecorder) GetItem(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetItem", reflect.TypeOf((*MockTakeHomeServiceServer)(nil).GetItem), arg0, arg1)
}

// GetItems mocks base method.
func (m *MockTakeHomeServiceServer) GetItems(arg0 context.Context, arg1 *EmptyRequest) (*GetItemsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetItems", arg0, arg1)
	ret0, _ := ret[0].(*GetItemsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetItems indicates an expected call of GetItems.
func (mr *MockTakeHomeServiceServerMockRecorder) GetItems(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetItems", reflect.TypeOf((*MockTakeHomeServiceServer)(nil).GetItems), arg0, arg1)
}

// mustEmbedUnimplementedTakeHomeServiceServer mocks base method.
func (m *MockTakeHomeServiceServer) mustEmbedUnimplementedTakeHomeServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedTakeHomeServiceServer")
}

// mustEmbedUnimplementedTakeHomeServiceServer indicates an expected call of mustEmbedUnimplementedTakeHomeServiceServer.
func (mr *MockTakeHomeServiceServerMockRecorder) mustEmbedUnimplementedTakeHomeServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedTakeHomeServiceServer", reflect.TypeOf((*MockTakeHomeServiceServer)(nil).mustEmbedUnimplementedTakeHomeServiceServer))
}

// MockUnsafeTakeHomeServiceServer is a mock of UnsafeTakeHomeServiceServer interface.
type MockUnsafeTakeHomeServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeTakeHomeServiceServerMockRecorder
	isgomock struct{}
}

// MockUnsafeTakeHomeServiceServerMockRecorder is the mock recorder for MockUnsafeTakeHomeServiceServer.
type MockUnsafeTakeHomeServiceServerMockRecorder struct {
	mock *MockUnsafeTakeHomeServiceServer
}

// NewMockUnsafeTakeHomeServiceServer creates a new mock instance.
func NewMockUnsafeTakeHomeServiceServer(ctrl *gomock.Controller) *MockUnsafeTakeHomeServiceServer {
	mock := &MockUnsafeTakeHomeServiceServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeTakeHomeServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeTakeHomeServiceServer) EXPECT() *MockUnsafeTakeHomeServiceServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedTakeHomeServiceServer mocks base method.
func (m *MockUnsafeTakeHomeServiceServer) mustEmbedUnimplementedTakeHomeServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedTakeHomeServiceServer")
}

// mustEmbedUnimplementedTakeHomeServiceServer indicates an expected call of mustEmbedUnimplementedTakeHomeServiceServer.
func (mr *MockUnsafeTakeHomeServiceServerMockRecorder) mustEmbedUnimplementedTakeHomeServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedTakeHomeServiceServer", reflect.TypeOf((*MockUnsafeTakeHomeServiceServer)(nil).mustEmbedUnimplementedTakeHomeServiceServer))
}
