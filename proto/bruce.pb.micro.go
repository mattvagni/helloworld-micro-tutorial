// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/bruce.proto

package bruce

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/micro/v3/service/api"
	client "github.com/micro/micro/v3/service/client"
	server "github.com/micro/micro/v3/service/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Bruce service

func NewBruceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Bruce service

type BruceService interface {
	Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (Bruce_StreamService, error)
	PingPong(ctx context.Context, opts ...client.CallOption) (Bruce_PingPongService, error)
}

type bruceService struct {
	c    client.Client
	name string
}

func NewBruceService(name string, c client.Client) BruceService {
	return &bruceService{
		c:    c,
		name: name,
	}
}

func (c *bruceService) Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Bruce.Call", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bruceService) Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (Bruce_StreamService, error) {
	req := c.c.NewRequest(c.name, "Bruce.Stream", &StreamingRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &bruceServiceStream{stream}, nil
}

type Bruce_StreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*StreamingResponse, error)
}

type bruceServiceStream struct {
	stream client.Stream
}

func (x *bruceServiceStream) Close() error {
	return x.stream.Close()
}

func (x *bruceServiceStream) Context() context.Context {
	return x.stream.Context()
}

func (x *bruceServiceStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *bruceServiceStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *bruceServiceStream) Recv() (*StreamingResponse, error) {
	m := new(StreamingResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *bruceService) PingPong(ctx context.Context, opts ...client.CallOption) (Bruce_PingPongService, error) {
	req := c.c.NewRequest(c.name, "Bruce.PingPong", &Ping{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &bruceServicePingPong{stream}, nil
}

type Bruce_PingPongService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Ping) error
	Recv() (*Pong, error)
}

type bruceServicePingPong struct {
	stream client.Stream
}

func (x *bruceServicePingPong) Close() error {
	return x.stream.Close()
}

func (x *bruceServicePingPong) Context() context.Context {
	return x.stream.Context()
}

func (x *bruceServicePingPong) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *bruceServicePingPong) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *bruceServicePingPong) Send(m *Ping) error {
	return x.stream.Send(m)
}

func (x *bruceServicePingPong) Recv() (*Pong, error) {
	m := new(Pong)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Bruce service

type BruceHandler interface {
	Call(context.Context, *Request, *Response) error
	Stream(context.Context, *StreamingRequest, Bruce_StreamStream) error
	PingPong(context.Context, Bruce_PingPongStream) error
}

func RegisterBruceHandler(s server.Server, hdlr BruceHandler, opts ...server.HandlerOption) error {
	type bruce interface {
		Call(ctx context.Context, in *Request, out *Response) error
		Stream(ctx context.Context, stream server.Stream) error
		PingPong(ctx context.Context, stream server.Stream) error
	}
	type Bruce struct {
		bruce
	}
	h := &bruceHandler{hdlr}
	return s.Handle(s.NewHandler(&Bruce{h}, opts...))
}

type bruceHandler struct {
	BruceHandler
}

func (h *bruceHandler) Call(ctx context.Context, in *Request, out *Response) error {
	return h.BruceHandler.Call(ctx, in, out)
}

func (h *bruceHandler) Stream(ctx context.Context, stream server.Stream) error {
	m := new(StreamingRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.BruceHandler.Stream(ctx, m, &bruceStreamStream{stream})
}

type Bruce_StreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*StreamingResponse) error
}

type bruceStreamStream struct {
	stream server.Stream
}

func (x *bruceStreamStream) Close() error {
	return x.stream.Close()
}

func (x *bruceStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *bruceStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *bruceStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *bruceStreamStream) Send(m *StreamingResponse) error {
	return x.stream.Send(m)
}

func (h *bruceHandler) PingPong(ctx context.Context, stream server.Stream) error {
	return h.BruceHandler.PingPong(ctx, &brucePingPongStream{stream})
}

type Bruce_PingPongStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Pong) error
	Recv() (*Ping, error)
}

type brucePingPongStream struct {
	stream server.Stream
}

func (x *brucePingPongStream) Close() error {
	return x.stream.Close()
}

func (x *brucePingPongStream) Context() context.Context {
	return x.stream.Context()
}

func (x *brucePingPongStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *brucePingPongStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *brucePingPongStream) Send(m *Pong) error {
	return x.stream.Send(m)
}

func (x *brucePingPongStream) Recv() (*Ping, error) {
	m := new(Ping)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}
