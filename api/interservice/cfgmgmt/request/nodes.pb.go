// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/interservice/cfgmgmt/request/nodes.proto

package request

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type NodeRun struct {
	NodeId               string               `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty" toml:"node_id,omitempty" mapstructure:"node_id,omitempty"`
	RunId                string               `protobuf:"bytes,2,opt,name=run_id,json=runId,proto3" json:"run_id,omitempty" toml:"run_id,omitempty" mapstructure:"run_id,omitempty"`
	EndTime              *timestamp.Timestamp `protobuf:"bytes,3,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty" toml:"end_time,omitempty" mapstructure:"end_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte               `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *NodeRun) Reset()         { *m = NodeRun{} }
func (m *NodeRun) String() string { return proto.CompactTextString(m) }
func (*NodeRun) ProtoMessage()    {}
func (*NodeRun) Descriptor() ([]byte, []int) {
	return fileDescriptor_2790a303f0e1261f, []int{0}
}

func (m *NodeRun) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeRun.Unmarshal(m, b)
}
func (m *NodeRun) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeRun.Marshal(b, m, deterministic)
}
func (m *NodeRun) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeRun.Merge(m, src)
}
func (m *NodeRun) XXX_Size() int {
	return xxx_messageInfo_NodeRun.Size(m)
}
func (m *NodeRun) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeRun.DiscardUnknown(m)
}

var xxx_messageInfo_NodeRun proto.InternalMessageInfo

func (m *NodeRun) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *NodeRun) GetRunId() string {
	if m != nil {
		return m.RunId
	}
	return ""
}

func (m *NodeRun) GetEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

type Nodes struct {
	Filter               []string    `protobuf:"bytes,1,rep,name=filter,proto3" json:"filter,omitempty" toml:"filter,omitempty" mapstructure:"filter,omitempty"`
	Pagination           *Pagination `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty" toml:"pagination,omitempty" mapstructure:"pagination,omitempty"`
	Sorting              *Sorting    `protobuf:"bytes,3,opt,name=sorting,proto3" json:"sorting,omitempty" toml:"sorting,omitempty" mapstructure:"sorting,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte      `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32       `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *Nodes) Reset()         { *m = Nodes{} }
func (m *Nodes) String() string { return proto.CompactTextString(m) }
func (*Nodes) ProtoMessage()    {}
func (*Nodes) Descriptor() ([]byte, []int) {
	return fileDescriptor_2790a303f0e1261f, []int{1}
}

func (m *Nodes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Nodes.Unmarshal(m, b)
}
func (m *Nodes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Nodes.Marshal(b, m, deterministic)
}
func (m *Nodes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Nodes.Merge(m, src)
}
func (m *Nodes) XXX_Size() int {
	return xxx_messageInfo_Nodes.Size(m)
}
func (m *Nodes) XXX_DiscardUnknown() {
	xxx_messageInfo_Nodes.DiscardUnknown(m)
}

var xxx_messageInfo_Nodes proto.InternalMessageInfo

func (m *Nodes) GetFilter() []string {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (m *Nodes) GetPagination() *Pagination {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func (m *Nodes) GetSorting() *Sorting {
	if m != nil {
		return m.Sorting
	}
	return nil
}

type Node struct {
	NodeId               string   `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty" toml:"node_id,omitempty" mapstructure:"node_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_2790a303f0e1261f, []int{2}
}

func (m *Node) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Node.Unmarshal(m, b)
}
func (m *Node) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Node.Marshal(b, m, deterministic)
}
func (m *Node) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Node.Merge(m, src)
}
func (m *Node) XXX_Size() int {
	return xxx_messageInfo_Node.Size(m)
}
func (m *Node) XXX_DiscardUnknown() {
	xxx_messageInfo_Node.DiscardUnknown(m)
}

var xxx_messageInfo_Node proto.InternalMessageInfo

func (m *Node) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

type Runs struct {
	NodeId     string      `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty" toml:"node_id,omitempty" mapstructure:"node_id,omitempty"`
	Filter     []string    `protobuf:"bytes,2,rep,name=filter,proto3" json:"filter,omitempty" toml:"filter,omitempty" mapstructure:"filter,omitempty"`
	Pagination *Pagination `protobuf:"bytes,3,opt,name=pagination,proto3" json:"pagination,omitempty" toml:"pagination,omitempty" mapstructure:"pagination,omitempty"`
	// TODO: (@afiune) Should we standardize these parameters as well?
	Start                string   `protobuf:"bytes,4,opt,name=start,proto3" json:"start,omitempty" toml:"start,omitempty" mapstructure:"start,omitempty"`
	End                  string   `protobuf:"bytes,5,opt,name=end,proto3" json:"end,omitempty" toml:"end,omitempty" mapstructure:"end,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *Runs) Reset()         { *m = Runs{} }
func (m *Runs) String() string { return proto.CompactTextString(m) }
func (*Runs) ProtoMessage()    {}
func (*Runs) Descriptor() ([]byte, []int) {
	return fileDescriptor_2790a303f0e1261f, []int{3}
}

func (m *Runs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Runs.Unmarshal(m, b)
}
func (m *Runs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Runs.Marshal(b, m, deterministic)
}
func (m *Runs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Runs.Merge(m, src)
}
func (m *Runs) XXX_Size() int {
	return xxx_messageInfo_Runs.Size(m)
}
func (m *Runs) XXX_DiscardUnknown() {
	xxx_messageInfo_Runs.DiscardUnknown(m)
}

var xxx_messageInfo_Runs proto.InternalMessageInfo

func (m *Runs) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *Runs) GetFilter() []string {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (m *Runs) GetPagination() *Pagination {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func (m *Runs) GetStart() string {
	if m != nil {
		return m.Start
	}
	return ""
}

func (m *Runs) GetEnd() string {
	if m != nil {
		return m.End
	}
	return ""
}

func init() {
	proto.RegisterType((*NodeRun)(nil), "chef.automate.domain.cfgmgmt.request.NodeRun")
	proto.RegisterType((*Nodes)(nil), "chef.automate.domain.cfgmgmt.request.Nodes")
	proto.RegisterType((*Node)(nil), "chef.automate.domain.cfgmgmt.request.Node")
	proto.RegisterType((*Runs)(nil), "chef.automate.domain.cfgmgmt.request.Runs")
}

func init() {
	proto.RegisterFile("api/interservice/cfgmgmt/request/nodes.proto", fileDescriptor_2790a303f0e1261f)
}

var fileDescriptor_2790a303f0e1261f = []byte{
	// 363 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0x3f, 0x6f, 0xdb, 0x30,
	0x10, 0xc5, 0x21, 0xcb, 0x92, 0x5a, 0x7a, 0x29, 0x88, 0xfe, 0x11, 0xbc, 0xd8, 0x10, 0x3a, 0x78,
	0x68, 0xc9, 0xd6, 0x45, 0x87, 0xa2, 0x5b, 0x96, 0xc0, 0x4b, 0x60, 0x28, 0x99, 0xb2, 0x18, 0xb4,
	0x78, 0x92, 0x09, 0x98, 0xa4, 0xc2, 0x3f, 0xf9, 0x5e, 0x59, 0xf3, 0xe9, 0x02, 0x4a, 0x56, 0x92,
	0xc5, 0xb1, 0x87, 0x6c, 0x7a, 0x07, 0xbd, 0x7b, 0xbf, 0xbb, 0x23, 0xfa, 0xc1, 0x5a, 0x41, 0x85,
	0x72, 0x60, 0x2c, 0x98, 0x7b, 0x51, 0x01, 0xad, 0xea, 0x46, 0x36, 0xd2, 0x51, 0x03, 0x77, 0x1e,
	0xac, 0xa3, 0x4a, 0x73, 0xb0, 0xa4, 0x35, 0xda, 0x69, 0xfc, 0xbd, 0xda, 0x41, 0x4d, 0x98, 0x77,
	0x5a, 0x32, 0x07, 0x84, 0x6b, 0xc9, 0x84, 0x22, 0x07, 0x07, 0x39, 0x38, 0xa6, 0xb3, 0x46, 0xeb,
	0x66, 0x0f, 0xb4, 0xf3, 0x6c, 0x7d, 0x4d, 0x9d, 0x90, 0x60, 0x1d, 0x93, 0x6d, 0xdf, 0x66, 0xfa,
	0xfb, 0x64, 0x68, 0xcb, 0x0c, 0x93, 0x10, 0x7e, 0xe8, 0x2d, 0x45, 0x8b, 0xb2, 0x2b, 0xcd, 0xa1,
	0xf4, 0x0a, 0x7f, 0x43, 0x59, 0x60, 0xda, 0x08, 0x9e, 0x47, 0xf3, 0x68, 0xf1, 0xb1, 0x4c, 0x83,
	0x5c, 0x71, 0xfc, 0x05, 0xa5, 0xc6, 0xab, 0x50, 0x1f, 0x75, 0xf5, 0xc4, 0x78, 0xb5, 0xe2, 0xf8,
	0x2f, 0xfa, 0x00, 0x8a, 0x6f, 0x02, 0x44, 0x1e, 0xcf, 0xa3, 0xc5, 0x64, 0x39, 0x25, 0x3d, 0x21,
	0x19, 0x08, 0xc9, 0xcd, 0x40, 0x58, 0x66, 0xa0, 0x78, 0x50, 0xc5, 0x63, 0x84, 0x92, 0x10, 0x69,
	0xf1, 0x57, 0x94, 0xd6, 0x62, 0xef, 0xc0, 0xe4, 0xd1, 0x3c, 0x0e, 0x79, 0xbd, 0xc2, 0x6b, 0x84,
	0x5a, 0xd6, 0x08, 0xc5, 0x9c, 0xd0, 0xaa, 0xcb, 0x9c, 0x2c, 0x7f, 0x91, 0x73, 0x56, 0x44, 0xd6,
	0xcf, 0xbe, 0xf2, 0x55, 0x0f, 0x7c, 0x89, 0x32, 0xab, 0x8d, 0x13, 0xaa, 0x39, 0x90, 0xfe, 0x3c,
	0xaf, 0xdd, 0x75, 0x6f, 0x2a, 0x07, 0x77, 0x31, 0x43, 0xe3, 0xc0, 0x7e, 0x74, 0x57, 0xc5, 0x43,
	0x84, 0xc6, 0xa5, 0x57, 0xf6, 0xf8, 0x36, 0x5f, 0xa6, 0x1e, 0xbd, 0x31, 0x75, 0xfc, 0x0e, 0x53,
	0x7f, 0x46, 0x89, 0x75, 0xcc, 0xb8, 0x7c, 0xdc, 0x9f, 0xad, 0x13, 0xf8, 0x13, 0x8a, 0x41, 0xf1,
	0x3c, 0xe9, 0x6a, 0xe1, 0xf3, 0xe2, 0xff, 0xed, 0xbf, 0x46, 0xb8, 0x9d, 0xdf, 0x92, 0x4a, 0x4b,
	0x1a, 0x12, 0xe9, 0x90, 0x48, 0x4f, 0xbd, 0xa8, 0x6d, 0xda, 0xdd, 0xfa, 0xcf, 0x53, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xd5, 0xd4, 0xed, 0x0e, 0xf1, 0x02, 0x00, 0x00,
}
