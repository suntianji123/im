// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: result.proto

package api

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
	HeartBeartService_HeartBeart_FullMethodName = "/api.HeartBeartService/HeartBeart"
)

// HeartBeartServiceClient is the client API for HeartBeartService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HeartBeartServiceClient interface {
	HeartBeart(ctx context.Context, in *PPing, opts ...grpc.CallOption) (*PPong, error)
}

type heartBeartServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHeartBeartServiceClient(cc grpc.ClientConnInterface) HeartBeartServiceClient {
	return &heartBeartServiceClient{cc}
}

func (c *heartBeartServiceClient) HeartBeart(ctx context.Context, in *PPing, opts ...grpc.CallOption) (*PPong, error) {
	out := new(PPong)
	err := c.cc.Invoke(ctx, HeartBeartService_HeartBeart_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HeartBeartServiceServer is the server API for HeartBeartService service.
// All implementations must embed UnimplementedHeartBeartServiceServer
// for forward compatibility
type HeartBeartServiceServer interface {
	HeartBeart(context.Context, *PPing) (*PPong, error)
	mustEmbedUnimplementedHeartBeartServiceServer()
}

// UnimplementedHeartBeartServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHeartBeartServiceServer struct {
}

func (UnimplementedHeartBeartServiceServer) HeartBeart(context.Context, *PPing) (*PPong, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HeartBeart not implemented")
}
func (UnimplementedHeartBeartServiceServer) mustEmbedUnimplementedHeartBeartServiceServer() {}

// UnsafeHeartBeartServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HeartBeartServiceServer will
// result in compilation errors.
type UnsafeHeartBeartServiceServer interface {
	mustEmbedUnimplementedHeartBeartServiceServer()
}

func RegisterHeartBeartServiceServer(s grpc.ServiceRegistrar, srv HeartBeartServiceServer) {
	s.RegisterService(&HeartBeartService_ServiceDesc, srv)
}

func _HeartBeartService_HeartBeart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PPing)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeartBeartServiceServer).HeartBeart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HeartBeartService_HeartBeart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeartBeartServiceServer).HeartBeart(ctx, req.(*PPing))
	}
	return interceptor(ctx, in, info, handler)
}

// HeartBeartService_ServiceDesc is the grpc.ServiceDesc for HeartBeartService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HeartBeartService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.HeartBeartService",
	HandlerType: (*HeartBeartServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HeartBeart",
			Handler:    _HeartBeartService_HeartBeart_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "result.proto",
}
