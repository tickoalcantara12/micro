// Code generated by protoc-gen-go. DO NOT EDIT.
// source: debug/debug.proto

package debug

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

type SpanType int32

const (
	SpanType_INBOUND  SpanType = 0
	SpanType_OUTBOUND SpanType = 1
)

var SpanType_name = map[int32]string{
	0: "INBOUND",
	1: "OUTBOUND",
}

var SpanType_value = map[string]int32{
	"INBOUND":  0,
	"OUTBOUND": 1,
}

func (x SpanType) String() string {
	return proto.EnumName(SpanType_name, int32(x))
}

func (SpanType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5ae24eab94cb53d5, []int{0}
}

type HealthRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthRequest) Reset()         { *m = HealthRequest{} }
func (m *HealthRequest) String() string { return proto.CompactTextString(m) }
func (*HealthRequest) ProtoMessage()    {}
func (*HealthRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ae24eab94cb53d5, []int{0}
}

func (m *HealthRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthRequest.Unmarshal(m, b)
}
func (m *HealthRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthRequest.Marshal(b, m, deterministic)
}
func (m *HealthRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthRequest.Merge(m, src)
}
func (m *HealthRequest) XXX_Size() int {
	return xxx_messageInfo_HealthRequest.Size(m)
}
func (m *HealthRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HealthRequest proto.InternalMessageInfo

type HealthResponse struct {
	// default: ok
	Status               string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthResponse) Reset()         { *m = HealthResponse{} }
func (m *HealthResponse) String() string { return proto.CompactTextString(m) }
func (*HealthResponse) ProtoMessage()    {}
func (*HealthResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ae24eab94cb53d5, []int{1}
}

func (m *HealthResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthResponse.Unmarshal(m, b)
}
func (m *HealthResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthResponse.Marshal(b, m, deterministic)
}
func (m *HealthResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthResponse.Merge(m, src)
}
func (m *HealthResponse) XXX_Size() int {
	return xxx_messageInfo_HealthResponse.Size(m)
}
func (m *HealthResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HealthResponse proto.InternalMessageInfo

func (m *HealthResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type StatsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatsRequest) Reset()         { *m = StatsRequest{} }
func (m *StatsRequest) String() string { return proto.CompactTextString(m) }
func (*StatsRequest) ProtoMessage()    {}
func (*StatsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ae24eab94cb53d5, []int{2}
}

func (m *StatsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatsRequest.Unmarshal(m, b)
}
func (m *StatsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatsRequest.Marshal(b, m, deterministic)
}
func (m *StatsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatsRequest.Merge(m, src)
}
func (m *StatsRequest) XXX_Size() int {
	return xxx_messageInfo_StatsRequest.Size(m)
}
func (m *StatsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StatsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StatsRequest proto.InternalMessageInfo

type StatsResponse struct {
	// timestamp of recording
	Timestamp uint64 `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// unix timestamp
	Started uint64 `protobuf:"varint,2,opt,name=started,proto3" json:"started,omitempty"`
	// in seconds
	Uptime uint64 `protobuf:"varint,3,opt,name=uptime,proto3" json:"uptime,omitempty"`
	// in bytes
	Memory uint64 `protobuf:"varint,4,opt,name=memory,proto3" json:"memory,omitempty"`
	// num threads
	Threads uint64 `protobuf:"varint,5,opt,name=threads,proto3" json:"threads,omitempty"`
	// total gc in nanoseconds
	Gc uint64 `protobuf:"varint,6,opt,name=gc,proto3" json:"gc,omitempty"`
	// total number of requests
	Requests uint64 `protobuf:"varint,7,opt,name=requests,proto3" json:"requests,omitempty"`
	// total number of errors
	Errors               uint64   `protobuf:"varint,8,opt,name=errors,proto3" json:"errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatsResponse) Reset()         { *m = StatsResponse{} }
func (m *StatsResponse) String() string { return proto.CompactTextString(m) }
func (*StatsResponse) ProtoMessage()    {}
func (*StatsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ae24eab94cb53d5, []int{3}
}

func (m *StatsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatsResponse.Unmarshal(m, b)
}
func (m *StatsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatsResponse.Marshal(b, m, deterministic)
}
func (m *StatsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatsResponse.Merge(m, src)
}
func (m *StatsResponse) XXX_Size() int {
	return xxx_messageInfo_StatsResponse.Size(m)
}
func (m *StatsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StatsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StatsResponse proto.InternalMessageInfo

func (m *StatsResponse) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *StatsResponse) GetStarted() uint64 {
	if m != nil {
		return m.Started
	}
	return 0
}

func (m *StatsResponse) GetUptime() uint64 {
	if m != nil {
		return m.Uptime
	}
	return 0
}

func (m *StatsResponse) GetMemory() uint64 {
	if m != nil {
		return m.Memory
	}
	return 0
}

func (m *StatsResponse) GetThreads() uint64 {
	if m != nil {
		return m.Threads
	}
	return 0
}

func (m *StatsResponse) GetGc() uint64 {
	if m != nil {
		return m.Gc
	}
	return 0
}

func (m *StatsResponse) GetRequests() uint64 {
	if m != nil {
		return m.Requests
	}
	return 0
}

func (m *StatsResponse) GetErrors() uint64 {
	if m != nil {
		return m.Errors
	}
	return 0
}

// LogRequest requests service logs
type LogRequest struct {
	// count of records to request
	Count int64 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	// relative time in seconds
	// before the current time
	// from which to show logs
	Since                int64    `protobuf:"varint,2,opt,name=since,proto3" json:"since,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogRequest) Reset()         { *m = LogRequest{} }
func (m *LogRequest) String() string { return proto.CompactTextString(m) }
func (*LogRequest) ProtoMessage()    {}
func (*LogRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ae24eab94cb53d5, []int{4}
}

func (m *LogRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogRequest.Unmarshal(m, b)
}
func (m *LogRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogRequest.Marshal(b, m, deterministic)
}
func (m *LogRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogRequest.Merge(m, src)
}
func (m *LogRequest) XXX_Size() int {
	return xxx_messageInfo_LogRequest.Size(m)
}
func (m *LogRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LogRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LogRequest proto.InternalMessageInfo

func (m *LogRequest) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *LogRequest) GetSince() int64 {
	if m != nil {
		return m.Since
	}
	return 0
}

// LogResponse returns a list of logs
type LogResponse struct {
	Records              []*Record `protobuf:"bytes,1,rep,name=records,proto3" json:"records,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *LogResponse) Reset()         { *m = LogResponse{} }
func (m *LogResponse) String() string { return proto.CompactTextString(m) }
func (*LogResponse) ProtoMessage()    {}
func (*LogResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ae24eab94cb53d5, []int{5}
}

func (m *LogResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogResponse.Unmarshal(m, b)
}
func (m *LogResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogResponse.Marshal(b, m, deterministic)
}
func (m *LogResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogResponse.Merge(m, src)
}
func (m *LogResponse) XXX_Size() int {
	return xxx_messageInfo_LogResponse.Size(m)
}
func (m *LogResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LogResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LogResponse proto.InternalMessageInfo

func (m *LogResponse) GetRecords() []*Record {
	if m != nil {
		return m.Records
	}
	return nil
}

// Record is service log record
type Record struct {
	// timestamp of log record
	Timestamp int64 `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// record metadata
	Metadata map[string]string `protobuf:"bytes,2,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// message
	Message              string   `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Record) Reset()         { *m = Record{} }
func (m *Record) String() string { return proto.CompactTextString(m) }
func (*Record) ProtoMessage()    {}
func (*Record) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ae24eab94cb53d5, []int{6}
}

func (m *Record) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Record.Unmarshal(m, b)
}
func (m *Record) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Record.Marshal(b, m, deterministic)
}
func (m *Record) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Record.Merge(m, src)
}
func (m *Record) XXX_Size() int {
	return xxx_messageInfo_Record.Size(m)
}
func (m *Record) XXX_DiscardUnknown() {
	xxx_messageInfo_Record.DiscardUnknown(m)
}

var xxx_messageInfo_Record proto.InternalMessageInfo

func (m *Record) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Record) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Record) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type TraceRequest struct {
	// trace id to retrieve
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TraceRequest) Reset()         { *m = TraceRequest{} }
func (m *TraceRequest) String() string { return proto.CompactTextString(m) }
func (*TraceRequest) ProtoMessage()    {}
func (*TraceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ae24eab94cb53d5, []int{7}
}

func (m *TraceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TraceRequest.Unmarshal(m, b)
}
func (m *TraceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TraceRequest.Marshal(b, m, deterministic)
}
func (m *TraceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TraceRequest.Merge(m, src)
}
func (m *TraceRequest) XXX_Size() int {
	return xxx_messageInfo_TraceRequest.Size(m)
}
func (m *TraceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TraceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TraceRequest proto.InternalMessageInfo

func (m *TraceRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type TraceResponse struct {
	Spans                []*Span  `protobuf:"bytes,1,rep,name=spans,proto3" json:"spans,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TraceResponse) Reset()         { *m = TraceResponse{} }
func (m *TraceResponse) String() string { return proto.CompactTextString(m) }
func (*TraceResponse) ProtoMessage()    {}
func (*TraceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ae24eab94cb53d5, []int{8}
}

func (m *TraceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TraceResponse.Unmarshal(m, b)
}
func (m *TraceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TraceResponse.Marshal(b, m, deterministic)
}
func (m *TraceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TraceResponse.Merge(m, src)
}
func (m *TraceResponse) XXX_Size() int {
	return xxx_messageInfo_TraceResponse.Size(m)
}
func (m *TraceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TraceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TraceResponse proto.InternalMessageInfo

func (m *TraceResponse) GetSpans() []*Span {
	if m != nil {
		return m.Spans
	}
	return nil
}

type Span struct {
	// the trace id
	Trace string `protobuf:"bytes,1,opt,name=trace,proto3" json:"trace,omitempty"`
	// id of the span
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// parent span
	Parent string `protobuf:"bytes,3,opt,name=parent,proto3" json:"parent,omitempty"`
	// name of the resource
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// time of start in nanoseconds
	Started uint64 `protobuf:"varint,5,opt,name=started,proto3" json:"started,omitempty"`
	// duration of the execution in nanoseconds
	Duration uint64 `protobuf:"varint,6,opt,name=duration,proto3" json:"duration,omitempty"`
	// associated metadata
	Metadata             map[string]string `protobuf:"bytes,7,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Type                 SpanType          `protobuf:"varint,8,opt,name=type,proto3,enum=debug.SpanType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Span) Reset()         { *m = Span{} }
func (m *Span) String() string { return proto.CompactTextString(m) }
func (*Span) ProtoMessage()    {}
func (*Span) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ae24eab94cb53d5, []int{9}
}

func (m *Span) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Span.Unmarshal(m, b)
}
func (m *Span) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Span.Marshal(b, m, deterministic)
}
func (m *Span) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Span.Merge(m, src)
}
func (m *Span) XXX_Size() int {
	return xxx_messageInfo_Span.Size(m)
}
func (m *Span) XXX_DiscardUnknown() {
	xxx_messageInfo_Span.DiscardUnknown(m)
}

var xxx_messageInfo_Span proto.InternalMessageInfo

func (m *Span) GetTrace() string {
	if m != nil {
		return m.Trace
	}
	return ""
}

func (m *Span) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Span) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *Span) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Span) GetStarted() uint64 {
	if m != nil {
		return m.Started
	}
	return 0
}

func (m *Span) GetDuration() uint64 {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *Span) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Span) GetType() SpanType {
	if m != nil {
		return m.Type
	}
	return SpanType_INBOUND
}

func init() {
	proto.RegisterEnum("debug.SpanType", SpanType_name, SpanType_value)
	proto.RegisterType((*HealthRequest)(nil), "debug.HealthRequest")
	proto.RegisterType((*HealthResponse)(nil), "debug.HealthResponse")
	proto.RegisterType((*StatsRequest)(nil), "debug.StatsRequest")
	proto.RegisterType((*StatsResponse)(nil), "debug.StatsResponse")
	proto.RegisterType((*LogRequest)(nil), "debug.LogRequest")
	proto.RegisterType((*LogResponse)(nil), "debug.LogResponse")
	proto.RegisterType((*Record)(nil), "debug.Record")
	proto.RegisterMapType((map[string]string)(nil), "debug.Record.MetadataEntry")
	proto.RegisterType((*TraceRequest)(nil), "debug.TraceRequest")
	proto.RegisterType((*TraceResponse)(nil), "debug.TraceResponse")
	proto.RegisterType((*Span)(nil), "debug.Span")
	proto.RegisterMapType((map[string]string)(nil), "debug.Span.MetadataEntry")
}

func init() { proto.RegisterFile("debug/debug.proto", fileDescriptor_5ae24eab94cb53d5) }

var fileDescriptor_5ae24eab94cb53d5 = []byte{
	// 623 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0xdd, 0x6e, 0xd3, 0x4c,
	0x10, 0xad, 0xed, 0x38, 0x71, 0x26, 0x4d, 0xda, 0xee, 0xd7, 0x0f, 0x19, 0x83, 0x10, 0x18, 0x21,
	0x2a, 0x10, 0xa9, 0x94, 0x02, 0xad, 0xe8, 0x5d, 0x55, 0x24, 0x90, 0x4a, 0x2b, 0x6d, 0xdb, 0x1b,
	0xee, 0x36, 0xf6, 0xca, 0xb1, 0xa8, 0xbd, 0x66, 0x77, 0x5d, 0x29, 0x4f, 0xc4, 0x1b, 0xf0, 0x26,
	0xdc, 0xf1, 0x30, 0x68, 0x7f, 0xec, 0xda, 0x20, 0xae, 0xb8, 0xb1, 0xf6, 0x9c, 0xf1, 0x99, 0x9d,
	0x39, 0x3b, 0x1a, 0xd8, 0x49, 0xe9, 0xb2, 0xce, 0xf6, 0xf5, 0x77, 0x5e, 0x71, 0x26, 0x19, 0xf2,
	0x35, 0x88, 0xb7, 0x60, 0xfa, 0x81, 0x92, 0x1b, 0xb9, 0xc2, 0xf4, 0x6b, 0x4d, 0x85, 0x8c, 0xf7,
	0x60, 0xd6, 0x10, 0xa2, 0x62, 0xa5, 0xa0, 0xe8, 0x1e, 0x0c, 0x85, 0x24, 0xb2, 0x16, 0xa1, 0xf3,
	0xd8, 0xd9, 0x1b, 0x63, 0x8b, 0xe2, 0x19, 0x6c, 0x5e, 0x4a, 0x22, 0x45, 0xa3, 0xfc, 0xe1, 0xc0,
	0xd4, 0x12, 0x56, 0xf9, 0x10, 0xc6, 0x32, 0x2f, 0xa8, 0x90, 0xa4, 0xa8, 0xb4, 0x78, 0x80, 0xef,
	0x08, 0x14, 0xc2, 0x48, 0x48, 0xc2, 0x25, 0x4d, 0x43, 0x57, 0xc7, 0x1a, 0xa8, 0x6e, 0xac, 0x2b,
	0xf5, 0x63, 0xe8, 0xe9, 0x80, 0x45, 0x8a, 0x2f, 0x68, 0xc1, 0xf8, 0x3a, 0x1c, 0x18, 0xde, 0x20,
	0x95, 0x49, 0xae, 0x38, 0x25, 0xa9, 0x08, 0x7d, 0x93, 0xc9, 0x42, 0x34, 0x03, 0x37, 0x4b, 0xc2,
	0xa1, 0x26, 0xdd, 0x2c, 0x41, 0x11, 0x04, 0xdc, 0x94, 0x2b, 0xc2, 0x91, 0x66, 0x5b, 0xac, 0xb2,
	0x53, 0xce, 0x19, 0x17, 0x61, 0x60, 0xb2, 0x1b, 0x14, 0x1f, 0x01, 0x9c, 0xb1, 0xcc, 0x76, 0x89,
	0x76, 0xc1, 0x4f, 0x58, 0x5d, 0x4a, 0xdd, 0x8f, 0x87, 0x0d, 0x50, 0xac, 0xc8, 0xcb, 0x84, 0xea,
	0x4e, 0x3c, 0x6c, 0x40, 0xfc, 0x16, 0x26, 0x5a, 0x69, 0xed, 0x78, 0x0e, 0x23, 0x4e, 0x13, 0xc6,
	0x53, 0xe5, 0xa4, 0xb7, 0x37, 0x59, 0x4c, 0xe7, 0xe6, 0x45, 0xb0, 0x66, 0x71, 0x13, 0x8d, 0xbf,
	0x3b, 0x30, 0x34, 0xdc, 0x9f, 0x16, 0x7a, 0x5d, 0x0b, 0x0f, 0x21, 0x28, 0xa8, 0x24, 0x29, 0x91,
	0x24, 0x74, 0x75, 0xca, 0x07, 0xbd, 0x94, 0xf3, 0x4f, 0x36, 0xfa, 0xbe, 0x94, 0x7c, 0x8d, 0xdb,
	0x9f, 0x95, 0x63, 0x05, 0x15, 0x82, 0x64, 0xc6, 0xe2, 0x31, 0x6e, 0x60, 0x74, 0x0c, 0xd3, 0x9e,
	0x08, 0x6d, 0x83, 0xf7, 0x85, 0xae, 0xed, 0xdb, 0xab, 0xa3, 0x6a, 0xf6, 0x96, 0xdc, 0xd4, 0xa6,
	0xd9, 0x31, 0x36, 0xe0, 0x9d, 0x7b, 0xe4, 0xc4, 0x8f, 0x60, 0xf3, 0x8a, 0x93, 0x84, 0x36, 0x66,
	0xcd, 0xc0, 0xcd, 0x53, 0x2b, 0x75, 0xf3, 0x34, 0x5e, 0xc0, 0xd4, 0xc6, 0xad, 0x25, 0x4f, 0xc0,
	0x17, 0x15, 0x29, 0x1b, 0x43, 0x26, 0xb6, 0xfa, 0xcb, 0x8a, 0x94, 0xd8, 0x44, 0xe2, 0x6f, 0x2e,
	0x0c, 0x14, 0x56, 0xd7, 0x4a, 0x25, 0xb6, 0xf9, 0x0c, 0xb0, 0x57, 0xb8, 0xcd, 0x15, 0xea, 0x15,
	0x2b, 0xc2, 0x69, 0x29, 0x6d, 0x63, 0x16, 0x21, 0x04, 0x83, 0x92, 0x14, 0x54, 0x4f, 0xce, 0x18,
	0xeb, 0x73, 0x77, 0x02, 0xfd, 0xfe, 0x04, 0x46, 0x10, 0xa4, 0x35, 0x27, 0x32, 0x67, 0xa5, 0x9d,
	0x9e, 0x16, 0xa3, 0x37, 0x1d, 0xd3, 0x47, 0xba, 0xec, 0xfb, 0x9d, 0xb2, 0xff, 0x6a, 0xf9, 0x53,
	0x18, 0xc8, 0x75, 0x45, 0xf5, 0x70, 0xcd, 0x16, 0x5b, 0x1d, 0xc9, 0xd5, 0xba, 0xa2, 0x58, 0x07,
	0xff, 0xc9, 0xfd, 0x17, 0xcf, 0x20, 0x68, 0xd2, 0xa1, 0x09, 0x8c, 0x3e, 0x9e, 0x9f, 0x5c, 0x5c,
	0x9f, 0x9f, 0x6e, 0x6f, 0xa0, 0x4d, 0x08, 0x2e, 0xae, 0xaf, 0x0c, 0x72, 0x16, 0x3f, 0x1d, 0xf0,
	0x4f, 0xd5, 0xe5, 0x68, 0x0e, 0xde, 0x19, 0xcb, 0xd0, 0x8e, 0xad, 0xe5, 0x6e, 0xca, 0x23, 0xd4,
	0xa5, 0xcc, 0x5b, 0xc5, 0x1b, 0xe8, 0x10, 0x86, 0x66, 0x37, 0xa0, 0x5d, 0x1b, 0xef, 0xed, 0x8e,
	0xe8, 0xff, 0xdf, 0xd8, 0x56, 0xf8, 0x1a, 0x7c, 0xbd, 0x19, 0xd0, 0x7f, 0x4d, 0xdb, 0x9d, 0xc5,
	0x11, 0xed, 0xf6, 0xc9, 0xae, 0x4a, 0x4f, 0x4b, 0xab, 0xea, 0xce, 0x56, 0xab, 0xea, 0x0d, 0x54,
	0xbc, 0x71, 0xf2, 0xea, 0xf3, 0xcb, 0x2c, 0x97, 0xab, 0x7a, 0x39, 0x4f, 0x58, 0xb1, 0x5f, 0xe4,
	0x09, 0x67, 0xf6, 0x7b, 0x7b, 0x60, 0x36, 0xe0, 0xbe, 0xde, 0x80, 0xc7, 0xfa, 0xbc, 0x1c, 0x6a,
	0x70, 0xf0, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x59, 0x82, 0x51, 0x15, 0x23, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DebugClient is the client API for Debug service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DebugClient interface {
	Log(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (*LogResponse, error)
	Health(ctx context.Context, in *HealthRequest, opts ...grpc.CallOption) (*HealthResponse, error)
	Stats(ctx context.Context, in *StatsRequest, opts ...grpc.CallOption) (*StatsResponse, error)
	Trace(ctx context.Context, in *TraceRequest, opts ...grpc.CallOption) (*TraceResponse, error)
}

type debugClient struct {
	cc *grpc.ClientConn
}

func NewDebugClient(cc *grpc.ClientConn) DebugClient {
	return &debugClient{cc}
}

func (c *debugClient) Log(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (*LogResponse, error) {
	out := new(LogResponse)
	err := c.cc.Invoke(ctx, "/debug.Debug/Log", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *debugClient) Health(ctx context.Context, in *HealthRequest, opts ...grpc.CallOption) (*HealthResponse, error) {
	out := new(HealthResponse)
	err := c.cc.Invoke(ctx, "/debug.Debug/Health", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *debugClient) Stats(ctx context.Context, in *StatsRequest, opts ...grpc.CallOption) (*StatsResponse, error) {
	out := new(StatsResponse)
	err := c.cc.Invoke(ctx, "/debug.Debug/Stats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *debugClient) Trace(ctx context.Context, in *TraceRequest, opts ...grpc.CallOption) (*TraceResponse, error) {
	out := new(TraceResponse)
	err := c.cc.Invoke(ctx, "/debug.Debug/Trace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DebugServer is the server API for Debug service.
type DebugServer interface {
	Log(context.Context, *LogRequest) (*LogResponse, error)
	Health(context.Context, *HealthRequest) (*HealthResponse, error)
	Stats(context.Context, *StatsRequest) (*StatsResponse, error)
	Trace(context.Context, *TraceRequest) (*TraceResponse, error)
}

func RegisterDebugServer(s *grpc.Server, srv DebugServer) {
	s.RegisterService(&_Debug_serviceDesc, srv)
}

func _Debug_Log_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DebugServer).Log(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/debug.Debug/Log",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DebugServer).Log(ctx, req.(*LogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Debug_Health_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DebugServer).Health(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/debug.Debug/Health",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DebugServer).Health(ctx, req.(*HealthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Debug_Stats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DebugServer).Stats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/debug.Debug/Stats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DebugServer).Stats(ctx, req.(*StatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Debug_Trace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TraceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DebugServer).Trace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/debug.Debug/Trace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DebugServer).Trace(ctx, req.(*TraceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Debug_serviceDesc = grpc.ServiceDesc{
	ServiceName: "debug.Debug",
	HandlerType: (*DebugServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Log",
			Handler:    _Debug_Log_Handler,
		},
		{
			MethodName: "Health",
			Handler:    _Debug_Health_Handler,
		},
		{
			MethodName: "Stats",
			Handler:    _Debug_Stats_Handler,
		},
		{
			MethodName: "Trace",
			Handler:    _Debug_Trace_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "debug/debug.proto",
}
