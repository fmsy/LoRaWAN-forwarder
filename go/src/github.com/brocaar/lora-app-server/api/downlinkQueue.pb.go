// Code generated by protoc-gen-go.
// source: downlinkQueue.proto
// DO NOT EDIT!

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type EnqueueDownlinkQueueItemRequest struct {
	// Hex encoded DevEUI of the node.
	DevEUI string `protobuf:"bytes,1,opt,name=devEUI" json:"devEUI,omitempty"`
	// Random reference (used on ack notification).
	Reference string `protobuf:"bytes,2,opt,name=reference" json:"reference,omitempty"`
	// Is an ACK required from the node.
	Confirmed bool `protobuf:"varint,3,opt,name=confirmed" json:"confirmed,omitempty"`
	// FPort used (must be >0)
	FPort uint32 `protobuf:"varint,4,opt,name=fPort" json:"fPort,omitempty"`
	// Base64 encoded data.
	Data []byte `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *EnqueueDownlinkQueueItemRequest) Reset()                    { *m = EnqueueDownlinkQueueItemRequest{} }
func (m *EnqueueDownlinkQueueItemRequest) String() string            { return proto.CompactTextString(m) }
func (*EnqueueDownlinkQueueItemRequest) ProtoMessage()               {}
func (*EnqueueDownlinkQueueItemRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *EnqueueDownlinkQueueItemRequest) GetDevEUI() string {
	if m != nil {
		return m.DevEUI
	}
	return ""
}

func (m *EnqueueDownlinkQueueItemRequest) GetReference() string {
	if m != nil {
		return m.Reference
	}
	return ""
}

func (m *EnqueueDownlinkQueueItemRequest) GetConfirmed() bool {
	if m != nil {
		return m.Confirmed
	}
	return false
}

func (m *EnqueueDownlinkQueueItemRequest) GetFPort() uint32 {
	if m != nil {
		return m.FPort
	}
	return 0
}

func (m *EnqueueDownlinkQueueItemRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type EnqueueDownlinkQueueItemResponse struct {
}

func (m *EnqueueDownlinkQueueItemResponse) Reset()         { *m = EnqueueDownlinkQueueItemResponse{} }
func (m *EnqueueDownlinkQueueItemResponse) String() string { return proto.CompactTextString(m) }
func (*EnqueueDownlinkQueueItemResponse) ProtoMessage()    {}
func (*EnqueueDownlinkQueueItemResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor3, []int{1}
}

type DeleteDownlinkQeueueItemRequest struct {
	// Hex encoded DevEUI of the node.
	DevEUI string `protobuf:"bytes,2,opt,name=devEUI" json:"devEUI,omitempty"`
	// ID of the queue item to delete.
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *DeleteDownlinkQeueueItemRequest) Reset()                    { *m = DeleteDownlinkQeueueItemRequest{} }
func (m *DeleteDownlinkQeueueItemRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteDownlinkQeueueItemRequest) ProtoMessage()               {}
func (*DeleteDownlinkQeueueItemRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

func (m *DeleteDownlinkQeueueItemRequest) GetDevEUI() string {
	if m != nil {
		return m.DevEUI
	}
	return ""
}

func (m *DeleteDownlinkQeueueItemRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type DeleteDownlinkQueueItemResponse struct {
}

func (m *DeleteDownlinkQueueItemResponse) Reset()                    { *m = DeleteDownlinkQueueItemResponse{} }
func (m *DeleteDownlinkQueueItemResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteDownlinkQueueItemResponse) ProtoMessage()               {}
func (*DeleteDownlinkQueueItemResponse) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{3} }

type DownlinkQueueItem struct {
	// ID of the queue item.
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	// Hex encoded DevEUI of the node.
	DevEUI string `protobuf:"bytes,2,opt,name=devEUI" json:"devEUI,omitempty"`
	// Random reference (used on ack notification).
	Reference string `protobuf:"bytes,3,opt,name=reference" json:"reference,omitempty"`
	// Is an ACK required from the node.
	Confirmed bool `protobuf:"varint,4,opt,name=confirmed" json:"confirmed,omitempty"`
	// Transmission is pending (waiting for an ack).
	Pending bool `protobuf:"varint,5,opt,name=pending" json:"pending,omitempty"`
	// FPort used (must be >0).
	FPort uint32 `protobuf:"varint,6,opt,name=fPort" json:"fPort,omitempty"`
	// Base64 encoded data.
	Data []byte `protobuf:"bytes,7,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *DownlinkQueueItem) Reset()                    { *m = DownlinkQueueItem{} }
func (m *DownlinkQueueItem) String() string            { return proto.CompactTextString(m) }
func (*DownlinkQueueItem) ProtoMessage()               {}
func (*DownlinkQueueItem) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{4} }

func (m *DownlinkQueueItem) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *DownlinkQueueItem) GetDevEUI() string {
	if m != nil {
		return m.DevEUI
	}
	return ""
}

func (m *DownlinkQueueItem) GetReference() string {
	if m != nil {
		return m.Reference
	}
	return ""
}

func (m *DownlinkQueueItem) GetConfirmed() bool {
	if m != nil {
		return m.Confirmed
	}
	return false
}

func (m *DownlinkQueueItem) GetPending() bool {
	if m != nil {
		return m.Pending
	}
	return false
}

func (m *DownlinkQueueItem) GetFPort() uint32 {
	if m != nil {
		return m.FPort
	}
	return 0
}

func (m *DownlinkQueueItem) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type ListDownlinkQueueItemsRequest struct {
	// Hex encoded DevEUI of the node.
	DevEUI string `protobuf:"bytes,1,opt,name=devEUI" json:"devEUI,omitempty"`
}

func (m *ListDownlinkQueueItemsRequest) Reset()                    { *m = ListDownlinkQueueItemsRequest{} }
func (m *ListDownlinkQueueItemsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListDownlinkQueueItemsRequest) ProtoMessage()               {}
func (*ListDownlinkQueueItemsRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{5} }

func (m *ListDownlinkQueueItemsRequest) GetDevEUI() string {
	if m != nil {
		return m.DevEUI
	}
	return ""
}

type ListDownlinkQueueItemsResponse struct {
	Items []*DownlinkQueueItem `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
}

func (m *ListDownlinkQueueItemsResponse) Reset()                    { *m = ListDownlinkQueueItemsResponse{} }
func (m *ListDownlinkQueueItemsResponse) String() string            { return proto.CompactTextString(m) }
func (*ListDownlinkQueueItemsResponse) ProtoMessage()               {}
func (*ListDownlinkQueueItemsResponse) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{6} }

func (m *ListDownlinkQueueItemsResponse) GetItems() []*DownlinkQueueItem {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
	proto.RegisterType((*EnqueueDownlinkQueueItemRequest)(nil), "api.EnqueueDownlinkQueueItemRequest")
	proto.RegisterType((*EnqueueDownlinkQueueItemResponse)(nil), "api.EnqueueDownlinkQueueItemResponse")
	proto.RegisterType((*DeleteDownlinkQeueueItemRequest)(nil), "api.DeleteDownlinkQeueueItemRequest")
	proto.RegisterType((*DeleteDownlinkQueueItemResponse)(nil), "api.DeleteDownlinkQueueItemResponse")
	proto.RegisterType((*DownlinkQueueItem)(nil), "api.DownlinkQueueItem")
	proto.RegisterType((*ListDownlinkQueueItemsRequest)(nil), "api.ListDownlinkQueueItemsRequest")
	proto.RegisterType((*ListDownlinkQueueItemsResponse)(nil), "api.ListDownlinkQueueItemsResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for DownlinkQueue service

type DownlinkQueueClient interface {
	// Enqueue adds the given item to the queue. When the node operates in
	// Class-C mode, the data will be pushed directly to the network-server.
	Enqueue(ctx context.Context, in *EnqueueDownlinkQueueItemRequest, opts ...grpc.CallOption) (*EnqueueDownlinkQueueItemResponse, error)
	// Delete deletes an item from the queue.
	Delete(ctx context.Context, in *DeleteDownlinkQeueueItemRequest, opts ...grpc.CallOption) (*DeleteDownlinkQueueItemResponse, error)
	// List lists the items in the queue for the given node.
	List(ctx context.Context, in *ListDownlinkQueueItemsRequest, opts ...grpc.CallOption) (*ListDownlinkQueueItemsResponse, error)
}

type downlinkQueueClient struct {
	cc *grpc.ClientConn
}

func NewDownlinkQueueClient(cc *grpc.ClientConn) DownlinkQueueClient {
	return &downlinkQueueClient{cc}
}

func (c *downlinkQueueClient) Enqueue(ctx context.Context, in *EnqueueDownlinkQueueItemRequest, opts ...grpc.CallOption) (*EnqueueDownlinkQueueItemResponse, error) {
	out := new(EnqueueDownlinkQueueItemResponse)
	err := grpc.Invoke(ctx, "/api.DownlinkQueue/Enqueue", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *downlinkQueueClient) Delete(ctx context.Context, in *DeleteDownlinkQeueueItemRequest, opts ...grpc.CallOption) (*DeleteDownlinkQueueItemResponse, error) {
	out := new(DeleteDownlinkQueueItemResponse)
	err := grpc.Invoke(ctx, "/api.DownlinkQueue/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *downlinkQueueClient) List(ctx context.Context, in *ListDownlinkQueueItemsRequest, opts ...grpc.CallOption) (*ListDownlinkQueueItemsResponse, error) {
	out := new(ListDownlinkQueueItemsResponse)
	err := grpc.Invoke(ctx, "/api.DownlinkQueue/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DownlinkQueue service

type DownlinkQueueServer interface {
	// Enqueue adds the given item to the queue. When the node operates in
	// Class-C mode, the data will be pushed directly to the network-server.
	Enqueue(context.Context, *EnqueueDownlinkQueueItemRequest) (*EnqueueDownlinkQueueItemResponse, error)
	// Delete deletes an item from the queue.
	Delete(context.Context, *DeleteDownlinkQeueueItemRequest) (*DeleteDownlinkQueueItemResponse, error)
	// List lists the items in the queue for the given node.
	List(context.Context, *ListDownlinkQueueItemsRequest) (*ListDownlinkQueueItemsResponse, error)
}

func RegisterDownlinkQueueServer(s *grpc.Server, srv DownlinkQueueServer) {
	s.RegisterService(&_DownlinkQueue_serviceDesc, srv)
}

func _DownlinkQueue_Enqueue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EnqueueDownlinkQueueItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DownlinkQueueServer).Enqueue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DownlinkQueue/Enqueue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DownlinkQueueServer).Enqueue(ctx, req.(*EnqueueDownlinkQueueItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DownlinkQueue_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDownlinkQeueueItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DownlinkQueueServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DownlinkQueue/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DownlinkQueueServer).Delete(ctx, req.(*DeleteDownlinkQeueueItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DownlinkQueue_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDownlinkQueueItemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DownlinkQueueServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DownlinkQueue/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DownlinkQueueServer).List(ctx, req.(*ListDownlinkQueueItemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DownlinkQueue_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.DownlinkQueue",
	HandlerType: (*DownlinkQueueServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Enqueue",
			Handler:    _DownlinkQueue_Enqueue_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _DownlinkQueue_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _DownlinkQueue_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "downlinkQueue.proto",
}

func init() { proto.RegisterFile("downlinkQueue.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 437 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0x41, 0x6e, 0xd3, 0x40,
	0x14, 0xd5, 0xd8, 0x8e, 0xd3, 0x7e, 0x28, 0x12, 0x03, 0xaa, 0x8c, 0x29, 0xa9, 0x3b, 0x14, 0x64,
	0x55, 0x28, 0x96, 0xca, 0x02, 0x89, 0x75, 0xbb, 0x88, 0x84, 0x10, 0x58, 0xe2, 0x00, 0xa6, 0xf3,
	0x13, 0x8d, 0x48, 0x66, 0x1c, 0x7b, 0x02, 0x8b, 0x24, 0x1b, 0xae, 0xc0, 0x01, 0x38, 0x08, 0xc7,
	0xe0, 0x0a, 0xdc, 0x03, 0xe4, 0xb1, 0x91, 0x93, 0x38, 0x8e, 0x77, 0x9e, 0xff, 0xdf, 0xff, 0xcf,
	0xef, 0xbd, 0xb1, 0xe1, 0x11, 0x57, 0xdf, 0xe4, 0x54, 0xc8, 0x2f, 0x1f, 0x17, 0xb8, 0xc0, 0x61,
	0x9a, 0x29, 0xad, 0xa8, 0x9d, 0xa4, 0xc2, 0x3f, 0x9b, 0x28, 0x35, 0x99, 0x62, 0x94, 0xa4, 0x22,
	0x4a, 0xa4, 0x54, 0x3a, 0xd1, 0x42, 0xc9, 0xbc, 0x84, 0xb0, 0x9f, 0x04, 0xce, 0x6f, 0xe5, 0xbc,
	0x18, 0xba, 0xd9, 0xdc, 0x30, 0xd2, 0x38, 0x8b, 0x71, 0xbe, 0xc0, 0x5c, 0xd3, 0x53, 0x70, 0x39,
	0x7e, 0xbd, 0xfd, 0x34, 0xf2, 0x48, 0x40, 0xc2, 0xe3, 0xb8, 0x3a, 0xd1, 0x33, 0x38, 0xce, 0x70,
	0x8c, 0x19, 0xca, 0x3b, 0xf4, 0x2c, 0xd3, 0xaa, 0x0b, 0x45, 0xf7, 0x4e, 0xc9, 0xb1, 0xc8, 0x66,
	0xc8, 0x3d, 0x3b, 0x20, 0xe1, 0x51, 0x5c, 0x17, 0xe8, 0x63, 0xe8, 0x8d, 0x3f, 0xa8, 0x4c, 0x7b,
	0x4e, 0x40, 0xc2, 0x93, 0xb8, 0x3c, 0x50, 0x0a, 0x0e, 0x4f, 0x74, 0xe2, 0xf5, 0x02, 0x12, 0xde,
	0x8f, 0xcd, 0x33, 0x63, 0x10, 0xb4, 0xbf, 0x60, 0x9e, 0x2a, 0x99, 0x23, 0x1b, 0xc1, 0xf9, 0x0d,
	0x4e, 0x51, 0xd7, 0x10, 0x6c, 0x17, 0x61, 0x6d, 0x89, 0x78, 0x00, 0x96, 0xe0, 0x46, 0x98, 0x1d,
	0x5b, 0x82, 0xb3, 0x8b, 0xc6, 0xaa, 0x06, 0xdb, 0x2f, 0x02, 0x0f, 0x1b, 0xdd, 0xdd, 0x45, 0xad,
	0x84, 0x5b, 0xae, 0xd9, 0x07, 0x5d, 0x73, 0x76, 0x5d, 0xf3, 0xa0, 0x9f, 0xa2, 0xe4, 0x42, 0x4e,
	0x8c, 0x45, 0x47, 0xf1, 0xff, 0x63, 0xed, 0xa7, 0xbb, 0xcf, 0xcf, 0xfe, 0x86, 0x9f, 0x6f, 0xe0,
	0xd9, 0x3b, 0x91, 0xeb, 0x86, 0x80, 0xbc, 0x23, 0x6e, 0xf6, 0x1e, 0x06, 0x6d, 0x83, 0xa5, 0x31,
	0xf4, 0x15, 0xf4, 0x44, 0x51, 0xf0, 0x48, 0x60, 0x87, 0xf7, 0xae, 0x4f, 0x87, 0x49, 0x2a, 0x86,
	0x4d, 0x1f, 0x4b, 0xd0, 0xf5, 0x5f, 0x0b, 0x4e, 0xb6, 0x9a, 0x74, 0x05, 0xfd, 0x2a, 0x6a, 0x7a,
	0x69, 0x66, 0x3b, 0x6e, 0xa6, 0xff, 0xa2, 0x03, 0x55, 0x05, 0x76, 0xf9, 0xfd, 0xf7, 0x9f, 0x1f,
	0xd6, 0x80, 0x3d, 0x31, 0x1f, 0x81, 0x54, 0x1c, 0xf3, 0x68, 0x59, 0xaa, 0x5a, 0x47, 0x66, 0xf6,
	0x2d, 0xb9, 0xa2, 0x2b, 0x70, 0xcb, 0xe4, 0x2b, 0xf2, 0x8e, 0x1b, 0xe5, 0xef, 0x45, 0x35, 0xb8,
	0x5f, 0x1a, 0xee, 0xe0, 0x6a, 0xd0, 0xca, 0x1d, 0x2d, 0x05, 0x5f, 0xd3, 0x0c, 0x9c, 0xc2, 0x5d,
	0xca, 0xcc, 0xd6, 0x83, 0x09, 0xf9, 0xcf, 0x0f, 0x62, 0x2a, 0xe2, 0x0b, 0x43, 0xfc, 0x94, 0xb6,
	0x8b, 0xfe, 0xec, 0x9a, 0x7f, 0xc0, 0xeb, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x16, 0xa7, 0x90,
	0xc0, 0x3d, 0x04, 0x00, 0x00,
}