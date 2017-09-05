// Code generated by protoc-gen-go. DO NOT EDIT.
// source: containers.proto

/*
Package containers is a generated protocol buffer package.

It is generated from these files:
	containers.proto

It has these top-level messages:
	Container
	Response
*/
package containers

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type Container struct {
	Id      string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Ip      string `protobuf:"bytes,3,opt,name=ip" json:"ip,omitempty"`
	Network string `protobuf:"bytes,4,opt,name=network" json:"network,omitempty"`
	Service string `protobuf:"bytes,5,opt,name=service" json:"service,omitempty"`
}

func (m *Container) Reset()                    { *m = Container{} }
func (m *Container) String() string            { return proto.CompactTextString(m) }
func (*Container) ProtoMessage()               {}
func (*Container) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Container) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Container) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Container) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *Container) GetNetwork() string {
	if m != nil {
		return m.Network
	}
	return ""
}

func (m *Container) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

type Response struct {
	Success bool `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Response) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterType((*Container)(nil), "containers.Container")
	proto.RegisterType((*Response)(nil), "containers.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ContainerService service

type ContainerServiceClient interface {
	AddNode(ctx context.Context, in *Container, opts ...grpc.CallOption) (*Response, error)
}

type containerServiceClient struct {
	cc *grpc.ClientConn
}

func NewContainerServiceClient(cc *grpc.ClientConn) ContainerServiceClient {
	return &containerServiceClient{cc}
}

func (c *containerServiceClient) AddNode(ctx context.Context, in *Container, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/containers.ContainerService/AddNode", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ContainerService service

type ContainerServiceServer interface {
	AddNode(context.Context, *Container) (*Response, error)
}

func RegisterContainerServiceServer(s *grpc.Server, srv ContainerServiceServer) {
	s.RegisterService(&_ContainerService_serviceDesc, srv)
}

func _ContainerService_AddNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Container)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContainerServiceServer).AddNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/containers.ContainerService/AddNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContainerServiceServer).AddNode(ctx, req.(*Container))
	}
	return interceptor(ctx, in, info, handler)
}

var _ContainerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "containers.ContainerService",
	HandlerType: (*ContainerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddNode",
			Handler:    _ContainerService_AddNode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "containers.proto",
}

func init() { proto.RegisterFile("containers.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 186 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8f, 0xb1, 0x0e, 0x82, 0x30,
	0x10, 0x86, 0x05, 0x51, 0xe0, 0x06, 0x43, 0x1a, 0x4d, 0x1a, 0x27, 0x43, 0x1c, 0x9c, 0x18, 0x74,
	0x71, 0x35, 0xae, 0xc6, 0x01, 0x9f, 0x00, 0xdb, 0x1b, 0x1a, 0x63, 0xdb, 0xb4, 0xa8, 0xaf, 0x6f,
	0x3c, 0x28, 0xb2, 0xf5, 0xff, 0xee, 0x6b, 0xee, 0x7e, 0x28, 0x84, 0xd1, 0x6d, 0xa3, 0x34, 0x3a,
	0x5f, 0x59, 0x67, 0x5a, 0xc3, 0xe0, 0x4f, 0x4a, 0x0f, 0xf9, 0x39, 0x24, 0xb6, 0x80, 0x58, 0x49,
	0x1e, 0x6d, 0xa2, 0x5d, 0x5e, 0xc7, 0x4a, 0x32, 0x06, 0x89, 0x6e, 0x9e, 0xc8, 0x63, 0x22, 0xf4,
	0x26, 0xc7, 0xf2, 0x69, 0xef, 0x58, 0xc6, 0x21, 0xd5, 0xd8, 0x7e, 0x8c, 0x7b, 0xf0, 0x84, 0x60,
	0x88, 0xbf, 0x89, 0x47, 0xf7, 0x56, 0x02, 0xf9, 0xac, 0x9b, 0xf4, 0xb1, 0xdc, 0x42, 0x56, 0xa3,
	0xb7, 0x46, 0x7b, 0x24, 0xeb, 0x25, 0x04, 0x7a, 0x4f, 0x8b, 0xb3, 0x3a, 0xc4, 0xfd, 0x05, 0x8a,
	0xe1, 0xb4, 0x5b, 0xf7, 0x93, 0x1d, 0x21, 0x3d, 0x49, 0x79, 0x35, 0x12, 0xd9, 0xaa, 0x1a, 0x15,
	0x1b, 0xc4, 0xf5, 0x72, 0x8c, 0xc3, 0x96, 0x72, 0x72, 0x9f, 0x53, 0xf7, 0xc3, 0x37, 0x00, 0x00,
	0xff, 0xff, 0xf7, 0xc7, 0xe9, 0x86, 0x0f, 0x01, 0x00, 0x00,
}