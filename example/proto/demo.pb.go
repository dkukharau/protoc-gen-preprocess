// Code generated by protoc-gen-go. DO NOT EDIT.
// source: demo.proto

package proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/infobloxopen/protoc-gen-preprocess/options"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
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

type Demo struct {
	// Also it is possible to specify additional method on field level
	PreprocessedField string `protobuf:"bytes,1,opt,name=preprocessedField,proto3" json:"preprocessedField,omitempty"`
	// Preprocessor automatically checks if field is repeated and generates methods accordingly
	PreprocessedRepeatedField []string `protobuf:"bytes,2,rep,name=preprocessedRepeatedField,proto3" json:"preprocessedRepeatedField,omitempty"`
	// If a field does not fit preprocess method, it is just ignored
	Ignored              int32    `protobuf:"varint,3,opt,name=ignored,proto3" json:"ignored,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Demo) Reset()         { *m = Demo{} }
func (m *Demo) String() string { return proto.CompactTextString(m) }
func (*Demo) ProtoMessage()    {}
func (*Demo) Descriptor() ([]byte, []int) {
	return fileDescriptor_demo_89d550965d436e19, []int{0}
}
func (m *Demo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Demo.Unmarshal(m, b)
}
func (m *Demo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Demo.Marshal(b, m, deterministic)
}
func (dst *Demo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Demo.Merge(dst, src)
}
func (m *Demo) XXX_Size() int {
	return xxx_messageInfo_Demo.Size(m)
}
func (m *Demo) XXX_DiscardUnknown() {
	xxx_messageInfo_Demo.DiscardUnknown(m)
}

var xxx_messageInfo_Demo proto.InternalMessageInfo

func (m *Demo) GetPreprocessedField() string {
	if m != nil {
		return m.PreprocessedField
	}
	return ""
}

func (m *Demo) GetPreprocessedRepeatedField() []string {
	if m != nil {
		return m.PreprocessedRepeatedField
	}
	return nil
}

func (m *Demo) GetIgnored() int32 {
	if m != nil {
		return m.Ignored
	}
	return 0
}

// This message left as is to show that we can provide our own preprocessors
type Custom struct {
	DoItYourself         string   `protobuf:"bytes,1,opt,name=doItYourself,proto3" json:"doItYourself,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Custom) Reset()         { *m = Custom{} }
func (m *Custom) String() string { return proto.CompactTextString(m) }
func (*Custom) ProtoMessage()    {}
func (*Custom) Descriptor() ([]byte, []int) {
	return fileDescriptor_demo_89d550965d436e19, []int{1}
}
func (m *Custom) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Custom.Unmarshal(m, b)
}
func (m *Custom) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Custom.Marshal(b, m, deterministic)
}
func (dst *Custom) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Custom.Merge(dst, src)
}
func (m *Custom) XXX_Size() int {
	return xxx_messageInfo_Custom.Size(m)
}
func (m *Custom) XXX_DiscardUnknown() {
	xxx_messageInfo_Custom.DiscardUnknown(m)
}

var xxx_messageInfo_Custom proto.InternalMessageInfo

func (m *Custom) GetDoItYourself() string {
	if m != nil {
		return m.DoItYourself
	}
	return ""
}

func init() {
	proto.RegisterType((*Demo)(nil), "proto.Demo")
	proto.RegisterType((*Custom)(nil), "proto.Custom")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DemoServiceClient is the client API for DemoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DemoServiceClient interface {
	Echo(ctx context.Context, in *Demo, opts ...grpc.CallOption) (*Demo, error)
}

type demoServiceClient struct {
	cc *grpc.ClientConn
}

func NewDemoServiceClient(cc *grpc.ClientConn) DemoServiceClient {
	return &demoServiceClient{cc}
}

func (c *demoServiceClient) Echo(ctx context.Context, in *Demo, opts ...grpc.CallOption) (*Demo, error) {
	out := new(Demo)
	err := c.cc.Invoke(ctx, "/proto.DemoService/Echo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DemoServiceServer is the server API for DemoService service.
type DemoServiceServer interface {
	Echo(context.Context, *Demo) (*Demo, error)
}

func RegisterDemoServiceServer(s *grpc.Server, srv DemoServiceServer) {
	s.RegisterService(&_DemoService_serviceDesc, srv)
}

func _DemoService_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Demo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DemoServiceServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.DemoService/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DemoServiceServer).Echo(ctx, req.(*Demo))
	}
	return interceptor(ctx, in, info, handler)
}

var _DemoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.DemoService",
	HandlerType: (*DemoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _DemoService_Echo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "demo.proto",
}

func init() { proto.RegisterFile("demo.proto", fileDescriptor_demo_89d550965d436e19) }

var fileDescriptor_demo_89d550965d436e19 = []byte{
	// 296 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x8f, 0xc1, 0x4e, 0x32, 0x31,
	0x14, 0x85, 0x53, 0x86, 0xf2, 0xc3, 0xe5, 0x5f, 0x68, 0x57, 0x48, 0x5c, 0x90, 0x71, 0x43, 0x8c,
	0xd0, 0x04, 0x13, 0x17, 0xc4, 0x8d, 0xa8, 0x24, 0x6e, 0xc7, 0x95, 0xcb, 0xa1, 0x73, 0x19, 0x9a,
	0xcc, 0xf4, 0x36, 0x9d, 0x8e, 0x71, 0xed, 0x2b, 0xf8, 0x1a, 0xae, 0x59, 0xf0, 0x1a, 0xbe, 0x82,
	0x0f, 0x62, 0x98, 0x09, 0x82, 0x31, 0xae, 0x9a, 0x7e, 0xdf, 0x39, 0x4d, 0x0f, 0x40, 0x82, 0x39,
	0x8d, 0xad, 0x23, 0x4f, 0x82, 0x57, 0x47, 0x7f, 0x9e, 0x6a, 0xbf, 0x2a, 0x17, 0x63, 0x45, 0xb9,
	0xd4, 0x66, 0x49, 0x8b, 0x8c, 0x5e, 0xc8, 0xa2, 0x91, 0x95, 0x56, 0xa3, 0x14, 0xcd, 0xc8, 0x3a,
	0xb4, 0x8e, 0x14, 0x16, 0x85, 0x24, 0xeb, 0x35, 0x99, 0x42, 0xee, 0x51, 0xfd, 0x5c, 0xff, 0x34,
	0x25, 0x4a, 0x33, 0x94, 0xb1, 0xd5, 0x32, 0x36, 0x86, 0x7c, 0x5c, 0x05, 0x6b, 0x1b, 0xbe, 0x33,
	0x68, 0xde, 0x61, 0x4e, 0xe2, 0x0a, 0x8e, 0xf7, 0x55, 0x4c, 0xe6, 0x1a, 0xb3, 0xa4, 0xc7, 0x06,
	0x6c, 0xd8, 0x99, 0xb5, 0x37, 0x6b, 0xde, 0x84, 0x46, 0x3b, 0x88, 0x7e, 0x47, 0xc4, 0x35, 0x9c,
	0x1c, 0xc2, 0x08, 0x2d, 0xc6, 0x7e, 0xd7, 0x6f, 0x0c, 0x82, 0x61, 0x27, 0xfa, 0x3b, 0x20, 0xce,
	0xe0, 0x9f, 0x4e, 0x0d, 0x39, 0x4c, 0x7a, 0xc1, 0x80, 0x0d, 0xf9, 0xac, 0xb3, 0x59, 0x73, 0x0e,
	0x01, 0xb0, 0x20, 0xda, 0x99, 0xe9, 0x37, 0x63, 0xe1, 0x05, 0xb4, 0x6e, 0xcb, 0xc2, 0x53, 0x2e,
	0x42, 0xf8, 0x9f, 0xd0, 0x83, 0x7f, 0xa2, 0xd2, 0x15, 0x98, 0x2d, 0xeb, 0xaf, 0x46, 0x3f, 0xd8,
	0xe4, 0x06, 0xba, 0xdb, 0x6d, 0x8f, 0xe8, 0x9e, 0xb5, 0x42, 0x31, 0x81, 0xe6, 0xbd, 0x5a, 0x91,
	0xe8, 0xd6, 0xdb, 0xc7, 0x5b, 0xd7, 0x3f, 0xbc, 0x84, 0x47, 0xaf, 0x1f, 0x9f, 0x6f, 0x0d, 0x08,
	0xb9, 0x44, 0xb5, 0xa2, 0x29, 0x3b, 0x5f, 0xb4, 0x2a, 0x7b, 0xf9, 0x15, 0x00, 0x00, 0xff, 0xff,
	0x37, 0xdf, 0xb4, 0xf1, 0xa1, 0x01, 0x00, 0x00,
}
