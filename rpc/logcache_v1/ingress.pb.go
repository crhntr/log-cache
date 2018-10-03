// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ingress.proto

package logcache_v1

import (
	fmt "fmt"
	math "math"

	loggregator_v2 "code.cloudfoundry.org/go-loggregator/rpc/loggregator_v2"
	proto "github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SendRequest struct {
	Envelopes            *loggregator_v2.EnvelopeBatch `protobuf:"bytes,1,opt,name=envelopes,proto3" json:"envelopes,omitempty"`
	LocalOnly            bool                          `protobuf:"varint,2,opt,name=local_only,json=localOnly,proto3" json:"local_only,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *SendRequest) Reset()         { *m = SendRequest{} }
func (m *SendRequest) String() string { return proto.CompactTextString(m) }
func (*SendRequest) ProtoMessage()    {}
func (*SendRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_62b704fb0e4a291c, []int{0}
}

func (m *SendRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendRequest.Unmarshal(m, b)
}
func (m *SendRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendRequest.Marshal(b, m, deterministic)
}
func (m *SendRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendRequest.Merge(m, src)
}
func (m *SendRequest) XXX_Size() int {
	return xxx_messageInfo_SendRequest.Size(m)
}
func (m *SendRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SendRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SendRequest proto.InternalMessageInfo

func (m *SendRequest) GetEnvelopes() *loggregator_v2.EnvelopeBatch {
	if m != nil {
		return m.Envelopes
	}
	return nil
}

func (m *SendRequest) GetLocalOnly() bool {
	if m != nil {
		return m.LocalOnly
	}
	return false
}

type SendResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendResponse) Reset()         { *m = SendResponse{} }
func (m *SendResponse) String() string { return proto.CompactTextString(m) }
func (*SendResponse) ProtoMessage()    {}
func (*SendResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_62b704fb0e4a291c, []int{1}
}

func (m *SendResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendResponse.Unmarshal(m, b)
}
func (m *SendResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendResponse.Marshal(b, m, deterministic)
}
func (m *SendResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendResponse.Merge(m, src)
}
func (m *SendResponse) XXX_Size() int {
	return xxx_messageInfo_SendResponse.Size(m)
}
func (m *SendResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SendResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SendResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*SendRequest)(nil), "logcache.v1.SendRequest")
	proto.RegisterType((*SendResponse)(nil), "logcache.v1.SendResponse")
}

func init() { proto.RegisterFile("ingress.proto", fileDescriptor_62b704fb0e4a291c) }

var fileDescriptor_62b704fb0e4a291c = []byte{
	// 199 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcd, 0xcc, 0x4b, 0x2f,
	0x4a, 0x2d, 0x2e, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xce, 0xc9, 0x4f, 0x4f, 0x4e,
	0x4c, 0xce, 0x48, 0xd5, 0x2b, 0x33, 0x94, 0x12, 0x2c, 0x33, 0xd2, 0x4f, 0xcd, 0x2b, 0x4b, 0xcd,
	0xc9, 0x2f, 0x48, 0x85, 0xc8, 0x2b, 0x65, 0x72, 0x71, 0x07, 0xa7, 0xe6, 0xa5, 0x04, 0xa5, 0x16,
	0x96, 0xa6, 0x16, 0x97, 0x08, 0x59, 0x73, 0x71, 0xc2, 0x14, 0x14, 0x4b, 0x30, 0x2a, 0x30, 0x6a,
	0x70, 0x1b, 0xc9, 0xea, 0xe5, 0xe4, 0xa7, 0xa7, 0x17, 0xa5, 0xa6, 0x27, 0x96, 0xe4, 0x17, 0xe9,
	0x95, 0x19, 0xe9, 0xb9, 0x42, 0x15, 0x38, 0x25, 0x96, 0x24, 0x67, 0x04, 0x21, 0xd4, 0x0b, 0xc9,
	0x72, 0x71, 0xe5, 0xe4, 0x27, 0x27, 0xe6, 0xc4, 0xe7, 0xe7, 0xe5, 0x54, 0x4a, 0x30, 0x29, 0x30,
	0x6a, 0x70, 0x04, 0x71, 0x82, 0x45, 0xfc, 0xf3, 0x72, 0x2a, 0x95, 0xf8, 0xb8, 0x78, 0x20, 0x56,
	0x15, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x1a, 0x79, 0x70, 0xb1, 0x7b, 0x42, 0xdc, 0x2a, 0x64, 0xcb,
	0xc5, 0x02, 0x92, 0x12, 0x92, 0xd0, 0x43, 0x72, 0xae, 0x1e, 0x92, 0xc3, 0xa4, 0x24, 0xb1, 0xc8,
	0x40, 0xcc, 0x51, 0x62, 0x48, 0x62, 0x03, 0xfb, 0xc5, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xc3,
	0xc0, 0x8c, 0xe4, 0xfc, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// IngressClient is the client API for Ingress service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type IngressClient interface {
	// Send is used to emit Envelopes batches into LogCache. The RPC function
	// will not return until the data has been stored.
	Send(ctx context.Context, in *SendRequest, opts ...grpc.CallOption) (*SendResponse, error)
}

type ingressClient struct {
	cc *grpc.ClientConn
}

func NewIngressClient(cc *grpc.ClientConn) IngressClient {
	return &ingressClient{cc}
}

func (c *ingressClient) Send(ctx context.Context, in *SendRequest, opts ...grpc.CallOption) (*SendResponse, error) {
	out := new(SendResponse)
	err := c.cc.Invoke(ctx, "/logcache.v1.Ingress/Send", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IngressServer is the server API for Ingress service.
type IngressServer interface {
	// Send is used to emit Envelopes batches into LogCache. The RPC function
	// will not return until the data has been stored.
	Send(context.Context, *SendRequest) (*SendResponse, error)
}

func RegisterIngressServer(s *grpc.Server, srv IngressServer) {
	s.RegisterService(&_Ingress_serviceDesc, srv)
}

func _Ingress_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IngressServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logcache.v1.Ingress/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IngressServer).Send(ctx, req.(*SendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Ingress_serviceDesc = grpc.ServiceDesc{
	ServiceName: "logcache.v1.Ingress",
	HandlerType: (*IngressServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Send",
			Handler:    _Ingress_Send_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ingress.proto",
}
