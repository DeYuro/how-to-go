// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package dummy

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

// FooServiceClient is the client API for FooService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FooServiceClient interface {
	Reverse(ctx context.Context, in *ReverseRequest, opts ...grpc.CallOption) (*ReverseResponse, error)
	GetBar(ctx context.Context, in *GetBarRequest, opts ...grpc.CallOption) (*GetBarResponse, error)
}

type fooServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFooServiceClient(cc grpc.ClientConnInterface) FooServiceClient {
	return &fooServiceClient{cc}
}

func (c *fooServiceClient) Reverse(ctx context.Context, in *ReverseRequest, opts ...grpc.CallOption) (*ReverseResponse, error) {
	out := new(ReverseResponse)
	err := c.cc.Invoke(ctx, "/dummy.v1.foo.FooService/reverse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fooServiceClient) GetBar(ctx context.Context, in *GetBarRequest, opts ...grpc.CallOption) (*GetBarResponse, error) {
	out := new(GetBarResponse)
	err := c.cc.Invoke(ctx, "/dummy.v1.foo.FooService/getBar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FooServiceServer is the server API for FooService service.
// All implementations should embed UnimplementedFooServiceServer
// for forward compatibility
type FooServiceServer interface {
	Reverse(context.Context, *ReverseRequest) (*ReverseResponse, error)
	GetBar(context.Context, *GetBarRequest) (*GetBarResponse, error)
}

// UnimplementedFooServiceServer should be embedded to have forward compatible implementations.
type UnimplementedFooServiceServer struct {
}

func (UnimplementedFooServiceServer) Reverse(context.Context, *ReverseRequest) (*ReverseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Reverse not implemented")
}
func (UnimplementedFooServiceServer) GetBar(context.Context, *GetBarRequest) (*GetBarResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBar not implemented")
}

// UnsafeFooServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FooServiceServer will
// result in compilation errors.
type UnsafeFooServiceServer interface {
	mustEmbedUnimplementedFooServiceServer()
}

func RegisterFooServiceServer(s grpc.ServiceRegistrar, srv FooServiceServer) {
	s.RegisterService(&FooService_ServiceDesc, srv)
}

func _FooService_Reverse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReverseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FooServiceServer).Reverse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dummy.v1.foo.FooService/reverse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FooServiceServer).Reverse(ctx, req.(*ReverseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FooService_GetBar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FooServiceServer).GetBar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dummy.v1.foo.FooService/getBar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FooServiceServer).GetBar(ctx, req.(*GetBarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FooService_ServiceDesc is the grpc.ServiceDesc for FooService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FooService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dummy.v1.foo.FooService",
	HandlerType: (*FooServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "reverse",
			Handler:    _FooService_Reverse_Handler,
		},
		{
			MethodName: "getBar",
			Handler:    _FooService_GetBar_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/v1/foo.proto",
}
