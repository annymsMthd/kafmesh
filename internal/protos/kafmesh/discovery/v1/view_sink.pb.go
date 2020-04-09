// Code generated by protoc-gen-go. DO NOT EDIT.
// source: kafmesh/discovery/v1/view_sink.proto

package discoveryv1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// View sink is a kafmesh view sink.
type ViewSink struct {
	Name                 string           `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Topic                *TopicDefinition `protobuf:"bytes,2,opt,name=topic,proto3" json:"topic,omitempty"`
	Description          string           `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ViewSink) Reset()         { *m = ViewSink{} }
func (m *ViewSink) String() string { return proto.CompactTextString(m) }
func (*ViewSink) ProtoMessage()    {}
func (*ViewSink) Descriptor() ([]byte, []int) {
	return fileDescriptor_632b722f78d6963c, []int{0}
}

func (m *ViewSink) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ViewSink.Unmarshal(m, b)
}
func (m *ViewSink) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ViewSink.Marshal(b, m, deterministic)
}
func (m *ViewSink) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ViewSink.Merge(m, src)
}
func (m *ViewSink) XXX_Size() int {
	return xxx_messageInfo_ViewSink.Size(m)
}
func (m *ViewSink) XXX_DiscardUnknown() {
	xxx_messageInfo_ViewSink.DiscardUnknown(m)
}

var xxx_messageInfo_ViewSink proto.InternalMessageInfo

func (m *ViewSink) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ViewSink) GetTopic() *TopicDefinition {
	if m != nil {
		return m.Topic
	}
	return nil
}

func (m *ViewSink) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterType((*ViewSink)(nil), "kafmesh.discovery.v1.ViewSink")
}

func init() {
	proto.RegisterFile("kafmesh/discovery/v1/view_sink.proto", fileDescriptor_632b722f78d6963c)
}

var fileDescriptor_632b722f78d6963c = []byte{
	// 217 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xc9, 0x4e, 0x4c, 0xcb,
	0x4d, 0x2d, 0xce, 0xd0, 0x4f, 0xc9, 0x2c, 0x4e, 0xce, 0x2f, 0x4b, 0x2d, 0xaa, 0xd4, 0x2f, 0x33,
	0xd4, 0x2f, 0xcb, 0x4c, 0x2d, 0x8f, 0x2f, 0xce, 0xcc, 0xcb, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x12, 0x81, 0xaa, 0xd2, 0x83, 0xab, 0xd2, 0x2b, 0x33, 0x94, 0xd2, 0xc6, 0xaa, 0xb7, 0x24,
	0xbf, 0x20, 0x33, 0x39, 0x3e, 0x25, 0x35, 0x2d, 0x33, 0x2f, 0xb3, 0x24, 0x33, 0x3f, 0x0f, 0x62,
	0x84, 0x52, 0x2d, 0x17, 0x47, 0x58, 0x66, 0x6a, 0x79, 0x70, 0x66, 0x5e, 0xb6, 0x90, 0x10, 0x17,
	0x4b, 0x5e, 0x62, 0x6e, 0xaa, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x98, 0x2d, 0x64, 0xcd,
	0xc5, 0x0a, 0xd6, 0x29, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x6d, 0xa4, 0xaa, 0x87, 0xcd, 0x4a, 0xbd,
	0x10, 0x90, 0x12, 0x17, 0xb8, 0xd9, 0x41, 0x10, 0x3d, 0x42, 0x0a, 0x5c, 0xdc, 0x29, 0xa9, 0xc5,
	0xc9, 0x45, 0x99, 0x05, 0x20, 0x51, 0x09, 0x66, 0xb0, 0xb9, 0xc8, 0x42, 0x4e, 0xa1, 0x5c, 0x12,
	0xc9, 0xf9, 0xb9, 0x58, 0x0d, 0x75, 0xe2, 0x85, 0x39, 0x2c, 0x00, 0xe4, 0xd2, 0x00, 0xc6, 0x28,
	0x6e, 0xb8, 0x74, 0x99, 0xe1, 0x22, 0x26, 0x66, 0x6f, 0x97, 0x88, 0x55, 0x4c, 0x22, 0xde, 0x50,
	0xad, 0x2e, 0x70, 0xad, 0x61, 0x86, 0x49, 0x6c, 0x60, 0xcf, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff,
	0xff, 0x7b, 0xea, 0x52, 0x26, 0x47, 0x01, 0x00, 0x00,
}