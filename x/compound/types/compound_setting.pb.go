// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: temporal/compound/compound_setting.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

type CompoundSetting struct {
	Delegator        string              `protobuf:"bytes,1,opt,name=delegator,proto3" json:"delegator,omitempty"`
	ValidatorSetting []*ValidatorSetting `protobuf:"bytes,2,rep,name=validatorSetting,proto3" json:"validatorSetting,omitempty"`
	AmountToRemain   types.Coin          `protobuf:"bytes,3,opt,name=amountToRemain,proto3" json:"amountToRemain"`
	Frequency        uint64              `protobuf:"varint,4,opt,name=frequency,proto3" json:"frequency,omitempty"`
}

func (m *CompoundSetting) Reset()         { *m = CompoundSetting{} }
func (m *CompoundSetting) String() string { return proto.CompactTextString(m) }
func (*CompoundSetting) ProtoMessage()    {}
func (*CompoundSetting) Descriptor() ([]byte, []int) {
	return fileDescriptor_43b8cac02d87dd16, []int{0}
}
func (m *CompoundSetting) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CompoundSetting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CompoundSetting.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CompoundSetting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CompoundSetting.Merge(m, src)
}
func (m *CompoundSetting) XXX_Size() int {
	return m.Size()
}
func (m *CompoundSetting) XXX_DiscardUnknown() {
	xxx_messageInfo_CompoundSetting.DiscardUnknown(m)
}

var xxx_messageInfo_CompoundSetting proto.InternalMessageInfo

func (m *CompoundSetting) GetDelegator() string {
	if m != nil {
		return m.Delegator
	}
	return ""
}

func (m *CompoundSetting) GetValidatorSetting() []*ValidatorSetting {
	if m != nil {
		return m.ValidatorSetting
	}
	return nil
}

func (m *CompoundSetting) GetAmountToRemain() types.Coin {
	if m != nil {
		return m.AmountToRemain
	}
	return types.Coin{}
}

func (m *CompoundSetting) GetFrequency() uint64 {
	if m != nil {
		return m.Frequency
	}
	return 0
}

func init() {
	proto.RegisterType((*CompoundSetting)(nil), "temporal.compound.CompoundSetting")
}

func init() {
	proto.RegisterFile("temporal/compound/compound_setting.proto", fileDescriptor_43b8cac02d87dd16)
}

var fileDescriptor_43b8cac02d87dd16 = []byte{
	// 341 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0xcf, 0x4e, 0xf2, 0x40,
	0x14, 0xc5, 0x5b, 0x20, 0x5f, 0x42, 0x49, 0x3e, 0xb5, 0x61, 0x51, 0x88, 0x19, 0x1b, 0xdd, 0xd4,
	0x05, 0x33, 0x01, 0x13, 0xf7, 0xc2, 0xc2, 0x8d, 0x89, 0x49, 0x31, 0x2e, 0xdc, 0x90, 0x69, 0x3b,
	0xd6, 0x49, 0xe8, 0xdc, 0xda, 0x99, 0x12, 0x71, 0xe7, 0x1b, 0xf8, 0x30, 0x3e, 0x04, 0x4b, 0xe2,
	0xca, 0x95, 0x31, 0xf0, 0x22, 0xa6, 0x74, 0x80, 0x08, 0xbb, 0xfb, 0xe7, 0xfc, 0x7a, 0xcf, 0xe9,
	0x58, 0x9e, 0x62, 0x49, 0x0a, 0x19, 0x1d, 0x93, 0x10, 0x92, 0x14, 0x72, 0x11, 0x6d, 0x8a, 0x91,
	0x64, 0x4a, 0x71, 0x11, 0xe3, 0x34, 0x03, 0x05, 0xf6, 0xd1, 0x5a, 0x89, 0xd7, 0x82, 0xf6, 0xf9,
	0x3e, 0x3c, 0xa1, 0x63, 0x1e, 0x51, 0x05, 0xd9, 0x5f, 0xba, 0xdd, 0x8c, 0x21, 0x86, 0x55, 0x49,
	0x8a, 0x4a, 0x4f, 0x51, 0x08, 0x32, 0x01, 0x49, 0x02, 0x2a, 0x19, 0x99, 0x74, 0x03, 0xa6, 0x68,
	0x97, 0x84, 0xc0, 0x85, 0xde, 0xb7, 0xca, 0xfd, 0xa8, 0x04, 0xcb, 0xa6, 0x5c, 0x9d, 0xbe, 0x55,
	0xac, 0x83, 0x81, 0xbe, 0x3a, 0x2c, 0x4f, 0xd9, 0x97, 0x56, 0x3d, 0x62, 0x63, 0x16, 0x17, 0xf7,
	0x1d, 0xd3, 0x35, 0xbd, 0x7a, 0xdf, 0xf9, 0xfc, 0xe8, 0x34, 0x35, 0x78, 0x15, 0x45, 0x19, 0x93,
	0x72, 0xa8, 0x32, 0x2e, 0x62, 0x7f, 0x2b, 0xb5, 0x6f, 0xad, 0xc3, 0x8d, 0x6f, 0xfd, 0x2d, 0xa7,
	0xe2, 0x56, 0xbd, 0x46, 0xef, 0x0c, 0xef, 0xa5, 0xc6, 0xf7, 0x3b, 0x52, 0x7f, 0x0f, 0xb6, 0xaf,
	0xad, 0xff, 0x34, 0x81, 0x5c, 0xa8, 0x3b, 0xf0, 0x59, 0x42, 0xb9, 0x70, 0xaa, 0xae, 0xe9, 0x35,
	0x7a, 0x2d, 0xac, 0xad, 0x14, 0x81, 0xb1, 0x0e, 0x8c, 0x07, 0xc0, 0x45, 0xbf, 0x36, 0xfb, 0x3e,
	0x31, 0xfc, 0x1d, 0xcc, 0x3e, 0xb6, 0xea, 0x8f, 0x19, 0x7b, 0xce, 0x99, 0x08, 0xa7, 0x4e, 0xcd,
	0x35, 0xbd, 0x9a, 0xbf, 0x1d, 0xf4, 0x6f, 0x66, 0x0b, 0x64, 0xce, 0x17, 0xc8, 0xfc, 0x59, 0x20,
	0xf3, 0x7d, 0x89, 0x8c, 0xf9, 0x12, 0x19, 0x5f, 0x4b, 0x64, 0x3c, 0xf4, 0x62, 0xae, 0x9e, 0xf2,
	0xa0, 0x30, 0x4d, 0xd6, 0x09, 0x3a, 0xaf, 0x20, 0xd8, 0xa6, 0x23, 0x2f, 0xdb, 0x47, 0x53, 0xd3,
	0x94, 0xc9, 0xe0, 0xdf, 0xea, 0xc7, 0x5e, 0xfc, 0x06, 0x00, 0x00, 0xff, 0xff, 0xe2, 0x60, 0x19,
	0xfd, 0x13, 0x02, 0x00, 0x00,
}

func (m *CompoundSetting) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CompoundSetting) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CompoundSetting) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Frequency != 0 {
		i = encodeVarintCompoundSetting(dAtA, i, uint64(m.Frequency))
		i--
		dAtA[i] = 0x20
	}
	{
		size, err := m.AmountToRemain.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintCompoundSetting(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.ValidatorSetting) > 0 {
		for iNdEx := len(m.ValidatorSetting) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ValidatorSetting[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCompoundSetting(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Delegator) > 0 {
		i -= len(m.Delegator)
		copy(dAtA[i:], m.Delegator)
		i = encodeVarintCompoundSetting(dAtA, i, uint64(len(m.Delegator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintCompoundSetting(dAtA []byte, offset int, v uint64) int {
	offset -= sovCompoundSetting(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CompoundSetting) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Delegator)
	if l > 0 {
		n += 1 + l + sovCompoundSetting(uint64(l))
	}
	if len(m.ValidatorSetting) > 0 {
		for _, e := range m.ValidatorSetting {
			l = e.Size()
			n += 1 + l + sovCompoundSetting(uint64(l))
		}
	}
	l = m.AmountToRemain.Size()
	n += 1 + l + sovCompoundSetting(uint64(l))
	if m.Frequency != 0 {
		n += 1 + sovCompoundSetting(uint64(m.Frequency))
	}
	return n
}

func sovCompoundSetting(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCompoundSetting(x uint64) (n int) {
	return sovCompoundSetting(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CompoundSetting) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCompoundSetting
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
			return fmt.Errorf("proto: CompoundSetting: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CompoundSetting: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Delegator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCompoundSetting
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
				return ErrInvalidLengthCompoundSetting
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCompoundSetting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Delegator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorSetting", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCompoundSetting
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
				return ErrInvalidLengthCompoundSetting
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCompoundSetting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorSetting = append(m.ValidatorSetting, &ValidatorSetting{})
			if err := m.ValidatorSetting[len(m.ValidatorSetting)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AmountToRemain", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCompoundSetting
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
				return ErrInvalidLengthCompoundSetting
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCompoundSetting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AmountToRemain.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Frequency", wireType)
			}
			m.Frequency = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCompoundSetting
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Frequency |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCompoundSetting(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCompoundSetting
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
func skipCompoundSetting(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCompoundSetting
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
					return 0, ErrIntOverflowCompoundSetting
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
					return 0, ErrIntOverflowCompoundSetting
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
				return 0, ErrInvalidLengthCompoundSetting
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCompoundSetting
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCompoundSetting
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCompoundSetting        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCompoundSetting          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCompoundSetting = fmt.Errorf("proto: unexpected end of group")
)
