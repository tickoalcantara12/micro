// Code generated by protoc-gen-go. DO NOT EDIT.
// source: server/grpc/proto/test.proto

package test

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Request struct {
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb9c685b7640cf1e, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *Request) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Response struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb9c685b7640cf1e, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "Request")
	proto.RegisterType((*Response)(nil), "Response")
}

func init() { proto.RegisterFile("server/grpc/proto/test.proto", fileDescriptor_bb9c685b7640cf1e) }

var fileDescriptor_bb9c685b7640cf1e = []byte{
	// 261 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x29, 0x4e, 0x2d, 0x2a,
	0x4b, 0x2d, 0xd2, 0x4f, 0x2f, 0x2a, 0x48, 0xd6, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x2f, 0x49,
	0x2d, 0x2e, 0xd1, 0x03, 0x33, 0xa5, 0x64, 0xd2, 0xf3, 0xf3, 0xd3, 0x73, 0x52, 0xf5, 0x13, 0x0b,
	0x32, 0xf5, 0x13, 0xf3, 0xf2, 0xf2, 0x4b, 0x12, 0x4b, 0x32, 0xf3, 0xf3, 0x8a, 0x21, 0xb2, 0x4a,
	0x86, 0x5c, 0xec, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x42, 0x5c, 0x2c, 0xa5, 0xa5,
	0x99, 0x29, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x60, 0x36, 0x48, 0x2c, 0x2f, 0x31, 0x37,
	0x55, 0x82, 0x09, 0x22, 0x06, 0x62, 0x2b, 0xc9, 0x70, 0x71, 0x04, 0xa5, 0x16, 0x17, 0xe4, 0xe7,
	0x15, 0xa7, 0x0a, 0x09, 0x70, 0x31, 0xe7, 0x16, 0xa7, 0x43, 0xb5, 0x80, 0x98, 0x46, 0xcf, 0x18,
	0xb9, 0x58, 0x42, 0x40, 0xc6, 0x39, 0x70, 0xb1, 0x38, 0x27, 0xe6, 0xe4, 0x08, 0x71, 0xe8, 0x41,
	0x2d, 0x90, 0xe2, 0xd4, 0x83, 0xe9, 0x53, 0x52, 0x6e, 0xba, 0xfc, 0x64, 0x32, 0x93, 0xac, 0x92,
	0x04, 0xd8, 0x59, 0x65, 0x06, 0x60, 0x07, 0xeb, 0x27, 0x27, 0xe6, 0xe4, 0xe8, 0x57, 0x83, 0x2c,
	0xae, 0xb5, 0x62, 0xd4, 0x12, 0x72, 0xe3, 0xe2, 0x00, 0x99, 0x10, 0x90, 0x5c, 0x94, 0x8a, 0xdd,
	0x14, 0x55, 0xb0, 0x29, 0xf2, 0x4a, 0x52, 0x71, 0x98, 0xc6, 0x14, 0x24, 0x17, 0xa5, 0xea, 0xdb,
	0xab, 0x80, 0xcc, 0x09, 0xe1, 0xe2, 0x87, 0x99, 0xe3, 0x99, 0x57, 0x96, 0x98, 0x93, 0x99, 0x82,
	0xdd, 0x38, 0x1d, 0xb0, 0x71, 0x6a, 0x4a, 0x8a, 0xb8, 0x8c, 0xcb, 0x84, 0xe8, 0xd6, 0xb7, 0xb7,
	0x62, 0xd4, 0x4a, 0x62, 0x03, 0x07, 0xa0, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xf6, 0x32, 0x18,
	0x0f, 0x7e, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TestClient is the client API for Test service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TestClient interface {
	Call(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	CallPcre(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	CallPcreInvalid(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type testClient struct {
	cc *grpc.ClientConn
}

func NewTestClient(cc *grpc.ClientConn) TestClient {
	return &testClient{cc}
}

func (c *testClient) Call(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/Test/Call", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testClient) CallPcre(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/Test/CallPcre", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testClient) CallPcreInvalid(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/Test/CallPcreInvalid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TestServer is the server API for Test service.
type TestServer interface {
	Call(context.Context, *Request) (*Response, error)
	CallPcre(context.Context, *Request) (*Response, error)
	CallPcreInvalid(context.Context, *Request) (*Response, error)
}

// UnimplementedTestServer can be embedded to have forward compatible implementations.
type UnimplementedTestServer struct {
}

func (*UnimplementedTestServer) Call(ctx context.Context, req *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Call not implemented")
}
func (*UnimplementedTestServer) CallPcre(ctx context.Context, req *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CallPcre not implemented")
}
func (*UnimplementedTestServer) CallPcreInvalid(ctx context.Context, req *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CallPcreInvalid not implemented")
}

func RegisterTestServer(s *grpc.Server, srv TestServer) {
	s.RegisterService(&_Test_serviceDesc, srv)
}

func _Test_Call_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServer).Call(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Test/Call",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServer).Call(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Test_CallPcre_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServer).CallPcre(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Test/CallPcre",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServer).CallPcre(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Test_CallPcreInvalid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServer).CallPcreInvalid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Test/CallPcreInvalid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServer).CallPcreInvalid(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Test_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Test",
	HandlerType: (*TestServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Call",
			Handler:    _Test_Call_Handler,
		},
		{
			MethodName: "CallPcre",
			Handler:    _Test_CallPcre_Handler,
		},
		{
			MethodName: "CallPcreInvalid",
			Handler:    _Test_CallPcreInvalid_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server/grpc/proto/test.proto",
}
