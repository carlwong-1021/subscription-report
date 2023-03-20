// webhook_configs.proto

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.2
// source: webhook_configs.proto

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
	WebhookConfigs_ShowAvailableWebhookConfig_FullMethodName = "/DeveloperApi.WebhookConfigs/ShowAvailableWebhookConfig"
)

// WebhookConfigsClient is the client API for WebhookConfigs service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WebhookConfigsClient interface {
	ShowAvailableWebhookConfig(ctx context.Context, in *WebhookConfigShowAvailableWebhookConfigRequest, opts ...grpc.CallOption) (*WebhookConfigShowAvailableWebhookConfigReply, error)
}

type webhookConfigsClient struct {
	cc grpc.ClientConnInterface
}

func NewWebhookConfigsClient(cc grpc.ClientConnInterface) WebhookConfigsClient {
	return &webhookConfigsClient{cc}
}

func (c *webhookConfigsClient) ShowAvailableWebhookConfig(ctx context.Context, in *WebhookConfigShowAvailableWebhookConfigRequest, opts ...grpc.CallOption) (*WebhookConfigShowAvailableWebhookConfigReply, error) {
	out := new(WebhookConfigShowAvailableWebhookConfigReply)
	err := c.cc.Invoke(ctx, WebhookConfigs_ShowAvailableWebhookConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WebhookConfigsServer is the server API for WebhookConfigs service.
// All implementations must embed UnimplementedWebhookConfigsServer
// for forward compatibility
type WebhookConfigsServer interface {
	ShowAvailableWebhookConfig(context.Context, *WebhookConfigShowAvailableWebhookConfigRequest) (*WebhookConfigShowAvailableWebhookConfigReply, error)
	mustEmbedUnimplementedWebhookConfigsServer()
}

// UnimplementedWebhookConfigsServer must be embedded to have forward compatible implementations.
type UnimplementedWebhookConfigsServer struct {
}

func (UnimplementedWebhookConfigsServer) ShowAvailableWebhookConfig(context.Context, *WebhookConfigShowAvailableWebhookConfigRequest) (*WebhookConfigShowAvailableWebhookConfigReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowAvailableWebhookConfig not implemented")
}
func (UnimplementedWebhookConfigsServer) mustEmbedUnimplementedWebhookConfigsServer() {}

// UnsafeWebhookConfigsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WebhookConfigsServer will
// result in compilation errors.
type UnsafeWebhookConfigsServer interface {
	mustEmbedUnimplementedWebhookConfigsServer()
}

func RegisterWebhookConfigsServer(s grpc.ServiceRegistrar, srv WebhookConfigsServer) {
	s.RegisterService(&WebhookConfigs_ServiceDesc, srv)
}

func _WebhookConfigs_ShowAvailableWebhookConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WebhookConfigShowAvailableWebhookConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebhookConfigsServer).ShowAvailableWebhookConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WebhookConfigs_ShowAvailableWebhookConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebhookConfigsServer).ShowAvailableWebhookConfig(ctx, req.(*WebhookConfigShowAvailableWebhookConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WebhookConfigs_ServiceDesc is the grpc.ServiceDesc for WebhookConfigs service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WebhookConfigs_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "DeveloperApi.WebhookConfigs",
	HandlerType: (*WebhookConfigsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ShowAvailableWebhookConfig",
			Handler:    _WebhookConfigs_ShowAvailableWebhookConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "webhook_configs.proto",
}
