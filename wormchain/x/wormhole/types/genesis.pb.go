// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: wormhole/genesis.proto

package types

import (
	fmt "fmt"
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

// GenesisState defines the wormhole module's genesis state.
type GenesisState struct {
	GuardianSetList           []GuardianSet              `protobuf:"bytes,1,rep,name=guardianSetList,proto3" json:"guardianSetList"`
	Config                    *Config                    `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
	ReplayProtectionList      []ReplayProtection         `protobuf:"bytes,3,rep,name=replayProtectionList,proto3" json:"replayProtectionList"`
	SequenceCounterList       []SequenceCounter          `protobuf:"bytes,4,rep,name=sequenceCounterList,proto3" json:"sequenceCounterList"`
	ConsensusGuardianSetIndex *ConsensusGuardianSetIndex `protobuf:"bytes,5,opt,name=consensusGuardianSetIndex,proto3" json:"consensusGuardianSetIndex,omitempty"`
	GuardianValidatorList     []GuardianValidator        `protobuf:"bytes,6,rep,name=guardianValidatorList,proto3" json:"guardianValidatorList"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a7ced3fe0304831, []int{0}
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

func (m *GenesisState) GetGuardianSetList() []GuardianSet {
	if m != nil {
		return m.GuardianSetList
	}
	return nil
}

func (m *GenesisState) GetConfig() *Config {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *GenesisState) GetReplayProtectionList() []ReplayProtection {
	if m != nil {
		return m.ReplayProtectionList
	}
	return nil
}

func (m *GenesisState) GetSequenceCounterList() []SequenceCounter {
	if m != nil {
		return m.SequenceCounterList
	}
	return nil
}

func (m *GenesisState) GetConsensusGuardianSetIndex() *ConsensusGuardianSetIndex {
	if m != nil {
		return m.ConsensusGuardianSetIndex
	}
	return nil
}

func (m *GenesisState) GetGuardianValidatorList() []GuardianValidator {
	if m != nil {
		return m.GuardianValidatorList
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "wormhole_foundation.wormchain.wormhole.GenesisState")
}

func init() { proto.RegisterFile("wormhole/genesis.proto", fileDescriptor_9a7ced3fe0304831) }

var fileDescriptor_9a7ced3fe0304831 = []byte{
	// 426 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xcf, 0x8b, 0xda, 0x40,
	0x14, 0xc7, 0x93, 0xfa, 0xe3, 0x30, 0x16, 0x0a, 0xa9, 0x16, 0x6b, 0x21, 0xda, 0x1e, 0x8a, 0x50,
	0x9a, 0x80, 0x1e, 0x5a, 0x8f, 0x55, 0xa8, 0x14, 0x7a, 0x28, 0x09, 0xf4, 0xb0, 0x97, 0x30, 0x26,
	0x63, 0x1c, 0xd0, 0x19, 0x37, 0x33, 0xd9, 0xd5, 0xd3, 0xfe, 0x0b, 0x7b, 0xdc, 0x3f, 0xc9, 0xa3,
	0xc7, 0x3d, 0x2d, 0x8b, 0xfe, 0x23, 0x4b, 0x66, 0x92, 0xd1, 0x75, 0xb3, 0x10, 0x6f, 0xc3, 0x9b,
	0xf7, 0x3e, 0xdf, 0xf7, 0xbe, 0xef, 0x81, 0x0f, 0xd7, 0x34, 0x5a, 0xcc, 0xe8, 0x1c, 0xd9, 0x21,
	0x22, 0x88, 0x61, 0x66, 0x2d, 0x23, 0xca, 0xa9, 0xf1, 0x35, 0x8b, 0x7b, 0x53, 0x1a, 0x93, 0x00,
	0x72, 0x4c, 0x89, 0x95, 0xc4, 0xfc, 0x19, 0xc4, 0xf2, 0x95, 0xfc, 0xb6, 0x3e, 0x1d, 0xea, 0x63,
	0x18, 0x05, 0x18, 0x12, 0x8f, 0x21, 0x2e, 0x21, 0xad, 0x86, 0xfa, 0xf4, 0x29, 0x99, 0xe2, 0x30,
	0x0d, 0x77, 0x54, 0x38, 0x42, 0xcb, 0x39, 0x5c, 0x7b, 0x49, 0x18, 0xf9, 0x42, 0x42, 0x66, 0xb4,
	0x55, 0x06, 0x43, 0x97, 0x31, 0x22, 0x3e, 0xf2, 0x7c, 0x1a, 0x13, 0x8e, 0xa2, 0x34, 0xe1, 0xdb,
	0x31, 0x99, 0x21, 0xc2, 0x62, 0xe6, 0x1d, 0x37, 0xe0, 0x61, 0x12, 0xa0, 0x55, 0x9a, 0xfc, 0xf9,
	0x65, 0x8f, 0x57, 0x70, 0x8e, 0x03, 0xc8, 0x69, 0xc6, 0xab, 0x87, 0x34, 0xa4, 0xe2, 0x69, 0x27,
	0x2f, 0x19, 0xfd, 0x72, 0x57, 0x01, 0x6f, 0xc7, 0xd2, 0x16, 0x97, 0x43, 0x8e, 0x0c, 0x1f, 0xbc,
	0xcb, 0x10, 0x2e, 0xe2, 0x7f, 0x31, 0xe3, 0x4d, 0xbd, 0x53, 0xea, 0xd6, 0x7a, 0x7d, 0xab, 0x98,
	0x5f, 0xd6, 0xf8, 0x50, 0x3e, 0x2c, 0x6f, 0x1e, 0xda, 0x9a, 0x73, 0x4a, 0x34, 0x7e, 0x83, 0xaa,
	0xb4, 0xab, 0xf9, 0xa6, 0xa3, 0x77, 0x6b, 0x3d, 0xab, 0x28, 0x7b, 0x24, 0xaa, 0x9c, 0xb4, 0xda,
	0x88, 0x40, 0x5d, 0xfa, 0xfb, 0x4f, 0xd9, 0x2b, 0x3a, 0x2e, 0x89, 0x8e, 0x7f, 0x16, 0xa5, 0x3a,
	0x27, 0x8c, 0xb4, 0xed, 0x5c, 0xb6, 0x41, 0xc1, 0xfb, 0x6c, 0x63, 0x23, 0xb9, 0x30, 0x21, 0x59,
	0x16, 0x92, 0x3f, 0x8a, 0x4a, 0xba, 0xcf, 0x11, 0xa9, 0x62, 0x1e, 0xd9, 0xb8, 0x01, 0x1f, 0xd5,
	0x05, 0x1c, 0x79, 0xfb, 0x27, 0x59, 0x7f, 0xb3, 0x22, 0xfc, 0xfb, 0x75, 0x86, 0x7f, 0xf9, 0x20,
	0xe7, 0x75, 0x0d, 0x23, 0x06, 0x8d, 0x6c, 0x81, 0xff, 0xb3, 0xa3, 0x12, 0x33, 0x57, 0xc5, 0xcc,
	0x83, 0x73, 0x0f, 0x43, 0x41, 0xd2, 0xa9, 0xf3, 0xe9, 0x43, 0x77, 0xb3, 0x33, 0xf5, 0xed, 0xce,
	0xd4, 0x1f, 0x77, 0xa6, 0x7e, 0xbb, 0x37, 0xb5, 0xed, 0xde, 0xd4, 0xee, 0xf7, 0xa6, 0x76, 0x31,
	0x08, 0x31, 0x9f, 0xc5, 0x13, 0xcb, 0xa7, 0x0b, 0x3b, 0xa3, 0x7f, 0x3f, 0x68, 0xdb, 0x4a, 0xdb,
	0x5e, 0xa9, 0x7f, 0x9b, 0xaf, 0x97, 0x88, 0x4d, 0xaa, 0xe2, 0xec, 0xfb, 0x4f, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x40, 0x4b, 0xe8, 0xcd, 0x15, 0x04, 0x00, 0x00,
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
	if len(m.GuardianValidatorList) > 0 {
		for iNdEx := len(m.GuardianValidatorList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.GuardianValidatorList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if m.ConsensusGuardianSetIndex != nil {
		{
			size, err := m.ConsensusGuardianSetIndex.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if len(m.SequenceCounterList) > 0 {
		for iNdEx := len(m.SequenceCounterList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SequenceCounterList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.ReplayProtectionList) > 0 {
		for iNdEx := len(m.ReplayProtectionList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ReplayProtectionList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if m.Config != nil {
		{
			size, err := m.Config.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.GuardianSetList) > 0 {
		for iNdEx := len(m.GuardianSetList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.GuardianSetList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
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
	if len(m.GuardianSetList) > 0 {
		for _, e := range m.GuardianSetList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.Config != nil {
		l = m.Config.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.ReplayProtectionList) > 0 {
		for _, e := range m.ReplayProtectionList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.SequenceCounterList) > 0 {
		for _, e := range m.SequenceCounterList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.ConsensusGuardianSetIndex != nil {
		l = m.ConsensusGuardianSetIndex.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.GuardianValidatorList) > 0 {
		for _, e := range m.GuardianValidatorList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
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
				return fmt.Errorf("proto: wrong wireType = %d for field GuardianSetList", wireType)
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
			m.GuardianSetList = append(m.GuardianSetList, GuardianSet{})
			if err := m.GuardianSetList[len(m.GuardianSetList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Config", wireType)
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
			if m.Config == nil {
				m.Config = &Config{}
			}
			if err := m.Config.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReplayProtectionList", wireType)
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
			m.ReplayProtectionList = append(m.ReplayProtectionList, ReplayProtection{})
			if err := m.ReplayProtectionList[len(m.ReplayProtectionList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SequenceCounterList", wireType)
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
			m.SequenceCounterList = append(m.SequenceCounterList, SequenceCounter{})
			if err := m.SequenceCounterList[len(m.SequenceCounterList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConsensusGuardianSetIndex", wireType)
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
			if m.ConsensusGuardianSetIndex == nil {
				m.ConsensusGuardianSetIndex = &ConsensusGuardianSetIndex{}
			}
			if err := m.ConsensusGuardianSetIndex.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GuardianValidatorList", wireType)
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
			m.GuardianValidatorList = append(m.GuardianValidatorList, GuardianValidator{})
			if err := m.GuardianValidatorList[len(m.GuardianValidatorList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
