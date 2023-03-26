// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: util/crypto/cryptoproto/protos/crypto.proto

package cryptoproto

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

type KeyType int32

const (
	KeyType_Ed25519Public  KeyType = 0
	KeyType_Ed25519Private KeyType = 1
	KeyType_AES            KeyType = 2
)

var KeyType_name = map[int32]string{
	0: "Ed25519Public",
	1: "Ed25519Private",
	2: "AES",
}

var KeyType_value = map[string]int32{
	"Ed25519Public":  0,
	"Ed25519Private": 1,
	"AES":            2,
}

func (x KeyType) String() string {
	return proto.EnumName(KeyType_name, int32(x))
}

func (KeyType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ddfeb19e486561de, []int{0}
}

type DerivationMethod int32

const (
	DerivationMethod_Slip21 DerivationMethod = 0
)

var DerivationMethod_name = map[int32]string{
	0: "Slip21",
}

var DerivationMethod_value = map[string]int32{
	"Slip21": 0,
}

func (x DerivationMethod) String() string {
	return proto.EnumName(DerivationMethod_name, int32(x))
}

func (DerivationMethod) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ddfeb19e486561de, []int{1}
}

type Key struct {
	Type KeyType `protobuf:"varint,1,opt,name=Type,proto3,enum=crypto.KeyType" json:"Type,omitempty"`
	Data []byte  `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (m *Key) Reset()         { *m = Key{} }
func (m *Key) String() string { return proto.CompactTextString(m) }
func (*Key) ProtoMessage()    {}
func (*Key) Descriptor() ([]byte, []int) {
	return fileDescriptor_ddfeb19e486561de, []int{0}
}
func (m *Key) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Key) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Key.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Key) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Key.Merge(m, src)
}
func (m *Key) XXX_Size() int {
	return m.Size()
}
func (m *Key) XXX_DiscardUnknown() {
	xxx_messageInfo_Key.DiscardUnknown(m)
}

var xxx_messageInfo_Key proto.InternalMessageInfo

func (m *Key) GetType() KeyType {
	if m != nil {
		return m.Type
	}
	return KeyType_Ed25519Public
}

func (m *Key) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type KeyDerivation struct {
	Method         DerivationMethod `protobuf:"varint,1,opt,name=method,proto3,enum=crypto.DerivationMethod" json:"method,omitempty"`
	DerivationPath string           `protobuf:"bytes,2,opt,name=derivationPath,proto3" json:"derivationPath,omitempty"`
}

func (m *KeyDerivation) Reset()         { *m = KeyDerivation{} }
func (m *KeyDerivation) String() string { return proto.CompactTextString(m) }
func (*KeyDerivation) ProtoMessage()    {}
func (*KeyDerivation) Descriptor() ([]byte, []int) {
	return fileDescriptor_ddfeb19e486561de, []int{1}
}
func (m *KeyDerivation) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *KeyDerivation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_KeyDerivation.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *KeyDerivation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyDerivation.Merge(m, src)
}
func (m *KeyDerivation) XXX_Size() int {
	return m.Size()
}
func (m *KeyDerivation) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyDerivation.DiscardUnknown(m)
}

var xxx_messageInfo_KeyDerivation proto.InternalMessageInfo

func (m *KeyDerivation) GetMethod() DerivationMethod {
	if m != nil {
		return m.Method
	}
	return DerivationMethod_Slip21
}

func (m *KeyDerivation) GetDerivationPath() string {
	if m != nil {
		return m.DerivationPath
	}
	return ""
}

func init() {
	proto.RegisterEnum("crypto.KeyType", KeyType_name, KeyType_value)
	proto.RegisterEnum("crypto.DerivationMethod", DerivationMethod_name, DerivationMethod_value)
	proto.RegisterType((*Key)(nil), "crypto.Key")
	proto.RegisterType((*KeyDerivation)(nil), "crypto.KeyDerivation")
}

func init() {
	proto.RegisterFile("util/crypto/cryptoproto/protos/crypto.proto", fileDescriptor_ddfeb19e486561de)
}

var fileDescriptor_ddfeb19e486561de = []byte{
	// 263 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x2e, 0x2d, 0xc9, 0xcc,
	0xd1, 0x4f, 0x2e, 0xaa, 0x2c, 0x28, 0xc9, 0x87, 0x52, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0xfa, 0x60,
	0xb2, 0x18, 0x2a, 0xa4, 0x07, 0xe6, 0x09, 0xb1, 0x41, 0x78, 0x4a, 0x76, 0x5c, 0xcc, 0xde, 0xa9,
	0x95, 0x42, 0xca, 0x5c, 0x2c, 0x21, 0x95, 0x05, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x7c, 0x46,
	0xfc, 0x7a, 0x50, 0xb5, 0xde, 0xa9, 0x95, 0x20, 0xe1, 0x20, 0xb0, 0xa4, 0x90, 0x10, 0x17, 0x8b,
	0x4b, 0x62, 0x49, 0xa2, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x4f, 0x10, 0x98, 0xad, 0x94, 0xc9, 0xc5,
	0xeb, 0x9d, 0x5a, 0xe9, 0x92, 0x5a, 0x94, 0x59, 0x96, 0x58, 0x92, 0x99, 0x9f, 0x27, 0x64, 0xc0,
	0xc5, 0x96, 0x9b, 0x5a, 0x92, 0x91, 0x9f, 0x02, 0x35, 0x4b, 0x02, 0x66, 0x16, 0x42, 0x8d, 0x2f,
	0x58, 0x3e, 0x08, 0xaa, 0x4e, 0x48, 0x8d, 0x8b, 0x2f, 0x05, 0x2e, 0x17, 0x90, 0x58, 0x92, 0x01,
	0xb6, 0x80, 0x33, 0x08, 0x4d, 0x54, 0xcb, 0x92, 0x8b, 0x1d, 0xea, 0x1e, 0x21, 0x41, 0x2e, 0x5e,
	0xd7, 0x14, 0x23, 0x53, 0x53, 0x43, 0xcb, 0x80, 0xd2, 0xa4, 0x9c, 0xcc, 0x64, 0x01, 0x06, 0x21,
	0x21, 0x2e, 0x3e, 0x98, 0x10, 0x58, 0x57, 0xaa, 0x00, 0xa3, 0x10, 0x3b, 0x17, 0xb3, 0xa3, 0x6b,
	0xb0, 0x00, 0x93, 0x96, 0x1c, 0x97, 0x00, 0xba, 0xf5, 0x42, 0x5c, 0x5c, 0x6c, 0xc1, 0x39, 0x99,
	0x05, 0x46, 0x86, 0x02, 0x0c, 0x4e, 0x86, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8,
	0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7,
	0x10, 0x25, 0x8e, 0x23, 0x50, 0x93, 0xd8, 0xc0, 0x94, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x65,
	0xdf, 0xa5, 0x11, 0x76, 0x01, 0x00, 0x00,
}

func (m *Key) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Key) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Key) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintCrypto(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x12
	}
	if m.Type != 0 {
		i = encodeVarintCrypto(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *KeyDerivation) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *KeyDerivation) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *KeyDerivation) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.DerivationPath) > 0 {
		i -= len(m.DerivationPath)
		copy(dAtA[i:], m.DerivationPath)
		i = encodeVarintCrypto(dAtA, i, uint64(len(m.DerivationPath)))
		i--
		dAtA[i] = 0x12
	}
	if m.Method != 0 {
		i = encodeVarintCrypto(dAtA, i, uint64(m.Method))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintCrypto(dAtA []byte, offset int, v uint64) int {
	offset -= sovCrypto(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Key) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovCrypto(uint64(m.Type))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovCrypto(uint64(l))
	}
	return n
}

func (m *KeyDerivation) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Method != 0 {
		n += 1 + sovCrypto(uint64(m.Method))
	}
	l = len(m.DerivationPath)
	if l > 0 {
		n += 1 + l + sovCrypto(uint64(l))
	}
	return n
}

func sovCrypto(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCrypto(x uint64) (n int) {
	return sovCrypto(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Key) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCrypto
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
			return fmt.Errorf("proto: Key: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Key: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCrypto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= KeyType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCrypto
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
				return ErrInvalidLengthCrypto
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCrypto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCrypto(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCrypto
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
func (m *KeyDerivation) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCrypto
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
			return fmt.Errorf("proto: KeyDerivation: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: KeyDerivation: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Method", wireType)
			}
			m.Method = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCrypto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Method |= DerivationMethod(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DerivationPath", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCrypto
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
				return ErrInvalidLengthCrypto
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCrypto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DerivationPath = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCrypto(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCrypto
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
func skipCrypto(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCrypto
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
					return 0, ErrIntOverflowCrypto
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
					return 0, ErrIntOverflowCrypto
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
				return 0, ErrInvalidLengthCrypto
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCrypto
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCrypto
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCrypto        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCrypto          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCrypto = fmt.Errorf("proto: unexpected end of group")
)
