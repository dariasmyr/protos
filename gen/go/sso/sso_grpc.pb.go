// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: sso/sso.proto

package gen

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Auth_RegisterClient_FullMethodName = "/auth.Auth/RegisterClient"
	Auth_Register_FullMethodName       = "/auth.Auth/Register"
	Auth_Login_FullMethodName          = "/auth.Auth/Login"
	Auth_Logout_FullMethodName         = "/auth.Auth/Logout"
	Auth_ChangePassword_FullMethodName = "/auth.Auth/ChangePassword"
	Auth_ChangeStatus_FullMethodName   = "/auth.Auth/ChangeStatus"
)

// AuthClient is the client API for Auth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthClient interface {
	RegisterClient(ctx context.Context, in *RegisterClientRequest, opts ...grpc.CallOption) (*RegisterClientResponse, error)
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	Logout(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*LogoutResponse, error)
	ChangePassword(ctx context.Context, in *ChangePasswordRequest, opts ...grpc.CallOption) (*ChangePasswordResponse, error)
	ChangeStatus(ctx context.Context, in *ChangeStatusRequest, opts ...grpc.CallOption) (*ChangeStatusResponse, error)
}

type authClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthClient(cc grpc.ClientConnInterface) AuthClient {
	return &authClient{cc}
}

func (c *authClient) RegisterClient(ctx context.Context, in *RegisterClientRequest, opts ...grpc.CallOption) (*RegisterClientResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RegisterClientResponse)
	err := c.cc.Invoke(ctx, Auth_RegisterClient_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, Auth_Register_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, Auth_Login_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) Logout(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*LogoutResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LogoutResponse)
	err := c.cc.Invoke(ctx, Auth_Logout_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) ChangePassword(ctx context.Context, in *ChangePasswordRequest, opts ...grpc.CallOption) (*ChangePasswordResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ChangePasswordResponse)
	err := c.cc.Invoke(ctx, Auth_ChangePassword_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) ChangeStatus(ctx context.Context, in *ChangeStatusRequest, opts ...grpc.CallOption) (*ChangeStatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ChangeStatusResponse)
	err := c.cc.Invoke(ctx, Auth_ChangeStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServer is the server API for Auth service.
// All implementations should embed UnimplementedAuthServer
// for forward compatibility.
type AuthServer interface {
	RegisterClient(context.Context, *RegisterClientRequest) (*RegisterClientResponse, error)
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	Logout(context.Context, *empty.Empty) (*LogoutResponse, error)
	ChangePassword(context.Context, *ChangePasswordRequest) (*ChangePasswordResponse, error)
	ChangeStatus(context.Context, *ChangeStatusRequest) (*ChangeStatusResponse, error)
}

// UnimplementedAuthServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAuthServer struct{}

func (UnimplementedAuthServer) RegisterClient(context.Context, *RegisterClientRequest) (*RegisterClientResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterClient not implemented")
}
func (UnimplementedAuthServer) Register(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedAuthServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedAuthServer) Logout(context.Context, *empty.Empty) (*LogoutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (UnimplementedAuthServer) ChangePassword(context.Context, *ChangePasswordRequest) (*ChangePasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePassword not implemented")
}
func (UnimplementedAuthServer) ChangeStatus(context.Context, *ChangeStatusRequest) (*ChangeStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeStatus not implemented")
}
func (UnimplementedAuthServer) testEmbeddedByValue() {}

// UnsafeAuthServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServer will
// result in compilation errors.
type UnsafeAuthServer interface {
	mustEmbedUnimplementedAuthServer()
}

func RegisterAuthServer(s grpc.ServiceRegistrar, srv AuthServer) {
	// If the following call pancis, it indicates UnimplementedAuthServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Auth_ServiceDesc, srv)
}

func _Auth_RegisterClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).RegisterClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_RegisterClient_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).RegisterClient(ctx, req.(*RegisterClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_Logout_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).Logout(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_ChangePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangePasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).ChangePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_ChangePassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).ChangePassword(ctx, req.(*ChangePasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_ChangeStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).ChangeStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_ChangeStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).ChangeStatus(ctx, req.(*ChangeStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Auth_ServiceDesc is the grpc.ServiceDesc for Auth service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Auth_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterClient",
			Handler:    _Auth_RegisterClient_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _Auth_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _Auth_Login_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _Auth_Logout_Handler,
		},
		{
			MethodName: "ChangePassword",
			Handler:    _Auth_ChangePassword_Handler,
		},
		{
			MethodName: "ChangeStatus",
			Handler:    _Auth_ChangeStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sso/sso.proto",
}

const (
	Sessions_GetActiveSessions_FullMethodName = "/auth.Sessions/GetActiveSessions"
	Sessions_RefreshSession_FullMethodName    = "/auth.Sessions/RefreshSession"
	Sessions_ValidateSession_FullMethodName   = "/auth.Sessions/ValidateSession"
	Sessions_RevokeSession_FullMethodName     = "/auth.Sessions/RevokeSession"
)

// SessionsClient is the client API for Sessions service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SessionsClient interface {
	GetActiveSessions(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetActiveAccountSessionsResponse, error)
	RefreshSession(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*RefreshAccountSessionResponse, error)
	ValidateSession(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ValidateAccountSessionResponse, error)
	RevokeSession(ctx context.Context, in *RevokeAccountSessionRequest, opts ...grpc.CallOption) (*RevokeAccountSessionResponse, error)
}

type sessionsClient struct {
	cc grpc.ClientConnInterface
}

func NewSessionsClient(cc grpc.ClientConnInterface) SessionsClient {
	return &sessionsClient{cc}
}

func (c *sessionsClient) GetActiveSessions(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetActiveAccountSessionsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetActiveAccountSessionsResponse)
	err := c.cc.Invoke(ctx, Sessions_GetActiveSessions_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionsClient) RefreshSession(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*RefreshAccountSessionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RefreshAccountSessionResponse)
	err := c.cc.Invoke(ctx, Sessions_RefreshSession_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionsClient) ValidateSession(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ValidateAccountSessionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ValidateAccountSessionResponse)
	err := c.cc.Invoke(ctx, Sessions_ValidateSession_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionsClient) RevokeSession(ctx context.Context, in *RevokeAccountSessionRequest, opts ...grpc.CallOption) (*RevokeAccountSessionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RevokeAccountSessionResponse)
	err := c.cc.Invoke(ctx, Sessions_RevokeSession_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SessionsServer is the server API for Sessions service.
// All implementations should embed UnimplementedSessionsServer
// for forward compatibility.
type SessionsServer interface {
	GetActiveSessions(context.Context, *empty.Empty) (*GetActiveAccountSessionsResponse, error)
	RefreshSession(context.Context, *empty.Empty) (*RefreshAccountSessionResponse, error)
	ValidateSession(context.Context, *empty.Empty) (*ValidateAccountSessionResponse, error)
	RevokeSession(context.Context, *RevokeAccountSessionRequest) (*RevokeAccountSessionResponse, error)
}

// UnimplementedSessionsServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSessionsServer struct{}

func (UnimplementedSessionsServer) GetActiveSessions(context.Context, *empty.Empty) (*GetActiveAccountSessionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetActiveSessions not implemented")
}
func (UnimplementedSessionsServer) RefreshSession(context.Context, *empty.Empty) (*RefreshAccountSessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshSession not implemented")
}
func (UnimplementedSessionsServer) ValidateSession(context.Context, *empty.Empty) (*ValidateAccountSessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateSession not implemented")
}
func (UnimplementedSessionsServer) RevokeSession(context.Context, *RevokeAccountSessionRequest) (*RevokeAccountSessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RevokeSession not implemented")
}
func (UnimplementedSessionsServer) testEmbeddedByValue() {}

// UnsafeSessionsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SessionsServer will
// result in compilation errors.
type UnsafeSessionsServer interface {
	mustEmbedUnimplementedSessionsServer()
}

func RegisterSessionsServer(s grpc.ServiceRegistrar, srv SessionsServer) {
	// If the following call pancis, it indicates UnimplementedSessionsServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Sessions_ServiceDesc, srv)
}

func _Sessions_GetActiveSessions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionsServer).GetActiveSessions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sessions_GetActiveSessions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionsServer).GetActiveSessions(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sessions_RefreshSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionsServer).RefreshSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sessions_RefreshSession_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionsServer).RefreshSession(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sessions_ValidateSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionsServer).ValidateSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sessions_ValidateSession_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionsServer).ValidateSession(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sessions_RevokeSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RevokeAccountSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionsServer).RevokeSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sessions_RevokeSession_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionsServer).RevokeSession(ctx, req.(*RevokeAccountSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Sessions_ServiceDesc is the grpc.ServiceDesc for Sessions service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Sessions_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.Sessions",
	HandlerType: (*SessionsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetActiveSessions",
			Handler:    _Sessions_GetActiveSessions_Handler,
		},
		{
			MethodName: "RefreshSession",
			Handler:    _Sessions_RefreshSession_Handler,
		},
		{
			MethodName: "ValidateSession",
			Handler:    _Sessions_ValidateSession_Handler,
		},
		{
			MethodName: "RevokeSession",
			Handler:    _Sessions_RevokeSession_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sso/sso.proto",
}
