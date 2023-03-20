// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.2
// source: webhook_resend_batches.proto

package developer_api

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

const (
	WebhookResendBatches_Create_FullMethodName = "/DeveloperApi.WebhookResendBatches/Create"
	WebhookResendBatches_Index_FullMethodName  = "/DeveloperApi.WebhookResendBatches/Index"
	WebhookResendBatches_Show_FullMethodName   = "/DeveloperApi.WebhookResendBatches/Show"
)

// WebhookResendBatchesClient is the client API for WebhookResendBatches service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WebhookResendBatchesClient interface {
	Create(ctx context.Context, in *WebhookResendBatchCreateRequest, opts ...grpc.CallOption) (*WebhookResendBatchCreateResponse, error)
	Index(ctx context.Context, in *WebhookResendBatchIndexRequest, opts ...grpc.CallOption) (*WebhookResendBatchIndexResponse, error)
	Show(ctx context.Context, in *WebhookResendBatchShowRequest, opts ...grpc.CallOption) (*WebhookResendBatchShowResponse, error)
}

type webhookResendBatchesClient struct {
	cc grpc.ClientConnInterface
}

func NewWebhookResendBatchesClient(cc grpc.ClientConnInterface) WebhookResendBatchesClient {
	return &webhookResendBatchesClient{cc}
}

func (c *webhookResendBatchesClient) Create(ctx context.Context, in *WebhookResendBatchCreateRequest, opts ...grpc.CallOption) (*WebhookResendBatchCreateResponse, error) {
	out := new(WebhookResendBatchCreateResponse)
	err := c.cc.Invoke(ctx, WebhookResendBatches_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webhookResendBatchesClient) Index(ctx context.Context, in *WebhookResendBatchIndexRequest, opts ...grpc.CallOption) (*WebhookResendBatchIndexResponse, error) {
	out := new(WebhookResendBatchIndexResponse)
	err := c.cc.Invoke(ctx, WebhookResendBatches_Index_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webhookResendBatchesClient) Show(ctx context.Context, in *WebhookResendBatchShowRequest, opts ...grpc.CallOption) (*WebhookResendBatchShowResponse, error) {
	out := new(WebhookResendBatchShowResponse)
	err := c.cc.Invoke(ctx, WebhookResendBatches_Show_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WebhookResendBatchesServer is the server API for WebhookResendBatches service.
// All implementations must embed UnimplementedWebhookResendBatchesServer
// for forward compatibility
type WebhookResendBatchesServer interface {
	Create(context.Context, *WebhookResendBatchCreateRequest) (*WebhookResendBatchCreateResponse, error)
	Index(context.Context, *WebhookResendBatchIndexRequest) (*WebhookResendBatchIndexResponse, error)
	Show(context.Context, *WebhookResendBatchShowRequest) (*WebhookResendBatchShowResponse, error)
	mustEmbedUnimplementedWebhookResendBatchesServer()
}

// UnimplementedWebhookResendBatchesServer must be embedded to have forward compatible implementations.
type UnimplementedWebhookResendBatchesServer struct {
}

func (UnimplementedWebhookResendBatchesServer) Create(context.Context, *WebhookResendBatchCreateRequest) (*WebhookResendBatchCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedWebhookResendBatchesServer) Index(context.Context, *WebhookResendBatchIndexRequest) (*WebhookResendBatchIndexResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Index not implemented")
}
func (UnimplementedWebhookResendBatchesServer) Show(context.Context, *WebhookResendBatchShowRequest) (*WebhookResendBatchShowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Show not implemented")
}
func (UnimplementedWebhookResendBatchesServer) mustEmbedUnimplementedWebhookResendBatchesServer() {}

// UnsafeWebhookResendBatchesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WebhookResendBatchesServer will
// result in compilation errors.
type UnsafeWebhookResendBatchesServer interface {
	mustEmbedUnimplementedWebhookResendBatchesServer()
}

func RegisterWebhookResendBatchesServer(s grpc.ServiceRegistrar, srv WebhookResendBatchesServer) {
	s.RegisterService(&WebhookResendBatches_ServiceDesc, srv)
}

func _WebhookResendBatches_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WebhookResendBatchCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebhookResendBatchesServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WebhookResendBatches_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebhookResendBatchesServer).Create(ctx, req.(*WebhookResendBatchCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WebhookResendBatches_Index_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WebhookResendBatchIndexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebhookResendBatchesServer).Index(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WebhookResendBatches_Index_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebhookResendBatchesServer).Index(ctx, req.(*WebhookResendBatchIndexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WebhookResendBatches_Show_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WebhookResendBatchShowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebhookResendBatchesServer).Show(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WebhookResendBatches_Show_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebhookResendBatchesServer).Show(ctx, req.(*WebhookResendBatchShowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WebhookResendBatches_ServiceDesc is the grpc.ServiceDesc for WebhookResendBatches service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WebhookResendBatches_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "DeveloperApi.WebhookResendBatches",
	HandlerType: (*WebhookResendBatchesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _WebhookResendBatches_Create_Handler,
		},
		{
			MethodName: "Index",
			Handler:    _WebhookResendBatches_Index_Handler,
		},
		{
			MethodName: "Show",
			Handler:    _WebhookResendBatches_Show_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "webhook_resend_batches.proto",
}
