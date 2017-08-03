// Code generated by protoc-gen-go. DO NOT EDIT.
// source: control_plane_ca.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	control_plane_ca.proto

It has these top-level messages:
	SignCsrRequest
	SignCsrResponse
	GenerateCsrRequest
	GenerateCsrResponse
	FetchCertificateRequest
	FetchCertificateResponse
	LoadCertificateRequest
	LoadCertificateResponse
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import proto2 "github.com/spiffe/control-plane/plugins/common/proto"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

// ConfigureRequest from public import github.com/spiffe/control-plane/plugins/common/proto/common.proto
type ConfigureRequest proto2.ConfigureRequest

func (m *ConfigureRequest) Reset()         { (*proto2.ConfigureRequest)(m).Reset() }
func (m *ConfigureRequest) String() string { return (*proto2.ConfigureRequest)(m).String() }
func (*ConfigureRequest) ProtoMessage()    {}
func (m *ConfigureRequest) GetConfiguration() string {
	return (*proto2.ConfigureRequest)(m).GetConfiguration()
}

// ConfigureResponse from public import github.com/spiffe/control-plane/plugins/common/proto/common.proto
type ConfigureResponse proto2.ConfigureResponse

func (m *ConfigureResponse) Reset()         { (*proto2.ConfigureResponse)(m).Reset() }
func (m *ConfigureResponse) String() string { return (*proto2.ConfigureResponse)(m).String() }
func (*ConfigureResponse) ProtoMessage()    {}
func (m *ConfigureResponse) GetErrorList() []string {
	return (*proto2.ConfigureResponse)(m).GetErrorList()
}

// GetPluginInfoRequest from public import github.com/spiffe/control-plane/plugins/common/proto/common.proto
type GetPluginInfoRequest proto2.GetPluginInfoRequest

func (m *GetPluginInfoRequest) Reset()         { (*proto2.GetPluginInfoRequest)(m).Reset() }
func (m *GetPluginInfoRequest) String() string { return (*proto2.GetPluginInfoRequest)(m).String() }
func (*GetPluginInfoRequest) ProtoMessage()    {}

// GetPluginInfoResponse from public import github.com/spiffe/control-plane/plugins/common/proto/common.proto
type GetPluginInfoResponse proto2.GetPluginInfoResponse

func (m *GetPluginInfoResponse) Reset()         { (*proto2.GetPluginInfoResponse)(m).Reset() }
func (m *GetPluginInfoResponse) String() string { return (*proto2.GetPluginInfoResponse)(m).String() }
func (*GetPluginInfoResponse) ProtoMessage()    {}
func (m *GetPluginInfoResponse) GetPluginName() string {
	return (*proto2.GetPluginInfoResponse)(m).GetPluginName()
}
func (m *GetPluginInfoResponse) GetDescription() string {
	return (*proto2.GetPluginInfoResponse)(m).GetDescription()
}
func (m *GetPluginInfoResponse) GetDateCreated() string {
	return (*proto2.GetPluginInfoResponse)(m).GetDateCreated()
}
func (m *GetPluginInfoResponse) GetLocation() string {
	return (*proto2.GetPluginInfoResponse)(m).GetLocation()
}
func (m *GetPluginInfoResponse) GetVersion() string {
	return (*proto2.GetPluginInfoResponse)(m).GetVersion()
}
func (m *GetPluginInfoResponse) GetAuthor() string {
	return (*proto2.GetPluginInfoResponse)(m).GetAuthor()
}
func (m *GetPluginInfoResponse) GetCompany() string {
	return (*proto2.GetPluginInfoResponse)(m).GetCompany()
}

// *Represents a request with a certificate signing request.
type SignCsrRequest struct {
	Csr []byte `protobuf:"bytes,1,opt,name=csr,proto3" json:"csr,omitempty"`
}

func (m *SignCsrRequest) Reset()                    { *m = SignCsrRequest{} }
func (m *SignCsrRequest) String() string            { return proto1.CompactTextString(m) }
func (*SignCsrRequest) ProtoMessage()               {}
func (*SignCsrRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SignCsrRequest) GetCsr() []byte {
	if m != nil {
		return m.Csr
	}
	return nil
}

// *Represents a response with a signed certificate.
type SignCsrResponse struct {
	SignedCertificate []byte `protobuf:"bytes,1,opt,name=signedCertificate,proto3" json:"signedCertificate,omitempty"`
}

func (m *SignCsrResponse) Reset()                    { *m = SignCsrResponse{} }
func (m *SignCsrResponse) String() string            { return proto1.CompactTextString(m) }
func (*SignCsrResponse) ProtoMessage()               {}
func (*SignCsrResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SignCsrResponse) GetSignedCertificate() []byte {
	if m != nil {
		return m.SignedCertificate
	}
	return nil
}

// *Represents an empty request.
type GenerateCsrRequest struct {
}

func (m *GenerateCsrRequest) Reset()                    { *m = GenerateCsrRequest{} }
func (m *GenerateCsrRequest) String() string            { return proto1.CompactTextString(m) }
func (*GenerateCsrRequest) ProtoMessage()               {}
func (*GenerateCsrRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// *Represents a response with a certificate signing request.
type GenerateCsrResponse struct {
	Csr []byte `protobuf:"bytes,1,opt,name=csr,proto3" json:"csr,omitempty"`
}

func (m *GenerateCsrResponse) Reset()                    { *m = GenerateCsrResponse{} }
func (m *GenerateCsrResponse) String() string            { return proto1.CompactTextString(m) }
func (*GenerateCsrResponse) ProtoMessage()               {}
func (*GenerateCsrResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GenerateCsrResponse) GetCsr() []byte {
	if m != nil {
		return m.Csr
	}
	return nil
}

// *Represents an empty request.
type FetchCertificateRequest struct {
}

func (m *FetchCertificateRequest) Reset()                    { *m = FetchCertificateRequest{} }
func (m *FetchCertificateRequest) String() string            { return proto1.CompactTextString(m) }
func (*FetchCertificateRequest) ProtoMessage()               {}
func (*FetchCertificateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

// *Represents a response with a stored intermediate certificate.
type FetchCertificateResponse struct {
	StoredIntermediateCert []byte `protobuf:"bytes,1,opt,name=storedIntermediateCert,proto3" json:"storedIntermediateCert,omitempty"`
}

func (m *FetchCertificateResponse) Reset()                    { *m = FetchCertificateResponse{} }
func (m *FetchCertificateResponse) String() string            { return proto1.CompactTextString(m) }
func (*FetchCertificateResponse) ProtoMessage()               {}
func (*FetchCertificateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *FetchCertificateResponse) GetStoredIntermediateCert() []byte {
	if m != nil {
		return m.StoredIntermediateCert
	}
	return nil
}

// *Represents a request with a signed intermediate certificate.
type LoadCertificateRequest struct {
	SignedIntermediateCert []byte `protobuf:"bytes,1,opt,name=signedIntermediateCert,proto3" json:"signedIntermediateCert,omitempty"`
}

func (m *LoadCertificateRequest) Reset()                    { *m = LoadCertificateRequest{} }
func (m *LoadCertificateRequest) String() string            { return proto1.CompactTextString(m) }
func (*LoadCertificateRequest) ProtoMessage()               {}
func (*LoadCertificateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *LoadCertificateRequest) GetSignedIntermediateCert() []byte {
	if m != nil {
		return m.SignedIntermediateCert
	}
	return nil
}

// *Represents an empty response.
type LoadCertificateResponse struct {
}

func (m *LoadCertificateResponse) Reset()                    { *m = LoadCertificateResponse{} }
func (m *LoadCertificateResponse) String() string            { return proto1.CompactTextString(m) }
func (*LoadCertificateResponse) ProtoMessage()               {}
func (*LoadCertificateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func init() {
	proto1.RegisterType((*SignCsrRequest)(nil), "proto.SignCsrRequest")
	proto1.RegisterType((*SignCsrResponse)(nil), "proto.SignCsrResponse")
	proto1.RegisterType((*GenerateCsrRequest)(nil), "proto.GenerateCsrRequest")
	proto1.RegisterType((*GenerateCsrResponse)(nil), "proto.GenerateCsrResponse")
	proto1.RegisterType((*FetchCertificateRequest)(nil), "proto.FetchCertificateRequest")
	proto1.RegisterType((*FetchCertificateResponse)(nil), "proto.FetchCertificateResponse")
	proto1.RegisterType((*LoadCertificateRequest)(nil), "proto.LoadCertificateRequest")
	proto1.RegisterType((*LoadCertificateResponse)(nil), "proto.LoadCertificateResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ControlPlaneCA service

type ControlPlaneCAClient interface {
	// * Responsible for configuration of the plugin.
	Configure(ctx context.Context, in *proto2.ConfigureRequest, opts ...grpc.CallOption) (*proto2.ConfigureResponse, error)
	// * Returns the  version and related metadata of the installed plugin.
	GetPluginInfo(ctx context.Context, in *proto2.GetPluginInfoRequest, opts ...grpc.CallOption) (*proto2.GetPluginInfoResponse, error)
	// * Interface will take in a CSR and sign it with the stored intermediate certificate.
	SignCsr(ctx context.Context, in *SignCsrRequest, opts ...grpc.CallOption) (*SignCsrResponse, error)
	// * Used for generating a CSR for the intermediate signing certificate. The CSR will then be submitted to the CA plugin for signing.
	GenerateCsr(ctx context.Context, in *GenerateCsrRequest, opts ...grpc.CallOption) (*GenerateCsrResponse, error)
	// * Used to read the stored Intermediate CP cert.
	FetchCertificate(ctx context.Context, in *FetchCertificateRequest, opts ...grpc.CallOption) (*FetchCertificateResponse, error)
	// * Used for setting/storing the signed intermediate certificate.
	LoadCertificate(ctx context.Context, in *LoadCertificateRequest, opts ...grpc.CallOption) (*LoadCertificateResponse, error)
}

type controlPlaneCAClient struct {
	cc *grpc.ClientConn
}

func NewControlPlaneCAClient(cc *grpc.ClientConn) ControlPlaneCAClient {
	return &controlPlaneCAClient{cc}
}

func (c *controlPlaneCAClient) Configure(ctx context.Context, in *proto2.ConfigureRequest, opts ...grpc.CallOption) (*proto2.ConfigureResponse, error) {
	out := new(proto2.ConfigureResponse)
	err := grpc.Invoke(ctx, "/proto.ControlPlaneCA/Configure", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controlPlaneCAClient) GetPluginInfo(ctx context.Context, in *proto2.GetPluginInfoRequest, opts ...grpc.CallOption) (*proto2.GetPluginInfoResponse, error) {
	out := new(proto2.GetPluginInfoResponse)
	err := grpc.Invoke(ctx, "/proto.ControlPlaneCA/GetPluginInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controlPlaneCAClient) SignCsr(ctx context.Context, in *SignCsrRequest, opts ...grpc.CallOption) (*SignCsrResponse, error) {
	out := new(SignCsrResponse)
	err := grpc.Invoke(ctx, "/proto.ControlPlaneCA/SignCsr", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controlPlaneCAClient) GenerateCsr(ctx context.Context, in *GenerateCsrRequest, opts ...grpc.CallOption) (*GenerateCsrResponse, error) {
	out := new(GenerateCsrResponse)
	err := grpc.Invoke(ctx, "/proto.ControlPlaneCA/GenerateCsr", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controlPlaneCAClient) FetchCertificate(ctx context.Context, in *FetchCertificateRequest, opts ...grpc.CallOption) (*FetchCertificateResponse, error) {
	out := new(FetchCertificateResponse)
	err := grpc.Invoke(ctx, "/proto.ControlPlaneCA/FetchCertificate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controlPlaneCAClient) LoadCertificate(ctx context.Context, in *LoadCertificateRequest, opts ...grpc.CallOption) (*LoadCertificateResponse, error) {
	out := new(LoadCertificateResponse)
	err := grpc.Invoke(ctx, "/proto.ControlPlaneCA/LoadCertificate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ControlPlaneCA service

type ControlPlaneCAServer interface {
	// * Responsible for configuration of the plugin.
	Configure(context.Context, *proto2.ConfigureRequest) (*proto2.ConfigureResponse, error)
	// * Returns the  version and related metadata of the installed plugin.
	GetPluginInfo(context.Context, *proto2.GetPluginInfoRequest) (*proto2.GetPluginInfoResponse, error)
	// * Interface will take in a CSR and sign it with the stored intermediate certificate.
	SignCsr(context.Context, *SignCsrRequest) (*SignCsrResponse, error)
	// * Used for generating a CSR for the intermediate signing certificate. The CSR will then be submitted to the CA plugin for signing.
	GenerateCsr(context.Context, *GenerateCsrRequest) (*GenerateCsrResponse, error)
	// * Used to read the stored Intermediate CP cert.
	FetchCertificate(context.Context, *FetchCertificateRequest) (*FetchCertificateResponse, error)
	// * Used for setting/storing the signed intermediate certificate.
	LoadCertificate(context.Context, *LoadCertificateRequest) (*LoadCertificateResponse, error)
}

func RegisterControlPlaneCAServer(s *grpc.Server, srv ControlPlaneCAServer) {
	s.RegisterService(&_ControlPlaneCA_serviceDesc, srv)
}

func _ControlPlaneCA_Configure_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(proto2.ConfigureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlPlaneCAServer).Configure(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ControlPlaneCA/Configure",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlPlaneCAServer).Configure(ctx, req.(*proto2.ConfigureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ControlPlaneCA_GetPluginInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(proto2.GetPluginInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlPlaneCAServer).GetPluginInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ControlPlaneCA/GetPluginInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlPlaneCAServer).GetPluginInfo(ctx, req.(*proto2.GetPluginInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ControlPlaneCA_SignCsr_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignCsrRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlPlaneCAServer).SignCsr(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ControlPlaneCA/SignCsr",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlPlaneCAServer).SignCsr(ctx, req.(*SignCsrRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ControlPlaneCA_GenerateCsr_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateCsrRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlPlaneCAServer).GenerateCsr(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ControlPlaneCA/GenerateCsr",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlPlaneCAServer).GenerateCsr(ctx, req.(*GenerateCsrRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ControlPlaneCA_FetchCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlPlaneCAServer).FetchCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ControlPlaneCA/FetchCertificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlPlaneCAServer).FetchCertificate(ctx, req.(*FetchCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ControlPlaneCA_LoadCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoadCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlPlaneCAServer).LoadCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ControlPlaneCA/LoadCertificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlPlaneCAServer).LoadCertificate(ctx, req.(*LoadCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ControlPlaneCA_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ControlPlaneCA",
	HandlerType: (*ControlPlaneCAServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Configure",
			Handler:    _ControlPlaneCA_Configure_Handler,
		},
		{
			MethodName: "GetPluginInfo",
			Handler:    _ControlPlaneCA_GetPluginInfo_Handler,
		},
		{
			MethodName: "SignCsr",
			Handler:    _ControlPlaneCA_SignCsr_Handler,
		},
		{
			MethodName: "GenerateCsr",
			Handler:    _ControlPlaneCA_GenerateCsr_Handler,
		},
		{
			MethodName: "FetchCertificate",
			Handler:    _ControlPlaneCA_FetchCertificate_Handler,
		},
		{
			MethodName: "LoadCertificate",
			Handler:    _ControlPlaneCA_LoadCertificate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "control_plane_ca.proto",
}

func init() { proto1.RegisterFile("control_plane_ca.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 389 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x6f, 0x8f, 0xd2, 0x40,
	0x10, 0xc6, 0x25, 0x44, 0x8d, 0xa3, 0x02, 0xae, 0x5a, 0xa0, 0x2a, 0x9a, 0xbe, 0xd1, 0x17, 0x4a,
	0x13, 0x4d, 0x8c, 0xaf, 0xbc, 0x90, 0x5e, 0x8e, 0x70, 0xb9, 0x17, 0x0d, 0x7c, 0x00, 0x52, 0xda,
	0x69, 0xd9, 0x84, 0xee, 0xf6, 0x76, 0xb7, 0x9f, 0xf0, 0xbe, 0xd8, 0xa5, 0xed, 0xc0, 0x41, 0xff,
	0xbc, 0x82, 0xce, 0x33, 0xcf, 0x6f, 0x76, 0xf2, 0x0c, 0x58, 0xa1, 0x14, 0x46, 0xc9, 0xc3, 0x36,
	0x3b, 0x04, 0x02, 0xb7, 0x61, 0x30, 0xcf, 0x94, 0x34, 0x92, 0x3d, 0x2f, 0x7f, 0xec, 0x45, 0xc2,
	0xcd, 0x3e, 0xdf, 0xcd, 0x43, 0x99, 0xba, 0x3a, 0xe3, 0x71, 0x8c, 0x2e, 0x19, 0x7e, 0x95, 0x06,
	0x37, 0x3b, 0xe4, 0x09, 0x17, 0xda, 0x0d, 0x65, 0x9a, 0x4a, 0xe1, 0x96, 0x2e, 0xfa, 0xa8, 0x48,
	0x8e, 0x03, 0x83, 0x0d, 0x4f, 0x84, 0xa7, 0xd5, 0x1a, 0xef, 0x73, 0xd4, 0x86, 0x8d, 0xa0, 0x1f,
	0x6a, 0x35, 0xe9, 0x7d, 0xeb, 0xfd, 0x78, 0xb3, 0x2e, 0xfe, 0x3a, 0x57, 0x30, 0x3c, 0xf5, 0xe8,
	0x4c, 0x0a, 0x8d, 0xec, 0x27, 0xbc, 0xd3, 0x3c, 0x11, 0x18, 0x79, 0xa8, 0x0c, 0x8f, 0x79, 0x18,
	0x18, 0x24, 0x4b, 0x53, 0x70, 0x3e, 0x00, 0x5b, 0xa2, 0x40, 0x15, 0x18, 0x7c, 0x1a, 0xe4, 0x7c,
	0x87, 0xf7, 0x17, 0x55, 0x42, 0x37, 0xe7, 0x4f, 0x61, 0x7c, 0x83, 0x26, 0xdc, 0x9f, 0x21, 0x8f,
	0x8c, 0x35, 0x4c, 0x9a, 0x12, 0x81, 0xfe, 0x82, 0xa5, 0x8d, 0x54, 0x18, 0xad, 0x84, 0x41, 0x95,
	0x62, 0xc4, 0x8b, 0x49, 0xa8, 0x0c, 0xb1, 0x3b, 0x54, 0xc7, 0x07, 0xeb, 0x4e, 0x06, 0x51, 0x73,
	0x5a, 0x49, 0x2c, 0x97, 0xeb, 0x24, 0xb6, 0xaa, 0xc5, 0x02, 0x0d, 0x62, 0xf5, 0xc8, 0xdf, 0x0f,
	0x7d, 0x18, 0x78, 0x55, 0x66, 0x7e, 0x11, 0x99, 0xb7, 0x60, 0xff, 0xe1, 0x95, 0x27, 0x45, 0xcc,
	0x93, 0x5c, 0x21, 0x1b, 0x57, 0x39, 0xcd, 0x4f, 0x15, 0x7a, 0x8b, 0x3d, 0x69, 0x0a, 0xb4, 0xf7,
	0x2d, 0xbc, 0x5d, 0xa2, 0xf1, 0xcb, 0xe8, 0x57, 0x22, 0x96, 0xec, 0x13, 0xb5, 0x5e, 0x54, 0x8f,
	0x9c, 0xcf, 0xed, 0x22, 0xb1, 0xfe, 0xc1, 0x4b, 0x8a, 0x9e, 0x7d, 0xa4, 0xc6, 0xcb, 0x73, 0xb1,
	0xad, 0x7a, 0x99, 0x9c, 0xd7, 0xf0, 0xfa, 0x2c, 0x5d, 0x36, 0x3d, 0x8d, 0xa9, 0xdf, 0x81, 0x6d,
	0xb7, 0x49, 0x44, 0xd9, 0xc0, 0xa8, 0x9e, 0x2f, 0x9b, 0x51, 0x7f, 0xc7, 0x4d, 0xd8, 0x5f, 0x3b,
	0x75, 0x82, 0xfa, 0x30, 0xac, 0xc5, 0xc1, 0xbe, 0x90, 0xa7, 0x3d, 0x78, 0x7b, 0xd6, 0x25, 0x57,
	0x44, 0xff, 0xd9, 0xee, 0x45, 0xd9, 0xf0, 0xe7, 0x31, 0x00, 0x00, 0xff, 0xff, 0x3b, 0x42, 0xee,
	0x32, 0xb2, 0x03, 0x00, 0x00,
}
