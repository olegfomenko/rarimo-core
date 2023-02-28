// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: rarimocore/vote.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	proto "github.com/gogo/protobuf/proto"
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

type VoteType int32

const (
	VoteType_YES VoteType = 0
	VoteType_NO  VoteType = 1
)

var VoteType_name = map[int32]string{
	0: "YES",
	1: "NO",
}

var VoteType_value = map[string]int32{
	"YES": 0,
	"NO":  1,
}

func (x VoteType) String() string {
	return proto.EnumName(VoteType_name, int32(x))
}

func (VoteType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b638e6b4bc245bc3, []int{0}
}

type VoteIndex struct {
	Operation string `protobuf:"bytes,1,opt,name=operation,proto3" json:"operation,omitempty"`
	Validator string `protobuf:"bytes,2,opt,name=validator,proto3" json:"validator,omitempty"`
}

func (m *VoteIndex) Reset()         { *m = VoteIndex{} }
func (m *VoteIndex) String() string { return proto.CompactTextString(m) }
func (*VoteIndex) ProtoMessage()    {}
func (*VoteIndex) Descriptor() ([]byte, []int) {
	return fileDescriptor_b638e6b4bc245bc3, []int{0}
}
func (m *VoteIndex) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VoteIndex) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VoteIndex.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *VoteIndex) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VoteIndex.Merge(m, src)
}
func (m *VoteIndex) XXX_Size() int {
	return m.Size()
}
func (m *VoteIndex) XXX_DiscardUnknown() {
	xxx_messageInfo_VoteIndex.DiscardUnknown(m)
}

var xxx_messageInfo_VoteIndex proto.InternalMessageInfo

func (m *VoteIndex) GetOperation() string {
	if m != nil {
		return m.Operation
	}
	return ""
}

func (m *VoteIndex) GetValidator() string {
	if m != nil {
		return m.Validator
	}
	return ""
}

type Vote struct {
	Index *VoteIndex `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	Vote  VoteType   `protobuf:"varint,2,opt,name=vote,proto3,enum=rarimo.rarimocore.rarimocore.VoteType" json:"vote,omitempty"`
}

func (m *Vote) Reset()         { *m = Vote{} }
func (m *Vote) String() string { return proto.CompactTextString(m) }
func (*Vote) ProtoMessage()    {}
func (*Vote) Descriptor() ([]byte, []int) {
	return fileDescriptor_b638e6b4bc245bc3, []int{1}
}
func (m *Vote) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Vote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Vote.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Vote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vote.Merge(m, src)
}
func (m *Vote) XXX_Size() int {
	return m.Size()
}
func (m *Vote) XXX_DiscardUnknown() {
	xxx_messageInfo_Vote.DiscardUnknown(m)
}

var xxx_messageInfo_Vote proto.InternalMessageInfo

func (m *Vote) GetIndex() *VoteIndex {
	if m != nil {
		return m.Index
	}
	return nil
}

func (m *Vote) GetVote() VoteType {
	if m != nil {
		return m.Vote
	}
	return VoteType_YES
}

func init() {
	proto.RegisterEnum("rarimo.rarimocore.rarimocore.VoteType", VoteType_name, VoteType_value)
	proto.RegisterType((*VoteIndex)(nil), "rarimo.rarimocore.rarimocore.VoteIndex")
	proto.RegisterType((*Vote)(nil), "rarimo.rarimocore.rarimocore.Vote")
}

func init() { proto.RegisterFile("rarimocore/vote.proto", fileDescriptor_b638e6b4bc245bc3) }

var fileDescriptor_b638e6b4bc245bc3 = []byte{
	// 262 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2d, 0x4a, 0x2c, 0xca,
	0xcc, 0xcd, 0x4f, 0xce, 0x2f, 0x4a, 0xd5, 0x2f, 0xcb, 0x2f, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x92, 0x81, 0x08, 0xeb, 0x21, 0x64, 0x91, 0x98, 0x52, 0x92, 0xe9, 0xf9, 0xf9, 0xe9,
	0x39, 0xa9, 0xfa, 0x60, 0xb5, 0x49, 0xa5, 0x69, 0xfa, 0x89, 0x79, 0x95, 0x10, 0x8d, 0x4a, 0xee,
	0x5c, 0x9c, 0x61, 0xf9, 0x25, 0xa9, 0x9e, 0x79, 0x29, 0xa9, 0x15, 0x42, 0x32, 0x5c, 0x9c, 0xf9,
	0x05, 0xa9, 0x45, 0x89, 0x25, 0x99, 0xf9, 0x79, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x08,
	0x01, 0x90, 0x6c, 0x59, 0x62, 0x4e, 0x66, 0x4a, 0x62, 0x49, 0x7e, 0x91, 0x04, 0x13, 0x44, 0x16,
	0x2e, 0xa0, 0xd4, 0xc8, 0xc8, 0xc5, 0x02, 0x32, 0x49, 0xc8, 0x96, 0x8b, 0x35, 0x13, 0x64, 0x1a,
	0xd8, 0x00, 0x6e, 0x23, 0x75, 0x3d, 0x7c, 0x4e, 0xd3, 0x83, 0x5b, 0x1e, 0x04, 0xd1, 0x25, 0x64,
	0xc5, 0xc5, 0x02, 0xf2, 0x17, 0xd8, 0x02, 0x3e, 0x23, 0x35, 0xc2, 0xba, 0x43, 0x2a, 0x0b, 0x52,
	0x83, 0xc0, 0x7a, 0xb4, 0xa4, 0xb9, 0x38, 0x60, 0x22, 0x42, 0xec, 0x5c, 0xcc, 0x91, 0xae, 0xc1,
	0x02, 0x0c, 0x42, 0x6c, 0x5c, 0x4c, 0x7e, 0xfe, 0x02, 0x8c, 0x4e, 0x5e, 0x27, 0x1e, 0xc9, 0x31,
	0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb,
	0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0x65, 0x90, 0x9e, 0x59, 0x92, 0x93, 0x98, 0xa4, 0x97, 0x9c,
	0x9f, 0xab, 0x0f, 0x31, 0x1c, 0x4a, 0xe9, 0x82, 0x83, 0xb9, 0x42, 0x1f, 0x29, 0xcc, 0x4b, 0x2a,
	0x0b, 0x52, 0x8b, 0x93, 0xd8, 0xc0, 0x81, 0x67, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xa5, 0x5f,
	0xdc, 0x8c, 0x8e, 0x01, 0x00, 0x00,
}

func (m *VoteIndex) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VoteIndex) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VoteIndex) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Validator) > 0 {
		i -= len(m.Validator)
		copy(dAtA[i:], m.Validator)
		i = encodeVarintVote(dAtA, i, uint64(len(m.Validator)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Operation) > 0 {
		i -= len(m.Operation)
		copy(dAtA[i:], m.Operation)
		i = encodeVarintVote(dAtA, i, uint64(len(m.Operation)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Vote) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Vote) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Vote) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Vote != 0 {
		i = encodeVarintVote(dAtA, i, uint64(m.Vote))
		i--
		dAtA[i] = 0x10
	}
	if m.Index != nil {
		{
			size, err := m.Index.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintVote(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintVote(dAtA []byte, offset int, v uint64) int {
	offset -= sovVote(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *VoteIndex) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Operation)
	if l > 0 {
		n += 1 + l + sovVote(uint64(l))
	}
	l = len(m.Validator)
	if l > 0 {
		n += 1 + l + sovVote(uint64(l))
	}
	return n
}

func (m *Vote) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Index != nil {
		l = m.Index.Size()
		n += 1 + l + sovVote(uint64(l))
	}
	if m.Vote != 0 {
		n += 1 + sovVote(uint64(m.Vote))
	}
	return n
}

func sovVote(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozVote(x uint64) (n int) {
	return sovVote(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *VoteIndex) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVote
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
			return fmt.Errorf("proto: VoteIndex: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VoteIndex: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Operation", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVote
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
				return ErrInvalidLengthVote
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVote
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Operation = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Validator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVote
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
				return ErrInvalidLengthVote
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVote
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Validator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVote(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVote
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
func (m *Vote) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVote
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
			return fmt.Errorf("proto: Vote: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Vote: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVote
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
				return ErrInvalidLengthVote
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVote
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Index == nil {
				m.Index = &VoteIndex{}
			}
			if err := m.Index.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Vote", wireType)
			}
			m.Vote = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVote
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Vote |= VoteType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipVote(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVote
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
func skipVote(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowVote
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
					return 0, ErrIntOverflowVote
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
					return 0, ErrIntOverflowVote
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
				return 0, ErrInvalidLengthVote
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupVote
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthVote
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthVote        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowVote          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupVote = fmt.Errorf("proto: unexpected end of group")
)
