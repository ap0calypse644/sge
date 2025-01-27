// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sge/sportevent/params.proto

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

// Params defines the parameters for the module.
// contains bet constraints associated to a sport event
type Params struct {
	// this would be required to set a default overall max bet capacity for sport event
	EventMinBetAmount github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,1,opt,name=event_min_bet_amount,json=eventMinBetAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"event_min_bet_amount" yaml:"event_min_bet_amount"`
	// this would be required to set a default overall max bet capacity for sport event
	EventMaxBetCap github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=event_max_bet_cap,json=eventMaxBetCap,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"event_max_bet_cap" yaml:"event_max_bet_cap"`
	// this would be required to set a default min bet fee for a sport event
	EventMinBetFee github_com_cosmos_cosmos_sdk_types.Coin `protobuf:"bytes,3,opt,name=event_min_bet_fee,json=eventMinBetFee,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Coin" json:"event_min_bet_fee" yaml:"event_min_bet_fee"`
	EventMaxLoss   github_com_cosmos_cosmos_sdk_types.Int  `protobuf:"bytes,4,opt,name=event_max_loss,json=eventMaxLoss,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"event_max_loss" yaml:"event_max_loss"`
	EventMaxVig    github_com_cosmos_cosmos_sdk_types.Dec  `protobuf:"bytes,5,opt,name=event_max_vig,json=eventMaxVig,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"event_max_vig" yaml:"event_max_vig"`
	EventMinVig    github_com_cosmos_cosmos_sdk_types.Dec  `protobuf:"bytes,6,opt,name=event_min_vig,json=eventMinVig,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"event_min_vig" yaml:"event_min_vig"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_60791a2b28c3b668, []int{0}
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

func init() {
	proto.RegisterType((*Params)(nil), "sgenetwork.sge.sportevent.Params")
}

func init() { proto.RegisterFile("sge/sportevent/params.proto", fileDescriptor_60791a2b28c3b668) }

var fileDescriptor_60791a2b28c3b668 = []byte{
	// 422 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0xd3, 0x3f, 0xef, 0xd2, 0x40,
	0x18, 0x07, 0xf0, 0x56, 0x7f, 0x3f, 0x12, 0xab, 0x92, 0xd0, 0x60, 0x52, 0x21, 0x69, 0x4d, 0x07,
	0x75, 0xa1, 0x17, 0xe2, 0xc6, 0x66, 0x31, 0x10, 0xff, 0x90, 0x98, 0x0e, 0x0c, 0x2e, 0xe4, 0x5a,
	0x1f, 0xcf, 0x13, 0x7a, 0xd7, 0xf4, 0x8e, 0x0a, 0x8b, 0xaf, 0xc1, 0xd1, 0xd1, 0x17, 0xe3, 0xc0,
	0xc8, 0x68, 0x1c, 0x1a, 0x03, 0xef, 0x80, 0x57, 0x60, 0xb8, 0x16, 0x8a, 0xe8, 0x20, 0x89, 0x13,
	0x97, 0xe3, 0xdb, 0xef, 0xe7, 0x9e, 0xe1, 0x31, 0xda, 0x82, 0x00, 0x12, 0x09, 0x4f, 0x25, 0x64,
	0xc0, 0x24, 0x4a, 0x70, 0x8a, 0x63, 0xe1, 0x25, 0x29, 0x97, 0xdc, 0xbc, 0x2f, 0x08, 0x30, 0x90,
	0x1f, 0x79, 0x3a, 0xf5, 0x04, 0x01, 0xaf, 0xca, 0xb5, 0x9a, 0x84, 0x13, 0xae, 0x52, 0x68, 0x7f,
	0x2a, 0x3e, 0x68, 0xd9, 0x11, 0x17, 0x31, 0x17, 0x28, 0xc4, 0x02, 0x50, 0xd6, 0x0d, 0x41, 0xe2,
	0x2e, 0x8a, 0x38, 0x65, 0xc5, 0xff, 0xee, 0xb7, 0x6b, 0xa3, 0xf6, 0x5a, 0x09, 0xe6, 0x27, 0xa3,
	0xa9, 0x9a, 0x26, 0x31, 0x65, 0x93, 0x10, 0xe4, 0x04, 0xc7, 0x7c, 0xce, 0xa4, 0xa5, 0x3f, 0xd0,
	0x1f, 0xdf, 0xf2, 0x47, 0xab, 0xdc, 0xd1, 0x7e, 0xe4, 0xce, 0x43, 0x42, 0xe5, 0xfb, 0x79, 0xe8,
	0x45, 0x3c, 0x46, 0x65, 0x77, 0xf1, 0xd3, 0x11, 0x6f, 0xa7, 0x48, 0x2e, 0x13, 0x10, 0xde, 0x73,
	0x26, 0x77, 0xb9, 0xd3, 0x5e, 0xe2, 0x78, 0xd6, 0x73, 0xff, 0xd6, 0xe9, 0x06, 0x0d, 0x75, 0x3d,
	0xa2, 0xcc, 0x07, 0xf9, 0x54, 0xdd, 0x99, 0x73, 0xa3, 0x51, 0x66, 0xf1, 0x42, 0x65, 0x23, 0x9c,
	0x58, 0x37, 0x14, 0xfe, 0xe2, 0x62, 0xdc, 0xfa, 0x0d, 0xaf, 0x0a, 0xdd, 0xa0, 0x5e, 0xc8, 0x78,
	0xe1, 0x83, 0xec, 0xe3, 0xc4, 0xcc, 0x8e, 0x6c, 0xf9, 0xc4, 0x77, 0x00, 0xd6, 0x4d, 0xc5, 0xbe,
	0x2c, 0xd9, 0x47, 0xff, 0xc0, 0xf6, 0x39, 0x65, 0x7f, 0xb8, 0x55, 0xe3, 0xd1, 0x55, 0x13, 0x0f,
	0x00, 0xcc, 0xd8, 0xa8, 0x57, 0xaf, 0x9b, 0x71, 0x21, 0xac, 0x2b, 0x85, 0x0e, 0x2f, 0x9e, 0xf5,
	0xde, 0xf9, 0xac, 0xfb, 0x36, 0x37, 0xb8, 0x73, 0x18, 0xf4, 0x15, 0x17, 0xc2, 0xfc, 0x60, 0xdc,
	0xad, 0x02, 0x19, 0x25, 0xd6, 0xb5, 0xd2, 0x06, 0x17, 0x68, 0xcf, 0x20, 0xda, 0xe5, 0x4e, 0xf3,
	0x5c, 0xcb, 0x28, 0x71, 0x83, 0xdb, 0x07, 0x6c, 0x4c, 0xc9, 0x89, 0x45, 0x99, 0xb2, 0x6a, 0xff,
	0xc5, 0x2a, 0xca, 0x8e, 0x16, 0x65, 0x63, 0x4a, 0x7a, 0x57, 0x5f, 0xbe, 0x3a, 0x9a, 0x3f, 0x5c,
	0x6d, 0x6c, 0x7d, 0xbd, 0xb1, 0xf5, 0x9f, 0x1b, 0x5b, 0xff, 0xbc, 0xb5, 0xb5, 0xf5, 0xd6, 0xd6,
	0xbe, 0x6f, 0x6d, 0xed, 0x4d, 0xe7, 0x04, 0x13, 0x04, 0x3a, 0xe5, 0xf6, 0xec, 0xcf, 0x68, 0x71,
	0xba, 0x67, 0xca, 0x0d, 0x6b, 0x6a, 0x2d, 0x9e, 0xfc, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x1a, 0x57,
	0x22, 0xbf, 0x86, 0x03, 0x00, 0x00,
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
	{
		size := m.EventMinVig.Size()
		i -= size
		if _, err := m.EventMinVig.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size := m.EventMaxVig.Size()
		i -= size
		if _, err := m.EventMaxVig.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.EventMaxLoss.Size()
		i -= size
		if _, err := m.EventMaxLoss.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.EventMinBetFee.Size()
		i -= size
		if _, err := m.EventMinBetFee.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.EventMaxBetCap.Size()
		i -= size
		if _, err := m.EventMaxBetCap.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.EventMinBetAmount.Size()
		i -= size
		if _, err := m.EventMinBetAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
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
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.EventMinBetAmount.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.EventMaxBetCap.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.EventMinBetFee.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.EventMaxLoss.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.EventMaxVig.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.EventMinVig.Size()
	n += 1 + l + sovParams(uint64(l))
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
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
				return fmt.Errorf("proto: wrong wireType = %d for field EventMinBetAmount", wireType)
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
			if err := m.EventMinBetAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EventMaxBetCap", wireType)
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
			if err := m.EventMaxBetCap.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EventMinBetFee", wireType)
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
			if err := m.EventMinBetFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EventMaxLoss", wireType)
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
			if err := m.EventMaxLoss.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EventMaxVig", wireType)
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
			if err := m.EventMaxVig.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EventMinVig", wireType)
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
			if err := m.EventMinVig.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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
