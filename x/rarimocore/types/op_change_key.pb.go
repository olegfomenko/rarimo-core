// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: rarimocore/op_change_key.proto

package types

import (
	fmt "fmt"
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

type ChangeKey struct {
	NewKey string `protobuf:"bytes,1,opt,name=newKey,proto3" json:"newKey,omitempty"`
}

func (m *ChangeKey) Reset()         { *m = ChangeKey{} }
func (m *ChangeKey) String() string { return proto.CompactTextString(m) }
func (*ChangeKey) ProtoMessage()    {}
func (*ChangeKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_2641444419e16034, []int{0}
}
func (m *ChangeKey) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ChangeKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ChangeKey.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ChangeKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangeKey.Merge(m, src)
}
func (m *ChangeKey) XXX_Size() int {
	return m.Size()
}
func (m *ChangeKey) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangeKey.DiscardUnknown(m)
}

var xxx_messageInfo_ChangeKey proto.InternalMessageInfo

func (m *ChangeKey) GetNewKey() string {
	if m != nil {
		return m.NewKey
	}
	return ""
}

func init() {
	proto.RegisterType((*ChangeKey)(nil), "rarifyprotocol.rarimocore.rarimocore.ChangeKey")
}

func init() { proto.RegisterFile("rarimocore/op_change_key.proto", fileDescriptor_2641444419e16034) }

var fileDescriptor_2641444419e16034 = []byte{
	// 167 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2b, 0x4a, 0x2c, 0xca,
	0xcc, 0xcd, 0x4f, 0xce, 0x2f, 0x4a, 0xd5, 0xcf, 0x2f, 0x88, 0x4f, 0xce, 0x48, 0xcc, 0x4b, 0x4f,
	0x8d, 0xcf, 0x4e, 0xad, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x52, 0x01, 0xc9, 0xa7, 0x55,
	0x82, 0x39, 0xc9, 0xf9, 0x39, 0x7a, 0x08, 0xe5, 0x48, 0x4c, 0x25, 0x65, 0x2e, 0x4e, 0x67, 0xb0,
	0x4e, 0xef, 0xd4, 0x4a, 0x21, 0x31, 0x2e, 0xb6, 0xbc, 0xd4, 0x72, 0xef, 0xd4, 0x4a, 0x09, 0x46,
	0x05, 0x46, 0x0d, 0xce, 0x20, 0x28, 0xcf, 0x29, 0xf8, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4,
	0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f,
	0xe5, 0x18, 0xa2, 0x2c, 0xd3, 0x33, 0x4b, 0x72, 0x12, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x21,
	0xf6, 0xe9, 0xc2, 0x2c, 0xd4, 0x87, 0xd8, 0xa2, 0x0b, 0x76, 0x60, 0x85, 0x3e, 0x92, 0x6b, 0x4b,
	0x2a, 0x0b, 0x52, 0x8b, 0x93, 0xd8, 0xc0, 0x0a, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x27,
	0xd0, 0xb8, 0x8c, 0xc8, 0x00, 0x00, 0x00,
}

func (m *ChangeKey) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ChangeKey) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ChangeKey) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.NewKey) > 0 {
		i -= len(m.NewKey)
		copy(dAtA[i:], m.NewKey)
		i = encodeVarintOpChangeKey(dAtA, i, uint64(len(m.NewKey)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintOpChangeKey(dAtA []byte, offset int, v uint64) int {
	offset -= sovOpChangeKey(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ChangeKey) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.NewKey)
	if l > 0 {
		n += 1 + l + sovOpChangeKey(uint64(l))
	}
	return n
}

func sovOpChangeKey(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozOpChangeKey(x uint64) (n int) {
	return sovOpChangeKey(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ChangeKey) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOpChangeKey
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
			return fmt.Errorf("proto: ChangeKey: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ChangeKey: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NewKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOpChangeKey
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
				return ErrInvalidLengthOpChangeKey
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOpChangeKey
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NewKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOpChangeKey(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOpChangeKey
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
func skipOpChangeKey(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOpChangeKey
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
					return 0, ErrIntOverflowOpChangeKey
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
					return 0, ErrIntOverflowOpChangeKey
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
				return 0, ErrInvalidLengthOpChangeKey
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupOpChangeKey
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthOpChangeKey
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthOpChangeKey        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOpChangeKey          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupOpChangeKey = fmt.Errorf("proto: unexpected end of group")
)
