// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: passage3d/claim/v1beta1/genesis.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
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

// GenesisState defines the claim module's genesis state.
type GenesisState struct {
	// balance of the claim module's account
	ModuleAccountBalance types.Coin `protobuf:"bytes,1,opt,name=module_account_balance,json=moduleAccountBalance,proto3" json:"module_account_balance" yaml:"module_account_balance"`
	// params defines all the parameters of the module.
	Params Params `protobuf:"bytes,2,opt,name=params,proto3" json:"params" yaml:"params"`
	// list of claim records, one for every airdrop recipient
	ClaimRecords []ClaimRecord `protobuf:"bytes,3,rep,name=claim_records,json=claimRecords,proto3" json:"claim_records" yaml:"claim_records"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_3c24d42126f7b0be, []int{0}
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

func (m *GenesisState) GetModuleAccountBalance() types.Coin {
	if m != nil {
		return m.ModuleAccountBalance
	}
	return types.Coin{}
}

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetClaimRecords() []ClaimRecord {
	if m != nil {
		return m.ClaimRecords
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "passage3d.claim.v1beta1.GenesisState")
}

func init() {
	proto.RegisterFile("passage3d/claim/v1beta1/genesis.proto", fileDescriptor_3c24d42126f7b0be)
}

var fileDescriptor_3c24d42126f7b0be = []byte{
	// 350 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xc1, 0x4e, 0xc2, 0x30,
	0x18, 0xc7, 0x37, 0x48, 0x38, 0x0c, 0xb8, 0x2c, 0xa8, 0x48, 0xb4, 0x90, 0x05, 0x12, 0xc2, 0xa1,
	0x0d, 0x70, 0xf3, 0xe6, 0x30, 0xf1, 0x66, 0xcc, 0xbc, 0x79, 0x21, 0x5d, 0xd7, 0xcc, 0x25, 0xdb,
	0xba, 0xac, 0x65, 0x91, 0xb7, 0xf0, 0x8d, 0xbc, 0x72, 0xe4, 0xe8, 0x89, 0x18, 0x78, 0x03, 0x9f,
	0xc0, 0xac, 0xad, 0xa0, 0x89, 0xbb, 0xad, 0xfd, 0x7e, 0xfb, 0xfd, 0xbf, 0xaf, 0x9f, 0x35, 0xca,
	0x30, 0xe7, 0x38, 0xa4, 0xf3, 0x00, 0x91, 0x18, 0x47, 0x09, 0x2a, 0xa6, 0x3e, 0x15, 0x78, 0x8a,
	0x42, 0x9a, 0x52, 0x1e, 0x71, 0x98, 0xe5, 0x4c, 0x30, 0xfb, 0xe2, 0x88, 0x41, 0x89, 0x41, 0x8d,
	0xf5, 0x3a, 0x21, 0x0b, 0x99, 0x64, 0x50, 0xf9, 0xa5, 0xf0, 0x1e, 0x20, 0x8c, 0x27, 0x8c, 0x23,
	0x1f, 0x73, 0x7a, 0x34, 0x12, 0x16, 0xa5, 0xba, 0x3e, 0xac, 0x4a, 0xcd, 0x70, 0x8e, 0x13, 0x1d,
	0xda, 0x9b, 0x54, 0x51, 0xf2, 0xb4, 0xcc, 0x29, 0x61, 0x79, 0xa0, 0x58, 0xe7, 0xbd, 0x66, 0xb5,
	0xee, 0x55, 0xcb, 0x4f, 0x02, 0x0b, 0x6a, 0x17, 0xd6, 0x79, 0xc2, 0x82, 0x55, 0x4c, 0x97, 0x98,
	0x10, 0xb6, 0x4a, 0xc5, 0xd2, 0xc7, 0x31, 0x4e, 0x09, 0xed, 0x9a, 0x03, 0x73, 0xdc, 0x9c, 0x5d,
	0x42, 0xd5, 0x23, 0x2c, 0x7b, 0xfc, 0x19, 0x07, 0x2e, 0x58, 0x94, 0xba, 0xa3, 0xcd, 0xae, 0x6f,
	0x7c, 0xed, 0xfa, 0xd7, 0x6b, 0x9c, 0xc4, 0x37, 0xce, 0xff, 0x1a, 0xc7, 0xeb, 0xa8, 0xc2, 0xad,
	0xba, 0x77, 0xd5, 0xb5, 0xfd, 0x60, 0x35, 0xd4, 0x10, 0xdd, 0x9a, 0xcc, 0xe9, 0xc3, 0x8a, 0xa7,
	0x83, 0x8f, 0x12, 0x73, 0xcf, 0x74, 0x5a, 0x5b, 0xa5, 0xa9, 0x9f, 0x1d, 0x4f, 0x5b, 0xec, 0xd0,
	0x6a, 0xff, 0x1e, 0x97, 0x77, 0xeb, 0x83, 0xfa, 0xb8, 0x39, 0x1b, 0x56, 0x6a, 0x17, 0xe5, 0xc9,
	0x93, 0xb0, 0x7b, 0xa5, 0xdd, 0x1d, 0xe5, 0xfe, 0x23, 0x72, 0xbc, 0x16, 0x39, 0xa1, 0xdc, 0xbd,
	0xdb, 0xec, 0x81, 0xb9, 0xdd, 0x03, 0xf3, 0x73, 0x0f, 0xcc, 0xb7, 0x03, 0x30, 0xb6, 0x07, 0x60,
	0x7c, 0x1c, 0x80, 0xf1, 0x3c, 0x09, 0x23, 0xf1, 0xb2, 0xf2, 0x21, 0x61, 0x09, 0xa2, 0x69, 0x81,
	0x83, 0xa8, 0x40, 0xa7, 0xd5, 0xbc, 0xea, 0xe5, 0x88, 0x75, 0x46, 0xb9, 0xdf, 0x90, 0xeb, 0x98,
	0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0x60, 0x7c, 0x4e, 0xe3, 0x58, 0x02, 0x00, 0x00,
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
	if len(m.ClaimRecords) > 0 {
		for iNdEx := len(m.ClaimRecords) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ClaimRecords[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.ModuleAccountBalance.MarshalToSizedBuffer(dAtA[:i])
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
	l = m.ModuleAccountBalance.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.ClaimRecords) > 0 {
		for _, e := range m.ClaimRecords {
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
				return fmt.Errorf("proto: wrong wireType = %d for field ModuleAccountBalance", wireType)
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
			if err := m.ModuleAccountBalance.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
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
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimRecords", wireType)
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
			m.ClaimRecords = append(m.ClaimRecords, ClaimRecord{})
			if err := m.ClaimRecords[len(m.ClaimRecords)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
