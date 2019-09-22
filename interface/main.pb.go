// Code generated by protoc-gen-go. DO NOT EDIT.
// source: main.proto

package kci

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

type BuildRequest struct {
	Repository           string   `protobuf:"bytes,1,opt,name=repository,proto3" json:"repository,omitempty"`
	Sshkey               string   `protobuf:"bytes,2,opt,name=sshkey,proto3" json:"sshkey,omitempty"`
	Steps                []*Step  `protobuf:"bytes,3,rep,name=steps,proto3" json:"steps,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BuildRequest) Reset()         { *m = BuildRequest{} }
func (m *BuildRequest) String() string { return proto.CompactTextString(m) }
func (*BuildRequest) ProtoMessage()    {}
func (*BuildRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ed94b0a22d11796, []int{0}
}

func (m *BuildRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuildRequest.Unmarshal(m, b)
}
func (m *BuildRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuildRequest.Marshal(b, m, deterministic)
}
func (m *BuildRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuildRequest.Merge(m, src)
}
func (m *BuildRequest) XXX_Size() int {
	return xxx_messageInfo_BuildRequest.Size(m)
}
func (m *BuildRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BuildRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BuildRequest proto.InternalMessageInfo

func (m *BuildRequest) GetRepository() string {
	if m != nil {
		return m.Repository
	}
	return ""
}

func (m *BuildRequest) GetSshkey() string {
	if m != nil {
		return m.Sshkey
	}
	return ""
}

func (m *BuildRequest) GetSteps() []*Step {
	if m != nil {
		return m.Steps
	}
	return nil
}

type Step struct {
	Image                string            `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	Args                 []string          `protobuf:"bytes,2,rep,name=args,proto3" json:"args,omitempty"`
	Env                  map[string]string `protobuf:"bytes,3,rep,name=env,proto3" json:"env,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Step) Reset()         { *m = Step{} }
func (m *Step) String() string { return proto.CompactTextString(m) }
func (*Step) ProtoMessage()    {}
func (*Step) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ed94b0a22d11796, []int{1}
}

func (m *Step) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Step.Unmarshal(m, b)
}
func (m *Step) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Step.Marshal(b, m, deterministic)
}
func (m *Step) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Step.Merge(m, src)
}
func (m *Step) XXX_Size() int {
	return xxx_messageInfo_Step.Size(m)
}
func (m *Step) XXX_DiscardUnknown() {
	xxx_messageInfo_Step.DiscardUnknown(m)
}

var xxx_messageInfo_Step proto.InternalMessageInfo

func (m *Step) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *Step) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *Step) GetEnv() map[string]string {
	if m != nil {
		return m.Env
	}
	return nil
}

type BuildResponse struct {
	Status               string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Step                 int32    `protobuf:"varint,2,opt,name=step,proto3" json:"step,omitempty"`
	Data                 string   `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BuildResponse) Reset()         { *m = BuildResponse{} }
func (m *BuildResponse) String() string { return proto.CompactTextString(m) }
func (*BuildResponse) ProtoMessage()    {}
func (*BuildResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ed94b0a22d11796, []int{2}
}

func (m *BuildResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuildResponse.Unmarshal(m, b)
}
func (m *BuildResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuildResponse.Marshal(b, m, deterministic)
}
func (m *BuildResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuildResponse.Merge(m, src)
}
func (m *BuildResponse) XXX_Size() int {
	return xxx_messageInfo_BuildResponse.Size(m)
}
func (m *BuildResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BuildResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BuildResponse proto.InternalMessageInfo

func (m *BuildResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *BuildResponse) GetStep() int32 {
	if m != nil {
		return m.Step
	}
	return 0
}

func (m *BuildResponse) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type AddSecretRequest struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddSecretRequest) Reset()         { *m = AddSecretRequest{} }
func (m *AddSecretRequest) String() string { return proto.CompactTextString(m) }
func (*AddSecretRequest) ProtoMessage()    {}
func (*AddSecretRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ed94b0a22d11796, []int{3}
}

func (m *AddSecretRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddSecretRequest.Unmarshal(m, b)
}
func (m *AddSecretRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddSecretRequest.Marshal(b, m, deterministic)
}
func (m *AddSecretRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddSecretRequest.Merge(m, src)
}
func (m *AddSecretRequest) XXX_Size() int {
	return xxx_messageInfo_AddSecretRequest.Size(m)
}
func (m *AddSecretRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddSecretRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddSecretRequest proto.InternalMessageInfo

func (m *AddSecretRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *AddSecretRequest) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type GenericStatus struct {
	Status               string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenericStatus) Reset()         { *m = GenericStatus{} }
func (m *GenericStatus) String() string { return proto.CompactTextString(m) }
func (*GenericStatus) ProtoMessage()    {}
func (*GenericStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ed94b0a22d11796, []int{4}
}

func (m *GenericStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenericStatus.Unmarshal(m, b)
}
func (m *GenericStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenericStatus.Marshal(b, m, deterministic)
}
func (m *GenericStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenericStatus.Merge(m, src)
}
func (m *GenericStatus) XXX_Size() int {
	return xxx_messageInfo_GenericStatus.Size(m)
}
func (m *GenericStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_GenericStatus.DiscardUnknown(m)
}

var xxx_messageInfo_GenericStatus proto.InternalMessageInfo

func (m *GenericStatus) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func init() {
	proto.RegisterType((*BuildRequest)(nil), "kci.BuildRequest")
	proto.RegisterType((*Step)(nil), "kci.Step")
	proto.RegisterMapType((map[string]string)(nil), "kci.Step.EnvEntry")
	proto.RegisterType((*BuildResponse)(nil), "kci.BuildResponse")
	proto.RegisterType((*AddSecretRequest)(nil), "kci.AddSecretRequest")
	proto.RegisterType((*GenericStatus)(nil), "kci.GenericStatus")
}

func init() { proto.RegisterFile("main.proto", fileDescriptor_7ed94b0a22d11796) }

var fileDescriptor_7ed94b0a22d11796 = []byte{
	// 328 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x5f, 0x4b, 0xf3, 0x30,
	0x14, 0xc6, 0xe9, 0xb2, 0x8e, 0xb7, 0xe7, 0x75, 0x30, 0x0f, 0x2a, 0x65, 0x17, 0x3a, 0x8a, 0xe0,
	0xae, 0xca, 0x98, 0x30, 0x64, 0x77, 0x0a, 0xc3, 0x0b, 0x2f, 0x84, 0xee, 0x13, 0x64, 0xed, 0x61,
	0x86, 0x6d, 0x4d, 0x4d, 0xd2, 0xc1, 0xbe, 0x84, 0x9f, 0x59, 0xf2, 0x67, 0x32, 0x05, 0xc1, 0xbb,
	0xf3, 0x3c, 0x49, 0x9e, 0xfc, 0x9e, 0xb4, 0x00, 0x3b, 0x2e, 0xea, 0xbc, 0x51, 0xd2, 0x48, 0x64,
	0x9b, 0x52, 0x64, 0x6b, 0x38, 0x7b, 0x6a, 0xc5, 0xb6, 0x2a, 0xe8, 0xbd, 0x25, 0x6d, 0xf0, 0x1a,
	0x40, 0x51, 0x23, 0xb5, 0x30, 0x52, 0x1d, 0xd2, 0x68, 0x14, 0x8d, 0x93, 0xe2, 0xc4, 0xc1, 0x2b,
	0xe8, 0x69, 0xfd, 0xb6, 0xa1, 0x43, 0xda, 0x71, 0x6b, 0x41, 0xe1, 0x0d, 0xc4, 0xda, 0x50, 0xa3,
	0x53, 0x36, 0x62, 0xe3, 0xff, 0xd3, 0x24, 0xdf, 0x94, 0x22, 0x5f, 0x1a, 0x6a, 0x0a, 0xef, 0x67,
	0x1f, 0x11, 0x74, 0xad, 0xc6, 0x0b, 0x88, 0xc5, 0x8e, 0xaf, 0x29, 0x84, 0x7b, 0x81, 0x08, 0x5d,
	0xae, 0xd6, 0x3a, 0xed, 0x8c, 0xd8, 0x38, 0x29, 0xdc, 0x8c, 0xb7, 0xc0, 0xa8, 0xde, 0x87, 0x44,
	0xfc, 0x4a, 0xcc, 0x17, 0xf5, 0x7e, 0x51, 0x1b, 0x75, 0x28, 0xec, 0xf2, 0x70, 0x06, 0xff, 0x8e,
	0x06, 0x0e, 0x80, 0x59, 0x34, 0x9f, 0x6c, 0x47, 0x7b, 0xdb, 0x9e, 0x6f, 0x5b, 0x0a, 0xb8, 0x5e,
	0xcc, 0x3b, 0x0f, 0x51, 0xf6, 0x0a, 0xfd, 0xd0, 0x5c, 0x37, 0xb2, 0xd6, 0xe4, 0xaa, 0x19, 0x6e,
	0x5a, 0x1d, 0xce, 0x07, 0x65, 0xd1, 0x6c, 0x05, 0x97, 0x10, 0x17, 0x6e, 0xb6, 0x5e, 0xc5, 0x0d,
	0x4f, 0x99, 0xdb, 0xe9, 0xe6, 0x6c, 0x0e, 0x83, 0xc7, 0xaa, 0x5a, 0x52, 0xa9, 0xc8, 0x1c, 0x9f,
	0xf3, 0x8f, 0x40, 0xd9, 0x1d, 0xf4, 0x9f, 0xa9, 0x26, 0x25, 0xca, 0xa5, 0xbf, 0xf4, 0x17, 0x98,
	0xa9, 0x04, 0xf6, 0x52, 0x0a, 0x9c, 0x40, 0xbc, 0xb2, 0xf0, 0x78, 0xee, 0x9e, 0xe5, 0xf4, 0x13,
	0x0e, 0xf1, 0xd4, 0xf2, 0xdd, 0x26, 0x11, 0xce, 0x20, 0xe1, 0x47, 0x3a, 0xbc, 0x74, 0x5b, 0x7e,
	0xd2, 0x86, 0x93, 0xdf, 0x40, 0x56, 0x3d, 0xf7, 0xb3, 0xdc, 0x7f, 0x06, 0x00, 0x00, 0xff, 0xff,
	0xb6, 0xba, 0xf9, 0xd0, 0x3a, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// KciClient is the client API for Kci service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KciClient interface {
	Build(ctx context.Context, in *BuildRequest, opts ...grpc.CallOption) (Kci_BuildClient, error)
	AddSecret(ctx context.Context, in *AddSecretRequest, opts ...grpc.CallOption) (*GenericStatus, error)
}

type kciClient struct {
	cc *grpc.ClientConn
}

func NewKciClient(cc *grpc.ClientConn) KciClient {
	return &kciClient{cc}
}

func (c *kciClient) Build(ctx context.Context, in *BuildRequest, opts ...grpc.CallOption) (Kci_BuildClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Kci_serviceDesc.Streams[0], "/kci.Kci/build", opts...)
	if err != nil {
		return nil, err
	}
	x := &kciBuildClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Kci_BuildClient interface {
	Recv() (*BuildResponse, error)
	grpc.ClientStream
}

type kciBuildClient struct {
	grpc.ClientStream
}

func (x *kciBuildClient) Recv() (*BuildResponse, error) {
	m := new(BuildResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *kciClient) AddSecret(ctx context.Context, in *AddSecretRequest, opts ...grpc.CallOption) (*GenericStatus, error) {
	out := new(GenericStatus)
	err := c.cc.Invoke(ctx, "/kci.Kci/addSecret", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KciServer is the server API for Kci service.
type KciServer interface {
	Build(*BuildRequest, Kci_BuildServer) error
	AddSecret(context.Context, *AddSecretRequest) (*GenericStatus, error)
}

// UnimplementedKciServer can be embedded to have forward compatible implementations.
type UnimplementedKciServer struct {
}

func (*UnimplementedKciServer) Build(req *BuildRequest, srv Kci_BuildServer) error {
	return status.Errorf(codes.Unimplemented, "method Build not implemented")
}
func (*UnimplementedKciServer) AddSecret(ctx context.Context, req *AddSecretRequest) (*GenericStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddSecret not implemented")
}

func RegisterKciServer(s *grpc.Server, srv KciServer) {
	s.RegisterService(&_Kci_serviceDesc, srv)
}

func _Kci_Build_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(BuildRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(KciServer).Build(m, &kciBuildServer{stream})
}

type Kci_BuildServer interface {
	Send(*BuildResponse) error
	grpc.ServerStream
}

type kciBuildServer struct {
	grpc.ServerStream
}

func (x *kciBuildServer) Send(m *BuildResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Kci_AddSecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddSecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KciServer).AddSecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kci.Kci/AddSecret",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KciServer).AddSecret(ctx, req.(*AddSecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Kci_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kci.Kci",
	HandlerType: (*KciServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "addSecret",
			Handler:    _Kci_AddSecret_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "build",
			Handler:       _Kci_Build_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "main.proto",
}
