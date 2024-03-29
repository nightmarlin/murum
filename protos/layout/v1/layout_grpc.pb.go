// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: layout/v1/layout.proto

package v1

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

// LayoutServiceClient is the client API for LayoutService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LayoutServiceClient interface {
	Generate(ctx context.Context, in *GenerateRequest, opts ...grpc.CallOption) (*GenerateResponse, error)
}

type layoutServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLayoutServiceClient(cc grpc.ClientConnInterface) LayoutServiceClient {
	return &layoutServiceClient{cc}
}

func (c *layoutServiceClient) Generate(ctx context.Context, in *GenerateRequest, opts ...grpc.CallOption) (*GenerateResponse, error) {
	out := new(GenerateResponse)
	err := c.cc.Invoke(ctx, "/layout.LayoutService/Generate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LayoutServiceServer is the server API for LayoutService service.
// All implementations must embed UnimplementedLayoutServiceServer
// for forward compatibility
type LayoutServiceServer interface {
	Generate(context.Context, *GenerateRequest) (*GenerateResponse, error)
	mustEmbedUnimplementedLayoutServiceServer()
}

// UnimplementedLayoutServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLayoutServiceServer struct {
}

func (UnimplementedLayoutServiceServer) Generate(context.Context, *GenerateRequest) (*GenerateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Generate not implemented")
}
func (UnimplementedLayoutServiceServer) mustEmbedUnimplementedLayoutServiceServer() {}

// UnsafeLayoutServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LayoutServiceServer will
// result in compilation errors.
type UnsafeLayoutServiceServer interface {
	mustEmbedUnimplementedLayoutServiceServer()
}

func RegisterLayoutServiceServer(s grpc.ServiceRegistrar, srv LayoutServiceServer) {
	s.RegisterService(&LayoutService_ServiceDesc, srv)
}

func _LayoutService_Generate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LayoutServiceServer).Generate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/layout.LayoutService/Generate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LayoutServiceServer).Generate(ctx, req.(*GenerateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LayoutService_ServiceDesc is the grpc.ServiceDesc for LayoutService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LayoutService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "layout.LayoutService",
	HandlerType: (*LayoutServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Generate",
			Handler:    _LayoutService_Generate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "layout/v1/layout.proto",
}
