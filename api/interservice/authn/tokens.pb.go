// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/interservice/authn/tokens.proto

package authn

import (
	context "context"
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type CreateTokenReq struct {
	// Match either a completely empty string or a valid id.
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" toml:"id,omitempty" mapstructure:"id,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty" toml:"description,omitempty" mapstructure:"description,omitempty"`
	Active               bool     `protobuf:"varint,3,opt,name=active,proto3" json:"active,omitempty" toml:"active,omitempty" mapstructure:"active,omitempty"`
	Projects             []string `protobuf:"bytes,4,rep,name=projects,proto3" json:"projects,omitempty" toml:"projects,omitempty" mapstructure:"projects,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *CreateTokenReq) Reset()         { *m = CreateTokenReq{} }
func (m *CreateTokenReq) String() string { return proto.CompactTextString(m) }
func (*CreateTokenReq) ProtoMessage()    {}
func (*CreateTokenReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_65ff3d3f7473d71c, []int{0}
}

func (m *CreateTokenReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTokenReq.Unmarshal(m, b)
}
func (m *CreateTokenReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTokenReq.Marshal(b, m, deterministic)
}
func (m *CreateTokenReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTokenReq.Merge(m, src)
}
func (m *CreateTokenReq) XXX_Size() int {
	return xxx_messageInfo_CreateTokenReq.Size(m)
}
func (m *CreateTokenReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTokenReq.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTokenReq proto.InternalMessageInfo

func (m *CreateTokenReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CreateTokenReq) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *CreateTokenReq) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

func (m *CreateTokenReq) GetProjects() []string {
	if m != nil {
		return m.Projects
	}
	return nil
}

type CreateTokenWithValueReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" toml:"id,omitempty" mapstructure:"id,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty" toml:"description,omitempty" mapstructure:"description,omitempty"`
	Active               bool     `protobuf:"varint,3,opt,name=active,proto3" json:"active,omitempty" toml:"active,omitempty" mapstructure:"active,omitempty"`
	Value                string   `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty" toml:"value,omitempty" mapstructure:"value,omitempty"`
	Projects             []string `protobuf:"bytes,5,rep,name=projects,proto3" json:"projects,omitempty" toml:"projects,omitempty" mapstructure:"projects,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *CreateTokenWithValueReq) Reset()         { *m = CreateTokenWithValueReq{} }
func (m *CreateTokenWithValueReq) String() string { return proto.CompactTextString(m) }
func (*CreateTokenWithValueReq) ProtoMessage()    {}
func (*CreateTokenWithValueReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_65ff3d3f7473d71c, []int{1}
}

func (m *CreateTokenWithValueReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTokenWithValueReq.Unmarshal(m, b)
}
func (m *CreateTokenWithValueReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTokenWithValueReq.Marshal(b, m, deterministic)
}
func (m *CreateTokenWithValueReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTokenWithValueReq.Merge(m, src)
}
func (m *CreateTokenWithValueReq) XXX_Size() int {
	return xxx_messageInfo_CreateTokenWithValueReq.Size(m)
}
func (m *CreateTokenWithValueReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTokenWithValueReq.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTokenWithValueReq proto.InternalMessageInfo

func (m *CreateTokenWithValueReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CreateTokenWithValueReq) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *CreateTokenWithValueReq) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

func (m *CreateTokenWithValueReq) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *CreateTokenWithValueReq) GetProjects() []string {
	if m != nil {
		return m.Projects
	}
	return nil
}

type UpdateTokenReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" toml:"id,omitempty" mapstructure:"id,omitempty"`
	Active               bool     `protobuf:"varint,2,opt,name=active,proto3" json:"active,omitempty" toml:"active,omitempty" mapstructure:"active,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty" toml:"description,omitempty" mapstructure:"description,omitempty"`
	Projects             []string `protobuf:"bytes,4,rep,name=projects,proto3" json:"projects,omitempty" toml:"projects,omitempty" mapstructure:"projects,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *UpdateTokenReq) Reset()         { *m = UpdateTokenReq{} }
func (m *UpdateTokenReq) String() string { return proto.CompactTextString(m) }
func (*UpdateTokenReq) ProtoMessage()    {}
func (*UpdateTokenReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_65ff3d3f7473d71c, []int{2}
}

func (m *UpdateTokenReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateTokenReq.Unmarshal(m, b)
}
func (m *UpdateTokenReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateTokenReq.Marshal(b, m, deterministic)
}
func (m *UpdateTokenReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateTokenReq.Merge(m, src)
}
func (m *UpdateTokenReq) XXX_Size() int {
	return xxx_messageInfo_UpdateTokenReq.Size(m)
}
func (m *UpdateTokenReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateTokenReq.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateTokenReq proto.InternalMessageInfo

func (m *UpdateTokenReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdateTokenReq) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

func (m *UpdateTokenReq) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *UpdateTokenReq) GetProjects() []string {
	if m != nil {
		return m.Projects
	}
	return nil
}

type Token struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" toml:"id,omitempty" mapstructure:"id,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty" toml:"description,omitempty" mapstructure:"description,omitempty"`
	Value                string   `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty" toml:"value,omitempty" mapstructure:"value,omitempty"`
	Active               bool     `protobuf:"varint,4,opt,name=active,proto3" json:"active,omitempty" toml:"active,omitempty" mapstructure:"active,omitempty"`
	Created              string   `protobuf:"bytes,5,opt,name=created,proto3" json:"created,omitempty" toml:"created,omitempty" mapstructure:"created,omitempty"`
	Updated              string   `protobuf:"bytes,6,opt,name=updated,proto3" json:"updated,omitempty" toml:"updated,omitempty" mapstructure:"updated,omitempty"`
	Projects             []string `protobuf:"bytes,7,rep,name=projects,proto3" json:"projects,omitempty" toml:"projects,omitempty" mapstructure:"projects,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *Token) Reset()         { *m = Token{} }
func (m *Token) String() string { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()    {}
func (*Token) Descriptor() ([]byte, []int) {
	return fileDescriptor_65ff3d3f7473d71c, []int{3}
}

func (m *Token) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Token.Unmarshal(m, b)
}
func (m *Token) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Token.Marshal(b, m, deterministic)
}
func (m *Token) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Token.Merge(m, src)
}
func (m *Token) XXX_Size() int {
	return xxx_messageInfo_Token.Size(m)
}
func (m *Token) XXX_DiscardUnknown() {
	xxx_messageInfo_Token.DiscardUnknown(m)
}

var xxx_messageInfo_Token proto.InternalMessageInfo

func (m *Token) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Token) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Token) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *Token) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

func (m *Token) GetCreated() string {
	if m != nil {
		return m.Created
	}
	return ""
}

func (m *Token) GetUpdated() string {
	if m != nil {
		return m.Updated
	}
	return ""
}

func (m *Token) GetProjects() []string {
	if m != nil {
		return m.Projects
	}
	return nil
}

type Tokens struct {
	Tokens               []*Token `protobuf:"bytes,1,rep,name=tokens,proto3" json:"tokens,omitempty" toml:"tokens,omitempty" mapstructure:"tokens,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *Tokens) Reset()         { *m = Tokens{} }
func (m *Tokens) String() string { return proto.CompactTextString(m) }
func (*Tokens) ProtoMessage()    {}
func (*Tokens) Descriptor() ([]byte, []int) {
	return fileDescriptor_65ff3d3f7473d71c, []int{4}
}

func (m *Tokens) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tokens.Unmarshal(m, b)
}
func (m *Tokens) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tokens.Marshal(b, m, deterministic)
}
func (m *Tokens) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tokens.Merge(m, src)
}
func (m *Tokens) XXX_Size() int {
	return xxx_messageInfo_Tokens.Size(m)
}
func (m *Tokens) XXX_DiscardUnknown() {
	xxx_messageInfo_Tokens.DiscardUnknown(m)
}

var xxx_messageInfo_Tokens proto.InternalMessageInfo

func (m *Tokens) GetTokens() []*Token {
	if m != nil {
		return m.Tokens
	}
	return nil
}

type Value struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty" toml:"value,omitempty" mapstructure:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *Value) Reset()         { *m = Value{} }
func (m *Value) String() string { return proto.CompactTextString(m) }
func (*Value) ProtoMessage()    {}
func (*Value) Descriptor() ([]byte, []int) {
	return fileDescriptor_65ff3d3f7473d71c, []int{5}
}

func (m *Value) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Value.Unmarshal(m, b)
}
func (m *Value) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Value.Marshal(b, m, deterministic)
}
func (m *Value) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Value.Merge(m, src)
}
func (m *Value) XXX_Size() int {
	return xxx_messageInfo_Value.Size(m)
}
func (m *Value) XXX_DiscardUnknown() {
	xxx_messageInfo_Value.DiscardUnknown(m)
}

var xxx_messageInfo_Value proto.InternalMessageInfo

func (m *Value) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type GetTokenReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" toml:"id,omitempty" mapstructure:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *GetTokenReq) Reset()         { *m = GetTokenReq{} }
func (m *GetTokenReq) String() string { return proto.CompactTextString(m) }
func (*GetTokenReq) ProtoMessage()    {}
func (*GetTokenReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_65ff3d3f7473d71c, []int{6}
}

func (m *GetTokenReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTokenReq.Unmarshal(m, b)
}
func (m *GetTokenReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTokenReq.Marshal(b, m, deterministic)
}
func (m *GetTokenReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTokenReq.Merge(m, src)
}
func (m *GetTokenReq) XXX_Size() int {
	return xxx_messageInfo_GetTokenReq.Size(m)
}
func (m *GetTokenReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTokenReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetTokenReq proto.InternalMessageInfo

func (m *GetTokenReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetTokensReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *GetTokensReq) Reset()         { *m = GetTokensReq{} }
func (m *GetTokensReq) String() string { return proto.CompactTextString(m) }
func (*GetTokensReq) ProtoMessage()    {}
func (*GetTokensReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_65ff3d3f7473d71c, []int{7}
}

func (m *GetTokensReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTokensReq.Unmarshal(m, b)
}
func (m *GetTokensReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTokensReq.Marshal(b, m, deterministic)
}
func (m *GetTokensReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTokensReq.Merge(m, src)
}
func (m *GetTokensReq) XXX_Size() int {
	return xxx_messageInfo_GetTokensReq.Size(m)
}
func (m *GetTokensReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTokensReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetTokensReq proto.InternalMessageInfo

type DeleteTokenReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" toml:"id,omitempty" mapstructure:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *DeleteTokenReq) Reset()         { *m = DeleteTokenReq{} }
func (m *DeleteTokenReq) String() string { return proto.CompactTextString(m) }
func (*DeleteTokenReq) ProtoMessage()    {}
func (*DeleteTokenReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_65ff3d3f7473d71c, []int{8}
}

func (m *DeleteTokenReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteTokenReq.Unmarshal(m, b)
}
func (m *DeleteTokenReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteTokenReq.Marshal(b, m, deterministic)
}
func (m *DeleteTokenReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteTokenReq.Merge(m, src)
}
func (m *DeleteTokenReq) XXX_Size() int {
	return xxx_messageInfo_DeleteTokenReq.Size(m)
}
func (m *DeleteTokenReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteTokenReq.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteTokenReq proto.InternalMessageInfo

func (m *DeleteTokenReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DeleteTokenResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *DeleteTokenResp) Reset()         { *m = DeleteTokenResp{} }
func (m *DeleteTokenResp) String() string { return proto.CompactTextString(m) }
func (*DeleteTokenResp) ProtoMessage()    {}
func (*DeleteTokenResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_65ff3d3f7473d71c, []int{9}
}

func (m *DeleteTokenResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteTokenResp.Unmarshal(m, b)
}
func (m *DeleteTokenResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteTokenResp.Marshal(b, m, deterministic)
}
func (m *DeleteTokenResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteTokenResp.Merge(m, src)
}
func (m *DeleteTokenResp) XXX_Size() int {
	return xxx_messageInfo_DeleteTokenResp.Size(m)
}
func (m *DeleteTokenResp) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteTokenResp.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteTokenResp proto.InternalMessageInfo

type ResetToV1Req struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ResetToV1Req) Reset()         { *m = ResetToV1Req{} }
func (m *ResetToV1Req) String() string { return proto.CompactTextString(m) }
func (*ResetToV1Req) ProtoMessage()    {}
func (*ResetToV1Req) Descriptor() ([]byte, []int) {
	return fileDescriptor_65ff3d3f7473d71c, []int{10}
}

func (m *ResetToV1Req) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResetToV1Req.Unmarshal(m, b)
}
func (m *ResetToV1Req) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResetToV1Req.Marshal(b, m, deterministic)
}
func (m *ResetToV1Req) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResetToV1Req.Merge(m, src)
}
func (m *ResetToV1Req) XXX_Size() int {
	return xxx_messageInfo_ResetToV1Req.Size(m)
}
func (m *ResetToV1Req) XXX_DiscardUnknown() {
	xxx_messageInfo_ResetToV1Req.DiscardUnknown(m)
}

var xxx_messageInfo_ResetToV1Req proto.InternalMessageInfo

type ResetToV1Resp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ResetToV1Resp) Reset()         { *m = ResetToV1Resp{} }
func (m *ResetToV1Resp) String() string { return proto.CompactTextString(m) }
func (*ResetToV1Resp) ProtoMessage()    {}
func (*ResetToV1Resp) Descriptor() ([]byte, []int) {
	return fileDescriptor_65ff3d3f7473d71c, []int{11}
}

func (m *ResetToV1Resp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResetToV1Resp.Unmarshal(m, b)
}
func (m *ResetToV1Resp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResetToV1Resp.Marshal(b, m, deterministic)
}
func (m *ResetToV1Resp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResetToV1Resp.Merge(m, src)
}
func (m *ResetToV1Resp) XXX_Size() int {
	return xxx_messageInfo_ResetToV1Resp.Size(m)
}
func (m *ResetToV1Resp) XXX_DiscardUnknown() {
	xxx_messageInfo_ResetToV1Resp.DiscardUnknown(m)
}

var xxx_messageInfo_ResetToV1Resp proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CreateTokenReq)(nil), "chef.automate.domain.authn.CreateTokenReq")
	proto.RegisterType((*CreateTokenWithValueReq)(nil), "chef.automate.domain.authn.CreateTokenWithValueReq")
	proto.RegisterType((*UpdateTokenReq)(nil), "chef.automate.domain.authn.UpdateTokenReq")
	proto.RegisterType((*Token)(nil), "chef.automate.domain.authn.Token")
	proto.RegisterType((*Tokens)(nil), "chef.automate.domain.authn.Tokens")
	proto.RegisterType((*Value)(nil), "chef.automate.domain.authn.Value")
	proto.RegisterType((*GetTokenReq)(nil), "chef.automate.domain.authn.GetTokenReq")
	proto.RegisterType((*GetTokensReq)(nil), "chef.automate.domain.authn.GetTokensReq")
	proto.RegisterType((*DeleteTokenReq)(nil), "chef.automate.domain.authn.DeleteTokenReq")
	proto.RegisterType((*DeleteTokenResp)(nil), "chef.automate.domain.authn.DeleteTokenResp")
	proto.RegisterType((*ResetToV1Req)(nil), "chef.automate.domain.authn.ResetToV1Req")
	proto.RegisterType((*ResetToV1Resp)(nil), "chef.automate.domain.authn.ResetToV1Resp")
}

func init() {
	proto.RegisterFile("api/interservice/authn/tokens.proto", fileDescriptor_65ff3d3f7473d71c)
}

var fileDescriptor_65ff3d3f7473d71c = []byte{
	// 616 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0xcd, 0x6e, 0xd3, 0x4e,
	0x10, 0xaf, 0xe3, 0x3a, 0x6d, 0xc6, 0xff, 0x7f, 0x2a, 0x56, 0xa5, 0x35, 0x16, 0x45, 0x66, 0xa9,
	0x44, 0xf8, 0x88, 0x4d, 0x52, 0x84, 0xd4, 0x0b, 0x87, 0x14, 0x89, 0x13, 0x17, 0x0b, 0x0a, 0xa2,
	0x22, 0x68, 0x63, 0x2f, 0xc9, 0x42, 0x62, 0xbb, 0xf6, 0x26, 0x07, 0x3e, 0x5e, 0x82, 0x23, 0x2f,
	0x82, 0xe8, 0x09, 0x89, 0xa7, 0xe1, 0x2d, 0xd0, 0xda, 0x71, 0xb2, 0x6e, 0x9b, 0xd4, 0x6a, 0x6f,
	0x9d, 0xee, 0xcc, 0xfc, 0x3e, 0x66, 0x26, 0x86, 0x3b, 0x24, 0x62, 0x0e, 0x0b, 0x38, 0x8d, 0x13,
	0x1a, 0x4f, 0x98, 0x47, 0x1d, 0x32, 0xe6, 0x83, 0xc0, 0xe1, 0xe1, 0x27, 0x1a, 0x24, 0x76, 0x14,
	0x87, 0x3c, 0x44, 0xa6, 0x37, 0xa0, 0x1f, 0x6c, 0x32, 0xe6, 0xe1, 0x88, 0x70, 0x6a, 0xfb, 0xe1,
	0x88, 0xb0, 0xc0, 0x4e, 0x13, 0xcd, 0x9b, 0xfd, 0x30, 0xec, 0x0f, 0xa9, 0x23, 0xfa, 0x90, 0x20,
	0x08, 0x39, 0xe1, 0x2c, 0xcc, 0x2b, 0xcd, 0xed, 0x09, 0x19, 0x32, 0x9f, 0x70, 0xea, 0xe4, 0x7f,
	0x64, 0x0f, 0xf8, 0x44, 0x81, 0xfa, 0x41, 0x4c, 0x09, 0xa7, 0x2f, 0x05, 0x92, 0x4b, 0x8f, 0x91,
	0x0d, 0x15, 0xe6, 0x1b, 0x8a, 0xa5, 0x34, 0x6a, 0x9d, 0x5b, 0x27, 0x7f, 0x7f, 0xab, 0x37, 0xe2,
	0xed, 0xf6, 0xf5, 0xee, 0x11, 0x69, 0x7e, 0x7e, 0xd4, 0xdc, 0x6f, 0xbe, 0x7f, 0xf7, 0xa5, 0xf5,
	0xf0, 0xc9, 0xe3, 0x6f, 0xbb, 0x5f, 0xbb, 0xbb, 0x6e, 0x85, 0xf9, 0xc8, 0x02, 0xdd, 0xa7, 0x89,
	0x17, 0xb3, 0x48, 0x20, 0x1a, 0x15, 0x51, 0xe8, 0xca, 0xff, 0x42, 0x5b, 0x50, 0x25, 0x1e, 0x67,
	0x13, 0x6a, 0xa8, 0x96, 0xd2, 0x58, 0x77, 0xa7, 0x11, 0x7a, 0x0a, 0xeb, 0x51, 0x1c, 0x7e, 0xa4,
	0x1e, 0x4f, 0x8c, 0x55, 0x4b, 0x6d, 0xd4, 0x3a, 0x58, 0xe0, 0xed, 0x7c, 0x57, 0x4c, 0x43, 0xc1,
	0x5b, 0xf1, 0x66, 0x1b, 0x9d, 0x85, 0x75, 0x67, 0x35, 0xf8, 0x97, 0x02, 0xdb, 0x12, 0xf9, 0xd7,
	0x8c, 0x0f, 0x0e, 0xc9, 0x70, 0x4c, 0x85, 0x8a, 0xfa, 0x5c, 0xc5, 0x15, 0x59, 0x6e, 0x82, 0x36,
	0x11, 0x5d, 0x8d, 0xd5, 0xb4, 0x26, 0x0b, 0x0a, 0xdc, 0xb5, 0x4b, 0x70, 0xff, 0xa1, 0x40, 0xfd,
	0x55, 0xe4, 0xcb, 0xc6, 0x9f, 0xa6, 0x3c, 0x27, 0x54, 0x29, 0x10, 0x3a, 0x25, 0x45, 0x3d, 0x2b,
	0xe5, 0xaa, 0xc6, 0xfe, 0x54, 0x40, 0x4b, 0x69, 0x5d, 0xc2, 0xc6, 0x99, 0x5d, 0xaa, 0x6c, 0xd7,
	0x5c, 0xcb, 0x6a, 0x41, 0x8b, 0x01, 0x6b, 0x5e, 0x3a, 0x41, 0xdf, 0xd0, 0xd2, 0xfc, 0x3c, 0x14,
	0x2f, 0xe3, 0xd4, 0x1f, 0xdf, 0xa8, 0x66, 0x2f, 0xd3, 0x10, 0x99, 0x92, 0xba, 0x35, 0xa1, 0x4e,
	0x62, 0x7e, 0x00, 0xd5, 0x94, 0x78, 0x82, 0xf6, 0xa1, 0x9a, 0x1d, 0x8f, 0xa1, 0x58, 0x6a, 0x43,
	0x6f, 0xdf, 0xb6, 0x17, 0x5f, 0x8f, 0x9d, 0xcd, 0x60, 0x5a, 0x80, 0x77, 0x40, 0x4b, 0xf7, 0x68,
	0xae, 0x45, 0x91, 0xb4, 0xe0, 0x1d, 0xd0, 0x9f, 0x53, 0xbe, 0x68, 0x6c, 0xb8, 0x0e, 0xff, 0xe5,
	0xcf, 0x89, 0x4b, 0x8f, 0xb1, 0x05, 0xf5, 0x67, 0x74, 0x48, 0x17, 0x0f, 0x1a, 0x5f, 0x83, 0x8d,
	0x42, 0x46, 0x12, 0x89, 0x26, 0x2e, 0x4d, 0x44, 0x9b, 0xc3, 0x96, 0x68, 0xb2, 0x01, 0xff, 0x4b,
	0x71, 0x12, 0xb5, 0xff, 0x68, 0x00, 0x19, 0xc6, 0x8b, 0xfe, 0x88, 0xa3, 0x23, 0xa8, 0xcd, 0x40,
	0x51, 0x63, 0x99, 0x54, 0x99, 0x9b, 0x89, 0x2f, 0x34, 0x25, 0xc1, 0x2b, 0xa8, 0x0b, 0xba, 0x74,
	0x66, 0xe8, 0xfe, 0xb2, 0xa2, 0xe2, 0x8f, 0x89, 0x79, 0xb1, 0xeb, 0x78, 0x05, 0x45, 0xb0, 0x79,
	0xde, 0x19, 0xa3, 0xbd, 0x92, 0x40, 0xf2, 0xe1, 0x97, 0x43, 0xec, 0x82, 0x2e, 0x1d, 0xdf, 0x72,
	0x45, 0xc5, 0x2b, 0x2d, 0xd7, 0xff, 0x0d, 0xac, 0xe7, 0x3e, 0xa3, 0xbb, 0x65, 0xa6, 0x51, 0xba,
	0xf3, 0x00, 0x74, 0x69, 0x57, 0x96, 0x33, 0x2f, 0xae, 0x9d, 0xf9, 0xa0, 0x74, 0x6e, 0x12, 0xe1,
	0x15, 0xd4, 0x83, 0xda, 0x6c, 0xe5, 0x96, 0xaf, 0x94, 0xbc, 0xa9, 0xe6, 0xbd, 0x92, 0x99, 0x02,
	0xa3, 0xd3, 0x7a, 0xeb, 0xf4, 0x19, 0x1f, 0x8c, 0x7b, 0xb6, 0x17, 0x8e, 0x1c, 0x51, 0xe8, 0xe4,
	0x85, 0xce, 0xf9, 0x5f, 0xc4, 0x5e, 0x35, 0xfd, 0x70, 0xed, 0xfd, 0x0b, 0x00, 0x00, 0xff, 0xff,
	0x01, 0xf7, 0x44, 0x4d, 0x32, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TokensMgmtClient is the client API for TokensMgmt service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TokensMgmtClient interface {
	GetTokens(ctx context.Context, in *GetTokensReq, opts ...grpc.CallOption) (*Tokens, error)
	CreateToken(ctx context.Context, in *CreateTokenReq, opts ...grpc.CallOption) (*Token, error)
	CreateTokenWithValue(ctx context.Context, in *CreateTokenWithValueReq, opts ...grpc.CallOption) (*Token, error)
	UpdateToken(ctx context.Context, in *UpdateTokenReq, opts ...grpc.CallOption) (*Token, error)
	GetToken(ctx context.Context, in *GetTokenReq, opts ...grpc.CallOption) (*Token, error)
	DeleteToken(ctx context.Context, in *DeleteTokenReq, opts ...grpc.CallOption) (*DeleteTokenResp, error)
	ResetToV1(ctx context.Context, in *ResetToV1Req, opts ...grpc.CallOption) (*ResetToV1Resp, error)
}

type tokensMgmtClient struct {
	cc *grpc.ClientConn
}

func NewTokensMgmtClient(cc *grpc.ClientConn) TokensMgmtClient {
	return &tokensMgmtClient{cc}
}

func (c *tokensMgmtClient) GetTokens(ctx context.Context, in *GetTokensReq, opts ...grpc.CallOption) (*Tokens, error) {
	out := new(Tokens)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.authn.TokensMgmt/GetTokens", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokensMgmtClient) CreateToken(ctx context.Context, in *CreateTokenReq, opts ...grpc.CallOption) (*Token, error) {
	out := new(Token)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.authn.TokensMgmt/CreateToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokensMgmtClient) CreateTokenWithValue(ctx context.Context, in *CreateTokenWithValueReq, opts ...grpc.CallOption) (*Token, error) {
	out := new(Token)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.authn.TokensMgmt/CreateTokenWithValue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokensMgmtClient) UpdateToken(ctx context.Context, in *UpdateTokenReq, opts ...grpc.CallOption) (*Token, error) {
	out := new(Token)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.authn.TokensMgmt/UpdateToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokensMgmtClient) GetToken(ctx context.Context, in *GetTokenReq, opts ...grpc.CallOption) (*Token, error) {
	out := new(Token)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.authn.TokensMgmt/GetToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokensMgmtClient) DeleteToken(ctx context.Context, in *DeleteTokenReq, opts ...grpc.CallOption) (*DeleteTokenResp, error) {
	out := new(DeleteTokenResp)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.authn.TokensMgmt/DeleteToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokensMgmtClient) ResetToV1(ctx context.Context, in *ResetToV1Req, opts ...grpc.CallOption) (*ResetToV1Resp, error) {
	out := new(ResetToV1Resp)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.authn.TokensMgmt/ResetToV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TokensMgmtServer is the server API for TokensMgmt service.
type TokensMgmtServer interface {
	GetTokens(context.Context, *GetTokensReq) (*Tokens, error)
	CreateToken(context.Context, *CreateTokenReq) (*Token, error)
	CreateTokenWithValue(context.Context, *CreateTokenWithValueReq) (*Token, error)
	UpdateToken(context.Context, *UpdateTokenReq) (*Token, error)
	GetToken(context.Context, *GetTokenReq) (*Token, error)
	DeleteToken(context.Context, *DeleteTokenReq) (*DeleteTokenResp, error)
	ResetToV1(context.Context, *ResetToV1Req) (*ResetToV1Resp, error)
}

// UnimplementedTokensMgmtServer can be embedded to have forward compatible implementations.
type UnimplementedTokensMgmtServer struct {
}

func (*UnimplementedTokensMgmtServer) GetTokens(ctx context.Context, req *GetTokensReq) (*Tokens, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTokens not implemented")
}
func (*UnimplementedTokensMgmtServer) CreateToken(ctx context.Context, req *CreateTokenReq) (*Token, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateToken not implemented")
}
func (*UnimplementedTokensMgmtServer) CreateTokenWithValue(ctx context.Context, req *CreateTokenWithValueReq) (*Token, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTokenWithValue not implemented")
}
func (*UnimplementedTokensMgmtServer) UpdateToken(ctx context.Context, req *UpdateTokenReq) (*Token, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateToken not implemented")
}
func (*UnimplementedTokensMgmtServer) GetToken(ctx context.Context, req *GetTokenReq) (*Token, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetToken not implemented")
}
func (*UnimplementedTokensMgmtServer) DeleteToken(ctx context.Context, req *DeleteTokenReq) (*DeleteTokenResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteToken not implemented")
}
func (*UnimplementedTokensMgmtServer) ResetToV1(ctx context.Context, req *ResetToV1Req) (*ResetToV1Resp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetToV1 not implemented")
}

func RegisterTokensMgmtServer(s *grpc.Server, srv TokensMgmtServer) {
	s.RegisterService(&_TokensMgmt_serviceDesc, srv)
}

func _TokensMgmt_GetTokens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTokensReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokensMgmtServer).GetTokens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.authn.TokensMgmt/GetTokens",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokensMgmtServer).GetTokens(ctx, req.(*GetTokensReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokensMgmt_CreateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokensMgmtServer).CreateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.authn.TokensMgmt/CreateToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokensMgmtServer).CreateToken(ctx, req.(*CreateTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokensMgmt_CreateTokenWithValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTokenWithValueReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokensMgmtServer).CreateTokenWithValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.authn.TokensMgmt/CreateTokenWithValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokensMgmtServer).CreateTokenWithValue(ctx, req.(*CreateTokenWithValueReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokensMgmt_UpdateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokensMgmtServer).UpdateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.authn.TokensMgmt/UpdateToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokensMgmtServer).UpdateToken(ctx, req.(*UpdateTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokensMgmt_GetToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokensMgmtServer).GetToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.authn.TokensMgmt/GetToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokensMgmtServer).GetToken(ctx, req.(*GetTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokensMgmt_DeleteToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokensMgmtServer).DeleteToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.authn.TokensMgmt/DeleteToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokensMgmtServer).DeleteToken(ctx, req.(*DeleteTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokensMgmt_ResetToV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetToV1Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokensMgmtServer).ResetToV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.authn.TokensMgmt/ResetToV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokensMgmtServer).ResetToV1(ctx, req.(*ResetToV1Req))
	}
	return interceptor(ctx, in, info, handler)
}

var _TokensMgmt_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chef.automate.domain.authn.TokensMgmt",
	HandlerType: (*TokensMgmtServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTokens",
			Handler:    _TokensMgmt_GetTokens_Handler,
		},
		{
			MethodName: "CreateToken",
			Handler:    _TokensMgmt_CreateToken_Handler,
		},
		{
			MethodName: "CreateTokenWithValue",
			Handler:    _TokensMgmt_CreateTokenWithValue_Handler,
		},
		{
			MethodName: "UpdateToken",
			Handler:    _TokensMgmt_UpdateToken_Handler,
		},
		{
			MethodName: "GetToken",
			Handler:    _TokensMgmt_GetToken_Handler,
		},
		{
			MethodName: "DeleteToken",
			Handler:    _TokensMgmt_DeleteToken_Handler,
		},
		{
			MethodName: "ResetToV1",
			Handler:    _TokensMgmt_ResetToV1_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/interservice/authn/tokens.proto",
}
