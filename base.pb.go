// Code generated by protoc-gen-go.
// source: base.proto
// DO NOT EDIT!

/*
Package common is a generated protocol buffer package.

It is generated from these files:
	base.proto

It has these top-level messages:
	Request
	Response
*/
package common

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
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

// Request request object
type Request struct {
	// Path request path
	Path string `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
	// Method request method
	Method string `protobuf:"bytes,2,opt,name=method" json:"method,omitempty"`
	// Headers request header
	Headers map[string]string `protobuf:"bytes,3,rep,name=headers" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Payload request payload
	Payload []byte `protobuf:"bytes,4,opt,name=payload,proto3" json:"payload,omitempty"`
	// Timestamp time stamp
	Timestamp uint64 `protobuf:"varint,5,opt,name=timestamp" json:"timestamp,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Request) GetHeaders() map[string]string {
	if m != nil {
		return m.Headers
	}
	return nil
}

// Response response object
type Response struct {
	// Code response code
	Code int32 `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	// Headers response headers
	Headers map[string]string `protobuf:"bytes,2,rep,name=headers" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Payload response payload
	Payload []byte `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	// Error error message
	Error string `protobuf:"bytes,4,opt,name=error" json:"error,omitempty"`
	// Time elapsed time
	Time uint64 `protobuf:"varint,5,opt,name=time" json:"time,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Response) GetHeaders() map[string]string {
	if m != nil {
		return m.Headers
	}
	return nil
}

func init() {
	proto.RegisterType((*Request)(nil), "common.Request")
	proto.RegisterType((*Response)(nil), "common.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for GenericService service

type GenericServiceClient interface {
	// Get get resource
	Get(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	// Post create resource
	Post(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	// Put modify resource
	Put(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	// Delete delete resource
	Delete(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	// Head check resource
	Head(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	// Options get resource option
	Options(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type genericServiceClient struct {
	c           client.Client
	serviceName string
}

func NewGenericServiceClient(serviceName string, c client.Client) GenericServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "common"
	}
	return &genericServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *genericServiceClient) Get(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "GenericService.Get", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *genericServiceClient) Post(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "GenericService.Post", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *genericServiceClient) Put(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "GenericService.Put", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *genericServiceClient) Delete(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "GenericService.Delete", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *genericServiceClient) Head(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "GenericService.Head", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *genericServiceClient) Options(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "GenericService.Options", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GenericService service

type GenericServiceHandler interface {
	// Get get resource
	Get(context.Context, *Request, *Response) error
	// Post create resource
	Post(context.Context, *Request, *Response) error
	// Put modify resource
	Put(context.Context, *Request, *Response) error
	// Delete delete resource
	Delete(context.Context, *Request, *Response) error
	// Head check resource
	Head(context.Context, *Request, *Response) error
	// Options get resource option
	Options(context.Context, *Request, *Response) error
}

func RegisterGenericServiceHandler(s server.Server, hdlr GenericServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&GenericService{hdlr}, opts...))
}

type GenericService struct {
	GenericServiceHandler
}

func (h *GenericService) Get(ctx context.Context, in *Request, out *Response) error {
	return h.GenericServiceHandler.Get(ctx, in, out)
}

func (h *GenericService) Post(ctx context.Context, in *Request, out *Response) error {
	return h.GenericServiceHandler.Post(ctx, in, out)
}

func (h *GenericService) Put(ctx context.Context, in *Request, out *Response) error {
	return h.GenericServiceHandler.Put(ctx, in, out)
}

func (h *GenericService) Delete(ctx context.Context, in *Request, out *Response) error {
	return h.GenericServiceHandler.Delete(ctx, in, out)
}

func (h *GenericService) Head(ctx context.Context, in *Request, out *Response) error {
	return h.GenericServiceHandler.Head(ctx, in, out)
}

func (h *GenericService) Options(ctx context.Context, in *Request, out *Response) error {
	return h.GenericServiceHandler.Options(ctx, in, out)
}

func init() { proto.RegisterFile("base.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 335 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x92, 0xcd, 0x4a, 0xc3, 0x40,
	0x14, 0x85, 0xcd, 0x4f, 0x13, 0x7b, 0x2d, 0x5a, 0x2e, 0x22, 0xa1, 0x54, 0x28, 0x5d, 0x15, 0xc5,
	0x2c, 0x2a, 0xa8, 0x74, 0xad, 0xd4, 0x9d, 0x65, 0x7c, 0x82, 0x69, 0x72, 0xa1, 0xc1, 0x26, 0x13,
	0x27, 0xd3, 0x42, 0x1f, 0xc7, 0xa7, 0x72, 0xed, 0x9b, 0x38, 0x99, 0xa4, 0x44, 0xbb, 0x31, 0xe0,
	0xee, 0x9e, 0x99, 0x73, 0x4f, 0xce, 0x97, 0x04, 0x60, 0xc9, 0x0b, 0x0a, 0x73, 0x29, 0x94, 0x40,
	0x2f, 0x12, 0x69, 0x2a, 0xb2, 0xf1, 0x97, 0x05, 0x3e, 0xa3, 0xf7, 0x0d, 0x15, 0x0a, 0x11, 0xdc,
	0x9c, 0xab, 0x55, 0x60, 0x8d, 0xac, 0x49, 0x97, 0x99, 0x19, 0x2f, 0xc0, 0x4b, 0x49, 0xad, 0x44,
	0x1c, 0xd8, 0xe6, 0xb4, 0x56, 0x78, 0x07, 0xfe, 0x8a, 0x78, 0x4c, 0xb2, 0x08, 0x9c, 0x91, 0x33,
	0x39, 0x99, 0x0e, 0xc3, 0x2a, 0x31, 0xac, 0xd3, 0xc2, 0xe7, 0xea, 0xfa, 0x29, 0x53, 0x72, 0xc7,
	0xf6, 0x66, 0x0c, 0xc0, 0xcf, 0xf9, 0x6e, 0x2d, 0x78, 0x1c, 0xb8, 0x3a, 0xb0, 0xc7, 0xf6, 0x12,
	0x87, 0xd0, 0x55, 0x49, 0xaa, 0x37, 0x79, 0x9a, 0x07, 0x1d, 0x7d, 0xe7, 0xb2, 0xe6, 0x60, 0x30,
	0x83, 0xde, 0xcf, 0x40, 0xec, 0x83, 0xf3, 0x46, 0xbb, 0xba, 0x6a, 0x39, 0xe2, 0x39, 0x74, 0xb6,
	0x7c, 0xbd, 0xa1, 0xba, 0x68, 0x25, 0x66, 0xf6, 0x83, 0x35, 0xfe, 0xb4, 0xe0, 0x98, 0x51, 0x91,
	0x8b, 0xac, 0xa0, 0x12, 0x32, 0x12, 0x31, 0x99, 0xcd, 0x0e, 0x33, 0x33, 0xde, 0x37, 0x30, 0xb6,
	0x81, 0xb9, 0x6c, 0x60, 0xaa, 0xb5, 0xbf, 0x69, 0x9c, 0xdf, 0x34, 0xba, 0x0d, 0x49, 0x29, 0xa4,
	0xa1, 0xd4, 0x6d, 0x8c, 0x28, 0x1f, 0x5e, 0x22, 0xd5, 0x78, 0x66, 0xfe, 0x0f, 0xd9, 0xf4, 0xc3,
	0x86, 0xd3, 0x39, 0x65, 0x24, 0x93, 0xe8, 0x95, 0xe4, 0x36, 0x89, 0x08, 0xaf, 0xc0, 0x99, 0x93,
	0xc2, 0xb3, 0x83, 0xcf, 0x31, 0xe8, 0x1f, 0x22, 0x8d, 0x8f, 0xf0, 0x1a, 0xdc, 0x85, 0x28, 0x5a,
	0x9a, 0x75, 0xf0, 0x62, 0xd3, 0xd2, 0x7b, 0x03, 0xde, 0x23, 0xad, 0x49, 0x51, 0xeb, 0x1e, 0xe5,
	0x2b, 0x68, 0x67, 0x0e, 0xc1, 0x7f, 0xc9, 0x55, 0xa2, 0x55, 0x2b, 0xff, 0xd2, 0x33, 0x3f, 0xfc,
	0xed, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x15, 0xef, 0x71, 0x9e, 0xfe, 0x02, 0x00, 0x00,
}
