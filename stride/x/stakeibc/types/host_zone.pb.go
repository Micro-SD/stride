// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stakeibc/host_zone.proto

package types

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

type HostZone struct {
	PortId    string `protobuf:"bytes,1,opt,name=portId,proto3" json:"portId,omitempty"`
	ChannelId string `protobuf:"bytes,2,opt,name=channelId,proto3" json:"channelId,omitempty"`
}

func (m *HostZone) Reset()         { *m = HostZone{} }
func (m *HostZone) String() string { return proto.CompactTextString(m) }
func (*HostZone) ProtoMessage()    {}
func (*HostZone) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1d300c62c2b2d54, []int{0}
}
func (m *HostZone) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HostZone) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HostZone.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HostZone) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HostZone.Merge(m, src)
}
func (m *HostZone) XXX_Size() int {
	return m.Size()
}
func (m *HostZone) XXX_DiscardUnknown() {
	xxx_messageInfo_HostZone.DiscardUnknown(m)
}

var xxx_messageInfo_HostZone proto.InternalMessageInfo

func (m *HostZone) GetPortId() string {
	if m != nil {
		return m.PortId
	}
	return ""
}

func (m *HostZone) GetChannelId() string {
	if m != nil {
		return m.ChannelId
	}
	return ""
}

func init() {
	proto.RegisterType((*HostZone)(nil), "Stridelabs.stride.stakeibc.HostZone")
}

func init() { proto.RegisterFile("stakeibc/host_zone.proto", fileDescriptor_a1d300c62c2b2d54) }

var fileDescriptor_a1d300c62c2b2d54 = []byte{
	// 183 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x28, 0x2e, 0x49, 0xcc,
	0x4e, 0xcd, 0x4c, 0x4a, 0xd6, 0xcf, 0xc8, 0x2f, 0x2e, 0x89, 0xaf, 0xca, 0xcf, 0x4b, 0xd5, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x0a, 0x2e, 0x29, 0xca, 0x4c, 0x49, 0xcd, 0x49, 0x4c, 0x2a,
	0xd6, 0x2b, 0x06, 0x33, 0xf5, 0x60, 0x6a, 0x95, 0x1c, 0xb8, 0x38, 0x3c, 0xf2, 0x8b, 0x4b, 0xa2,
	0xf2, 0xf3, 0x52, 0x85, 0xc4, 0xb8, 0xd8, 0x0a, 0xf2, 0x8b, 0x4a, 0x3c, 0x53, 0x24, 0x18, 0x15,
	0x18, 0x35, 0x38, 0x83, 0xa0, 0x3c, 0x21, 0x19, 0x2e, 0xce, 0xe4, 0x8c, 0xc4, 0xbc, 0xbc, 0xd4,
	0x1c, 0xcf, 0x14, 0x09, 0x26, 0xb0, 0x14, 0x42, 0xc0, 0xc9, 0xe3, 0xc4, 0x23, 0x39, 0xc6, 0x0b,
	0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86,
	0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0xf4, 0xd2, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73,
	0xf5, 0x21, 0x4e, 0xd0, 0x05, 0xb9, 0x41, 0x1f, 0xe2, 0x06, 0xfd, 0x0a, 0x7d, 0xb8, 0x8b, 0x4b,
	0x2a, 0x0b, 0x52, 0x8b, 0x93, 0xd8, 0xc0, 0xce, 0x35, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x01,
	0x08, 0x81, 0x6a, 0xca, 0x00, 0x00, 0x00,
}

func (m *HostZone) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HostZone) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HostZone) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ChannelId) > 0 {
		i -= len(m.ChannelId)
		copy(dAtA[i:], m.ChannelId)
		i = encodeVarintHostZone(dAtA, i, uint64(len(m.ChannelId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.PortId) > 0 {
		i -= len(m.PortId)
		copy(dAtA[i:], m.PortId)
		i = encodeVarintHostZone(dAtA, i, uint64(len(m.PortId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintHostZone(dAtA []byte, offset int, v uint64) int {
	offset -= sovHostZone(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *HostZone) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PortId)
	if l > 0 {
		n += 1 + l + sovHostZone(uint64(l))
	}
	l = len(m.ChannelId)
	if l > 0 {
		n += 1 + l + sovHostZone(uint64(l))
	}
	return n
}

func sovHostZone(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozHostZone(x uint64) (n int) {
	return sovHostZone(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *HostZone) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHostZone
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
			return fmt.Errorf("proto: HostZone: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HostZone: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PortId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHostZone
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
				return ErrInvalidLengthHostZone
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHostZone
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PortId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChannelId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHostZone
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
				return ErrInvalidLengthHostZone
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHostZone
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChannelId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHostZone(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthHostZone
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
func skipHostZone(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHostZone
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
					return 0, ErrIntOverflowHostZone
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
					return 0, ErrIntOverflowHostZone
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
				return 0, ErrInvalidLengthHostZone
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupHostZone
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthHostZone
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthHostZone        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHostZone          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupHostZone = fmt.Errorf("proto: unexpected end of group")
)
