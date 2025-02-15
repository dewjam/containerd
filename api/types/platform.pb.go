// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/containerd/containerd/api/types/platform.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
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

// Platform follows the structure of the OCI platform specification, from
// descriptors.
type Platform struct {
	OS                   string   `protobuf:"bytes,1,opt,name=os,proto3" json:"os,omitempty"`
	Architecture         string   `protobuf:"bytes,2,opt,name=architecture,proto3" json:"architecture,omitempty"`
	Variant              string   `protobuf:"bytes,3,opt,name=variant,proto3" json:"variant,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Platform) Reset()      { *m = Platform{} }
func (*Platform) ProtoMessage() {}
func (*Platform) Descriptor() ([]byte, []int) {
	return fileDescriptor_24ba7a4b83e2367e, []int{0}
}
func (m *Platform) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Platform) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Platform.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Platform) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Platform.Merge(m, src)
}
func (m *Platform) XXX_Size() int {
	return m.Size()
}
func (m *Platform) XXX_DiscardUnknown() {
	xxx_messageInfo_Platform.DiscardUnknown(m)
}

var xxx_messageInfo_Platform proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Platform)(nil), "containerd.types.Platform")
}

func init() {
	proto.RegisterFile("github.com/containerd/containerd/api/types/platform.proto", fileDescriptor_24ba7a4b83e2367e)
}

var fileDescriptor_24ba7a4b83e2367e = []byte{
	// 194 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x4c, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xce, 0xcf, 0x2b, 0x49, 0xcc, 0xcc, 0x4b, 0x2d,
	0x4a, 0x41, 0x66, 0x26, 0x16, 0x64, 0xea, 0x97, 0x54, 0x16, 0xa4, 0x16, 0xeb, 0x17, 0xe4, 0x24,
	0x96, 0xa4, 0xe5, 0x17, 0xe5, 0xea, 0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b, 0x09, 0x20, 0x14, 0xe9,
	0x81, 0x15, 0x48, 0x89, 0xa4, 0xe7, 0xa7, 0xe7, 0x83, 0x25, 0xf5, 0x41, 0x2c, 0x88, 0x3a, 0xa5,
	0x08, 0x2e, 0x8e, 0x00, 0xa8, 0x4e, 0x21, 0x3e, 0x2e, 0xa6, 0xfc, 0x62, 0x09, 0x46, 0x05, 0x46,
	0x0d, 0xce, 0x20, 0xa6, 0xfc, 0x62, 0x21, 0x25, 0x2e, 0x9e, 0xc4, 0xa2, 0xe4, 0x8c, 0xcc, 0x92,
	0xd4, 0xe4, 0x92, 0xd2, 0xa2, 0x54, 0x09, 0x26, 0xb0, 0x0c, 0x8a, 0x98, 0x90, 0x04, 0x17, 0x7b,
	0x59, 0x62, 0x51, 0x66, 0x62, 0x5e, 0x89, 0x04, 0x33, 0x58, 0x1a, 0xc6, 0x75, 0xf2, 0x3a, 0xf1,
	0x50, 0x8e, 0xe1, 0xc6, 0x43, 0x39, 0x86, 0x86, 0x47, 0x72, 0x8c, 0x27, 0x1e, 0xc9, 0x31, 0x5e,
	0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0x63, 0x94, 0x01, 0xf1, 0xde, 0xb2, 0x06, 0x93, 0x49,
	0x6c, 0x60, 0xc7, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x90, 0x53, 0xf1, 0x6b, 0x11, 0x01,
	0x00, 0x00,
}

func (m *Platform) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Platform) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Platform) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Variant) > 0 {
		i -= len(m.Variant)
		copy(dAtA[i:], m.Variant)
		i = encodeVarintPlatform(dAtA, i, uint64(len(m.Variant)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Architecture) > 0 {
		i -= len(m.Architecture)
		copy(dAtA[i:], m.Architecture)
		i = encodeVarintPlatform(dAtA, i, uint64(len(m.Architecture)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.OS) > 0 {
		i -= len(m.OS)
		copy(dAtA[i:], m.OS)
		i = encodeVarintPlatform(dAtA, i, uint64(len(m.OS)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPlatform(dAtA []byte, offset int, v uint64) int {
	offset -= sovPlatform(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Platform) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.OS)
	if l > 0 {
		n += 1 + l + sovPlatform(uint64(l))
	}
	l = len(m.Architecture)
	if l > 0 {
		n += 1 + l + sovPlatform(uint64(l))
	}
	l = len(m.Variant)
	if l > 0 {
		n += 1 + l + sovPlatform(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovPlatform(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPlatform(x uint64) (n int) {
	return sovPlatform(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Platform) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Platform{`,
		`Os:` + fmt.Sprintf("%v", this.OS) + `,`,
		`Architecture:` + fmt.Sprintf("%v", this.Architecture) + `,`,
		`Variant:` + fmt.Sprintf("%v", this.Variant) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringPlatform(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Platform) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPlatform
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
			return fmt.Errorf("proto: Platform: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Platform: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Os", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlatform
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
				return ErrInvalidLengthPlatform
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPlatform
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OS = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Architecture", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlatform
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
				return ErrInvalidLengthPlatform
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPlatform
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Architecture = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Variant", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlatform
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
				return ErrInvalidLengthPlatform
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPlatform
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Variant = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPlatform(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPlatform
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipPlatform(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPlatform
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
					return 0, ErrIntOverflowPlatform
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
					return 0, ErrIntOverflowPlatform
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
				return 0, ErrInvalidLengthPlatform
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPlatform
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPlatform
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPlatform        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPlatform          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPlatform = fmt.Errorf("proto: unexpected end of group")
)
