// Code generated by protoc-gen-gogo.
// source: kv.proto
// DO NOT EDIT!

/*
	Package mvccpb is a generated protocol buffer package.

	It is generated from these files:
		kv.proto

	It has these top-level messages:
		KeyValue
		Event
*/
package mvccpb

import (
	"fmt"

	proto "github.com/gogo/protobuf/proto"

	math "math"
)

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.GoGoProtoPackageIsVersion1

type Event_EventType int32

const (
	PUT    Event_EventType = 0
	DELETE Event_EventType = 1
)

var Event_EventType_name = map[int32]string{
	0: "PUT",
	1: "DELETE",
}
var Event_EventType_value = map[string]int32{
	"PUT":    0,
	"DELETE": 1,
}

func (x Event_EventType) String() string {
	return proto.EnumName(Event_EventType_name, int32(x))
}
func (Event_EventType) EnumDescriptor() ([]byte, []int) { return fileDescriptorKv, []int{1, 0} }

type KeyValue struct {
	// key is the key in bytes. An empty key is not allowed.
	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// create_revision is the revision of last creation on this key.
	CreateRevision int64 `protobuf:"varint,2,opt,name=create_revision,json=createRevision,proto3" json:"create_revision,omitempty"`
	// mod_revision is the revision of last modification on this key.
	ModRevision int64 `protobuf:"varint,3,opt,name=mod_revision,json=modRevision,proto3" json:"mod_revision,omitempty"`
	// version is the version of the key. A deletion resets
	// the version to zero and any modification of the key
	// increases its version.
	Version int64 `protobuf:"varint,4,opt,name=version,proto3" json:"version,omitempty"`
	// value is the value held by the key, in bytes.
	Value []byte `protobuf:"bytes,5,opt,name=value,proto3" json:"value,omitempty"`
	// lease is the ID of the lease that attached to key.
	// When the attached lease expires, the key will be deleted.
	// If lease is 0, then no lease is attached to the key.
	Lease int64 `protobuf:"varint,6,opt,name=lease,proto3" json:"lease,omitempty"`
}

func (m *KeyValue) Reset()                    { *m = KeyValue{} }
func (m *KeyValue) String() string            { return proto.CompactTextString(m) }
func (*KeyValue) ProtoMessage()               {}
func (*KeyValue) Descriptor() ([]byte, []int) { return fileDescriptorKv, []int{0} }

type Event struct {
	// type is the kind of event. If type is a PUT, it indicates
	// new data has been stored to the key. If type is a DELETE,
	// it indicates the key was deleted.
	Type Event_EventType `protobuf:"varint,1,opt,name=type,proto3,enum=mvccpb.Event_EventType" json:"type,omitempty"`
	// kv holds the KeyValue for the event.
	// A PUT event contains current kv pair.
	// A PUT event with kv.Version=1 indicates the creation of a key.
	// A DELETE/EXPIRE event contains the deleted key with
	// its modification revision set to the revision of deletion.
	Kv *KeyValue `protobuf:"bytes,2,opt,name=kv" json:"kv,omitempty"`
	// prev_kv holds the key-value pair before the event happens.
	PrevKv *KeyValue `protobuf:"bytes,3,opt,name=prev_kv,json=prevKv" json:"prev_kv,omitempty"`
}

func (m *Event) Reset()                    { *m = Event{} }
func (m *Event) String() string            { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()               {}
func (*Event) Descriptor() ([]byte, []int) { return fileDescriptorKv, []int{1} }

func init() {
	proto.RegisterType((*KeyValue)(nil), "mvccpb.KeyValue")
	proto.RegisterType((*Event)(nil), "mvccpb.Event")
	proto.RegisterEnum("mvccpb.Event_EventType", Event_EventType_name, Event_EventType_value)
}
func (m *KeyValue) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *KeyValue) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Key) > 0 {
		data[i] = 0xa
		i++
		i = encodeVarintKv(data, i, uint64(len(m.Key)))
		i += copy(data[i:], m.Key)
	}
	if m.CreateRevision != 0 {
		data[i] = 0x10
		i++
		i = encodeVarintKv(data, i, uint64(m.CreateRevision))
	}
	if m.ModRevision != 0 {
		data[i] = 0x18
		i++
		i = encodeVarintKv(data, i, uint64(m.ModRevision))
	}
	if m.Version != 0 {
		data[i] = 0x20
		i++
		i = encodeVarintKv(data, i, uint64(m.Version))
	}
	if len(m.Value) > 0 {
		data[i] = 0x2a
		i++
		i = encodeVarintKv(data, i, uint64(len(m.Value)))
		i += copy(data[i:], m.Value)
	}
	if m.Lease != 0 {
		data[i] = 0x30
		i++
		i = encodeVarintKv(data, i, uint64(m.Lease))
	}
	return i, nil
}

func (m *Event) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Event) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Type != 0 {
		data[i] = 0x8
		i++
		i = encodeVarintKv(data, i, uint64(m.Type))
	}
	if m.Kv != nil {
		data[i] = 0x12
		i++
		i = encodeVarintKv(data, i, uint64(m.Kv.Size()))
		n1, err := m.Kv.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.PrevKv != nil {
		data[i] = 0x1a
		i++
		i = encodeVarintKv(data, i, uint64(m.PrevKv.Size()))
		n2, err := m.PrevKv.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func encodeFixed64Kv(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Kv(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintKv(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *KeyValue) Size() (n int) {
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovKv(uint64(l))
	}
	if m.CreateRevision != 0 {
		n += 1 + sovKv(uint64(m.CreateRevision))
	}
	if m.ModRevision != 0 {
		n += 1 + sovKv(uint64(m.ModRevision))
	}
	if m.Version != 0 {
		n += 1 + sovKv(uint64(m.Version))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovKv(uint64(l))
	}
	if m.Lease != 0 {
		n += 1 + sovKv(uint64(m.Lease))
	}
	return n
}

func (m *Event) Size() (n int) {
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovKv(uint64(m.Type))
	}
	if m.Kv != nil {
		l = m.Kv.Size()
		n += 1 + l + sovKv(uint64(l))
	}
	if m.PrevKv != nil {
		l = m.PrevKv.Size()
		n += 1 + l + sovKv(uint64(l))
	}
	return n
}

func sovKv(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozKv(x uint64) (n int) {
	return sovKv(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *KeyValue) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowKv
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: KeyValue: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: KeyValue: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKv
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthKv
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = append(m.Key[:0], data[iNdEx:postIndex]...)
			if m.Key == nil {
				m.Key = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreateRevision", wireType)
			}
			m.CreateRevision = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKv
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.CreateRevision |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ModRevision", wireType)
			}
			m.ModRevision = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKv
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.ModRevision |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKv
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Version |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKv
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthKv
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = append(m.Value[:0], data[iNdEx:postIndex]...)
			if m.Value == nil {
				m.Value = []byte{}
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Lease", wireType)
			}
			m.Lease = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKv
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Lease |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipKv(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthKv
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
func (m *Event) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowKv
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Event: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Event: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKv
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Type |= (Event_EventType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Kv", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKv
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthKv
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Kv == nil {
				m.Kv = &KeyValue{}
			}
			if err := m.Kv.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrevKv", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKv
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthKv
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PrevKv == nil {
				m.PrevKv = &KeyValue{}
			}
			if err := m.PrevKv.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipKv(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthKv
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
func skipKv(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowKv
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
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
					return 0, ErrIntOverflowKv
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowKv
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthKv
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowKv
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipKv(data[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthKv = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowKv   = fmt.Errorf("proto: integer overflow")
)

var fileDescriptorKv = []byte{
	// 292 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x90, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0xbb, 0x4d, 0x9b, 0xd4, 0x69, 0xa9, 0x61, 0x09, 0x18, 0x3c, 0x94, 0x98, 0x8b, 0x8a,
	0x10, 0xa1, 0xbe, 0x81, 0x98, 0x53, 0x3d, 0x48, 0x88, 0x5e, 0x4b, 0x1a, 0x87, 0x52, 0xd2, 0x76,
	0x43, 0x1a, 0x17, 0xf2, 0x3c, 0xde, 0x7d, 0x8e, 0x1e, 0xfb, 0x08, 0xea, 0x93, 0xb8, 0x3b, 0x6b,
	0xea, 0xc9, 0xc3, 0x2c, 0x33, 0xff, 0xff, 0xb1, 0xfb, 0xcf, 0xc2, 0xa0, 0x90, 0x51, 0x59, 0x89,
	0x5a, 0x70, 0x7b, 0x23, 0xf3, 0xbc, 0x5c, 0x9c, 0x7b, 0x4b, 0xb1, 0x14, 0x24, 0xdd, 0xea, 0xce,
	0xb8, 0xe1, 0x07, 0x83, 0xc1, 0x0c, 0x9b, 0x97, 0x6c, 0xfd, 0x86, 0xdc, 0x05, 0xab, 0xc0, 0xc6,
	0x67, 0x01, 0xbb, 0x1a, 0x25, 0xba, 0xe5, 0x97, 0x70, 0x9a, 0x57, 0x98, 0xd5, 0x38, 0xaf, 0x50,
	0xae, 0x76, 0x2b, 0xb1, 0xf5, 0xbb, 0xca, 0xb5, 0x92, 0xb1, 0x91, 0x93, 0x5f, 0x95, 0x5f, 0xc0,
	0x68, 0x23, 0x5e, 0xff, 0x28, 0x8b, 0xa8, 0xa1, 0xd2, 0x8e, 0x88, 0x0f, 0x8e, 0xc4, 0x8a, 0xdc,
	0x1e, 0xb9, 0xed, 0xc8, 0x3d, 0xe8, 0x4b, 0x1d, 0xc0, 0xef, 0xd3, 0xcb, 0x66, 0xd0, 0xea, 0x1a,
	0xb3, 0x1d, 0xfa, 0x36, 0xd1, 0x66, 0x08, 0xdf, 0x19, 0xf4, 0x63, 0x89, 0xdb, 0x9a, 0xdf, 0x40,
	0xaf, 0x6e, 0x4a, 0xa4, 0xb8, 0xe3, 0xe9, 0x59, 0x64, 0xf6, 0x8c, 0xc8, 0x34, 0x67, 0xaa, 0xec,
	0x84, 0x20, 0x1e, 0x40, 0xb7, 0x90, 0x94, 0x7d, 0x38, 0x75, 0x5b, 0xb4, 0x5d, 0x3c, 0x51, 0x1e,
	0xbf, 0x06, 0xa7, 0x54, 0xf1, 0xe7, 0x0a, 0xb3, 0xfe, 0xc1, 0x6c, 0x0d, 0xcc, 0x64, 0x18, 0xc0,
	0xc9, 0xf1, 0x7e, 0xee, 0x80, 0xf5, 0xf4, 0x9c, 0xba, 0x1d, 0x0e, 0x60, 0x3f, 0xc4, 0x8f, 0x71,
	0x1a, 0xbb, 0xec, 0xde, 0xdb, 0x7f, 0x4d, 0x3a, 0x07, 0x55, 0xfb, 0xef, 0x09, 0x3b, 0xa8, 0xfa,
	0x54, 0xb5, 0xb0, 0xe9, 0xcf, 0xef, 0x7e, 0x02, 0x00, 0x00, 0xff, 0xff, 0xbb, 0x86, 0x36, 0x07,
	0x9d, 0x01, 0x00, 0x00,
}
