// Code generated by protoc-gen-go. DO NOT EDIT.
// source: components/automate-gateway/api/iam/v2/response/teams.proto

package response

import (
	fmt "fmt"
	common "github.com/chef/automate/components/automate-gateway/api/iam/v2/common"
	proto "github.com/golang/protobuf/proto"
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

type GetTeamResp struct {
	Team                 *common.Team `protobuf:"bytes,1,opt,name=team,proto3" json:"team,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GetTeamResp) Reset()         { *m = GetTeamResp{} }
func (m *GetTeamResp) String() string { return proto.CompactTextString(m) }
func (*GetTeamResp) ProtoMessage()    {}
func (*GetTeamResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a0fc0626d013215, []int{0}
}

func (m *GetTeamResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTeamResp.Unmarshal(m, b)
}
func (m *GetTeamResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTeamResp.Marshal(b, m, deterministic)
}
func (m *GetTeamResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTeamResp.Merge(m, src)
}
func (m *GetTeamResp) XXX_Size() int {
	return xxx_messageInfo_GetTeamResp.Size(m)
}
func (m *GetTeamResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTeamResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetTeamResp proto.InternalMessageInfo

func (m *GetTeamResp) GetTeam() *common.Team {
	if m != nil {
		return m.Team
	}
	return nil
}

type ListTeamsResp struct {
	Teams                []*common.Team `protobuf:"bytes,1,rep,name=teams,proto3" json:"teams,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ListTeamsResp) Reset()         { *m = ListTeamsResp{} }
func (m *ListTeamsResp) String() string { return proto.CompactTextString(m) }
func (*ListTeamsResp) ProtoMessage()    {}
func (*ListTeamsResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a0fc0626d013215, []int{1}
}

func (m *ListTeamsResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTeamsResp.Unmarshal(m, b)
}
func (m *ListTeamsResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTeamsResp.Marshal(b, m, deterministic)
}
func (m *ListTeamsResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTeamsResp.Merge(m, src)
}
func (m *ListTeamsResp) XXX_Size() int {
	return xxx_messageInfo_ListTeamsResp.Size(m)
}
func (m *ListTeamsResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTeamsResp.DiscardUnknown(m)
}

var xxx_messageInfo_ListTeamsResp proto.InternalMessageInfo

func (m *ListTeamsResp) GetTeams() []*common.Team {
	if m != nil {
		return m.Teams
	}
	return nil
}

type CreateTeamResp struct {
	Team                 *common.Team `protobuf:"bytes,1,opt,name=team,proto3" json:"team,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *CreateTeamResp) Reset()         { *m = CreateTeamResp{} }
func (m *CreateTeamResp) String() string { return proto.CompactTextString(m) }
func (*CreateTeamResp) ProtoMessage()    {}
func (*CreateTeamResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a0fc0626d013215, []int{2}
}

func (m *CreateTeamResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTeamResp.Unmarshal(m, b)
}
func (m *CreateTeamResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTeamResp.Marshal(b, m, deterministic)
}
func (m *CreateTeamResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTeamResp.Merge(m, src)
}
func (m *CreateTeamResp) XXX_Size() int {
	return xxx_messageInfo_CreateTeamResp.Size(m)
}
func (m *CreateTeamResp) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTeamResp.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTeamResp proto.InternalMessageInfo

func (m *CreateTeamResp) GetTeam() *common.Team {
	if m != nil {
		return m.Team
	}
	return nil
}

type UpdateTeamResp struct {
	Team                 *common.Team `protobuf:"bytes,1,opt,name=team,proto3" json:"team,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *UpdateTeamResp) Reset()         { *m = UpdateTeamResp{} }
func (m *UpdateTeamResp) String() string { return proto.CompactTextString(m) }
func (*UpdateTeamResp) ProtoMessage()    {}
func (*UpdateTeamResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a0fc0626d013215, []int{3}
}

func (m *UpdateTeamResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateTeamResp.Unmarshal(m, b)
}
func (m *UpdateTeamResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateTeamResp.Marshal(b, m, deterministic)
}
func (m *UpdateTeamResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateTeamResp.Merge(m, src)
}
func (m *UpdateTeamResp) XXX_Size() int {
	return xxx_messageInfo_UpdateTeamResp.Size(m)
}
func (m *UpdateTeamResp) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateTeamResp.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateTeamResp proto.InternalMessageInfo

func (m *UpdateTeamResp) GetTeam() *common.Team {
	if m != nil {
		return m.Team
	}
	return nil
}

type DeleteTeamResp struct {
	Team                 *common.Team `protobuf:"bytes,1,opt,name=team,proto3" json:"team,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *DeleteTeamResp) Reset()         { *m = DeleteTeamResp{} }
func (m *DeleteTeamResp) String() string { return proto.CompactTextString(m) }
func (*DeleteTeamResp) ProtoMessage()    {}
func (*DeleteTeamResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a0fc0626d013215, []int{4}
}

func (m *DeleteTeamResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteTeamResp.Unmarshal(m, b)
}
func (m *DeleteTeamResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteTeamResp.Marshal(b, m, deterministic)
}
func (m *DeleteTeamResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteTeamResp.Merge(m, src)
}
func (m *DeleteTeamResp) XXX_Size() int {
	return xxx_messageInfo_DeleteTeamResp.Size(m)
}
func (m *DeleteTeamResp) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteTeamResp.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteTeamResp proto.InternalMessageInfo

func (m *DeleteTeamResp) GetTeam() *common.Team {
	if m != nil {
		return m.Team
	}
	return nil
}

type AddTeamMembersResp struct {
	UserIds              []string `protobuf:"bytes,1,rep,name=user_ids,json=userIds,proto3" json:"user_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddTeamMembersResp) Reset()         { *m = AddTeamMembersResp{} }
func (m *AddTeamMembersResp) String() string { return proto.CompactTextString(m) }
func (*AddTeamMembersResp) ProtoMessage()    {}
func (*AddTeamMembersResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a0fc0626d013215, []int{5}
}

func (m *AddTeamMembersResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddTeamMembersResp.Unmarshal(m, b)
}
func (m *AddTeamMembersResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddTeamMembersResp.Marshal(b, m, deterministic)
}
func (m *AddTeamMembersResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddTeamMembersResp.Merge(m, src)
}
func (m *AddTeamMembersResp) XXX_Size() int {
	return xxx_messageInfo_AddTeamMembersResp.Size(m)
}
func (m *AddTeamMembersResp) XXX_DiscardUnknown() {
	xxx_messageInfo_AddTeamMembersResp.DiscardUnknown(m)
}

var xxx_messageInfo_AddTeamMembersResp proto.InternalMessageInfo

func (m *AddTeamMembersResp) GetUserIds() []string {
	if m != nil {
		return m.UserIds
	}
	return nil
}

type GetTeamMembershipResp struct {
	UserIds              []string `protobuf:"bytes,1,rep,name=user_ids,json=userIds,proto3" json:"user_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTeamMembershipResp) Reset()         { *m = GetTeamMembershipResp{} }
func (m *GetTeamMembershipResp) String() string { return proto.CompactTextString(m) }
func (*GetTeamMembershipResp) ProtoMessage()    {}
func (*GetTeamMembershipResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a0fc0626d013215, []int{6}
}

func (m *GetTeamMembershipResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTeamMembershipResp.Unmarshal(m, b)
}
func (m *GetTeamMembershipResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTeamMembershipResp.Marshal(b, m, deterministic)
}
func (m *GetTeamMembershipResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTeamMembershipResp.Merge(m, src)
}
func (m *GetTeamMembershipResp) XXX_Size() int {
	return xxx_messageInfo_GetTeamMembershipResp.Size(m)
}
func (m *GetTeamMembershipResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTeamMembershipResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetTeamMembershipResp proto.InternalMessageInfo

func (m *GetTeamMembershipResp) GetUserIds() []string {
	if m != nil {
		return m.UserIds
	}
	return nil
}

type RemoveTeamMembersResp struct {
	UserIds              []string `protobuf:"bytes,1,rep,name=user_ids,json=userIds,proto3" json:"user_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveTeamMembersResp) Reset()         { *m = RemoveTeamMembersResp{} }
func (m *RemoveTeamMembersResp) String() string { return proto.CompactTextString(m) }
func (*RemoveTeamMembersResp) ProtoMessage()    {}
func (*RemoveTeamMembersResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a0fc0626d013215, []int{7}
}

func (m *RemoveTeamMembersResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveTeamMembersResp.Unmarshal(m, b)
}
func (m *RemoveTeamMembersResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveTeamMembersResp.Marshal(b, m, deterministic)
}
func (m *RemoveTeamMembersResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveTeamMembersResp.Merge(m, src)
}
func (m *RemoveTeamMembersResp) XXX_Size() int {
	return xxx_messageInfo_RemoveTeamMembersResp.Size(m)
}
func (m *RemoveTeamMembersResp) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveTeamMembersResp.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveTeamMembersResp proto.InternalMessageInfo

func (m *RemoveTeamMembersResp) GetUserIds() []string {
	if m != nil {
		return m.UserIds
	}
	return nil
}

type GetTeamsForMemberResp struct {
	Teams                []*common.Team `protobuf:"bytes,1,rep,name=teams,proto3" json:"teams,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GetTeamsForMemberResp) Reset()         { *m = GetTeamsForMemberResp{} }
func (m *GetTeamsForMemberResp) String() string { return proto.CompactTextString(m) }
func (*GetTeamsForMemberResp) ProtoMessage()    {}
func (*GetTeamsForMemberResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a0fc0626d013215, []int{8}
}

func (m *GetTeamsForMemberResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTeamsForMemberResp.Unmarshal(m, b)
}
func (m *GetTeamsForMemberResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTeamsForMemberResp.Marshal(b, m, deterministic)
}
func (m *GetTeamsForMemberResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTeamsForMemberResp.Merge(m, src)
}
func (m *GetTeamsForMemberResp) XXX_Size() int {
	return xxx_messageInfo_GetTeamsForMemberResp.Size(m)
}
func (m *GetTeamsForMemberResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTeamsForMemberResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetTeamsForMemberResp proto.InternalMessageInfo

func (m *GetTeamsForMemberResp) GetTeams() []*common.Team {
	if m != nil {
		return m.Teams
	}
	return nil
}

type ApplyV2DataMigrationsResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ApplyV2DataMigrationsResp) Reset()         { *m = ApplyV2DataMigrationsResp{} }
func (m *ApplyV2DataMigrationsResp) String() string { return proto.CompactTextString(m) }
func (*ApplyV2DataMigrationsResp) ProtoMessage()    {}
func (*ApplyV2DataMigrationsResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a0fc0626d013215, []int{9}
}

func (m *ApplyV2DataMigrationsResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApplyV2DataMigrationsResp.Unmarshal(m, b)
}
func (m *ApplyV2DataMigrationsResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApplyV2DataMigrationsResp.Marshal(b, m, deterministic)
}
func (m *ApplyV2DataMigrationsResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplyV2DataMigrationsResp.Merge(m, src)
}
func (m *ApplyV2DataMigrationsResp) XXX_Size() int {
	return xxx_messageInfo_ApplyV2DataMigrationsResp.Size(m)
}
func (m *ApplyV2DataMigrationsResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplyV2DataMigrationsResp.DiscardUnknown(m)
}

var xxx_messageInfo_ApplyV2DataMigrationsResp proto.InternalMessageInfo

type ResetAllTeamProjectsResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResetAllTeamProjectsResp) Reset()         { *m = ResetAllTeamProjectsResp{} }
func (m *ResetAllTeamProjectsResp) String() string { return proto.CompactTextString(m) }
func (*ResetAllTeamProjectsResp) ProtoMessage()    {}
func (*ResetAllTeamProjectsResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a0fc0626d013215, []int{10}
}

func (m *ResetAllTeamProjectsResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResetAllTeamProjectsResp.Unmarshal(m, b)
}
func (m *ResetAllTeamProjectsResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResetAllTeamProjectsResp.Marshal(b, m, deterministic)
}
func (m *ResetAllTeamProjectsResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResetAllTeamProjectsResp.Merge(m, src)
}
func (m *ResetAllTeamProjectsResp) XXX_Size() int {
	return xxx_messageInfo_ResetAllTeamProjectsResp.Size(m)
}
func (m *ResetAllTeamProjectsResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ResetAllTeamProjectsResp.DiscardUnknown(m)
}

var xxx_messageInfo_ResetAllTeamProjectsResp proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GetTeamResp)(nil), "chef.automate.api.iam.v2.GetTeamResp")
	proto.RegisterType((*ListTeamsResp)(nil), "chef.automate.api.iam.v2.ListTeamsResp")
	proto.RegisterType((*CreateTeamResp)(nil), "chef.automate.api.iam.v2.CreateTeamResp")
	proto.RegisterType((*UpdateTeamResp)(nil), "chef.automate.api.iam.v2.UpdateTeamResp")
	proto.RegisterType((*DeleteTeamResp)(nil), "chef.automate.api.iam.v2.DeleteTeamResp")
	proto.RegisterType((*AddTeamMembersResp)(nil), "chef.automate.api.iam.v2.AddTeamMembersResp")
	proto.RegisterType((*GetTeamMembershipResp)(nil), "chef.automate.api.iam.v2.GetTeamMembershipResp")
	proto.RegisterType((*RemoveTeamMembersResp)(nil), "chef.automate.api.iam.v2.RemoveTeamMembersResp")
	proto.RegisterType((*GetTeamsForMemberResp)(nil), "chef.automate.api.iam.v2.GetTeamsForMemberResp")
	proto.RegisterType((*ApplyV2DataMigrationsResp)(nil), "chef.automate.api.iam.v2.ApplyV2DataMigrationsResp")
	proto.RegisterType((*ResetAllTeamProjectsResp)(nil), "chef.automate.api.iam.v2.ResetAllTeamProjectsResp")
}

func init() {
	proto.RegisterFile("components/automate-gateway/api/iam/v2/response/teams.proto", fileDescriptor_6a0fc0626d013215)
}

var fileDescriptor_6a0fc0626d013215 = []byte{
	// 329 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xcb, 0x4b, 0xc3, 0x40,
	0x10, 0xc6, 0x29, 0xbe, 0xb7, 0xd8, 0x43, 0x41, 0x48, 0x2b, 0x48, 0xc9, 0xc9, 0x8b, 0xbb, 0x10,
	0xbd, 0x88, 0xa7, 0x68, 0x7c, 0x62, 0x41, 0x82, 0x7a, 0xf0, 0x22, 0xd3, 0x64, 0x6c, 0x57, 0xb2,
	0xd9, 0x65, 0x77, 0x5a, 0xe9, 0x7f, 0x2f, 0xd9, 0xd4, 0xe2, 0x45, 0x68, 0x89, 0xc7, 0xdd, 0x99,
	0xef, 0xf7, 0x0d, 0xf3, 0x60, 0x17, 0x99, 0x56, 0x46, 0x97, 0x58, 0x92, 0x13, 0x30, 0x25, 0xad,
	0x80, 0xf0, 0x64, 0x0c, 0x84, 0x5f, 0x30, 0x17, 0x60, 0xa4, 0x90, 0xa0, 0xc4, 0x2c, 0x12, 0x16,
	0x9d, 0xd1, 0xa5, 0x43, 0x41, 0x08, 0xca, 0x71, 0x63, 0x35, 0xe9, 0x6e, 0x90, 0x4d, 0xf0, 0x83,
	0xff, 0xc8, 0x38, 0x18, 0xc9, 0x25, 0x28, 0x3e, 0x8b, 0xfa, 0xe7, 0x2b, 0x62, 0x33, 0xad, 0x94,
	0x2e, 0x7f, 0x43, 0xc3, 0x98, 0xb5, 0x6f, 0x91, 0x9e, 0x11, 0x54, 0x8a, 0xce, 0x74, 0x23, 0xb6,
	0x59, 0x45, 0x83, 0xd6, 0xa0, 0x75, 0xdc, 0x8e, 0x8e, 0xf8, 0x5f, 0x96, 0xdc, 0x2b, 0x7c, 0x6e,
	0x78, 0xcd, 0xf6, 0x1f, 0xa5, 0xf3, 0x0c, 0xe7, 0x21, 0x67, 0x6c, 0xcb, 0x5b, 0x04, 0xad, 0xc1,
	0xc6, 0x0a, 0x94, 0x3a, 0x39, 0x4c, 0x58, 0xe7, 0xca, 0x22, 0x10, 0x36, 0x2a, 0x26, 0x61, 0x9d,
	0x17, 0x93, 0xff, 0x03, 0x25, 0xc1, 0x02, 0x1b, 0x52, 0x04, 0xeb, 0xc6, 0x79, 0x5e, 0x7d, 0x0c,
	0x51, 0x8d, 0xd0, 0xd6, 0xdd, 0xe9, 0xb1, 0xdd, 0xa9, 0x43, 0xfb, 0x2e, 0xf3, 0xba, 0x41, 0x7b,
	0xe9, 0x4e, 0xf5, 0xbe, 0xcf, 0x5d, 0x18, 0xb1, 0x83, 0xc5, 0x30, 0x16, 0x82, 0x89, 0x34, 0x2b,
	0x68, 0x52, 0x54, 0x7a, 0x86, 0x6b, 0xf8, 0x0c, 0x97, 0x3e, 0xee, 0x46, 0xdb, 0x5a, 0xd4, 0x60,
	0x72, 0x87, 0xac, 0x17, 0x1b, 0x53, 0xcc, 0x5f, 0xa3, 0x04, 0x08, 0x86, 0x72, 0x6c, 0x81, 0xa4,
	0x2e, 0x7d, 0x19, 0x61, 0x9f, 0x05, 0x29, 0x3a, 0xa4, 0xb8, 0x28, 0x2a, 0xcd, 0x93, 0xd5, 0x9f,
	0x98, 0x91, 0x8f, 0x5d, 0x3e, 0xbc, 0xdd, 0x8d, 0x25, 0x4d, 0xa6, 0x23, 0x9e, 0x69, 0x25, 0x2a,
	0xaf, 0xe5, 0xfa, 0x8a, 0x35, 0x2f, 0x65, 0xb4, 0xed, 0xf7, 0xf9, 0xf4, 0x3b, 0x00, 0x00, 0xff,
	0xff, 0x13, 0xb6, 0x93, 0xd7, 0x63, 0x03, 0x00, 0x00,
}
