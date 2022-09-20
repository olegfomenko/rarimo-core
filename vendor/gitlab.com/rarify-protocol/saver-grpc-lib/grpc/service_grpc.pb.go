// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.3.0
// source: proto/service.proto

package grpc

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

// SaverClient is the client API for Saver service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SaverClient interface {
	GetDepositInfo(ctx context.Context, in *MsgTransactionInfoRequest, opts ...grpc.CallOption) (*MsgDepositResponse, error)
	GetWithdrawInfo(ctx context.Context, in *MsgTransactionInfoRequest, opts ...grpc.CallOption) (*MsgWithdrawResponse, error)
}

type saverClient struct {
	cc grpc.ClientConnInterface
}

func NewSaverClient(cc grpc.ClientConnInterface) SaverClient {
	return &saverClient{cc}
}

func (c *saverClient) GetDepositInfo(ctx context.Context, in *MsgTransactionInfoRequest, opts ...grpc.CallOption) (*MsgDepositResponse, error) {
	out := new(MsgDepositResponse)
	err := c.cc.Invoke(ctx, "/Saver/GetDepositInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *saverClient) GetWithdrawInfo(ctx context.Context, in *MsgTransactionInfoRequest, opts ...grpc.CallOption) (*MsgWithdrawResponse, error) {
	out := new(MsgWithdrawResponse)
	err := c.cc.Invoke(ctx, "/Saver/GetWithdrawInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SaverServer is the server API for Saver service.
// All implementations must embed UnimplementedSaverServer
// for forward compatibility
type SaverServer interface {
	GetDepositInfo(context.Context, *MsgTransactionInfoRequest) (*MsgDepositResponse, error)
	GetWithdrawInfo(context.Context, *MsgTransactionInfoRequest) (*MsgWithdrawResponse, error)
	mustEmbedUnimplementedSaverServer()
}

// UnimplementedSaverServer must be embedded to have forward compatible implementations.
type UnimplementedSaverServer struct {
}

func (UnimplementedSaverServer) GetDepositInfo(context.Context, *MsgTransactionInfoRequest) (*MsgDepositResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDepositInfo not implemented")
}
func (UnimplementedSaverServer) GetWithdrawInfo(context.Context, *MsgTransactionInfoRequest) (*MsgWithdrawResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWithdrawInfo not implemented")
}
func (UnimplementedSaverServer) mustEmbedUnimplementedSaverServer() {}

// UnsafeSaverServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SaverServer will
// result in compilation errors.
type UnsafeSaverServer interface {
	mustEmbedUnimplementedSaverServer()
}

func RegisterSaverServer(s grpc.ServiceRegistrar, srv SaverServer) {
	s.RegisterService(&Saver_ServiceDesc, srv)
}

func _Saver_GetDepositInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgTransactionInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SaverServer).GetDepositInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Saver/GetDepositInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SaverServer).GetDepositInfo(ctx, req.(*MsgTransactionInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Saver_GetWithdrawInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgTransactionInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SaverServer).GetWithdrawInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Saver/GetWithdrawInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SaverServer).GetWithdrawInfo(ctx, req.(*MsgTransactionInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Saver_ServiceDesc is the grpc.ServiceDesc for Saver service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Saver_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Saver",
	HandlerType: (*SaverServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDepositInfo",
			Handler:    _Saver_GetDepositInfo_Handler,
		},
		{
			MethodName: "GetWithdrawInfo",
			Handler:    _Saver_GetWithdrawInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/service.proto",
}
