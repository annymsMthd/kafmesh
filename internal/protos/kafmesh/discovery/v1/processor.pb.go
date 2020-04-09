// Code generated by protoc-gen-go. DO NOT EDIT.
// source: kafmesh/discovery/v1/processor.proto

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

// Processor is a stateful kafmesh processor.
type Processor struct {
	Name                 string       `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	GroupName            string       `protobuf:"bytes,2,opt,name=group_name,json=groupName,proto3" json:"group_name,omitempty"`
	Description          string       `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Inputs               []*Input     `protobuf:"bytes,4,rep,name=inputs,proto3" json:"inputs,omitempty"`
	Lookups              []*Lookup    `protobuf:"bytes,5,rep,name=lookups,proto3" json:"lookups,omitempty"`
	Joins                []*Join      `protobuf:"bytes,6,rep,name=joins,proto3" json:"joins,omitempty"`
	Outputs              []*Output    `protobuf:"bytes,7,rep,name=outputs,proto3" json:"outputs,omitempty"`
	Persistence          *Persistence `protobuf:"bytes,8,opt,name=persistence,proto3" json:"persistence,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Processor) Reset()         { *m = Processor{} }
func (m *Processor) String() string { return proto.CompactTextString(m) }
func (*Processor) ProtoMessage()    {}
func (*Processor) Descriptor() ([]byte, []int) {
	return fileDescriptor_698127296455087b, []int{0}
}

func (m *Processor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Processor.Unmarshal(m, b)
}
func (m *Processor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Processor.Marshal(b, m, deterministic)
}
func (m *Processor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Processor.Merge(m, src)
}
func (m *Processor) XXX_Size() int {
	return xxx_messageInfo_Processor.Size(m)
}
func (m *Processor) XXX_DiscardUnknown() {
	xxx_messageInfo_Processor.DiscardUnknown(m)
}

var xxx_messageInfo_Processor proto.InternalMessageInfo

func (m *Processor) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Processor) GetGroupName() string {
	if m != nil {
		return m.GroupName
	}
	return ""
}

func (m *Processor) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Processor) GetInputs() []*Input {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *Processor) GetLookups() []*Lookup {
	if m != nil {
		return m.Lookups
	}
	return nil
}

func (m *Processor) GetJoins() []*Join {
	if m != nil {
		return m.Joins
	}
	return nil
}

func (m *Processor) GetOutputs() []*Output {
	if m != nil {
		return m.Outputs
	}
	return nil
}

func (m *Processor) GetPersistence() *Persistence {
	if m != nil {
		return m.Persistence
	}
	return nil
}

// Input is the input to a processor.
type Input struct {
	Topic                *TopicDefinition `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Input) Reset()         { *m = Input{} }
func (m *Input) String() string { return proto.CompactTextString(m) }
func (*Input) ProtoMessage()    {}
func (*Input) Descriptor() ([]byte, []int) {
	return fileDescriptor_698127296455087b, []int{1}
}

func (m *Input) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Input.Unmarshal(m, b)
}
func (m *Input) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Input.Marshal(b, m, deterministic)
}
func (m *Input) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Input.Merge(m, src)
}
func (m *Input) XXX_Size() int {
	return xxx_messageInfo_Input.Size(m)
}
func (m *Input) XXX_DiscardUnknown() {
	xxx_messageInfo_Input.DiscardUnknown(m)
}

var xxx_messageInfo_Input proto.InternalMessageInfo

func (m *Input) GetTopic() *TopicDefinition {
	if m != nil {
		return m.Topic
	}
	return nil
}

// Join is a join for a processor.
type Join struct {
	Topic                *TopicDefinition `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Join) Reset()         { *m = Join{} }
func (m *Join) String() string { return proto.CompactTextString(m) }
func (*Join) ProtoMessage()    {}
func (*Join) Descriptor() ([]byte, []int) {
	return fileDescriptor_698127296455087b, []int{2}
}

func (m *Join) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Join.Unmarshal(m, b)
}
func (m *Join) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Join.Marshal(b, m, deterministic)
}
func (m *Join) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Join.Merge(m, src)
}
func (m *Join) XXX_Size() int {
	return xxx_messageInfo_Join.Size(m)
}
func (m *Join) XXX_DiscardUnknown() {
	xxx_messageInfo_Join.DiscardUnknown(m)
}

var xxx_messageInfo_Join proto.InternalMessageInfo

func (m *Join) GetTopic() *TopicDefinition {
	if m != nil {
		return m.Topic
	}
	return nil
}

// Lookup is a lookup for a processor.
type Lookup struct {
	Topic                *TopicDefinition `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Lookup) Reset()         { *m = Lookup{} }
func (m *Lookup) String() string { return proto.CompactTextString(m) }
func (*Lookup) ProtoMessage()    {}
func (*Lookup) Descriptor() ([]byte, []int) {
	return fileDescriptor_698127296455087b, []int{3}
}

func (m *Lookup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Lookup.Unmarshal(m, b)
}
func (m *Lookup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Lookup.Marshal(b, m, deterministic)
}
func (m *Lookup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Lookup.Merge(m, src)
}
func (m *Lookup) XXX_Size() int {
	return xxx_messageInfo_Lookup.Size(m)
}
func (m *Lookup) XXX_DiscardUnknown() {
	xxx_messageInfo_Lookup.DiscardUnknown(m)
}

var xxx_messageInfo_Lookup proto.InternalMessageInfo

func (m *Lookup) GetTopic() *TopicDefinition {
	if m != nil {
		return m.Topic
	}
	return nil
}

// Output is an output for a processor.
type Output struct {
	Topic                *TopicDefinition `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	Description          string           `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Output) Reset()         { *m = Output{} }
func (m *Output) String() string { return proto.CompactTextString(m) }
func (*Output) ProtoMessage()    {}
func (*Output) Descriptor() ([]byte, []int) {
	return fileDescriptor_698127296455087b, []int{4}
}

func (m *Output) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Output.Unmarshal(m, b)
}
func (m *Output) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Output.Marshal(b, m, deterministic)
}
func (m *Output) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Output.Merge(m, src)
}
func (m *Output) XXX_Size() int {
	return xxx_messageInfo_Output.Size(m)
}
func (m *Output) XXX_DiscardUnknown() {
	xxx_messageInfo_Output.DiscardUnknown(m)
}

var xxx_messageInfo_Output proto.InternalMessageInfo

func (m *Output) GetTopic() *TopicDefinition {
	if m != nil {
		return m.Topic
	}
	return nil
}

func (m *Output) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

// Persistence is the stateful persistence for a processor.
type Persistence struct {
	Topic                *TopicDefinition `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Persistence) Reset()         { *m = Persistence{} }
func (m *Persistence) String() string { return proto.CompactTextString(m) }
func (*Persistence) ProtoMessage()    {}
func (*Persistence) Descriptor() ([]byte, []int) {
	return fileDescriptor_698127296455087b, []int{5}
}

func (m *Persistence) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Persistence.Unmarshal(m, b)
}
func (m *Persistence) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Persistence.Marshal(b, m, deterministic)
}
func (m *Persistence) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Persistence.Merge(m, src)
}
func (m *Persistence) XXX_Size() int {
	return xxx_messageInfo_Persistence.Size(m)
}
func (m *Persistence) XXX_DiscardUnknown() {
	xxx_messageInfo_Persistence.DiscardUnknown(m)
}

var xxx_messageInfo_Persistence proto.InternalMessageInfo

func (m *Persistence) GetTopic() *TopicDefinition {
	if m != nil {
		return m.Topic
	}
	return nil
}

func init() {
	proto.RegisterType((*Processor)(nil), "kafmesh.discovery.v1.Processor")
	proto.RegisterType((*Input)(nil), "kafmesh.discovery.v1.Input")
	proto.RegisterType((*Join)(nil), "kafmesh.discovery.v1.Join")
	proto.RegisterType((*Lookup)(nil), "kafmesh.discovery.v1.Lookup")
	proto.RegisterType((*Output)(nil), "kafmesh.discovery.v1.Output")
	proto.RegisterType((*Persistence)(nil), "kafmesh.discovery.v1.Persistence")
}

func init() {
	proto.RegisterFile("kafmesh/discovery/v1/processor.proto", fileDescriptor_698127296455087b)
}

var fileDescriptor_698127296455087b = []byte{
	// 386 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0x41, 0x4b, 0xe3, 0x40,
	0x14, 0xc7, 0x49, 0xda, 0xa4, 0xdb, 0x17, 0xd8, 0xc3, 0xd0, 0x43, 0xe8, 0xee, 0x42, 0x36, 0x28,
	0x14, 0x84, 0xd4, 0xb4, 0xe0, 0xc5, 0x5b, 0x1b, 0x0f, 0xb6, 0xa2, 0x21, 0x48, 0x11, 0x2f, 0xa5,
	0xa6, 0xd3, 0x3a, 0xd6, 0x66, 0x86, 0x4c, 0x12, 0xf0, 0xeb, 0x78, 0xf4, 0x33, 0x7a, 0x90, 0xbc,
	0xb4, 0xb1, 0x68, 0xf0, 0x92, 0x5b, 0xc8, 0xfb, 0xfd, 0xfe, 0x33, 0xf3, 0x66, 0x1e, 0x1c, 0x6d,
	0x16, 0xab, 0x2d, 0x95, 0x8f, 0xfd, 0x25, 0x93, 0x21, 0xcf, 0x68, 0xfc, 0xd2, 0xcf, 0xdc, 0xbe,
	0x88, 0x79, 0x48, 0xa5, 0xe4, 0xb1, 0x23, 0x62, 0x9e, 0x70, 0xd2, 0xd9, 0x51, 0x4e, 0x49, 0x39,
	0x99, 0xdb, 0x3d, 0xa9, 0x74, 0x13, 0x2e, 0x58, 0x38, 0x5f, 0xd2, 0x15, 0x8b, 0x58, 0xc2, 0x78,
	0x54, 0x44, 0xd8, 0xef, 0x2a, 0xb4, 0xfd, 0x7d, 0x2c, 0x21, 0xd0, 0x8c, 0x16, 0x5b, 0x6a, 0x2a,
	0x96, 0xd2, 0x6b, 0x07, 0xf8, 0x4d, 0xfe, 0x01, 0xac, 0x63, 0x9e, 0x8a, 0x39, 0x56, 0x54, 0xac,
	0xb4, 0xf1, 0xcf, 0x75, 0x5e, 0xb6, 0xc0, 0x58, 0x52, 0x19, 0xc6, 0x4c, 0xe4, 0xa9, 0x66, 0x03,
	0xeb, 0x87, 0xbf, 0xc8, 0x10, 0x74, 0x16, 0x89, 0x34, 0x91, 0x66, 0xd3, 0x6a, 0xf4, 0x8c, 0xc1,
	0x1f, 0xa7, 0x6a, 0xdb, 0xce, 0x65, 0xce, 0x04, 0x3b, 0x94, 0x9c, 0x41, 0xeb, 0x99, 0xf3, 0x4d,
	0x2a, 0xa4, 0xa9, 0xa1, 0xf5, 0xb7, 0xda, 0xba, 0x42, 0x28, 0xd8, 0xc3, 0xe4, 0x14, 0xb4, 0x27,
	0xce, 0x22, 0x69, 0xea, 0x68, 0x75, 0xab, 0xad, 0x09, 0x67, 0x51, 0x50, 0x80, 0xf9, 0x4a, 0x3c,
	0x4d, 0x70, 0x7f, 0xad, 0x9f, 0x56, 0xba, 0x41, 0x28, 0xd8, 0xc3, 0x64, 0x0c, 0x86, 0xa0, 0xb1,
	0x64, 0x32, 0xa1, 0x51, 0x48, 0xcd, 0x5f, 0x96, 0xd2, 0x33, 0x06, 0xff, 0xab, 0x5d, 0xff, 0x13,
	0x0c, 0x0e, 0x2d, 0xdb, 0x03, 0x0d, 0xcf, 0x4d, 0xce, 0x41, 0xc3, 0x1b, 0xc2, 0xd6, 0x1b, 0x83,
	0xe3, 0xea, 0x9c, 0xdb, 0x1c, 0xf1, 0xca, 0x3b, 0x0c, 0x0a, 0xc7, 0x1e, 0x43, 0x33, 0x3f, 0x51,
	0xbd, 0x90, 0x0b, 0xd0, 0x8b, 0x66, 0xd6, 0x8b, 0x59, 0x83, 0x5e, 0x74, 0xaa, 0x56, 0xcc, 0xd7,
	0x67, 0xa5, 0x7e, 0x7b, 0x56, 0xf6, 0x04, 0x8c, 0x83, 0xb6, 0xd6, 0x5a, 0x6d, 0x34, 0x03, 0x33,
	0xe4, 0xdb, 0x4a, 0x65, 0xf4, 0xbb, 0x1c, 0x0f, 0x3f, 0x9f, 0x18, 0x5f, 0xb9, 0x37, 0xca, 0x7a,
	0xe6, 0xbe, 0xaa, 0x8d, 0xa9, 0x77, 0xf7, 0xa6, 0x76, 0xa6, 0x3b, 0xd7, 0x2b, 0xdd, 0x99, 0xfb,
	0xa0, 0xe3, 0x90, 0x0d, 0x3f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x7d, 0x7d, 0x24, 0xa0, 0xcf, 0x03,
	0x00, 0x00,
}