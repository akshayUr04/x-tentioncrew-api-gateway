// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: pkg/user/pb/user.proto

package pb

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

// UserSvcClient is the client API for UserSvc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserSvcClient interface {
	CreatUser(ctx context.Context, in *CreatUserRequest, opts ...grpc.CallOption) (*CommonResponse, error)
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*CommonResponse, error)
	GetUserById(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponce, error)
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*CommonResponse, error)
}

type userSvcClient struct {
	cc grpc.ClientConnInterface
}

func NewUserSvcClient(cc grpc.ClientConnInterface) UserSvcClient {
	return &userSvcClient{cc}
}

func (c *userSvcClient) CreatUser(ctx context.Context, in *CreatUserRequest, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, "/product.UserSvc/CreatUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userSvcClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, "/product.UserSvc/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userSvcClient) GetUserById(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponce, error) {
	out := new(GetUserResponce)
	err := c.cc.Invoke(ctx, "/product.UserSvc/GetUserById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userSvcClient) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, "/product.UserSvc/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserSvcServer is the server API for UserSvc service.
// All implementations must embed UnimplementedUserSvcServer
// for forward compatibility
type UserSvcServer interface {
	CreatUser(context.Context, *CreatUserRequest) (*CommonResponse, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*CommonResponse, error)
	GetUserById(context.Context, *GetUserRequest) (*GetUserResponce, error)
	DeleteUser(context.Context, *DeleteUserRequest) (*CommonResponse, error)
	mustEmbedUnimplementedUserSvcServer()
}

// UnimplementedUserSvcServer must be embedded to have forward compatible implementations.
type UnimplementedUserSvcServer struct {
}

func (UnimplementedUserSvcServer) CreatUser(context.Context, *CreatUserRequest) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatUser not implemented")
}
func (UnimplementedUserSvcServer) UpdateUser(context.Context, *UpdateUserRequest) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedUserSvcServer) GetUserById(context.Context, *GetUserRequest) (*GetUserResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserById not implemented")
}
func (UnimplementedUserSvcServer) DeleteUser(context.Context, *DeleteUserRequest) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedUserSvcServer) mustEmbedUnimplementedUserSvcServer() {}

// UnsafeUserSvcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserSvcServer will
// result in compilation errors.
type UnsafeUserSvcServer interface {
	mustEmbedUnimplementedUserSvcServer()
}

func RegisterUserSvcServer(s grpc.ServiceRegistrar, srv UserSvcServer) {
	s.RegisterService(&UserSvc_ServiceDesc, srv)
}

func _UserSvc_CreatUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserSvcServer).CreatUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.UserSvc/CreatUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserSvcServer).CreatUser(ctx, req.(*CreatUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserSvc_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserSvcServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.UserSvc/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserSvcServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserSvc_GetUserById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserSvcServer).GetUserById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.UserSvc/GetUserById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserSvcServer).GetUserById(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserSvc_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserSvcServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.UserSvc/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserSvcServer).DeleteUser(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserSvc_ServiceDesc is the grpc.ServiceDesc for UserSvc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserSvc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "product.UserSvc",
	HandlerType: (*UserSvcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatUser",
			Handler:    _UserSvc_CreatUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _UserSvc_UpdateUser_Handler,
		},
		{
			MethodName: "GetUserById",
			Handler:    _UserSvc_GetUserById_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _UserSvc_DeleteUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/user/pb/user.proto",
}
