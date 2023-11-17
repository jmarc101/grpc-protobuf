// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.9
// source: proto/greeting/goodbye.proto

package greeting

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
	GoodbyeService_Goodbye_FullMethodName = "/GoodbyeService/Goodbye"
)

// GoodbyeServiceClient is the client API for GoodbyeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GoodbyeServiceClient interface {
	Goodbye(ctx context.Context, in *GoodbyeRequest, opts ...grpc.CallOption) (*GoodbyeResponse, error)
}

type goodbyeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGoodbyeServiceClient(cc grpc.ClientConnInterface) GoodbyeServiceClient {
	return &goodbyeServiceClient{cc}
}

func (c *goodbyeServiceClient) Goodbye(ctx context.Context, in *GoodbyeRequest, opts ...grpc.CallOption) (*GoodbyeResponse, error) {
	out := new(GoodbyeResponse)
	err := c.cc.Invoke(ctx, GoodbyeService_Goodbye_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GoodbyeServiceServer is the server API for GoodbyeService service.
// All implementations must embed UnimplementedGoodbyeServiceServer
// for forward compatibility
type GoodbyeServiceServer interface {
	Goodbye(context.Context, *GoodbyeRequest) (*GoodbyeResponse, error)
	mustEmbedUnimplementedGoodbyeServiceServer()
}

// UnimplementedGoodbyeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGoodbyeServiceServer struct {
}

func (UnimplementedGoodbyeServiceServer) Goodbye(context.Context, *GoodbyeRequest) (*GoodbyeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Goodbye not implemented")
}
func (UnimplementedGoodbyeServiceServer) mustEmbedUnimplementedGoodbyeServiceServer() {}

// UnsafeGoodbyeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GoodbyeServiceServer will
// result in compilation errors.
type UnsafeGoodbyeServiceServer interface {
	mustEmbedUnimplementedGoodbyeServiceServer()
}

func RegisterGoodbyeServiceServer(s grpc.ServiceRegistrar, srv GoodbyeServiceServer) {
	s.RegisterService(&GoodbyeService_ServiceDesc, srv)
}

func _GoodbyeService_Goodbye_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GoodbyeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodbyeServiceServer).Goodbye(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GoodbyeService_Goodbye_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodbyeServiceServer).Goodbye(ctx, req.(*GoodbyeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GoodbyeService_ServiceDesc is the grpc.ServiceDesc for GoodbyeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GoodbyeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "GoodbyeService",
	HandlerType: (*GoodbyeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Goodbye",
			Handler:    _GoodbyeService_Goodbye_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/greeting/goodbye.proto",
}
