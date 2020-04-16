// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: codec/std/keys.proto

package std

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

// PubKey defines the application-level PubKey type.
type PublicKey struct {
	// TODO: uncomment once tendermint has made the key types slices instead of arrays
	// option (cosmos_proto.interface_type) = "*github.com/tendermint/tendermint/crypto.PubKey";
	// sum defines a set of all acceptable concrete PubKey implementations.
	//
	// Types that are valid to be assigned to Sum:
	//	*PublicKey_Ed25519
	//	*PublicKey_Secp256K1
	//	*PublicKey_Sr25519
	//	*PublicKey_Multisig
	Sum isPublicKey_Sum `protobuf_oneof:"sum"`
}

func (m *PublicKey) Reset()         { *m = PublicKey{} }
func (m *PublicKey) String() string { return proto.CompactTextString(m) }
func (*PublicKey) ProtoMessage()    {}
func (*PublicKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5440939592b0641, []int{0}
}
func (m *PublicKey) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PublicKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PublicKey.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PublicKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublicKey.Merge(m, src)
}
func (m *PublicKey) XXX_Size() int {
	return m.Size()
}
func (m *PublicKey) XXX_DiscardUnknown() {
	xxx_messageInfo_PublicKey.DiscardUnknown(m)
}

var xxx_messageInfo_PublicKey proto.InternalMessageInfo

type isPublicKey_Sum interface {
	isPublicKey_Sum()
	MarshalTo([]byte) (int, error)
	Size() int
}

type PublicKey_Ed25519 struct {
	Ed25519 []byte `protobuf:"bytes,1,opt,name=ed25519,proto3,oneof" json:"ed25519,omitempty"`
}
type PublicKey_Secp256K1 struct {
	Secp256K1 []byte `protobuf:"bytes,2,opt,name=secp256k1,proto3,oneof" json:"secp256k1,omitempty"`
}
type PublicKey_Sr25519 struct {
	Sr25519 []byte `protobuf:"bytes,3,opt,name=sr25519,proto3,oneof" json:"sr25519,omitempty"`
}
type PublicKey_Multisig struct {
	Multisig *PubKeyMultisigThreshold `protobuf:"bytes,4,opt,name=multisig,proto3,oneof" json:"multisig,omitempty"`
}

func (*PublicKey_Ed25519) isPublicKey_Sum()   {}
func (*PublicKey_Secp256K1) isPublicKey_Sum() {}
func (*PublicKey_Sr25519) isPublicKey_Sum()   {}
func (*PublicKey_Multisig) isPublicKey_Sum()  {}

func (m *PublicKey) GetSum() isPublicKey_Sum {
	if m != nil {
		return m.Sum
	}
	return nil
}

func (m *PublicKey) GetEd25519() []byte {
	if x, ok := m.GetSum().(*PublicKey_Ed25519); ok {
		return x.Ed25519
	}
	return nil
}

func (m *PublicKey) GetSecp256K1() []byte {
	if x, ok := m.GetSum().(*PublicKey_Secp256K1); ok {
		return x.Secp256K1
	}
	return nil
}

func (m *PublicKey) GetSr25519() []byte {
	if x, ok := m.GetSum().(*PublicKey_Sr25519); ok {
		return x.Sr25519
	}
	return nil
}

func (m *PublicKey) GetMultisig() *PubKeyMultisigThreshold {
	if x, ok := m.GetSum().(*PublicKey_Multisig); ok {
		return x.Multisig
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*PublicKey) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*PublicKey_Ed25519)(nil),
		(*PublicKey_Secp256K1)(nil),
		(*PublicKey_Sr25519)(nil),
		(*PublicKey_Multisig)(nil),
	}
}

// PrivKey defines the application-level PrivKey type.
type PrivateKey struct {
	// sum defines a set of all acceptable concrete PrivKey implementations.
	//
	// Types that are valid to be assigned to Sum:
	//	*PrivateKey_Ed25519
	//	*PrivateKey_Secp256K1
	//	*PrivateKey_Sr25519
	Sum isPrivateKey_Sum `protobuf_oneof:"sum"`
}

func (m *PrivateKey) Reset()         { *m = PrivateKey{} }
func (m *PrivateKey) String() string { return proto.CompactTextString(m) }
func (*PrivateKey) ProtoMessage()    {}
func (*PrivateKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5440939592b0641, []int{1}
}
func (m *PrivateKey) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PrivateKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PrivateKey.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PrivateKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrivateKey.Merge(m, src)
}
func (m *PrivateKey) XXX_Size() int {
	return m.Size()
}
func (m *PrivateKey) XXX_DiscardUnknown() {
	xxx_messageInfo_PrivateKey.DiscardUnknown(m)
}

var xxx_messageInfo_PrivateKey proto.InternalMessageInfo

type isPrivateKey_Sum interface {
	isPrivateKey_Sum()
	MarshalTo([]byte) (int, error)
	Size() int
}

type PrivateKey_Ed25519 struct {
	Ed25519 []byte `protobuf:"bytes,1,opt,name=ed25519,proto3,oneof" json:"ed25519,omitempty"`
}
type PrivateKey_Secp256K1 struct {
	Secp256K1 []byte `protobuf:"bytes,2,opt,name=secp256k1,proto3,oneof" json:"secp256k1,omitempty"`
}
type PrivateKey_Sr25519 struct {
	Sr25519 []byte `protobuf:"bytes,3,opt,name=sr25519,proto3,oneof" json:"sr25519,omitempty"`
}

func (*PrivateKey_Ed25519) isPrivateKey_Sum()   {}
func (*PrivateKey_Secp256K1) isPrivateKey_Sum() {}
func (*PrivateKey_Sr25519) isPrivateKey_Sum()   {}

func (m *PrivateKey) GetSum() isPrivateKey_Sum {
	if m != nil {
		return m.Sum
	}
	return nil
}

func (m *PrivateKey) GetEd25519() []byte {
	if x, ok := m.GetSum().(*PrivateKey_Ed25519); ok {
		return x.Ed25519
	}
	return nil
}

func (m *PrivateKey) GetSecp256K1() []byte {
	if x, ok := m.GetSum().(*PrivateKey_Secp256K1); ok {
		return x.Secp256K1
	}
	return nil
}

func (m *PrivateKey) GetSr25519() []byte {
	if x, ok := m.GetSum().(*PrivateKey_Sr25519); ok {
		return x.Sr25519
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*PrivateKey) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*PrivateKey_Ed25519)(nil),
		(*PrivateKey_Secp256K1)(nil),
		(*PrivateKey_Sr25519)(nil),
	}
}

// PubKeyMultisigThreshold implements a K of N threshold multisig.
type PubKeyMultisigThreshold struct {
	K       uint32       `protobuf:"varint,1,opt,name=threshold,proto3" json:"threshold,omitempty" yaml:"threshold"`
	PubKeys []*PublicKey `protobuf:"bytes,2,rep,name=pubkeys,proto3" json:"pubkeys,omitempty" yaml:"pubkeys"`
}

func (m *PubKeyMultisigThreshold) Reset()         { *m = PubKeyMultisigThreshold{} }
func (m *PubKeyMultisigThreshold) String() string { return proto.CompactTextString(m) }
func (*PubKeyMultisigThreshold) ProtoMessage()    {}
func (*PubKeyMultisigThreshold) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5440939592b0641, []int{2}
}
func (m *PubKeyMultisigThreshold) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PubKeyMultisigThreshold) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PubKeyMultisigThreshold.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PubKeyMultisigThreshold) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PubKeyMultisigThreshold.Merge(m, src)
}
func (m *PubKeyMultisigThreshold) XXX_Size() int {
	return m.Size()
}
func (m *PubKeyMultisigThreshold) XXX_DiscardUnknown() {
	xxx_messageInfo_PubKeyMultisigThreshold.DiscardUnknown(m)
}

var xxx_messageInfo_PubKeyMultisigThreshold proto.InternalMessageInfo

func (m *PubKeyMultisigThreshold) GetK() uint32 {
	if m != nil {
		return m.K
	}
	return 0
}

func (m *PubKeyMultisigThreshold) GetPubKeys() []*PublicKey {
	if m != nil {
		return m.PubKeys
	}
	return nil
}

// CompactBitArray is an implementation of a space efficient bit array.
// This is used to ensure that the encoded data takes up a minimal amount of
// space after proto encoding.
// This is not thread safe, and is not intended for concurrent usage.
type CompactBitArray struct {
	ExtraBitsStored []byte `protobuf:"bytes,1,opt,name=extra_bits,json=extraBits,proto3" json:"extra_bits,omitempty" yaml:"extra_bits"`
	Elems           []byte `protobuf:"bytes,2,opt,name=bits,proto3" json:"bits,omitempty" yaml:"bits"`
}

func (m *CompactBitArray) Reset()         { *m = CompactBitArray{} }
func (m *CompactBitArray) String() string { return proto.CompactTextString(m) }
func (*CompactBitArray) ProtoMessage()    {}
func (*CompactBitArray) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5440939592b0641, []int{3}
}
func (m *CompactBitArray) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CompactBitArray.Unmarshal(m, b)
}
func (m *CompactBitArray) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CompactBitArray.Marshal(b, m, deterministic)
}
func (m *CompactBitArray) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CompactBitArray.Merge(m, src)
}
func (m *CompactBitArray) XXX_Size() int {
	return xxx_messageInfo_CompactBitArray.Size(m)
}
func (m *CompactBitArray) XXX_DiscardUnknown() {
	xxx_messageInfo_CompactBitArray.DiscardUnknown(m)
}

var xxx_messageInfo_CompactBitArray proto.InternalMessageInfo

func (m *CompactBitArray) GetExtraBitsStored() []byte {
	if m != nil {
		return m.ExtraBitsStored
	}
	return nil
}

func (m *CompactBitArray) GetElems() []byte {
	if m != nil {
		return m.Elems
	}
	return nil
}

// Multisignature is used to represent the signature object used in the multisigs.
// Sigs is a list of signatures, sorted by corresponding index.message PrivKeyEd25519 {
type Multisignature struct {
	BitArray *CompactBitArray `protobuf:"bytes,1,opt,name=bitarray,proto3" json:"bitarray,omitempty"`
	Sigs     [][]byte         `protobuf:"bytes,2,rep,name=sigs,proto3" json:"sigs,omitempty"`
}

func (m *Multisignature) Reset()         { *m = Multisignature{} }
func (m *Multisignature) String() string { return proto.CompactTextString(m) }
func (*Multisignature) ProtoMessage()    {}
func (*Multisignature) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5440939592b0641, []int{4}
}
func (m *Multisignature) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Multisignature) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Multisignature.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Multisignature) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Multisignature.Merge(m, src)
}
func (m *Multisignature) XXX_Size() int {
	return m.Size()
}
func (m *Multisignature) XXX_DiscardUnknown() {
	xxx_messageInfo_Multisignature.DiscardUnknown(m)
}

var xxx_messageInfo_Multisignature proto.InternalMessageInfo

func (m *Multisignature) GetBitArray() *CompactBitArray {
	if m != nil {
		return m.BitArray
	}
	return nil
}

func (m *Multisignature) GetSigs() [][]byte {
	if m != nil {
		return m.Sigs
	}
	return nil
}

func init() {
	proto.RegisterType((*PublicKey)(nil), "cosmos_sdk.codec.std.v1.PublicKey")
	proto.RegisterType((*PrivateKey)(nil), "cosmos_sdk.codec.std.v1.PrivateKey")
	proto.RegisterType((*PubKeyMultisigThreshold)(nil), "cosmos_sdk.codec.std.v1.PubKeyMultisigThreshold")
	proto.RegisterType((*CompactBitArray)(nil), "cosmos_sdk.codec.std.v1.CompactBitArray")
	proto.RegisterType((*Multisignature)(nil), "cosmos_sdk.codec.std.v1.Multisignature")
}

func init() { proto.RegisterFile("codec/std/keys.proto", fileDescriptor_c5440939592b0641) }

var fileDescriptor_c5440939592b0641 = []byte{
	// 518 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x53, 0xbd, 0x8e, 0xd3, 0x30,
	0x1c, 0x8f, 0xef, 0x7a, 0xb4, 0x75, 0xcb, 0x1d, 0x58, 0x48, 0x57, 0x4e, 0x22, 0xae, 0x32, 0x9c,
	0x3a, 0x40, 0x42, 0x8b, 0x0a, 0xa2, 0x13, 0x04, 0x9d, 0x38, 0xa9, 0x02, 0x55, 0x81, 0x09, 0x86,
	0x2a, 0x1f, 0x56, 0x6b, 0xb5, 0xc1, 0x91, 0xed, 0x9c, 0xc8, 0xc2, 0xcc, 0x78, 0x62, 0x62, 0xe4,
	0x9e, 0x80, 0x9d, 0x27, 0x60, 0xbc, 0x91, 0x29, 0x42, 0xe9, 0x1b, 0xf4, 0x09, 0x50, 0xe2, 0xa4,
	0x95, 0x90, 0xca, 0xc6, 0x14, 0xfb, 0xef, 0xdf, 0x57, 0x7e, 0x96, 0xe1, 0x1d, 0x9f, 0x05, 0xc4,
	0xb7, 0x84, 0x0c, 0xac, 0x05, 0x49, 0x84, 0x19, 0x71, 0x26, 0x19, 0x3a, 0xf6, 0x99, 0x08, 0x99,
	0x98, 0x8a, 0x60, 0x61, 0x16, 0x00, 0x53, 0xc8, 0xc0, 0xbc, 0xe8, 0x9f, 0x9c, 0xca, 0x39, 0xe5,
	0xc1, 0x34, 0x72, 0xb9, 0x4c, 0xac, 0x02, 0x6b, 0xcd, 0xd8, 0x8c, 0x6d, 0x57, 0x4a, 0xc0, 0xf8,
	0x01, 0x60, 0x73, 0x12, 0x7b, 0x4b, 0xea, 0x8f, 0x49, 0x82, 0x4e, 0x60, 0x9d, 0x04, 0x83, 0xe1,
	0xb0, 0xff, 0xb4, 0x03, 0xba, 0xa0, 0xd7, 0x3e, 0xd7, 0x9c, 0x6a, 0x80, 0x74, 0xd8, 0x14, 0xc4,
	0x8f, 0x06, 0xc3, 0xc7, 0x8b, 0x7e, 0x67, 0xaf, 0x3c, 0xdd, 0x8e, 0x72, 0xae, 0xe0, 0x8a, 0xbb,
	0x5f, 0x71, 0xcb, 0x01, 0x7a, 0x0d, 0x1b, 0x61, 0xbc, 0x94, 0x54, 0xd0, 0x59, 0xa7, 0xd6, 0x05,
	0xbd, 0xd6, 0xe0, 0xa1, 0xb9, 0x23, 0xb9, 0x39, 0x89, 0xbd, 0x31, 0x49, 0x5e, 0x95, 0xf0, 0xb7,
	0x73, 0x4e, 0xc4, 0x9c, 0x2d, 0x83, 0x73, 0xcd, 0xd9, 0x68, 0xd8, 0x07, 0x70, 0x5f, 0xc4, 0xa1,
	0xb1, 0x80, 0x70, 0xc2, 0xe9, 0x85, 0x2b, 0xc9, 0x7f, 0x0c, 0x5f, 0x99, 0x7d, 0x07, 0xf0, 0x78,
	0x47, 0x36, 0xf4, 0x04, 0x36, 0x65, 0xb5, 0x29, 0xcc, 0x6f, 0xda, 0x77, 0xb3, 0x14, 0x83, 0xf1,
	0x3a, 0xc5, 0xb7, 0x12, 0x37, 0x5c, 0x8e, 0x8c, 0xcd, 0xb9, 0xe1, 0x6c, 0xb1, 0xe8, 0x3d, 0xac,
	0x47, 0xb1, 0x97, 0x5f, 0x68, 0x67, 0xaf, 0xbb, 0xdf, 0x6b, 0x0d, 0x8c, 0x7f, 0xf5, 0xa2, 0x6e,
	0xc9, 0xbe, 0x97, 0xa5, 0xb8, 0xae, 0xa2, 0x88, 0x75, 0x8a, 0x0f, 0x95, 0x41, 0xa9, 0x63, 0x38,
	0x95, 0xa2, 0xf1, 0x05, 0xc0, 0xa3, 0x17, 0x2c, 0x8c, 0x5c, 0x5f, 0xda, 0x54, 0x3e, 0xe7, 0xdc,
	0x4d, 0xd0, 0x4b, 0x08, 0xc9, 0x47, 0xc9, 0xdd, 0xa9, 0x47, 0xa5, 0x50, 0x3d, 0xd9, 0xbd, 0x2c,
	0xc5, 0x47, 0x67, 0xf9, 0xd4, 0xa6, 0x52, 0xbc, 0x91, 0x8c, 0x93, 0x60, 0x9d, 0xe2, 0xdb, 0x4a,
	0x77, 0x0b, 0x37, 0x9c, 0x26, 0xa9, 0x50, 0xe8, 0x3e, 0xac, 0x15, 0x12, 0x45, 0x99, 0x76, 0x27,
	0x4b, 0xf1, 0xc1, 0xd9, 0x92, 0x84, 0x79, 0xa0, 0x96, 0x22, 0x2a, 0x4a, 0x81, 0x1a, 0x35, 0x3e,
	0x5f, 0x61, 0xed, 0xf2, 0x0a, 0x6b, 0xc6, 0x27, 0x78, 0x58, 0xf5, 0xf7, 0xc1, 0x95, 0x31, 0x27,
	0xc8, 0x81, 0x0d, 0x8f, 0x4a, 0x37, 0x8f, 0x57, 0x04, 0x6a, 0x0d, 0x7a, 0x3b, 0x4b, 0xf8, 0xeb,
	0x77, 0xec, 0x76, 0x96, 0xe2, 0x46, 0xb5, 0x73, 0x36, 0x3a, 0x08, 0xc1, 0x9a, 0xa0, 0x33, 0x55,
	0x6a, 0xdb, 0x29, 0xd6, 0xa3, 0xda, 0xd7, 0x6f, 0x18, 0xd8, 0xcf, 0x7e, 0x66, 0x3a, 0xb8, 0xce,
	0x74, 0xf0, 0x3b, 0xd3, 0xc1, 0xe5, 0x4a, 0xd7, 0xae, 0x57, 0xba, 0xf6, 0x6b, 0xa5, 0x6b, 0xef,
	0x4e, 0x67, 0x54, 0xce, 0x63, 0xcf, 0xf4, 0x59, 0x68, 0x29, 0xff, 0xf2, 0xf3, 0x40, 0x04, 0x0b,
	0x6b, 0xf3, 0xfc, 0xbc, 0x1b, 0xc5, 0xcb, 0x79, 0xf4, 0x27, 0x00, 0x00, 0xff, 0xff, 0x9c, 0xc5,
	0x2d, 0x98, 0x92, 0x03, 0x00, 0x00,
}

func (m *PublicKey) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PublicKey) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PublicKey) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Sum != nil {
		{
			size := m.Sum.Size()
			i -= size
			if _, err := m.Sum.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *PublicKey_Ed25519) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PublicKey_Ed25519) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Ed25519 != nil {
		i -= len(m.Ed25519)
		copy(dAtA[i:], m.Ed25519)
		i = encodeVarintKeys(dAtA, i, uint64(len(m.Ed25519)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *PublicKey_Secp256K1) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PublicKey_Secp256K1) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Secp256K1 != nil {
		i -= len(m.Secp256K1)
		copy(dAtA[i:], m.Secp256K1)
		i = encodeVarintKeys(dAtA, i, uint64(len(m.Secp256K1)))
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *PublicKey_Sr25519) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PublicKey_Sr25519) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Sr25519 != nil {
		i -= len(m.Sr25519)
		copy(dAtA[i:], m.Sr25519)
		i = encodeVarintKeys(dAtA, i, uint64(len(m.Sr25519)))
		i--
		dAtA[i] = 0x1a
	}
	return len(dAtA) - i, nil
}
func (m *PublicKey_Multisig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PublicKey_Multisig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Multisig != nil {
		{
			size, err := m.Multisig.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintKeys(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	return len(dAtA) - i, nil
}
func (m *PrivateKey) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PrivateKey) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PrivateKey) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Sum != nil {
		{
			size := m.Sum.Size()
			i -= size
			if _, err := m.Sum.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *PrivateKey_Ed25519) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PrivateKey_Ed25519) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Ed25519 != nil {
		i -= len(m.Ed25519)
		copy(dAtA[i:], m.Ed25519)
		i = encodeVarintKeys(dAtA, i, uint64(len(m.Ed25519)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *PrivateKey_Secp256K1) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PrivateKey_Secp256K1) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Secp256K1 != nil {
		i -= len(m.Secp256K1)
		copy(dAtA[i:], m.Secp256K1)
		i = encodeVarintKeys(dAtA, i, uint64(len(m.Secp256K1)))
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *PrivateKey_Sr25519) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PrivateKey_Sr25519) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Sr25519 != nil {
		i -= len(m.Sr25519)
		copy(dAtA[i:], m.Sr25519)
		i = encodeVarintKeys(dAtA, i, uint64(len(m.Sr25519)))
		i--
		dAtA[i] = 0x1a
	}
	return len(dAtA) - i, nil
}
func (m *PubKeyMultisigThreshold) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PubKeyMultisigThreshold) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PubKeyMultisigThreshold) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PubKeys) > 0 {
		for iNdEx := len(m.PubKeys) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PubKeys[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintKeys(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.K != 0 {
		i = encodeVarintKeys(dAtA, i, uint64(m.K))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Multisignature) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Multisignature) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Multisignature) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Sigs) > 0 {
		for iNdEx := len(m.Sigs) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Sigs[iNdEx])
			copy(dAtA[i:], m.Sigs[iNdEx])
			i = encodeVarintKeys(dAtA, i, uint64(len(m.Sigs[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if m.BitArray != nil {
		{
			size, err := m.BitArray.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintKeys(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintKeys(dAtA []byte, offset int, v uint64) int {
	offset -= sovKeys(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PublicKey) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Sum != nil {
		n += m.Sum.Size()
	}
	return n
}

func (m *PublicKey_Ed25519) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Ed25519 != nil {
		l = len(m.Ed25519)
		n += 1 + l + sovKeys(uint64(l))
	}
	return n
}
func (m *PublicKey_Secp256K1) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Secp256K1 != nil {
		l = len(m.Secp256K1)
		n += 1 + l + sovKeys(uint64(l))
	}
	return n
}
func (m *PublicKey_Sr25519) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Sr25519 != nil {
		l = len(m.Sr25519)
		n += 1 + l + sovKeys(uint64(l))
	}
	return n
}
func (m *PublicKey_Multisig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Multisig != nil {
		l = m.Multisig.Size()
		n += 1 + l + sovKeys(uint64(l))
	}
	return n
}
func (m *PrivateKey) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Sum != nil {
		n += m.Sum.Size()
	}
	return n
}

func (m *PrivateKey_Ed25519) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Ed25519 != nil {
		l = len(m.Ed25519)
		n += 1 + l + sovKeys(uint64(l))
	}
	return n
}
func (m *PrivateKey_Secp256K1) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Secp256K1 != nil {
		l = len(m.Secp256K1)
		n += 1 + l + sovKeys(uint64(l))
	}
	return n
}
func (m *PrivateKey_Sr25519) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Sr25519 != nil {
		l = len(m.Sr25519)
		n += 1 + l + sovKeys(uint64(l))
	}
	return n
}
func (m *PubKeyMultisigThreshold) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.K != 0 {
		n += 1 + sovKeys(uint64(m.K))
	}
	if len(m.PubKeys) > 0 {
		for _, e := range m.PubKeys {
			l = e.Size()
			n += 1 + l + sovKeys(uint64(l))
		}
	}
	return n
}

func (m *CompactBitArray) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ExtraBitsStored)
	if l > 0 {
		n += 1 + l + sovKeys(uint64(l))
	}
	l = len(m.Elems)
	if l > 0 {
		n += 1 + l + sovKeys(uint64(l))
	}
	return n
}

func (m *Multisignature) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BitArray != nil {
		l = m.BitArray.Size()
		n += 1 + l + sovKeys(uint64(l))
	}
	if len(m.Sigs) > 0 {
		for _, b := range m.Sigs {
			l = len(b)
			n += 1 + l + sovKeys(uint64(l))
		}
	}
	return n
}

func sovKeys(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozKeys(x uint64) (n int) {
	return sovKeys(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PublicKey) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowKeys
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
			return fmt.Errorf("proto: PublicKey: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PublicKey: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ed25519", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeys
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
				return ErrInvalidLengthKeys
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthKeys
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := make([]byte, postIndex-iNdEx)
			copy(v, dAtA[iNdEx:postIndex])
			m.Sum = &PublicKey_Ed25519{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Secp256K1", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeys
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
				return ErrInvalidLengthKeys
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthKeys
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := make([]byte, postIndex-iNdEx)
			copy(v, dAtA[iNdEx:postIndex])
			m.Sum = &PublicKey_Secp256K1{v}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sr25519", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeys
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
				return ErrInvalidLengthKeys
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthKeys
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := make([]byte, postIndex-iNdEx)
			copy(v, dAtA[iNdEx:postIndex])
			m.Sum = &PublicKey_Sr25519{v}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Multisig", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeys
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
				return ErrInvalidLengthKeys
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthKeys
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &PubKeyMultisigThreshold{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &PublicKey_Multisig{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipKeys(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthKeys
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthKeys
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
func (m *PrivateKey) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowKeys
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
			return fmt.Errorf("proto: PrivateKey: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PrivateKey: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ed25519", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeys
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
				return ErrInvalidLengthKeys
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthKeys
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := make([]byte, postIndex-iNdEx)
			copy(v, dAtA[iNdEx:postIndex])
			m.Sum = &PrivateKey_Ed25519{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Secp256K1", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeys
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
				return ErrInvalidLengthKeys
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthKeys
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := make([]byte, postIndex-iNdEx)
			copy(v, dAtA[iNdEx:postIndex])
			m.Sum = &PrivateKey_Secp256K1{v}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sr25519", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeys
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
				return ErrInvalidLengthKeys
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthKeys
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := make([]byte, postIndex-iNdEx)
			copy(v, dAtA[iNdEx:postIndex])
			m.Sum = &PrivateKey_Sr25519{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipKeys(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthKeys
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthKeys
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
func (m *PubKeyMultisigThreshold) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowKeys
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
			return fmt.Errorf("proto: PubKeyMultisigThreshold: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PubKeyMultisigThreshold: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field K", wireType)
			}
			m.K = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeys
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.K |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PubKeys", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeys
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
				return ErrInvalidLengthKeys
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthKeys
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PubKeys = append(m.PubKeys, &PublicKey{})
			if err := m.PubKeys[len(m.PubKeys)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipKeys(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthKeys
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthKeys
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
func (m *Multisignature) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowKeys
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
			return fmt.Errorf("proto: Multisignature: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Multisignature: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BitArray", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeys
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
				return ErrInvalidLengthKeys
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthKeys
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BitArray == nil {
				m.BitArray = &CompactBitArray{}
			}
			if err := m.BitArray.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sigs", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeys
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
				return ErrInvalidLengthKeys
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthKeys
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sigs = append(m.Sigs, make([]byte, postIndex-iNdEx))
			copy(m.Sigs[len(m.Sigs)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipKeys(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthKeys
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthKeys
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
func skipKeys(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowKeys
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
					return 0, ErrIntOverflowKeys
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
					return 0, ErrIntOverflowKeys
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
				return 0, ErrInvalidLengthKeys
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupKeys
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthKeys
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthKeys        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowKeys          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupKeys = fmt.Errorf("proto: unexpected end of group")
)
