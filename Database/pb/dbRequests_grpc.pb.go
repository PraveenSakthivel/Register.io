// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package dbRequests

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

// DatabseWrapperClient is the client API for DatabseWrapper service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DatabseWrapperClient interface {
	RetrieveClasses(ctx context.Context, in *ReceiveClassesParams, opts ...grpc.CallOption) (*ClassesResponse, error)
}

type databseWrapperClient struct {
	cc grpc.ClientConnInterface
}

func NewDatabseWrapperClient(cc grpc.ClientConnInterface) DatabseWrapperClient {
	return &databseWrapperClient{cc}
}

func (c *databseWrapperClient) RetrieveClasses(ctx context.Context, in *ReceiveClassesParams, opts ...grpc.CallOption) (*ClassesResponse, error) {
	out := new(ClassesResponse)
	err := c.cc.Invoke(ctx, "/dbRequests.DatabseWrapper/RetrieveClasses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DatabseWrapperServer is the server API for DatabseWrapper service.
// All implementations must embed UnimplementedDatabseWrapperServer
// for forward compatibility
type DatabseWrapperServer interface {
	RetrieveClasses(context.Context, *ReceiveClassesParams) (*ClassesResponse, error)
	mustEmbedUnimplementedDatabseWrapperServer()
}

// UnimplementedDatabseWrapperServer must be embedded to have forward compatible implementations.
type UnimplementedDatabseWrapperServer struct {
}

func (UnimplementedDatabseWrapperServer) RetrieveClasses(context.Context, *ReceiveClassesParams) (*ClassesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RetrieveClasses not implemented")
}
func (UnimplementedDatabseWrapperServer) mustEmbedUnimplementedDatabseWrapperServer() {}

// UnsafeDatabseWrapperServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DatabseWrapperServer will
// result in compilation errors.
type UnsafeDatabseWrapperServer interface {
	mustEmbedUnimplementedDatabseWrapperServer()
}

func RegisterDatabseWrapperServer(s grpc.ServiceRegistrar, srv DatabseWrapperServer) {
	s.RegisterService(&DatabseWrapper_ServiceDesc, srv)
}

func _DatabseWrapper_RetrieveClasses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReceiveClassesParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabseWrapperServer).RetrieveClasses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dbRequests.DatabseWrapper/RetrieveClasses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabseWrapperServer).RetrieveClasses(ctx, req.(*ReceiveClassesParams))
	}
	return interceptor(ctx, in, info, handler)
}

// DatabseWrapper_ServiceDesc is the grpc.ServiceDesc for DatabseWrapper service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DatabseWrapper_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dbRequests.DatabseWrapper",
	HandlerType: (*DatabseWrapperServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RetrieveClasses",
			Handler:    _DatabseWrapper_RetrieveClasses_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/dbRequests.proto",
}
