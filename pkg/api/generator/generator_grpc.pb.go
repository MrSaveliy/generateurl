// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: api/generator/generator.proto

package generator

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
	Generator_GetUrl_FullMethodName = "/generator.Generator/GetUrl"
	Generator_SetUrl_FullMethodName = "/generator.Generator/SetUrl"
)

// GeneratorClient is the client API for Generator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GeneratorClient interface {
	GetUrl(ctx context.Context, in *GetUrlRequest, opts ...grpc.CallOption) (*GetUrlResponse, error)
	SetUrl(ctx context.Context, in *SetUrlRequest, opts ...grpc.CallOption) (*SetUrlResponse, error)
}

type generatorClient struct {
	cc grpc.ClientConnInterface
}

func NewGeneratorClient(cc grpc.ClientConnInterface) GeneratorClient {
	return &generatorClient{cc}
}

func (c *generatorClient) GetUrl(ctx context.Context, in *GetUrlRequest, opts ...grpc.CallOption) (*GetUrlResponse, error) {
	out := new(GetUrlResponse)
	err := c.cc.Invoke(ctx, Generator_GetUrl_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *generatorClient) SetUrl(ctx context.Context, in *SetUrlRequest, opts ...grpc.CallOption) (*SetUrlResponse, error) {
	out := new(SetUrlResponse)
	err := c.cc.Invoke(ctx, Generator_SetUrl_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GeneratorServer is the server API for Generator service.
// All implementations must embed UnimplementedGeneratorServer
// for forward compatibility
type GeneratorServer interface {
	GetUrl(context.Context, *GetUrlRequest) (*GetUrlResponse, error)
	SetUrl(context.Context, *SetUrlRequest) (*SetUrlResponse, error)
	mustEmbedUnimplementedGeneratorServer()
}

// UnimplementedGeneratorServer must be embedded to have forward compatible implementations.
type UnimplementedGeneratorServer struct {
}

func (UnimplementedGeneratorServer) GetUrl(context.Context, *GetUrlRequest) (*GetUrlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUrl not implemented")
}
func (UnimplementedGeneratorServer) SetUrl(context.Context, *SetUrlRequest) (*SetUrlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetUrl not implemented")
}
func (UnimplementedGeneratorServer) mustEmbedUnimplementedGeneratorServer() {}

// UnsafeGeneratorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GeneratorServer will
// result in compilation errors.
type UnsafeGeneratorServer interface {
	mustEmbedUnimplementedGeneratorServer()
}

func RegisterGeneratorServer(s grpc.ServiceRegistrar, srv GeneratorServer) {
	s.RegisterService(&Generator_ServiceDesc, srv)
}

func _Generator_GetUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUrlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeneratorServer).GetUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Generator_GetUrl_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeneratorServer).GetUrl(ctx, req.(*GetUrlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Generator_SetUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetUrlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeneratorServer).SetUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Generator_SetUrl_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeneratorServer).SetUrl(ctx, req.(*SetUrlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Generator_ServiceDesc is the grpc.ServiceDesc for Generator service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Generator_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "generator.Generator",
	HandlerType: (*GeneratorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUrl",
			Handler:    _Generator_GetUrl_Handler,
		},
		{
			MethodName: "SetUrl",
			Handler:    _Generator_SetUrl_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/generator/generator.proto",
}
