// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.3
// source: core/v1/auth.proto

package corev1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	AuthenticateService_PreLogin_FullMethodName = "/core.v1.AuthenticateService/preLogin"
	AuthenticateService_Login_FullMethodName    = "/core.v1.AuthenticateService/login"
	AuthenticateService_Logout_FullMethodName   = "/core.v1.AuthenticateService/logout"
)

// AuthenticateServiceClient is the client API for AuthenticateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthenticateServiceClient interface {
	PreLogin(ctx context.Context, in *User, opts ...grpc.CallOption) (*ActionReply, error)
	Login(ctx context.Context, in *User, opts ...grpc.CallOption) (*LoginReply, error)
	Logout(ctx context.Context, in *User, opts ...grpc.CallOption) (*ActionReply, error)
}

type authenticateServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthenticateServiceClient(cc grpc.ClientConnInterface) AuthenticateServiceClient {
	return &authenticateServiceClient{cc}
}

func (c *authenticateServiceClient) PreLogin(ctx context.Context, in *User, opts ...grpc.CallOption) (*ActionReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ActionReply)
	err := c.cc.Invoke(ctx, AuthenticateService_PreLogin_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticateServiceClient) Login(ctx context.Context, in *User, opts ...grpc.CallOption) (*LoginReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LoginReply)
	err := c.cc.Invoke(ctx, AuthenticateService_Login_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticateServiceClient) Logout(ctx context.Context, in *User, opts ...grpc.CallOption) (*ActionReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ActionReply)
	err := c.cc.Invoke(ctx, AuthenticateService_Logout_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthenticateServiceServer is the server API for AuthenticateService service.
// All implementations must embed UnimplementedAuthenticateServiceServer
// for forward compatibility.
type AuthenticateServiceServer interface {
	PreLogin(context.Context, *User) (*ActionReply, error)
	Login(context.Context, *User) (*LoginReply, error)
	Logout(context.Context, *User) (*ActionReply, error)
	mustEmbedUnimplementedAuthenticateServiceServer()
}

// UnimplementedAuthenticateServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAuthenticateServiceServer struct{}

func (UnimplementedAuthenticateServiceServer) PreLogin(context.Context, *User) (*ActionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreLogin not implemented")
}
func (UnimplementedAuthenticateServiceServer) Login(context.Context, *User) (*LoginReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedAuthenticateServiceServer) Logout(context.Context, *User) (*ActionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (UnimplementedAuthenticateServiceServer) mustEmbedUnimplementedAuthenticateServiceServer() {}
func (UnimplementedAuthenticateServiceServer) testEmbeddedByValue()                             {}

// UnsafeAuthenticateServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthenticateServiceServer will
// result in compilation errors.
type UnsafeAuthenticateServiceServer interface {
	mustEmbedUnimplementedAuthenticateServiceServer()
}

func RegisterAuthenticateServiceServer(s grpc.ServiceRegistrar, srv AuthenticateServiceServer) {
	// If the following call pancis, it indicates UnimplementedAuthenticateServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AuthenticateService_ServiceDesc, srv)
}

func _AuthenticateService_PreLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticateServiceServer).PreLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthenticateService_PreLogin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticateServiceServer).PreLogin(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthenticateService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticateServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthenticateService_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticateServiceServer).Login(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthenticateService_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticateServiceServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthenticateService_Logout_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticateServiceServer).Logout(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthenticateService_ServiceDesc is the grpc.ServiceDesc for AuthenticateService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthenticateService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "core.v1.AuthenticateService",
	HandlerType: (*AuthenticateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "preLogin",
			Handler:    _AuthenticateService_PreLogin_Handler,
		},
		{
			MethodName: "login",
			Handler:    _AuthenticateService_Login_Handler,
		},
		{
			MethodName: "logout",
			Handler:    _AuthenticateService_Logout_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "core/v1/auth.proto",
}
