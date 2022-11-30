// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cfevesting/genesis.proto

package types

import (
	fmt "fmt"
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

// GenesisState defines the cfevesting module's genesis state.
type GenesisState struct {
	Params              Params                 `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	VestingTypes        []GenesisVestingType   `protobuf:"bytes,2,rep,name=vesting_types,json=vestingTypes,proto3" json:"vesting_types" yaml:"vesting_types"`
	AccountVestingPools []*AccountVestingPools `protobuf:"bytes,3,rep,name=account_vesting_pools,json=accountVestingPools,proto3" json:"account_vesting_pools,omitempty" yaml:"account_vesting_pools"`
	VestingAccountList  []VestingAccount       `protobuf:"bytes,4,rep,name=vesting_account_list,json=vestingAccountList,proto3" json:"vesting_account_list"`
	VestingAccountCount uint64                 `protobuf:"varint,5,opt,name=vesting_account_count,json=vestingAccountCount,proto3" json:"vesting_account_count,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff4bc386047a9a81, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetVestingTypes() []GenesisVestingType {
	if m != nil {
		return m.VestingTypes
	}
	return nil
}

func (m *GenesisState) GetAccountVestingPools() []*AccountVestingPools {
	if m != nil {
		return m.AccountVestingPools
	}
	return nil
}

func (m *GenesisState) GetVestingAccountList() []VestingAccount {
	if m != nil {
		return m.VestingAccountList
	}
	return nil
}

func (m *GenesisState) GetVestingAccountCount() uint64 {
	if m != nil {
		return m.VestingAccountCount
	}
	return 0
}

type GenesisVestingType struct {
	// vesting type name
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// period of locked coins from vesting start
	LockupPeriod     int64  `protobuf:"varint,2,opt,name=lockup_period,json=lockupPeriod,proto3" json:"lockup_period,omitempty"`
	LockupPeriodUnit string `protobuf:"bytes,3,opt,name=lockup_period_unit,json=lockupPeriodUnit,proto3" json:"lockup_period_unit,omitempty"`
	// period of veesting coins from lockup period end
	VestingPeriod     int64                                  `protobuf:"varint,4,opt,name=vesting_period,json=vestingPeriod,proto3" json:"vesting_period,omitempty"`
	VestingPeriodUnit string                                 `protobuf:"bytes,5,opt,name=vesting_period_unit,json=vestingPeriodUnit,proto3" json:"vesting_period_unit,omitempty"`
	InitialBonus      github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,6,opt,name=initial_bonus,json=initialBonus,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"initial_bonus"`
}

func (m *GenesisVestingType) Reset()         { *m = GenesisVestingType{} }
func (m *GenesisVestingType) String() string { return proto.CompactTextString(m) }
func (*GenesisVestingType) ProtoMessage()    {}
func (*GenesisVestingType) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff4bc386047a9a81, []int{1}
}
func (m *GenesisVestingType) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisVestingType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisVestingType.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisVestingType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisVestingType.Merge(m, src)
}
func (m *GenesisVestingType) XXX_Size() int {
	return m.Size()
}
func (m *GenesisVestingType) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisVestingType.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisVestingType proto.InternalMessageInfo

func (m *GenesisVestingType) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GenesisVestingType) GetLockupPeriod() int64 {
	if m != nil {
		return m.LockupPeriod
	}
	return 0
}

func (m *GenesisVestingType) GetLockupPeriodUnit() string {
	if m != nil {
		return m.LockupPeriodUnit
	}
	return ""
}

func (m *GenesisVestingType) GetVestingPeriod() int64 {
	if m != nil {
		return m.VestingPeriod
	}
	return 0
}

func (m *GenesisVestingType) GetVestingPeriodUnit() string {
	if m != nil {
		return m.VestingPeriodUnit
	}
	return ""
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "chain4energy.c4echain.cfevesting.GenesisState")
	proto.RegisterType((*GenesisVestingType)(nil), "chain4energy.c4echain.cfevesting.GenesisVestingType")
}

func init() { proto.RegisterFile("cfevesting/genesis.proto", fileDescriptor_ff4bc386047a9a81) }

var fileDescriptor_ff4bc386047a9a81 = []byte{
	// 524 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0x9b, 0xb6, 0xab, 0x84, 0xd7, 0x22, 0x70, 0x3b, 0x11, 0x55, 0x53, 0x1a, 0x15, 0x0d,
	0xf5, 0xc0, 0x12, 0x54, 0xc2, 0x85, 0x1b, 0x01, 0xc1, 0x05, 0xa1, 0x2a, 0x03, 0x0e, 0x5c, 0xa2,
	0xd4, 0x33, 0xa9, 0xb5, 0xc4, 0x8e, 0x6a, 0xa7, 0xd0, 0xcf, 0xc0, 0x85, 0x6f, 0xc1, 0x57, 0xd9,
	0x71, 0x47, 0xc4, 0xa1, 0x42, 0xed, 0x37, 0xd8, 0x95, 0x0b, 0x8a, 0xed, 0x6e, 0xc9, 0x36, 0xa9,
	0x17, 0xc7, 0x7e, 0xfe, 0xbf, 0xdf, 0xfb, 0xeb, 0xe5, 0x19, 0x98, 0xe8, 0x2b, 0x5e, 0x60, 0x2e,
	0x08, 0x8d, 0xdd, 0x18, 0x53, 0xcc, 0x09, 0x77, 0xb2, 0x39, 0x13, 0x0c, 0xda, 0x68, 0x16, 0x11,
	0xea, 0x61, 0x8a, 0xe7, 0xf1, 0xd2, 0x41, 0x1e, 0x96, 0x67, 0xe7, 0x5a, 0xdf, 0xef, 0xc5, 0x2c,
	0x66, 0x52, 0xec, 0x16, 0x3b, 0x95, 0xd7, 0x7f, 0x54, 0x22, 0x66, 0xd1, 0x3c, 0x4a, 0x35, 0xb0,
	0x7f, 0x54, 0xba, 0x88, 0x10, 0x62, 0x39, 0x15, 0xa1, 0x3e, 0x87, 0x19, 0x63, 0x89, 0x96, 0xd9,
	0x25, 0xd9, 0xf6, 0x5a, 0xcb, 0x95, 0x62, 0xf8, 0xaf, 0x01, 0xda, 0xef, 0x94, 0xd7, 0x13, 0x11,
	0x09, 0x0c, 0xdf, 0x82, 0x96, 0xaa, 0x64, 0x1a, 0xb6, 0x31, 0xda, 0x1f, 0x8f, 0x9c, 0x5d, 0xde,
	0x9d, 0x89, 0xd4, 0xfb, 0xcd, 0xf3, 0xd5, 0xa0, 0x16, 0xe8, 0x6c, 0xf8, 0x0d, 0x74, 0xb6, 0x15,
	0xc5, 0x32, 0xc3, 0xdc, 0xac, 0xdb, 0x8d, 0xd1, 0xfe, 0xd8, 0xdb, 0x8d, 0xd3, 0x76, 0x3e, 0xab,
	0xe3, 0xc7, 0x65, 0x86, 0xfd, 0xc3, 0x02, 0x7d, 0xb9, 0x1a, 0xf4, 0x96, 0x51, 0x9a, 0xbc, 0x1c,
	0x56, 0xc0, 0xc3, 0xa0, 0xbd, 0xb8, 0x96, 0x72, 0xf8, 0xc3, 0x00, 0x07, 0x77, 0xb5, 0x84, 0x9b,
	0x0d, 0xe9, 0xe0, 0xc5, 0x6e, 0x07, 0xaf, 0x54, 0xba, 0x76, 0x30, 0x29, 0x92, 0x7d, 0xfb, 0x72,
	0x35, 0x38, 0x54, 0xe5, 0xef, 0xa4, 0x0f, 0x83, 0x6e, 0x74, 0x3b, 0x0d, 0xce, 0x40, 0xef, 0x46,
	0xe3, 0xc3, 0x84, 0x70, 0x61, 0x36, 0xa5, 0x97, 0x67, 0xbb, 0xbd, 0x68, 0x9a, 0xb6, 0xa4, 0x9b,
	0x0c, 0x17, 0x95, 0xe8, 0x7b, 0xc2, 0x05, 0x1c, 0x83, 0x83, 0x9b, 0x95, 0xe4, 0x6a, 0xee, 0xd9,
	0xc6, 0xa8, 0x19, 0x74, 0xab, 0x29, 0xaf, 0x8b, 0x65, 0xf8, 0xab, 0x0e, 0xe0, 0xed, 0x76, 0x43,
	0x08, 0x9a, 0x34, 0x4a, 0xb1, 0x9c, 0x80, 0x7b, 0x81, 0xdc, 0xc3, 0xc7, 0xa0, 0x93, 0x30, 0x74,
	0x96, 0x67, 0x61, 0x86, 0xe7, 0x84, 0x9d, 0x9a, 0x75, 0xdb, 0x18, 0x35, 0x82, 0xb6, 0x0a, 0x4e,
	0x64, 0x0c, 0x3e, 0x05, 0xb0, 0x22, 0x0a, 0x73, 0x4a, 0x84, 0xd9, 0x90, 0x98, 0x07, 0x65, 0xe5,
	0x27, 0x4a, 0x04, 0x3c, 0x02, 0xf7, 0xaf, 0x5a, 0xa8, 0x98, 0x4d, 0xc9, 0xdc, 0x0e, 0x8e, 0x86,
	0x3a, 0xa0, 0x5b, 0x95, 0x29, 0xea, 0x9e, 0xa4, 0x3e, 0xac, 0x68, 0x25, 0xf6, 0x04, 0x74, 0x08,
	0x25, 0x82, 0x44, 0x49, 0x38, 0x65, 0x34, 0xe7, 0x66, 0xab, 0x50, 0xfa, 0x4e, 0xd1, 0xb9, 0x3f,
	0xab, 0xc1, 0x93, 0x98, 0x88, 0x59, 0x3e, 0x75, 0x10, 0x4b, 0x5d, 0xc4, 0x78, 0xca, 0xb8, 0xfe,
	0x1c, 0xf3, 0xd3, 0x33, 0x57, 0x8e, 0x94, 0xf3, 0x06, 0xa3, 0xa0, 0xad, 0x21, 0x7e, 0xc1, 0xf0,
	0x3f, 0x9c, 0xaf, 0x2d, 0xe3, 0x62, 0x6d, 0x19, 0x7f, 0xd7, 0x96, 0xf1, 0x73, 0x63, 0xd5, 0x2e,
	0x36, 0x56, 0xed, 0xf7, 0xc6, 0xaa, 0x7d, 0xf1, 0xca, 0xbc, 0xd2, 0xdf, 0x74, 0x91, 0x87, 0x8f,
	0x65, 0xc0, 0xfd, 0xee, 0x96, 0xde, 0xa1, 0xac, 0x30, 0x6d, 0xc9, 0xe7, 0xf7, 0xfc, 0x7f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xc4, 0x4b, 0x2d, 0x44, 0x34, 0x04, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.VestingAccountCount != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.VestingAccountCount))
		i--
		dAtA[i] = 0x28
	}
	if len(m.VestingAccountList) > 0 {
		for iNdEx := len(m.VestingAccountList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.VestingAccountList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.AccountVestingPools) > 0 {
		for iNdEx := len(m.AccountVestingPools) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AccountVestingPools[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.VestingTypes) > 0 {
		for iNdEx := len(m.VestingTypes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.VestingTypes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *GenesisVestingType) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisVestingType) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisVestingType) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.InitialBonus.Size()
		i -= size
		if _, err := m.InitialBonus.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	if len(m.VestingPeriodUnit) > 0 {
		i -= len(m.VestingPeriodUnit)
		copy(dAtA[i:], m.VestingPeriodUnit)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.VestingPeriodUnit)))
		i--
		dAtA[i] = 0x2a
	}
	if m.VestingPeriod != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.VestingPeriod))
		i--
		dAtA[i] = 0x20
	}
	if len(m.LockupPeriodUnit) > 0 {
		i -= len(m.LockupPeriodUnit)
		copy(dAtA[i:], m.LockupPeriodUnit)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.LockupPeriodUnit)))
		i--
		dAtA[i] = 0x1a
	}
	if m.LockupPeriod != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.LockupPeriod))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.VestingTypes) > 0 {
		for _, e := range m.VestingTypes {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.AccountVestingPools) > 0 {
		for _, e := range m.AccountVestingPools {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.VestingAccountList) > 0 {
		for _, e := range m.VestingAccountList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.VestingAccountCount != 0 {
		n += 1 + sovGenesis(uint64(m.VestingAccountCount))
	}
	return n
}

func (m *GenesisVestingType) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.LockupPeriod != 0 {
		n += 1 + sovGenesis(uint64(m.LockupPeriod))
	}
	l = len(m.LockupPeriodUnit)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.VestingPeriod != 0 {
		n += 1 + sovGenesis(uint64(m.VestingPeriod))
	}
	l = len(m.VestingPeriodUnit)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = m.InitialBonus.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VestingTypes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VestingTypes = append(m.VestingTypes, GenesisVestingType{})
			if err := m.VestingTypes[len(m.VestingTypes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccountVestingPools", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AccountVestingPools = append(m.AccountVestingPools, &AccountVestingPools{})
			if err := m.AccountVestingPools[len(m.AccountVestingPools)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VestingAccountList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VestingAccountList = append(m.VestingAccountList, VestingAccount{})
			if err := m.VestingAccountList[len(m.VestingAccountList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VestingAccountCount", wireType)
			}
			m.VestingAccountCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VestingAccountCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *GenesisVestingType) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisVestingType: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisVestingType: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LockupPeriod", wireType)
			}
			m.LockupPeriod = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LockupPeriod |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LockupPeriodUnit", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LockupPeriodUnit = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VestingPeriod", wireType)
			}
			m.VestingPeriod = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VestingPeriod |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VestingPeriodUnit", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VestingPeriodUnit = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitialBonus", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.InitialBonus.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
