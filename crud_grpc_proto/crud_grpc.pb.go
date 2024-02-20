// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: crud_grpc_proto/crud.proto

package grpc_crud

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
	Crud_CreateUser_FullMethodName    = "/crud_grpc.crud/CreateUser"
	Crud_Listuser_FullMethodName      = "/crud_grpc.crud/Listuser"
	Crud_GetUserByName_FullMethodName = "/crud_grpc.crud/GetUserByName"
	Crud_DeleteUser_FullMethodName    = "/crud_grpc.crud/DeleteUser"
)

// CrudClient is the client API for Crud service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CrudClient interface {
	CreateUser(ctx context.Context, in *UserDetails, opts ...grpc.CallOption) (*Response, error)
	Listuser(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListUserResponse, error)
	GetUserByName(ctx context.Context, in *Name, opts ...grpc.CallOption) (*UserDetails, error)
	// rpc UpdateUser(Name)returns(Response);
	DeleteUser(ctx context.Context, in *Name, opts ...grpc.CallOption) (*Response, error)
}

type crudClient struct {
	cc grpc.ClientConnInterface
}

func NewCrudClient(cc grpc.ClientConnInterface) CrudClient {
	return &crudClient{cc}
}

func (c *crudClient) CreateUser(ctx context.Context, in *UserDetails, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, Crud_CreateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crudClient) Listuser(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListUserResponse, error) {
	out := new(ListUserResponse)
	err := c.cc.Invoke(ctx, Crud_Listuser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crudClient) GetUserByName(ctx context.Context, in *Name, opts ...grpc.CallOption) (*UserDetails, error) {
	out := new(UserDetails)
	err := c.cc.Invoke(ctx, Crud_GetUserByName_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crudClient) DeleteUser(ctx context.Context, in *Name, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, Crud_DeleteUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CrudServer is the server API for Crud service.
// All implementations must embed UnimplementedCrudServer
// for forward compatibility
type CrudServer interface {
	CreateUser(context.Context, *UserDetails) (*Response, error)
	Listuser(context.Context, *Empty) (*ListUserResponse, error)
	GetUserByName(context.Context, *Name) (*UserDetails, error)
	// rpc UpdateUser(Name)returns(Response);
	DeleteUser(context.Context, *Name) (*Response, error)
	mustEmbedUnimplementedCrudServer()
}

// UnimplementedCrudServer must be embedded to have forward compatible implementations.
type UnimplementedCrudServer struct {
}

func (UnimplementedCrudServer) CreateUser(context.Context, *UserDetails) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedCrudServer) Listuser(context.Context, *Empty) (*ListUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Listuser not implemented")
}
func (UnimplementedCrudServer) GetUserByName(context.Context, *Name) (*UserDetails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByName not implemented")
}
func (UnimplementedCrudServer) DeleteUser(context.Context, *Name) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedCrudServer) mustEmbedUnimplementedCrudServer() {}

// UnsafeCrudServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CrudServer will
// result in compilation errors.
type UnsafeCrudServer interface {
	mustEmbedUnimplementedCrudServer()
}

func RegisterCrudServer(s grpc.ServiceRegistrar, srv CrudServer) {
	s.RegisterService(&Crud_ServiceDesc, srv)
}

func _Crud_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserDetails)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrudServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Crud_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrudServer).CreateUser(ctx, req.(*UserDetails))
	}
	return interceptor(ctx, in, info, handler)
}

func _Crud_Listuser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrudServer).Listuser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Crud_Listuser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrudServer).Listuser(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Crud_GetUserByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Name)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrudServer).GetUserByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Crud_GetUserByName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrudServer).GetUserByName(ctx, req.(*Name))
	}
	return interceptor(ctx, in, info, handler)
}

func _Crud_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Name)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrudServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Crud_DeleteUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrudServer).DeleteUser(ctx, req.(*Name))
	}
	return interceptor(ctx, in, info, handler)
}

// Crud_ServiceDesc is the grpc.ServiceDesc for Crud service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Crud_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "crud_grpc.crud",
	HandlerType: (*CrudServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _Crud_CreateUser_Handler,
		},
		{
			MethodName: "Listuser",
			Handler:    _Crud_Listuser_Handler,
		},
		{
			MethodName: "GetUserByName",
			Handler:    _Crud_GetUserByName_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _Crud_DeleteUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "crud_grpc_proto/crud.proto",
}
