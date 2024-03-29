// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: sync_api.proto

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
	BatchService_Batch_FullMethodName = "/api.BatchService/Batch"
)

// BatchServiceClient is the client API for BatchService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BatchServiceClient interface {
	Batch(ctx context.Context, in *PBatch, opts ...grpc.CallOption) (*Result, error)
}

type batchServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBatchServiceClient(cc grpc.ClientConnInterface) BatchServiceClient {
	return &batchServiceClient{cc}
}

func (c *batchServiceClient) Batch(ctx context.Context, in *PBatch, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, BatchService_Batch_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BatchServiceServer is the server API for BatchService service.
// All implementations must embed UnimplementedBatchServiceServer
// for forward compatibility
type BatchServiceServer interface {
	Batch(context.Context, *PBatch) (*Result, error)
	mustEmbedUnimplementedBatchServiceServer()
}

// UnimplementedBatchServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBatchServiceServer struct {
}

func (UnimplementedBatchServiceServer) Batch(context.Context, *PBatch) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Batch not implemented")
}
func (UnimplementedBatchServiceServer) mustEmbedUnimplementedBatchServiceServer() {}

// UnsafeBatchServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BatchServiceServer will
// result in compilation errors.
type UnsafeBatchServiceServer interface {
	mustEmbedUnimplementedBatchServiceServer()
}

func RegisterBatchServiceServer(s grpc.ServiceRegistrar, srv BatchServiceServer) {
	s.RegisterService(&BatchService_ServiceDesc, srv)
}

func _BatchService_Batch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PBatch)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BatchServiceServer).Batch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BatchService_Batch_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BatchServiceServer).Batch(ctx, req.(*PBatch))
	}
	return interceptor(ctx, in, info, handler)
}

// BatchService_ServiceDesc is the grpc.ServiceDesc for BatchService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BatchService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.BatchService",
	HandlerType: (*BatchServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Batch",
			Handler:    _BatchService_Batch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sync_api.proto",
}
