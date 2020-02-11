// Code generated by protoc-gen-go. DO NOT EDIT.
// source: components/automate-gateway/api/iam/v2/rules.proto

package v2

import (
	context "context"
	fmt "fmt"
	request "github.com/chef/automate/components/automate-gateway/api/iam/v2/request"
	response "github.com/chef/automate/components/automate-gateway/api/iam/v2/response"
	_ "github.com/chef/automate/components/automate-grpc/protoc-gen-policy/api"
	_ "github.com/chef/automate/components/automate-grpc/protoc-gen-policy/iam"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

func init() {
	proto.RegisterFile("components/automate-gateway/api/iam/v2/rules.proto", fileDescriptor_28c6f49a8332221c)
}

var fileDescriptor_28c6f49a8332221c = []byte{
	// 695 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x96, 0xbf, 0x6f, 0xd4, 0x3c,
	0x18, 0xc7, 0x75, 0xf7, 0xaa, 0x7d, 0xdf, 0xd7, 0xef, 0x2b, 0xb5, 0xf7, 0x5c, 0x11, 0x69, 0x5a,
	0x09, 0x11, 0x40, 0xad, 0x4e, 0xbd, 0xa4, 0x3d, 0xb6, 0x63, 0xa0, 0xa5, 0x08, 0x16, 0x06, 0xd4,
	0x0a, 0x86, 0x2c, 0xe0, 0xa6, 0x26, 0x0d, 0xba, 0x8b, 0xdd, 0xd8, 0x29, 0xaa, 0xaa, 0x0e, 0x1c,
	0xdb, 0x6d, 0x08, 0x24, 0x06, 0x36, 0x06, 0x26, 0xd6, 0xac, 0x88, 0x81, 0x89, 0x95, 0x7f, 0x81,
	0xbf, 0x80, 0x85, 0x15, 0xd9, 0xb9, 0x1f, 0xb9, 0x36, 0x47, 0x9c, 0xa5, 0x5b, 0x12, 0x7f, 0x1e,
	0xfb, 0xf9, 0x7e, 0x1c, 0x4b, 0x46, 0x2d, 0x8f, 0x76, 0x19, 0x0d, 0x49, 0x28, 0xb8, 0x83, 0x63,
	0x41, 0xbb, 0x58, 0x90, 0xa6, 0x8f, 0x05, 0x79, 0x81, 0x8f, 0x1d, 0xcc, 0x02, 0x27, 0xc0, 0x5d,
	0xe7, 0xa8, 0xe5, 0x44, 0x71, 0x87, 0x70, 0x9b, 0x45, 0x54, 0x50, 0x30, 0xbc, 0x03, 0xf2, 0xcc,
	0x1e, 0xd2, 0x36, 0x66, 0x81, 0x1d, 0xe0, 0xae, 0x7d, 0xd4, 0x32, 0x97, 0x7d, 0x4a, 0xfd, 0x0e,
	0x51, 0x85, 0x38, 0x0c, 0xa9, 0xc0, 0x22, 0xa0, 0xe1, 0xa0, 0xce, 0x6c, 0xeb, 0xae, 0x45, 0x0e,
	0x63, 0xc2, 0x45, 0x76, 0x4d, 0xf3, 0x96, 0x76, 0x2d, 0x67, 0x34, 0xe4, 0x64, 0xa2, 0x78, 0x33,
	0xb7, 0x38, 0x62, 0x9e, 0xa3, 0xc6, 0xbd, 0xa6, 0x4f, 0xc2, 0x26, 0xa3, 0x9d, 0xc0, 0x3b, 0x9e,
	0xd2, 0x7a, 0x99, 0x19, 0x64, 0x27, 0xe7, 0x66, 0x68, 0xbd, 0xaf, 0xa1, 0x99, 0x1d, 0xd9, 0x13,
	0x7c, 0xa8, 0x22, 0xb4, 0x1d, 0x11, 0x2c, 0x88, 0x7c, 0x87, 0x15, 0x7b, 0x9a, 0x4e, 0x7b, 0x4c,
	0xed, 0x90, 0x43, 0x73, 0x55, 0x0f, 0xe4, 0xcc, 0xfa, 0x52, 0xe9, 0x25, 0xc6, 0x32, 0x32, 0x71,
	0x2c, 0x0e, 0xda, 0x2c, 0xa2, 0xcf, 0x89, 0x27, 0x78, 0xfb, 0x64, 0xf0, 0xf4, 0x24, 0xd8, 0x3f,
	0xed, 0x25, 0xc6, 0x3f, 0x30, 0x1b, 0xb3, 0x7d, 0x2c, 0x48, 0x3f, 0x31, 0x96, 0xd0, 0x62, 0x80,
	0xbb, 0xf9, 0x68, 0x3f, 0x31, 0x2e, 0x41, 0x7d, 0x62, 0x38, 0xad, 0xeb, 0x7d, 0xff, 0xf1, 0xa6,
	0xfa, 0xd8, 0xba, 0x36, 0x54, 0x3f, 0x1c, 0x76, 0xb2, 0xd5, 0xe9, 0x3e, 0xb4, 0x2b, 0x0d, 0x77,
	0xcd, 0x5a, 0x19, 0x90, 0x7b, 0x44, 0xe0, 0x02, 0x1a, 0x3e, 0x55, 0x11, 0x7a, 0xa4, 0x96, 0x29,
	0x92, 0x34, 0xa6, 0x0a, 0x24, 0x65, 0x41, 0xce, 0xac, 0x6f, 0x17, 0x27, 0xe9, 0xa9, 0xb9, 0xaa,
	0x21, 0xc9, 0x39, 0x09, 0xf6, 0x4f, 0xa5, 0xa9, 0x0d, 0x73, 0x4d, 0xd3, 0xd4, 0xb0, 0x04, 0xde,
	0x55, 0xd1, 0xdf, 0xf7, 0x89, 0x50, 0xae, 0xae, 0x4f, 0x57, 0x30, 0x40, 0xa4, 0xa8, 0x1b, 0x1a,
	0x14, 0x67, 0xd6, 0x67, 0x1d, 0x4b, 0x33, 0xf0, 0x97, 0x4f, 0x44, 0xb1, 0x22, 0x80, 0xf9, 0x89,
	0x61, 0x9f, 0x08, 0xe5, 0xc7, 0x05, 0x6d, 0x3f, 0xae, 0x0d, 0xa5, 0xe4, 0xc0, 0xeb, 0x2a, 0xaa,
	0x3f, 0x08, 0xb8, 0x4a, 0xc4, 0xef, 0xd1, 0xe8, 0x61, 0x8a, 0xc1, 0xfa, 0xf4, 0xfc, 0x39, 0xb8,
	0x34, 0xb6, 0x51, 0xb2, 0x82, 0x33, 0xeb, 0xad, 0xb4, 0xb7, 0x80, 0xe0, 0x8c, 0xbd, 0x33, 0xd6,
	0xea, 0xa8, 0x36, 0x69, 0xed, 0xcf, 0xb6, 0x36, 0x61, 0xe9, 0xbc, 0xad, 0x51, 0x6a, 0xf7, 0x2a,
	0x5c, 0xc9, 0x17, 0x34, 0x42, 0xe0, 0x63, 0x15, 0xa1, 0xbb, 0xa4, 0x43, 0x8a, 0x0f, 0xd7, 0x98,
	0x2a, 0x38, 0x5c, 0x59, 0x90, 0x33, 0xeb, 0xeb, 0xc5, 0x1d, 0x2e, 0xb7, 0x51, 0xe2, 0xe7, 0x69,
	0x94, 0xfb, 0x79, 0x7e, 0x56, 0xd0, 0xdc, 0x16, 0x63, 0x9d, 0x63, 0xb5, 0xb9, 0xbb, 0x02, 0x47,
	0x02, 0xd6, 0xa6, 0x4b, 0x38, 0x83, 0x4a, 0x65, 0xcd, 0x12, 0x34, 0x67, 0xd6, 0x4b, 0xe9, 0xed,
	0x7f, 0x84, 0x94, 0x37, 0xd5, 0x48, 0x2f, 0x31, 0xea, 0x50, 0xc3, 0x92, 0x6d, 0xaa, 0x0f, 0x4d,
	0x2e, 0xe9, 0x7e, 0x62, 0xfc, 0x87, 0xfe, 0x95, 0x4e, 0xd4, 0xc7, 0x7e, 0x62, 0xd4, 0x60, 0x6e,
	0xf4, 0xda, 0x56, 0xbc, 0xd2, 0xb3, 0x6e, 0xd5, 0x87, 0x7a, 0x32, 0x93, 0xb8, 0x8b, 0xd6, 0xe5,
	0xac, 0x89, 0xcc, 0x10, 0xfc, 0xaa, 0xa0, 0xf9, 0x71, 0x6f, 0xdb, 0x38, 0xf4, 0x48, 0x07, 0xb4,
	0x72, 0xa4, 0xac, 0x8c, 0x6d, 0x97, 0xc1, 0x39, 0xb3, 0x5e, 0xe5, 0xe5, 0x5e, 0x00, 0xc8, 0xe6,
	0xf6, 0x14, 0x7e, 0x3e, 0xf8, 0xe0, 0x9c, 0xa4, 0xc1, 0x53, 0x28, 0x4d, 0xde, 0xc8, 0x4f, 0xde,
	0xd0, 0x4c, 0xbe, 0x2b, 0xb0, 0x88, 0x39, 0xe8, 0xee, 0xa0, 0x88, 0xb9, 0x76, 0xf2, 0x21, 0xae,
	0x99, 0x9c, 0x2b, 0xbc, 0x20, 0x79, 0x0a, 0xa5, 0xc9, 0x21, 0x3f, 0x39, 0x4c, 0x4b, 0x7e, 0x67,
	0xcb, 0xbd, 0xed, 0x07, 0xe2, 0x20, 0xde, 0xb3, 0x3d, 0xda, 0x75, 0x64, 0x82, 0xd1, 0x35, 0xc7,
	0xd1, 0xbb, 0x79, 0xed, 0xcd, 0xaa, 0x7b, 0xce, 0xcd, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x3e,
	0xe8, 0xf0, 0x31, 0x52, 0x0a, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RulesClient is the client API for Rules service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RulesClient interface {
	CreateRule(ctx context.Context, in *request.CreateRuleReq, opts ...grpc.CallOption) (*response.CreateRuleResp, error)
	UpdateRule(ctx context.Context, in *request.UpdateRuleReq, opts ...grpc.CallOption) (*response.UpdateRuleResp, error)
	GetRule(ctx context.Context, in *request.GetRuleReq, opts ...grpc.CallOption) (*response.GetRuleResp, error)
	ListRulesForProject(ctx context.Context, in *request.ListRulesForProjectReq, opts ...grpc.CallOption) (*response.ListRulesForProjectResp, error)
	DeleteRule(ctx context.Context, in *request.DeleteRuleReq, opts ...grpc.CallOption) (*response.DeleteRuleResp, error)
	ApplyRulesStart(ctx context.Context, in *request.ApplyRulesStartReq, opts ...grpc.CallOption) (*response.ApplyRulesStartResp, error)
	ApplyRulesCancel(ctx context.Context, in *request.ApplyRulesCancelReq, opts ...grpc.CallOption) (*response.ApplyRulesCancelResp, error)
	ApplyRulesStatus(ctx context.Context, in *request.ApplyRulesStatusReq, opts ...grpc.CallOption) (*response.ApplyRulesStatusResp, error)
}

type rulesClient struct {
	cc *grpc.ClientConn
}

func NewRulesClient(cc *grpc.ClientConn) RulesClient {
	return &rulesClient{cc}
}

func (c *rulesClient) CreateRule(ctx context.Context, in *request.CreateRuleReq, opts ...grpc.CallOption) (*response.CreateRuleResp, error) {
	out := new(response.CreateRuleResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.iam.v2.Rules/CreateRule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rulesClient) UpdateRule(ctx context.Context, in *request.UpdateRuleReq, opts ...grpc.CallOption) (*response.UpdateRuleResp, error) {
	out := new(response.UpdateRuleResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.iam.v2.Rules/UpdateRule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rulesClient) GetRule(ctx context.Context, in *request.GetRuleReq, opts ...grpc.CallOption) (*response.GetRuleResp, error) {
	out := new(response.GetRuleResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.iam.v2.Rules/GetRule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rulesClient) ListRulesForProject(ctx context.Context, in *request.ListRulesForProjectReq, opts ...grpc.CallOption) (*response.ListRulesForProjectResp, error) {
	out := new(response.ListRulesForProjectResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.iam.v2.Rules/ListRulesForProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rulesClient) DeleteRule(ctx context.Context, in *request.DeleteRuleReq, opts ...grpc.CallOption) (*response.DeleteRuleResp, error) {
	out := new(response.DeleteRuleResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.iam.v2.Rules/DeleteRule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rulesClient) ApplyRulesStart(ctx context.Context, in *request.ApplyRulesStartReq, opts ...grpc.CallOption) (*response.ApplyRulesStartResp, error) {
	out := new(response.ApplyRulesStartResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.iam.v2.Rules/ApplyRulesStart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rulesClient) ApplyRulesCancel(ctx context.Context, in *request.ApplyRulesCancelReq, opts ...grpc.CallOption) (*response.ApplyRulesCancelResp, error) {
	out := new(response.ApplyRulesCancelResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.iam.v2.Rules/ApplyRulesCancel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rulesClient) ApplyRulesStatus(ctx context.Context, in *request.ApplyRulesStatusReq, opts ...grpc.CallOption) (*response.ApplyRulesStatusResp, error) {
	out := new(response.ApplyRulesStatusResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.iam.v2.Rules/ApplyRulesStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RulesServer is the server API for Rules service.
type RulesServer interface {
	CreateRule(context.Context, *request.CreateRuleReq) (*response.CreateRuleResp, error)
	UpdateRule(context.Context, *request.UpdateRuleReq) (*response.UpdateRuleResp, error)
	GetRule(context.Context, *request.GetRuleReq) (*response.GetRuleResp, error)
	ListRulesForProject(context.Context, *request.ListRulesForProjectReq) (*response.ListRulesForProjectResp, error)
	DeleteRule(context.Context, *request.DeleteRuleReq) (*response.DeleteRuleResp, error)
	ApplyRulesStart(context.Context, *request.ApplyRulesStartReq) (*response.ApplyRulesStartResp, error)
	ApplyRulesCancel(context.Context, *request.ApplyRulesCancelReq) (*response.ApplyRulesCancelResp, error)
	ApplyRulesStatus(context.Context, *request.ApplyRulesStatusReq) (*response.ApplyRulesStatusResp, error)
}

// UnimplementedRulesServer can be embedded to have forward compatible implementations.
type UnimplementedRulesServer struct {
}

func (*UnimplementedRulesServer) CreateRule(ctx context.Context, req *request.CreateRuleReq) (*response.CreateRuleResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRule not implemented")
}
func (*UnimplementedRulesServer) UpdateRule(ctx context.Context, req *request.UpdateRuleReq) (*response.UpdateRuleResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRule not implemented")
}
func (*UnimplementedRulesServer) GetRule(ctx context.Context, req *request.GetRuleReq) (*response.GetRuleResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRule not implemented")
}
func (*UnimplementedRulesServer) ListRulesForProject(ctx context.Context, req *request.ListRulesForProjectReq) (*response.ListRulesForProjectResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRulesForProject not implemented")
}
func (*UnimplementedRulesServer) DeleteRule(ctx context.Context, req *request.DeleteRuleReq) (*response.DeleteRuleResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRule not implemented")
}
func (*UnimplementedRulesServer) ApplyRulesStart(ctx context.Context, req *request.ApplyRulesStartReq) (*response.ApplyRulesStartResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApplyRulesStart not implemented")
}
func (*UnimplementedRulesServer) ApplyRulesCancel(ctx context.Context, req *request.ApplyRulesCancelReq) (*response.ApplyRulesCancelResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApplyRulesCancel not implemented")
}
func (*UnimplementedRulesServer) ApplyRulesStatus(ctx context.Context, req *request.ApplyRulesStatusReq) (*response.ApplyRulesStatusResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApplyRulesStatus not implemented")
}

func RegisterRulesServer(s *grpc.Server, srv RulesServer) {
	s.RegisterService(&_Rules_serviceDesc, srv)
}

func _Rules_CreateRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.CreateRuleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RulesServer).CreateRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.iam.v2.Rules/CreateRule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RulesServer).CreateRule(ctx, req.(*request.CreateRuleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rules_UpdateRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.UpdateRuleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RulesServer).UpdateRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.iam.v2.Rules/UpdateRule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RulesServer).UpdateRule(ctx, req.(*request.UpdateRuleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rules_GetRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.GetRuleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RulesServer).GetRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.iam.v2.Rules/GetRule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RulesServer).GetRule(ctx, req.(*request.GetRuleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rules_ListRulesForProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.ListRulesForProjectReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RulesServer).ListRulesForProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.iam.v2.Rules/ListRulesForProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RulesServer).ListRulesForProject(ctx, req.(*request.ListRulesForProjectReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rules_DeleteRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.DeleteRuleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RulesServer).DeleteRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.iam.v2.Rules/DeleteRule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RulesServer).DeleteRule(ctx, req.(*request.DeleteRuleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rules_ApplyRulesStart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.ApplyRulesStartReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RulesServer).ApplyRulesStart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.iam.v2.Rules/ApplyRulesStart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RulesServer).ApplyRulesStart(ctx, req.(*request.ApplyRulesStartReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rules_ApplyRulesCancel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.ApplyRulesCancelReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RulesServer).ApplyRulesCancel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.iam.v2.Rules/ApplyRulesCancel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RulesServer).ApplyRulesCancel(ctx, req.(*request.ApplyRulesCancelReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rules_ApplyRulesStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.ApplyRulesStatusReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RulesServer).ApplyRulesStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.iam.v2.Rules/ApplyRulesStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RulesServer).ApplyRulesStatus(ctx, req.(*request.ApplyRulesStatusReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Rules_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chef.automate.api.iam.v2.Rules",
	HandlerType: (*RulesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRule",
			Handler:    _Rules_CreateRule_Handler,
		},
		{
			MethodName: "UpdateRule",
			Handler:    _Rules_UpdateRule_Handler,
		},
		{
			MethodName: "GetRule",
			Handler:    _Rules_GetRule_Handler,
		},
		{
			MethodName: "ListRulesForProject",
			Handler:    _Rules_ListRulesForProject_Handler,
		},
		{
			MethodName: "DeleteRule",
			Handler:    _Rules_DeleteRule_Handler,
		},
		{
			MethodName: "ApplyRulesStart",
			Handler:    _Rules_ApplyRulesStart_Handler,
		},
		{
			MethodName: "ApplyRulesCancel",
			Handler:    _Rules_ApplyRulesCancel_Handler,
		},
		{
			MethodName: "ApplyRulesStatus",
			Handler:    _Rules_ApplyRulesStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "components/automate-gateway/api/iam/v2/rules.proto",
}
