// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.2
// source: merchant.proto

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
	Merchant_PurchaseAppIfNotBoughtAndFree_FullMethodName = "/DeveloperApi.Merchant/PurchaseAppIfNotBoughtAndFree"
	Merchant_ReauthorizeApp_FullMethodName                = "/DeveloperApi.Merchant/ReauthorizeApp"
	Merchant_InstallFreeTrialApp_FullMethodName           = "/DeveloperApi.Merchant/InstallFreeTrialApp"
	Merchant_RecoverMerchantWebhookStatus_FullMethodName  = "/DeveloperApi.Merchant/RecoverMerchantWebhookStatus"
)

// MerchantClient is the client API for Merchant service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MerchantClient interface {
	PurchaseAppIfNotBoughtAndFree(ctx context.Context, in *MerchantPurchaseAppIfNotBoughtAndFreeRequest, opts ...grpc.CallOption) (*MerchantPurchaseAppIfNotBoughtAndFreeReply, error)
	ReauthorizeApp(ctx context.Context, in *ReauthorizeAppRequest, opts ...grpc.CallOption) (*ReauthorizeAppReply, error)
	InstallFreeTrialApp(ctx context.Context, in *InstallFreeTrialAppRequest, opts ...grpc.CallOption) (*InstallFreeTrialAppReply, error)
	RecoverMerchantWebhookStatus(ctx context.Context, in *RecoverMerchantWebhookStatusRequest, opts ...grpc.CallOption) (*RecoverMerchantWebhookStatusReply, error)
}

type merchantClient struct {
	cc grpc.ClientConnInterface
}

func NewMerchantClient(cc grpc.ClientConnInterface) MerchantClient {
	return &merchantClient{cc}
}

func (c *merchantClient) PurchaseAppIfNotBoughtAndFree(ctx context.Context, in *MerchantPurchaseAppIfNotBoughtAndFreeRequest, opts ...grpc.CallOption) (*MerchantPurchaseAppIfNotBoughtAndFreeReply, error) {
	out := new(MerchantPurchaseAppIfNotBoughtAndFreeReply)
	err := c.cc.Invoke(ctx, Merchant_PurchaseAppIfNotBoughtAndFree_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *merchantClient) ReauthorizeApp(ctx context.Context, in *ReauthorizeAppRequest, opts ...grpc.CallOption) (*ReauthorizeAppReply, error) {
	out := new(ReauthorizeAppReply)
	err := c.cc.Invoke(ctx, Merchant_ReauthorizeApp_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *merchantClient) InstallFreeTrialApp(ctx context.Context, in *InstallFreeTrialAppRequest, opts ...grpc.CallOption) (*InstallFreeTrialAppReply, error) {
	out := new(InstallFreeTrialAppReply)
	err := c.cc.Invoke(ctx, Merchant_InstallFreeTrialApp_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *merchantClient) RecoverMerchantWebhookStatus(ctx context.Context, in *RecoverMerchantWebhookStatusRequest, opts ...grpc.CallOption) (*RecoverMerchantWebhookStatusReply, error) {
	out := new(RecoverMerchantWebhookStatusReply)
	err := c.cc.Invoke(ctx, Merchant_RecoverMerchantWebhookStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MerchantServer is the server API for Merchant service.
// All implementations must embed UnimplementedMerchantServer
// for forward compatibility
type MerchantServer interface {
	PurchaseAppIfNotBoughtAndFree(context.Context, *MerchantPurchaseAppIfNotBoughtAndFreeRequest) (*MerchantPurchaseAppIfNotBoughtAndFreeReply, error)
	ReauthorizeApp(context.Context, *ReauthorizeAppRequest) (*ReauthorizeAppReply, error)
	InstallFreeTrialApp(context.Context, *InstallFreeTrialAppRequest) (*InstallFreeTrialAppReply, error)
	RecoverMerchantWebhookStatus(context.Context, *RecoverMerchantWebhookStatusRequest) (*RecoverMerchantWebhookStatusReply, error)
	mustEmbedUnimplementedMerchantServer()
}

// UnimplementedMerchantServer must be embedded to have forward compatible implementations.
type UnimplementedMerchantServer struct {
}

func (UnimplementedMerchantServer) PurchaseAppIfNotBoughtAndFree(context.Context, *MerchantPurchaseAppIfNotBoughtAndFreeRequest) (*MerchantPurchaseAppIfNotBoughtAndFreeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PurchaseAppIfNotBoughtAndFree not implemented")
}
func (UnimplementedMerchantServer) ReauthorizeApp(context.Context, *ReauthorizeAppRequest) (*ReauthorizeAppReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReauthorizeApp not implemented")
}
func (UnimplementedMerchantServer) InstallFreeTrialApp(context.Context, *InstallFreeTrialAppRequest) (*InstallFreeTrialAppReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InstallFreeTrialApp not implemented")
}
func (UnimplementedMerchantServer) RecoverMerchantWebhookStatus(context.Context, *RecoverMerchantWebhookStatusRequest) (*RecoverMerchantWebhookStatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecoverMerchantWebhookStatus not implemented")
}
func (UnimplementedMerchantServer) mustEmbedUnimplementedMerchantServer() {}

// UnsafeMerchantServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MerchantServer will
// result in compilation errors.
type UnsafeMerchantServer interface {
	mustEmbedUnimplementedMerchantServer()
}

func RegisterMerchantServer(s grpc.ServiceRegistrar, srv MerchantServer) {
	s.RegisterService(&Merchant_ServiceDesc, srv)
}

func _Merchant_PurchaseAppIfNotBoughtAndFree_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MerchantPurchaseAppIfNotBoughtAndFreeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MerchantServer).PurchaseAppIfNotBoughtAndFree(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Merchant_PurchaseAppIfNotBoughtAndFree_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MerchantServer).PurchaseAppIfNotBoughtAndFree(ctx, req.(*MerchantPurchaseAppIfNotBoughtAndFreeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Merchant_ReauthorizeApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReauthorizeAppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MerchantServer).ReauthorizeApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Merchant_ReauthorizeApp_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MerchantServer).ReauthorizeApp(ctx, req.(*ReauthorizeAppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Merchant_InstallFreeTrialApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InstallFreeTrialAppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MerchantServer).InstallFreeTrialApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Merchant_InstallFreeTrialApp_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MerchantServer).InstallFreeTrialApp(ctx, req.(*InstallFreeTrialAppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Merchant_RecoverMerchantWebhookStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecoverMerchantWebhookStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MerchantServer).RecoverMerchantWebhookStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Merchant_RecoverMerchantWebhookStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MerchantServer).RecoverMerchantWebhookStatus(ctx, req.(*RecoverMerchantWebhookStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Merchant_ServiceDesc is the grpc.ServiceDesc for Merchant service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Merchant_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "DeveloperApi.Merchant",
	HandlerType: (*MerchantServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PurchaseAppIfNotBoughtAndFree",
			Handler:    _Merchant_PurchaseAppIfNotBoughtAndFree_Handler,
		},
		{
			MethodName: "ReauthorizeApp",
			Handler:    _Merchant_ReauthorizeApp_Handler,
		},
		{
			MethodName: "InstallFreeTrialApp",
			Handler:    _Merchant_InstallFreeTrialApp_Handler,
		},
		{
			MethodName: "RecoverMerchantWebhookStatus",
			Handler:    _Merchant_RecoverMerchantWebhookStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "merchant.proto",
}