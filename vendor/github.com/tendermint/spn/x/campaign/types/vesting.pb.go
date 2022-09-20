// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: campaign/vesting.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
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

type ShareVestingOptions struct {
	// Types that are valid to be assigned to Options:
	//	*ShareVestingOptions_DelayedVesting
	Options isShareVestingOptions_Options `protobuf_oneof:"options"`
}

func (m *ShareVestingOptions) Reset()         { *m = ShareVestingOptions{} }
func (m *ShareVestingOptions) String() string { return proto.CompactTextString(m) }
func (*ShareVestingOptions) ProtoMessage()    {}
func (*ShareVestingOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_57113a5fc3a0f84e, []int{0}
}
func (m *ShareVestingOptions) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ShareVestingOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ShareVestingOptions.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ShareVestingOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShareVestingOptions.Merge(m, src)
}
func (m *ShareVestingOptions) XXX_Size() int {
	return m.Size()
}
func (m *ShareVestingOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_ShareVestingOptions.DiscardUnknown(m)
}

var xxx_messageInfo_ShareVestingOptions proto.InternalMessageInfo

type isShareVestingOptions_Options interface {
	isShareVestingOptions_Options()
	MarshalTo([]byte) (int, error)
	Size() int
}

type ShareVestingOptions_DelayedVesting struct {
	DelayedVesting *ShareDelayedVesting `protobuf:"bytes,1,opt,name=delayedVesting,proto3,oneof" json:"delayedVesting,omitempty"`
}

func (*ShareVestingOptions_DelayedVesting) isShareVestingOptions_Options() {}

func (m *ShareVestingOptions) GetOptions() isShareVestingOptions_Options {
	if m != nil {
		return m.Options
	}
	return nil
}

func (m *ShareVestingOptions) GetDelayedVesting() *ShareDelayedVesting {
	if x, ok := m.GetOptions().(*ShareVestingOptions_DelayedVesting); ok {
		return x.DelayedVesting
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ShareVestingOptions) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ShareVestingOptions_DelayedVesting)(nil),
	}
}

// ShareDelayedVesting represents options for share delayed vesting
// Delayed vesting is the type of vesting where all vesting coins are vested
// once end time is reached
type ShareDelayedVesting struct {
	TotalShares Shares `protobuf:"bytes,1,rep,name=totalShares,proto3,casttype=github.com/cosmos/cosmos-sdk/types.Coin,castrepeated=Shares" json:"totalShares"`
	Vesting     Shares `protobuf:"bytes,2,rep,name=vesting,proto3,casttype=github.com/cosmos/cosmos-sdk/types.Coin,castrepeated=Shares" json:"vesting"`
	EndTime     int64  `protobuf:"varint,3,opt,name=endTime,proto3" json:"endTime,omitempty"`
}

func (m *ShareDelayedVesting) Reset()         { *m = ShareDelayedVesting{} }
func (m *ShareDelayedVesting) String() string { return proto.CompactTextString(m) }
func (*ShareDelayedVesting) ProtoMessage()    {}
func (*ShareDelayedVesting) Descriptor() ([]byte, []int) {
	return fileDescriptor_57113a5fc3a0f84e, []int{1}
}
func (m *ShareDelayedVesting) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ShareDelayedVesting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ShareDelayedVesting.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ShareDelayedVesting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShareDelayedVesting.Merge(m, src)
}
func (m *ShareDelayedVesting) XXX_Size() int {
	return m.Size()
}
func (m *ShareDelayedVesting) XXX_DiscardUnknown() {
	xxx_messageInfo_ShareDelayedVesting.DiscardUnknown(m)
}

var xxx_messageInfo_ShareDelayedVesting proto.InternalMessageInfo

func (m *ShareDelayedVesting) GetTotalShares() Shares {
	if m != nil {
		return m.TotalShares
	}
	return nil
}

func (m *ShareDelayedVesting) GetVesting() Shares {
	if m != nil {
		return m.Vesting
	}
	return nil
}

func (m *ShareDelayedVesting) GetEndTime() int64 {
	if m != nil {
		return m.EndTime
	}
	return 0
}

func init() {
	proto.RegisterType((*ShareVestingOptions)(nil), "tendermint.spn.campaign.ShareVestingOptions")
	proto.RegisterType((*ShareDelayedVesting)(nil), "tendermint.spn.campaign.ShareDelayedVesting")
}

func init() { proto.RegisterFile("campaign/vesting.proto", fileDescriptor_57113a5fc3a0f84e) }

var fileDescriptor_57113a5fc3a0f84e = []byte{
	// 337 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x92, 0xb1, 0x4e, 0x02, 0x31,
	0x1c, 0xc6, 0xaf, 0x90, 0x40, 0x2c, 0x89, 0xc3, 0x69, 0xf4, 0x64, 0x28, 0x84, 0x45, 0x62, 0xb4,
	0x0d, 0x38, 0xb9, 0x22, 0x83, 0x9b, 0x09, 0x1a, 0x06, 0xb7, 0xde, 0x5d, 0x73, 0x34, 0x72, 0xfd,
	0x5f, 0x68, 0x25, 0xf0, 0x16, 0x6e, 0xbe, 0x83, 0x4f, 0xc2, 0xc8, 0xe8, 0x84, 0x06, 0xde, 0xc2,
	0xc9, 0xd0, 0x1e, 0x0a, 0x26, 0x8e, 0x4e, 0x77, 0xcd, 0xff, 0xfb, 0xbe, 0x5f, 0xbe, 0xfe, 0x8b,
	0x8f, 0x22, 0x9e, 0x66, 0x5c, 0x26, 0x8a, 0x8d, 0x85, 0x36, 0x52, 0x25, 0x34, 0x1b, 0x81, 0x01,
	0xff, 0xd8, 0x08, 0x15, 0x8b, 0x51, 0x2a, 0x95, 0xa1, 0x3a, 0x53, 0x74, 0x23, 0xab, 0x1e, 0x26,
	0x90, 0x80, 0xd5, 0xb0, 0xf5, 0x9f, 0x93, 0x57, 0x49, 0x04, 0x3a, 0x05, 0xcd, 0x42, 0xae, 0x05,
	0x1b, 0xb7, 0x42, 0x61, 0x78, 0x8b, 0x45, 0x20, 0x95, 0x9b, 0x37, 0x26, 0xf8, 0xe0, 0x6e, 0xc0,
	0x47, 0xa2, 0xef, 0x20, 0xb7, 0x99, 0x91, 0xa0, 0xb4, 0xdf, 0xc7, 0xfb, 0xb1, 0x18, 0xf2, 0xa9,
	0x88, 0xf3, 0x41, 0x80, 0xea, 0xa8, 0x59, 0x69, 0x9f, 0xd3, 0x3f, 0xf0, 0xd4, 0xa6, 0x74, 0x77,
	0x3c, 0x37, 0x5e, 0xef, 0x57, 0x4a, 0x67, 0x0f, 0x97, 0xc1, 0x21, 0x1a, 0x2f, 0x85, 0x1c, 0xbd,
	0x6b, 0xf2, 0xc7, 0xb8, 0x62, 0xc0, 0xf0, 0xa1, 0x9d, 0xe9, 0x00, 0xd5, 0x8b, 0xcd, 0x4a, 0xfb,
	0x84, 0xba, 0x1e, 0x74, 0xdd, 0x83, 0xe6, 0x3d, 0xe8, 0x35, 0x48, 0xd5, 0xb9, 0x9a, 0x2d, 0x6a,
	0xde, 0xe7, 0xa2, 0x76, 0x9a, 0x48, 0x33, 0x78, 0x0a, 0x69, 0x04, 0x29, 0xcb, 0x4b, 0xbb, 0xcf,
	0x85, 0x8e, 0x1f, 0x99, 0x99, 0x66, 0x42, 0x5b, 0xc3, 0xeb, 0x7b, 0xad, 0xe4, 0xb2, 0x7b, 0xdb,
	0x20, 0x5f, 0xe1, 0x72, 0x7e, 0xd3, 0x41, 0xe1, 0x1f, 0x99, 0x1b, 0x88, 0x1f, 0xe0, 0xb2, 0x50,
	0xf1, 0xbd, 0x4c, 0x45, 0x50, 0xac, 0xa3, 0x66, 0xb1, 0xb7, 0x39, 0x76, 0xba, 0xb3, 0x25, 0x41,
	0xf3, 0x25, 0x41, 0x1f, 0x4b, 0x82, 0x9e, 0x57, 0xc4, 0x9b, 0xaf, 0x88, 0xf7, 0xb6, 0x22, 0xde,
	0xc3, 0xd9, 0x16, 0xef, 0x67, 0x11, 0x4c, 0x67, 0x8a, 0x4d, 0xd8, 0xf7, 0x83, 0xb1, 0xdc, 0xb0,
	0x64, 0x17, 0x7c, 0xf9, 0x15, 0x00, 0x00, 0xff, 0xff, 0xa3, 0x51, 0xd2, 0xb3, 0x49, 0x02, 0x00,
	0x00,
}

func (m *ShareVestingOptions) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ShareVestingOptions) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ShareVestingOptions) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Options != nil {
		{
			size := m.Options.Size()
			i -= size
			if _, err := m.Options.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *ShareVestingOptions_DelayedVesting) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ShareVestingOptions_DelayedVesting) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.DelayedVesting != nil {
		{
			size, err := m.DelayedVesting.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintVesting(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *ShareDelayedVesting) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ShareDelayedVesting) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ShareDelayedVesting) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.EndTime != 0 {
		i = encodeVarintVesting(dAtA, i, uint64(m.EndTime))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Vesting) > 0 {
		for iNdEx := len(m.Vesting) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Vesting[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintVesting(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.TotalShares) > 0 {
		for iNdEx := len(m.TotalShares) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TotalShares[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintVesting(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintVesting(dAtA []byte, offset int, v uint64) int {
	offset -= sovVesting(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ShareVestingOptions) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Options != nil {
		n += m.Options.Size()
	}
	return n
}

func (m *ShareVestingOptions_DelayedVesting) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.DelayedVesting != nil {
		l = m.DelayedVesting.Size()
		n += 1 + l + sovVesting(uint64(l))
	}
	return n
}
func (m *ShareDelayedVesting) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.TotalShares) > 0 {
		for _, e := range m.TotalShares {
			l = e.Size()
			n += 1 + l + sovVesting(uint64(l))
		}
	}
	if len(m.Vesting) > 0 {
		for _, e := range m.Vesting {
			l = e.Size()
			n += 1 + l + sovVesting(uint64(l))
		}
	}
	if m.EndTime != 0 {
		n += 1 + sovVesting(uint64(m.EndTime))
	}
	return n
}

func sovVesting(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozVesting(x uint64) (n int) {
	return sovVesting(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ShareVestingOptions) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVesting
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
			return fmt.Errorf("proto: ShareVestingOptions: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ShareVestingOptions: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelayedVesting", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVesting
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
				return ErrInvalidLengthVesting
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVesting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &ShareDelayedVesting{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Options = &ShareVestingOptions_DelayedVesting{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVesting(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVesting
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
func (m *ShareDelayedVesting) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVesting
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
			return fmt.Errorf("proto: ShareDelayedVesting: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ShareDelayedVesting: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalShares", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVesting
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
				return ErrInvalidLengthVesting
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVesting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TotalShares = append(m.TotalShares, github_com_cosmos_cosmos_sdk_types.Coin{})
			if err := m.TotalShares[len(m.TotalShares)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Vesting", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVesting
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
				return ErrInvalidLengthVesting
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVesting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Vesting = append(m.Vesting, github_com_cosmos_cosmos_sdk_types.Coin{})
			if err := m.Vesting[len(m.Vesting)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndTime", wireType)
			}
			m.EndTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVesting
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EndTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipVesting(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVesting
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
func skipVesting(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowVesting
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
					return 0, ErrIntOverflowVesting
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
					return 0, ErrIntOverflowVesting
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
				return 0, ErrInvalidLengthVesting
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupVesting
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthVesting
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthVesting        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowVesting          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupVesting = fmt.Errorf("proto: unexpected end of group")
)
