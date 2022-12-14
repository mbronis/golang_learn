// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: arithmetic_svc.proto

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

// ArithemticServiceClient is the client API for ArithemticService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ArithemticServiceClient interface {
	GetAddition(ctx context.Context, in *OperationParameters, opts ...grpc.CallOption) (*Answer, error)
	GetSubtraction(ctx context.Context, in *OperationParameters, opts ...grpc.CallOption) (*Answer, error)
	GetMultiplication(ctx context.Context, in *OperationParameters, opts ...grpc.CallOption) (*Answer, error)
	GetDivision(ctx context.Context, in *OperationParameters, opts ...grpc.CallOption) (*Answer, error)
}

type arithemticServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewArithemticServiceClient(cc grpc.ClientConnInterface) ArithemticServiceClient {
	return &arithemticServiceClient{cc}
}

func (c *arithemticServiceClient) GetAddition(ctx context.Context, in *OperationParameters, opts ...grpc.CallOption) (*Answer, error) {
	out := new(Answer)
	err := c.cc.Invoke(ctx, "/pb.ArithemticService/GetAddition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *arithemticServiceClient) GetSubtraction(ctx context.Context, in *OperationParameters, opts ...grpc.CallOption) (*Answer, error) {
	out := new(Answer)
	err := c.cc.Invoke(ctx, "/pb.ArithemticService/GetSubtraction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *arithemticServiceClient) GetMultiplication(ctx context.Context, in *OperationParameters, opts ...grpc.CallOption) (*Answer, error) {
	out := new(Answer)
	err := c.cc.Invoke(ctx, "/pb.ArithemticService/GetMultiplication", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *arithemticServiceClient) GetDivision(ctx context.Context, in *OperationParameters, opts ...grpc.CallOption) (*Answer, error) {
	out := new(Answer)
	err := c.cc.Invoke(ctx, "/pb.ArithemticService/GetDivision", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArithemticServiceServer is the server API for ArithemticService service.
// All implementations must embed UnimplementedArithemticServiceServer
// for forward compatibility
type ArithemticServiceServer interface {
	GetAddition(context.Context, *OperationParameters) (*Answer, error)
	GetSubtraction(context.Context, *OperationParameters) (*Answer, error)
	GetMultiplication(context.Context, *OperationParameters) (*Answer, error)
	GetDivision(context.Context, *OperationParameters) (*Answer, error)
	mustEmbedUnimplementedArithemticServiceServer()
}

// UnimplementedArithemticServiceServer must be embedded to have forward compatible implementations.
type UnimplementedArithemticServiceServer struct {
}

func (UnimplementedArithemticServiceServer) GetAddition(context.Context, *OperationParameters) (*Answer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAddition not implemented")
}
func (UnimplementedArithemticServiceServer) GetSubtraction(context.Context, *OperationParameters) (*Answer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubtraction not implemented")
}
func (UnimplementedArithemticServiceServer) GetMultiplication(context.Context, *OperationParameters) (*Answer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMultiplication not implemented")
}
func (UnimplementedArithemticServiceServer) GetDivision(context.Context, *OperationParameters) (*Answer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDivision not implemented")
}
func (UnimplementedArithemticServiceServer) mustEmbedUnimplementedArithemticServiceServer() {}

// UnsafeArithemticServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ArithemticServiceServer will
// result in compilation errors.
type UnsafeArithemticServiceServer interface {
	mustEmbedUnimplementedArithemticServiceServer()
}

func RegisterArithemticServiceServer(s grpc.ServiceRegistrar, srv ArithemticServiceServer) {
	s.RegisterService(&ArithemticService_ServiceDesc, srv)
}

func _ArithemticService_GetAddition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OperationParameters)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArithemticServiceServer).GetAddition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ArithemticService/GetAddition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArithemticServiceServer).GetAddition(ctx, req.(*OperationParameters))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArithemticService_GetSubtraction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OperationParameters)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArithemticServiceServer).GetSubtraction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ArithemticService/GetSubtraction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArithemticServiceServer).GetSubtraction(ctx, req.(*OperationParameters))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArithemticService_GetMultiplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OperationParameters)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArithemticServiceServer).GetMultiplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ArithemticService/GetMultiplication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArithemticServiceServer).GetMultiplication(ctx, req.(*OperationParameters))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArithemticService_GetDivision_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OperationParameters)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArithemticServiceServer).GetDivision(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ArithemticService/GetDivision",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArithemticServiceServer).GetDivision(ctx, req.(*OperationParameters))
	}
	return interceptor(ctx, in, info, handler)
}

// ArithemticService_ServiceDesc is the grpc.ServiceDesc for ArithemticService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ArithemticService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ArithemticService",
	HandlerType: (*ArithemticServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAddition",
			Handler:    _ArithemticService_GetAddition_Handler,
		},
		{
			MethodName: "GetSubtraction",
			Handler:    _ArithemticService_GetSubtraction_Handler,
		},
		{
			MethodName: "GetMultiplication",
			Handler:    _ArithemticService_GetMultiplication_Handler,
		},
		{
			MethodName: "GetDivision",
			Handler:    _ArithemticService_GetDivision_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "arithmetic_svc.proto",
}
