// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: moment/moment.proto

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

// MomentServiceClient is the client API for MomentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MomentServiceClient interface {
	PostMoment(ctx context.Context, in *PostMomentRequest, opts ...grpc.CallOption) (*PostMomentResponse, error)
	GetMoment(ctx context.Context, in *GetMomentRequest, opts ...grpc.CallOption) (*GetMomentResponse, error)
	GetMoments(ctx context.Context, in *GetMomentsRequest, opts ...grpc.CallOption) (*GetMomentsResponse, error)
	GetMomentsReported(ctx context.Context, in *GetMomentsReportedRequest, opts ...grpc.CallOption) (*GetMomentsReportedResponse, error)
	GetRetailersBlocked(ctx context.Context, in *GetRetailersBlockedRequest, opts ...grpc.CallOption) (*GetRetailersBlockedResponse, error)
	GetRetailersNearby(ctx context.Context, in *GetRetailersNearbyRequest, opts ...grpc.CallOption) (*GetRetailersNearbyResponse, error)
	PutMoment(ctx context.Context, in *PutMomentRequest, opts ...grpc.CallOption) (*PutMomentResponse, error)
	DeleteMoment(ctx context.Context, in *DeleteMomentRequest, opts ...grpc.CallOption) (*DeleteMomentResponse, error)
	PostMomentReported(ctx context.Context, in *PostMomentReportedRequest, opts ...grpc.CallOption) (*PostMomentReportedResponse, error)
	DeleteMomentReported(ctx context.Context, in *DeleteMomentReportedRequest, opts ...grpc.CallOption) (*DeleteMomentReportedResponse, error)
	PostRetailerBlocked(ctx context.Context, in *PostRetailerBlockedRequest, opts ...grpc.CallOption) (*PostRetailerBlockedResponse, error)
	DeleteRetailerBlocked(ctx context.Context, in *DeleteRetailerBlockedRequest, opts ...grpc.CallOption) (*DeleteRetailerBlockedResponse, error)
}

type momentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMomentServiceClient(cc grpc.ClientConnInterface) MomentServiceClient {
	return &momentServiceClient{cc}
}

func (c *momentServiceClient) PostMoment(ctx context.Context, in *PostMomentRequest, opts ...grpc.CallOption) (*PostMomentResponse, error) {
	out := new(PostMomentResponse)
	err := c.cc.Invoke(ctx, "/moment.MomentService/PostMoment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *momentServiceClient) GetMoment(ctx context.Context, in *GetMomentRequest, opts ...grpc.CallOption) (*GetMomentResponse, error) {
	out := new(GetMomentResponse)
	err := c.cc.Invoke(ctx, "/moment.MomentService/GetMoment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *momentServiceClient) GetMoments(ctx context.Context, in *GetMomentsRequest, opts ...grpc.CallOption) (*GetMomentsResponse, error) {
	out := new(GetMomentsResponse)
	err := c.cc.Invoke(ctx, "/moment.MomentService/GetMoments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *momentServiceClient) GetMomentsReported(ctx context.Context, in *GetMomentsReportedRequest, opts ...grpc.CallOption) (*GetMomentsReportedResponse, error) {
	out := new(GetMomentsReportedResponse)
	err := c.cc.Invoke(ctx, "/moment.MomentService/GetMomentsReported", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *momentServiceClient) GetRetailersBlocked(ctx context.Context, in *GetRetailersBlockedRequest, opts ...grpc.CallOption) (*GetRetailersBlockedResponse, error) {
	out := new(GetRetailersBlockedResponse)
	err := c.cc.Invoke(ctx, "/moment.MomentService/GetRetailersBlocked", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *momentServiceClient) GetRetailersNearby(ctx context.Context, in *GetRetailersNearbyRequest, opts ...grpc.CallOption) (*GetRetailersNearbyResponse, error) {
	out := new(GetRetailersNearbyResponse)
	err := c.cc.Invoke(ctx, "/moment.MomentService/GetRetailersNearby", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *momentServiceClient) PutMoment(ctx context.Context, in *PutMomentRequest, opts ...grpc.CallOption) (*PutMomentResponse, error) {
	out := new(PutMomentResponse)
	err := c.cc.Invoke(ctx, "/moment.MomentService/PutMoment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *momentServiceClient) DeleteMoment(ctx context.Context, in *DeleteMomentRequest, opts ...grpc.CallOption) (*DeleteMomentResponse, error) {
	out := new(DeleteMomentResponse)
	err := c.cc.Invoke(ctx, "/moment.MomentService/DeleteMoment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *momentServiceClient) PostMomentReported(ctx context.Context, in *PostMomentReportedRequest, opts ...grpc.CallOption) (*PostMomentReportedResponse, error) {
	out := new(PostMomentReportedResponse)
	err := c.cc.Invoke(ctx, "/moment.MomentService/PostMomentReported", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *momentServiceClient) DeleteMomentReported(ctx context.Context, in *DeleteMomentReportedRequest, opts ...grpc.CallOption) (*DeleteMomentReportedResponse, error) {
	out := new(DeleteMomentReportedResponse)
	err := c.cc.Invoke(ctx, "/moment.MomentService/DeleteMomentReported", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *momentServiceClient) PostRetailerBlocked(ctx context.Context, in *PostRetailerBlockedRequest, opts ...grpc.CallOption) (*PostRetailerBlockedResponse, error) {
	out := new(PostRetailerBlockedResponse)
	err := c.cc.Invoke(ctx, "/moment.MomentService/PostRetailerBlocked", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *momentServiceClient) DeleteRetailerBlocked(ctx context.Context, in *DeleteRetailerBlockedRequest, opts ...grpc.CallOption) (*DeleteRetailerBlockedResponse, error) {
	out := new(DeleteRetailerBlockedResponse)
	err := c.cc.Invoke(ctx, "/moment.MomentService/DeleteRetailerBlocked", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MomentServiceServer is the server API for MomentService service.
// All implementations must embed UnimplementedMomentServiceServer
// for forward compatibility
type MomentServiceServer interface {
	PostMoment(context.Context, *PostMomentRequest) (*PostMomentResponse, error)
	GetMoment(context.Context, *GetMomentRequest) (*GetMomentResponse, error)
	GetMoments(context.Context, *GetMomentsRequest) (*GetMomentsResponse, error)
	GetMomentsReported(context.Context, *GetMomentsReportedRequest) (*GetMomentsReportedResponse, error)
	GetRetailersBlocked(context.Context, *GetRetailersBlockedRequest) (*GetRetailersBlockedResponse, error)
	GetRetailersNearby(context.Context, *GetRetailersNearbyRequest) (*GetRetailersNearbyResponse, error)
	PutMoment(context.Context, *PutMomentRequest) (*PutMomentResponse, error)
	DeleteMoment(context.Context, *DeleteMomentRequest) (*DeleteMomentResponse, error)
	PostMomentReported(context.Context, *PostMomentReportedRequest) (*PostMomentReportedResponse, error)
	DeleteMomentReported(context.Context, *DeleteMomentReportedRequest) (*DeleteMomentReportedResponse, error)
	PostRetailerBlocked(context.Context, *PostRetailerBlockedRequest) (*PostRetailerBlockedResponse, error)
	DeleteRetailerBlocked(context.Context, *DeleteRetailerBlockedRequest) (*DeleteRetailerBlockedResponse, error)
	mustEmbedUnimplementedMomentServiceServer()
}

// UnimplementedMomentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMomentServiceServer struct {
}

func (UnimplementedMomentServiceServer) PostMoment(context.Context, *PostMomentRequest) (*PostMomentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostMoment not implemented")
}
func (UnimplementedMomentServiceServer) GetMoment(context.Context, *GetMomentRequest) (*GetMomentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMoment not implemented")
}
func (UnimplementedMomentServiceServer) GetMoments(context.Context, *GetMomentsRequest) (*GetMomentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMoments not implemented")
}
func (UnimplementedMomentServiceServer) GetMomentsReported(context.Context, *GetMomentsReportedRequest) (*GetMomentsReportedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMomentsReported not implemented")
}
func (UnimplementedMomentServiceServer) GetRetailersBlocked(context.Context, *GetRetailersBlockedRequest) (*GetRetailersBlockedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRetailersBlocked not implemented")
}
func (UnimplementedMomentServiceServer) GetRetailersNearby(context.Context, *GetRetailersNearbyRequest) (*GetRetailersNearbyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRetailersNearby not implemented")
}
func (UnimplementedMomentServiceServer) PutMoment(context.Context, *PutMomentRequest) (*PutMomentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutMoment not implemented")
}
func (UnimplementedMomentServiceServer) DeleteMoment(context.Context, *DeleteMomentRequest) (*DeleteMomentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMoment not implemented")
}
func (UnimplementedMomentServiceServer) PostMomentReported(context.Context, *PostMomentReportedRequest) (*PostMomentReportedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostMomentReported not implemented")
}
func (UnimplementedMomentServiceServer) DeleteMomentReported(context.Context, *DeleteMomentReportedRequest) (*DeleteMomentReportedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMomentReported not implemented")
}
func (UnimplementedMomentServiceServer) PostRetailerBlocked(context.Context, *PostRetailerBlockedRequest) (*PostRetailerBlockedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostRetailerBlocked not implemented")
}
func (UnimplementedMomentServiceServer) DeleteRetailerBlocked(context.Context, *DeleteRetailerBlockedRequest) (*DeleteRetailerBlockedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRetailerBlocked not implemented")
}
func (UnimplementedMomentServiceServer) mustEmbedUnimplementedMomentServiceServer() {}

// UnsafeMomentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MomentServiceServer will
// result in compilation errors.
type UnsafeMomentServiceServer interface {
	mustEmbedUnimplementedMomentServiceServer()
}

func RegisterMomentServiceServer(s grpc.ServiceRegistrar, srv MomentServiceServer) {
	s.RegisterService(&MomentService_ServiceDesc, srv)
}

func _MomentService_PostMoment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostMomentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MomentServiceServer).PostMoment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moment.MomentService/PostMoment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MomentServiceServer).PostMoment(ctx, req.(*PostMomentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MomentService_GetMoment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMomentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MomentServiceServer).GetMoment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moment.MomentService/GetMoment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MomentServiceServer).GetMoment(ctx, req.(*GetMomentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MomentService_GetMoments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMomentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MomentServiceServer).GetMoments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moment.MomentService/GetMoments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MomentServiceServer).GetMoments(ctx, req.(*GetMomentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MomentService_GetMomentsReported_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMomentsReportedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MomentServiceServer).GetMomentsReported(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moment.MomentService/GetMomentsReported",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MomentServiceServer).GetMomentsReported(ctx, req.(*GetMomentsReportedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MomentService_GetRetailersBlocked_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRetailersBlockedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MomentServiceServer).GetRetailersBlocked(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moment.MomentService/GetRetailersBlocked",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MomentServiceServer).GetRetailersBlocked(ctx, req.(*GetRetailersBlockedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MomentService_GetRetailersNearby_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRetailersNearbyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MomentServiceServer).GetRetailersNearby(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moment.MomentService/GetRetailersNearby",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MomentServiceServer).GetRetailersNearby(ctx, req.(*GetRetailersNearbyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MomentService_PutMoment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutMomentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MomentServiceServer).PutMoment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moment.MomentService/PutMoment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MomentServiceServer).PutMoment(ctx, req.(*PutMomentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MomentService_DeleteMoment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMomentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MomentServiceServer).DeleteMoment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moment.MomentService/DeleteMoment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MomentServiceServer).DeleteMoment(ctx, req.(*DeleteMomentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MomentService_PostMomentReported_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostMomentReportedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MomentServiceServer).PostMomentReported(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moment.MomentService/PostMomentReported",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MomentServiceServer).PostMomentReported(ctx, req.(*PostMomentReportedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MomentService_DeleteMomentReported_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMomentReportedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MomentServiceServer).DeleteMomentReported(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moment.MomentService/DeleteMomentReported",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MomentServiceServer).DeleteMomentReported(ctx, req.(*DeleteMomentReportedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MomentService_PostRetailerBlocked_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostRetailerBlockedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MomentServiceServer).PostRetailerBlocked(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moment.MomentService/PostRetailerBlocked",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MomentServiceServer).PostRetailerBlocked(ctx, req.(*PostRetailerBlockedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MomentService_DeleteRetailerBlocked_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRetailerBlockedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MomentServiceServer).DeleteRetailerBlocked(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moment.MomentService/DeleteRetailerBlocked",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MomentServiceServer).DeleteRetailerBlocked(ctx, req.(*DeleteRetailerBlockedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MomentService_ServiceDesc is the grpc.ServiceDesc for MomentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MomentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "moment.MomentService",
	HandlerType: (*MomentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PostMoment",
			Handler:    _MomentService_PostMoment_Handler,
		},
		{
			MethodName: "GetMoment",
			Handler:    _MomentService_GetMoment_Handler,
		},
		{
			MethodName: "GetMoments",
			Handler:    _MomentService_GetMoments_Handler,
		},
		{
			MethodName: "GetMomentsReported",
			Handler:    _MomentService_GetMomentsReported_Handler,
		},
		{
			MethodName: "GetRetailersBlocked",
			Handler:    _MomentService_GetRetailersBlocked_Handler,
		},
		{
			MethodName: "GetRetailersNearby",
			Handler:    _MomentService_GetRetailersNearby_Handler,
		},
		{
			MethodName: "PutMoment",
			Handler:    _MomentService_PutMoment_Handler,
		},
		{
			MethodName: "DeleteMoment",
			Handler:    _MomentService_DeleteMoment_Handler,
		},
		{
			MethodName: "PostMomentReported",
			Handler:    _MomentService_PostMomentReported_Handler,
		},
		{
			MethodName: "DeleteMomentReported",
			Handler:    _MomentService_DeleteMomentReported_Handler,
		},
		{
			MethodName: "PostRetailerBlocked",
			Handler:    _MomentService_PostRetailerBlocked_Handler,
		},
		{
			MethodName: "DeleteRetailerBlocked",
			Handler:    _MomentService_DeleteRetailerBlocked_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "moment/moment.proto",
}