// users.proto

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.2
// source: users.proto

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
	Users_Index_FullMethodName                               = "/DeveloperApi.Users/Index"
	Users_Show_FullMethodName                                = "/DeveloperApi.Users/Show"
	Users_ShowWithPassword_FullMethodName                    = "/DeveloperApi.Users/ShowWithPassword"
	Users_Create_FullMethodName                              = "/DeveloperApi.Users/Create"
	Users_Update_FullMethodName                              = "/DeveloperApi.Users/Update"
	Users_ForgotPassword_FullMethodName                      = "/DeveloperApi.Users/ForgotPassword"
	Users_VerifyForgotPasswordToken_FullMethodName           = "/DeveloperApi.Users/VerifyForgotPasswordToken"
	Users_ConfirmEmail_FullMethodName                        = "/DeveloperApi.Users/ConfirmEmail"
	Users_IsStaffAccessableToMerchantAppStore_FullMethodName = "/DeveloperApi.Users/IsStaffAccessableToMerchantAppStore"
)

// UsersClient is the client API for Users service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UsersClient interface {
	Index(ctx context.Context, in *UserIndexRequest, opts ...grpc.CallOption) (*UserIndexReply, error)
	Show(ctx context.Context, in *UserShowRequest, opts ...grpc.CallOption) (*User, error)
	ShowWithPassword(ctx context.Context, in *UserShowWithPasswordRequest, opts ...grpc.CallOption) (*UserWithPassword, error)
	Create(ctx context.Context, in *UserCreateRequest, opts ...grpc.CallOption) (*User, error)
	Update(ctx context.Context, in *UserUpdateRequest, opts ...grpc.CallOption) (*User, error)
	ForgotPassword(ctx context.Context, in *UserForgotPasswordRequest, opts ...grpc.CallOption) (*UserEmptyReply, error)
	VerifyForgotPasswordToken(ctx context.Context, in *UserVerifyForgotPasswordTokenRequest, opts ...grpc.CallOption) (*User, error)
	ConfirmEmail(ctx context.Context, in *ConfirmEmailRequest, opts ...grpc.CallOption) (*User, error)
	IsStaffAccessableToMerchantAppStore(ctx context.Context, in *IsStaffAccessableToMerchantAppStoreRequest, opts ...grpc.CallOption) (*IsStaffAccessableToMerchantAppStoreReply, error)
}

type usersClient struct {
	cc grpc.ClientConnInterface
}

func NewUsersClient(cc grpc.ClientConnInterface) UsersClient {
	return &usersClient{cc}
}

func (c *usersClient) Index(ctx context.Context, in *UserIndexRequest, opts ...grpc.CallOption) (*UserIndexReply, error) {
	out := new(UserIndexReply)
	err := c.cc.Invoke(ctx, Users_Index_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) Show(ctx context.Context, in *UserShowRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, Users_Show_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) ShowWithPassword(ctx context.Context, in *UserShowWithPasswordRequest, opts ...grpc.CallOption) (*UserWithPassword, error) {
	out := new(UserWithPassword)
	err := c.cc.Invoke(ctx, Users_ShowWithPassword_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) Create(ctx context.Context, in *UserCreateRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, Users_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) Update(ctx context.Context, in *UserUpdateRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, Users_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) ForgotPassword(ctx context.Context, in *UserForgotPasswordRequest, opts ...grpc.CallOption) (*UserEmptyReply, error) {
	out := new(UserEmptyReply)
	err := c.cc.Invoke(ctx, Users_ForgotPassword_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) VerifyForgotPasswordToken(ctx context.Context, in *UserVerifyForgotPasswordTokenRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, Users_VerifyForgotPasswordToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) ConfirmEmail(ctx context.Context, in *ConfirmEmailRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, Users_ConfirmEmail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) IsStaffAccessableToMerchantAppStore(ctx context.Context, in *IsStaffAccessableToMerchantAppStoreRequest, opts ...grpc.CallOption) (*IsStaffAccessableToMerchantAppStoreReply, error) {
	out := new(IsStaffAccessableToMerchantAppStoreReply)
	err := c.cc.Invoke(ctx, Users_IsStaffAccessableToMerchantAppStore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsersServer is the server API for Users service.
// All implementations must embed UnimplementedUsersServer
// for forward compatibility
type UsersServer interface {
	Index(context.Context, *UserIndexRequest) (*UserIndexReply, error)
	Show(context.Context, *UserShowRequest) (*User, error)
	ShowWithPassword(context.Context, *UserShowWithPasswordRequest) (*UserWithPassword, error)
	Create(context.Context, *UserCreateRequest) (*User, error)
	Update(context.Context, *UserUpdateRequest) (*User, error)
	ForgotPassword(context.Context, *UserForgotPasswordRequest) (*UserEmptyReply, error)
	VerifyForgotPasswordToken(context.Context, *UserVerifyForgotPasswordTokenRequest) (*User, error)
	ConfirmEmail(context.Context, *ConfirmEmailRequest) (*User, error)
	IsStaffAccessableToMerchantAppStore(context.Context, *IsStaffAccessableToMerchantAppStoreRequest) (*IsStaffAccessableToMerchantAppStoreReply, error)
	mustEmbedUnimplementedUsersServer()
}

// UnimplementedUsersServer must be embedded to have forward compatible implementations.
type UnimplementedUsersServer struct {
}

func (UnimplementedUsersServer) Index(context.Context, *UserIndexRequest) (*UserIndexReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Index not implemented")
}
func (UnimplementedUsersServer) Show(context.Context, *UserShowRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Show not implemented")
}
func (UnimplementedUsersServer) ShowWithPassword(context.Context, *UserShowWithPasswordRequest) (*UserWithPassword, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowWithPassword not implemented")
}
func (UnimplementedUsersServer) Create(context.Context, *UserCreateRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedUsersServer) Update(context.Context, *UserUpdateRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedUsersServer) ForgotPassword(context.Context, *UserForgotPasswordRequest) (*UserEmptyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ForgotPassword not implemented")
}
func (UnimplementedUsersServer) VerifyForgotPasswordToken(context.Context, *UserVerifyForgotPasswordTokenRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyForgotPasswordToken not implemented")
}
func (UnimplementedUsersServer) ConfirmEmail(context.Context, *ConfirmEmailRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmEmail not implemented")
}
func (UnimplementedUsersServer) IsStaffAccessableToMerchantAppStore(context.Context, *IsStaffAccessableToMerchantAppStoreRequest) (*IsStaffAccessableToMerchantAppStoreReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsStaffAccessableToMerchantAppStore not implemented")
}
func (UnimplementedUsersServer) mustEmbedUnimplementedUsersServer() {}

// UnsafeUsersServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UsersServer will
// result in compilation errors.
type UnsafeUsersServer interface {
	mustEmbedUnimplementedUsersServer()
}

func RegisterUsersServer(s grpc.ServiceRegistrar, srv UsersServer) {
	s.RegisterService(&Users_ServiceDesc, srv)
}

func _Users_Index_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserIndexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).Index(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Users_Index_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).Index(ctx, req.(*UserIndexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_Show_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserShowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).Show(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Users_Show_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).Show(ctx, req.(*UserShowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_ShowWithPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserShowWithPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).ShowWithPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Users_ShowWithPassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).ShowWithPassword(ctx, req.(*UserShowWithPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Users_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).Create(ctx, req.(*UserCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Users_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).Update(ctx, req.(*UserUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_ForgotPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserForgotPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).ForgotPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Users_ForgotPassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).ForgotPassword(ctx, req.(*UserForgotPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_VerifyForgotPasswordToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserVerifyForgotPasswordTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).VerifyForgotPasswordToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Users_VerifyForgotPasswordToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).VerifyForgotPasswordToken(ctx, req.(*UserVerifyForgotPasswordTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_ConfirmEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfirmEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).ConfirmEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Users_ConfirmEmail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).ConfirmEmail(ctx, req.(*ConfirmEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_IsStaffAccessableToMerchantAppStore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsStaffAccessableToMerchantAppStoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).IsStaffAccessableToMerchantAppStore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Users_IsStaffAccessableToMerchantAppStore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).IsStaffAccessableToMerchantAppStore(ctx, req.(*IsStaffAccessableToMerchantAppStoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Users_ServiceDesc is the grpc.ServiceDesc for Users service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Users_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "DeveloperApi.Users",
	HandlerType: (*UsersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Index",
			Handler:    _Users_Index_Handler,
		},
		{
			MethodName: "Show",
			Handler:    _Users_Show_Handler,
		},
		{
			MethodName: "ShowWithPassword",
			Handler:    _Users_ShowWithPassword_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Users_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Users_Update_Handler,
		},
		{
			MethodName: "ForgotPassword",
			Handler:    _Users_ForgotPassword_Handler,
		},
		{
			MethodName: "VerifyForgotPasswordToken",
			Handler:    _Users_VerifyForgotPasswordToken_Handler,
		},
		{
			MethodName: "ConfirmEmail",
			Handler:    _Users_ConfirmEmail_Handler,
		},
		{
			MethodName: "IsStaffAccessableToMerchantAppStore",
			Handler:    _Users_IsStaffAccessableToMerchantAppStore_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "users.proto",
}
