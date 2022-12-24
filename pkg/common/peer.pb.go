// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: common/peer.proto

// Package common defines commonly used types agnostic to the node role on the Sonr network.

package common

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

// Peers Active Status
type Peer_Type int32

const (
	Peer_UNKNOWN     Peer_Type = 0
	Peer_HIGHWAY     Peer_Type = 1
	Peer_MOTOR       Peer_Type = 2
	Peer_VALIDATOR   Peer_Type = 3
	Peer_THIRD_PARTY Peer_Type = 4
)

var Peer_Type_name = map[int32]string{
	0: "UNKNOWN",
	1: "HIGHWAY",
	2: "MOTOR",
	3: "VALIDATOR",
	4: "THIRD_PARTY",
}

var Peer_Type_value = map[string]int32{
	"UNKNOWN":     0,
	"HIGHWAY":     1,
	"MOTOR":       2,
	"VALIDATOR":   3,
	"THIRD_PARTY": 4,
}

func (x Peer_Type) String() string {
	return proto.EnumName(Peer_Type_name, int32(x))
}

func (Peer_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_253be8a7e67f50a6, []int{0, 0}
}

// Basic Info Sent to Peers to Establish Connections
type Peer struct {
	Id        string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string    `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	PeerId    string    `protobuf:"bytes,3,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty"`
	Multiaddr string    `protobuf:"bytes,4,opt,name=multiaddr,proto3" json:"multiaddr,omitempty"`
	Type      Peer_Type `protobuf:"varint,5,opt,name=type,proto3,enum=sonrhq.common.Peer_Type" json:"type,omitempty"`
}

func (m *Peer) Reset()         { *m = Peer{} }
func (m *Peer) String() string { return proto.CompactTextString(m) }
func (*Peer) ProtoMessage()    {}
func (*Peer) Descriptor() ([]byte, []int) {
	return fileDescriptor_253be8a7e67f50a6, []int{0}
}
func (m *Peer) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Peer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Peer.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Peer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Peer.Merge(m, src)
}
func (m *Peer) XXX_Size() int {
	return m.Size()
}
func (m *Peer) XXX_DiscardUnknown() {
	xxx_messageInfo_Peer.DiscardUnknown(m)
}

var xxx_messageInfo_Peer proto.InternalMessageInfo

func (m *Peer) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Peer) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Peer) GetPeerId() string {
	if m != nil {
		return m.PeerId
	}
	return ""
}

func (m *Peer) GetMultiaddr() string {
	if m != nil {
		return m.Multiaddr
	}
	return ""
}

func (m *Peer) GetType() Peer_Type {
	if m != nil {
		return m.Type
	}
	return Peer_UNKNOWN
}

// AuthInfo is a object used by Motor clients to store authentication information in Biometric storage
type AuthInfo struct {
	Address   string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Did       string `protobuf:"bytes,2,opt,name=did,proto3" json:"did,omitempty"`
	AesDscKey []byte `protobuf:"bytes,3,opt,name=aes_dsc_key,json=aesDscKey,proto3" json:"aes_dsc_key,omitempty"`
	AesPskKey []byte `protobuf:"bytes,4,opt,name=aes_psk_key,json=aesPskKey,proto3" json:"aes_psk_key,omitempty"`
	Password  string `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
	Timestamp int64  `protobuf:"varint,6,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (m *AuthInfo) Reset()         { *m = AuthInfo{} }
func (m *AuthInfo) String() string { return proto.CompactTextString(m) }
func (*AuthInfo) ProtoMessage()    {}
func (*AuthInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_253be8a7e67f50a6, []int{1}
}
func (m *AuthInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AuthInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AuthInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AuthInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthInfo.Merge(m, src)
}
func (m *AuthInfo) XXX_Size() int {
	return m.Size()
}
func (m *AuthInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthInfo.DiscardUnknown(m)
}

var xxx_messageInfo_AuthInfo proto.InternalMessageInfo

func (m *AuthInfo) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *AuthInfo) GetDid() string {
	if m != nil {
		return m.Did
	}
	return ""
}

func (m *AuthInfo) GetAesDscKey() []byte {
	if m != nil {
		return m.AesDscKey
	}
	return nil
}

func (m *AuthInfo) GetAesPskKey() []byte {
	if m != nil {
		return m.AesPskKey
	}
	return nil
}

func (m *AuthInfo) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *AuthInfo) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func init() {
	proto.RegisterEnum("sonrhq.common.Peer_Type", Peer_Type_name, Peer_Type_value)
	proto.RegisterType((*Peer)(nil), "sonrhq.common.Peer")
	proto.RegisterType((*AuthInfo)(nil), "sonrhq.common.AuthInfo")
}

func init() { proto.RegisterFile("common/peer.proto", fileDescriptor_253be8a7e67f50a6) }

var fileDescriptor_253be8a7e67f50a6 = []byte{
	// 392 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0x4f, 0x8f, 0x93, 0x40,
	0x18, 0xc6, 0x19, 0x60, 0xdb, 0xe5, 0xad, 0xbb, 0xe2, 0x5c, 0x9c, 0x18, 0x43, 0x1a, 0x4e, 0x3d,
	0x28, 0x24, 0x7a, 0xf5, 0x82, 0x69, 0x62, 0x49, 0xb5, 0x6d, 0x26, 0xe8, 0x66, 0xbd, 0x10, 0xb6,
	0x33, 0x2e, 0x04, 0x81, 0x59, 0x86, 0xc6, 0xf0, 0x2d, 0xfc, 0x28, 0x7e, 0x0c, 0x8f, 0x7b, 0xf4,
	0xe0, 0xc1, 0xb4, 0x5f, 0xc4, 0xcc, 0x6c, 0xff, 0x64, 0x4f, 0xbc, 0xcf, 0xf3, 0x23, 0x33, 0xef,
	0xf3, 0x0c, 0x3c, 0x5b, 0x37, 0x55, 0xd5, 0xd4, 0xa1, 0xe0, 0xbc, 0x0d, 0x44, 0xdb, 0x74, 0x0d,
	0xbe, 0x90, 0x4d, 0xdd, 0xe6, 0x77, 0xc1, 0x03, 0xf1, 0xff, 0x22, 0xb0, 0x57, 0x9c, 0xb7, 0xf8,
	0x12, 0xcc, 0x82, 0x11, 0x34, 0x46, 0x13, 0x87, 0x9a, 0x05, 0xc3, 0x18, 0xec, 0x3a, 0xab, 0x38,
	0x31, 0xb5, 0xa3, 0x67, 0xfc, 0x1c, 0x86, 0xea, 0xa4, 0xb4, 0x60, 0xc4, 0xd2, 0xf6, 0x40, 0xc9,
	0x98, 0xe1, 0x97, 0xe0, 0x54, 0x9b, 0xef, 0x5d, 0x91, 0x31, 0xd6, 0x12, 0x5b, 0xa3, 0x93, 0x81,
	0x5f, 0x81, 0xdd, 0xf5, 0x82, 0x93, 0xb3, 0x31, 0x9a, 0x5c, 0xbe, 0x21, 0xc1, 0xa3, 0x0d, 0x02,
	0x75, 0x7b, 0x90, 0xf4, 0x82, 0x53, 0xfd, 0x97, 0x3f, 0x07, 0x5b, 0x29, 0x3c, 0x82, 0xe1, 0xe7,
	0xc5, 0x7c, 0xb1, 0xbc, 0x5a, 0xb8, 0x86, 0x12, 0xb3, 0xf8, 0xc3, 0xec, 0x2a, 0xba, 0x76, 0x11,
	0x76, 0xe0, 0xec, 0xd3, 0x32, 0x59, 0x52, 0xd7, 0xc4, 0x17, 0xe0, 0x7c, 0x89, 0x3e, 0xc6, 0xd3,
	0x48, 0x49, 0x0b, 0x3f, 0x85, 0x51, 0x32, 0x8b, 0xe9, 0x34, 0x5d, 0x45, 0x34, 0xb9, 0x76, 0x6d,
	0xff, 0x17, 0x82, 0xf3, 0x68, 0xd3, 0xe5, 0x71, 0xfd, 0xad, 0xc1, 0x04, 0x86, 0x6a, 0x1f, 0x2e,
	0xe5, 0x3e, 0xe7, 0x41, 0x62, 0x17, 0x2c, 0x56, 0xb0, 0x7d, 0x56, 0x35, 0x62, 0x0f, 0x46, 0x19,
	0x97, 0x29, 0x93, 0xeb, 0xb4, 0xe4, 0xbd, 0x8e, 0xfb, 0x84, 0x3a, 0x19, 0x97, 0x53, 0xb9, 0x9e,
	0xf3, 0xfe, 0xc0, 0x85, 0x2c, 0x35, 0xb7, 0x8f, 0x7c, 0x25, 0x4b, 0xc5, 0x5f, 0xc0, 0xb9, 0xc8,
	0xa4, 0xfc, 0xd1, 0xb4, 0x4c, 0xe7, 0x76, 0xe8, 0x51, 0xab, 0xb6, 0xba, 0xa2, 0xe2, 0xb2, 0xcb,
	0x2a, 0x41, 0x06, 0x63, 0x34, 0xb1, 0xe8, 0xc9, 0x78, 0xff, 0xee, 0xf7, 0xd6, 0x43, 0xf7, 0x5b,
	0x0f, 0xfd, 0xdb, 0x7a, 0xe8, 0xe7, 0xce, 0x33, 0xee, 0x77, 0x9e, 0xf1, 0x67, 0xe7, 0x19, 0x5f,
	0xfd, 0xdb, 0xa2, 0xcb, 0x37, 0x37, 0xaa, 0xb8, 0x50, 0x75, 0xf8, 0x3a, 0xbf, 0xd3, 0xdf, 0x50,
	0x94, 0xb7, 0xe1, 0x43, 0x9b, 0x37, 0x03, 0xfd, 0xca, 0x6f, 0xff, 0x07, 0x00, 0x00, 0xff, 0xff,
	0xdb, 0x90, 0x3a, 0x6a, 0xfa, 0x01, 0x00, 0x00,
}

func (m *Peer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Peer) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Peer) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Type != 0 {
		i = encodeVarintPeer(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Multiaddr) > 0 {
		i -= len(m.Multiaddr)
		copy(dAtA[i:], m.Multiaddr)
		i = encodeVarintPeer(dAtA, i, uint64(len(m.Multiaddr)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.PeerId) > 0 {
		i -= len(m.PeerId)
		copy(dAtA[i:], m.PeerId)
		i = encodeVarintPeer(dAtA, i, uint64(len(m.PeerId)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintPeer(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintPeer(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AuthInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuthInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AuthInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Timestamp != 0 {
		i = encodeVarintPeer(dAtA, i, uint64(m.Timestamp))
		i--
		dAtA[i] = 0x30
	}
	if len(m.Password) > 0 {
		i -= len(m.Password)
		copy(dAtA[i:], m.Password)
		i = encodeVarintPeer(dAtA, i, uint64(len(m.Password)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.AesPskKey) > 0 {
		i -= len(m.AesPskKey)
		copy(dAtA[i:], m.AesPskKey)
		i = encodeVarintPeer(dAtA, i, uint64(len(m.AesPskKey)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.AesDscKey) > 0 {
		i -= len(m.AesDscKey)
		copy(dAtA[i:], m.AesDscKey)
		i = encodeVarintPeer(dAtA, i, uint64(len(m.AesDscKey)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Did) > 0 {
		i -= len(m.Did)
		copy(dAtA[i:], m.Did)
		i = encodeVarintPeer(dAtA, i, uint64(len(m.Did)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintPeer(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPeer(dAtA []byte, offset int, v uint64) int {
	offset -= sovPeer(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Peer) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovPeer(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovPeer(uint64(l))
	}
	l = len(m.PeerId)
	if l > 0 {
		n += 1 + l + sovPeer(uint64(l))
	}
	l = len(m.Multiaddr)
	if l > 0 {
		n += 1 + l + sovPeer(uint64(l))
	}
	if m.Type != 0 {
		n += 1 + sovPeer(uint64(m.Type))
	}
	return n
}

func (m *AuthInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovPeer(uint64(l))
	}
	l = len(m.Did)
	if l > 0 {
		n += 1 + l + sovPeer(uint64(l))
	}
	l = len(m.AesDscKey)
	if l > 0 {
		n += 1 + l + sovPeer(uint64(l))
	}
	l = len(m.AesPskKey)
	if l > 0 {
		n += 1 + l + sovPeer(uint64(l))
	}
	l = len(m.Password)
	if l > 0 {
		n += 1 + l + sovPeer(uint64(l))
	}
	if m.Timestamp != 0 {
		n += 1 + sovPeer(uint64(m.Timestamp))
	}
	return n
}

func sovPeer(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPeer(x uint64) (n int) {
	return sovPeer(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Peer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPeer
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
			return fmt.Errorf("proto: Peer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Peer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPeer
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
				return ErrInvalidLengthPeer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPeer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPeer
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
				return ErrInvalidLengthPeer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPeer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeerId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPeer
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
				return ErrInvalidLengthPeer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPeer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PeerId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Multiaddr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPeer
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
				return ErrInvalidLengthPeer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPeer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Multiaddr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPeer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= Peer_Type(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPeer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPeer
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
func (m *AuthInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPeer
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
			return fmt.Errorf("proto: AuthInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPeer
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
				return ErrInvalidLengthPeer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPeer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Did", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPeer
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
				return ErrInvalidLengthPeer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPeer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Did = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AesDscKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPeer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthPeer
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPeer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AesDscKey = append(m.AesDscKey[:0], dAtA[iNdEx:postIndex]...)
			if m.AesDscKey == nil {
				m.AesDscKey = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AesPskKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPeer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthPeer
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPeer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AesPskKey = append(m.AesPskKey[:0], dAtA[iNdEx:postIndex]...)
			if m.AesPskKey == nil {
				m.AesPskKey = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Password", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPeer
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
				return ErrInvalidLengthPeer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPeer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Password = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPeer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timestamp |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPeer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPeer
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
func skipPeer(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPeer
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
					return 0, ErrIntOverflowPeer
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
					return 0, ErrIntOverflowPeer
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
				return 0, ErrInvalidLengthPeer
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPeer
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPeer
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPeer        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPeer          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPeer = fmt.Errorf("proto: unexpected end of group")
)