// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lavanet/lava/plans/plans_proposal.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

type PlansAddProposal struct {
	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description"`
	Plans       []Plan `protobuf:"bytes,3,rep,name=plans,proto3" json:"plans"`
	Modify      bool   `protobuf:"varint,4,opt,name=modify,proto3" json:"modify"`
}

func (m *PlansAddProposal) Reset()      { *m = PlansAddProposal{} }
func (*PlansAddProposal) ProtoMessage() {}
func (*PlansAddProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_781f7efff31affc8, []int{0}
}
func (m *PlansAddProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PlansAddProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PlansAddProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PlansAddProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlansAddProposal.Merge(m, src)
}
func (m *PlansAddProposal) XXX_Size() int {
	return m.Size()
}
func (m *PlansAddProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_PlansAddProposal.DiscardUnknown(m)
}

var xxx_messageInfo_PlansAddProposal proto.InternalMessageInfo

type PlansDelProposal struct {
	Title       string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title"`
	Description string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description"`
	Plans       []string `protobuf:"bytes,3,rep,name=plans,proto3" json:"plans"`
}

func (m *PlansDelProposal) Reset()      { *m = PlansDelProposal{} }
func (*PlansDelProposal) ProtoMessage() {}
func (*PlansDelProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_781f7efff31affc8, []int{1}
}
func (m *PlansDelProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PlansDelProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PlansDelProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PlansDelProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlansDelProposal.Merge(m, src)
}
func (m *PlansDelProposal) XXX_Size() int {
	return m.Size()
}
func (m *PlansDelProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_PlansDelProposal.DiscardUnknown(m)
}

var xxx_messageInfo_PlansDelProposal proto.InternalMessageInfo

func init() {
	proto.RegisterType((*PlansAddProposal)(nil), "lavanet.lava.plans.PlansAddProposal")
	proto.RegisterType((*PlansDelProposal)(nil), "lavanet.lava.plans.PlansDelProposal")
}

func init() {
	proto.RegisterFile("lavanet/lava/plans/plans_proposal.proto", fileDescriptor_781f7efff31affc8)
}

var fileDescriptor_781f7efff31affc8 = []byte{
	// 308 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xcf, 0x49, 0x2c, 0x4b,
	0xcc, 0x4b, 0x2d, 0xd1, 0x07, 0xd1, 0xfa, 0x05, 0x39, 0x89, 0x79, 0xc5, 0x10, 0x32, 0xbe, 0xa0,
	0x28, 0xbf, 0x20, 0xbf, 0x38, 0x31, 0x47, 0xaf, 0xa0, 0x28, 0xbf, 0x24, 0x5f, 0x48, 0x08, 0xaa,
	0x50, 0x0f, 0x44, 0xeb, 0x81, 0x95, 0x48, 0x89, 0xa4, 0xe7, 0xa7, 0xe7, 0x83, 0xa5, 0xf5, 0x41,
	0x2c, 0x88, 0x4a, 0x29, 0x59, 0x1c, 0x46, 0x42, 0xa4, 0x95, 0x2e, 0x33, 0x72, 0x09, 0x04, 0x80,
	0x04, 0x1d, 0x53, 0x52, 0x02, 0xa0, 0x76, 0x08, 0xc9, 0x73, 0xb1, 0x96, 0x64, 0x96, 0xe4, 0xa4,
	0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x3a, 0x71, 0xbe, 0xba, 0x27, 0x0f, 0x11, 0x08, 0x82, 0x50,
	0x42, 0x86, 0x5c, 0xdc, 0x29, 0xa9, 0xc5, 0xc9, 0x45, 0x99, 0x05, 0x25, 0x99, 0xf9, 0x79, 0x12,
	0x4c, 0x60, 0x65, 0xfc, 0xaf, 0xee, 0xc9, 0x23, 0x0b, 0x07, 0x21, 0x73, 0x84, 0x6c, 0xb9, 0x58,
	0xc1, 0x96, 0x4b, 0x30, 0x2b, 0x30, 0x6b, 0x70, 0x1b, 0x49, 0xe8, 0x61, 0xfa, 0x40, 0x0f, 0xe4,
	0x10, 0x27, 0xde, 0x13, 0xf7, 0xe4, 0x19, 0x40, 0x36, 0x82, 0xc5, 0x82, 0x20, 0x94, 0x90, 0x12,
	0x17, 0x5b, 0x6e, 0x7e, 0x4a, 0x66, 0x5a, 0xa5, 0x04, 0x8b, 0x02, 0xa3, 0x06, 0x87, 0x13, 0xd7,
	0xab, 0x7b, 0xf2, 0x50, 0x91, 0x20, 0x28, 0x6d, 0xc5, 0xd1, 0xb1, 0x40, 0x9e, 0x61, 0xc6, 0x02,
	0x79, 0x06, 0xa5, 0x89, 0x30, 0x5f, 0xb9, 0xa4, 0xe6, 0xd0, 0xd4, 0x57, 0xf2, 0xc8, 0xbe, 0x82,
	0x9a, 0x89, 0xec, 0x6e, 0x84, 0x9b, 0x9c, 0xdc, 0x56, 0x3c, 0x92, 0x63, 0x3c, 0xf1, 0x48, 0x8e,
	0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58,
	0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86, 0x28, 0x8d, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4,
	0xfc, 0x5c, 0x7d, 0x94, 0x18, 0x2b, 0x33, 0xd1, 0xaf, 0x80, 0x46, 0x5b, 0x49, 0x65, 0x41, 0x6a,
	0x71, 0x12, 0x1b, 0x38, 0xe2, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xc9, 0xf7, 0x78, 0xaa,
	0x2c, 0x02, 0x00, 0x00,
}

func (this *PlansAddProposal) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*PlansAddProposal)
	if !ok {
		that2, ok := that.(PlansAddProposal)
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
	if this.Title != that1.Title {
		return false
	}
	if this.Description != that1.Description {
		return false
	}
	if len(this.Plans) != len(that1.Plans) {
		return false
	}
	for i := range this.Plans {
		if !this.Plans[i].Equal(&that1.Plans[i]) {
			return false
		}
	}
	if this.Modify != that1.Modify {
		return false
	}
	return true
}
func (this *PlansDelProposal) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*PlansDelProposal)
	if !ok {
		that2, ok := that.(PlansDelProposal)
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
	if this.Title != that1.Title {
		return false
	}
	if this.Description != that1.Description {
		return false
	}
	if len(this.Plans) != len(that1.Plans) {
		return false
	}
	for i := range this.Plans {
		if this.Plans[i] != that1.Plans[i] {
			return false
		}
	}
	return true
}
func (m *PlansAddProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PlansAddProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PlansAddProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Modify {
		i--
		if m.Modify {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if len(m.Plans) > 0 {
		for iNdEx := len(m.Plans) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Plans[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintPlansProposal(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintPlansProposal(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintPlansProposal(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PlansDelProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PlansDelProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PlansDelProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Plans) > 0 {
		for iNdEx := len(m.Plans) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Plans[iNdEx])
			copy(dAtA[i:], m.Plans[iNdEx])
			i = encodeVarintPlansProposal(dAtA, i, uint64(len(m.Plans[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintPlansProposal(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintPlansProposal(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPlansProposal(dAtA []byte, offset int, v uint64) int {
	offset -= sovPlansProposal(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PlansAddProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovPlansProposal(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovPlansProposal(uint64(l))
	}
	if len(m.Plans) > 0 {
		for _, e := range m.Plans {
			l = e.Size()
			n += 1 + l + sovPlansProposal(uint64(l))
		}
	}
	if m.Modify {
		n += 2
	}
	return n
}

func (m *PlansDelProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovPlansProposal(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovPlansProposal(uint64(l))
	}
	if len(m.Plans) > 0 {
		for _, s := range m.Plans {
			l = len(s)
			n += 1 + l + sovPlansProposal(uint64(l))
		}
	}
	return n
}

func sovPlansProposal(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPlansProposal(x uint64) (n int) {
	return sovPlansProposal(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PlansAddProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPlansProposal
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
			return fmt.Errorf("proto: PlansAddProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PlansAddProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlansProposal
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
				return ErrInvalidLengthPlansProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPlansProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlansProposal
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
				return ErrInvalidLengthPlansProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPlansProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Plans", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlansProposal
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
				return ErrInvalidLengthPlansProposal
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPlansProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Plans = append(m.Plans, Plan{})
			if err := m.Plans[len(m.Plans)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Modify", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlansProposal
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
			m.Modify = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipPlansProposal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPlansProposal
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
func (m *PlansDelProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPlansProposal
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
			return fmt.Errorf("proto: PlansDelProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PlansDelProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlansProposal
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
				return ErrInvalidLengthPlansProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPlansProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlansProposal
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
				return ErrInvalidLengthPlansProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPlansProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Plans", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlansProposal
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
				return ErrInvalidLengthPlansProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPlansProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Plans = append(m.Plans, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPlansProposal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPlansProposal
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
func skipPlansProposal(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPlansProposal
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
					return 0, ErrIntOverflowPlansProposal
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
					return 0, ErrIntOverflowPlansProposal
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
				return 0, ErrInvalidLengthPlansProposal
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPlansProposal
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPlansProposal
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPlansProposal        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPlansProposal          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPlansProposal = fmt.Errorf("proto: unexpected end of group")
)
