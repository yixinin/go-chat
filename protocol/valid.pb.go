// Code generated by protoc-gen-go. DO NOT EDIT.
// source: valid.proto

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

type GetValidReq struct {
	Header               *ReqHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	ValidType            int32      `protobuf:"varint,2,opt,name=ValidType,proto3" json:"ValidType,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *GetValidReq) Reset()         { *m = GetValidReq{} }
func (m *GetValidReq) String() string { return proto.CompactTextString(m) }
func (*GetValidReq) ProtoMessage()    {}
func (*GetValidReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_6bda6afcf0898d4a, []int{0}
}

func (m *GetValidReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetValidReq.Unmarshal(m, b)
}
func (m *GetValidReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetValidReq.Marshal(b, m, deterministic)
}
func (m *GetValidReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetValidReq.Merge(m, src)
}
func (m *GetValidReq) XXX_Size() int {
	return xxx_messageInfo_GetValidReq.Size(m)
}
func (m *GetValidReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetValidReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetValidReq proto.InternalMessageInfo

func (m *GetValidReq) GetHeader() *ReqHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *GetValidReq) GetValidType() int32 {
	if m != nil {
		return m.ValidType
	}
	return 0
}

type GetValidAck struct {
	Header               *AckHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	ValidId              string     `protobuf:"bytes,2,opt,name=ValidId,proto3" json:"ValidId,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *GetValidAck) Reset()         { *m = GetValidAck{} }
func (m *GetValidAck) String() string { return proto.CompactTextString(m) }
func (*GetValidAck) ProtoMessage()    {}
func (*GetValidAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_6bda6afcf0898d4a, []int{1}
}

func (m *GetValidAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetValidAck.Unmarshal(m, b)
}
func (m *GetValidAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetValidAck.Marshal(b, m, deterministic)
}
func (m *GetValidAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetValidAck.Merge(m, src)
}
func (m *GetValidAck) XXX_Size() int {
	return xxx_messageInfo_GetValidAck.Size(m)
}
func (m *GetValidAck) XXX_DiscardUnknown() {
	xxx_messageInfo_GetValidAck.DiscardUnknown(m)
}

var xxx_messageInfo_GetValidAck proto.InternalMessageInfo

func (m *GetValidAck) GetHeader() *AckHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *GetValidAck) GetValidId() string {
	if m != nil {
		return m.ValidId
	}
	return ""
}

//检查用户名、deviceCode、phone是否可用
type CheckUserIdReq struct {
	Header               *ReqHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Value                string     `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	CheckType            int32      `protobuf:"varint,3,opt,name=checkType,proto3" json:"checkType,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *CheckUserIdReq) Reset()         { *m = CheckUserIdReq{} }
func (m *CheckUserIdReq) String() string { return proto.CompactTextString(m) }
func (*CheckUserIdReq) ProtoMessage()    {}
func (*CheckUserIdReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_6bda6afcf0898d4a, []int{2}
}

func (m *CheckUserIdReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckUserIdReq.Unmarshal(m, b)
}
func (m *CheckUserIdReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckUserIdReq.Marshal(b, m, deterministic)
}
func (m *CheckUserIdReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckUserIdReq.Merge(m, src)
}
func (m *CheckUserIdReq) XXX_Size() int {
	return xxx_messageInfo_CheckUserIdReq.Size(m)
}
func (m *CheckUserIdReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckUserIdReq.DiscardUnknown(m)
}

var xxx_messageInfo_CheckUserIdReq proto.InternalMessageInfo

func (m *CheckUserIdReq) GetHeader() *ReqHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *CheckUserIdReq) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *CheckUserIdReq) GetCheckType() int32 {
	if m != nil {
		return m.CheckType
	}
	return 0
}

type CheckUserIdAck struct {
	Header               *AckHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *CheckUserIdAck) Reset()         { *m = CheckUserIdAck{} }
func (m *CheckUserIdAck) String() string { return proto.CompactTextString(m) }
func (*CheckUserIdAck) ProtoMessage()    {}
func (*CheckUserIdAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_6bda6afcf0898d4a, []int{3}
}

func (m *CheckUserIdAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckUserIdAck.Unmarshal(m, b)
}
func (m *CheckUserIdAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckUserIdAck.Marshal(b, m, deterministic)
}
func (m *CheckUserIdAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckUserIdAck.Merge(m, src)
}
func (m *CheckUserIdAck) XXX_Size() int {
	return xxx_messageInfo_CheckUserIdAck.Size(m)
}
func (m *CheckUserIdAck) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckUserIdAck.DiscardUnknown(m)
}

var xxx_messageInfo_CheckUserIdAck proto.InternalMessageInfo

func (m *CheckUserIdAck) GetHeader() *AckHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func init() {
	proto.RegisterType((*GetValidReq)(nil), "protocol.GetValidReq")
	proto.RegisterType((*GetValidAck)(nil), "protocol.GetValidAck")
	proto.RegisterType((*CheckUserIdReq)(nil), "protocol.CheckUserIdReq")
	proto.RegisterType((*CheckUserIdAck)(nil), "protocol.CheckUserIdAck")
}

func init() { proto.RegisterFile("valid.proto", fileDescriptor_6bda6afcf0898d4a) }

var fileDescriptor_6bda6afcf0898d4a = []byte{
	// 193 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0x4b, 0xcc, 0xc9,
	0x4c, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00, 0x53, 0xc9, 0xf9, 0x39, 0x52, 0x3c,
	0x19, 0xa9, 0x89, 0x29, 0xa9, 0x45, 0x10, 0x71, 0xa5, 0x08, 0x2e, 0x6e, 0xf7, 0xd4, 0x92, 0x30,
	0x90, 0xca, 0xa0, 0xd4, 0x42, 0x21, 0x6d, 0x2e, 0x36, 0x88, 0xb4, 0x04, 0xa3, 0x02, 0xa3, 0x06,
	0xb7, 0x91, 0xb0, 0x1e, 0x4c, 0x9f, 0x5e, 0x50, 0x6a, 0xa1, 0x07, 0x58, 0x2a, 0x08, 0xaa, 0x44,
	0x48, 0x86, 0x8b, 0x13, 0xac, 0x31, 0xa4, 0xb2, 0x20, 0x55, 0x82, 0x49, 0x81, 0x51, 0x83, 0x35,
	0x08, 0x21, 0xa0, 0x14, 0x82, 0x30, 0xd9, 0x31, 0x39, 0x1b, 0x9f, 0xc9, 0x8e, 0xc9, 0xd9, 0x68,
	0x26, 0x4b, 0x70, 0xb1, 0x83, 0x35, 0x7a, 0xa6, 0x80, 0xcd, 0xe5, 0x0c, 0x82, 0x71, 0x95, 0x0a,
	0xb9, 0xf8, 0x9c, 0x33, 0x52, 0x93, 0xb3, 0x43, 0x8b, 0x53, 0x8b, 0x3c, 0x49, 0x77, 0xb2, 0x08,
	0x17, 0x6b, 0x59, 0x62, 0x4e, 0x69, 0x2a, 0xd4, 0x58, 0x08, 0x07, 0xe4, 0x91, 0x64, 0x90, 0xa1,
	0x60, 0x8f, 0x30, 0x43, 0x3c, 0x02, 0x17, 0x50, 0xb2, 0x45, 0xb1, 0x92, 0x54, 0xbf, 0x24, 0xb1,
	0x81, 0xe5, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x9d, 0x8f, 0x59, 0x64, 0x8f, 0x01, 0x00,
	0x00,
}
