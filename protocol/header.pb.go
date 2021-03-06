// Code generated by protoc-gen-go. DO NOT EDIT.
// source: header.proto

package protocol

import (
	fmt "fmt"
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

type ReqHeader struct {
	Uid                  int64    `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqHeader) Reset()         { *m = ReqHeader{} }
func (m *ReqHeader) String() string { return proto.CompactTextString(m) }
func (*ReqHeader) ProtoMessage()    {}
func (*ReqHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_6398613e36d6c2ce, []int{0}
}

func (m *ReqHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqHeader.Unmarshal(m, b)
}
func (m *ReqHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqHeader.Marshal(b, m, deterministic)
}
func (m *ReqHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqHeader.Merge(m, src)
}
func (m *ReqHeader) XXX_Size() int {
	return xxx_messageInfo_ReqHeader.Size(m)
}
func (m *ReqHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqHeader.DiscardUnknown(m)
}

var xxx_messageInfo_ReqHeader proto.InternalMessageInfo

func (m *ReqHeader) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *ReqHeader) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type AckHeader struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Uid                  int64    `protobuf:"varint,3,opt,name=uid,proto3" json:"uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AckHeader) Reset()         { *m = AckHeader{} }
func (m *AckHeader) String() string { return proto.CompactTextString(m) }
func (*AckHeader) ProtoMessage()    {}
func (*AckHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_6398613e36d6c2ce, []int{1}
}

func (m *AckHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AckHeader.Unmarshal(m, b)
}
func (m *AckHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AckHeader.Marshal(b, m, deterministic)
}
func (m *AckHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AckHeader.Merge(m, src)
}
func (m *AckHeader) XXX_Size() int {
	return xxx_messageInfo_AckHeader.Size(m)
}
func (m *AckHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_AckHeader.DiscardUnknown(m)
}

var xxx_messageInfo_AckHeader proto.InternalMessageInfo

func (m *AckHeader) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *AckHeader) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *AckHeader) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

type NotifyHeader struct {
	Uid                  int64    `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NotifyHeader) Reset()         { *m = NotifyHeader{} }
func (m *NotifyHeader) String() string { return proto.CompactTextString(m) }
func (*NotifyHeader) ProtoMessage()    {}
func (*NotifyHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_6398613e36d6c2ce, []int{2}
}

func (m *NotifyHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotifyHeader.Unmarshal(m, b)
}
func (m *NotifyHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotifyHeader.Marshal(b, m, deterministic)
}
func (m *NotifyHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotifyHeader.Merge(m, src)
}
func (m *NotifyHeader) XXX_Size() int {
	return xxx_messageInfo_NotifyHeader.Size(m)
}
func (m *NotifyHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_NotifyHeader.DiscardUnknown(m)
}

var xxx_messageInfo_NotifyHeader proto.InternalMessageInfo

func (m *NotifyHeader) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

type CallAckHeader struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CallAckHeader) Reset()         { *m = CallAckHeader{} }
func (m *CallAckHeader) String() string { return proto.CompactTextString(m) }
func (*CallAckHeader) ProtoMessage()    {}
func (*CallAckHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_6398613e36d6c2ce, []int{3}
}

func (m *CallAckHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CallAckHeader.Unmarshal(m, b)
}
func (m *CallAckHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CallAckHeader.Marshal(b, m, deterministic)
}
func (m *CallAckHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CallAckHeader.Merge(m, src)
}
func (m *CallAckHeader) XXX_Size() int {
	return xxx_messageInfo_CallAckHeader.Size(m)
}
func (m *CallAckHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_CallAckHeader.DiscardUnknown(m)
}

var xxx_messageInfo_CallAckHeader proto.InternalMessageInfo

func (m *CallAckHeader) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *CallAckHeader) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*ReqHeader)(nil), "protocol.ReqHeader")
	proto.RegisterType((*AckHeader)(nil), "protocol.AckHeader")
	proto.RegisterType((*NotifyHeader)(nil), "protocol.NotifyHeader")
	proto.RegisterType((*CallAckHeader)(nil), "protocol.CallAckHeader")
}

func init() { proto.RegisterFile("header.proto", fileDescriptor_6398613e36d6c2ce) }

var fileDescriptor_6398613e36d6c2ce = []byte{
	// 152 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xc9, 0x48, 0x4d, 0x4c,
	0x49, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00, 0x53, 0xc9, 0xf9, 0x39, 0x4a,
	0xc6, 0x5c, 0x9c, 0x41, 0xa9, 0x85, 0x1e, 0x60, 0x49, 0x21, 0x01, 0x2e, 0xe6, 0xd2, 0xcc, 0x14,
	0x09, 0x46, 0x05, 0x46, 0x0d, 0xe6, 0x20, 0x10, 0x53, 0x48, 0x84, 0x8b, 0xb5, 0x24, 0x3f, 0x3b,
	0x35, 0x4f, 0x82, 0x49, 0x81, 0x51, 0x83, 0x33, 0x08, 0xc2, 0x51, 0x72, 0xe6, 0xe2, 0x74, 0x4c,
	0xce, 0x86, 0x6a, 0x12, 0xe2, 0x62, 0x49, 0xce, 0x4f, 0x49, 0x05, 0xeb, 0x62, 0x0d, 0x02, 0xb3,
	0x41, 0x06, 0xe5, 0x16, 0xa7, 0x43, 0x35, 0x81, 0x98, 0x30, 0xa3, 0x99, 0xe1, 0x46, 0x2b, 0x29,
	0x70, 0xf1, 0xf8, 0xe5, 0x97, 0x64, 0xa6, 0x55, 0xe2, 0xb2, 0x5c, 0xc9, 0x94, 0x8b, 0xd7, 0x39,
	0x31, 0x27, 0x87, 0x44, 0xab, 0x92, 0xd8, 0xc0, 0x9e, 0x33, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff,
	0x73, 0x9e, 0x0f, 0x9b, 0xf3, 0x00, 0x00, 0x00,
}
