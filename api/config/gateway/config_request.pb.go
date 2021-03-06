// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/config/gateway/config_request.proto

package gateway

import (
	fmt "fmt"
	shared "github.com/chef/automate/api/config/shared"
	_ "github.com/chef/automate/components/automate-grpc/protoc-gen-a2-config/api/a2conf"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type ConfigRequest struct {
	V1                   *ConfigRequest_V1 `protobuf:"bytes,3,opt,name=v1,proto3" json:"v1,omitempty" toml:"v1,omitempty" mapstructure:"v1,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte            `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32             `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest) Reset()         { *m = ConfigRequest{} }
func (m *ConfigRequest) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest) ProtoMessage()    {}
func (*ConfigRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9be4cf630cc5f57c, []int{0}
}

func (m *ConfigRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest.Unmarshal(m, b)
}
func (m *ConfigRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest.Marshal(b, m, deterministic)
}
func (m *ConfigRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest.Merge(m, src)
}
func (m *ConfigRequest) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest.Size(m)
}
func (m *ConfigRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest proto.InternalMessageInfo

func (m *ConfigRequest) GetV1() *ConfigRequest_V1 {
	if m != nil {
		return m.V1
	}
	return nil
}

type ConfigRequest_V1 struct {
	Sys                  *ConfigRequest_V1_System  `protobuf:"bytes,1,opt,name=sys,proto3" json:"sys,omitempty" toml:"sys,omitempty" mapstructure:"sys,omitempty"`
	Svc                  *ConfigRequest_V1_Service `protobuf:"bytes,2,opt,name=svc,proto3" json:"svc,omitempty" toml:"svc,omitempty" mapstructure:"svc,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte                    `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                     `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1) Reset()         { *m = ConfigRequest_V1{} }
func (m *ConfigRequest_V1) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1) ProtoMessage()    {}
func (*ConfigRequest_V1) Descriptor() ([]byte, []int) {
	return fileDescriptor_9be4cf630cc5f57c, []int{0, 0}
}

func (m *ConfigRequest_V1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest_V1.Unmarshal(m, b)
}
func (m *ConfigRequest_V1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest_V1.Marshal(b, m, deterministic)
}
func (m *ConfigRequest_V1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest_V1.Merge(m, src)
}
func (m *ConfigRequest_V1) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest_V1.Size(m)
}
func (m *ConfigRequest_V1) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest_V1.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest_V1 proto.InternalMessageInfo

func (m *ConfigRequest_V1) GetSys() *ConfigRequest_V1_System {
	if m != nil {
		return m.Sys
	}
	return nil
}

func (m *ConfigRequest_V1) GetSvc() *ConfigRequest_V1_Service {
	if m != nil {
		return m.Svc
	}
	return nil
}

type ConfigRequest_V1_System struct {
	Mlsa                 *shared.Mlsa                     `protobuf:"bytes,1,opt,name=mlsa,proto3" json:"mlsa,omitempty" toml:"mlsa,omitempty" mapstructure:"mlsa,omitempty"`
	Tls                  *shared.TLSCredentials           `protobuf:"bytes,2,opt,name=tls,proto3" json:"tls,omitempty" toml:"tls,omitempty" mapstructure:"tls,omitempty"`
	Service              *ConfigRequest_V1_System_Service `protobuf:"bytes,3,opt,name=service,proto3" json:"service,omitempty" toml:"service,omitempty" mapstructure:"service,omitempty"`
	Log                  *ConfigRequest_V1_System_Log     `protobuf:"bytes,4,opt,name=log,proto3" json:"log,omitempty" toml:"log,omitempty" mapstructure:"log,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte                           `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                            `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1_System) Reset()         { *m = ConfigRequest_V1_System{} }
func (m *ConfigRequest_V1_System) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1_System) ProtoMessage()    {}
func (*ConfigRequest_V1_System) Descriptor() ([]byte, []int) {
	return fileDescriptor_9be4cf630cc5f57c, []int{0, 0, 0}
}

func (m *ConfigRequest_V1_System) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest_V1_System.Unmarshal(m, b)
}
func (m *ConfigRequest_V1_System) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest_V1_System.Marshal(b, m, deterministic)
}
func (m *ConfigRequest_V1_System) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest_V1_System.Merge(m, src)
}
func (m *ConfigRequest_V1_System) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest_V1_System.Size(m)
}
func (m *ConfigRequest_V1_System) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest_V1_System.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest_V1_System proto.InternalMessageInfo

func (m *ConfigRequest_V1_System) GetMlsa() *shared.Mlsa {
	if m != nil {
		return m.Mlsa
	}
	return nil
}

func (m *ConfigRequest_V1_System) GetTls() *shared.TLSCredentials {
	if m != nil {
		return m.Tls
	}
	return nil
}

func (m *ConfigRequest_V1_System) GetService() *ConfigRequest_V1_System_Service {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *ConfigRequest_V1_System) GetLog() *ConfigRequest_V1_System_Log {
	if m != nil {
		return m.Log
	}
	return nil
}

type ConfigRequest_V1_System_Service struct {
	Host                 *wrappers.StringValue        `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty" toml:"host,omitempty" mapstructure:"host,omitempty"`
	Port                 *wrappers.Int32Value         `protobuf:"bytes,2,opt,name=port,proto3" json:"port,omitempty" toml:"port,omitempty" mapstructure:"port,omitempty"`
	ExternalFqdn         *wrappers.StringValue        `protobuf:"bytes,3,opt,name=external_fqdn,json=externalFqdn,proto3" json:"external_fqdn,omitempty" toml:"external_fqdn,omitempty" mapstructure:"external_fqdn,omitempty"`
	GrpcPort             *wrappers.Int32Value         `protobuf:"bytes,4,opt,name=grpc_port,json=grpcPort,proto3" json:"grpc_port,omitempty" toml:"grpc_port,omitempty" mapstructure:"grpc_port,omitempty"`
	TrialLicenseUrl      *wrappers.StringValue        `protobuf:"bytes,5,opt,name=trial_license_url,json=trialLicenseUrl,proto3" json:"trial_license_url,omitempty" toml:"trial_license_url,omitempty" mapstructure:"trial_license_url,omitempty"`
	Log                  *ConfigRequest_V1_System_Log `protobuf:"bytes,6,opt,name=log,proto3" json:"log,omitempty" toml:"log,omitempty" mapstructure:"log,omitempty"`
	AuthMiddleware       *wrappers.StringValue        `protobuf:"bytes,7,opt,name=auth_middleware,json=authMiddleware,proto3" json:"auth_middleware,omitempty" toml:"auth_middleware,omitempty" mapstructure:"auth_middleware,omitempty"` // Deprecated: Do not use.
	XXX_NoUnkeyedLiteral struct{}                     `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte                       `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                        `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1_System_Service) Reset()         { *m = ConfigRequest_V1_System_Service{} }
func (m *ConfigRequest_V1_System_Service) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1_System_Service) ProtoMessage()    {}
func (*ConfigRequest_V1_System_Service) Descriptor() ([]byte, []int) {
	return fileDescriptor_9be4cf630cc5f57c, []int{0, 0, 0, 0}
}

func (m *ConfigRequest_V1_System_Service) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest_V1_System_Service.Unmarshal(m, b)
}
func (m *ConfigRequest_V1_System_Service) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest_V1_System_Service.Marshal(b, m, deterministic)
}
func (m *ConfigRequest_V1_System_Service) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest_V1_System_Service.Merge(m, src)
}
func (m *ConfigRequest_V1_System_Service) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest_V1_System_Service.Size(m)
}
func (m *ConfigRequest_V1_System_Service) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest_V1_System_Service.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest_V1_System_Service proto.InternalMessageInfo

func (m *ConfigRequest_V1_System_Service) GetHost() *wrappers.StringValue {
	if m != nil {
		return m.Host
	}
	return nil
}

func (m *ConfigRequest_V1_System_Service) GetPort() *wrappers.Int32Value {
	if m != nil {
		return m.Port
	}
	return nil
}

func (m *ConfigRequest_V1_System_Service) GetExternalFqdn() *wrappers.StringValue {
	if m != nil {
		return m.ExternalFqdn
	}
	return nil
}

func (m *ConfigRequest_V1_System_Service) GetGrpcPort() *wrappers.Int32Value {
	if m != nil {
		return m.GrpcPort
	}
	return nil
}

func (m *ConfigRequest_V1_System_Service) GetTrialLicenseUrl() *wrappers.StringValue {
	if m != nil {
		return m.TrialLicenseUrl
	}
	return nil
}

func (m *ConfigRequest_V1_System_Service) GetLog() *ConfigRequest_V1_System_Log {
	if m != nil {
		return m.Log
	}
	return nil
}

// Deprecated: Do not use.
func (m *ConfigRequest_V1_System_Service) GetAuthMiddleware() *wrappers.StringValue {
	if m != nil {
		return m.AuthMiddleware
	}
	return nil
}

type ConfigRequest_V1_System_Log struct {
	Level                *wrappers.StringValue `protobuf:"bytes,1,opt,name=level,proto3" json:"level,omitempty" toml:"level,omitempty" mapstructure:"level,omitempty"`
	Format               *wrappers.StringValue `protobuf:"bytes,2,opt,name=format,proto3" json:"format,omitempty" toml:"format,omitempty" mapstructure:"format,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte                `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                 `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1_System_Log) Reset()         { *m = ConfigRequest_V1_System_Log{} }
func (m *ConfigRequest_V1_System_Log) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1_System_Log) ProtoMessage()    {}
func (*ConfigRequest_V1_System_Log) Descriptor() ([]byte, []int) {
	return fileDescriptor_9be4cf630cc5f57c, []int{0, 0, 0, 1}
}

func (m *ConfigRequest_V1_System_Log) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest_V1_System_Log.Unmarshal(m, b)
}
func (m *ConfigRequest_V1_System_Log) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest_V1_System_Log.Marshal(b, m, deterministic)
}
func (m *ConfigRequest_V1_System_Log) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest_V1_System_Log.Merge(m, src)
}
func (m *ConfigRequest_V1_System_Log) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest_V1_System_Log.Size(m)
}
func (m *ConfigRequest_V1_System_Log) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest_V1_System_Log.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest_V1_System_Log proto.InternalMessageInfo

func (m *ConfigRequest_V1_System_Log) GetLevel() *wrappers.StringValue {
	if m != nil {
		return m.Level
	}
	return nil
}

func (m *ConfigRequest_V1_System_Log) GetFormat() *wrappers.StringValue {
	if m != nil {
		return m.Format
	}
	return nil
}

type ConfigRequest_V1_Service struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1_Service) Reset()         { *m = ConfigRequest_V1_Service{} }
func (m *ConfigRequest_V1_Service) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1_Service) ProtoMessage()    {}
func (*ConfigRequest_V1_Service) Descriptor() ([]byte, []int) {
	return fileDescriptor_9be4cf630cc5f57c, []int{0, 0, 1}
}

func (m *ConfigRequest_V1_Service) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest_V1_Service.Unmarshal(m, b)
}
func (m *ConfigRequest_V1_Service) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest_V1_Service.Marshal(b, m, deterministic)
}
func (m *ConfigRequest_V1_Service) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest_V1_Service.Merge(m, src)
}
func (m *ConfigRequest_V1_Service) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest_V1_Service.Size(m)
}
func (m *ConfigRequest_V1_Service) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest_V1_Service.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest_V1_Service proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ConfigRequest)(nil), "chef.automate.api.config.ConfigRequest")
	proto.RegisterType((*ConfigRequest_V1)(nil), "chef.automate.api.config.ConfigRequest.V1")
	proto.RegisterType((*ConfigRequest_V1_System)(nil), "chef.automate.api.config.ConfigRequest.V1.System")
	proto.RegisterType((*ConfigRequest_V1_System_Service)(nil), "chef.automate.api.config.ConfigRequest.V1.System.Service")
	proto.RegisterType((*ConfigRequest_V1_System_Log)(nil), "chef.automate.api.config.ConfigRequest.V1.System.Log")
	proto.RegisterType((*ConfigRequest_V1_Service)(nil), "chef.automate.api.config.ConfigRequest.V1.Service")
}

func init() {
	proto.RegisterFile("api/config/gateway/config_request.proto", fileDescriptor_9be4cf630cc5f57c)
}

var fileDescriptor_9be4cf630cc5f57c = []byte{
	// 634 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0x4d, 0x6f, 0xd3, 0x30,
	0x18, 0xc7, 0xd5, 0x26, 0xeb, 0x8b, 0x61, 0xb4, 0x58, 0x08, 0x45, 0x19, 0x9a, 0x26, 0x2e, 0xa0,
	0xa1, 0x26, 0xb4, 0x1b, 0x07, 0x26, 0x2e, 0x6c, 0xbc, 0x6d, 0xea, 0xa4, 0x29, 0x85, 0x1d, 0xb8,
	0x54, 0x6e, 0xfa, 0x34, 0x89, 0xe4, 0xd8, 0x99, 0xed, 0x74, 0xec, 0x23, 0x70, 0xe5, 0x43, 0xf0,
	0x41, 0xf6, 0x09, 0xe0, 0xb3, 0xec, 0xc8, 0x05, 0x39, 0x71, 0xca, 0xcb, 0xc4, 0xb6, 0x8a, 0x4b,
	0x55, 0xcb, 0xcf, 0xef, 0xd7, 0xbf, 0x9f, 0xc7, 0x35, 0x7a, 0x44, 0xb2, 0xc4, 0x0f, 0x39, 0x9b,
	0x25, 0x91, 0x1f, 0x11, 0x05, 0xa7, 0xe4, 0xcc, 0x2c, 0xc7, 0x02, 0x4e, 0x72, 0x90, 0xca, 0xcb,
	0x04, 0x57, 0x1c, 0x3b, 0x61, 0x0c, 0x33, 0x8f, 0xe4, 0x8a, 0xa7, 0x44, 0x81, 0x47, 0xb2, 0xc4,
	0x2b, 0xeb, 0xdc, 0xf5, 0xdf, 0x14, 0x32, 0x26, 0x02, 0xa6, 0x7e, 0x44, 0xf9, 0x84, 0xd0, 0x92,
	0x74, 0xd7, 0x2e, 0xef, 0x2b, 0x2a, 0xcd, 0xe6, 0x41, 0xc8, 0xd3, 0x8c, 0x33, 0x60, 0x4a, 0xfa,
	0x95, 0xbc, 0x17, 0x89, 0x2c, 0xf4, 0x8b, 0xfd, 0xb0, 0x17, 0x01, 0xeb, 0x91, 0x41, 0xcf, 0xf0,
	0x5a, 0x45, 0x06, 0x7a, 0xe1, 0x13, 0xc6, 0xb8, 0x22, 0x2a, 0xe1, 0xac, 0x72, 0xad, 0x47, 0x9c,
	0x47, 0x14, 0x4a, 0x72, 0x92, 0xcf, 0xfc, 0x53, 0x41, 0xb2, 0x0c, 0x84, 0xd9, 0x7f, 0xf8, 0xa5,
	0x8d, 0x56, 0xf7, 0x0a, 0x4f, 0x50, 0x1e, 0x0d, 0xef, 0xa0, 0xfa, 0xbc, 0xef, 0x58, 0x1b, 0xb5,
	0xc7, 0xb7, 0x06, 0x9b, 0xde, 0xbf, 0x4e, 0xe8, 0xfd, 0x01, 0x79, 0xc7, 0xfd, 0xa0, 0x3e, 0xef,
	0xbb, 0x5f, 0x5b, 0xa8, 0x7e, 0xdc, 0xc7, 0x7b, 0xc8, 0x92, 0x67, 0xd2, 0xa9, 0x15, 0x8e, 0xfe,
	0xcd, 0x1d, 0xde, 0xe8, 0x4c, 0x2a, 0x48, 0x03, 0x4d, 0xe3, 0x57, 0xc8, 0x92, 0xf3, 0xd0, 0xa9,
	0x17, 0x92, 0xc1, 0x32, 0x12, 0x10, 0xf3, 0x24, 0x84, 0x40, 0xe3, 0xee, 0xe7, 0x26, 0x6a, 0x94,
	0x56, 0xbc, 0x8d, 0xec, 0x94, 0x4a, 0x62, 0x62, 0x6d, 0xfc, 0x65, 0x4c, 0xd8, 0x4c, 0x90, 0xca,
	0x79, 0x48, 0x25, 0x09, 0x8a, 0x6a, 0xfc, 0x02, 0x59, 0x8a, 0x4a, 0x13, 0x63, 0xf3, 0x2a, 0xe8,
	0xfd, 0x70, 0xb4, 0x27, 0x60, 0x0a, 0x4c, 0x25, 0x84, 0xca, 0x40, 0x63, 0x78, 0x84, 0x9a, 0xb2,
	0x8c, 0x63, 0x3a, 0xfa, 0x7c, 0xe9, 0x6e, 0x2c, 0xce, 0x53, 0x99, 0xf0, 0x5b, 0x64, 0x51, 0x1e,
	0x39, 0x76, 0x21, 0x7c, 0xb6, 0xbc, 0x70, 0xc8, 0xa3, 0x40, 0x1b, 0xdc, 0x1f, 0x16, 0x6a, 0x1a,
	0x3b, 0x7e, 0x8a, 0xec, 0x98, 0x4b, 0x65, 0xba, 0xf3, 0xc0, 0x2b, 0xef, 0x8d, 0x57, 0xdd, 0x1b,
	0x6f, 0xa4, 0x44, 0xc2, 0xa2, 0x63, 0x42, 0x73, 0x08, 0x8a, 0x4a, 0xfc, 0x1a, 0xd9, 0x19, 0x17,
	0xca, 0xb4, 0x66, 0xed, 0x12, 0xb1, 0xcf, 0xd4, 0xd6, 0xa0, 0x00, 0x76, 0xef, 0x9d, 0x5f, 0x38,
	0x5d, 0x64, 0xc7, 0x4a, 0x65, 0xdd, 0x6f, 0x1d, 0x77, 0x45, 0x7f, 0x91, 0x41, 0x81, 0xe3, 0x97,
	0x68, 0x15, 0x3e, 0x29, 0x10, 0x8c, 0xd0, 0xf1, 0xec, 0x64, 0xca, 0x4c, 0xa3, 0xae, 0x4e, 0x70,
	0xbb, 0x42, 0xde, 0x9c, 0x4c, 0x19, 0x3e, 0x42, 0x6d, 0xfd, 0xf7, 0x18, 0x17, 0x71, 0xec, 0xeb,
	0xe3, 0xdc, 0x3f, 0xbf, 0x70, 0xf0, 0x62, 0x32, 0xdd, 0xef, 0x1d, 0xd7, 0xd6, 0x7c, 0xd0, 0xd2,
	0x9f, 0x47, 0x3a, 0xd4, 0x3b, 0x74, 0x57, 0x89, 0x84, 0xd0, 0x31, 0x4d, 0x42, 0x60, 0x12, 0xc6,
	0xb9, 0xa0, 0xce, 0xca, 0x0d, 0x82, 0x75, 0x0a, 0x6c, 0x58, 0x52, 0x1f, 0x04, 0xad, 0x86, 0xd5,
	0xf8, 0xdf, 0x61, 0xe1, 0x7d, 0xd4, 0x21, 0xb9, 0x8a, 0xc7, 0x69, 0x32, 0x9d, 0x52, 0x38, 0x25,
	0x02, 0x9c, 0xe6, 0xf5, 0x81, 0x76, 0xeb, 0x4e, 0x2d, 0xb8, 0xa3, 0xc1, 0xc3, 0x05, 0x77, 0x60,
	0xb7, 0xda, 0x5d, 0xe4, 0x72, 0x64, 0x0d, 0x79, 0x84, 0x07, 0x68, 0x85, 0xc2, 0x1c, 0xe8, 0x8d,
	0x26, 0x5f, 0x96, 0xe2, 0x6d, 0xd4, 0x98, 0x71, 0x91, 0x92, 0x6a, 0xf8, 0x57, 0x43, 0xa6, 0xd6,
	0x6d, 0x2f, 0x6e, 0xdb, 0x4e, 0x39, 0x81, 0xee, 0xaf, 0xd7, 0xad, 0x7c, 0x66, 0x0f, 0xec, 0x56,
	0xad, 0x6b, 0xed, 0xf6, 0x3e, 0x3e, 0x89, 0x12, 0x15, 0xe7, 0x13, 0x2f, 0xe4, 0xa9, 0xaf, 0x5b,
	0xb6, 0x78, 0x07, 0xfd, 0xcb, 0x6f, 0xf3, 0xa4, 0x51, 0xfc, 0xea, 0xd6, 0xcf, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xf1, 0x6a, 0x05, 0x88, 0xb8, 0x05, 0x00, 0x00,
}
