// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package execution

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ExecutionAPIClient is the client API for ExecutionAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExecutionAPIClient interface {
	// Ping is used to check if the access node is alive and healthy.
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	// GetAccountAtBlockID gets an account by address at the given block ID
	GetAccountAtBlockID(ctx context.Context, in *GetAccountAtBlockIDRequest, opts ...grpc.CallOption) (*GetAccountAtBlockIDResponse, error)
	// ExecuteScriptAtBlockID executes a ready-only Cadence script against the
	// execution state at the block with the given ID.
	ExecuteScriptAtBlockID(ctx context.Context, in *ExecuteScriptAtBlockIDRequest, opts ...grpc.CallOption) (*ExecuteScriptAtBlockIDResponse, error)
	// GetEventsForBlockIDs retrieves events for all the specified block IDs that
	// have the given type
	GetEventsForBlockIDs(ctx context.Context, in *GetEventsForBlockIDsRequest, opts ...grpc.CallOption) (*GetEventsForBlockIDsResponse, error)
	// GetTransactionResult gets the result of a transaction.
	GetTransactionResult(ctx context.Context, in *GetTransactionResultRequest, opts ...grpc.CallOption) (*GetTransactionResultResponse, error)
	// GetTransactionResultByIndex gets the result of a transaction at the index.
	GetTransactionResultByIndex(ctx context.Context, in *GetTransactionByIndexRequest, opts ...grpc.CallOption) (*GetTransactionResultResponse, error)
	// GetTransactionResultByIndex gets the results of all transactions in the
	// block ordered by transaction index.
	GetTransactionResultsByBlockID(ctx context.Context, in *GetTransactionsByBlockIDRequest, opts ...grpc.CallOption) (*GetTransactionResultsResponse, error)
	// GetTransactionErrorMessage gets the error messages of a failed transaction by id.
	GetTransactionErrorMessage(ctx context.Context, in *GetTransactionErrorMessageRequest, opts ...grpc.CallOption) (*GetTransactionErrorMessagesResponse, error)
	// GetTransactionErrorMessageByIndex gets the error messages of a failed transaction at the index.
	GetTransactionErrorMessageByIndex(ctx context.Context, in *GetTransactionErrorMessageByIndexRequest, opts ...grpc.CallOption) (*GetTransactionErrorMessagesResponse, error)
	// GetTransactionErrorMessagesByBlockID gets the error messages of all failed transactions in the
	// block ordered by transaction index.
	GetTransactionErrorMessagesByBlockID(ctx context.Context, in *GetTransactionErrorMessagesByBlockIDRequest, opts ...grpc.CallOption) (*GetTransactionErrorMessagesResponse, error)
	// GetRegisterAtBlockID collects a register at the block with the given ID (if
	// available).
	GetRegisterAtBlockID(ctx context.Context, in *GetRegisterAtBlockIDRequest, opts ...grpc.CallOption) (*GetRegisterAtBlockIDResponse, error)
	// GetLatestBlockHeader gets the latest sealed or unsealed block header.
	GetLatestBlockHeader(ctx context.Context, in *GetLatestBlockHeaderRequest, opts ...grpc.CallOption) (*BlockHeaderResponse, error)
	// GetBlockHeaderByID gets a block header by ID.
	GetBlockHeaderByID(ctx context.Context, in *GetBlockHeaderByIDRequest, opts ...grpc.CallOption) (*BlockHeaderResponse, error)
}

type executionAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewExecutionAPIClient(cc grpc.ClientConnInterface) ExecutionAPIClient {
	return &executionAPIClient{cc}
}

func (c *executionAPIClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, "/flow.execution.ExecutionAPI/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *executionAPIClient) GetAccountAtBlockID(ctx context.Context, in *GetAccountAtBlockIDRequest, opts ...grpc.CallOption) (*GetAccountAtBlockIDResponse, error) {
	out := new(GetAccountAtBlockIDResponse)
	err := c.cc.Invoke(ctx, "/flow.execution.ExecutionAPI/GetAccountAtBlockID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *executionAPIClient) ExecuteScriptAtBlockID(ctx context.Context, in *ExecuteScriptAtBlockIDRequest, opts ...grpc.CallOption) (*ExecuteScriptAtBlockIDResponse, error) {
	out := new(ExecuteScriptAtBlockIDResponse)
	err := c.cc.Invoke(ctx, "/flow.execution.ExecutionAPI/ExecuteScriptAtBlockID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *executionAPIClient) GetEventsForBlockIDs(ctx context.Context, in *GetEventsForBlockIDsRequest, opts ...grpc.CallOption) (*GetEventsForBlockIDsResponse, error) {
	out := new(GetEventsForBlockIDsResponse)
	err := c.cc.Invoke(ctx, "/flow.execution.ExecutionAPI/GetEventsForBlockIDs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *executionAPIClient) GetTransactionResult(ctx context.Context, in *GetTransactionResultRequest, opts ...grpc.CallOption) (*GetTransactionResultResponse, error) {
	out := new(GetTransactionResultResponse)
	err := c.cc.Invoke(ctx, "/flow.execution.ExecutionAPI/GetTransactionResult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *executionAPIClient) GetTransactionResultByIndex(ctx context.Context, in *GetTransactionByIndexRequest, opts ...grpc.CallOption) (*GetTransactionResultResponse, error) {
	out := new(GetTransactionResultResponse)
	err := c.cc.Invoke(ctx, "/flow.execution.ExecutionAPI/GetTransactionResultByIndex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *executionAPIClient) GetTransactionResultsByBlockID(ctx context.Context, in *GetTransactionsByBlockIDRequest, opts ...grpc.CallOption) (*GetTransactionResultsResponse, error) {
	out := new(GetTransactionResultsResponse)
	err := c.cc.Invoke(ctx, "/flow.execution.ExecutionAPI/GetTransactionResultsByBlockID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *executionAPIClient) GetTransactionErrorMessage(ctx context.Context, in *GetTransactionErrorMessageRequest, opts ...grpc.CallOption) (*GetTransactionErrorMessagesResponse, error) {
	out := new(GetTransactionErrorMessagesResponse)
	err := c.cc.Invoke(ctx, "/flow.execution.ExecutionAPI/GetTransactionErrorMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *executionAPIClient) GetTransactionErrorMessageByIndex(ctx context.Context, in *GetTransactionErrorMessageByIndexRequest, opts ...grpc.CallOption) (*GetTransactionErrorMessagesResponse, error) {
	out := new(GetTransactionErrorMessagesResponse)
	err := c.cc.Invoke(ctx, "/flow.execution.ExecutionAPI/GetTransactionErrorMessageByIndex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *executionAPIClient) GetTransactionErrorMessagesByBlockID(ctx context.Context, in *GetTransactionErrorMessagesByBlockIDRequest, opts ...grpc.CallOption) (*GetTransactionErrorMessagesResponse, error) {
	out := new(GetTransactionErrorMessagesResponse)
	err := c.cc.Invoke(ctx, "/flow.execution.ExecutionAPI/GetTransactionErrorMessagesByBlockID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *executionAPIClient) GetRegisterAtBlockID(ctx context.Context, in *GetRegisterAtBlockIDRequest, opts ...grpc.CallOption) (*GetRegisterAtBlockIDResponse, error) {
	out := new(GetRegisterAtBlockIDResponse)
	err := c.cc.Invoke(ctx, "/flow.execution.ExecutionAPI/GetRegisterAtBlockID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *executionAPIClient) GetLatestBlockHeader(ctx context.Context, in *GetLatestBlockHeaderRequest, opts ...grpc.CallOption) (*BlockHeaderResponse, error) {
	out := new(BlockHeaderResponse)
	err := c.cc.Invoke(ctx, "/flow.execution.ExecutionAPI/GetLatestBlockHeader", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *executionAPIClient) GetBlockHeaderByID(ctx context.Context, in *GetBlockHeaderByIDRequest, opts ...grpc.CallOption) (*BlockHeaderResponse, error) {
	out := new(BlockHeaderResponse)
	err := c.cc.Invoke(ctx, "/flow.execution.ExecutionAPI/GetBlockHeaderByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExecutionAPIServer is the server API for ExecutionAPI service.
// All implementations should embed UnimplementedExecutionAPIServer
// for forward compatibility
type ExecutionAPIServer interface {
	// Ping is used to check if the access node is alive and healthy.
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	// GetAccountAtBlockID gets an account by address at the given block ID
	GetAccountAtBlockID(context.Context, *GetAccountAtBlockIDRequest) (*GetAccountAtBlockIDResponse, error)
	// ExecuteScriptAtBlockID executes a ready-only Cadence script against the
	// execution state at the block with the given ID.
	ExecuteScriptAtBlockID(context.Context, *ExecuteScriptAtBlockIDRequest) (*ExecuteScriptAtBlockIDResponse, error)
	// GetEventsForBlockIDs retrieves events for all the specified block IDs that
	// have the given type
	GetEventsForBlockIDs(context.Context, *GetEventsForBlockIDsRequest) (*GetEventsForBlockIDsResponse, error)
	// GetTransactionResult gets the result of a transaction.
	GetTransactionResult(context.Context, *GetTransactionResultRequest) (*GetTransactionResultResponse, error)
	// GetTransactionResultByIndex gets the result of a transaction at the index.
	GetTransactionResultByIndex(context.Context, *GetTransactionByIndexRequest) (*GetTransactionResultResponse, error)
	// GetTransactionResultByIndex gets the results of all transactions in the
	// block ordered by transaction index.
	GetTransactionResultsByBlockID(context.Context, *GetTransactionsByBlockIDRequest) (*GetTransactionResultsResponse, error)
	// GetTransactionErrorMessage gets the error messages of a failed transaction by id.
	GetTransactionErrorMessage(context.Context, *GetTransactionErrorMessageRequest) (*GetTransactionErrorMessagesResponse, error)
	// GetTransactionErrorMessageByIndex gets the error messages of a failed transaction at the index.
	GetTransactionErrorMessageByIndex(context.Context, *GetTransactionErrorMessageByIndexRequest) (*GetTransactionErrorMessagesResponse, error)
	// GetTransactionErrorMessagesByBlockID gets the error messages of all failed transactions in the
	// block ordered by transaction index.
	GetTransactionErrorMessagesByBlockID(context.Context, *GetTransactionErrorMessagesByBlockIDRequest) (*GetTransactionErrorMessagesResponse, error)
	// GetRegisterAtBlockID collects a register at the block with the given ID (if
	// available).
	GetRegisterAtBlockID(context.Context, *GetRegisterAtBlockIDRequest) (*GetRegisterAtBlockIDResponse, error)
	// GetLatestBlockHeader gets the latest sealed or unsealed block header.
	GetLatestBlockHeader(context.Context, *GetLatestBlockHeaderRequest) (*BlockHeaderResponse, error)
	// GetBlockHeaderByID gets a block header by ID.
	GetBlockHeaderByID(context.Context, *GetBlockHeaderByIDRequest) (*BlockHeaderResponse, error)
}

// UnimplementedExecutionAPIServer should be embedded to have forward compatible implementations.
type UnimplementedExecutionAPIServer struct {
}

func (UnimplementedExecutionAPIServer) Ping(context.Context, *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedExecutionAPIServer) GetAccountAtBlockID(context.Context, *GetAccountAtBlockIDRequest) (*GetAccountAtBlockIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccountAtBlockID not implemented")
}
func (UnimplementedExecutionAPIServer) ExecuteScriptAtBlockID(context.Context, *ExecuteScriptAtBlockIDRequest) (*ExecuteScriptAtBlockIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecuteScriptAtBlockID not implemented")
}
func (UnimplementedExecutionAPIServer) GetEventsForBlockIDs(context.Context, *GetEventsForBlockIDsRequest) (*GetEventsForBlockIDsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEventsForBlockIDs not implemented")
}
func (UnimplementedExecutionAPIServer) GetTransactionResult(context.Context, *GetTransactionResultRequest) (*GetTransactionResultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransactionResult not implemented")
}
func (UnimplementedExecutionAPIServer) GetTransactionResultByIndex(context.Context, *GetTransactionByIndexRequest) (*GetTransactionResultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransactionResultByIndex not implemented")
}
func (UnimplementedExecutionAPIServer) GetTransactionResultsByBlockID(context.Context, *GetTransactionsByBlockIDRequest) (*GetTransactionResultsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransactionResultsByBlockID not implemented")
}
func (UnimplementedExecutionAPIServer) GetTransactionErrorMessage(context.Context, *GetTransactionErrorMessageRequest) (*GetTransactionErrorMessagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransactionErrorMessage not implemented")
}
func (UnimplementedExecutionAPIServer) GetTransactionErrorMessageByIndex(context.Context, *GetTransactionErrorMessageByIndexRequest) (*GetTransactionErrorMessagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransactionErrorMessageByIndex not implemented")
}
func (UnimplementedExecutionAPIServer) GetTransactionErrorMessagesByBlockID(context.Context, *GetTransactionErrorMessagesByBlockIDRequest) (*GetTransactionErrorMessagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransactionErrorMessagesByBlockID not implemented")
}
func (UnimplementedExecutionAPIServer) GetRegisterAtBlockID(context.Context, *GetRegisterAtBlockIDRequest) (*GetRegisterAtBlockIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRegisterAtBlockID not implemented")
}
func (UnimplementedExecutionAPIServer) GetLatestBlockHeader(context.Context, *GetLatestBlockHeaderRequest) (*BlockHeaderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLatestBlockHeader not implemented")
}
func (UnimplementedExecutionAPIServer) GetBlockHeaderByID(context.Context, *GetBlockHeaderByIDRequest) (*BlockHeaderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlockHeaderByID not implemented")
}

// UnsafeExecutionAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExecutionAPIServer will
// result in compilation errors.
type UnsafeExecutionAPIServer interface {
	mustEmbedUnimplementedExecutionAPIServer()
}

func RegisterExecutionAPIServer(s grpc.ServiceRegistrar, srv ExecutionAPIServer) {
	s.RegisterService(&ExecutionAPI_ServiceDesc, srv)
}

func _ExecutionAPI_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutionAPIServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flow.execution.ExecutionAPI/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutionAPIServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExecutionAPI_GetAccountAtBlockID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccountAtBlockIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutionAPIServer).GetAccountAtBlockID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flow.execution.ExecutionAPI/GetAccountAtBlockID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutionAPIServer).GetAccountAtBlockID(ctx, req.(*GetAccountAtBlockIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExecutionAPI_ExecuteScriptAtBlockID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecuteScriptAtBlockIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutionAPIServer).ExecuteScriptAtBlockID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flow.execution.ExecutionAPI/ExecuteScriptAtBlockID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutionAPIServer).ExecuteScriptAtBlockID(ctx, req.(*ExecuteScriptAtBlockIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExecutionAPI_GetEventsForBlockIDs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEventsForBlockIDsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutionAPIServer).GetEventsForBlockIDs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flow.execution.ExecutionAPI/GetEventsForBlockIDs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutionAPIServer).GetEventsForBlockIDs(ctx, req.(*GetEventsForBlockIDsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExecutionAPI_GetTransactionResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransactionResultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutionAPIServer).GetTransactionResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flow.execution.ExecutionAPI/GetTransactionResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutionAPIServer).GetTransactionResult(ctx, req.(*GetTransactionResultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExecutionAPI_GetTransactionResultByIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransactionByIndexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutionAPIServer).GetTransactionResultByIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flow.execution.ExecutionAPI/GetTransactionResultByIndex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutionAPIServer).GetTransactionResultByIndex(ctx, req.(*GetTransactionByIndexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExecutionAPI_GetTransactionResultsByBlockID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransactionsByBlockIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutionAPIServer).GetTransactionResultsByBlockID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flow.execution.ExecutionAPI/GetTransactionResultsByBlockID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutionAPIServer).GetTransactionResultsByBlockID(ctx, req.(*GetTransactionsByBlockIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExecutionAPI_GetTransactionErrorMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransactionErrorMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutionAPIServer).GetTransactionErrorMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flow.execution.ExecutionAPI/GetTransactionErrorMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutionAPIServer).GetTransactionErrorMessage(ctx, req.(*GetTransactionErrorMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExecutionAPI_GetTransactionErrorMessageByIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransactionErrorMessageByIndexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutionAPIServer).GetTransactionErrorMessageByIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flow.execution.ExecutionAPI/GetTransactionErrorMessageByIndex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutionAPIServer).GetTransactionErrorMessageByIndex(ctx, req.(*GetTransactionErrorMessageByIndexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExecutionAPI_GetTransactionErrorMessagesByBlockID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransactionErrorMessagesByBlockIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutionAPIServer).GetTransactionErrorMessagesByBlockID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flow.execution.ExecutionAPI/GetTransactionErrorMessagesByBlockID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutionAPIServer).GetTransactionErrorMessagesByBlockID(ctx, req.(*GetTransactionErrorMessagesByBlockIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExecutionAPI_GetRegisterAtBlockID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRegisterAtBlockIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutionAPIServer).GetRegisterAtBlockID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flow.execution.ExecutionAPI/GetRegisterAtBlockID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutionAPIServer).GetRegisterAtBlockID(ctx, req.(*GetRegisterAtBlockIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExecutionAPI_GetLatestBlockHeader_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLatestBlockHeaderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutionAPIServer).GetLatestBlockHeader(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flow.execution.ExecutionAPI/GetLatestBlockHeader",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutionAPIServer).GetLatestBlockHeader(ctx, req.(*GetLatestBlockHeaderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExecutionAPI_GetBlockHeaderByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlockHeaderByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutionAPIServer).GetBlockHeaderByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flow.execution.ExecutionAPI/GetBlockHeaderByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutionAPIServer).GetBlockHeaderByID(ctx, req.(*GetBlockHeaderByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ExecutionAPI_ServiceDesc is the grpc.ServiceDesc for ExecutionAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ExecutionAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "flow.execution.ExecutionAPI",
	HandlerType: (*ExecutionAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _ExecutionAPI_Ping_Handler,
		},
		{
			MethodName: "GetAccountAtBlockID",
			Handler:    _ExecutionAPI_GetAccountAtBlockID_Handler,
		},
		{
			MethodName: "ExecuteScriptAtBlockID",
			Handler:    _ExecutionAPI_ExecuteScriptAtBlockID_Handler,
		},
		{
			MethodName: "GetEventsForBlockIDs",
			Handler:    _ExecutionAPI_GetEventsForBlockIDs_Handler,
		},
		{
			MethodName: "GetTransactionResult",
			Handler:    _ExecutionAPI_GetTransactionResult_Handler,
		},
		{
			MethodName: "GetTransactionResultByIndex",
			Handler:    _ExecutionAPI_GetTransactionResultByIndex_Handler,
		},
		{
			MethodName: "GetTransactionResultsByBlockID",
			Handler:    _ExecutionAPI_GetTransactionResultsByBlockID_Handler,
		},
		{
			MethodName: "GetTransactionErrorMessage",
			Handler:    _ExecutionAPI_GetTransactionErrorMessage_Handler,
		},
		{
			MethodName: "GetTransactionErrorMessageByIndex",
			Handler:    _ExecutionAPI_GetTransactionErrorMessageByIndex_Handler,
		},
		{
			MethodName: "GetTransactionErrorMessagesByBlockID",
			Handler:    _ExecutionAPI_GetTransactionErrorMessagesByBlockID_Handler,
		},
		{
			MethodName: "GetRegisterAtBlockID",
			Handler:    _ExecutionAPI_GetRegisterAtBlockID_Handler,
		},
		{
			MethodName: "GetLatestBlockHeader",
			Handler:    _ExecutionAPI_GetLatestBlockHeader_Handler,
		},
		{
			MethodName: "GetBlockHeaderByID",
			Handler:    _ExecutionAPI_GetBlockHeaderByID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "flow/execution/execution.proto",
}
