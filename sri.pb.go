// Code generated by protoc-gen-go.
// source: sri.proto
// DO NOT EDIT!

/*
Package gateway is a generated protocol buffer package.

It is generated from these files:
	sri.proto

It has these top-level messages:
	SriMessage
	SriResponse
*/
package gateway

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

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SriMessage struct {
	MsgId   int32  `protobuf:"varint,1,opt,name=msgId" json:"msgId,omitempty"`
	TextMsg string `protobuf:"bytes,2,opt,name=textMsg" json:"textMsg,omitempty"`
	When    int64  `protobuf:"varint,3,opt,name=when" json:"when,omitempty"`
	Name    string `protobuf:"bytes,4,opt,name=name" json:"name,omitempty"`
}

func (m *SriMessage) Reset()                    { *m = SriMessage{} }
func (m *SriMessage) String() string            { return proto.CompactTextString(m) }
func (*SriMessage) ProtoMessage()               {}
func (*SriMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type SriResponse struct {
	Ok         bool   `protobuf:"varint,1,opt,name=ok" json:"ok,omitempty"`
	ErrMessage string `protobuf:"bytes,2,opt,name=errMessage" json:"errMessage,omitempty"`
	MsgId      int32  `protobuf:"varint,3,opt,name=msgId" json:"msgId,omitempty"`
}

func (m *SriResponse) Reset()                    { *m = SriResponse{} }
func (m *SriResponse) String() string            { return proto.CompactTextString(m) }
func (*SriResponse) ProtoMessage()               {}
func (*SriResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*SriMessage)(nil), "gateway.SriMessage")
	proto.RegisterType((*SriResponse)(nil), "gateway.SriResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for MessageService service

type MessageServiceClient interface {
	SendMessage(ctx context.Context, in *SriMessage, opts ...grpc.CallOption) (*SriResponse, error)
}

type messageServiceClient struct {
	cc *grpc.ClientConn
}

func NewMessageServiceClient(cc *grpc.ClientConn) MessageServiceClient {
	return &messageServiceClient{cc}
}

func (c *messageServiceClient) SendMessage(ctx context.Context, in *SriMessage, opts ...grpc.CallOption) (*SriResponse, error) {
	out := new(SriResponse)
	err := grpc.Invoke(ctx, "/gateway.MessageService/SendMessage", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MessageService service

type MessageServiceServer interface {
	SendMessage(context.Context, *SriMessage) (*SriResponse, error)
}

func RegisterMessageServiceServer(s *grpc.Server, srv MessageServiceServer) {
	s.RegisterService(&_MessageService_serviceDesc, srv)
}

func _MessageService_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SriMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gateway.MessageService/SendMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).SendMessage(ctx, req.(*SriMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _MessageService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gateway.MessageService",
	HandlerType: (*MessageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMessage",
			Handler:    _MessageService_SendMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sri.proto",
}

func init() { proto.RegisterFile("sri.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 254 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x4c, 0x90, 0x41, 0x4e, 0xc3, 0x30,
	0x10, 0x45, 0x95, 0xa4, 0xa5, 0x74, 0x8a, 0x2a, 0x61, 0x2a, 0x61, 0x55, 0x08, 0x45, 0x59, 0x45,
	0x2c, 0x12, 0x09, 0x76, 0xdc, 0x80, 0x45, 0x17, 0xd8, 0x27, 0x30, 0xcd, 0xc8, 0x58, 0x25, 0x76,
	0xe4, 0xb1, 0x28, 0x6c, 0xb9, 0x02, 0x47, 0xe3, 0x0a, 0x1c, 0x04, 0xe1, 0x24, 0x6a, 0x76, 0x33,
	0x5f, 0xcf, 0x9a, 0xe7, 0x0f, 0x4b, 0xf2, 0xa6, 0xea, 0xbc, 0x0b, 0x8e, 0x2d, 0xb4, 0x0a, 0x78,
	0x54, 0x9f, 0xdb, 0x1b, 0xed, 0x9c, 0x7e, 0xc3, 0x5a, 0x75, 0xa6, 0x56, 0xd6, 0xba, 0xa0, 0x82,
	0x71, 0x96, 0x7a, 0xac, 0x68, 0x00, 0xa4, 0x37, 0x3b, 0x24, 0x52, 0x1a, 0xd9, 0x06, 0xe6, 0x2d,
	0xe9, 0xa7, 0x86, 0x27, 0x79, 0x52, 0xce, 0x45, 0xbf, 0x30, 0x0e, 0x8b, 0x80, 0x1f, 0x61, 0x47,
	0x9a, 0xa7, 0x79, 0x52, 0x2e, 0xc5, 0xb8, 0x32, 0x06, 0xb3, 0xe3, 0x2b, 0x5a, 0x9e, 0xe5, 0x49,
	0x99, 0x89, 0x38, 0xff, 0x67, 0x56, 0xb5, 0xc8, 0x67, 0x11, 0x8d, 0x73, 0x21, 0x61, 0x25, 0xbd,
	0x11, 0x48, 0x9d, 0xb3, 0x84, 0x6c, 0x0d, 0xa9, 0x3b, 0xc4, 0x1b, 0xe7, 0x22, 0x75, 0x07, 0x76,
	0x0b, 0x80, 0xde, 0x0f, 0x12, 0xc3, 0x8d, 0x49, 0x72, 0xd2, 0xca, 0x26, 0x5a, 0xf7, 0x7b, 0x58,
	0x0f, 0x80, 0x44, 0xff, 0x6e, 0xf6, 0xc8, 0x9e, 0x61, 0x25, 0xd1, 0x36, 0xe3, 0xb3, 0xab, 0x6a,
	0xe8, 0xa0, 0x3a, 0x7d, 0x71, 0xbb, 0x99, 0x86, 0xa3, 0x51, 0x71, 0xfd, 0xf5, 0xf3, 0xfb, 0x9d,
	0x5e, 0x16, 0x17, 0x35, 0x79, 0x53, 0xb7, 0x3d, 0xfb, 0x98, 0xdc, 0xbd, 0x9c, 0xc5, 0x9a, 0x1e,
	0xfe, 0x02, 0x00, 0x00, 0xff, 0xff, 0x04, 0x47, 0xa5, 0x99, 0x5a, 0x01, 0x00, 0x00,
}
