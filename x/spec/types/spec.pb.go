// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: spec/spec.proto

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

type Spec struct {
	Index                string       `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	Name                 string       `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Apis                 []ServiceApi `protobuf:"bytes,3,rep,name=apis,proto3" json:"apis"`
	Enabled              bool         `protobuf:"varint,4,opt,name=enabled,proto3" json:"enabled,omitempty"`
	ReliabilityThreshold uint32       `protobuf:"varint,5,opt,name=reliabilityThreshold,proto3" json:"reliabilityThreshold,omitempty"`
	ComparesHashes       bool         `protobuf:"varint,6,opt,name=comparesHashes,proto3" json:"comparesHashes,omitempty"`
	FinalizationCriteria uint32       `protobuf:"varint,7,opt,name=finalizationCriteria,proto3" json:"finalizationCriteria,omitempty"`
	SavedBlocks          uint32       `protobuf:"varint,8,opt,name=saved_blocks,json=savedBlocks,proto3" json:"saved_blocks,omitempty"`
}

func (m *Spec) Reset()         { *m = Spec{} }
func (m *Spec) String() string { return proto.CompactTextString(m) }
func (*Spec) ProtoMessage()    {}
func (*Spec) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4cc771ffab81d0a, []int{0}
}
func (m *Spec) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Spec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Spec.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Spec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Spec.Merge(m, src)
}
func (m *Spec) XXX_Size() int {
	return m.Size()
}
func (m *Spec) XXX_DiscardUnknown() {
	xxx_messageInfo_Spec.DiscardUnknown(m)
}

var xxx_messageInfo_Spec proto.InternalMessageInfo

func (m *Spec) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *Spec) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Spec) GetApis() []ServiceApi {
	if m != nil {
		return m.Apis
	}
	return nil
}

func (m *Spec) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *Spec) GetReliabilityThreshold() uint32 {
	if m != nil {
		return m.ReliabilityThreshold
	}
	return 0
}

func (m *Spec) GetComparesHashes() bool {
	if m != nil {
		return m.ComparesHashes
	}
	return false
}

func (m *Spec) GetFinalizationCriteria() uint32 {
	if m != nil {
		return m.FinalizationCriteria
	}
	return 0
}

func (m *Spec) GetSavedBlocks() uint32 {
	if m != nil {
		return m.SavedBlocks
	}
	return 0
}

func init() {
	proto.RegisterType((*Spec)(nil), "lavanet.lava.spec.Spec")
}

func init() { proto.RegisterFile("spec/spec.proto", fileDescriptor_c4cc771ffab81d0a) }

var fileDescriptor_c4cc771ffab81d0a = []byte{
	// 341 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xbf, 0x4e, 0xfb, 0x30,
	0x10, 0xc7, 0xe3, 0x36, 0xfd, 0xf3, 0x73, 0x7f, 0x80, 0xb0, 0x2a, 0x64, 0x55, 0x22, 0x04, 0x84,
	0x50, 0xa6, 0x44, 0x2a, 0x03, 0x33, 0x61, 0x61, 0x4e, 0x99, 0x58, 0x2a, 0x27, 0x39, 0x1a, 0x8b,
	0x34, 0xb6, 0x62, 0x53, 0xb5, 0x3c, 0x05, 0x8f, 0xc1, 0x13, 0xf0, 0x0c, 0x1d, 0x3b, 0x32, 0x21,
	0x94, 0xbe, 0x08, 0x8a, 0x03, 0x0b, 0x74, 0xf1, 0xdd, 0x7d, 0xbe, 0x77, 0x5f, 0x5b, 0x3e, 0x7c,
	0xa0, 0x24, 0x24, 0x41, 0x7d, 0xf8, 0xb2, 0x14, 0x5a, 0x90, 0xc3, 0x9c, 0x2d, 0x58, 0x01, 0xda,
	0xaf, 0xa3, 0x5f, 0x0b, 0xa3, 0xe1, 0x4c, 0xcc, 0x84, 0x51, 0x83, 0x3a, 0x6b, 0x1a, 0x47, 0x47,
	0xcd, 0x24, 0x94, 0x0b, 0x9e, 0xc0, 0x94, 0x49, 0xde, 0xf0, 0xb3, 0xb7, 0x16, 0xb6, 0x27, 0x12,
	0x12, 0x32, 0xc4, 0x1d, 0x5e, 0xa4, 0xb0, 0xa4, 0xc8, 0x45, 0xde, 0xbf, 0xa8, 0x29, 0x08, 0xc1,
	0x76, 0xc1, 0xe6, 0x40, 0x5b, 0x06, 0x9a, 0x9c, 0x5c, 0x61, 0x9b, 0x49, 0xae, 0x68, 0xdb, 0x6d,
	0x7b, 0x83, 0xf1, 0xb1, 0xff, 0xe7, 0x09, 0xfe, 0xa4, 0xb9, 0xe6, 0x5a, 0xf2, 0xd0, 0x5e, 0x7f,
	0x9c, 0x58, 0x91, 0x19, 0x20, 0x14, 0xf7, 0xa0, 0x60, 0x71, 0x0e, 0x29, 0xb5, 0x5d, 0xe4, 0xf5,
	0xa3, 0x9f, 0x92, 0x8c, 0xf1, 0xb0, 0x84, 0x9c, 0xb3, 0x98, 0xe7, 0x5c, 0xaf, 0xee, 0xb2, 0x12,
	0x54, 0x26, 0xf2, 0x94, 0x76, 0x5c, 0xe4, 0xed, 0x45, 0x3b, 0x35, 0x72, 0x81, 0xf7, 0x13, 0x31,
	0x97, 0xac, 0x04, 0x75, 0xcb, 0x54, 0x06, 0x8a, 0x76, 0x8d, 0xe9, 0x2f, 0x5a, 0x7b, 0x3f, 0xf0,
	0x82, 0xe5, 0xfc, 0x99, 0x69, 0x2e, 0x8a, 0x9b, 0x92, 0x6b, 0x28, 0x39, 0xa3, 0xbd, 0xc6, 0x7b,
	0x97, 0x46, 0x4e, 0xf1, 0x7f, 0xc5, 0x16, 0x90, 0x4e, 0xe3, 0x5c, 0x24, 0x8f, 0x8a, 0xf6, 0x4d,
	0xef, 0xc0, 0xb0, 0xd0, 0xa0, 0x30, 0x7c, 0xad, 0x1c, 0xb4, 0xae, 0x1c, 0xb4, 0xa9, 0x1c, 0xf4,
	0x59, 0x39, 0xe8, 0x65, 0xeb, 0x58, 0x9b, 0xad, 0x63, 0xbd, 0x6f, 0x1d, 0xeb, 0xfe, 0x7c, 0xc6,
	0x75, 0xf6, 0x14, 0xfb, 0x89, 0x98, 0x07, 0xdf, 0xff, 0x63, 0x62, 0xb0, 0x34, 0xdb, 0x0b, 0xf4,
	0x4a, 0x82, 0x8a, 0xbb, 0x66, 0x07, 0x97, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa7, 0xc7, 0xdc,
	0xd4, 0xd7, 0x01, 0x00, 0x00,
}

func (this *Spec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Spec)
	if !ok {
		that2, ok := that.(Spec)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Index != that1.Index {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if len(this.Apis) != len(that1.Apis) {
		return false
	}
	for i := range this.Apis {
		if !this.Apis[i].Equal(&that1.Apis[i]) {
			return false
		}
	}
	if this.Enabled != that1.Enabled {
		return false
	}
	if this.ReliabilityThreshold != that1.ReliabilityThreshold {
		return false
	}
	if this.ComparesHashes != that1.ComparesHashes {
		return false
	}
	if this.FinalizationCriteria != that1.FinalizationCriteria {
		return false
	}
	if this.SavedBlocks != that1.SavedBlocks {
		return false
	}
	return true
}
func (m *Spec) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Spec) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Spec) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.SavedBlocks != 0 {
		i = encodeVarintSpec(dAtA, i, uint64(m.SavedBlocks))
		i--
		dAtA[i] = 0x40
	}
	if m.FinalizationCriteria != 0 {
		i = encodeVarintSpec(dAtA, i, uint64(m.FinalizationCriteria))
		i--
		dAtA[i] = 0x38
	}
	if m.ComparesHashes {
		i--
		if m.ComparesHashes {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	if m.ReliabilityThreshold != 0 {
		i = encodeVarintSpec(dAtA, i, uint64(m.ReliabilityThreshold))
		i--
		dAtA[i] = 0x28
	}
	if m.Enabled {
		i--
		if m.Enabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if len(m.Apis) > 0 {
		for iNdEx := len(m.Apis) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Apis[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintSpec(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintSpec(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintSpec(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintSpec(dAtA []byte, offset int, v uint64) int {
	offset -= sovSpec(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Spec) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovSpec(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovSpec(uint64(l))
	}
	if len(m.Apis) > 0 {
		for _, e := range m.Apis {
			l = e.Size()
			n += 1 + l + sovSpec(uint64(l))
		}
	}
	if m.Enabled {
		n += 2
	}
	if m.ReliabilityThreshold != 0 {
		n += 1 + sovSpec(uint64(m.ReliabilityThreshold))
	}
	if m.ComparesHashes {
		n += 2
	}
	if m.FinalizationCriteria != 0 {
		n += 1 + sovSpec(uint64(m.FinalizationCriteria))
	}
	if m.SavedBlocks != 0 {
		n += 1 + sovSpec(uint64(m.SavedBlocks))
	}
	return n
}

func sovSpec(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSpec(x uint64) (n int) {
	return sovSpec(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Spec) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSpec
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
			return fmt.Errorf("proto: Spec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Spec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpec
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
				return ErrInvalidLengthSpec
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSpec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpec
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
				return ErrInvalidLengthSpec
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSpec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Apis", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpec
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
				return ErrInvalidLengthSpec
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSpec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Apis = append(m.Apis, ServiceApi{})
			if err := m.Apis[len(m.Apis)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Enabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpec
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
			m.Enabled = bool(v != 0)
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReliabilityThreshold", wireType)
			}
			m.ReliabilityThreshold = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ReliabilityThreshold |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ComparesHashes", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpec
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
			m.ComparesHashes = bool(v != 0)
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FinalizationCriteria", wireType)
			}
			m.FinalizationCriteria = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FinalizationCriteria |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SavedBlocks", wireType)
			}
			m.SavedBlocks = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SavedBlocks |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSpec(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSpec
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
func skipSpec(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSpec
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
					return 0, ErrIntOverflowSpec
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
					return 0, ErrIntOverflowSpec
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
				return 0, ErrInvalidLengthSpec
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSpec
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSpec
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSpec        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSpec          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSpec = fmt.Errorf("proto: unexpected end of group")
)
