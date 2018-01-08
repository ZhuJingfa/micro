// Code generated by protoc-gen-go.
// source: github.com/micro/micro/run/proto/run.proto
// DO NOT EDIT!

/*
Package go_micro_run is a generated protocol buffer package.

It is generated from these files:
	github.com/micro/micro/run/proto/run.proto

It has these top-level messages:
	Source
	Binary
	Process
	FetchRequest
	FetchResponse
	BuildRequest
	BuildResponse
	ExecRequest
	ExecResponse
	KillRequest
	KillResponse
	WaitRequest
	WaitResponse
	RunRequest
	RunResponse
	StopRequest
	StopResponse
	StatusRequest
	StatusResponse
*/
package go_micro_run

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
  client "micro/go-micro/client"
  server "micro/go-micro/server"
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

type Source struct {
	Url string `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
	Dir string `protobuf:"bytes,2,opt,name=dir" json:"dir,omitempty"`
}

func (m *Source) Reset()                    { *m = Source{} }
func (m *Source) String() string            { return proto.CompactTextString(m) }
func (*Source) ProtoMessage()               {}
func (*Source) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Source) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Source) GetDir() string {
	if m != nil {
		return m.Dir
	}
	return ""
}

type Binary struct {
	Path   string  `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
	Source *Source `protobuf:"bytes,2,opt,name=source" json:"source,omitempty"`
}

func (m *Binary) Reset()                    { *m = Binary{} }
func (m *Binary) String() string            { return proto.CompactTextString(m) }
func (*Binary) ProtoMessage()               {}
func (*Binary) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Binary) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Binary) GetSource() *Source {
	if m != nil {
		return m.Source
	}
	return nil
}

type Process struct {
	Id     string  `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Binary *Binary `protobuf:"bytes,2,opt,name=binary" json:"binary,omitempty"`
}

func (m *Process) Reset()                    { *m = Process{} }
func (m *Process) String() string            { return proto.CompactTextString(m) }
func (*Process) ProtoMessage()               {}
func (*Process) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Process) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Process) GetBinary() *Binary {
	if m != nil {
		return m.Binary
	}
	return nil
}

type FetchRequest struct {
	Url    string `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
	Update bool   `protobuf:"varint,2,opt,name=update" json:"update,omitempty"`
}

func (m *FetchRequest) Reset()                    { *m = FetchRequest{} }
func (m *FetchRequest) String() string            { return proto.CompactTextString(m) }
func (*FetchRequest) ProtoMessage()               {}
func (*FetchRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *FetchRequest) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *FetchRequest) GetUpdate() bool {
	if m != nil {
		return m.Update
	}
	return false
}

type FetchResponse struct {
	Source *Source `protobuf:"bytes,1,opt,name=source" json:"source,omitempty"`
}

func (m *FetchResponse) Reset()                    { *m = FetchResponse{} }
func (m *FetchResponse) String() string            { return proto.CompactTextString(m) }
func (*FetchResponse) ProtoMessage()               {}
func (*FetchResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *FetchResponse) GetSource() *Source {
	if m != nil {
		return m.Source
	}
	return nil
}

type BuildRequest struct {
	Source *Source `protobuf:"bytes,1,opt,name=source" json:"source,omitempty"`
}

func (m *BuildRequest) Reset()                    { *m = BuildRequest{} }
func (m *BuildRequest) String() string            { return proto.CompactTextString(m) }
func (*BuildRequest) ProtoMessage()               {}
func (*BuildRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *BuildRequest) GetSource() *Source {
	if m != nil {
		return m.Source
	}
	return nil
}

type BuildResponse struct {
	Binary *Binary `protobuf:"bytes,1,opt,name=binary" json:"binary,omitempty"`
}

func (m *BuildResponse) Reset()                    { *m = BuildResponse{} }
func (m *BuildResponse) String() string            { return proto.CompactTextString(m) }
func (*BuildResponse) ProtoMessage()               {}
func (*BuildResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *BuildResponse) GetBinary() *Binary {
	if m != nil {
		return m.Binary
	}
	return nil
}

type ExecRequest struct {
	Binary *Binary `protobuf:"bytes,1,opt,name=binary" json:"binary,omitempty"`
}

func (m *ExecRequest) Reset()                    { *m = ExecRequest{} }
func (m *ExecRequest) String() string            { return proto.CompactTextString(m) }
func (*ExecRequest) ProtoMessage()               {}
func (*ExecRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ExecRequest) GetBinary() *Binary {
	if m != nil {
		return m.Binary
	}
	return nil
}

type ExecResponse struct {
	Process *Process `protobuf:"bytes,1,opt,name=process" json:"process,omitempty"`
}

func (m *ExecResponse) Reset()                    { *m = ExecResponse{} }
func (m *ExecResponse) String() string            { return proto.CompactTextString(m) }
func (*ExecResponse) ProtoMessage()               {}
func (*ExecResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *ExecResponse) GetProcess() *Process {
	if m != nil {
		return m.Process
	}
	return nil
}

type KillRequest struct {
	Process *Process `protobuf:"bytes,1,opt,name=process" json:"process,omitempty"`
}

func (m *KillRequest) Reset()                    { *m = KillRequest{} }
func (m *KillRequest) String() string            { return proto.CompactTextString(m) }
func (*KillRequest) ProtoMessage()               {}
func (*KillRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *KillRequest) GetProcess() *Process {
	if m != nil {
		return m.Process
	}
	return nil
}

type KillResponse struct {
}

func (m *KillResponse) Reset()                    { *m = KillResponse{} }
func (m *KillResponse) String() string            { return proto.CompactTextString(m) }
func (*KillResponse) ProtoMessage()               {}
func (*KillResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

type WaitRequest struct {
	Process *Process `protobuf:"bytes,1,opt,name=process" json:"process,omitempty"`
}

func (m *WaitRequest) Reset()                    { *m = WaitRequest{} }
func (m *WaitRequest) String() string            { return proto.CompactTextString(m) }
func (*WaitRequest) ProtoMessage()               {}
func (*WaitRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *WaitRequest) GetProcess() *Process {
	if m != nil {
		return m.Process
	}
	return nil
}

type WaitResponse struct {
	Error string `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
}

func (m *WaitResponse) Reset()                    { *m = WaitResponse{} }
func (m *WaitResponse) String() string            { return proto.CompactTextString(m) }
func (*WaitResponse) ProtoMessage()               {}
func (*WaitResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *WaitResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type RunRequest struct {
	Url     string `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
	Restart bool   `protobuf:"varint,2,opt,name=restart" json:"restart,omitempty"`
	Update  bool   `protobuf:"varint,3,opt,name=update" json:"update,omitempty"`
}

func (m *RunRequest) Reset()                    { *m = RunRequest{} }
func (m *RunRequest) String() string            { return proto.CompactTextString(m) }
func (*RunRequest) ProtoMessage()               {}
func (*RunRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *RunRequest) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *RunRequest) GetRestart() bool {
	if m != nil {
		return m.Restart
	}
	return false
}

func (m *RunRequest) GetUpdate() bool {
	if m != nil {
		return m.Update
	}
	return false
}

type RunResponse struct {
}

func (m *RunResponse) Reset()                    { *m = RunResponse{} }
func (m *RunResponse) String() string            { return proto.CompactTextString(m) }
func (*RunResponse) ProtoMessage()               {}
func (*RunResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

type StopRequest struct {
	Url string `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
}

func (m *StopRequest) Reset()                    { *m = StopRequest{} }
func (m *StopRequest) String() string            { return proto.CompactTextString(m) }
func (*StopRequest) ProtoMessage()               {}
func (*StopRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *StopRequest) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type StopResponse struct {
}

func (m *StopResponse) Reset()                    { *m = StopResponse{} }
func (m *StopResponse) String() string            { return proto.CompactTextString(m) }
func (*StopResponse) ProtoMessage()               {}
func (*StopResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

type StatusRequest struct {
	Url string `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
}

func (m *StatusRequest) Reset()                    { *m = StatusRequest{} }
func (m *StatusRequest) String() string            { return proto.CompactTextString(m) }
func (*StatusRequest) ProtoMessage()               {}
func (*StatusRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{17} }

func (m *StatusRequest) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type StatusResponse struct {
	Info string `protobuf:"bytes,1,opt,name=info" json:"info,omitempty"`
}

func (m *StatusResponse) Reset()                    { *m = StatusResponse{} }
func (m *StatusResponse) String() string            { return proto.CompactTextString(m) }
func (*StatusResponse) ProtoMessage()               {}
func (*StatusResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{18} }

func (m *StatusResponse) GetInfo() string {
	if m != nil {
		return m.Info
	}
	return ""
}

func init() {
	proto.RegisterType((*Source)(nil), "go.micro.run.Source")
	proto.RegisterType((*Binary)(nil), "go.micro.run.Binary")
	proto.RegisterType((*Process)(nil), "go.micro.run.Process")
	proto.RegisterType((*FetchRequest)(nil), "go.micro.run.FetchRequest")
	proto.RegisterType((*FetchResponse)(nil), "go.micro.run.FetchResponse")
	proto.RegisterType((*BuildRequest)(nil), "go.micro.run.BuildRequest")
	proto.RegisterType((*BuildResponse)(nil), "go.micro.run.BuildResponse")
	proto.RegisterType((*ExecRequest)(nil), "go.micro.run.ExecRequest")
	proto.RegisterType((*ExecResponse)(nil), "go.micro.run.ExecResponse")
	proto.RegisterType((*KillRequest)(nil), "go.micro.run.KillRequest")
	proto.RegisterType((*KillResponse)(nil), "go.micro.run.KillResponse")
	proto.RegisterType((*WaitRequest)(nil), "go.micro.run.WaitRequest")
	proto.RegisterType((*WaitResponse)(nil), "go.micro.run.WaitResponse")
	proto.RegisterType((*RunRequest)(nil), "go.micro.run.RunRequest")
	proto.RegisterType((*RunResponse)(nil), "go.micro.run.RunResponse")
	proto.RegisterType((*StopRequest)(nil), "go.micro.run.StopRequest")
	proto.RegisterType((*StopResponse)(nil), "go.micro.run.StopResponse")
	proto.RegisterType((*StatusRequest)(nil), "go.micro.run.StatusRequest")
	proto.RegisterType((*StatusResponse)(nil), "go.micro.run.StatusResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Runtime service

type RuntimeClient interface {
	Fetch(ctx context.Context, in *FetchRequest, opts ...client.CallOption) (*FetchResponse, error)
	Build(ctx context.Context, in *BuildRequest, opts ...client.CallOption) (*BuildResponse, error)
	Exec(ctx context.Context, in *ExecRequest, opts ...client.CallOption) (*ExecResponse, error)
	Kill(ctx context.Context, in *KillRequest, opts ...client.CallOption) (*KillResponse, error)
	Wait(ctx context.Context, in *WaitRequest, opts ...client.CallOption) (Runtime_WaitClient, error)
}

type runtimeClient struct {
	c           client.Client
	serviceName string
}

func NewRuntimeClient(serviceName string, c client.Client) RuntimeClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "go.micro.run"
	}
	return &runtimeClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *runtimeClient) Fetch(ctx context.Context, in *FetchRequest, opts ...client.CallOption) (*FetchResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Runtime.Fetch", in)
	out := new(FetchResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *runtimeClient) Build(ctx context.Context, in *BuildRequest, opts ...client.CallOption) (*BuildResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Runtime.Build", in)
	out := new(BuildResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *runtimeClient) Exec(ctx context.Context, in *ExecRequest, opts ...client.CallOption) (*ExecResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Runtime.Exec", in)
	out := new(ExecResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *runtimeClient) Kill(ctx context.Context, in *KillRequest, opts ...client.CallOption) (*KillResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Runtime.Kill", in)
	out := new(KillResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *runtimeClient) Wait(ctx context.Context, in *WaitRequest, opts ...client.CallOption) (Runtime_WaitClient, error) {
	req := c.c.NewRequest(c.serviceName, "Runtime.Wait", &WaitRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &runtimeWaitClient{stream}, nil
}

type Runtime_WaitClient interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*WaitResponse, error)
}

type runtimeWaitClient struct {
	stream client.Streamer
}

func (x *runtimeWaitClient) Close() error {
	return x.stream.Close()
}

func (x *runtimeWaitClient) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *runtimeWaitClient) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *runtimeWaitClient) Recv() (*WaitResponse, error) {
	m := new(WaitResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Runtime service

type RuntimeHandler interface {
	Fetch(context.Context, *FetchRequest, *FetchResponse) error
	Build(context.Context, *BuildRequest, *BuildResponse) error
	Exec(context.Context, *ExecRequest, *ExecResponse) error
	Kill(context.Context, *KillRequest, *KillResponse) error
	Wait(context.Context, *WaitRequest, Runtime_WaitStream) error
}

func RegisterRuntimeHandler(s server.Server, hdlr RuntimeHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Runtime{hdlr}, opts...))
}

type Runtime struct {
	RuntimeHandler
}

func (h *Runtime) Fetch(ctx context.Context, in *FetchRequest, out *FetchResponse) error {
	return h.RuntimeHandler.Fetch(ctx, in, out)
}

func (h *Runtime) Build(ctx context.Context, in *BuildRequest, out *BuildResponse) error {
	return h.RuntimeHandler.Build(ctx, in, out)
}

func (h *Runtime) Exec(ctx context.Context, in *ExecRequest, out *ExecResponse) error {
	return h.RuntimeHandler.Exec(ctx, in, out)
}

func (h *Runtime) Kill(ctx context.Context, in *KillRequest, out *KillResponse) error {
	return h.RuntimeHandler.Kill(ctx, in, out)
}

func (h *Runtime) Wait(ctx context.Context, stream server.Streamer) error {
	m := new(WaitRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.RuntimeHandler.Wait(ctx, m, &runtimeWaitStream{stream})
}

type Runtime_WaitStream interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*WaitResponse) error
}

type runtimeWaitStream struct {
	stream server.Streamer
}

func (x *runtimeWaitStream) Close() error {
	return x.stream.Close()
}

func (x *runtimeWaitStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *runtimeWaitStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *runtimeWaitStream) Send(m *WaitResponse) error {
	return x.stream.Send(m)
}

// Client API for Service service

type ServiceClient interface {
	Run(ctx context.Context, in *RunRequest, opts ...client.CallOption) (*RunResponse, error)
	Stop(ctx context.Context, in *StopRequest, opts ...client.CallOption) (*StopResponse, error)
	Status(ctx context.Context, in *StatusRequest, opts ...client.CallOption) (*StatusResponse, error)
}

type serviceClient struct {
	c           client.Client
	serviceName string
}

func NewServiceClient(serviceName string, c client.Client) ServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "go.micro.run"
	}
	return &serviceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *serviceClient) Run(ctx context.Context, in *RunRequest, opts ...client.CallOption) (*RunResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Service.Run", in)
	out := new(RunResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) Stop(ctx context.Context, in *StopRequest, opts ...client.CallOption) (*StopResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Service.Stop", in)
	out := new(StopResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) Status(ctx context.Context, in *StatusRequest, opts ...client.CallOption) (*StatusResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Service.Status", in)
	out := new(StatusResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Service service

type ServiceHandler interface {
	Run(context.Context, *RunRequest, *RunResponse) error
	Stop(context.Context, *StopRequest, *StopResponse) error
	Status(context.Context, *StatusRequest, *StatusResponse) error
}

func RegisterServiceHandler(s server.Server, hdlr ServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Service{hdlr}, opts...))
}

type Service struct {
	ServiceHandler
}

func (h *Service) Run(ctx context.Context, in *RunRequest, out *RunResponse) error {
	return h.ServiceHandler.Run(ctx, in, out)
}

func (h *Service) Stop(ctx context.Context, in *StopRequest, out *StopResponse) error {
	return h.ServiceHandler.Stop(ctx, in, out)
}

func (h *Service) Status(ctx context.Context, in *StatusRequest, out *StatusResponse) error {
	return h.ServiceHandler.Status(ctx, in, out)
}

func init() { proto.RegisterFile("micro/micro/run/proto/run.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 547 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x95, 0xc1, 0x6f, 0xd3, 0x30,
	0x14, 0xc6, 0x9b, 0xb6, 0x4b, 0xe1, 0x35, 0xad, 0x90, 0x35, 0x50, 0x97, 0x21, 0x01, 0xd1, 0x0e,
	0x08, 0x4d, 0x29, 0x2a, 0x17, 0x24, 0x36, 0x26, 0x2a, 0x0d, 0x24, 0xb8, 0x4c, 0xe9, 0x81, 0x73,
	0x9a, 0x98, 0xd5, 0x52, 0x1b, 0x07, 0xc7, 0x46, 0xf0, 0x57, 0x72, 0xe3, 0xef, 0x41, 0xb6, 0x5f,
	0x34, 0x27, 0x4a, 0xa5, 0x75, 0x97, 0xca, 0x8e, 0xdf, 0xf7, 0xbd, 0xf7, 0xd5, 0xbf, 0x28, 0xf0,
	0xe6, 0x96, 0xc9, 0x8d, 0x5a, 0xc7, 0x19, 0xdf, 0xcd, 0x77, 0x2c, 0x13, 0x1c, 0x7f, 0x85, 0x2a,
	0xe6, 0xa5, 0xe0, 0xd2, 0xac, 0x62, 0xb3, 0x22, 0xc1, 0x2d, 0x8f, 0xcd, 0x69, 0x2c, 0x54, 0x11,
	0x9d, 0x83, 0xbf, 0xe2, 0x4a, 0x64, 0x94, 0x3c, 0x81, 0x81, 0x12, 0xdb, 0x99, 0xf7, 0xd2, 0x7b,
	0xfd, 0x38, 0xd1, 0x4b, 0xfd, 0x24, 0x67, 0x62, 0xd6, 0xb7, 0x4f, 0x72, 0x26, 0xa2, 0xaf, 0xe0,
	0x2f, 0x59, 0x91, 0x8a, 0x3f, 0x84, 0xc0, 0xb0, 0x4c, 0xe5, 0x06, 0xcb, 0xcd, 0x9a, 0x9c, 0x83,
	0x5f, 0x19, 0x2f, 0x23, 0x19, 0x2f, 0x8e, 0x63, 0xb7, 0x55, 0x6c, 0xfb, 0x24, 0x58, 0x13, 0x7d,
	0x81, 0xd1, 0x8d, 0xe0, 0x19, 0xad, 0x2a, 0x32, 0x85, 0x3e, 0xcb, 0xd1, 0xaa, 0xcf, 0x72, 0x6d,
	0xb4, 0x36, 0x6d, 0xba, 0x8d, 0xec, 0x08, 0x09, 0xd6, 0x44, 0xef, 0x21, 0xf8, 0x4c, 0x65, 0xb6,
	0x49, 0xe8, 0x4f, 0x45, 0x2b, 0xd9, 0x11, 0xe4, 0x19, 0xf8, 0xaa, 0xcc, 0x53, 0x69, 0x07, 0x7b,
	0x94, 0xe0, 0x2e, 0xba, 0x84, 0x09, 0x2a, 0xab, 0x92, 0x17, 0x15, 0x75, 0x12, 0x78, 0xf7, 0x48,
	0x70, 0x01, 0xc1, 0x52, 0xb1, 0x6d, 0x5e, 0x37, 0x3e, 0x4c, 0x7d, 0x09, 0x13, 0x54, 0xdf, 0x35,
	0xc7, 0xd4, 0xde, 0x3d, 0x52, 0x7f, 0x80, 0xf1, 0xf5, 0x6f, 0x9a, 0x39, 0xbd, 0x0f, 0x10, 0x5f,
	0x41, 0x60, 0xc5, 0xd8, 0x7a, 0x0e, 0xa3, 0xd2, 0xde, 0x05, 0xca, 0x9f, 0x36, 0xe5, 0x78, 0x51,
	0x49, 0x5d, 0x15, 0x7d, 0x84, 0xf1, 0x37, 0xb6, 0xdd, 0xd6, 0xdd, 0x0f, 0xd6, 0x4f, 0x21, 0xb0,
	0x7a, 0x3b, 0x80, 0xf6, 0xfb, 0x9e, 0x32, 0xf9, 0x60, 0xbf, 0x33, 0x08, 0xac, 0x1e, 0x03, 0x1d,
	0xc3, 0x11, 0x15, 0x82, 0x0b, 0xa4, 0xc0, 0x6e, 0xa2, 0x1b, 0x80, 0x44, 0x15, 0xfb, 0x39, 0x99,
	0xc1, 0x48, 0xd0, 0x4a, 0xa6, 0x42, 0x22, 0x28, 0xf5, 0xd6, 0x21, 0x68, 0xd0, 0x20, 0x68, 0x02,
	0x63, 0xe3, 0x88, 0x31, 0x5e, 0xc0, 0x78, 0x25, 0x79, 0xb9, 0xb7, 0x83, 0xce, 0x6d, 0x0b, 0x50,
	0xf0, 0x0a, 0x26, 0x2b, 0x99, 0x4a, 0x55, 0xed, 0x97, 0x9c, 0xc1, 0xb4, 0x2e, 0xc1, 0x70, 0x04,
	0x86, 0xac, 0xf8, 0xc1, 0xeb, 0x77, 0x4f, 0xaf, 0x17, 0xff, 0xfa, 0x30, 0x4a, 0x54, 0x21, 0xd9,
	0x8e, 0x92, 0x25, 0x1c, 0x19, 0xac, 0x49, 0xd8, 0xfc, 0xd7, 0xdc, 0xb7, 0x24, 0x3c, 0xed, 0x3c,
	0xc3, 0xb1, 0x7a, 0xda, 0xc3, 0xd0, 0xd9, 0xf6, 0x70, 0x81, 0x6f, 0x7b, 0x34, 0x70, 0x8e, 0x7a,
	0xe4, 0x0a, 0x86, 0x9a, 0x32, 0x72, 0xd2, 0x2c, 0x73, 0xb0, 0x0d, 0xc3, 0xae, 0x23, 0xd7, 0x40,
	0x53, 0xd2, 0x36, 0x70, 0xc8, 0x6b, 0x1b, 0x34, 0xa0, 0xea, 0x91, 0x4f, 0x30, 0xd4, 0x58, 0xb4,
	0x0d, 0x1c, 0xd4, 0xda, 0x06, 0x2e, 0x45, 0x51, 0xef, 0xad, 0xb7, 0xf8, 0xeb, 0xc1, 0x68, 0x45,
	0xc5, 0x2f, 0x96, 0x51, 0x72, 0x01, 0x83, 0x44, 0x15, 0x64, 0xd6, 0x94, 0xdc, 0x21, 0x15, 0x9e,
	0x74, 0x9c, 0xb8, 0x69, 0xf4, 0xdd, 0xb7, 0x87, 0x71, 0x80, 0x69, 0x0f, 0xd3, 0x40, 0xa5, 0x47,
	0xae, 0xc1, 0xb7, 0x24, 0x90, 0xd3, 0x76, 0x9d, 0x83, 0x50, 0xf8, 0xbc, 0xfb, 0xb0, 0xb6, 0x59,
	0xfb, 0xe6, 0x3b, 0xf0, 0xee, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x45, 0x01, 0x3a, 0x90, 0x35,
	0x06, 0x00, 0x00,
}
