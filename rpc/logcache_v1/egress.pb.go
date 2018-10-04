// Code generated by protoc-gen-go. DO NOT EDIT.
// source: egress.proto

package logcache_v1

import (
	loggregator_v2 "code.cloudfoundry.org/go-loggregator/rpc/loggregator_v2"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type EnvelopeType int32

const (
	EnvelopeType_ANY     EnvelopeType = 0
	EnvelopeType_LOG     EnvelopeType = 1
	EnvelopeType_COUNTER EnvelopeType = 2
	EnvelopeType_GAUGE   EnvelopeType = 3
	EnvelopeType_TIMER   EnvelopeType = 4
	EnvelopeType_EVENT   EnvelopeType = 5
)

var EnvelopeType_name = map[int32]string{
	0: "ANY",
	1: "LOG",
	2: "COUNTER",
	3: "GAUGE",
	4: "TIMER",
	5: "EVENT",
}

var EnvelopeType_value = map[string]int32{
	"ANY":     0,
	"LOG":     1,
	"COUNTER": 2,
	"GAUGE":   3,
	"TIMER":   4,
	"EVENT":   5,
}

func (x EnvelopeType) String() string {
	return proto.EnumName(EnvelopeType_name, int32(x))
}

func (EnvelopeType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e640ebe2e4e4631c, []int{0}
}

type ReadRequest struct {
	SourceId             string         `protobuf:"bytes,1,opt,name=source_id,json=sourceId,proto3" json:"source_id,omitempty"`
	StartTime            int64          `protobuf:"varint,2,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime              int64          `protobuf:"varint,3,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	Limit                int64          `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
	EnvelopeTypes        []EnvelopeType `protobuf:"varint,5,rep,packed,name=envelope_types,json=envelopeTypes,proto3,enum=logcache.v1.EnvelopeType" json:"envelope_types,omitempty"`
	Descending           bool           `protobuf:"varint,6,opt,name=descending,proto3" json:"descending,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ReadRequest) Reset()         { *m = ReadRequest{} }
func (m *ReadRequest) String() string { return proto.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()    {}
func (*ReadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e640ebe2e4e4631c, []int{0}
}

func (m *ReadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadRequest.Unmarshal(m, b)
}
func (m *ReadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadRequest.Marshal(b, m, deterministic)
}
func (m *ReadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadRequest.Merge(m, src)
}
func (m *ReadRequest) XXX_Size() int {
	return xxx_messageInfo_ReadRequest.Size(m)
}
func (m *ReadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadRequest proto.InternalMessageInfo

func (m *ReadRequest) GetSourceId() string {
	if m != nil {
		return m.SourceId
	}
	return ""
}

func (m *ReadRequest) GetStartTime() int64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *ReadRequest) GetEndTime() int64 {
	if m != nil {
		return m.EndTime
	}
	return 0
}

func (m *ReadRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ReadRequest) GetEnvelopeTypes() []EnvelopeType {
	if m != nil {
		return m.EnvelopeTypes
	}
	return nil
}

func (m *ReadRequest) GetDescending() bool {
	if m != nil {
		return m.Descending
	}
	return false
}

type ReadResponse struct {
	Envelopes            *loggregator_v2.EnvelopeBatch `protobuf:"bytes,1,opt,name=envelopes,proto3" json:"envelopes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *ReadResponse) Reset()         { *m = ReadResponse{} }
func (m *ReadResponse) String() string { return proto.CompactTextString(m) }
func (*ReadResponse) ProtoMessage()    {}
func (*ReadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e640ebe2e4e4631c, []int{1}
}

func (m *ReadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadResponse.Unmarshal(m, b)
}
func (m *ReadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadResponse.Marshal(b, m, deterministic)
}
func (m *ReadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadResponse.Merge(m, src)
}
func (m *ReadResponse) XXX_Size() int {
	return xxx_messageInfo_ReadResponse.Size(m)
}
func (m *ReadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReadResponse proto.InternalMessageInfo

func (m *ReadResponse) GetEnvelopes() *loggregator_v2.EnvelopeBatch {
	if m != nil {
		return m.Envelopes
	}
	return nil
}

type MetaRequest struct {
	LocalOnly            bool     `protobuf:"varint,1,opt,name=local_only,json=localOnly,proto3" json:"local_only,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MetaRequest) Reset()         { *m = MetaRequest{} }
func (m *MetaRequest) String() string { return proto.CompactTextString(m) }
func (*MetaRequest) ProtoMessage()    {}
func (*MetaRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e640ebe2e4e4631c, []int{2}
}

func (m *MetaRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetaRequest.Unmarshal(m, b)
}
func (m *MetaRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetaRequest.Marshal(b, m, deterministic)
}
func (m *MetaRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetaRequest.Merge(m, src)
}
func (m *MetaRequest) XXX_Size() int {
	return xxx_messageInfo_MetaRequest.Size(m)
}
func (m *MetaRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MetaRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MetaRequest proto.InternalMessageInfo

func (m *MetaRequest) GetLocalOnly() bool {
	if m != nil {
		return m.LocalOnly
	}
	return false
}

type MetaResponse struct {
	Meta                 map[string]*MetaInfo `protobuf:"bytes,1,rep,name=meta,proto3" json:"meta,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *MetaResponse) Reset()         { *m = MetaResponse{} }
func (m *MetaResponse) String() string { return proto.CompactTextString(m) }
func (*MetaResponse) ProtoMessage()    {}
func (*MetaResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e640ebe2e4e4631c, []int{3}
}

func (m *MetaResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetaResponse.Unmarshal(m, b)
}
func (m *MetaResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetaResponse.Marshal(b, m, deterministic)
}
func (m *MetaResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetaResponse.Merge(m, src)
}
func (m *MetaResponse) XXX_Size() int {
	return xxx_messageInfo_MetaResponse.Size(m)
}
func (m *MetaResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MetaResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MetaResponse proto.InternalMessageInfo

func (m *MetaResponse) GetMeta() map[string]*MetaInfo {
	if m != nil {
		return m.Meta
	}
	return nil
}

type MetaInfo struct {
	Count                int64    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Expired              int64    `protobuf:"varint,2,opt,name=expired,proto3" json:"expired,omitempty"`
	OldestTimestamp      int64    `protobuf:"varint,3,opt,name=oldest_timestamp,json=oldestTimestamp,proto3" json:"oldest_timestamp,omitempty"`
	NewestTimestamp      int64    `protobuf:"varint,4,opt,name=newest_timestamp,json=newestTimestamp,proto3" json:"newest_timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MetaInfo) Reset()         { *m = MetaInfo{} }
func (m *MetaInfo) String() string { return proto.CompactTextString(m) }
func (*MetaInfo) ProtoMessage()    {}
func (*MetaInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_e640ebe2e4e4631c, []int{4}
}

func (m *MetaInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetaInfo.Unmarshal(m, b)
}
func (m *MetaInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetaInfo.Marshal(b, m, deterministic)
}
func (m *MetaInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetaInfo.Merge(m, src)
}
func (m *MetaInfo) XXX_Size() int {
	return xxx_messageInfo_MetaInfo.Size(m)
}
func (m *MetaInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_MetaInfo.DiscardUnknown(m)
}

var xxx_messageInfo_MetaInfo proto.InternalMessageInfo

func (m *MetaInfo) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *MetaInfo) GetExpired() int64 {
	if m != nil {
		return m.Expired
	}
	return 0
}

func (m *MetaInfo) GetOldestTimestamp() int64 {
	if m != nil {
		return m.OldestTimestamp
	}
	return 0
}

func (m *MetaInfo) GetNewestTimestamp() int64 {
	if m != nil {
		return m.NewestTimestamp
	}
	return 0
}

func init() {
	proto.RegisterEnum("logcache.v1.EnvelopeType", EnvelopeType_name, EnvelopeType_value)
	proto.RegisterType((*ReadRequest)(nil), "logcache.v1.ReadRequest")
	proto.RegisterType((*ReadResponse)(nil), "logcache.v1.ReadResponse")
	proto.RegisterType((*MetaRequest)(nil), "logcache.v1.MetaRequest")
	proto.RegisterType((*MetaResponse)(nil), "logcache.v1.MetaResponse")
	proto.RegisterMapType((map[string]*MetaInfo)(nil), "logcache.v1.MetaResponse.MetaEntry")
	proto.RegisterType((*MetaInfo)(nil), "logcache.v1.MetaInfo")
}

func init() { proto.RegisterFile("egress.proto", fileDescriptor_e640ebe2e4e4631c) }

var fileDescriptor_e640ebe2e4e4631c = []byte{
	// 591 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x53, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xc5, 0x71, 0xd2, 0x24, 0xe3, 0xd0, 0x9a, 0x55, 0x91, 0xdc, 0x94, 0xa0, 0xc8, 0xbd, 0x84,
	0x82, 0x12, 0xd5, 0x1c, 0x40, 0x20, 0x24, 0x0a, 0xb2, 0xaa, 0x0a, 0x9a, 0x0a, 0x2b, 0x45, 0xe2,
	0x14, 0x16, 0x7b, 0x70, 0x2d, 0x9c, 0x5d, 0xe3, 0xdd, 0x04, 0x22, 0xc4, 0x85, 0x2f, 0x40, 0xe2,
	0xc0, 0xd7, 0xf0, 0x15, 0x5c, 0x39, 0xf2, 0x21, 0xc8, 0xbb, 0x71, 0x70, 0x0b, 0xb7, 0x99, 0x37,
	0xcf, 0xf3, 0x3c, 0xfb, 0x66, 0xa0, 0x83, 0x71, 0x8e, 0x42, 0x0c, 0xb3, 0x9c, 0x4b, 0x4e, 0xac,
	0x94, 0xc7, 0x21, 0x0d, 0xcf, 0x71, 0xb8, 0x38, 0xe8, 0x5e, 0x5b, 0x78, 0x23, 0x64, 0x0b, 0x4c,
	0x79, 0x86, 0xba, 0xde, 0xbd, 0x11, 0x73, 0x1e, 0xa7, 0x38, 0xa2, 0x59, 0x32, 0xa2, 0x8c, 0x71,
	0x49, 0x65, 0xc2, 0xd9, 0xea, 0x6b, 0xf7, 0x97, 0x01, 0x56, 0x80, 0x34, 0x0a, 0xf0, 0xfd, 0x1c,
	0x85, 0x24, 0xbb, 0xd0, 0x16, 0x7c, 0x9e, 0x87, 0x38, 0x4d, 0x22, 0xc7, 0xe8, 0x1b, 0x83, 0x76,
	0xd0, 0xd2, 0xc0, 0x71, 0x44, 0x7a, 0x00, 0x42, 0xd2, 0x5c, 0x4e, 0x65, 0x32, 0x43, 0xa7, 0xd6,
	0x37, 0x06, 0x66, 0xd0, 0x56, 0xc8, 0x24, 0x99, 0x21, 0xd9, 0x81, 0x16, 0xb2, 0x48, 0x17, 0x4d,
	0x55, 0x6c, 0x22, 0x8b, 0x54, 0x69, 0x1b, 0x1a, 0x69, 0x32, 0x4b, 0xa4, 0x53, 0x57, 0xb8, 0x4e,
	0xc8, 0x63, 0xd8, 0x2c, 0x7f, 0x76, 0x2a, 0x97, 0x19, 0x0a, 0xa7, 0xd1, 0x37, 0x07, 0x9b, 0xde,
	0xce, 0xb0, 0x32, 0xd3, 0xd0, 0x5f, 0x51, 0x26, 0xcb, 0x0c, 0x83, 0xab, 0x58, 0xc9, 0x04, 0xb9,
	0x09, 0x10, 0xa1, 0x08, 0x91, 0x45, 0x09, 0x8b, 0x9d, 0x8d, 0xbe, 0x31, 0x68, 0x05, 0x15, 0xc4,
	0x7d, 0x06, 0x1d, 0x3d, 0x9d, 0xc8, 0x38, 0x13, 0x48, 0x1e, 0x42, 0xbb, 0x6c, 0x20, 0xd4, 0x78,
	0x96, 0xd7, 0x2b, 0xc4, 0xe2, 0x1c, 0x63, 0x2a, 0x79, 0x3e, 0x5c, 0x78, 0x6b, 0xbd, 0x27, 0x54,
	0x86, 0xe7, 0xc1, 0x5f, 0xbe, 0x7b, 0x07, 0xac, 0x13, 0x94, 0xb4, 0x7c, 0xaa, 0x1e, 0x40, 0xca,
	0x43, 0x9a, 0x4e, 0x39, 0x4b, 0x97, 0xaa, 0x59, 0x2b, 0x68, 0x2b, 0xe4, 0x94, 0xa5, 0x4b, 0xf7,
	0xbb, 0x01, 0x1d, 0x4d, 0x5f, 0x69, 0xdf, 0x83, 0xfa, 0x0c, 0x25, 0x75, 0x8c, 0xbe, 0x39, 0xb0,
	0xbc, 0xbd, 0x0b, 0x33, 0x56, 0x89, 0x2a, 0xf1, 0x99, 0xcc, 0x97, 0x81, 0xfa, 0xa0, 0x3b, 0x86,
	0xf6, 0x1a, 0x22, 0x36, 0x98, 0xef, 0x70, 0xb9, 0xb2, 0xa6, 0x08, 0xc9, 0x6d, 0x68, 0x2c, 0x68,
	0x3a, 0xd7, 0x86, 0x58, 0xde, 0xf5, 0x7f, 0x1a, 0x1f, 0xb3, 0xb7, 0x3c, 0xd0, 0x9c, 0x07, 0xb5,
	0xfb, 0x86, 0xfb, 0xd5, 0x80, 0x56, 0x89, 0x17, 0xce, 0x84, 0x7c, 0xce, 0xa4, 0xea, 0x68, 0x06,
	0x3a, 0x21, 0x0e, 0x34, 0xf1, 0x63, 0x96, 0xe4, 0x18, 0xad, 0x6c, 0x2e, 0x53, 0x72, 0x0b, 0x6c,
	0x9e, 0x46, 0x28, 0xf4, 0x12, 0x08, 0x49, 0x67, 0xd9, 0xca, 0xec, 0x2d, 0x8d, 0x4f, 0x4a, 0xb8,
	0xa0, 0x32, 0xfc, 0x70, 0x91, 0xaa, 0xfd, 0xdf, 0xd2, 0xf8, 0x9a, 0xba, 0x3f, 0x86, 0x4e, 0xd5,
	0x66, 0xd2, 0x04, 0xf3, 0x70, 0xfc, 0xca, 0xbe, 0x52, 0x04, 0xcf, 0x4f, 0x8f, 0x6c, 0x83, 0x58,
	0xd0, 0x7c, 0x7a, 0x7a, 0x36, 0x9e, 0xf8, 0x81, 0x5d, 0x23, 0x6d, 0x68, 0x1c, 0x1d, 0x9e, 0x1d,
	0xf9, 0xb6, 0x59, 0x84, 0x93, 0xe3, 0x13, 0x3f, 0xb0, 0xeb, 0x45, 0xe8, 0xbf, 0xf4, 0xc7, 0x13,
	0xbb, 0xe1, 0xfd, 0x30, 0x60, 0xc3, 0x57, 0x57, 0x42, 0x5e, 0x43, 0xbd, 0x58, 0x01, 0xe2, 0x5c,
	0x78, 0x97, 0xca, 0xce, 0x77, 0x77, 0xfe, 0x53, 0xd1, 0x56, 0xb8, 0x7b, 0x5f, 0x7e, 0xfe, 0xfe,
	0x56, 0xeb, 0x91, 0x5d, 0x75, 0x3e, 0x8b, 0x83, 0x51, 0x8e, 0x34, 0x1a, 0x7d, 0x5a, 0x9f, 0xc8,
	0xa3, 0xfd, 0xfd, 0xcf, 0xe4, 0x05, 0xd4, 0x8b, 0xe7, 0xbc, 0xa4, 0x50, 0x59, 0x95, 0x4b, 0x0a,
	0x55, 0xb3, 0xdd, 0x6d, 0xa5, 0xb0, 0x49, 0x3a, 0xa5, 0x42, 0x61, 0xf9, 0x9b, 0x0d, 0x75, 0x9d,
	0x77, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0xf4, 0xe7, 0x66, 0x79, 0xeb, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EgressClient is the client API for Egress service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EgressClient interface {
	Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error)
	Meta(ctx context.Context, in *MetaRequest, opts ...grpc.CallOption) (*MetaResponse, error)
}

type egressClient struct {
	cc *grpc.ClientConn
}

func NewEgressClient(cc *grpc.ClientConn) EgressClient {
	return &egressClient{cc}
}

func (c *egressClient) Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error) {
	out := new(ReadResponse)
	err := c.cc.Invoke(ctx, "/logcache.v1.Egress/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *egressClient) Meta(ctx context.Context, in *MetaRequest, opts ...grpc.CallOption) (*MetaResponse, error) {
	out := new(MetaResponse)
	err := c.cc.Invoke(ctx, "/logcache.v1.Egress/Meta", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EgressServer is the server API for Egress service.
type EgressServer interface {
	Read(context.Context, *ReadRequest) (*ReadResponse, error)
	Meta(context.Context, *MetaRequest) (*MetaResponse, error)
}

func RegisterEgressServer(s *grpc.Server, srv EgressServer) {
	s.RegisterService(&_Egress_serviceDesc, srv)
}

func _Egress_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EgressServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logcache.v1.Egress/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EgressServer).Read(ctx, req.(*ReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Egress_Meta_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MetaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EgressServer).Meta(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logcache.v1.Egress/Meta",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EgressServer).Meta(ctx, req.(*MetaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Egress_serviceDesc = grpc.ServiceDesc{
	ServiceName: "logcache.v1.Egress",
	HandlerType: (*EgressServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Read",
			Handler:    _Egress_Read_Handler,
		},
		{
			MethodName: "Meta",
			Handler:    _Egress_Meta_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "egress.proto",
}