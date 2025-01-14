// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test/dep-test/dep-test-service/proto/imports/api.proto

package go_api

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

type Pair struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Values               []string `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pair) Reset()         { *m = Pair{} }
func (m *Pair) String() string { return proto.CompactTextString(m) }
func (*Pair) ProtoMessage()    {}
func (*Pair) Descriptor() ([]byte, []int) {
	return fileDescriptor_143fde9720551958, []int{0}
}

func (m *Pair) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pair.Unmarshal(m, b)
}
func (m *Pair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pair.Marshal(b, m, deterministic)
}
func (m *Pair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pair.Merge(m, src)
}
func (m *Pair) XXX_Size() int {
	return xxx_messageInfo_Pair.Size(m)
}
func (m *Pair) XXX_DiscardUnknown() {
	xxx_messageInfo_Pair.DiscardUnknown(m)
}

var xxx_messageInfo_Pair proto.InternalMessageInfo

func (m *Pair) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Pair) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

// A HTTP request as RPC
// Forward by the api handler
type Request struct {
	Method               string           `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	Path                 string           `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Header               map[string]*Pair `protobuf:"bytes,3,rep,name=header,proto3" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Get                  map[string]*Pair `protobuf:"bytes,4,rep,name=get,proto3" json:"get,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Post                 map[string]*Pair `protobuf:"bytes,5,rep,name=post,proto3" json:"post,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Body                 string           `protobuf:"bytes,6,opt,name=body,proto3" json:"body,omitempty"`
	Url                  string           `protobuf:"bytes,7,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_143fde9720551958, []int{1}
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

func (m *Request) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *Request) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Request) GetHeader() map[string]*Pair {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Request) GetGet() map[string]*Pair {
	if m != nil {
		return m.Get
	}
	return nil
}

func (m *Request) GetPost() map[string]*Pair {
	if m != nil {
		return m.Post
	}
	return nil
}

func (m *Request) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *Request) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

// A HTTP response as RPC
// Expected response for the api handler
type Response struct {
	StatusCode           int32            `protobuf:"varint,1,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
	Header               map[string]*Pair `protobuf:"bytes,2,rep,name=header,proto3" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Body                 string           `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_143fde9720551958, []int{2}
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

func (m *Response) GetStatusCode() int32 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

func (m *Response) GetHeader() map[string]*Pair {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Response) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

// A HTTP event as RPC
// Forwarded by the event handler
type Event struct {
	// e.g login
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// uuid
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// unix timestamp of event
	Timestamp int64 `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// event headers
	Header map[string]*Pair `protobuf:"bytes,4,rep,name=header,proto3" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// the event data
	Data                 string   `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_143fde9720551958, []int{3}
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

func (m *Event) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Event) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Event) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Event) GetHeader() map[string]*Pair {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Event) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func init() {
	proto.RegisterType((*Pair)(nil), "go.api.Pair")
	proto.RegisterType((*Request)(nil), "go.api.Request")
	proto.RegisterMapType((map[string]*Pair)(nil), "go.api.Request.GetEntry")
	proto.RegisterMapType((map[string]*Pair)(nil), "go.api.Request.HeaderEntry")
	proto.RegisterMapType((map[string]*Pair)(nil), "go.api.Request.PostEntry")
	proto.RegisterType((*Response)(nil), "go.api.Response")
	proto.RegisterMapType((map[string]*Pair)(nil), "go.api.Response.HeaderEntry")
	proto.RegisterType((*Event)(nil), "go.api.Event")
	proto.RegisterMapType((map[string]*Pair)(nil), "go.api.Event.HeaderEntry")
}

func init() {
	proto.RegisterFile("test/dep-test/dep-test-service/proto/imports/api.proto", fileDescriptor_143fde9720551958)
}

var fileDescriptor_143fde9720551958 = []byte{
	// 411 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0x4d, 0x6f, 0x13, 0x31,
	0x10, 0xd5, 0x7e, 0xb6, 0x99, 0x20, 0x84, 0x7c, 0x40, 0xa6, 0x54, 0x28, 0xda, 0x53, 0x84, 0x94,
	0x0d, 0xb4, 0x08, 0x21, 0xae, 0x10, 0x95, 0x63, 0xb5, 0xff, 0xc0, 0x65, 0x47, 0x8d, 0x45, 0x36,
	0x36, 0xf6, 0x6c, 0xa4, 0xfc, 0x38, 0x0e, 0xfc, 0x0c, 0xfe, 0x0d, 0xf2, 0xac, 0x93, 0xb4, 0x55,
	0xb9, 0xd0, 0xdc, 0xde, 0x8e, 0x9f, 0xdf, 0xbc, 0x79, 0xe3, 0x85, 0x8f, 0x84, 0x9e, 0xe6, 0x2d,
	0xda, 0xd9, 0x3d, 0x30, 0xf3, 0xe8, 0x36, 0xfa, 0x3b, 0xce, 0xad, 0x33, 0x64, 0xe6, 0xba, 0xb3,
	0xc6, 0x91, 0x9f, 0x2b, 0xab, 0x6b, 0xae, 0x88, 0xf2, 0xd6, 0xd4, 0xca, 0xea, 0xea, 0x1d, 0xe4,
	0xd7, 0x4a, 0x3b, 0xf1, 0x02, 0xb2, 0x1f, 0xb8, 0x95, 0xc9, 0x24, 0x99, 0x8e, 0x9a, 0x00, 0xc5,
	0x4b, 0x28, 0x37, 0x6a, 0xd5, 0xa3, 0x97, 0xe9, 0x24, 0x9b, 0x8e, 0x9a, 0xf8, 0x55, 0xfd, 0xce,
	0xe0, 0xa4, 0xc1, 0x9f, 0x3d, 0x7a, 0x0a, 0x9c, 0x0e, 0x69, 0x69, 0xda, 0x78, 0x31, 0x7e, 0x09,
	0x01, 0xb9, 0x55, 0xb4, 0x94, 0x29, 0x57, 0x19, 0x8b, 0x4b, 0x28, 0x97, 0xa8, 0x5a, 0x74, 0x32,
	0x9b, 0x64, 0xd3, 0xf1, 0xc5, 0xeb, 0x7a, 0xb0, 0x50, 0x47, 0xb1, 0xfa, 0x1b, 0x9f, 0x2e, 0xd6,
	0xe4, 0xb6, 0x4d, 0xa4, 0x8a, 0xb7, 0x90, 0xdd, 0x22, 0xc9, 0x9c, 0x6f, 0xc8, 0x87, 0x37, 0xae,
	0x90, 0x06, 0x7a, 0x20, 0x89, 0x19, 0xe4, 0xd6, 0x78, 0x92, 0x05, 0x93, 0x5f, 0x3d, 0x24, 0x5f,
	0x1b, 0x1f, 0xd9, 0x4c, 0x0b, 0x1e, 0x6f, 0x4c, 0xbb, 0x95, 0xe5, 0xe0, 0x31, 0xe0, 0x90, 0x42,
	0xef, 0x56, 0xf2, 0x64, 0x48, 0xa1, 0x77, 0xab, 0xb3, 0x2b, 0x18, 0xdf, 0xf1, 0xf5, 0x48, 0x4c,
	0x15, 0x14, 0x1c, 0x0c, 0xcf, 0x3a, 0xbe, 0x78, 0xb6, 0x6b, 0x1b, 0x52, 0x6d, 0x86, 0xa3, 0xcf,
	0xe9, 0xa7, 0xe4, 0xec, 0x2b, 0x9c, 0xee, 0xec, 0x3e, 0x41, 0x65, 0x01, 0xa3, 0xfd, 0x1c, 0xff,
	0x2f, 0x53, 0xfd, 0x4a, 0xe0, 0xb4, 0x41, 0x6f, 0xcd, 0xda, 0xa3, 0x78, 0x03, 0xe0, 0x49, 0x51,
	0xef, 0xbf, 0x98, 0x16, 0x59, 0xad, 0x68, 0xee, 0x54, 0xc4, 0x87, 0xfd, 0xe2, 0x52, 0x4e, 0xf6,
	0xfc, 0x90, 0xec, 0xa0, 0xf0, 0xe8, 0xe6, 0x76, 0xf1, 0x66, 0x87, 0x78, 0x8f, 0x16, 0x66, 0xf5,
	0x27, 0x81, 0x62, 0xb1, 0xc1, 0x35, 0x6f, 0x71, 0xad, 0x3a, 0x8c, 0x22, 0x8c, 0xc5, 0x73, 0x48,
	0x75, 0x1b, 0xdf, 0x5e, 0xaa, 0x5b, 0x71, 0x0e, 0x23, 0xd2, 0x1d, 0x7a, 0x52, 0x9d, 0x65, 0x3f,
	0x59, 0x73, 0x28, 0x88, 0xf7, 0xfb, 0xf1, 0xf2, 0xfb, 0x0f, 0x87, 0x1b, 0xfc, 0x6b, 0xb6, 0x56,
	0x91, 0x92, 0xc5, 0xd0, 0x34, 0xe0, 0xa3, 0xcd, 0x76, 0x53, 0xf2, 0x0f, 0x7a, 0xf9, 0x37, 0x00,
	0x00, 0xff, 0xff, 0x06, 0xd2, 0x97, 0x45, 0xda, 0x03, 0x00, 0x00,
}
