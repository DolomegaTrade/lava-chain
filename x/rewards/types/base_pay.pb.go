// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lavanet/lava/rewards/base_pay.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
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

// aggregated rewards for the provider through out the month
type BasePay struct {
	Total         github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,1,opt,name=total,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"total"`
	TotalAdjusted github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=totalAdjusted,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"totalAdjusted"`
	IprpcCu       uint64                                 `protobuf:"varint,3,opt,name=iprpc_cu,json=iprpcCu,proto3" json:"iprpc_cu,omitempty"`
}

func (m *BasePay) Reset()         { *m = BasePay{} }
func (m *BasePay) String() string { return proto.CompactTextString(m) }
func (*BasePay) ProtoMessage()    {}
func (*BasePay) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2fb0eb917a4ee4e, []int{0}
}
func (m *BasePay) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BasePay) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BasePay.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BasePay) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BasePay.Merge(m, src)
}
func (m *BasePay) XXX_Size() int {
	return m.Size()
}
func (m *BasePay) XXX_DiscardUnknown() {
	xxx_messageInfo_BasePay.DiscardUnknown(m)
}

var xxx_messageInfo_BasePay proto.InternalMessageInfo

func (m *BasePay) GetIprpcCu() uint64 {
	if m != nil {
		return m.IprpcCu
	}
	return 0
}

// aggregated rewards for the provider through out the month
type BasePayGenesis struct {
	Index   string  `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	BasePay BasePay `protobuf:"bytes,2,opt,name=base_pay,json=basePay,proto3" json:"base_pay"`
}

func (m *BasePayGenesis) Reset()         { *m = BasePayGenesis{} }
func (m *BasePayGenesis) String() string { return proto.CompactTextString(m) }
func (*BasePayGenesis) ProtoMessage()    {}
func (*BasePayGenesis) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2fb0eb917a4ee4e, []int{1}
}
func (m *BasePayGenesis) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BasePayGenesis) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BasePayGenesis.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BasePayGenesis) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BasePayGenesis.Merge(m, src)
}
func (m *BasePayGenesis) XXX_Size() int {
	return m.Size()
}
func (m *BasePayGenesis) XXX_DiscardUnknown() {
	xxx_messageInfo_BasePayGenesis.DiscardUnknown(m)
}

var xxx_messageInfo_BasePayGenesis proto.InternalMessageInfo

func (m *BasePayGenesis) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *BasePayGenesis) GetBasePay() BasePay {
	if m != nil {
		return m.BasePay
	}
	return BasePay{}
}

func init() {
	proto.RegisterType((*BasePay)(nil), "lavanet.lava.rewards.BasePay")
	proto.RegisterType((*BasePayGenesis)(nil), "lavanet.lava.rewards.BasePayGenesis")
}

func init() {
	proto.RegisterFile("lavanet/lava/rewards/base_pay.proto", fileDescriptor_a2fb0eb917a4ee4e)
}

var fileDescriptor_a2fb0eb917a4ee4e = []byte{
	// 349 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x51, 0xc1, 0x4e, 0x02, 0x31,
	0x14, 0xdc, 0x2a, 0x08, 0xd6, 0x68, 0xe2, 0x86, 0x03, 0x90, 0x58, 0x08, 0x26, 0xca, 0x85, 0x36,
	0xe8, 0xd5, 0x98, 0x80, 0x26, 0x86, 0x9b, 0xd9, 0x78, 0xf2, 0x42, 0xba, 0xbb, 0x15, 0x57, 0xa1,
	0xdd, 0x6c, 0xbb, 0x08, 0x7f, 0xe1, 0xc7, 0xf8, 0x11, 0x9c, 0x0c, 0xf1, 0x64, 0x3c, 0x10, 0x03,
	0x3f, 0x62, 0xb6, 0xad, 0x46, 0x12, 0x2f, 0x5e, 0xfa, 0xde, 0xf4, 0x4d, 0xa7, 0x6f, 0x32, 0xf0,
	0x70, 0x48, 0xc7, 0x94, 0x33, 0x45, 0xb2, 0x4a, 0x12, 0xf6, 0x44, 0x93, 0x50, 0x12, 0x9f, 0x4a,
	0xd6, 0x8f, 0xe9, 0x14, 0xc7, 0x89, 0x50, 0xc2, 0x2d, 0x59, 0x12, 0xce, 0x2a, 0xb6, 0xa4, 0x6a,
	0x69, 0x20, 0x06, 0x42, 0x13, 0x48, 0xd6, 0x19, 0x6e, 0x15, 0x05, 0x42, 0x8e, 0x84, 0x91, 0x20,
	0xe3, 0xb6, 0xcf, 0x14, 0x6d, 0x93, 0x40, 0x44, 0xdc, 0xce, 0x2b, 0x66, 0xde, 0x37, 0x0f, 0x0d,
	0xb0, 0xa3, 0x7d, 0x3a, 0x8a, 0xb8, 0x20, 0xfa, 0x34, 0x57, 0x8d, 0x57, 0x00, 0x0b, 0x5d, 0x2a,
	0xd9, 0x35, 0x9d, 0xba, 0x1e, 0xcc, 0x2b, 0xa1, 0xe8, 0xb0, 0x0c, 0xea, 0xa0, 0xb9, 0xdd, 0x3d,
	0x9b, 0x2d, 0x6a, 0xce, 0xc7, 0xa2, 0x76, 0x34, 0x88, 0xd4, 0x7d, 0xea, 0xe3, 0x40, 0x8c, 0xac,
	0x9c, 0x2d, 0x2d, 0x19, 0x3e, 0x12, 0x35, 0x8d, 0x99, 0xc4, 0x3d, 0xae, 0xde, 0x5e, 0x5a, 0xd0,
	0xfe, 0xd6, 0xe3, 0xca, 0x33, 0x52, 0xee, 0x0d, 0xdc, 0xd5, 0x4d, 0x27, 0x7c, 0x48, 0xa5, 0x62,
	0x61, 0x79, 0x43, 0x6b, 0xe3, 0x7f, 0x68, 0x5f, 0xb2, 0xc0, 0x5b, 0x17, 0x71, 0x2b, 0xb0, 0x18,
	0xc5, 0x49, 0x1c, 0xf4, 0x83, 0xb4, 0xbc, 0x59, 0x07, 0xcd, 0x9c, 0x57, 0xd0, 0xf8, 0x22, 0x6d,
	0xdc, 0xc1, 0x3d, 0xeb, 0xe7, 0x8a, 0x71, 0x26, 0x23, 0xe9, 0x96, 0x60, 0x3e, 0xe2, 0x21, 0x9b,
	0x18, 0x5b, 0x9e, 0x01, 0xee, 0x39, 0x2c, 0x7e, 0x87, 0xa0, 0x77, 0xda, 0x39, 0x39, 0xc0, 0x7f,
	0xa5, 0x80, 0xad, 0x5a, 0x37, 0x97, 0xad, 0xec, 0x15, 0x7c, 0x0b, 0x3b, 0xb3, 0x25, 0x02, 0xf3,
	0x25, 0x02, 0x9f, 0x4b, 0x04, 0x9e, 0x57, 0xc8, 0x99, 0xaf, 0x90, 0xf3, 0xbe, 0x42, 0xce, 0xed,
	0xf1, 0x2f, 0x4f, 0x6b, 0xe1, 0x4f, 0x7e, 0xe2, 0xd7, 0xc6, 0xfc, 0x2d, 0x1d, 0xc1, 0xe9, 0x57,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x0f, 0x24, 0x2a, 0x65, 0x23, 0x02, 0x00, 0x00,
}

func (m *BasePay) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BasePay) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BasePay) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.IprpcCu != 0 {
		i = encodeVarintBasePay(dAtA, i, uint64(m.IprpcCu))
		i--
		dAtA[i] = 0x18
	}
	{
		size := m.TotalAdjusted.Size()
		i -= size
		if _, err := m.TotalAdjusted.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintBasePay(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.Total.Size()
		i -= size
		if _, err := m.Total.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintBasePay(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *BasePayGenesis) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BasePayGenesis) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BasePayGenesis) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.BasePay.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintBasePay(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintBasePay(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintBasePay(dAtA []byte, offset int, v uint64) int {
	offset -= sovBasePay(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BasePay) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Total.Size()
	n += 1 + l + sovBasePay(uint64(l))
	l = m.TotalAdjusted.Size()
	n += 1 + l + sovBasePay(uint64(l))
	if m.IprpcCu != 0 {
		n += 1 + sovBasePay(uint64(m.IprpcCu))
	}
	return n
}

func (m *BasePayGenesis) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovBasePay(uint64(l))
	}
	l = m.BasePay.Size()
	n += 1 + l + sovBasePay(uint64(l))
	return n
}

func sovBasePay(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBasePay(x uint64) (n int) {
	return sovBasePay(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BasePay) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBasePay
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
			return fmt.Errorf("proto: BasePay: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BasePay: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Total", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBasePay
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
				return ErrInvalidLengthBasePay
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBasePay
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Total.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalAdjusted", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBasePay
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
				return ErrInvalidLengthBasePay
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBasePay
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TotalAdjusted.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IprpcCu", wireType)
			}
			m.IprpcCu = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBasePay
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.IprpcCu |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipBasePay(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBasePay
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
func (m *BasePayGenesis) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBasePay
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
			return fmt.Errorf("proto: BasePayGenesis: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BasePayGenesis: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBasePay
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
				return ErrInvalidLengthBasePay
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBasePay
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BasePay", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBasePay
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
				return ErrInvalidLengthBasePay
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBasePay
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BasePay.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBasePay(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBasePay
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
func skipBasePay(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBasePay
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
					return 0, ErrIntOverflowBasePay
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
					return 0, ErrIntOverflowBasePay
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
				return 0, ErrInvalidLengthBasePay
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBasePay
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBasePay
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBasePay        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBasePay          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBasePay = fmt.Errorf("proto: unexpected end of group")
)
