// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: internal/infra/grpc/protofiles/schema_input.proto

package pb

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
	SchemaInputService_CreateSchemaInput_FullMethodName = "/pb.SchemaInputService/CreateSchemaInput"
)

// SchemaInputServiceClient is the client API for SchemaInputService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SchemaInputServiceClient interface {
	CreateSchemaInput(ctx context.Context, in *CreateSchemaInputRequest, opts ...grpc.CallOption) (*CreateSchemaInputResponse, error)
}

type schemaInputServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSchemaInputServiceClient(cc grpc.ClientConnInterface) SchemaInputServiceClient {
	return &schemaInputServiceClient{cc}
}

func (c *schemaInputServiceClient) CreateSchemaInput(ctx context.Context, in *CreateSchemaInputRequest, opts ...grpc.CallOption) (*CreateSchemaInputResponse, error) {
	out := new(CreateSchemaInputResponse)
	err := c.cc.Invoke(ctx, SchemaInputService_CreateSchemaInput_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SchemaInputServiceServer is the server API for SchemaInputService service.
// All implementations must embed UnimplementedSchemaInputServiceServer
// for forward compatibility
type SchemaInputServiceServer interface {
	CreateSchemaInput(context.Context, *CreateSchemaInputRequest) (*CreateSchemaInputResponse, error)
	mustEmbedUnimplementedSchemaInputServiceServer()
}

// UnimplementedSchemaInputServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSchemaInputServiceServer struct {
}

func (UnimplementedSchemaInputServiceServer) CreateSchemaInput(context.Context, *CreateSchemaInputRequest) (*CreateSchemaInputResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSchemaInput not implemented")
}
func (UnimplementedSchemaInputServiceServer) mustEmbedUnimplementedSchemaInputServiceServer() {}

// UnsafeSchemaInputServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SchemaInputServiceServer will
// result in compilation errors.
type UnsafeSchemaInputServiceServer interface {
	mustEmbedUnimplementedSchemaInputServiceServer()
}

func RegisterSchemaInputServiceServer(s grpc.ServiceRegistrar, srv SchemaInputServiceServer) {
	s.RegisterService(&SchemaInputService_ServiceDesc, srv)
}

func _SchemaInputService_CreateSchemaInput_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSchemaInputRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchemaInputServiceServer).CreateSchemaInput(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SchemaInputService_CreateSchemaInput_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchemaInputServiceServer).CreateSchemaInput(ctx, req.(*CreateSchemaInputRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SchemaInputService_ServiceDesc is the grpc.ServiceDesc for SchemaInputService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SchemaInputService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.SchemaInputService",
	HandlerType: (*SchemaInputServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSchemaInput",
			Handler:    _SchemaInputService_CreateSchemaInput_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/infra/grpc/protofiles/schema_input.proto",
}
