// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.11.4
// source: github.com/micro/go-micro/network/mucp/proto/network.proto

package go_micro_network_mucp

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// AdvertType defines the type of advert
type AdvertType int32

const (
	AdvertType_AdvertAnnounce AdvertType = 0
	AdvertType_AdvertUpdate   AdvertType = 1
)

// Enum value maps for AdvertType.
var (
	AdvertType_name = map[int32]string{
		0: "AdvertAnnounce",
		1: "AdvertUpdate",
	}
	AdvertType_value = map[string]int32{
		"AdvertAnnounce": 0,
		"AdvertUpdate":   1,
	}
)

func (x AdvertType) Enum() *AdvertType {
	p := new(AdvertType)
	*p = x
	return p
}

func (x AdvertType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AdvertType) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_micro_go_micro_network_mucp_proto_network_proto_enumTypes[0].Descriptor()
}

func (AdvertType) Type() protoreflect.EnumType {
	return &file_github_com_micro_go_micro_network_mucp_proto_network_proto_enumTypes[0]
}

func (x AdvertType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AdvertType.Descriptor instead.
func (AdvertType) EnumDescriptor() ([]byte, []int) {
	return file_github_com_micro_go_micro_network_mucp_proto_network_proto_rawDescGZIP(), []int{0}
}

// EventType defines the type of event
type EventType int32

const (
	EventType_Create EventType = 0
	EventType_Delete EventType = 1
	EventType_Update EventType = 2
)

// Enum value maps for EventType.
var (
	EventType_name = map[int32]string{
		0: "Create",
		1: "Delete",
		2: "Update",
	}
	EventType_value = map[string]int32{
		"Create": 0,
		"Delete": 1,
		"Update": 2,
	}
)

func (x EventType) Enum() *EventType {
	p := new(EventType)
	*p = x
	return p
}

func (x EventType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EventType) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_micro_go_micro_network_mucp_proto_network_proto_enumTypes[1].Descriptor()
}

func (EventType) Type() protoreflect.EnumType {
	return &file_github_com_micro_go_micro_network_mucp_proto_network_proto_enumTypes[1]
}

func (x EventType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EventType.Descriptor instead.
func (EventType) EnumDescriptor() ([]byte, []int) {
	return file_github_com_micro_go_micro_network_mucp_proto_network_proto_rawDescGZIP(), []int{1}
}

// Advert is router advertsement streamed by Watch
type Advert struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id of the advertising router
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// type of advertisement
	Type AdvertType `protobuf:"varint,2,opt,name=type,proto3,enum=go.micro.network.mucp.AdvertType" json:"type,omitempty"`
	// unix timestamp of the advertisement
	Timestamp int64 `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// TTL of the Advert
	Ttl int64 `protobuf:"varint,4,opt,name=ttl,proto3" json:"ttl,omitempty"`
	// events is a list of advertised events
	Events []*Event `protobuf:"bytes,5,rep,name=events,proto3" json:"events,omitempty"`
}

func (x *Advert) Reset() {
	*x = Advert{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_micro_go_micro_network_mucp_proto_network_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Advert) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Advert) ProtoMessage() {}

func (x *Advert) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_micro_go_micro_network_mucp_proto_network_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Advert.ProtoReflect.Descriptor instead.
func (*Advert) Descriptor() ([]byte, []int) {
	return file_github_com_micro_go_micro_network_mucp_proto_network_proto_rawDescGZIP(), []int{0}
}

func (x *Advert) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Advert) GetType() AdvertType {
	if x != nil {
		return x.Type
	}
	return AdvertType_AdvertAnnounce
}

func (x *Advert) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Advert) GetTtl() int64 {
	if x != nil {
		return x.Ttl
	}
	return 0
}

func (x *Advert) GetEvents() []*Event {
	if x != nil {
		return x.Events
	}
	return nil
}

// Event is routing table event
type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the unique event id
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// type of event
	Type EventType `protobuf:"varint,2,opt,name=type,proto3,enum=go.micro.network.mucp.EventType" json:"type,omitempty"`
	// unix timestamp of event
	Timestamp int64 `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// service route
	Route *Route `protobuf:"bytes,4,opt,name=route,proto3" json:"route,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_micro_go_micro_network_mucp_proto_network_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_micro_go_micro_network_mucp_proto_network_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_github_com_micro_go_micro_network_mucp_proto_network_proto_rawDescGZIP(), []int{1}
}

func (x *Event) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Event) GetType() EventType {
	if x != nil {
		return x.Type
	}
	return EventType_Create
}

func (x *Event) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Event) GetRoute() *Route {
	if x != nil {
		return x.Route
	}
	return nil
}

// Route is a service route
type Route struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// service for the route
	Service string `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	// the address that advertise this route
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	// gateway as the next hop
	Gateway string `protobuf:"bytes,3,opt,name=gateway,proto3" json:"gateway,omitempty"`
	// the network for this destination
	Network string `protobuf:"bytes,4,opt,name=network,proto3" json:"network,omitempty"`
	// router if the router id
	Router string `protobuf:"bytes,5,opt,name=router,proto3" json:"router,omitempty"`
	// the network link
	Link string `protobuf:"bytes,6,opt,name=link,proto3" json:"link,omitempty"`
	// the metric / score of this route
	Metric int64 `protobuf:"varint,7,opt,name=metric,proto3" json:"metric,omitempty"`
	// metadata for the route
	Metadata map[string]string `protobuf:"bytes,8,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Route) Reset() {
	*x = Route{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_micro_go_micro_network_mucp_proto_network_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Route) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Route) ProtoMessage() {}

func (x *Route) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_micro_go_micro_network_mucp_proto_network_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Route.ProtoReflect.Descriptor instead.
func (*Route) Descriptor() ([]byte, []int) {
	return file_github_com_micro_go_micro_network_mucp_proto_network_proto_rawDescGZIP(), []int{2}
}

func (x *Route) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *Route) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Route) GetGateway() string {
	if x != nil {
		return x.Gateway
	}
	return ""
}

func (x *Route) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *Route) GetRouter() string {
	if x != nil {
		return x.Router
	}
	return ""
}

func (x *Route) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

func (x *Route) GetMetric() int64 {
	if x != nil {
		return x.Metric
	}
	return 0
}

func (x *Route) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

// Error tracks network errors
type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count uint32 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Msg   string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_micro_go_micro_network_mucp_proto_network_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_micro_go_micro_network_mucp_proto_network_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_github_com_micro_go_micro_network_mucp_proto_network_proto_rawDescGZIP(), []int{3}
}

func (x *Error) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *Error) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

// Status is node status
type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error *Error `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_micro_go_micro_network_mucp_proto_network_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_micro_go_micro_network_mucp_proto_network_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_github_com_micro_go_micro_network_mucp_proto_network_proto_rawDescGZIP(), []int{4}
}

func (x *Status) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

// Node is network node
type Node struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// node id
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// node address
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	// the network
	Network string `protobuf:"bytes,3,opt,name=network,proto3" json:"network,omitempty"`
	// associated metadata
	Metadata map[string]string `protobuf:"bytes,4,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// node status
	Status *Status `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Node) Reset() {
	*x = Node{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_micro_go_micro_network_mucp_proto_network_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Node) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Node) ProtoMessage() {}

func (x *Node) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_micro_go_micro_network_mucp_proto_network_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Node.ProtoReflect.Descriptor instead.
func (*Node) Descriptor() ([]byte, []int) {
	return file_github_com_micro_go_micro_network_mucp_proto_network_proto_rawDescGZIP(), []int{5}
}

func (x *Node) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Node) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Node) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *Node) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Node) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

// Connect is sent when the node connects to the network
type Connect struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// network mode
	Node *Node `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
}

func (x *Connect) Reset() {
	*x = Connect{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_micro_go_micro_network_mucp_proto_network_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Connect) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Connect) ProtoMessage() {}

func (x *Connect) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_micro_go_micro_network_mucp_proto_network_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Connect.ProtoReflect.Descriptor instead.
func (*Connect) Descriptor() ([]byte, []int) {
	return file_github_com_micro_go_micro_network_mucp_proto_network_proto_rawDescGZIP(), []int{6}
}

func (x *Connect) GetNode() *Node {
	if x != nil {
		return x.Node
	}
	return nil
}

// Close is sent when the node disconnects from the network
type Close struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// network node
	Node *Node `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
}

func (x *Close) Reset() {
	*x = Close{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_micro_go_micro_network_mucp_proto_network_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Close) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Close) ProtoMessage() {}

func (x *Close) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_micro_go_micro_network_mucp_proto_network_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Close.ProtoReflect.Descriptor instead.
func (*Close) Descriptor() ([]byte, []int) {
	return file_github_com_micro_go_micro_network_mucp_proto_network_proto_rawDescGZIP(), []int{7}
}

func (x *Close) GetNode() *Node {
	if x != nil {
		return x.Node
	}
	return nil
}

// Peer is used to advertise node peers
type Peer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// network node
	Node *Node `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	// node peers
	Peers []*Peer `protobuf:"bytes,2,rep,name=peers,proto3" json:"peers,omitempty"`
}

func (x *Peer) Reset() {
	*x = Peer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_micro_go_micro_network_mucp_proto_network_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Peer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Peer) ProtoMessage() {}

func (x *Peer) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_micro_go_micro_network_mucp_proto_network_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Peer.ProtoReflect.Descriptor instead.
func (*Peer) Descriptor() ([]byte, []int) {
	return file_github_com_micro_go_micro_network_mucp_proto_network_proto_rawDescGZIP(), []int{8}
}

func (x *Peer) GetNode() *Node {
	if x != nil {
		return x.Node
	}
	return nil
}

func (x *Peer) GetPeers() []*Peer {
	if x != nil {
		return x.Peers
	}
	return nil
}

// Sync is network sync message
type Sync struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// peer origin
	Peer *Peer `protobuf:"bytes,1,opt,name=peer,proto3" json:"peer,omitempty"`
	// node routes
	Routes []*Route `protobuf:"bytes,2,rep,name=routes,proto3" json:"routes,omitempty"`
}

func (x *Sync) Reset() {
	*x = Sync{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_micro_go_micro_network_mucp_proto_network_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sync) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sync) ProtoMessage() {}

func (x *Sync) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_micro_go_micro_network_mucp_proto_network_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sync.ProtoReflect.Descriptor instead.
func (*Sync) Descriptor() ([]byte, []int) {
	return file_github_com_micro_go_micro_network_mucp_proto_network_proto_rawDescGZIP(), []int{9}
}

func (x *Sync) GetPeer() *Peer {
	if x != nil {
		return x.Peer
	}
	return nil
}

func (x *Sync) GetRoutes() []*Route {
	if x != nil {
		return x.Routes
	}
	return nil
}

var File_github_com_micro_go_micro_network_mucp_proto_network_proto protoreflect.FileDescriptor

var file_github_com_micro_go_micro_network_mucp_proto_network_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x69, 0x63,
	0x72, 0x6f, 0x2f, 0x67, 0x6f, 0x2d, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2f, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x2f, 0x6d, 0x75, 0x63, 0x70, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x67, 0x6f,
	0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x6d,
	0x75, 0x63, 0x70, 0x22, 0xb5, 0x01, 0x0a, 0x06, 0x41, 0x64, 0x76, 0x65, 0x72, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x35,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x67,
	0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e,
	0x6d, 0x75, 0x63, 0x70, 0x2e, 0x41, 0x64, 0x76, 0x65, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x74, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x03, 0x74, 0x74, 0x6c, 0x12, 0x34, 0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f,
	0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x6d, 0x75, 0x63, 0x70, 0x2e, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x52, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x9f, 0x01, 0x0a, 0x05,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x34, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x6d, 0x75, 0x63, 0x70, 0x2e, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x32, 0x0a, 0x05, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69,
	0x63, 0x72, 0x6f, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x6d, 0x75, 0x63, 0x70,
	0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x05, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x22, 0xb8, 0x02,
	0x0a, 0x05, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x6d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x12, 0x46, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f,
	0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x6d, 0x75, 0x63, 0x70, 0x2e, 0x52, 0x6f,
	0x75, 0x74, 0x65, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x2f, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x3c, 0x0a, 0x06, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x32, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x6d, 0x75, 0x63, 0x70, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x85, 0x02, 0x0a, 0x04, 0x4e, 0x6f, 0x64, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x12, 0x45, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72,
	0x6f, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x6d, 0x75, 0x63, 0x70, 0x2e, 0x4e,
	0x6f, 0x64, 0x65, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x35, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x67, 0x6f,
	0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x6d,
	0x75, 0x63, 0x70, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22,
	0x3a, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x2f, 0x0a, 0x04, 0x6e, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69,
	0x63, 0x72, 0x6f, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x6d, 0x75, 0x63, 0x70,
	0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x22, 0x38, 0x0a, 0x05, 0x43,
	0x6c, 0x6f, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x6d, 0x75, 0x63, 0x70, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52,
	0x04, 0x6e, 0x6f, 0x64, 0x65, 0x22, 0x6a, 0x0a, 0x04, 0x50, 0x65, 0x65, 0x72, 0x12, 0x2f, 0x0a,
	0x04, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f,
	0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x6d,
	0x75, 0x63, 0x70, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x12, 0x31,
	0x0a, 0x05, 0x70, 0x65, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x2e, 0x6d, 0x75, 0x63, 0x70, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x52, 0x05, 0x70, 0x65, 0x65, 0x72,
	0x73, 0x22, 0x6d, 0x0a, 0x04, 0x53, 0x79, 0x6e, 0x63, 0x12, 0x2f, 0x0a, 0x04, 0x70, 0x65, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63,
	0x72, 0x6f, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x6d, 0x75, 0x63, 0x70, 0x2e,
	0x50, 0x65, 0x65, 0x72, 0x52, 0x04, 0x70, 0x65, 0x65, 0x72, 0x12, 0x34, 0x0a, 0x06, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x2e,
	0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x6d, 0x75,
	0x63, 0x70, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73,
	0x2a, 0x32, 0x0a, 0x0a, 0x41, 0x64, 0x76, 0x65, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12,
	0x0a, 0x0e, 0x41, 0x64, 0x76, 0x65, 0x72, 0x74, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65,
	0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x41, 0x64, 0x76, 0x65, 0x72, 0x74, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x10, 0x01, 0x2a, 0x2f, 0x0a, 0x09, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x10, 0x00, 0x12, 0x0a, 0x0a,
	0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_micro_go_micro_network_mucp_proto_network_proto_rawDescOnce sync.Once
	file_github_com_micro_go_micro_network_mucp_proto_network_proto_rawDescData = file_github_com_micro_go_micro_network_mucp_proto_network_proto_rawDesc
)

func file_github_com_micro_go_micro_network_mucp_proto_network_proto_rawDescGZIP() []byte {
	file_github_com_micro_go_micro_network_mucp_proto_network_proto_rawDescOnce.Do(func() {
		file_github_com_micro_go_micro_network_mucp_proto_network_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_micro_go_micro_network_mucp_proto_network_proto_rawDescData)
	})
	return file_github_com_micro_go_micro_network_mucp_proto_network_proto_rawDescData
}

var file_github_com_micro_go_micro_network_mucp_proto_network_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_github_com_micro_go_micro_network_mucp_proto_network_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_github_com_micro_go_micro_network_mucp_proto_network_proto_goTypes = []interface{}{
	(AdvertType)(0), // 0: go.micro.network.mucp.AdvertType
	(EventType)(0),  // 1: go.micro.network.mucp.EventType
	(*Advert)(nil),  // 2: go.micro.network.mucp.Advert
	(*Event)(nil),   // 3: go.micro.network.mucp.Event
	(*Route)(nil),   // 4: go.micro.network.mucp.Route
	(*Error)(nil),   // 5: go.micro.network.mucp.Error
	(*Status)(nil),  // 6: go.micro.network.mucp.Status
	(*Node)(nil),    // 7: go.micro.network.mucp.Node
	(*Connect)(nil), // 8: go.micro.network.mucp.Connect
	(*Close)(nil),   // 9: go.micro.network.mucp.Close
	(*Peer)(nil),    // 10: go.micro.network.mucp.Peer
	(*Sync)(nil),    // 11: go.micro.network.mucp.Sync
	nil,             // 12: go.micro.network.mucp.Route.MetadataEntry
	nil,             // 13: go.micro.network.mucp.Node.MetadataEntry
}
var file_github_com_micro_go_micro_network_mucp_proto_network_proto_depIdxs = []int32{
	0,  // 0: go.micro.network.mucp.Advert.type:type_name -> go.micro.network.mucp.AdvertType
	3,  // 1: go.micro.network.mucp.Advert.events:type_name -> go.micro.network.mucp.Event
	1,  // 2: go.micro.network.mucp.Event.type:type_name -> go.micro.network.mucp.EventType
	4,  // 3: go.micro.network.mucp.Event.route:type_name -> go.micro.network.mucp.Route
	12, // 4: go.micro.network.mucp.Route.metadata:type_name -> go.micro.network.mucp.Route.MetadataEntry
	5,  // 5: go.micro.network.mucp.Status.error:type_name -> go.micro.network.mucp.Error
	13, // 6: go.micro.network.mucp.Node.metadata:type_name -> go.micro.network.mucp.Node.MetadataEntry
	6,  // 7: go.micro.network.mucp.Node.status:type_name -> go.micro.network.mucp.Status
	7,  // 8: go.micro.network.mucp.Connect.node:type_name -> go.micro.network.mucp.Node
	7,  // 9: go.micro.network.mucp.Close.node:type_name -> go.micro.network.mucp.Node
	7,  // 10: go.micro.network.mucp.Peer.node:type_name -> go.micro.network.mucp.Node
	10, // 11: go.micro.network.mucp.Peer.peers:type_name -> go.micro.network.mucp.Peer
	10, // 12: go.micro.network.mucp.Sync.peer:type_name -> go.micro.network.mucp.Peer
	4,  // 13: go.micro.network.mucp.Sync.routes:type_name -> go.micro.network.mucp.Route
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_github_com_micro_go_micro_network_mucp_proto_network_proto_init() }
func file_github_com_micro_go_micro_network_mucp_proto_network_proto_init() {
	if File_github_com_micro_go_micro_network_mucp_proto_network_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_micro_go_micro_network_mucp_proto_network_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Advert); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_micro_go_micro_network_mucp_proto_network_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_micro_go_micro_network_mucp_proto_network_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Route); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_micro_go_micro_network_mucp_proto_network_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_micro_go_micro_network_mucp_proto_network_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_micro_go_micro_network_mucp_proto_network_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Node); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_micro_go_micro_network_mucp_proto_network_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Connect); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_micro_go_micro_network_mucp_proto_network_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Close); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_micro_go_micro_network_mucp_proto_network_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Peer); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_micro_go_micro_network_mucp_proto_network_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sync); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_micro_go_micro_network_mucp_proto_network_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_micro_go_micro_network_mucp_proto_network_proto_goTypes,
		DependencyIndexes: file_github_com_micro_go_micro_network_mucp_proto_network_proto_depIdxs,
		EnumInfos:         file_github_com_micro_go_micro_network_mucp_proto_network_proto_enumTypes,
		MessageInfos:      file_github_com_micro_go_micro_network_mucp_proto_network_proto_msgTypes,
	}.Build()
	File_github_com_micro_go_micro_network_mucp_proto_network_proto = out.File
	file_github_com_micro_go_micro_network_mucp_proto_network_proto_rawDesc = nil
	file_github_com_micro_go_micro_network_mucp_proto_network_proto_goTypes = nil
	file_github_com_micro_go_micro_network_mucp_proto_network_proto_depIdxs = nil
}
