// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package xxxxtttttt

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

// XxxxxxClient is the client API for Xxxxxx service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type XxxxxxClient interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type xxxxxxClient struct {
	cc grpc.ClientConnInterface
}

func NewXxxxxxClient(cc grpc.ClientConnInterface) XxxxxxClient {
	return &xxxxxxClient{cc}
}

func (c *xxxxxxClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/xxxttt.xxxxtttttt.xxxxxx/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// XxxxxxServer is the server API for Xxxxxx service.
// All implementations must embed UnimplementedXxxxxxServer
// for forward compatibility
type XxxxxxServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	mustEmbedUnimplementedXxxxxxServer()
}

// UnimplementedXxxxxxServer must be embedded to have forward compatible implementations.
type UnimplementedXxxxxxServer struct {
}

func (UnimplementedXxxxxxServer) SayHello(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedXxxxxxServer) mustEmbedUnimplementedXxxxxxServer() {}

// UnsafeXxxxxxServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to XxxxxxServer will
// result in compilation errors.
type UnsafeXxxxxxServer interface {
	mustEmbedUnimplementedXxxxxxServer()
}

func RegisterXxxxxxServer(s grpc.ServiceRegistrar, srv XxxxxxServer) {
	s.RegisterService(&Xxxxxx_ServiceDesc, srv)
}

func _Xxxxxx_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XxxxxxServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xxxttt.xxxxtttttt.xxxxxx/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XxxxxxServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Xxxxxx_ServiceDesc is the grpc.ServiceDesc for Xxxxxx service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Xxxxxx_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "xxxttt.xxxxtttttt.xxxxxx",
	HandlerType: (*XxxxxxServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Xxxxxx_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "xxxxxx.proto",
}
