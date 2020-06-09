// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/service.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type BenchmarkRequest_Command int32

const (
	BenchmarkRequest_start BenchmarkRequest_Command = 0
	BenchmarkRequest_stop  BenchmarkRequest_Command = 1
	BenchmarkRequest_stats BenchmarkRequest_Command = 2
)

var BenchmarkRequest_Command_name = map[int32]string{
	0: "start",
	1: "stop",
	2: "stats",
}

var BenchmarkRequest_Command_value = map[string]int32{
	"start": 0,
	"stop":  1,
	"stats": 2,
}

func (x BenchmarkRequest_Command) String() string {
	return proto.EnumName(BenchmarkRequest_Command_name, int32(x))
}

func (BenchmarkRequest_Command) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6ff5ab49d8a5fcc4, []int{0, 0}
}

type Event_EventType int32

const (
	Event_Creat  Event_EventType = 0
	Event_Update Event_EventType = 1
	Event_Delete Event_EventType = 2
)

var Event_EventType_name = map[int32]string{
	0: "Creat",
	1: "Update",
	2: "Delete",
}

var Event_EventType_value = map[string]int32{
	"Creat":  0,
	"Update": 1,
	"Delete": 2,
}

func (x Event_EventType) String() string {
	return proto.EnumName(Event_EventType_name, int32(x))
}

func (Event_EventType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6ff5ab49d8a5fcc4, []int{7, 0}
}

type BenchmarkRequest struct {
	Command              BenchmarkRequest_Command `protobuf:"varint,1,opt,name=command,proto3,enum=pb.BenchmarkRequest_Command" json:"command,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *BenchmarkRequest) Reset()         { *m = BenchmarkRequest{} }
func (m *BenchmarkRequest) String() string { return proto.CompactTextString(m) }
func (*BenchmarkRequest) ProtoMessage()    {}
func (*BenchmarkRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ff5ab49d8a5fcc4, []int{0}
}

func (m *BenchmarkRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BenchmarkRequest.Unmarshal(m, b)
}
func (m *BenchmarkRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BenchmarkRequest.Marshal(b, m, deterministic)
}
func (m *BenchmarkRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BenchmarkRequest.Merge(m, src)
}
func (m *BenchmarkRequest) XXX_Size() int {
	return xxx_messageInfo_BenchmarkRequest.Size(m)
}
func (m *BenchmarkRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BenchmarkRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BenchmarkRequest proto.InternalMessageInfo

func (m *BenchmarkRequest) GetCommand() BenchmarkRequest_Command {
	if m != nil {
		return m.Command
	}
	return BenchmarkRequest_start
}

type BenchmarkResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BenchmarkResponse) Reset()         { *m = BenchmarkResponse{} }
func (m *BenchmarkResponse) String() string { return proto.CompactTextString(m) }
func (*BenchmarkResponse) ProtoMessage()    {}
func (*BenchmarkResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ff5ab49d8a5fcc4, []int{1}
}

func (m *BenchmarkResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BenchmarkResponse.Unmarshal(m, b)
}
func (m *BenchmarkResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BenchmarkResponse.Marshal(b, m, deterministic)
}
func (m *BenchmarkResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BenchmarkResponse.Merge(m, src)
}
func (m *BenchmarkResponse) XXX_Size() int {
	return xxx_messageInfo_BenchmarkResponse.Size(m)
}
func (m *BenchmarkResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BenchmarkResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BenchmarkResponse proto.InternalMessageInfo

type PutRequest struct {
	Key                  string   `protobuf:"bytes,1,opt,name=Key,proto3" json:"Key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PutRequest) Reset()         { *m = PutRequest{} }
func (m *PutRequest) String() string { return proto.CompactTextString(m) }
func (*PutRequest) ProtoMessage()    {}
func (*PutRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ff5ab49d8a5fcc4, []int{2}
}

func (m *PutRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PutRequest.Unmarshal(m, b)
}
func (m *PutRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PutRequest.Marshal(b, m, deterministic)
}
func (m *PutRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutRequest.Merge(m, src)
}
func (m *PutRequest) XXX_Size() int {
	return xxx_messageInfo_PutRequest.Size(m)
}
func (m *PutRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PutRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PutRequest proto.InternalMessageInfo

func (m *PutRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *PutRequest) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type PutResponse struct {
	Update               bool     `protobuf:"varint,1,opt,name=update,proto3" json:"update,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PutResponse) Reset()         { *m = PutResponse{} }
func (m *PutResponse) String() string { return proto.CompactTextString(m) }
func (*PutResponse) ProtoMessage()    {}
func (*PutResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ff5ab49d8a5fcc4, []int{3}
}

func (m *PutResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PutResponse.Unmarshal(m, b)
}
func (m *PutResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PutResponse.Marshal(b, m, deterministic)
}
func (m *PutResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutResponse.Merge(m, src)
}
func (m *PutResponse) XXX_Size() int {
	return xxx_messageInfo_PutResponse.Size(m)
}
func (m *PutResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PutResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PutResponse proto.InternalMessageInfo

func (m *PutResponse) GetUpdate() bool {
	if m != nil {
		return m.Update
	}
	return false
}

type DeleteRequest struct {
	Key                  string   `protobuf:"bytes,1,opt,name=Key,proto3" json:"Key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ff5ab49d8a5fcc4, []int{4}
}

func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(m, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type DeleteResponse struct {
	Removed              bool     `protobuf:"varint,1,opt,name=removed,proto3" json:"removed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ff5ab49d8a5fcc4, []int{5}
}

func (m *DeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResponse.Unmarshal(m, b)
}
func (m *DeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResponse.Marshal(b, m, deterministic)
}
func (m *DeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResponse.Merge(m, src)
}
func (m *DeleteResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteResponse.Size(m)
}
func (m *DeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResponse proto.InternalMessageInfo

func (m *DeleteResponse) GetRemoved() bool {
	if m != nil {
		return m.Removed
	}
	return false
}

type Filter struct {
	Create               bool     `protobuf:"varint,1,opt,name=create,proto3" json:"create,omitempty"`
	Update               bool     `protobuf:"varint,2,opt,name=update,proto3" json:"update,omitempty"`
	Delete               bool     `protobuf:"varint,3,opt,name=delete,proto3" json:"delete,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Filter) Reset()         { *m = Filter{} }
func (m *Filter) String() string { return proto.CompactTextString(m) }
func (*Filter) ProtoMessage()    {}
func (*Filter) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ff5ab49d8a5fcc4, []int{6}
}

func (m *Filter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Filter.Unmarshal(m, b)
}
func (m *Filter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Filter.Marshal(b, m, deterministic)
}
func (m *Filter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Filter.Merge(m, src)
}
func (m *Filter) XXX_Size() int {
	return xxx_messageInfo_Filter.Size(m)
}
func (m *Filter) XXX_DiscardUnknown() {
	xxx_messageInfo_Filter.DiscardUnknown(m)
}

var xxx_messageInfo_Filter proto.InternalMessageInfo

func (m *Filter) GetCreate() bool {
	if m != nil {
		return m.Create
	}
	return false
}

func (m *Filter) GetUpdate() bool {
	if m != nil {
		return m.Update
	}
	return false
}

func (m *Filter) GetDelete() bool {
	if m != nil {
		return m.Delete
	}
	return false
}

type Event struct {
	Type                 Event_EventType `protobuf:"varint,1,opt,name=type,proto3,enum=pb.Event_EventType" json:"type,omitempty"`
	Key                  string          `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Value                string          `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	OldValue             string          `protobuf:"bytes,4,opt,name=old_value,json=oldValue,proto3" json:"old_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ff5ab49d8a5fcc4, []int{7}
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

func (m *Event) GetType() Event_EventType {
	if m != nil {
		return m.Type
	}
	return Event_Creat
}

func (m *Event) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Event) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *Event) GetOldValue() string {
	if m != nil {
		return m.OldValue
	}
	return ""
}

func init() {
	proto.RegisterEnum("pb.BenchmarkRequest_Command", BenchmarkRequest_Command_name, BenchmarkRequest_Command_value)
	proto.RegisterEnum("pb.Event_EventType", Event_EventType_name, Event_EventType_value)
	proto.RegisterType((*BenchmarkRequest)(nil), "pb.BenchmarkRequest")
	proto.RegisterType((*BenchmarkResponse)(nil), "pb.BenchmarkResponse")
	proto.RegisterType((*PutRequest)(nil), "pb.PutRequest")
	proto.RegisterType((*PutResponse)(nil), "pb.PutResponse")
	proto.RegisterType((*DeleteRequest)(nil), "pb.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "pb.DeleteResponse")
	proto.RegisterType((*Filter)(nil), "pb.Filter")
	proto.RegisterType((*Event)(nil), "pb.Event")
}

func init() {
	proto.RegisterFile("pb/service.proto", fileDescriptor_6ff5ab49d8a5fcc4)
}

var fileDescriptor_6ff5ab49d8a5fcc4 = []byte{
	// 430 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x53, 0xcb, 0x6e, 0xd4, 0x40,
	0x10, 0x5c, 0x7b, 0x9f, 0x6e, 0xc4, 0x32, 0xe9, 0x04, 0xb4, 0x32, 0x1c, 0x60, 0x24, 0x44, 0xe0,
	0x60, 0x20, 0x20, 0xae, 0x48, 0x04, 0xb8, 0xe4, 0xb2, 0x32, 0x10, 0x8e, 0xc8, 0x8f, 0x96, 0x12,
	0xc5, 0xf6, 0x0c, 0xf6, 0xd8, 0x92, 0x3f, 0x88, 0xcf, 0xe1, 0x9f, 0xd0, 0x3c, 0xec, 0x98, 0x15,
	0x5c, 0x56, 0xd3, 0xd5, 0x35, 0xd5, 0xb5, 0xd5, 0x63, 0x60, 0x32, 0x7d, 0xd9, 0x50, 0xdd, 0x5d,
	0x67, 0x14, 0xc9, 0x5a, 0x28, 0x81, 0xbe, 0x4c, 0x79, 0x0b, 0xec, 0x03, 0x55, 0xd9, 0x55, 0x99,
	0xd4, 0x37, 0x31, 0xfd, 0x6c, 0xa9, 0x51, 0xf8, 0x0e, 0xd6, 0x99, 0x28, 0xcb, 0xa4, 0xca, 0x77,
	0xde, 0x63, 0xef, 0x74, 0x7b, 0xf6, 0x28, 0x92, 0x69, 0x74, 0x48, 0x8b, 0xce, 0x2d, 0x27, 0x1e,
	0xc8, 0xfc, 0x39, 0xac, 0x1d, 0x86, 0x01, 0x2c, 0x1b, 0x95, 0xd4, 0x8a, 0xcd, 0x70, 0x03, 0x8b,
	0x46, 0x09, 0xc9, 0x3c, 0x07, 0xaa, 0x86, 0xf9, 0xfc, 0x18, 0x8e, 0x26, 0x7a, 0x8d, 0x14, 0x55,
	0x43, 0xfc, 0x2d, 0xc0, 0xbe, 0x55, 0x83, 0x0b, 0x06, 0xf3, 0x0b, 0xea, 0x8d, 0x83, 0x20, 0xd6,
	0x47, 0x3c, 0x81, 0x65, 0x97, 0x14, 0x2d, 0xed, 0x7c, 0x83, 0xd9, 0x82, 0x3f, 0x85, 0x3b, 0xe6,
	0x96, 0x15, 0xc1, 0x07, 0xb0, 0x6a, 0x65, 0x9e, 0x28, 0x32, 0x37, 0x37, 0xb1, 0xab, 0xf8, 0x13,
	0xb8, 0xfb, 0x91, 0x0a, 0x52, 0xf4, 0x5f, 0x7d, 0xfe, 0x02, 0xb6, 0x03, 0xc5, 0x89, 0xed, 0x60,
	0x5d, 0x53, 0x29, 0x3a, 0xca, 0x9d, 0xda, 0x50, 0xf2, 0x3d, 0xac, 0x3e, 0x5f, 0x17, 0x8a, 0x6a,
	0x3d, 0x30, 0xab, 0x69, 0x32, 0xd0, 0x56, 0x13, 0x23, 0xfe, 0xd4, 0x88, 0xc6, 0x73, 0x33, 0x65,
	0x37, 0xb7, 0xb8, 0xad, 0xf8, 0x2f, 0x0f, 0x96, 0x9f, 0x3a, 0xaa, 0x14, 0x3e, 0x83, 0x85, 0xea,
	0x25, 0xb9, 0xf0, 0x8f, 0x75, 0xf8, 0xa6, 0x61, 0x7f, 0xbf, 0xf6, 0x92, 0x62, 0x43, 0xd0, 0x7f,
	0xe1, 0x86, 0x7a, 0x17, 0x87, 0x3e, 0xde, 0x46, 0x34, 0x9f, 0x44, 0x84, 0x0f, 0x21, 0x10, 0x45,
	0xfe, 0xc3, 0x76, 0x16, 0xa6, 0xb3, 0x11, 0x45, 0x7e, 0x69, 0xf2, 0x8b, 0x20, 0x18, 0x75, 0xf5,
	0x8a, 0xce, 0xb5, 0x7d, 0x36, 0x43, 0x80, 0xd5, 0x37, 0xe3, 0x98, 0x79, 0xfa, 0x6c, 0x93, 0x61,
	0xfe, 0xd9, 0x6f, 0x0f, 0x82, 0x8b, 0xcb, 0x2f, 0xf6, 0x25, 0xe1, 0x29, 0xcc, 0xf7, 0xad, 0xc2,
	0xad, 0x36, 0x79, 0xbb, 0xbc, 0xf0, 0xde, 0x58, 0xbb, 0xdd, 0xce, 0xf0, 0xf5, 0xa0, 0x81, 0x47,
	0xba, 0xf9, 0xd7, 0x32, 0x42, 0x9c, 0x42, 0xe3, 0x15, 0x0e, 0xcb, 0xef, 0x89, 0xca, 0xae, 0x10,
	0x74, 0xdb, 0xe6, 0x1d, 0x06, 0x63, 0x1e, 0x7c, 0xf6, 0xca, 0xc3, 0xf7, 0xb0, 0x35, 0x9c, 0xf1,
	0x39, 0xe1, 0xc9, 0xbf, 0x5e, 0x6b, 0x78, 0xff, 0x00, 0x1d, 0x86, 0xa4, 0x2b, 0xf3, 0x31, 0xbc,
	0xf9, 0x13, 0x00, 0x00, 0xff, 0xff, 0xd2, 0xee, 0x73, 0xf6, 0x20, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// KVServiceClient is the client API for KVService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KVServiceClient interface {
	Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*PutResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	Watch(ctx context.Context, in *Filter, opts ...grpc.CallOption) (KVService_WatchClient, error)
	// for testing
	WatchBenchmark(ctx context.Context, in *BenchmarkRequest, opts ...grpc.CallOption) (*BenchmarkResponse, error)
}

type kVServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewKVServiceClient(cc grpc.ClientConnInterface) KVServiceClient {
	return &kVServiceClient{cc}
}

func (c *kVServiceClient) Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*PutResponse, error) {
	out := new(PutResponse)
	err := c.cc.Invoke(ctx, "/pb.KVService/Put", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVServiceClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/pb.KVService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVServiceClient) Watch(ctx context.Context, in *Filter, opts ...grpc.CallOption) (KVService_WatchClient, error) {
	stream, err := c.cc.NewStream(ctx, &_KVService_serviceDesc.Streams[0], "/pb.KVService/Watch", opts...)
	if err != nil {
		return nil, err
	}
	x := &kVServiceWatchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type KVService_WatchClient interface {
	Recv() (*Event, error)
	grpc.ClientStream
}

type kVServiceWatchClient struct {
	grpc.ClientStream
}

func (x *kVServiceWatchClient) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *kVServiceClient) WatchBenchmark(ctx context.Context, in *BenchmarkRequest, opts ...grpc.CallOption) (*BenchmarkResponse, error) {
	out := new(BenchmarkResponse)
	err := c.cc.Invoke(ctx, "/pb.KVService/WatchBenchmark", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KVServiceServer is the server API for KVService service.
type KVServiceServer interface {
	Put(context.Context, *PutRequest) (*PutResponse, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	Watch(*Filter, KVService_WatchServer) error
	// for testing
	WatchBenchmark(context.Context, *BenchmarkRequest) (*BenchmarkResponse, error)
}

// UnimplementedKVServiceServer can be embedded to have forward compatible implementations.
type UnimplementedKVServiceServer struct {
}

func (*UnimplementedKVServiceServer) Put(ctx context.Context, req *PutRequest) (*PutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Put not implemented")
}
func (*UnimplementedKVServiceServer) Delete(ctx context.Context, req *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedKVServiceServer) Watch(req *Filter, srv KVService_WatchServer) error {
	return status.Errorf(codes.Unimplemented, "method Watch not implemented")
}
func (*UnimplementedKVServiceServer) WatchBenchmark(ctx context.Context, req *BenchmarkRequest) (*BenchmarkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WatchBenchmark not implemented")
}

func RegisterKVServiceServer(s *grpc.Server, srv KVServiceServer) {
	s.RegisterService(&_KVService_serviceDesc, srv)
}

func _KVService_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVServiceServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.KVService/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVServiceServer).Put(ctx, req.(*PutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KVService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.KVService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVServiceServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KVService_Watch_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Filter)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(KVServiceServer).Watch(m, &kVServiceWatchServer{stream})
}

type KVService_WatchServer interface {
	Send(*Event) error
	grpc.ServerStream
}

type kVServiceWatchServer struct {
	grpc.ServerStream
}

func (x *kVServiceWatchServer) Send(m *Event) error {
	return x.ServerStream.SendMsg(m)
}

func _KVService_WatchBenchmark_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BenchmarkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVServiceServer).WatchBenchmark(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.KVService/WatchBenchmark",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVServiceServer).WatchBenchmark(ctx, req.(*BenchmarkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _KVService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.KVService",
	HandlerType: (*KVServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Put",
			Handler:    _KVService_Put_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _KVService_Delete_Handler,
		},
		{
			MethodName: "WatchBenchmark",
			Handler:    _KVService_WatchBenchmark_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Watch",
			Handler:       _KVService_Watch_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pb/service.proto",
}
