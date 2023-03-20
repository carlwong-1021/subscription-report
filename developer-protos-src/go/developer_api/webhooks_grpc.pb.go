// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.2
// source: webhooks.proto

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
	Webhooks_Index_FullMethodName  = "/DeveloperApi.Webhooks/Index"
	Webhooks_Create_FullMethodName = "/DeveloperApi.Webhooks/Create"
	Webhooks_Update_FullMethodName = "/DeveloperApi.Webhooks/Update"
)

// WebhooksClient is the client API for Webhooks service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WebhooksClient interface {
	Index(ctx context.Context, in *WebhooksIndexRequest, opts ...grpc.CallOption) (*WebhooksIndexResponse, error)
	Create(ctx context.Context, in *WebhooksCreateRequest, opts ...grpc.CallOption) (*Webhook, error)
	Update(ctx context.Context, in *WebhooksUpdateRequest, opts ...grpc.CallOption) (*Webhook, error)
}

type webhooksClient struct {
	cc grpc.ClientConnInterface
}

func NewWebhooksClient(cc grpc.ClientConnInterface) WebhooksClient {
	return &webhooksClient{cc}
}

func (c *webhooksClient) Index(ctx context.Context, in *WebhooksIndexRequest, opts ...grpc.CallOption) (*WebhooksIndexResponse, error) {
	out := new(WebhooksIndexResponse)
	err := c.cc.Invoke(ctx, Webhooks_Index_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webhooksClient) Create(ctx context.Context, in *WebhooksCreateRequest, opts ...grpc.CallOption) (*Webhook, error) {
	out := new(Webhook)
	err := c.cc.Invoke(ctx, Webhooks_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webhooksClient) Update(ctx context.Context, in *WebhooksUpdateRequest, opts ...grpc.CallOption) (*Webhook, error) {
	out := new(Webhook)
	err := c.cc.Invoke(ctx, Webhooks_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WebhooksServer is the server API for Webhooks service.
// All implementations must embed UnimplementedWebhooksServer
// for forward compatibility
type WebhooksServer interface {
	Index(context.Context, *WebhooksIndexRequest) (*WebhooksIndexResponse, error)
	Create(context.Context, *WebhooksCreateRequest) (*Webhook, error)
	Update(context.Context, *WebhooksUpdateRequest) (*Webhook, error)
	mustEmbedUnimplementedWebhooksServer()
}

// UnimplementedWebhooksServer must be embedded to have forward compatible implementations.
type UnimplementedWebhooksServer struct {
}

func (UnimplementedWebhooksServer) Index(context.Context, *WebhooksIndexRequest) (*WebhooksIndexResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Index not implemented")
}
func (UnimplementedWebhooksServer) Create(context.Context, *WebhooksCreateRequest) (*Webhook, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedWebhooksServer) Update(context.Context, *WebhooksUpdateRequest) (*Webhook, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedWebhooksServer) mustEmbedUnimplementedWebhooksServer() {}

// UnsafeWebhooksServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WebhooksServer will
// result in compilation errors.
type UnsafeWebhooksServer interface {
	mustEmbedUnimplementedWebhooksServer()
}

func RegisterWebhooksServer(s grpc.ServiceRegistrar, srv WebhooksServer) {
	s.RegisterService(&Webhooks_ServiceDesc, srv)
}

func _Webhooks_Index_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WebhooksIndexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebhooksServer).Index(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Webhooks_Index_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebhooksServer).Index(ctx, req.(*WebhooksIndexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Webhooks_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WebhooksCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebhooksServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Webhooks_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebhooksServer).Create(ctx, req.(*WebhooksCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Webhooks_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WebhooksUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebhooksServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Webhooks_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebhooksServer).Update(ctx, req.(*WebhooksUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Webhooks_ServiceDesc is the grpc.ServiceDesc for Webhooks service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Webhooks_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "DeveloperApi.Webhooks",
	HandlerType: (*WebhooksServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Index",
			Handler:    _Webhooks_Index_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Webhooks_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Webhooks_Update_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "webhooks.proto",
}
