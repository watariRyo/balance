// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: user_tag.proto

package proto

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

// UserTagServiceClient is the client API for UserTagService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserTagServiceClient interface {
	ListUserTag(ctx context.Context, in *UserTagListRequest, opts ...grpc.CallOption) (*UserTagRequest, error)
	GetUserTag(ctx context.Context, in *UserTagID, opts ...grpc.CallOption) (*UserTagListResponse, error)
	RegisterUserTag(ctx context.Context, in *UserTagRequest, opts ...grpc.CallOption) (*UserTagListResponse, error)
	UpdateUserTag(ctx context.Context, in *UserTagRequest, opts ...grpc.CallOption) (*UserTagListResponse, error)
	DeleteUserTag(ctx context.Context, in *UserTagID, opts ...grpc.CallOption) (*UserTagID, error)
}

type userTagServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserTagServiceClient(cc grpc.ClientConnInterface) UserTagServiceClient {
	return &userTagServiceClient{cc}
}

func (c *userTagServiceClient) ListUserTag(ctx context.Context, in *UserTagListRequest, opts ...grpc.CallOption) (*UserTagRequest, error) {
	out := new(UserTagRequest)
	err := c.cc.Invoke(ctx, "/server.UserTagService/ListUserTag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userTagServiceClient) GetUserTag(ctx context.Context, in *UserTagID, opts ...grpc.CallOption) (*UserTagListResponse, error) {
	out := new(UserTagListResponse)
	err := c.cc.Invoke(ctx, "/server.UserTagService/GetUserTag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userTagServiceClient) RegisterUserTag(ctx context.Context, in *UserTagRequest, opts ...grpc.CallOption) (*UserTagListResponse, error) {
	out := new(UserTagListResponse)
	err := c.cc.Invoke(ctx, "/server.UserTagService/RegisterUserTag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userTagServiceClient) UpdateUserTag(ctx context.Context, in *UserTagRequest, opts ...grpc.CallOption) (*UserTagListResponse, error) {
	out := new(UserTagListResponse)
	err := c.cc.Invoke(ctx, "/server.UserTagService/UpdateUserTag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userTagServiceClient) DeleteUserTag(ctx context.Context, in *UserTagID, opts ...grpc.CallOption) (*UserTagID, error) {
	out := new(UserTagID)
	err := c.cc.Invoke(ctx, "/server.UserTagService/DeleteUserTag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserTagServiceServer is the server API for UserTagService service.
// All implementations must embed UnimplementedUserTagServiceServer
// for forward compatibility
type UserTagServiceServer interface {
	ListUserTag(context.Context, *UserTagListRequest) (*UserTagRequest, error)
	GetUserTag(context.Context, *UserTagID) (*UserTagListResponse, error)
	RegisterUserTag(context.Context, *UserTagRequest) (*UserTagListResponse, error)
	UpdateUserTag(context.Context, *UserTagRequest) (*UserTagListResponse, error)
	DeleteUserTag(context.Context, *UserTagID) (*UserTagID, error)
	mustEmbedUnimplementedUserTagServiceServer()
}

// UnimplementedUserTagServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserTagServiceServer struct {
}

func (UnimplementedUserTagServiceServer) ListUserTag(context.Context, *UserTagListRequest) (*UserTagRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUserTag not implemented")
}
func (UnimplementedUserTagServiceServer) GetUserTag(context.Context, *UserTagID) (*UserTagListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserTag not implemented")
}
func (UnimplementedUserTagServiceServer) RegisterUserTag(context.Context, *UserTagRequest) (*UserTagListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterUserTag not implemented")
}
func (UnimplementedUserTagServiceServer) UpdateUserTag(context.Context, *UserTagRequest) (*UserTagListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserTag not implemented")
}
func (UnimplementedUserTagServiceServer) DeleteUserTag(context.Context, *UserTagID) (*UserTagID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUserTag not implemented")
}
func (UnimplementedUserTagServiceServer) mustEmbedUnimplementedUserTagServiceServer() {}

// UnsafeUserTagServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserTagServiceServer will
// result in compilation errors.
type UnsafeUserTagServiceServer interface {
	mustEmbedUnimplementedUserTagServiceServer()
}

func RegisterUserTagServiceServer(s grpc.ServiceRegistrar, srv UserTagServiceServer) {
	s.RegisterService(&UserTagService_ServiceDesc, srv)
}

func _UserTagService_ListUserTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserTagListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserTagServiceServer).ListUserTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.UserTagService/ListUserTag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserTagServiceServer).ListUserTag(ctx, req.(*UserTagListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserTagService_GetUserTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserTagID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserTagServiceServer).GetUserTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.UserTagService/GetUserTag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserTagServiceServer).GetUserTag(ctx, req.(*UserTagID))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserTagService_RegisterUserTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserTagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserTagServiceServer).RegisterUserTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.UserTagService/RegisterUserTag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserTagServiceServer).RegisterUserTag(ctx, req.(*UserTagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserTagService_UpdateUserTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserTagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserTagServiceServer).UpdateUserTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.UserTagService/UpdateUserTag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserTagServiceServer).UpdateUserTag(ctx, req.(*UserTagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserTagService_DeleteUserTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserTagID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserTagServiceServer).DeleteUserTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.UserTagService/DeleteUserTag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserTagServiceServer).DeleteUserTag(ctx, req.(*UserTagID))
	}
	return interceptor(ctx, in, info, handler)
}

// UserTagService_ServiceDesc is the grpc.ServiceDesc for UserTagService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserTagService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "server.UserTagService",
	HandlerType: (*UserTagServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListUserTag",
			Handler:    _UserTagService_ListUserTag_Handler,
		},
		{
			MethodName: "GetUserTag",
			Handler:    _UserTagService_GetUserTag_Handler,
		},
		{
			MethodName: "RegisterUserTag",
			Handler:    _UserTagService_RegisterUserTag_Handler,
		},
		{
			MethodName: "UpdateUserTag",
			Handler:    _UserTagService_UpdateUserTag_Handler,
		},
		{
			MethodName: "DeleteUserTag",
			Handler:    _UserTagService_DeleteUserTag_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user_tag.proto",
}
