// Code generated by protoc-gen-go. DO NOT EDIT.
// source: helloworld.proto

package protobuf

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

// The request message containing the user's name.
type Dubbo3HelloRequest struct {
	Myname               string   `protobuf:"bytes,1,opt,name=myname,proto3" json:"myname,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Dubbo3HelloRequest) Reset()         { *m = Dubbo3HelloRequest{} }
func (m *Dubbo3HelloRequest) String() string { return proto.CompactTextString(m) }
func (*Dubbo3HelloRequest) ProtoMessage()    {}
func (*Dubbo3HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_17b8c58d586b62f2, []int{0}
}

func (m *Dubbo3HelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Dubbo3HelloRequest.Unmarshal(m, b)
}
func (m *Dubbo3HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Dubbo3HelloRequest.Marshal(b, m, deterministic)
}
func (m *Dubbo3HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Dubbo3HelloRequest.Merge(m, src)
}
func (m *Dubbo3HelloRequest) XXX_Size() int {
	return xxx_messageInfo_Dubbo3HelloRequest.Size(m)
}
func (m *Dubbo3HelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_Dubbo3HelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_Dubbo3HelloRequest proto.InternalMessageInfo

func (m *Dubbo3HelloRequest) GetMyname() string {
	if m != nil {
		return m.Myname
	}
	return ""
}

// The response message containing the greetings
type Dubbo3HelloReply struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Dubbo3HelloReply) Reset()         { *m = Dubbo3HelloReply{} }
func (m *Dubbo3HelloReply) String() string { return proto.CompactTextString(m) }
func (*Dubbo3HelloReply) ProtoMessage()    {}
func (*Dubbo3HelloReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_17b8c58d586b62f2, []int{1}
}

func (m *Dubbo3HelloReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Dubbo3HelloReply.Unmarshal(m, b)
}
func (m *Dubbo3HelloReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Dubbo3HelloReply.Marshal(b, m, deterministic)
}
func (m *Dubbo3HelloReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Dubbo3HelloReply.Merge(m, src)
}
func (m *Dubbo3HelloReply) XXX_Size() int {
	return xxx_messageInfo_Dubbo3HelloReply.Size(m)
}
func (m *Dubbo3HelloReply) XXX_DiscardUnknown() {
	xxx_messageInfo_Dubbo3HelloReply.DiscardUnknown(m)
}

var xxx_messageInfo_Dubbo3HelloReply proto.InternalMessageInfo

func (m *Dubbo3HelloReply) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*Dubbo3HelloRequest)(nil), "protobuf.Dubbo3HelloRequest")
	proto.RegisterType((*Dubbo3HelloReply)(nil), "protobuf.Dubbo3HelloReply")
}

func init() { proto.RegisterFile("helloworld.proto", fileDescriptor_17b8c58d586b62f2) }

var fileDescriptor_17b8c58d586b62f2 = []byte{
	// 205 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xc8, 0x48, 0xcd, 0xc9,
	0xc9, 0x2f, 0xcf, 0x2f, 0xca, 0x49, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00, 0x53,
	0x49, 0xa5, 0x69, 0x4a, 0x3a, 0x5c, 0x42, 0x2e, 0xa5, 0x49, 0x49, 0xf9, 0xc6, 0x1e, 0x20, 0x35,
	0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x62, 0x5c, 0x6c, 0xb9, 0x95, 0x79, 0x89, 0xb9,
	0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x50, 0x9e, 0x92, 0x0a, 0x97, 0x00, 0x8a, 0xea,
	0x82, 0x9c, 0x4a, 0x21, 0x01, 0x2e, 0xe6, 0xdc, 0xe2, 0x74, 0xa8, 0x42, 0x10, 0xd3, 0x68, 0x03,
	0x23, 0x17, 0x2f, 0x44, 0x99, 0x7b, 0x51, 0x6a, 0x6a, 0x49, 0x6a, 0x91, 0x50, 0x00, 0x17, 0x1f,
	0x44, 0x20, 0x38, 0xb1, 0x12, 0xac, 0x55, 0x48, 0x46, 0x0f, 0xe6, 0x04, 0x3d, 0x4c, 0xfb, 0xa5,
	0xa4, 0x70, 0xc8, 0x16, 0xe4, 0x54, 0x2a, 0x31, 0x68, 0x30, 0x1a, 0x30, 0x0a, 0xf9, 0x72, 0xf1,
	0xa3, 0x9a, 0x68, 0x44, 0x89, 0x91, 0x4e, 0xda, 0x5c, 0x02, 0xf9, 0x45, 0xe9, 0x7a, 0x89, 0x05,
	0x89, 0xc9, 0x19, 0xa9, 0x7a, 0x29, 0x20, 0x05, 0x4e, 0xfc, 0x60, 0x15, 0xe1, 0xa0, 0x60, 0x0b,
	0x00, 0x69, 0x0d, 0x60, 0x5c, 0xc4, 0xc4, 0xec, 0xe1, 0x13, 0x9e, 0xc4, 0x06, 0x36, 0xc9, 0x18,
	0x10, 0x00, 0x00, 0xff, 0xff, 0x66, 0x6c, 0x45, 0x54, 0x58, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Dubbo3GreeterClient is the grpc-grpc-client API for Dubbo3Greeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type Dubbo3GreeterClient interface {
	// Sends a greeting
	Dubbo3SayHello(ctx context.Context, opts ...grpc.CallOption) (Dubbo3Greeter_Dubbo3SayHelloClient, error)
	Dubbo3SayHello2(ctx context.Context, in *Dubbo3HelloRequest, opts ...grpc.CallOption) (*Dubbo3HelloReply, error)
}

type dubbo3GreeterClient struct {
	cc *grpc.ClientConn
}

func NewDubbo3GreeterClient(cc *grpc.ClientConn) Dubbo3GreeterClient {
	return &dubbo3GreeterClient{cc}
}

func (c *dubbo3GreeterClient) Dubbo3SayHello(ctx context.Context, opts ...grpc.CallOption) (Dubbo3Greeter_Dubbo3SayHelloClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Dubbo3Greeter_serviceDesc.Streams[0], "/protobuf.Dubbo3Greeter/Dubbo3SayHello", opts...)
	if err != nil {
		return nil, err
	}
	x := &dubbo3GreeterDubbo3SayHelloClient{stream}
	return x, nil
}

type Dubbo3Greeter_Dubbo3SayHelloClient interface {
	Send(*Dubbo3HelloRequest) error
	Recv() (*Dubbo3HelloReply, error)
	grpc.ClientStream
}

type dubbo3GreeterDubbo3SayHelloClient struct {
	grpc.ClientStream
}

func (x *dubbo3GreeterDubbo3SayHelloClient) Send(m *Dubbo3HelloRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *dubbo3GreeterDubbo3SayHelloClient) Recv() (*Dubbo3HelloReply, error) {
	m := new(Dubbo3HelloReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dubbo3GreeterClient) Dubbo3SayHello2(ctx context.Context, in *Dubbo3HelloRequest, opts ...grpc.CallOption) (*Dubbo3HelloReply, error) {
	out := new(Dubbo3HelloReply)
	err := c.cc.Invoke(ctx, "/protobuf.Dubbo3Greeter/Dubbo3SayHello2", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Dubbo3GreeterServer is the server API for Dubbo3Greeter service.
type Dubbo3GreeterServer interface {
	// Sends a greeting
	Dubbo3SayHello(Dubbo3Greeter_Dubbo3SayHelloServer) error
	Dubbo3SayHello2(context.Context, *Dubbo3HelloRequest) (*Dubbo3HelloReply, error)
}

// UnimplementedDubbo3GreeterServer can be embedded to have forward compatible implementations.
type UnimplementedDubbo3GreeterServer struct {
}

func (*UnimplementedDubbo3GreeterServer) Dubbo3SayHello(srv Dubbo3Greeter_Dubbo3SayHelloServer) error {
	return status.Errorf(codes.Unimplemented, "method Dubbo3SayHello not implemented")
}
func (*UnimplementedDubbo3GreeterServer) Dubbo3SayHello2(ctx context.Context, req *Dubbo3HelloRequest) (*Dubbo3HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Dubbo3SayHello2 not implemented")
}

func RegisterDubbo3GreeterServer(s *grpc.Server, srv Dubbo3GreeterServer) {
	s.RegisterService(&_Dubbo3Greeter_serviceDesc, srv)
}

func _Dubbo3Greeter_Dubbo3SayHello_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(Dubbo3GreeterServer).Dubbo3SayHello(&dubbo3GreeterDubbo3SayHelloServer{stream})
}

type Dubbo3Greeter_Dubbo3SayHelloServer interface {
	Send(*Dubbo3HelloReply) error
	Recv() (*Dubbo3HelloRequest, error)
	grpc.ServerStream
}

type dubbo3GreeterDubbo3SayHelloServer struct {
	grpc.ServerStream
}

func (x *dubbo3GreeterDubbo3SayHelloServer) Send(m *Dubbo3HelloReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *dubbo3GreeterDubbo3SayHelloServer) Recv() (*Dubbo3HelloRequest, error) {
	m := new(Dubbo3HelloRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Dubbo3Greeter_Dubbo3SayHello2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Dubbo3HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Dubbo3GreeterServer).Dubbo3SayHello2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.Dubbo3Greeter/Dubbo3SayHello2",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Dubbo3GreeterServer).Dubbo3SayHello2(ctx, req.(*Dubbo3HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Dubbo3Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.Dubbo3Greeter",
	HandlerType: (*Dubbo3GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Dubbo3SayHello2",
			Handler:    _Dubbo3Greeter_Dubbo3SayHello2_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Dubbo3SayHello",
			Handler:       _Dubbo3Greeter_Dubbo3SayHello_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "helloworld.proto",
}
