// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package xxxxx

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

// AAAAAAClient is the client API for AAAAAA service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AAAAAAClient interface {
	Heathcheck(ctx context.Context, in *HealthcheckRequest, opts ...grpc.CallOption) (*HealthcheckResponse, error)
	Helloworld(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type aAAAAAClient struct {
	cc grpc.ClientConnInterface
}

func NewAAAAAAClient(cc grpc.ClientConnInterface) AAAAAAClient {
	return &aAAAAAClient{cc}
}

func (c *aAAAAAClient) Heathcheck(ctx context.Context, in *HealthcheckRequest, opts ...grpc.CallOption) (*HealthcheckResponse, error) {
	out := new(HealthcheckResponse)
	err := c.cc.Invoke(ctx, "/xxxxx.xxxxx.AAAAAA/Heathcheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aAAAAAClient) Helloworld(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/xxxxx.xxxxx.AAAAAA/Helloworld", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AAAAAAServer is the server API for AAAAAA service.
// All implementations must embed UnimplementedAAAAAAServer
// for forward compatibility
type AAAAAAServer interface {
	Heathcheck(context.Context, *HealthcheckRequest) (*HealthcheckResponse, error)
	Helloworld(context.Context, *HelloRequest) (*HelloReply, error)
	mustEmbedUnimplementedAAAAAAServer()
}

// UnimplementedAAAAAAServer must be embedded to have forward compatible implementations.
type UnimplementedAAAAAAServer struct {
}

func (UnimplementedAAAAAAServer) Heathcheck(context.Context, *HealthcheckRequest) (*HealthcheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Heathcheck not implemented")
}
func (UnimplementedAAAAAAServer) Helloworld(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Helloworld not implemented")
}
func (UnimplementedAAAAAAServer) mustEmbedUnimplementedAAAAAAServer() {}

// UnsafeAAAAAAServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AAAAAAServer will
// result in compilation errors.
type UnsafeAAAAAAServer interface {
	mustEmbedUnimplementedAAAAAAServer()
}

func RegisterAAAAAAServer(s grpc.ServiceRegistrar, srv AAAAAAServer) {
	s.RegisterService(&AAAAAA_ServiceDesc, srv)
}

func _AAAAAA_Heathcheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthcheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AAAAAAServer).Heathcheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xxxxx.xxxxx.AAAAAA/Heathcheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AAAAAAServer).Heathcheck(ctx, req.(*HealthcheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AAAAAA_Helloworld_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AAAAAAServer).Helloworld(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xxxxx.xxxxx.AAAAAA/Helloworld",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AAAAAAServer).Helloworld(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AAAAAA_ServiceDesc is the grpc.ServiceDesc for AAAAAA service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AAAAAA_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "xxxxx.xxxxx.AAAAAA",
	HandlerType: (*AAAAAAServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Heathcheck",
			Handler:    _AAAAAA_Heathcheck_Handler,
		},
		{
			MethodName: "Helloworld",
			Handler:    _AAAAAA_Helloworld_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "AAAAAA.proto",
}
