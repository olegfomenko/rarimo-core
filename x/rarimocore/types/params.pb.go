// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: rarimocore/params.proto

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

type ParamsUpdateType int32

const (
	ParamsUpdateType_CHANGE_SET ParamsUpdateType = 0
	ParamsUpdateType_OTHER      ParamsUpdateType = 1
)

var ParamsUpdateType_name = map[int32]string{
	0: "CHANGE_SET",
	1: "OTHER",
}

var ParamsUpdateType_value = map[string]int32{
	"CHANGE_SET": 0,
	"OTHER":      1,
}

func (x ParamsUpdateType) String() string {
	return proto.EnumName(ParamsUpdateType_name, int32(x))
}

func (ParamsUpdateType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_997c11d1d0285e72, []int{0}
}

type Party struct {
	// PublicKey in hex format
	PubKey string `protobuf:"bytes,1,opt,name=pubKey,proto3" json:"pubKey,omitempty"`
	// Server address connect to
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	// Rarimo core account
	Account  string `protobuf:"bytes,3,opt,name=account,proto3" json:"account,omitempty"`
	Verified bool   `protobuf:"varint,4,opt,name=verified,proto3" json:"verified,omitempty"`
}

func (m *Party) Reset()         { *m = Party{} }
func (m *Party) String() string { return proto.CompactTextString(m) }
func (*Party) ProtoMessage()    {}
func (*Party) Descriptor() ([]byte, []int) {
	return fileDescriptor_997c11d1d0285e72, []int{0}
}
func (m *Party) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Party) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Party.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Party) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Party.Merge(m, src)
}
func (m *Party) XXX_Size() int {
	return m.Size()
}
func (m *Party) XXX_DiscardUnknown() {
	xxx_messageInfo_Party.DiscardUnknown(m)
}

var xxx_messageInfo_Party proto.InternalMessageInfo

func (m *Party) GetPubKey() string {
	if m != nil {
		return m.PubKey
	}
	return ""
}

func (m *Party) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Party) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *Party) GetVerified() bool {
	if m != nil {
		return m.Verified
	}
	return false
}

// Params defines the parameters for the module.
type Params struct {
	KeyECDSA         string   `protobuf:"bytes,1,opt,name=keyECDSA,proto3" json:"keyECDSA,omitempty"`
	Threshold        uint64   `protobuf:"varint,2,opt,name=threshold,proto3" json:"threshold,omitempty"`
	Parties          []*Party `protobuf:"bytes,3,rep,name=parties,proto3" json:"parties,omitempty"`
	IsUpdateRequired bool     `protobuf:"varint,5,opt,name=isUpdateRequired,proto3" json:"isUpdateRequired,omitempty"`
	LastSignature    string   `protobuf:"bytes,6,opt,name=lastSignature,proto3" json:"lastSignature,omitempty"`
	VoteQuorum       string   `protobuf:"bytes,7,opt,name=voteQuorum,proto3" json:"voteQuorum,omitempty"`
	VoteThreshold    string   `protobuf:"bytes,8,opt,name=voteThreshold,proto3" json:"voteThreshold,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_997c11d1d0285e72, []int{1}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetKeyECDSA() string {
	if m != nil {
		return m.KeyECDSA
	}
	return ""
}

func (m *Params) GetThreshold() uint64 {
	if m != nil {
		return m.Threshold
	}
	return 0
}

func (m *Params) GetParties() []*Party {
	if m != nil {
		return m.Parties
	}
	return nil
}

func (m *Params) GetIsUpdateRequired() bool {
	if m != nil {
		return m.IsUpdateRequired
	}
	return false
}

func (m *Params) GetLastSignature() string {
	if m != nil {
		return m.LastSignature
	}
	return ""
}

func (m *Params) GetVoteQuorum() string {
	if m != nil {
		return m.VoteQuorum
	}
	return ""
}

func (m *Params) GetVoteThreshold() string {
	if m != nil {
		return m.VoteThreshold
	}
	return ""
}

func init() {
	proto.RegisterEnum("rarimo.rarimocore.rarimocore.ParamsUpdateType", ParamsUpdateType_name, ParamsUpdateType_value)
	proto.RegisterType((*Party)(nil), "rarimo.rarimocore.rarimocore.Party")
	proto.RegisterType((*Params)(nil), "rarimo.rarimocore.rarimocore.Params")
}

func init() { proto.RegisterFile("rarimocore/params.proto", fileDescriptor_997c11d1d0285e72) }

var fileDescriptor_997c11d1d0285e72 = []byte{
	// 382 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0x4d, 0x6f, 0xda, 0x40,
	0x10, 0xb5, 0xf9, 0x30, 0x30, 0x55, 0x2b, 0x6b, 0x0f, 0xed, 0xaa, 0x42, 0x16, 0xa2, 0x3d, 0x20,
	0x24, 0x4c, 0xd5, 0x9e, 0x7b, 0xa0, 0xd4, 0x2a, 0x6a, 0xa5, 0x96, 0x1a, 0xf7, 0xd2, 0x4b, 0xb4,
	0xd8, 0x1b, 0x58, 0x05, 0xb3, 0xce, 0xee, 0x1a, 0xc5, 0xff, 0x22, 0xd7, 0xfc, 0xa3, 0x1c, 0x39,
	0xe6, 0x18, 0xc1, 0x1f, 0x89, 0x58, 0x9b, 0x2f, 0x45, 0xca, 0xc9, 0x7e, 0xef, 0xcd, 0xce, 0xbc,
	0x19, 0x3d, 0x78, 0x27, 0x88, 0x60, 0x31, 0x0f, 0xb9, 0xa0, 0xfd, 0x84, 0x08, 0x12, 0x4b, 0x37,
	0x11, 0x5c, 0x71, 0xd4, 0xcc, 0x05, 0xf7, 0xa8, 0x9f, 0xfc, 0xb6, 0x39, 0x54, 0xc7, 0x44, 0xa8,
	0x0c, 0xbd, 0x05, 0x2b, 0x49, 0xa7, 0xbf, 0x68, 0x86, 0xcd, 0x96, 0xd9, 0x69, 0xf8, 0x05, 0x42,
	0x18, 0x6a, 0x24, 0x8a, 0x04, 0x95, 0x12, 0x97, 0xb4, 0xb0, 0x87, 0x5a, 0x09, 0x43, 0x9e, 0x2e,
	0x15, 0x2e, 0x17, 0x4a, 0x0e, 0xd1, 0x7b, 0xa8, 0xaf, 0xa8, 0x60, 0x97, 0x8c, 0x46, 0xb8, 0xd2,
	0x32, 0x3b, 0x75, 0xff, 0x80, 0xdb, 0x77, 0x25, 0xb0, 0xc6, 0xda, 0xdf, 0xae, 0xec, 0x8a, 0x66,
	0xde, 0xf0, 0xfb, 0x64, 0x50, 0x0c, 0x3d, 0x60, 0xd4, 0x84, 0x86, 0x9a, 0x0b, 0x2a, 0xe7, 0x7c,
	0x11, 0xe9, 0xc1, 0x15, 0xff, 0x48, 0xa0, 0xaf, 0x50, 0x4b, 0x88, 0x50, 0x8c, 0x4a, 0x5c, 0x6e,
	0x95, 0x3b, 0xaf, 0x3e, 0x7f, 0x70, 0x5f, 0xda, 0xd2, 0xd5, 0x2b, 0xfa, 0xfb, 0x37, 0xa8, 0x0b,
	0x36, 0x93, 0xff, 0x92, 0x88, 0x28, 0xea, 0xd3, 0xeb, 0x94, 0x09, 0x1a, 0xe1, 0xaa, 0xf6, 0xf9,
	0x8c, 0x47, 0x1f, 0xe1, 0xf5, 0x82, 0x48, 0x35, 0x61, 0xb3, 0x25, 0x51, 0xa9, 0xa0, 0xd8, 0xd2,
	0x4e, 0xcf, 0x49, 0xe4, 0x00, 0xac, 0xb8, 0xa2, 0x7f, 0x53, 0x2e, 0xd2, 0x18, 0xd7, 0x74, 0xc9,
	0x09, 0xb3, 0xeb, 0xb2, 0x43, 0xc1, 0x61, 0xa5, 0x7a, 0xde, 0xe5, 0x8c, 0xec, 0xf6, 0xc0, 0xce,
	0x4f, 0x93, 0x7b, 0x08, 0xb2, 0x84, 0xa2, 0x37, 0x00, 0xc3, 0xd1, 0xe0, 0xf7, 0x0f, 0xef, 0x62,
	0xe2, 0x05, 0xb6, 0x81, 0x1a, 0x50, 0xfd, 0x13, 0x8c, 0x3c, 0xdf, 0x36, 0xbf, 0xfd, 0xbc, 0xdf,
	0x38, 0xe6, 0x7a, 0xe3, 0x98, 0x8f, 0x1b, 0xc7, 0xbc, 0xdd, 0x3a, 0xc6, 0x7a, 0xeb, 0x18, 0x0f,
	0x5b, 0xc7, 0xf8, 0xff, 0x69, 0xc6, 0xd4, 0x82, 0x4c, 0xdd, 0x90, 0xc7, 0xfd, 0xfc, 0x0c, 0xc5,
	0xa7, 0xa7, 0xf3, 0x71, 0xd3, 0x3f, 0x09, 0x8b, 0xca, 0x12, 0x2a, 0xa7, 0x96, 0x0e, 0xcb, 0x97,
	0xa7, 0x00, 0x00, 0x00, 0xff, 0xff, 0x04, 0x28, 0x3d, 0xa8, 0x47, 0x02, 0x00, 0x00,
}

func (m *Party) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Party) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Party) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Verified {
		i--
		if m.Verified {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if len(m.Account) > 0 {
		i -= len(m.Account)
		copy(dAtA[i:], m.Account)
		i = encodeVarintParams(dAtA, i, uint64(len(m.Account)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintParams(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.PubKey) > 0 {
		i -= len(m.PubKey)
		copy(dAtA[i:], m.PubKey)
		i = encodeVarintParams(dAtA, i, uint64(len(m.PubKey)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.VoteThreshold) > 0 {
		i -= len(m.VoteThreshold)
		copy(dAtA[i:], m.VoteThreshold)
		i = encodeVarintParams(dAtA, i, uint64(len(m.VoteThreshold)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.VoteQuorum) > 0 {
		i -= len(m.VoteQuorum)
		copy(dAtA[i:], m.VoteQuorum)
		i = encodeVarintParams(dAtA, i, uint64(len(m.VoteQuorum)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.LastSignature) > 0 {
		i -= len(m.LastSignature)
		copy(dAtA[i:], m.LastSignature)
		i = encodeVarintParams(dAtA, i, uint64(len(m.LastSignature)))
		i--
		dAtA[i] = 0x32
	}
	if m.IsUpdateRequired {
		i--
		if m.IsUpdateRequired {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if len(m.Parties) > 0 {
		for iNdEx := len(m.Parties) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Parties[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.Threshold != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.Threshold))
		i--
		dAtA[i] = 0x10
	}
	if len(m.KeyECDSA) > 0 {
		i -= len(m.KeyECDSA)
		copy(dAtA[i:], m.KeyECDSA)
		i = encodeVarintParams(dAtA, i, uint64(len(m.KeyECDSA)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Party) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PubKey)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.Account)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	if m.Verified {
		n += 2
	}
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.KeyECDSA)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	if m.Threshold != 0 {
		n += 1 + sovParams(uint64(m.Threshold))
	}
	if len(m.Parties) > 0 {
		for _, e := range m.Parties {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	if m.IsUpdateRequired {
		n += 2
	}
	l = len(m.LastSignature)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.VoteQuorum)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.VoteThreshold)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Party) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Party: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Party: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PubKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PubKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Account", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Account = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Verified", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Verified = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeyECDSA", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.KeyECDSA = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Threshold", wireType)
			}
			m.Threshold = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Threshold |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Parties", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Parties = append(m.Parties, &Party{})
			if err := m.Parties[len(m.Parties)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsUpdateRequired", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsUpdateRequired = bool(v != 0)
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastSignature", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LastSignature = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VoteQuorum", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VoteQuorum = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VoteThreshold", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VoteThreshold = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
