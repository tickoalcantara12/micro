// Code generated by protoc-gen-go. DO NOT EDIT.
// source: registry/registry.proto

package registry

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

// EventType defines the type of event
type EventType int32

const (
	EventType_Create EventType = 0
	EventType_Delete EventType = 1
	EventType_Update EventType = 2
)

var EventType_name = map[int32]string{
	0: "Create",
	1: "Delete",
	2: "Update",
}

var EventType_value = map[string]int32{
	"Create": 0,
	"Delete": 1,
	"Update": 2,
}

func (x EventType) String() string {
	return proto.EnumName(EventType_name, int32(x))
}

func (EventType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f3f64dc2c9630278, []int{0}
}

// Service represents a go-micro service
type Service struct {
	Name                 string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version              string            `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Metadata             map[string]string `protobuf:"bytes,3,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Endpoints            []*Endpoint       `protobuf:"bytes,4,rep,name=endpoints,proto3" json:"endpoints,omitempty"`
	Nodes                []*Node           `protobuf:"bytes,5,rep,name=nodes,proto3" json:"nodes,omitempty"`
	Options              *Options          `protobuf:"bytes,6,opt,name=options,proto3" json:"options,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Service) Reset()         { *m = Service{} }
func (m *Service) String() string { return proto.CompactTextString(m) }
func (*Service) ProtoMessage()    {}
func (*Service) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3f64dc2c9630278, []int{0}
}

func (m *Service) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Service.Unmarshal(m, b)
}
func (m *Service) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Service.Marshal(b, m, deterministic)
}
func (m *Service) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Service.Merge(m, src)
}
func (m *Service) XXX_Size() int {
	return xxx_messageInfo_Service.Size(m)
}
func (m *Service) XXX_DiscardUnknown() {
	xxx_messageInfo_Service.DiscardUnknown(m)
}

var xxx_messageInfo_Service proto.InternalMessageInfo

func (m *Service) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Service) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Service) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Service) GetEndpoints() []*Endpoint {
	if m != nil {
		return m.Endpoints
	}
	return nil
}

func (m *Service) GetNodes() []*Node {
	if m != nil {
		return m.Nodes
	}
	return nil
}

func (m *Service) GetOptions() *Options {
	if m != nil {
		return m.Options
	}
	return nil
}

// Node represents the node the service is on
type Node struct {
	Id                   string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Address              string            `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Port                 int64             `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"`
	Metadata             map[string]string `protobuf:"bytes,4,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3f64dc2c9630278, []int{1}
}

func (m *Node) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Node.Unmarshal(m, b)
}
func (m *Node) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Node.Marshal(b, m, deterministic)
}
func (m *Node) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Node.Merge(m, src)
}
func (m *Node) XXX_Size() int {
	return xxx_messageInfo_Node.Size(m)
}
func (m *Node) XXX_DiscardUnknown() {
	xxx_messageInfo_Node.DiscardUnknown(m)
}

var xxx_messageInfo_Node proto.InternalMessageInfo

func (m *Node) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Node) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Node) GetPort() int64 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *Node) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

// Endpoint is a endpoint provided by a service
type Endpoint struct {
	Name                 string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Request              *Value            `protobuf:"bytes,2,opt,name=request,proto3" json:"request,omitempty"`
	Response             *Value            `protobuf:"bytes,3,opt,name=response,proto3" json:"response,omitempty"`
	Metadata             map[string]string `protobuf:"bytes,4,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Endpoint) Reset()         { *m = Endpoint{} }
func (m *Endpoint) String() string { return proto.CompactTextString(m) }
func (*Endpoint) ProtoMessage()    {}
func (*Endpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3f64dc2c9630278, []int{2}
}

func (m *Endpoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Endpoint.Unmarshal(m, b)
}
func (m *Endpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Endpoint.Marshal(b, m, deterministic)
}
func (m *Endpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Endpoint.Merge(m, src)
}
func (m *Endpoint) XXX_Size() int {
	return xxx_messageInfo_Endpoint.Size(m)
}
func (m *Endpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_Endpoint.DiscardUnknown(m)
}

var xxx_messageInfo_Endpoint proto.InternalMessageInfo

func (m *Endpoint) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Endpoint) GetRequest() *Value {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *Endpoint) GetResponse() *Value {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *Endpoint) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

// Value is an opaque value for a request or response
type Value struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Values               []*Value `protobuf:"bytes,3,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Value) Reset()         { *m = Value{} }
func (m *Value) String() string { return proto.CompactTextString(m) }
func (*Value) ProtoMessage()    {}
func (*Value) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3f64dc2c9630278, []int{3}
}

func (m *Value) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Value.Unmarshal(m, b)
}
func (m *Value) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Value.Marshal(b, m, deterministic)
}
func (m *Value) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Value.Merge(m, src)
}
func (m *Value) XXX_Size() int {
	return xxx_messageInfo_Value.Size(m)
}
func (m *Value) XXX_DiscardUnknown() {
	xxx_messageInfo_Value.DiscardUnknown(m)
}

var xxx_messageInfo_Value proto.InternalMessageInfo

func (m *Value) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Value) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Value) GetValues() []*Value {
	if m != nil {
		return m.Values
	}
	return nil
}

// Options are registry options
type Options struct {
	Ttl                  int64    `protobuf:"varint,1,opt,name=ttl,proto3" json:"ttl,omitempty"`
	Domain               string   `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Options) Reset()         { *m = Options{} }
func (m *Options) String() string { return proto.CompactTextString(m) }
func (*Options) ProtoMessage()    {}
func (*Options) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3f64dc2c9630278, []int{4}
}

func (m *Options) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Options.Unmarshal(m, b)
}
func (m *Options) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Options.Marshal(b, m, deterministic)
}
func (m *Options) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Options.Merge(m, src)
}
func (m *Options) XXX_Size() int {
	return xxx_messageInfo_Options.Size(m)
}
func (m *Options) XXX_DiscardUnknown() {
	xxx_messageInfo_Options.DiscardUnknown(m)
}

var xxx_messageInfo_Options proto.InternalMessageInfo

func (m *Options) GetTtl() int64 {
	if m != nil {
		return m.Ttl
	}
	return 0
}

func (m *Options) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

// Result is returns by the watcher
type Result struct {
	Action               string   `protobuf:"bytes,1,opt,name=action,proto3" json:"action,omitempty"`
	Service              *Service `protobuf:"bytes,2,opt,name=service,proto3" json:"service,omitempty"`
	Timestamp            int64    `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3f64dc2c9630278, []int{5}
}

func (m *Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Result.Unmarshal(m, b)
}
func (m *Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Result.Marshal(b, m, deterministic)
}
func (m *Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Result.Merge(m, src)
}
func (m *Result) XXX_Size() int {
	return xxx_messageInfo_Result.Size(m)
}
func (m *Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Result.DiscardUnknown(m)
}

var xxx_messageInfo_Result proto.InternalMessageInfo

func (m *Result) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *Result) GetService() *Service {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *Result) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type EmptyResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmptyResponse) Reset()         { *m = EmptyResponse{} }
func (m *EmptyResponse) String() string { return proto.CompactTextString(m) }
func (*EmptyResponse) ProtoMessage()    {}
func (*EmptyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3f64dc2c9630278, []int{6}
}

func (m *EmptyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmptyResponse.Unmarshal(m, b)
}
func (m *EmptyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmptyResponse.Marshal(b, m, deterministic)
}
func (m *EmptyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmptyResponse.Merge(m, src)
}
func (m *EmptyResponse) XXX_Size() int {
	return xxx_messageInfo_EmptyResponse.Size(m)
}
func (m *EmptyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EmptyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EmptyResponse proto.InternalMessageInfo

type GetRequest struct {
	Service              string   `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	Options              *Options `protobuf:"bytes,2,opt,name=options,proto3" json:"options,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3f64dc2c9630278, []int{7}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *GetRequest) GetOptions() *Options {
	if m != nil {
		return m.Options
	}
	return nil
}

type GetResponse struct {
	Services             []*Service `protobuf:"bytes,1,rep,name=services,proto3" json:"services,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3f64dc2c9630278, []int{8}
}

func (m *GetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResponse.Unmarshal(m, b)
}
func (m *GetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResponse.Marshal(b, m, deterministic)
}
func (m *GetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResponse.Merge(m, src)
}
func (m *GetResponse) XXX_Size() int {
	return xxx_messageInfo_GetResponse.Size(m)
}
func (m *GetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetResponse proto.InternalMessageInfo

func (m *GetResponse) GetServices() []*Service {
	if m != nil {
		return m.Services
	}
	return nil
}

type ListRequest struct {
	Options              *Options `protobuf:"bytes,1,opt,name=options,proto3" json:"options,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3f64dc2c9630278, []int{9}
}

func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (m *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(m, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

func (m *ListRequest) GetOptions() *Options {
	if m != nil {
		return m.Options
	}
	return nil
}

type ListResponse struct {
	Services             []*Service `protobuf:"bytes,1,rep,name=services,proto3" json:"services,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3f64dc2c9630278, []int{10}
}

func (m *ListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse.Unmarshal(m, b)
}
func (m *ListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse.Marshal(b, m, deterministic)
}
func (m *ListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse.Merge(m, src)
}
func (m *ListResponse) XXX_Size() int {
	return xxx_messageInfo_ListResponse.Size(m)
}
func (m *ListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse proto.InternalMessageInfo

func (m *ListResponse) GetServices() []*Service {
	if m != nil {
		return m.Services
	}
	return nil
}

type WatchRequest struct {
	// service is optional
	Service              string   `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	Options              *Options `protobuf:"bytes,2,opt,name=options,proto3" json:"options,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WatchRequest) Reset()         { *m = WatchRequest{} }
func (m *WatchRequest) String() string { return proto.CompactTextString(m) }
func (*WatchRequest) ProtoMessage()    {}
func (*WatchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3f64dc2c9630278, []int{11}
}

func (m *WatchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WatchRequest.Unmarshal(m, b)
}
func (m *WatchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WatchRequest.Marshal(b, m, deterministic)
}
func (m *WatchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WatchRequest.Merge(m, src)
}
func (m *WatchRequest) XXX_Size() int {
	return xxx_messageInfo_WatchRequest.Size(m)
}
func (m *WatchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WatchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WatchRequest proto.InternalMessageInfo

func (m *WatchRequest) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *WatchRequest) GetOptions() *Options {
	if m != nil {
		return m.Options
	}
	return nil
}

// Event is registry event
type Event struct {
	// Event Id
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// type of event
	Type EventType `protobuf:"varint,2,opt,name=type,proto3,enum=registry.EventType" json:"type,omitempty"`
	// unix timestamp of event
	Timestamp int64 `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// service entry
	Service              *Service `protobuf:"bytes,4,opt,name=service,proto3" json:"service,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3f64dc2c9630278, []int{12}
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

func (m *Event) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Event) GetType() EventType {
	if m != nil {
		return m.Type
	}
	return EventType_Create
}

func (m *Event) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Event) GetService() *Service {
	if m != nil {
		return m.Service
	}
	return nil
}

func init() {
	proto.RegisterEnum("registry.EventType", EventType_name, EventType_value)
	proto.RegisterType((*Service)(nil), "registry.Service")
	proto.RegisterMapType((map[string]string)(nil), "registry.Service.MetadataEntry")
	proto.RegisterType((*Node)(nil), "registry.Node")
	proto.RegisterMapType((map[string]string)(nil), "registry.Node.MetadataEntry")
	proto.RegisterType((*Endpoint)(nil), "registry.Endpoint")
	proto.RegisterMapType((map[string]string)(nil), "registry.Endpoint.MetadataEntry")
	proto.RegisterType((*Value)(nil), "registry.Value")
	proto.RegisterType((*Options)(nil), "registry.Options")
	proto.RegisterType((*Result)(nil), "registry.Result")
	proto.RegisterType((*EmptyResponse)(nil), "registry.EmptyResponse")
	proto.RegisterType((*GetRequest)(nil), "registry.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "registry.GetResponse")
	proto.RegisterType((*ListRequest)(nil), "registry.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "registry.ListResponse")
	proto.RegisterType((*WatchRequest)(nil), "registry.WatchRequest")
	proto.RegisterType((*Event)(nil), "registry.Event")
}

func init() { proto.RegisterFile("registry/registry.proto", fileDescriptor_f3f64dc2c9630278) }

var fileDescriptor_f3f64dc2c9630278 = []byte{
	// 709 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0xdb, 0x6e, 0xd3, 0x4c,
	0x10, 0x8e, 0x9d, 0xf3, 0xa4, 0x87, 0x74, 0xff, 0xfe, 0xad, 0x15, 0x55, 0x22, 0xb2, 0x90, 0x1a,
	0xa8, 0x48, 0x4a, 0x22, 0xa4, 0xaa, 0x29, 0x42, 0x82, 0x46, 0xdc, 0x70, 0x90, 0x5c, 0x0a, 0x88,
	0x3b, 0x37, 0x1e, 0xb5, 0x56, 0x63, 0xaf, 0xd9, 0xdd, 0x44, 0xca, 0x33, 0xf0, 0x3c, 0x70, 0xc3,
	0x53, 0xf1, 0x06, 0x68, 0xd7, 0xeb, 0x43, 0x9b, 0x40, 0xa4, 0x02, 0x37, 0xd1, 0xec, 0xcc, 0x37,
	0xe7, 0x6f, 0x62, 0xd8, 0x65, 0x78, 0xe9, 0x73, 0xc1, 0xe6, 0xbd, 0x44, 0xe8, 0x46, 0x8c, 0x0a,
	0x4a, 0x6a, 0xc9, 0xdb, 0xfe, 0x66, 0x42, 0xf5, 0x0c, 0xd9, 0xcc, 0x1f, 0x23, 0x21, 0x50, 0x0a,
	0xdd, 0x00, 0x2d, 0xa3, 0x6d, 0x74, 0xea, 0x8e, 0x92, 0x89, 0x05, 0xd5, 0x19, 0x32, 0xee, 0xd3,
	0xd0, 0x32, 0x95, 0x3a, 0x79, 0x92, 0x21, 0xd4, 0x02, 0x14, 0xae, 0xe7, 0x0a, 0xd7, 0x2a, 0xb6,
	0x8b, 0x9d, 0x46, 0xff, 0x5e, 0x37, 0x4d, 0xa3, 0x43, 0x76, 0x5f, 0x6b, 0xc4, 0x28, 0x14, 0x6c,
	0xee, 0xa4, 0x0e, 0xe4, 0x10, 0xea, 0x18, 0x7a, 0x11, 0xf5, 0x43, 0xc1, 0xad, 0x92, 0xf2, 0x26,
	0x99, 0xf7, 0x48, 0x9b, 0x9c, 0x0c, 0x44, 0xee, 0x43, 0x39, 0xa4, 0x1e, 0x72, 0xab, 0xac, 0xd0,
	0x1b, 0x19, 0xfa, 0x0d, 0xf5, 0xd0, 0x89, 0x8d, 0xe4, 0x00, 0xaa, 0x34, 0x12, 0x3e, 0x0d, 0xb9,
	0x55, 0x69, 0x1b, 0x9d, 0x46, 0x7f, 0x2b, 0xc3, 0xbd, 0x8d, 0x0d, 0x4e, 0x82, 0x68, 0x0d, 0x61,
	0xfd, 0x46, 0x7d, 0xa4, 0x09, 0xc5, 0x6b, 0x9c, 0xeb, 0xfe, 0xa5, 0x48, 0xb6, 0xa1, 0x3c, 0x73,
	0x27, 0x53, 0xd4, 0xcd, 0xc7, 0x8f, 0x63, 0xf3, 0xc8, 0xb0, 0xbf, 0x1b, 0x50, 0x92, 0x99, 0xc9,
	0x06, 0x98, 0xbe, 0xa7, 0x7d, 0x4c, 0xdf, 0x93, 0x13, 0x73, 0x3d, 0x8f, 0x21, 0xe7, 0xc9, 0xc4,
	0xf4, 0x53, 0xce, 0x37, 0xa2, 0x4c, 0x58, 0xc5, 0xb6, 0xd1, 0x29, 0x3a, 0x4a, 0x26, 0x47, 0xb9,
	0x29, 0xc6, 0x73, 0xd8, 0xbb, 0xd9, 0xd9, 0xaf, 0x46, 0xf8, 0x67, 0xd5, 0xff, 0x30, 0xa0, 0x96,
	0x4c, 0x79, 0xe9, 0xde, 0x1f, 0x40, 0x95, 0xe1, 0xe7, 0x29, 0x72, 0xa1, 0x9c, 0x1b, 0xfd, 0xcd,
	0xac, 0xac, 0xf7, 0x32, 0x8c, 0x93, 0xd8, 0xc9, 0x01, 0xd4, 0x18, 0xf2, 0x88, 0x86, 0x1c, 0x55,
	0x6b, 0x4b, 0xb0, 0x29, 0x80, 0x9c, 0x2c, 0xf4, 0xdb, 0x5e, 0xdc, 0xfb, 0xbf, 0xe9, 0xf9, 0x23,
	0x94, 0x55, 0x35, 0x4b, 0xfb, 0x25, 0x50, 0x12, 0xf3, 0x28, 0xf1, 0x52, 0x32, 0xd9, 0x87, 0x8a,
	0xf2, 0xe6, 0x9a, 0xdf, 0x0b, 0x6d, 0x69, 0xb3, 0x3d, 0x80, 0xaa, 0x26, 0x97, 0x2c, 0x48, 0x88,
	0x89, 0x0a, 0x5d, 0x74, 0xa4, 0x48, 0x76, 0xa0, 0xe2, 0xd1, 0xc0, 0xf5, 0x93, 0x03, 0xd2, 0x2f,
	0xfb, 0x1a, 0x2a, 0x0e, 0xf2, 0xe9, 0x44, 0x48, 0x84, 0x3b, 0x96, 0xee, 0xba, 0x22, 0xfd, 0x92,
	0x64, 0xe6, 0xf1, 0x1d, 0xe9, 0x1d, 0x6c, 0x2d, 0x1c, 0x98, 0x93, 0x20, 0xc8, 0x1e, 0xd4, 0x85,
	0x1f, 0x20, 0x17, 0x6e, 0x10, 0x69, 0x86, 0x65, 0x0a, 0x7b, 0x13, 0xd6, 0x47, 0x41, 0x24, 0xe6,
	0x8e, 0xde, 0x83, 0x7d, 0x06, 0xf0, 0x12, 0x85, 0xa3, 0x57, 0x68, 0x65, 0x99, 0xe2, 0x12, 0xd2,
	0xb0, 0xb9, 0x83, 0x32, 0x57, 0x1d, 0x94, 0x7d, 0x02, 0x0d, 0x15, 0x54, 0xef, 0xfa, 0x11, 0xd4,
	0x74, 0x18, 0x6e, 0x19, 0x6a, 0x82, 0x4b, 0x1a, 0x48, 0x21, 0xf6, 0x31, 0x34, 0x5e, 0xf9, 0x3c,
	0xad, 0x29, 0x97, 0xd9, 0x58, 0x99, 0xf9, 0x29, 0xac, 0xc5, 0xbe, 0x77, 0x4b, 0x7d, 0x0e, 0x6b,
	0x1f, 0x5c, 0x31, 0xbe, 0xfa, 0xcb, 0xf3, 0xf8, 0x62, 0x40, 0x79, 0x34, 0xc3, 0x50, 0x2c, 0xfc,
	0x49, 0xec, 0xe7, 0xe8, 0xb6, 0xd1, 0xff, 0x2f, 0x77, 0x02, 0x12, 0xfe, 0x6e, 0x1e, 0xa1, 0xe6,
	0xe0, 0x6f, 0xd7, 0x9a, 0x67, 0x48, 0x69, 0x15, 0x43, 0x1e, 0xf6, 0xa0, 0x9e, 0x46, 0x27, 0x00,
	0x95, 0x17, 0x0c, 0x5d, 0x81, 0xcd, 0x82, 0x94, 0x4f, 0x71, 0x82, 0x02, 0x9b, 0x86, 0x94, 0xcf,
	0x23, 0x4f, 0xea, 0xcd, 0xfe, 0x57, 0x13, 0x6a, 0x8e, 0x0e, 0x47, 0x86, 0x8a, 0x30, 0xc9, 0xa7,
	0x62, 0x3b, 0xcb, 0x93, 0xd1, 0xa8, 0xf5, 0xff, 0x2d, 0xad, 0xe6, 0x5a, 0x41, 0xfe, 0xcb, 0xc5,
	0x81, 0x90, 0x91, 0xc5, 0x12, 0x5b, 0xbb, 0xb9, 0xfe, 0x6f, 0xb0, 0xb4, 0x40, 0x8e, 0x01, 0x4e,
	0x91, 0xdd, 0xcd, 0xf7, 0x59, 0x4c, 0x0a, 0x8d, 0xe4, 0x24, 0x57, 0x5e, 0x8e, 0x68, 0xad, 0x9d,
	0xdb, 0xea, 0x34, 0xc0, 0x13, 0x28, 0x2b, 0x5a, 0x90, 0x1c, 0x24, 0xcf, 0x93, 0x56, 0x33, 0xd3,
	0xc7, 0xb7, 0x6c, 0x17, 0x0e, 0x8d, 0xe7, 0x83, 0x4f, 0x8f, 0x2f, 0x7d, 0x71, 0x35, 0xbd, 0xe8,
	0x8e, 0x69, 0xd0, 0x0b, 0xfc, 0x31, 0xa3, 0xfa, 0x77, 0x36, 0xe8, 0xa9, 0x0f, 0x70, 0xfa, 0x3d,
	0x1e, 0x26, 0xc2, 0x45, 0x45, 0x19, 0x06, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xde, 0x1c, 0x89,
	0xfa, 0xb4, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RegistryClient is the client API for Registry service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RegistryClient interface {
	GetService(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	Register(ctx context.Context, in *Service, opts ...grpc.CallOption) (*EmptyResponse, error)
	Deregister(ctx context.Context, in *Service, opts ...grpc.CallOption) (*EmptyResponse, error)
	ListServices(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	Watch(ctx context.Context, in *WatchRequest, opts ...grpc.CallOption) (Registry_WatchClient, error)
}

type registryClient struct {
	cc *grpc.ClientConn
}

func NewRegistryClient(cc *grpc.ClientConn) RegistryClient {
	return &registryClient{cc}
}

func (c *registryClient) GetService(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/registry.Registry/GetService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) Register(ctx context.Context, in *Service, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/registry.Registry/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) Deregister(ctx context.Context, in *Service, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/registry.Registry/Deregister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) ListServices(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/registry.Registry/ListServices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) Watch(ctx context.Context, in *WatchRequest, opts ...grpc.CallOption) (Registry_WatchClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Registry_serviceDesc.Streams[0], "/registry.Registry/Watch", opts...)
	if err != nil {
		return nil, err
	}
	x := &registryWatchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Registry_WatchClient interface {
	Recv() (*Result, error)
	grpc.ClientStream
}

type registryWatchClient struct {
	grpc.ClientStream
}

func (x *registryWatchClient) Recv() (*Result, error) {
	m := new(Result)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RegistryServer is the server API for Registry service.
type RegistryServer interface {
	GetService(context.Context, *GetRequest) (*GetResponse, error)
	Register(context.Context, *Service) (*EmptyResponse, error)
	Deregister(context.Context, *Service) (*EmptyResponse, error)
	ListServices(context.Context, *ListRequest) (*ListResponse, error)
	Watch(*WatchRequest, Registry_WatchServer) error
}

func RegisterRegistryServer(s *grpc.Server, srv RegistryServer) {
	s.RegisterService(&_Registry_serviceDesc, srv)
}

func _Registry_GetService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).GetService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/registry.Registry/GetService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).GetService(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Service)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/registry.Registry/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).Register(ctx, req.(*Service))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_Deregister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Service)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).Deregister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/registry.Registry/Deregister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).Deregister(ctx, req.(*Service))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_ListServices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).ListServices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/registry.Registry/ListServices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).ListServices(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_Watch_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(WatchRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RegistryServer).Watch(m, &registryWatchServer{stream})
}

type Registry_WatchServer interface {
	Send(*Result) error
	grpc.ServerStream
}

type registryWatchServer struct {
	grpc.ServerStream
}

func (x *registryWatchServer) Send(m *Result) error {
	return x.ServerStream.SendMsg(m)
}

var _Registry_serviceDesc = grpc.ServiceDesc{
	ServiceName: "registry.Registry",
	HandlerType: (*RegistryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetService",
			Handler:    _Registry_GetService_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _Registry_Register_Handler,
		},
		{
			MethodName: "Deregister",
			Handler:    _Registry_Deregister_Handler,
		},
		{
			MethodName: "ListServices",
			Handler:    _Registry_ListServices_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Watch",
			Handler:       _Registry_Watch_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "registry/registry.proto",
}
