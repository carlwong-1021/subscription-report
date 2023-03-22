// applications.proto

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.2
// source: application_reviews.proto

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
	ApplicationReviews_Index_FullMethodName  = "/DeveloperApi.ApplicationReviews/Index"
	ApplicationReviews_Upsert_FullMethodName = "/DeveloperApi.ApplicationReviews/Upsert"
)

// ApplicationReviewsClient is the client API for ApplicationReviews service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ApplicationReviewsClient interface {
	Index(ctx context.Context, in *IndexRequest, opts ...grpc.CallOption) (*IndexResponse, error)
	Upsert(ctx context.Context, in *UpsertRequest, opts ...grpc.CallOption) (*Review, error)
}

type applicationReviewsClient struct {
	cc grpc.ClientConnInterface
}

func NewApplicationReviewsClient(cc grpc.ClientConnInterface) ApplicationReviewsClient {
	return &applicationReviewsClient{cc}
}

func (c *applicationReviewsClient) Index(ctx context.Context, in *IndexRequest, opts ...grpc.CallOption) (*IndexResponse, error) {
	out := new(IndexResponse)
	err := c.cc.Invoke(ctx, ApplicationReviews_Index_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationReviewsClient) Upsert(ctx context.Context, in *UpsertRequest, opts ...grpc.CallOption) (*Review, error) {
	out := new(Review)
	err := c.cc.Invoke(ctx, ApplicationReviews_Upsert_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApplicationReviewsServer is the server API for ApplicationReviews service.
// All implementations must embed UnimplementedApplicationReviewsServer
// for forward compatibility
type ApplicationReviewsServer interface {
	Index(context.Context, *IndexRequest) (*IndexResponse, error)
	Upsert(context.Context, *UpsertRequest) (*Review, error)
	mustEmbedUnimplementedApplicationReviewsServer()
}

// UnimplementedApplicationReviewsServer must be embedded to have forward compatible implementations.
type UnimplementedApplicationReviewsServer struct {
}

func (UnimplementedApplicationReviewsServer) Index(context.Context, *IndexRequest) (*IndexResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Index not implemented")
}
func (UnimplementedApplicationReviewsServer) Upsert(context.Context, *UpsertRequest) (*Review, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Upsert not implemented")
}
func (UnimplementedApplicationReviewsServer) mustEmbedUnimplementedApplicationReviewsServer() {}

// UnsafeApplicationReviewsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ApplicationReviewsServer will
// result in compilation errors.
type UnsafeApplicationReviewsServer interface {
	mustEmbedUnimplementedApplicationReviewsServer()
}

func RegisterApplicationReviewsServer(s grpc.ServiceRegistrar, srv ApplicationReviewsServer) {
	s.RegisterService(&ApplicationReviews_ServiceDesc, srv)
}

func _ApplicationReviews_Index_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IndexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationReviewsServer).Index(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApplicationReviews_Index_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationReviewsServer).Index(ctx, req.(*IndexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApplicationReviews_Upsert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationReviewsServer).Upsert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApplicationReviews_Upsert_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationReviewsServer).Upsert(ctx, req.(*UpsertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ApplicationReviews_ServiceDesc is the grpc.ServiceDesc for ApplicationReviews service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ApplicationReviews_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "DeveloperApi.ApplicationReviews",
	HandlerType: (*ApplicationReviewsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Index",
			Handler:    _ApplicationReviews_Index_Handler,
		},
		{
			MethodName: "Upsert",
			Handler:    _ApplicationReviews_Upsert_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "application_reviews.proto",
}