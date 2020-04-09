// Code generated by protoc-gen-go. DO NOT EDIT.
// source: kafmesh/discovery/v1/discovery_api.proto

package discoveryv1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type GetServiceInfoRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetServiceInfoRequest) Reset()         { *m = GetServiceInfoRequest{} }
func (m *GetServiceInfoRequest) String() string { return proto.CompactTextString(m) }
func (*GetServiceInfoRequest) ProtoMessage()    {}
func (*GetServiceInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_87f32b5f68aaa02c, []int{0}
}

func (m *GetServiceInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetServiceInfoRequest.Unmarshal(m, b)
}
func (m *GetServiceInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetServiceInfoRequest.Marshal(b, m, deterministic)
}
func (m *GetServiceInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetServiceInfoRequest.Merge(m, src)
}
func (m *GetServiceInfoRequest) XXX_Size() int {
	return xxx_messageInfo_GetServiceInfoRequest.Size(m)
}
func (m *GetServiceInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetServiceInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetServiceInfoRequest proto.InternalMessageInfo

type GetServiceInfoResponse struct {
	Service              *Service `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetServiceInfoResponse) Reset()         { *m = GetServiceInfoResponse{} }
func (m *GetServiceInfoResponse) String() string { return proto.CompactTextString(m) }
func (*GetServiceInfoResponse) ProtoMessage()    {}
func (*GetServiceInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_87f32b5f68aaa02c, []int{1}
}

func (m *GetServiceInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetServiceInfoResponse.Unmarshal(m, b)
}
func (m *GetServiceInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetServiceInfoResponse.Marshal(b, m, deterministic)
}
func (m *GetServiceInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetServiceInfoResponse.Merge(m, src)
}
func (m *GetServiceInfoResponse) XXX_Size() int {
	return xxx_messageInfo_GetServiceInfoResponse.Size(m)
}
func (m *GetServiceInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetServiceInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetServiceInfoResponse proto.InternalMessageInfo

func (m *GetServiceInfoResponse) GetService() *Service {
	if m != nil {
		return m.Service
	}
	return nil
}

func init() {
	proto.RegisterType((*GetServiceInfoRequest)(nil), "kafmesh.discovery.v1.GetServiceInfoRequest")
	proto.RegisterType((*GetServiceInfoResponse)(nil), "kafmesh.discovery.v1.GetServiceInfoResponse")
}

func init() {
	proto.RegisterFile("kafmesh/discovery/v1/discovery_api.proto", fileDescriptor_87f32b5f68aaa02c)
}

var fileDescriptor_87f32b5f68aaa02c = []byte{
	// 217 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0xc8, 0x4e, 0x4c, 0xcb,
	0x4d, 0x2d, 0xce, 0xd0, 0x4f, 0xc9, 0x2c, 0x4e, 0xce, 0x2f, 0x4b, 0x2d, 0xaa, 0xd4, 0x2f, 0x33,
	0x44, 0x70, 0xe2, 0x13, 0x0b, 0x32, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x44, 0xa0, 0x2a,
	0xf5, 0xe0, 0x92, 0x7a, 0x65, 0x86, 0x52, 0x4a, 0x58, 0xf5, 0x17, 0xa7, 0x16, 0x95, 0x65, 0x26,
	0xa7, 0x42, 0x74, 0x2a, 0x89, 0x73, 0x89, 0xba, 0xa7, 0x96, 0x04, 0x43, 0xc4, 0x3c, 0xf3, 0xd2,
	0xf2, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x94, 0x02, 0xb9, 0xc4, 0xd0, 0x25, 0x8a, 0x0b,
	0xf2, 0xf3, 0x8a, 0x53, 0x85, 0xcc, 0xb9, 0xd8, 0xa1, 0x66, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70,
	0x1b, 0xc9, 0xea, 0x61, 0xb3, 0x5e, 0x0f, 0xaa, 0x37, 0x08, 0xa6, 0xda, 0xa8, 0x9a, 0x8b, 0xc7,
	0x05, 0xa6, 0xc0, 0x31, 0xc0, 0x53, 0x28, 0x9b, 0x8b, 0x0f, 0xd5, 0x0a, 0x21, 0x6d, 0xec, 0x26,
	0x61, 0x75, 0xa1, 0x94, 0x0e, 0x71, 0x8a, 0x21, 0xae, 0x76, 0x8a, 0xe4, 0x92, 0x48, 0xce, 0xcf,
	0xc5, 0xaa, 0xc5, 0x49, 0x10, 0xe1, 0xac, 0x82, 0xcc, 0x00, 0x50, 0xb8, 0x04, 0x30, 0x46, 0x71,
	0xc3, 0x95, 0x94, 0x19, 0x2e, 0x62, 0x62, 0xf6, 0x76, 0x89, 0x58, 0xc5, 0x24, 0xe2, 0x0d, 0xd5,
	0x0e, 0xd7, 0xa0, 0x17, 0x66, 0x98, 0xc4, 0x06, 0x0e, 0x4a, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x55, 0x19, 0xaf, 0x65, 0xb0, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DiscoveryAPIClient is the client API for DiscoveryAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DiscoveryAPIClient interface {
	// GetServiceInfo retreives information about the kafmesh service from the
	// node.
	GetServiceInfo(ctx context.Context, in *GetServiceInfoRequest, opts ...grpc.CallOption) (*GetServiceInfoResponse, error)
}

type discoveryAPIClient struct {
	cc *grpc.ClientConn
}

func NewDiscoveryAPIClient(cc *grpc.ClientConn) DiscoveryAPIClient {
	return &discoveryAPIClient{cc}
}

func (c *discoveryAPIClient) GetServiceInfo(ctx context.Context, in *GetServiceInfoRequest, opts ...grpc.CallOption) (*GetServiceInfoResponse, error) {
	out := new(GetServiceInfoResponse)
	err := c.cc.Invoke(ctx, "/kafmesh.discovery.v1.DiscoveryAPI/GetServiceInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DiscoveryAPIServer is the server API for DiscoveryAPI service.
type DiscoveryAPIServer interface {
	// GetServiceInfo retreives information about the kafmesh service from the
	// node.
	GetServiceInfo(context.Context, *GetServiceInfoRequest) (*GetServiceInfoResponse, error)
}

func RegisterDiscoveryAPIServer(s *grpc.Server, srv DiscoveryAPIServer) {
	s.RegisterService(&_DiscoveryAPI_serviceDesc, srv)
}

func _DiscoveryAPI_GetServiceInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetServiceInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscoveryAPIServer).GetServiceInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kafmesh.discovery.v1.DiscoveryAPI/GetServiceInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscoveryAPIServer).GetServiceInfo(ctx, req.(*GetServiceInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DiscoveryAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kafmesh.discovery.v1.DiscoveryAPI",
	HandlerType: (*DiscoveryAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetServiceInfo",
			Handler:    _DiscoveryAPI_GetServiceInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kafmesh/discovery/v1/discovery_api.proto",
}