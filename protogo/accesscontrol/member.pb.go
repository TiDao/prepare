// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: accesscontrol/member.proto

package accesscontrol

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

// Serialized member of blockchain
type SerializedMember struct {
	// organization identifier of the member
	OrgId string `protobuf:"bytes,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	// member identity related info bytes
	MemberInfo []byte `protobuf:"bytes,2,opt,name=member_info,json=memberInfo,proto3" json:"member_info,omitempty"`
	// use cert compression
	// todo: is_full_cert -> compressed
	IsFullCert bool `protobuf:"varint,3,opt,name=is_full_cert,json=isFullCert,proto3" json:"is_full_cert,omitempty"`
}

func (m *SerializedMember) Reset()         { *m = SerializedMember{} }
func (m *SerializedMember) String() string { return proto.CompactTextString(m) }
func (*SerializedMember) ProtoMessage()    {}
func (*SerializedMember) Descriptor() ([]byte, []int) {
	return fileDescriptor_d279366e50bb6647, []int{0}
}
func (m *SerializedMember) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SerializedMember) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SerializedMember.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SerializedMember) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SerializedMember.Merge(m, src)
}
func (m *SerializedMember) XXX_Size() int {
	return m.Size()
}
func (m *SerializedMember) XXX_DiscardUnknown() {
	xxx_messageInfo_SerializedMember.DiscardUnknown(m)
}

var xxx_messageInfo_SerializedMember proto.InternalMessageInfo

func (m *SerializedMember) GetOrgId() string {
	if m != nil {
		return m.OrgId
	}
	return ""
}

func (m *SerializedMember) GetMemberInfo() []byte {
	if m != nil {
		return m.MemberInfo
	}
	return nil
}

func (m *SerializedMember) GetIsFullCert() bool {
	if m != nil {
		return m.IsFullCert
	}
	return false
}

func init() {
	proto.RegisterType((*SerializedMember)(nil), "accesscontrol.SerializedMember")
}

func init() { proto.RegisterFile("accesscontrol/member.proto", fileDescriptor_d279366e50bb6647) }

var fileDescriptor_d279366e50bb6647 = []byte{
	// 221 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4a, 0x4c, 0x4e, 0x4e,
	0x2d, 0x2e, 0x4e, 0xce, 0xcf, 0x2b, 0x29, 0xca, 0xcf, 0xd1, 0xcf, 0x4d, 0xcd, 0x4d, 0x4a, 0x2d,
	0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x45, 0x91, 0x53, 0xca, 0xe1, 0x12, 0x08, 0x4e,
	0x2d, 0xca, 0x4c, 0xcc, 0xc9, 0xac, 0x4a, 0x4d, 0xf1, 0x05, 0x2b, 0x14, 0x12, 0xe5, 0x62, 0xcb,
	0x2f, 0x4a, 0x8f, 0xcf, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x62, 0xcd, 0x2f, 0x4a,
	0xf7, 0x4c, 0x11, 0x92, 0xe7, 0xe2, 0x86, 0x98, 0x14, 0x9f, 0x99, 0x97, 0x96, 0x2f, 0xc1, 0xa4,
	0xc0, 0xa8, 0xc1, 0x13, 0xc4, 0x05, 0x11, 0xf2, 0xcc, 0x4b, 0xcb, 0x17, 0x52, 0xe0, 0xe2, 0xc9,
	0x2c, 0x8e, 0x4f, 0x2b, 0xcd, 0xc9, 0x89, 0x4f, 0x4e, 0x2d, 0x2a, 0x91, 0x60, 0x56, 0x60, 0xd4,
	0xe0, 0x08, 0xe2, 0xca, 0x2c, 0x76, 0x2b, 0xcd, 0xc9, 0x71, 0x4e, 0x2d, 0x2a, 0x71, 0xf2, 0x3f,
	0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63,
	0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86, 0x28, 0xd3, 0xe4, 0x8c, 0xc4, 0xcc, 0xbc,
	0xdc, 0xc4, 0xec, 0xd4, 0x22, 0xbd, 0xfc, 0xa2, 0x74, 0x7d, 0x04, 0x57, 0x37, 0x3d, 0x5f, 0xbf,
	0x20, 0x49, 0x1f, 0xec, 0xfa, 0xf4, 0x7c, 0x7d, 0x14, 0xe7, 0x27, 0xb1, 0x81, 0x85, 0x8d, 0x01,
	0x01, 0x00, 0x00, 0xff, 0xff, 0xc4, 0x3e, 0xc2, 0x89, 0xf2, 0x00, 0x00, 0x00,
}

func (m *SerializedMember) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SerializedMember) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SerializedMember) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.IsFullCert {
		i--
		if m.IsFullCert {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if len(m.MemberInfo) > 0 {
		i -= len(m.MemberInfo)
		copy(dAtA[i:], m.MemberInfo)
		i = encodeVarintMember(dAtA, i, uint64(len(m.MemberInfo)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.OrgId) > 0 {
		i -= len(m.OrgId)
		copy(dAtA[i:], m.OrgId)
		i = encodeVarintMember(dAtA, i, uint64(len(m.OrgId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMember(dAtA []byte, offset int, v uint64) int {
	offset -= sovMember(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SerializedMember) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.OrgId)
	if l > 0 {
		n += 1 + l + sovMember(uint64(l))
	}
	l = len(m.MemberInfo)
	if l > 0 {
		n += 1 + l + sovMember(uint64(l))
	}
	if m.IsFullCert {
		n += 2
	}
	return n
}

func sovMember(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMember(x uint64) (n int) {
	return sovMember(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SerializedMember) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMember
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
			return fmt.Errorf("proto: SerializedMember: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SerializedMember: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrgId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMember
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
				return ErrInvalidLengthMember
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMember
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OrgId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MemberInfo", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMember
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
				return ErrInvalidLengthMember
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMember
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MemberInfo = append(m.MemberInfo[:0], dAtA[iNdEx:postIndex]...)
			if m.MemberInfo == nil {
				m.MemberInfo = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsFullCert", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMember
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsFullCert = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipMember(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMember
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
func skipMember(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMember
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
					return 0, ErrIntOverflowMember
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
					return 0, ErrIntOverflowMember
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
				return 0, ErrInvalidLengthMember
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMember
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMember
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMember        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMember          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMember = fmt.Errorf("proto: unexpected end of group")
)
