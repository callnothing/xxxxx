// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package XXXXTT

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

// XXXXClient is the client API for XXXX service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type XXXXClient interface {
	SayHello(ctx context.Context, in *HelloRequestt, opts ...grpc.CallOption) (*HelloReply, error)
}

type xXXXClient struct {
	cc grpc.ClientConnInterface
}

func NewXXXXClient(cc grpc.ClientConnInterface) XXXXClient {
	return &xXXXClient{cc}
}

func (c *xXXXClient) SayHello(ctx context.Context, in *HelloRequestt, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/XXXXTTT.XXXXTT.XXXX/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// XXXXServer is the server API for XXXX service.
// All implementations must embed UnimplementedXXXXServer
// for forward compatibility
type XXXXServer interface {
	SayHello(context.Context, *HelloRequestt) (*HelloReply, error)
	mustEmbedUnimplementedXXXXServer()
}

// UnimplementedXXXXServer must be embedded to have forward compatible implementations.
type UnimplementedXXXXServer struct {
}

func (UnimplementedXXXXServer) SayHello(context.Context, *HelloRequestt) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedXXXXServer) mustEmbedUnimplementedXXXXServer() {}

// UnsafeXXXXServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to XXXXServer will
// result in compilation errors.
type UnsafeXXXXServer interface {
	mustEmbedUnimplementedXXXXServer()
}

func RegisterXXXXServer(s grpc.ServiceRegistrar, srv XXXXServer) {
	s.RegisterService(&XXXX_ServiceDesc, srv)
}

func _XXXX_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequestt)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XXXXServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/XXXXTTT.XXXXTT.XXXX/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XXXXServer).SayHello(ctx, req.(*HelloRequestt))
	}
	return interceptor(ctx, in, info, handler)
}

// XXXX_ServiceDesc is the grpc.ServiceDesc for XXXX service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var XXXX_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "XXXXTTT.XXXXTT.XXXX",
	HandlerType: (*XXXXServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _XXXX_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "XXXX.proto",
}
