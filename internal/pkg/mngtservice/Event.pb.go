// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal/pkg/mngtservice/Event.proto

package mngtservice

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
	return fileDescriptor_70cfc8220f7e5de1, []int{0}
}

type Event struct {
	Id                   int32                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string               `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description          string               `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Type                 EventType            `protobuf:"varint,4,opt,name=type,proto3,enum=mngtservice.EventType" json:"type,omitempty"`
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
	return fileDescriptor_70cfc8220f7e5de1, []int{0}
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
	proto.RegisterEnum("mngtservice.EventType", EventType_name, EventType_value)
	proto.RegisterType((*Event)(nil), "mngtservice.Event")
}

func init() {
	proto.RegisterFile("internal/pkg/mngtservice/Event.proto", fileDescriptor_70cfc8220f7e5de1)
}

var fileDescriptor_70cfc8220f7e5de1 = []byte{
	// 277 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0x41, 0x6b, 0xb3, 0x40,
	0x10, 0x86, 0xbf, 0xdd, 0x68, 0xf8, 0x1c, 0x49, 0x90, 0xa5, 0x94, 0x25, 0x97, 0x4a, 0xe9, 0x41,
	0x3c, 0xac, 0x90, 0x9e, 0x7a, 0xb4, 0xc5, 0x43, 0xc0, 0xf4, 0x60, 0x6d, 0x0f, 0xbd, 0x25, 0x71,
	0x23, 0x4b, 0xd5, 0x5d, 0x74, 0x1a, 0xc8, 0x2f, 0xed, 0xdf, 0x29, 0x5d, 0x49, 0xf1, 0xd6, 0xdb,
	0xee, 0xbc, 0xcf, 0xf0, 0x3e, 0x0c, 0xdc, 0xa9, 0x0e, 0x65, 0xdf, 0xed, 0x9a, 0xc4, 0x7c, 0xd4,
	0x49, 0xdb, 0xd5, 0x38, 0xc8, 0xfe, 0xa4, 0x0e, 0x32, 0xc9, 0x4e, 0xb2, 0x43, 0x61, 0x7a, 0x8d,
	0x9a, 0xf9, 0x93, 0x60, 0x75, 0x53, 0x6b, 0x5d, 0x37, 0x32, 0xb1, 0xd1, 0xfe, 0xf3, 0x98, 0xa0,
	0x6a, 0xe5, 0x80, 0xbb, 0xd6, 0x8c, 0xf4, 0xed, 0x17, 0x01, 0xd7, 0x6e, 0xb3, 0x25, 0x50, 0x55,
	0x71, 0x12, 0x92, 0xc8, 0x2d, 0xa8, 0xaa, 0xd8, 0x15, 0xb8, 0xa8, 0xb0, 0x91, 0x9c, 0x86, 0x24,
	0xf2, 0x8a, 0xf1, 0xc3, 0x42, 0xf0, 0x2b, 0x39, 0x1c, 0x7a, 0x65, 0x50, 0xe9, 0x8e, 0xcf, 0x6c,
	0x36, 0x1d, 0xb1, 0x18, 0x1c, 0x3c, 0x1b, 0xc9, 0x9d, 0x90, 0x44, 0xcb, 0xf5, 0xb5, 0x98, 0xe8,
	0x08, 0xdb, 0x54, 0x9e, 0x8d, 0x2c, 0x2c, 0xc3, 0x04, 0x38, 0xc7, 0x5e, 0xb7, 0xdc, 0x0d, 0x49,
	0xe4, 0xaf, 0x57, 0x62, 0xb4, 0x15, 0x17, 0x5b, 0x51, 0x5e, 0x6c, 0x0b, 0xcb, 0xb1, 0x18, 0x28,
	0x6a, 0x3e, 0xff, 0x93, 0xa6, 0xa8, 0xe3, 0x07, 0xf0, 0x7e, 0xeb, 0xd8, 0x02, 0xbc, 0xd7, 0xe7,
	0xcd, 0x5b, 0x56, 0xbc, 0xa4, 0x79, 0xf0, 0x8f, 0xfd, 0x07, 0x67, 0x9b, 0x65, 0x65, 0x40, 0x7e,
	0x5e, 0x4f, 0x69, 0x9e, 0x07, 0x94, 0x79, 0xe0, 0x66, 0xdb, 0x74, 0x93, 0x07, 0xb3, 0xc7, 0xc5,
	0xfb, 0xf4, 0x88, 0xfb, 0xb9, 0x6d, 0xb8, 0xff, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x5d, 0xb0, 0x95,
	0x6b, 0x80, 0x01, 0x00, 0x00,
}