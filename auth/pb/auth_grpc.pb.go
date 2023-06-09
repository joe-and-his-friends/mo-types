// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: auth/auth.proto

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

// AuthClient is the client API for Auth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthClient interface {
	AuthenticateUser(ctx context.Context, in *AuthenticateUserRequest, opts ...grpc.CallOption) (*AuthenticateUserResponse, error)
	RefreshAccessToken(ctx context.Context, in *RefreshAccessTokenRequest, opts ...grpc.CallOption) (*RefreshAccessTokenResponse, error)
	AuthenticateUserFromGoogle(ctx context.Context, in *AuthenticateUserFromGoogleRequest, opts ...grpc.CallOption) (*AuthenticateUserFromGoogleResponse, error)
	GetGoogleUserId(ctx context.Context, in *GetGoogleUserIdRequest, opts ...grpc.CallOption) (*GetGoogleUserIdResponse, error)
	AuthenticateUserFromApple(ctx context.Context, in *AuthenticateUserFromAppleRequest, opts ...grpc.CallOption) (*AuthenticateUserFromAppleResponse, error)
	GetAppleUserId(ctx context.Context, in *GetAppleUserIdRequest, opts ...grpc.CallOption) (*GetAppleUserIdResponse, error)
}

type authClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthClient(cc grpc.ClientConnInterface) AuthClient {
	return &authClient{cc}
}

func (c *authClient) AuthenticateUser(ctx context.Context, in *AuthenticateUserRequest, opts ...grpc.CallOption) (*AuthenticateUserResponse, error) {
	out := new(AuthenticateUserResponse)
	err := c.cc.Invoke(ctx, "/auth.Auth/AuthenticateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) RefreshAccessToken(ctx context.Context, in *RefreshAccessTokenRequest, opts ...grpc.CallOption) (*RefreshAccessTokenResponse, error) {
	out := new(RefreshAccessTokenResponse)
	err := c.cc.Invoke(ctx, "/auth.Auth/RefreshAccessToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) AuthenticateUserFromGoogle(ctx context.Context, in *AuthenticateUserFromGoogleRequest, opts ...grpc.CallOption) (*AuthenticateUserFromGoogleResponse, error) {
	out := new(AuthenticateUserFromGoogleResponse)
	err := c.cc.Invoke(ctx, "/auth.Auth/AuthenticateUserFromGoogle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) GetGoogleUserId(ctx context.Context, in *GetGoogleUserIdRequest, opts ...grpc.CallOption) (*GetGoogleUserIdResponse, error) {
	out := new(GetGoogleUserIdResponse)
	err := c.cc.Invoke(ctx, "/auth.Auth/GetGoogleUserId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) AuthenticateUserFromApple(ctx context.Context, in *AuthenticateUserFromAppleRequest, opts ...grpc.CallOption) (*AuthenticateUserFromAppleResponse, error) {
	out := new(AuthenticateUserFromAppleResponse)
	err := c.cc.Invoke(ctx, "/auth.Auth/AuthenticateUserFromApple", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) GetAppleUserId(ctx context.Context, in *GetAppleUserIdRequest, opts ...grpc.CallOption) (*GetAppleUserIdResponse, error) {
	out := new(GetAppleUserIdResponse)
	err := c.cc.Invoke(ctx, "/auth.Auth/GetAppleUserId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServer is the server API for Auth service.
// All implementations must embed UnimplementedAuthServer
// for forward compatibility
type AuthServer interface {
	AuthenticateUser(context.Context, *AuthenticateUserRequest) (*AuthenticateUserResponse, error)
	RefreshAccessToken(context.Context, *RefreshAccessTokenRequest) (*RefreshAccessTokenResponse, error)
	AuthenticateUserFromGoogle(context.Context, *AuthenticateUserFromGoogleRequest) (*AuthenticateUserFromGoogleResponse, error)
	GetGoogleUserId(context.Context, *GetGoogleUserIdRequest) (*GetGoogleUserIdResponse, error)
	AuthenticateUserFromApple(context.Context, *AuthenticateUserFromAppleRequest) (*AuthenticateUserFromAppleResponse, error)
	GetAppleUserId(context.Context, *GetAppleUserIdRequest) (*GetAppleUserIdResponse, error)
	mustEmbedUnimplementedAuthServer()
}

// UnimplementedAuthServer must be embedded to have forward compatible implementations.
type UnimplementedAuthServer struct {
}

func (UnimplementedAuthServer) AuthenticateUser(context.Context, *AuthenticateUserRequest) (*AuthenticateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthenticateUser not implemented")
}
func (UnimplementedAuthServer) RefreshAccessToken(context.Context, *RefreshAccessTokenRequest) (*RefreshAccessTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshAccessToken not implemented")
}
func (UnimplementedAuthServer) AuthenticateUserFromGoogle(context.Context, *AuthenticateUserFromGoogleRequest) (*AuthenticateUserFromGoogleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthenticateUserFromGoogle not implemented")
}
func (UnimplementedAuthServer) GetGoogleUserId(context.Context, *GetGoogleUserIdRequest) (*GetGoogleUserIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGoogleUserId not implemented")
}
func (UnimplementedAuthServer) AuthenticateUserFromApple(context.Context, *AuthenticateUserFromAppleRequest) (*AuthenticateUserFromAppleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthenticateUserFromApple not implemented")
}
func (UnimplementedAuthServer) GetAppleUserId(context.Context, *GetAppleUserIdRequest) (*GetAppleUserIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppleUserId not implemented")
}
func (UnimplementedAuthServer) mustEmbedUnimplementedAuthServer() {}

// UnsafeAuthServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServer will
// result in compilation errors.
type UnsafeAuthServer interface {
	mustEmbedUnimplementedAuthServer()
}

func RegisterAuthServer(s grpc.ServiceRegistrar, srv AuthServer) {
	s.RegisterService(&Auth_ServiceDesc, srv)
}

func _Auth_AuthenticateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthenticateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).AuthenticateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/AuthenticateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).AuthenticateUser(ctx, req.(*AuthenticateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_RefreshAccessToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshAccessTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).RefreshAccessToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/RefreshAccessToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).RefreshAccessToken(ctx, req.(*RefreshAccessTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_AuthenticateUserFromGoogle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthenticateUserFromGoogleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).AuthenticateUserFromGoogle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/AuthenticateUserFromGoogle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).AuthenticateUserFromGoogle(ctx, req.(*AuthenticateUserFromGoogleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_GetGoogleUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGoogleUserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).GetGoogleUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/GetGoogleUserId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).GetGoogleUserId(ctx, req.(*GetGoogleUserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_AuthenticateUserFromApple_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthenticateUserFromAppleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).AuthenticateUserFromApple(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/AuthenticateUserFromApple",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).AuthenticateUserFromApple(ctx, req.(*AuthenticateUserFromAppleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_GetAppleUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppleUserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).GetAppleUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/GetAppleUserId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).GetAppleUserId(ctx, req.(*GetAppleUserIdRequest))
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
			MethodName: "AuthenticateUser",
			Handler:    _Auth_AuthenticateUser_Handler,
		},
		{
			MethodName: "RefreshAccessToken",
			Handler:    _Auth_RefreshAccessToken_Handler,
		},
		{
			MethodName: "AuthenticateUserFromGoogle",
			Handler:    _Auth_AuthenticateUserFromGoogle_Handler,
		},
		{
			MethodName: "GetGoogleUserId",
			Handler:    _Auth_GetGoogleUserId_Handler,
		},
		{
			MethodName: "AuthenticateUserFromApple",
			Handler:    _Auth_AuthenticateUserFromApple_Handler,
		},
		{
			MethodName: "GetAppleUserId",
			Handler:    _Auth_GetAppleUserId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth/auth.proto",
}
