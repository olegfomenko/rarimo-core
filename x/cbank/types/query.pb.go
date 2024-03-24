// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cbank/query.proto

package types

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// QueryParamsRequest is request type for the Query/Params RPC method.
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_17d639329e3d4b77, []int{0}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

// QueryParamsResponse is response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	// params holds all the parameters of this module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_17d639329e3d4b77, []int{1}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

type QueryGetAllCommitmentRequest struct {
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryGetAllCommitmentRequest) Reset()         { *m = QueryGetAllCommitmentRequest{} }
func (m *QueryGetAllCommitmentRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetAllCommitmentRequest) ProtoMessage()    {}
func (*QueryGetAllCommitmentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_17d639329e3d4b77, []int{2}
}
func (m *QueryGetAllCommitmentRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetAllCommitmentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetAllCommitmentRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetAllCommitmentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetAllCommitmentRequest.Merge(m, src)
}
func (m *QueryGetAllCommitmentRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetAllCommitmentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetAllCommitmentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetAllCommitmentRequest proto.InternalMessageInfo

func (m *QueryGetAllCommitmentRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type QueryGetAllCommitmentResponse struct {
	Commitment []Commitment        `protobuf:"bytes,1,rep,name=commitment,proto3" json:"commitment"`
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryGetAllCommitmentResponse) Reset()         { *m = QueryGetAllCommitmentResponse{} }
func (m *QueryGetAllCommitmentResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetAllCommitmentResponse) ProtoMessage()    {}
func (*QueryGetAllCommitmentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_17d639329e3d4b77, []int{3}
}
func (m *QueryGetAllCommitmentResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetAllCommitmentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetAllCommitmentResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetAllCommitmentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetAllCommitmentResponse.Merge(m, src)
}
func (m *QueryGetAllCommitmentResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetAllCommitmentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetAllCommitmentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetAllCommitmentResponse proto.InternalMessageInfo

func (m *QueryGetAllCommitmentResponse) GetCommitment() []Commitment {
	if m != nil {
		return m.Commitment
	}
	return nil
}

func (m *QueryGetAllCommitmentResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type QueryGetCommitmentRequest struct {
	// compressed form point
	Index string `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
}

func (m *QueryGetCommitmentRequest) Reset()         { *m = QueryGetCommitmentRequest{} }
func (m *QueryGetCommitmentRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetCommitmentRequest) ProtoMessage()    {}
func (*QueryGetCommitmentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_17d639329e3d4b77, []int{4}
}
func (m *QueryGetCommitmentRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetCommitmentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetCommitmentRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetCommitmentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetCommitmentRequest.Merge(m, src)
}
func (m *QueryGetCommitmentRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetCommitmentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetCommitmentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetCommitmentRequest proto.InternalMessageInfo

func (m *QueryGetCommitmentRequest) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

type QueryGetCommitmentResponse struct {
	Commitment Commitment `protobuf:"bytes,1,opt,name=commitment,proto3" json:"commitment"`
}

func (m *QueryGetCommitmentResponse) Reset()         { *m = QueryGetCommitmentResponse{} }
func (m *QueryGetCommitmentResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetCommitmentResponse) ProtoMessage()    {}
func (*QueryGetCommitmentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_17d639329e3d4b77, []int{5}
}
func (m *QueryGetCommitmentResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetCommitmentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetCommitmentResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetCommitmentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetCommitmentResponse.Merge(m, src)
}
func (m *QueryGetCommitmentResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetCommitmentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetCommitmentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetCommitmentResponse proto.InternalMessageInfo

func (m *QueryGetCommitmentResponse) GetCommitment() Commitment {
	if m != nil {
		return m.Commitment
	}
	return Commitment{}
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "rarimo.rarimocore.cbank.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "rarimo.rarimocore.cbank.QueryParamsResponse")
	proto.RegisterType((*QueryGetAllCommitmentRequest)(nil), "rarimo.rarimocore.cbank.QueryGetAllCommitmentRequest")
	proto.RegisterType((*QueryGetAllCommitmentResponse)(nil), "rarimo.rarimocore.cbank.QueryGetAllCommitmentResponse")
	proto.RegisterType((*QueryGetCommitmentRequest)(nil), "rarimo.rarimocore.cbank.QueryGetCommitmentRequest")
	proto.RegisterType((*QueryGetCommitmentResponse)(nil), "rarimo.rarimocore.cbank.QueryGetCommitmentResponse")
}

func init() { proto.RegisterFile("cbank/query.proto", fileDescriptor_17d639329e3d4b77) }

var fileDescriptor_17d639329e3d4b77 = []byte{
	// 508 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0xc1, 0x6e, 0xd3, 0x30,
	0x1c, 0xc6, 0xeb, 0x8d, 0x55, 0xe2, 0x3f, 0x71, 0xc0, 0x54, 0x30, 0xa2, 0x91, 0x55, 0x01, 0x8d,
	0x8a, 0x15, 0x5b, 0xed, 0x80, 0x1b, 0x87, 0x0d, 0xc1, 0xc4, 0x6d, 0x54, 0x9c, 0xb8, 0x39, 0xc1,
	0x84, 0x88, 0xc4, 0xce, 0x62, 0x17, 0x6d, 0x42, 0x5c, 0xb8, 0x23, 0x21, 0xf1, 0x18, 0x70, 0xe3,
	0x25, 0x76, 0x9c, 0xc4, 0x85, 0x13, 0x42, 0x2d, 0x8f, 0xc0, 0x03, 0xa0, 0xda, 0x1e, 0x49, 0xd5,
	0x85, 0xad, 0x3b, 0xa5, 0x75, 0xbf, 0xef, 0xfb, 0xff, 0x3e, 0xf7, 0xaf, 0xc0, 0xe5, 0x28, 0x64,
	0xe2, 0x0d, 0xdd, 0x1b, 0xf2, 0xe2, 0x80, 0xe4, 0x85, 0xd4, 0x12, 0x5f, 0x2b, 0x58, 0x91, 0x64,
	0x92, 0xd8, 0x47, 0x24, 0x0b, 0x4e, 0x8c, 0xc8, 0x6b, 0xc5, 0x32, 0x96, 0x46, 0x43, 0x27, 0x9f,
	0xac, 0xdc, 0x5b, 0x8d, 0xa5, 0x8c, 0x53, 0x4e, 0x59, 0x9e, 0x50, 0x26, 0x84, 0xd4, 0x4c, 0x27,
	0x52, 0x28, 0xf7, 0xeb, 0x9d, 0x48, 0xaa, 0x4c, 0x2a, 0x1a, 0x32, 0xc5, 0xed, 0x14, 0xfa, 0xb6,
	0x17, 0x72, 0xcd, 0x7a, 0x34, 0x67, 0x71, 0x22, 0x8c, 0xd8, 0x69, 0xb1, 0x65, 0xc9, 0x59, 0xc1,
	0xb2, 0x63, 0xff, 0x55, 0x7b, 0x16, 0xc9, 0x2c, 0x4b, 0x74, 0xc6, 0x85, 0xb6, 0xe7, 0x41, 0x0b,
	0xf0, 0xb3, 0x49, 0xda, 0xae, 0x11, 0x0f, 0xf8, 0xde, 0x90, 0x2b, 0x1d, 0x3c, 0x87, 0x2b, 0x53,
	0xa7, 0x2a, 0x97, 0x42, 0x71, 0xfc, 0x10, 0x9a, 0x36, 0x74, 0x05, 0xb5, 0x51, 0x67, 0xb9, 0xbf,
	0x46, 0x6a, 0x2a, 0x12, 0x6b, 0xdc, 0xbe, 0x70, 0xf8, 0x73, 0xad, 0x31, 0x70, 0xa6, 0xe0, 0x15,
	0xac, 0x9a, 0xd4, 0x1d, 0xae, 0xb7, 0xd2, 0xf4, 0xd1, 0x3f, 0x14, 0x37, 0x15, 0x3f, 0x01, 0x28,
	0xbb, 0xb8, 0x11, 0xeb, 0xc4, 0x16, 0x27, 0x93, 0xe2, 0xc4, 0x5e, 0xaf, 0x2b, 0x4e, 0x76, 0x59,
	0xcc, 0x9d, 0x77, 0x50, 0x71, 0x06, 0xdf, 0x10, 0xdc, 0xa8, 0x19, 0xe4, 0x8a, 0x3c, 0x05, 0x28,
	0x6f, 0x62, 0x05, 0xb5, 0x17, 0x3b, 0xcb, 0xfd, 0x9b, 0xb5, 0x65, 0xca, 0x00, 0x57, 0xa8, 0x62,
	0xc6, 0x3b, 0x53, 0xd0, 0x0b, 0x06, 0xfa, 0xf6, 0xa9, 0xd0, 0x96, 0x63, 0x8a, 0xba, 0x07, 0xd7,
	0x8f, 0xa1, 0x67, 0xaf, 0xa6, 0x05, 0x4b, 0x89, 0x78, 0xc9, 0xf7, 0xcd, 0xad, 0x5c, 0x1c, 0xd8,
	0x2f, 0x41, 0x0c, 0xde, 0x49, 0x96, 0x9a, 0x92, 0xe8, 0xdc, 0x25, 0xfb, 0x7f, 0x16, 0x61, 0xc9,
	0x4c, 0xc2, 0x1f, 0x11, 0x34, 0xed, 0x9f, 0x8b, 0x37, 0x6a, 0xb3, 0x66, 0x37, 0xca, 0xeb, 0x9e,
	0x4d, 0x6c, 0xd1, 0x83, 0xce, 0x87, 0xef, 0xbf, 0x3f, 0x2f, 0x04, 0xb8, 0x4d, 0xad, 0xdc, 0x3d,
	0xee, 0x4e, 0x6c, 0xb4, 0xba, 0xdd, 0xf8, 0x2b, 0x82, 0x4b, 0x25, 0xfa, 0x56, 0x9a, 0xe2, 0xfb,
	0xff, 0x9f, 0x54, 0xb3, 0x7c, 0xde, 0x83, 0x79, 0x6d, 0x0e, 0xb5, 0x6b, 0x50, 0xd7, 0xf1, 0xad,
	0x7a, 0xd4, 0xca, 0xb6, 0x7c, 0x41, 0x00, 0x65, 0x08, 0xee, 0x9f, 0x3a, 0x74, 0x16, 0x74, 0x73,
	0x2e, 0x8f, 0xa3, 0xbc, 0x67, 0x28, 0x09, 0xee, 0x9e, 0x85, 0x92, 0xbe, 0x33, 0xeb, 0xf5, 0x7e,
	0xfb, 0xf1, 0xe1, 0xc8, 0x47, 0x47, 0x23, 0x1f, 0xfd, 0x1a, 0xf9, 0xe8, 0xd3, 0xd8, 0x6f, 0x1c,
	0x8d, 0xfd, 0xc6, 0x8f, 0xb1, 0xdf, 0x78, 0xb1, 0x11, 0x27, 0xfa, 0xf5, 0x30, 0x24, 0x91, 0xcc,
	0x4e, 0x4a, 0xdc, 0x77, 0x99, 0xfa, 0x20, 0xe7, 0x2a, 0x6c, 0x9a, 0x57, 0xcd, 0xe6, 0xdf, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x07, 0xdd, 0x86, 0x5e, 0x24, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	CommitmentAll(ctx context.Context, in *QueryGetAllCommitmentRequest, opts ...grpc.CallOption) (*QueryGetAllCommitmentResponse, error)
	Commitment(ctx context.Context, in *QueryGetCommitmentRequest, opts ...grpc.CallOption) (*QueryGetCommitmentResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/rarimo.rarimocore.cbank.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) CommitmentAll(ctx context.Context, in *QueryGetAllCommitmentRequest, opts ...grpc.CallOption) (*QueryGetAllCommitmentResponse, error) {
	out := new(QueryGetAllCommitmentResponse)
	err := c.cc.Invoke(ctx, "/rarimo.rarimocore.cbank.Query/CommitmentAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Commitment(ctx context.Context, in *QueryGetCommitmentRequest, opts ...grpc.CallOption) (*QueryGetCommitmentResponse, error) {
	out := new(QueryGetCommitmentResponse)
	err := c.cc.Invoke(ctx, "/rarimo.rarimocore.cbank.Query/Commitment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	CommitmentAll(context.Context, *QueryGetAllCommitmentRequest) (*QueryGetAllCommitmentResponse, error)
	Commitment(context.Context, *QueryGetCommitmentRequest) (*QueryGetCommitmentResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) CommitmentAll(ctx context.Context, req *QueryGetAllCommitmentRequest) (*QueryGetAllCommitmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommitmentAll not implemented")
}
func (*UnimplementedQueryServer) Commitment(ctx context.Context, req *QueryGetCommitmentRequest) (*QueryGetCommitmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Commitment not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rarimo.rarimocore.cbank.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_CommitmentAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetAllCommitmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).CommitmentAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rarimo.rarimocore.cbank.Query/CommitmentAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).CommitmentAll(ctx, req.(*QueryGetAllCommitmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Commitment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetCommitmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Commitment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rarimo.rarimocore.cbank.Query/Commitment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Commitment(ctx, req.(*QueryGetCommitmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rarimo.rarimocore.cbank.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "CommitmentAll",
			Handler:    _Query_CommitmentAll_Handler,
		},
		{
			MethodName: "Commitment",
			Handler:    _Query_Commitment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cbank/query.proto",
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryGetAllCommitmentRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetAllCommitmentRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetAllCommitmentRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryGetAllCommitmentResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetAllCommitmentResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetAllCommitmentResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Commitment) > 0 {
		for iNdEx := len(m.Commitment) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Commitment[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *QueryGetCommitmentRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetCommitmentRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetCommitmentRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryGetCommitmentResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetCommitmentResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetCommitmentResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Commitment.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryGetAllCommitmentRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryGetAllCommitmentResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Commitment) > 0 {
		for _, e := range m.Commitment {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryGetCommitmentRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryGetCommitmentResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Commitment.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryGetAllCommitmentRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryGetAllCommitmentRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetAllCommitmentRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryGetAllCommitmentResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryGetAllCommitmentResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetAllCommitmentResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Commitment", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Commitment = append(m.Commitment, Commitment{})
			if err := m.Commitment[len(m.Commitment)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryGetCommitmentRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryGetCommitmentRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetCommitmentRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryGetCommitmentResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryGetCommitmentResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetCommitmentResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Commitment", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Commitment.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
