// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal/domain/entities/Event.proto

package entities

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type EventType int32

const (
	EventType_UNIVERSAL EventType = 0
	EventType_MEET      EventType = 1
	EventType_CALL      EventType = 2
	EventType_EMAIL     EventType = 3
)

var EventType_name = map[int32]string{
	0: "UNIVERSAL",
	1: "MEET",
	2: "CALL",
	3: "EMAIL",
}

var EventType_value = map[string]int32{
	"UNIVERSAL": 0,
	"MEET":      1,
	"CALL":      2,
	"EMAIL":     3,
}

func (x EventType) String() string {
	return proto.EnumName(EventType_name, int32(x))
}

func (EventType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2022c372d6ef8cab, []int{0}
}

type Event struct {
	Id                   int32                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string               `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description          string               `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Type                 EventType            `protobuf:"varint,4,opt,name=type,proto3,enum=entities.EventType" json:"type,omitempty"`
	From                 *timestamp.Timestamp `protobuf:"bytes,5,opt,name=from,proto3" json:"from,omitempty"`
	To                   *timestamp.Timestamp `protobuf:"bytes,6,opt,name=to,proto3" json:"to,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_2022c372d6ef8cab, []int{0}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Event) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Event) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Event) GetType() EventType {
	if m != nil {
		return m.Type
	}
	return EventType_UNIVERSAL
}

func (m *Event) GetFrom() *timestamp.Timestamp {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *Event) GetTo() *timestamp.Timestamp {
	if m != nil {
		return m.To
	}
	return nil
}

func init() {
	proto.RegisterEnum("entities.EventType", EventType_name, EventType_value)
	proto.RegisterType((*Event)(nil), "entities.Event")
}

func init() {
	proto.RegisterFile("internal/domain/entities/Event.proto", fileDescriptor_2022c372d6ef8cab)
}

var fileDescriptor_2022c372d6ef8cab = []byte{
	// 300 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0xdd, 0x34, 0x29, 0xcd, 0x14, 0x4b, 0x58, 0x3d, 0x84, 0x5e, 0x0c, 0x22, 0x18, 0x7a,
	0xd8, 0x85, 0x8a, 0x88, 0xc7, 0x2a, 0x39, 0x14, 0x52, 0x0f, 0xb1, 0x7a, 0xf0, 0x96, 0x36, 0xdb,
	0x3a, 0x98, 0xec, 0x86, 0x64, 0x2a, 0xd4, 0xe7, 0xf4, 0x81, 0xc4, 0x2d, 0x11, 0x2f, 0xe2, 0x6d,
	0xf7, 0x9f, 0x6f, 0xf8, 0x3f, 0x18, 0xb8, 0x40, 0x4d, 0xaa, 0xd1, 0x79, 0x29, 0x0b, 0x53, 0xe5,
	0xa8, 0xa5, 0xd2, 0x84, 0x84, 0xaa, 0x95, 0xc9, 0xbb, 0xd2, 0x24, 0xea, 0xc6, 0x90, 0xe1, 0x83,
	0x2e, 0x1d, 0x9f, 0x6d, 0x8d, 0xd9, 0x96, 0x4a, 0xda, 0x7c, 0xb5, 0xdb, 0x48, 0xc2, 0x4a, 0xb5,
	0x94, 0x57, 0xf5, 0x01, 0x3d, 0xff, 0x64, 0xe0, 0xd9, 0x55, 0x3e, 0x02, 0x07, 0x8b, 0x90, 0x45,
	0x2c, 0xf6, 0x32, 0x07, 0x0b, 0x7e, 0x0a, 0x1e, 0x21, 0x95, 0x2a, 0x74, 0x22, 0x16, 0xfb, 0xd9,
	0xe1, 0xc3, 0x23, 0x18, 0x16, 0xaa, 0x5d, 0x37, 0x58, 0x13, 0x1a, 0x1d, 0xf6, 0xec, 0xec, 0x77,
	0xc4, 0x2f, 0xc1, 0xa5, 0x7d, 0xad, 0x42, 0x37, 0x62, 0xf1, 0x68, 0x7a, 0x22, 0x3a, 0x17, 0x61,
	0x6b, 0x96, 0xfb, 0x5a, 0x65, 0x16, 0xe0, 0x02, 0xdc, 0x4d, 0x63, 0xaa, 0xd0, 0x8b, 0x58, 0x3c,
	0x9c, 0x8e, 0xc5, 0x41, 0x55, 0x74, 0xaa, 0x62, 0xd9, 0xa9, 0x66, 0x96, 0xe3, 0x13, 0x70, 0xc8,
	0x84, 0xfd, 0x7f, 0x69, 0x87, 0xcc, 0xe4, 0x16, 0xfc, 0x9f, 0x3a, 0x7e, 0x0c, 0xfe, 0xd3, 0xc3,
	0xfc, 0x39, 0xc9, 0x1e, 0x67, 0x69, 0x70, 0xc4, 0x07, 0xe0, 0x2e, 0x92, 0x64, 0x19, 0xb0, 0xef,
	0xd7, 0xfd, 0x2c, 0x4d, 0x03, 0x87, 0xfb, 0xe0, 0x25, 0x8b, 0xd9, 0x3c, 0x0d, 0x7a, 0x77, 0x37,
	0x2f, 0xd7, 0x5b, 0xa4, 0xd7, 0xdd, 0x4a, 0xac, 0x4d, 0x25, 0xf3, 0xb6, 0xc5, 0x8f, 0x37, 0x44,
	0xb9, 0xce, 0x4b, 0xa5, 0x8b, 0xbc, 0x91, 0x7f, 0x5d, 0x60, 0xd5, 0xb7, 0x2e, 0x57, 0x5f, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x42, 0xf0, 0xf3, 0x29, 0xa4, 0x01, 0x00, 0x00,
}