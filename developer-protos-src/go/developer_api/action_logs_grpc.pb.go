// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.2
// source: action_logs.proto

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
	ActionLogs_Index_FullMethodName = "/DeveloperApi.ActionLogs/Index"
)

// ActionLogsClient is the client API for ActionLogs service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ActionLogsClient interface {
	Index(ctx context.Context, in *ActionLogsIndexRequest, opts ...grpc.CallOption) (*ActionLogsIndexResponse, error)
}

type actionLogsClient struct {
	cc grpc.ClientConnInterface
}

func NewActionLogsClient(cc grpc.ClientConnInterface) ActionLogsClient {
	return &actionLogsClient{cc}
}

func (c *actionLogsClient) Index(ctx context.Context, in *ActionLogsIndexRequest, opts ...grpc.CallOption) (*ActionLogsIndexResponse, error) {
	out := new(ActionLogsIndexResponse)
	err := c.cc.Invoke(ctx, ActionLogs_Index_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ActionLogsServer is the server API for ActionLogs service.
// All implementations must embed UnimplementedActionLogsServer
// for forward compatibility
type ActionLogsServer interface {
	Index(context.Context, *ActionLogsIndexRequest) (*ActionLogsIndexResponse, error)
	mustEmbedUnimplementedActionLogsServer()
}

// UnimplementedActionLogsServer must be embedded to have forward compatible implementations.
type UnimplementedActionLogsServer struct {
}

func (UnimplementedActionLogsServer) Index(context.Context, *ActionLogsIndexRequest) (*ActionLogsIndexResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Index not implemented")
}
func (UnimplementedActionLogsServer) mustEmbedUnimplementedActionLogsServer() {}

// UnsafeActionLogsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ActionLogsServer will
// result in compilation errors.
type UnsafeActionLogsServer interface {
	mustEmbedUnimplementedActionLogsServer()
}

func RegisterActionLogsServer(s grpc.ServiceRegistrar, srv ActionLogsServer) {
	s.RegisterService(&ActionLogs_ServiceDesc, srv)
}

func _ActionLogs_Index_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActionLogsIndexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionLogsServer).Index(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ActionLogs_Index_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionLogsServer).Index(ctx, req.(*ActionLogsIndexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ActionLogs_ServiceDesc is the grpc.ServiceDesc for ActionLogs service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ActionLogs_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "DeveloperApi.ActionLogs",
	HandlerType: (*ActionLogsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Index",
			Handler:    _ActionLogs_Index_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "action_logs.proto",
}