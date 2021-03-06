// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Event.proto

package golang

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// An event emitted
type Event struct {
	Id       string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Type     string `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	Data     []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	ParentId string `protobuf:"bytes,4,opt,name=parentId" json:"parentId,omitempty"`
}

func (m *Event) Reset()                    { *m = Event{} }
func (m *Event) String() string            { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()               {}
func (*Event) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Event) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Event) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Event) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Event) GetParentId() string {
	if m != nil {
		return m.ParentId
	}
	return ""
}

func init() {
	proto.RegisterType((*Event)(nil), "Event")
}

func init() { proto.RegisterFile("Event.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 120 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x76, 0x2d, 0x4b, 0xcd,
	0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x8a, 0xe6, 0x62, 0x05, 0x73, 0x85, 0xf8, 0xb8,
	0x98, 0x32, 0x53, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x98, 0x32, 0x53, 0x84, 0x84, 0xb8,
	0x58, 0x4a, 0x2a, 0x0b, 0x52, 0x25, 0x98, 0xc0, 0x22, 0x60, 0x36, 0x48, 0x2c, 0x25, 0xb1, 0x24,
	0x51, 0x82, 0x59, 0x81, 0x51, 0x83, 0x27, 0x08, 0xcc, 0x16, 0x92, 0xe2, 0xe2, 0x28, 0x48, 0x2c,
	0x4a, 0xcd, 0x2b, 0xf1, 0x4c, 0x91, 0x60, 0x01, 0xab, 0x85, 0xf3, 0x9d, 0x38, 0xa2, 0xd8, 0xd2,
	0xf3, 0x73, 0x12, 0xf3, 0xd2, 0x93, 0xd8, 0xc0, 0xb6, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff,
	0xb8, 0x3e, 0x03, 0x19, 0x7c, 0x00, 0x00, 0x00,
}
