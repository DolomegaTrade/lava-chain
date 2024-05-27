// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lavanet/lava/epochstorage/stake_entry.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	types1 "github.com/cosmos/cosmos-sdk/x/staking/types"
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

type StakeEntry struct {
	Stake              types.Coin         `protobuf:"bytes,1,opt,name=stake,proto3" json:"stake"`
	Address            string             `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	StakeAppliedBlock  uint64             `protobuf:"varint,3,opt,name=stake_applied_block,json=stakeAppliedBlock,proto3" json:"stake_applied_block,omitempty"`
	Endpoints          []Endpoint         `protobuf:"bytes,4,rep,name=endpoints,proto3" json:"endpoints"`
	Geolocation        int32              `protobuf:"varint,5,opt,name=geolocation,proto3" json:"geolocation,omitempty"`
	Chain              string             `protobuf:"bytes,6,opt,name=chain,proto3" json:"chain,omitempty"`
	Moniker            string             `protobuf:"bytes,8,opt,name=moniker,proto3" json:"moniker,omitempty"`
	DelegateTotal      types.Coin         `protobuf:"bytes,9,opt,name=delegate_total,json=delegateTotal,proto3" json:"delegate_total"`
	DelegateLimit      types.Coin         `protobuf:"bytes,10,opt,name=delegate_limit,json=delegateLimit,proto3" json:"delegate_limit"`
	DelegateCommission uint64             `protobuf:"varint,11,opt,name=delegate_commission,json=delegateCommission,proto3" json:"delegate_commission,omitempty"`
	LastChange         uint64             `protobuf:"varint,12,opt,name=last_change,json=lastChange,proto3" json:"last_change,omitempty"`
	BlockReport        *BlockReport       `protobuf:"bytes,13,opt,name=block_report,json=blockReport,proto3" json:"block_report,omitempty"`
	Vault              string             `protobuf:"bytes,14,opt,name=vault,proto3" json:"vault,omitempty"`
	Description        types1.Description `protobuf:"bytes,15,opt,name=description,proto3" json:"description"`
	Jails              uint64             `protobuf:"varint,16,opt,name=jails,proto3" json:"jails,omitempty"`
	JailTime           int64              `protobuf:"varint,17,opt,name=jail_time,json=jailTime,proto3" json:"jail_time,omitempty"`
}

func (m *StakeEntry) Reset()         { *m = StakeEntry{} }
func (m *StakeEntry) String() string { return proto.CompactTextString(m) }
func (*StakeEntry) ProtoMessage()    {}
func (*StakeEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_df6302d6b53c056e, []int{0}
}
func (m *StakeEntry) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StakeEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StakeEntry.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StakeEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StakeEntry.Merge(m, src)
}
func (m *StakeEntry) XXX_Size() int {
	return m.Size()
}
func (m *StakeEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_StakeEntry.DiscardUnknown(m)
}

var xxx_messageInfo_StakeEntry proto.InternalMessageInfo

func (m *StakeEntry) GetStake() types.Coin {
	if m != nil {
		return m.Stake
	}
	return types.Coin{}
}

func (m *StakeEntry) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *StakeEntry) GetStakeAppliedBlock() uint64 {
	if m != nil {
		return m.StakeAppliedBlock
	}
	return 0
}

func (m *StakeEntry) GetEndpoints() []Endpoint {
	if m != nil {
		return m.Endpoints
	}
	return nil
}

func (m *StakeEntry) GetGeolocation() int32 {
	if m != nil {
		return m.Geolocation
	}
	return 0
}

func (m *StakeEntry) GetChain() string {
	if m != nil {
		return m.Chain
	}
	return ""
}

func (m *StakeEntry) GetMoniker() string {
	if m != nil {
		return m.Moniker
	}
	return ""
}

func (m *StakeEntry) GetDelegateTotal() types.Coin {
	if m != nil {
		return m.DelegateTotal
	}
	return types.Coin{}
}

func (m *StakeEntry) GetDelegateLimit() types.Coin {
	if m != nil {
		return m.DelegateLimit
	}
	return types.Coin{}
}

func (m *StakeEntry) GetDelegateCommission() uint64 {
	if m != nil {
		return m.DelegateCommission
	}
	return 0
}

func (m *StakeEntry) GetLastChange() uint64 {
	if m != nil {
		return m.LastChange
	}
	return 0
}

func (m *StakeEntry) GetBlockReport() *BlockReport {
	if m != nil {
		return m.BlockReport
	}
	return nil
}

func (m *StakeEntry) GetVault() string {
	if m != nil {
		return m.Vault
	}
	return ""
}

func (m *StakeEntry) GetDescription() types1.Description {
	if m != nil {
		return m.Description
	}
	return types1.Description{}
}

func (m *StakeEntry) GetJails() uint64 {
	if m != nil {
		return m.Jails
	}
	return 0
}

func (m *StakeEntry) GetJailTime() int64 {
	if m != nil {
		return m.JailTime
	}
	return 0
}

// BlockReport holds the most up-to-date info regarding blocks of the provider
// It is set in the relay payment TX logic
// used by the consumer to calculate the provider's sync score
type BlockReport struct {
	Epoch       uint64 `protobuf:"varint,1,opt,name=epoch,proto3" json:"epoch,omitempty"`
	LatestBlock uint64 `protobuf:"varint,2,opt,name=latest_block,json=latestBlock,proto3" json:"latest_block,omitempty"`
}

func (m *BlockReport) Reset()         { *m = BlockReport{} }
func (m *BlockReport) String() string { return proto.CompactTextString(m) }
func (*BlockReport) ProtoMessage()    {}
func (*BlockReport) Descriptor() ([]byte, []int) {
	return fileDescriptor_df6302d6b53c056e, []int{1}
}
func (m *BlockReport) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BlockReport) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BlockReport.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BlockReport) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockReport.Merge(m, src)
}
func (m *BlockReport) XXX_Size() int {
	return m.Size()
}
func (m *BlockReport) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockReport.DiscardUnknown(m)
}

var xxx_messageInfo_BlockReport proto.InternalMessageInfo

func (m *BlockReport) GetEpoch() uint64 {
	if m != nil {
		return m.Epoch
	}
	return 0
}

func (m *BlockReport) GetLatestBlock() uint64 {
	if m != nil {
		return m.LatestBlock
	}
	return 0
}

func init() {
	proto.RegisterType((*StakeEntry)(nil), "lavanet.lava.epochstorage.StakeEntry")
	proto.RegisterType((*BlockReport)(nil), "lavanet.lava.epochstorage.BlockReport")
}

func init() {
	proto.RegisterFile("lavanet/lava/epochstorage/stake_entry.proto", fileDescriptor_df6302d6b53c056e)
}

var fileDescriptor_df6302d6b53c056e = []byte{
	// 603 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x4f, 0x6f, 0xd3, 0x3e,
	0x18, 0x6e, 0xb6, 0x76, 0x6b, 0x9d, 0x6d, 0xbf, 0xd5, 0xdb, 0xc1, 0xdb, 0x4f, 0xca, 0xc2, 0x86,
	0x50, 0x04, 0x28, 0xd1, 0x86, 0xf8, 0x00, 0xb4, 0xac, 0x08, 0xc4, 0x01, 0x85, 0x9d, 0xb8, 0x44,
	0x4e, 0x62, 0xa5, 0xa6, 0x89, 0x1d, 0xc5, 0x5e, 0xc5, 0xbe, 0x05, 0x1f, 0x83, 0x23, 0x1f, 0x63,
	0xc7, 0x1d, 0x39, 0x21, 0xd4, 0x1e, 0xf8, 0x06, 0x9c, 0x91, 0xed, 0xa4, 0x7f, 0x0e, 0x43, 0x70,
	0x69, 0xfc, 0xbe, 0xef, 0xf3, 0x3e, 0x7e, 0x9f, 0xc7, 0x76, 0xc1, 0x93, 0x1c, 0x4f, 0x31, 0x23,
	0x32, 0x50, 0xdf, 0x80, 0x94, 0x3c, 0x19, 0x0b, 0xc9, 0x2b, 0x9c, 0x91, 0x40, 0x48, 0x3c, 0x21,
	0x11, 0x61, 0xb2, 0xba, 0xf1, 0xcb, 0x8a, 0x4b, 0x0e, 0x8f, 0x6a, 0xb0, 0xaf, 0xbe, 0xfe, 0x2a,
	0xf8, 0xd8, 0xbb, 0x9f, 0x87, 0xb0, 0xb4, 0xe4, 0x94, 0x49, 0x43, 0x72, 0x7c, 0x98, 0xf1, 0x8c,
	0xeb, 0x65, 0xa0, 0x56, 0x75, 0xd6, 0x49, 0xb8, 0x28, 0xb8, 0x08, 0x62, 0x2c, 0x48, 0x30, 0x3d,
	0x8f, 0x89, 0xc4, 0xe7, 0x41, 0xc2, 0x29, 0xab, 0xeb, 0x0f, 0xeb, 0xba, 0x1a, 0x8a, 0xb2, 0x6c,
	0x01, 0xa9, 0xe3, 0x1a, 0xd5, 0xc7, 0x05, 0x65, 0x3c, 0xd0, 0xbf, 0x26, 0x75, 0xfa, 0xab, 0x03,
	0xc0, 0x7b, 0xa5, 0xe4, 0x52, 0x09, 0x81, 0xcf, 0x41, 0x47, 0xeb, 0x42, 0x96, 0x6b, 0x79, 0xf6,
	0xc5, 0x91, 0x6f, 0x78, 0x7d, 0xb5, 0xaf, 0x5f, 0x93, 0xfa, 0x43, 0x4e, 0xd9, 0xa0, 0x7d, 0xfb,
	0xfd, 0xa4, 0x15, 0x1a, 0x34, 0x44, 0x60, 0x1b, 0xa7, 0x69, 0x45, 0x84, 0x40, 0x1b, 0xae, 0xe5,
	0xf5, 0xc2, 0x26, 0x84, 0x3e, 0x38, 0x30, 0x46, 0xe1, 0xb2, 0xcc, 0x29, 0x49, 0xa3, 0x38, 0xe7,
	0xc9, 0x04, 0x6d, 0xba, 0x96, 0xd7, 0x0e, 0xfb, 0xba, 0xf4, 0xc2, 0x54, 0x06, 0xaa, 0x00, 0x5f,
	0x81, 0x5e, 0x63, 0x88, 0x40, 0x6d, 0x77, 0xd3, 0xb3, 0x2f, 0xce, 0xfc, 0x7b, 0x7d, 0xf5, 0x2f,
	0x6b, 0x6c, 0x3d, 0xce, 0xb2, 0x17, 0xba, 0xc0, 0xce, 0x08, 0xcf, 0x79, 0x82, 0x25, 0xe5, 0x0c,
	0x75, 0x5c, 0xcb, 0xeb, 0x84, 0xab, 0x29, 0x78, 0x08, 0x3a, 0xc9, 0x18, 0x53, 0x86, 0xb6, 0xf4,
	0xc8, 0x26, 0x50, 0x52, 0x0a, 0xce, 0xe8, 0x84, 0x54, 0xa8, 0x6b, 0xa4, 0xd4, 0x21, 0x1c, 0x81,
	0xbd, 0x94, 0xe4, 0x24, 0xc3, 0x92, 0x44, 0x92, 0x4b, 0x9c, 0xa3, 0xde, 0xdf, 0x99, 0xb4, 0xdb,
	0xb4, 0x5d, 0xa9, 0xae, 0x35, 0x9e, 0x9c, 0x16, 0x54, 0x22, 0xf0, 0x8f, 0x3c, 0x6f, 0x55, 0x17,
	0x0c, 0xc0, 0xc1, 0x82, 0x27, 0xe1, 0x45, 0x41, 0x85, 0x50, 0x4a, 0x6d, 0x6d, 0x2d, 0x6c, 0x4a,
	0xc3, 0x45, 0x05, 0x9e, 0x00, 0x3b, 0xc7, 0x42, 0x46, 0xc9, 0x18, 0xb3, 0x8c, 0xa0, 0x1d, 0x0d,
	0x04, 0x2a, 0x35, 0xd4, 0x19, 0xf8, 0x1a, 0xec, 0xe8, 0xe3, 0x89, 0x2a, 0x52, 0xf2, 0x4a, 0xa2,
	0x5d, 0x3d, 0xd7, 0xa3, 0x3f, 0xf8, 0xaf, 0x0f, 0x2d, 0xd4, 0xe8, 0xd0, 0x8e, 0x97, 0x81, 0x32,
	0x77, 0x8a, 0xaf, 0x73, 0x89, 0xf6, 0x8c, 0xb9, 0x3a, 0x80, 0xef, 0x80, 0x9d, 0x12, 0x91, 0x54,
	0xb4, 0xd4, 0x87, 0xf2, 0x9f, 0xe6, 0x3f, 0x6b, 0x74, 0x37, 0x97, 0xb5, 0x91, 0xfe, 0x72, 0x09,
	0x1d, 0xf4, 0x94, 0x03, 0x5f, 0x7e, 0x7e, 0x7d, 0x6c, 0x85, 0xab, 0x14, 0x6a, 0x9f, 0x8f, 0x98,
	0xe6, 0x02, 0xed, 0x6b, 0x35, 0x26, 0x80, 0xff, 0x83, 0x9e, 0x5a, 0x44, 0x92, 0x16, 0x04, 0xf5,
	0x5d, 0xcb, 0xdb, 0x0c, 0xbb, 0x2a, 0x71, 0x45, 0x0b, 0xf2, 0xa6, 0xdd, 0xdd, 0xde, 0xef, 0x9e,
	0x8e, 0x80, 0x3d, 0x58, 0x9f, 0x57, 0x0b, 0xd3, 0x17, 0xbf, 0x1d, 0x9a, 0x00, 0x3e, 0x00, 0x3b,
	0x39, 0x96, 0x44, 0xc8, 0xfa, 0xda, 0x6e, 0xe8, 0xa2, 0x6d, 0x72, 0xba, 0x7d, 0x30, 0xba, 0x9d,
	0x39, 0xd6, 0xdd, 0xcc, 0xb1, 0x7e, 0xcc, 0x1c, 0xeb, 0xf3, 0xdc, 0x69, 0xdd, 0xcd, 0x9d, 0xd6,
	0xb7, 0xb9, 0xd3, 0xfa, 0xf0, 0x34, 0xa3, 0x72, 0x7c, 0x1d, 0xfb, 0x09, 0x2f, 0x82, 0xb5, 0xe7,
	0xff, 0x69, 0xfd, 0x0f, 0x40, 0xde, 0x94, 0x44, 0xc4, 0x5b, 0xfa, 0x3d, 0x3e, 0xfb, 0x1d, 0x00,
	0x00, 0xff, 0xff, 0xd6, 0x19, 0x94, 0xba, 0x72, 0x04, 0x00, 0x00,
}

func (m *StakeEntry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StakeEntry) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StakeEntry) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.JailTime != 0 {
		i = encodeVarintStakeEntry(dAtA, i, uint64(m.JailTime))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x88
	}
	if m.Jails != 0 {
		i = encodeVarintStakeEntry(dAtA, i, uint64(m.Jails))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x80
	}
	{
		size, err := m.Description.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintStakeEntry(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x7a
	if len(m.Vault) > 0 {
		i -= len(m.Vault)
		copy(dAtA[i:], m.Vault)
		i = encodeVarintStakeEntry(dAtA, i, uint64(len(m.Vault)))
		i--
		dAtA[i] = 0x72
	}
	if m.BlockReport != nil {
		{
			size, err := m.BlockReport.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintStakeEntry(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x6a
	}
	if m.LastChange != 0 {
		i = encodeVarintStakeEntry(dAtA, i, uint64(m.LastChange))
		i--
		dAtA[i] = 0x60
	}
	if m.DelegateCommission != 0 {
		i = encodeVarintStakeEntry(dAtA, i, uint64(m.DelegateCommission))
		i--
		dAtA[i] = 0x58
	}
	{
		size, err := m.DelegateLimit.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintStakeEntry(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x52
	{
		size, err := m.DelegateTotal.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintStakeEntry(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x4a
	if len(m.Moniker) > 0 {
		i -= len(m.Moniker)
		copy(dAtA[i:], m.Moniker)
		i = encodeVarintStakeEntry(dAtA, i, uint64(len(m.Moniker)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.Chain) > 0 {
		i -= len(m.Chain)
		copy(dAtA[i:], m.Chain)
		i = encodeVarintStakeEntry(dAtA, i, uint64(len(m.Chain)))
		i--
		dAtA[i] = 0x32
	}
	if m.Geolocation != 0 {
		i = encodeVarintStakeEntry(dAtA, i, uint64(m.Geolocation))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Endpoints) > 0 {
		for iNdEx := len(m.Endpoints) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Endpoints[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintStakeEntry(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if m.StakeAppliedBlock != 0 {
		i = encodeVarintStakeEntry(dAtA, i, uint64(m.StakeAppliedBlock))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintStakeEntry(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.Stake.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintStakeEntry(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *BlockReport) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BlockReport) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BlockReport) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.LatestBlock != 0 {
		i = encodeVarintStakeEntry(dAtA, i, uint64(m.LatestBlock))
		i--
		dAtA[i] = 0x10
	}
	if m.Epoch != 0 {
		i = encodeVarintStakeEntry(dAtA, i, uint64(m.Epoch))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintStakeEntry(dAtA []byte, offset int, v uint64) int {
	offset -= sovStakeEntry(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *StakeEntry) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Stake.Size()
	n += 1 + l + sovStakeEntry(uint64(l))
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovStakeEntry(uint64(l))
	}
	if m.StakeAppliedBlock != 0 {
		n += 1 + sovStakeEntry(uint64(m.StakeAppliedBlock))
	}
	if len(m.Endpoints) > 0 {
		for _, e := range m.Endpoints {
			l = e.Size()
			n += 1 + l + sovStakeEntry(uint64(l))
		}
	}
	if m.Geolocation != 0 {
		n += 1 + sovStakeEntry(uint64(m.Geolocation))
	}
	l = len(m.Chain)
	if l > 0 {
		n += 1 + l + sovStakeEntry(uint64(l))
	}
	l = len(m.Moniker)
	if l > 0 {
		n += 1 + l + sovStakeEntry(uint64(l))
	}
	l = m.DelegateTotal.Size()
	n += 1 + l + sovStakeEntry(uint64(l))
	l = m.DelegateLimit.Size()
	n += 1 + l + sovStakeEntry(uint64(l))
	if m.DelegateCommission != 0 {
		n += 1 + sovStakeEntry(uint64(m.DelegateCommission))
	}
	if m.LastChange != 0 {
		n += 1 + sovStakeEntry(uint64(m.LastChange))
	}
	if m.BlockReport != nil {
		l = m.BlockReport.Size()
		n += 1 + l + sovStakeEntry(uint64(l))
	}
	l = len(m.Vault)
	if l > 0 {
		n += 1 + l + sovStakeEntry(uint64(l))
	}
	l = m.Description.Size()
	n += 1 + l + sovStakeEntry(uint64(l))
	if m.Jails != 0 {
		n += 2 + sovStakeEntry(uint64(m.Jails))
	}
	if m.JailTime != 0 {
		n += 2 + sovStakeEntry(uint64(m.JailTime))
	}
	return n
}

func (m *BlockReport) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Epoch != 0 {
		n += 1 + sovStakeEntry(uint64(m.Epoch))
	}
	if m.LatestBlock != 0 {
		n += 1 + sovStakeEntry(uint64(m.LatestBlock))
	}
	return n
}

func sovStakeEntry(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozStakeEntry(x uint64) (n int) {
	return sovStakeEntry(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *StakeEntry) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStakeEntry
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
			return fmt.Errorf("proto: StakeEntry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StakeEntry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stake", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStakeEntry
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
				return ErrInvalidLengthStakeEntry
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStakeEntry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Stake.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStakeEntry
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
				return ErrInvalidLengthStakeEntry
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStakeEntry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StakeAppliedBlock", wireType)
			}
			m.StakeAppliedBlock = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStakeEntry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StakeAppliedBlock |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Endpoints", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStakeEntry
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
				return ErrInvalidLengthStakeEntry
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStakeEntry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Endpoints = append(m.Endpoints, Endpoint{})
			if err := m.Endpoints[len(m.Endpoints)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Geolocation", wireType)
			}
			m.Geolocation = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStakeEntry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Geolocation |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStakeEntry
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
				return ErrInvalidLengthStakeEntry
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStakeEntry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Chain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Moniker", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStakeEntry
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
				return ErrInvalidLengthStakeEntry
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStakeEntry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Moniker = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelegateTotal", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStakeEntry
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
				return ErrInvalidLengthStakeEntry
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStakeEntry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DelegateTotal.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelegateLimit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStakeEntry
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
				return ErrInvalidLengthStakeEntry
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStakeEntry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DelegateLimit.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelegateCommission", wireType)
			}
			m.DelegateCommission = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStakeEntry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DelegateCommission |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastChange", wireType)
			}
			m.LastChange = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStakeEntry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastChange |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockReport", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStakeEntry
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
				return ErrInvalidLengthStakeEntry
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStakeEntry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BlockReport == nil {
				m.BlockReport = &BlockReport{}
			}
			if err := m.BlockReport.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Vault", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStakeEntry
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
				return ErrInvalidLengthStakeEntry
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStakeEntry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Vault = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 15:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStakeEntry
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
				return ErrInvalidLengthStakeEntry
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStakeEntry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Description.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 16:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Jails", wireType)
			}
			m.Jails = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStakeEntry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Jails |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 17:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field JailTime", wireType)
			}
			m.JailTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStakeEntry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.JailTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipStakeEntry(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStakeEntry
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
func (m *BlockReport) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStakeEntry
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
			return fmt.Errorf("proto: BlockReport: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BlockReport: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Epoch", wireType)
			}
			m.Epoch = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStakeEntry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Epoch |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LatestBlock", wireType)
			}
			m.LatestBlock = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStakeEntry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LatestBlock |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipStakeEntry(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStakeEntry
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
func skipStakeEntry(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStakeEntry
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
					return 0, ErrIntOverflowStakeEntry
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
					return 0, ErrIntOverflowStakeEntry
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
				return 0, ErrInvalidLengthStakeEntry
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupStakeEntry
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthStakeEntry
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthStakeEntry        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStakeEntry          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupStakeEntry = fmt.Errorf("proto: unexpected end of group")
)
