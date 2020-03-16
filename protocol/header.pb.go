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
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
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

func (m *ReqHeader) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
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

type NotiHeader struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NotiHeader) Reset()         { *m = NotiHeader{} }
func (m *NotiHeader) String() string { return proto.CompactTextString(m) }
func (*NotiHeader) ProtoMessage()    {}
func (*NotiHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_6398613e36d6c2ce, []int{2}
}

func (m *NotiHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotiHeader.Unmarshal(m, b)
}
func (m *NotiHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotiHeader.Marshal(b, m, deterministic)
}
func (m *NotiHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotiHeader.Merge(m, src)
}
func (m *NotiHeader) XXX_Size() int {
	return xxx_messageInfo_NotiHeader.Size(m)
}
func (m *NotiHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_NotiHeader.DiscardUnknown(m)
}

var xxx_messageInfo_NotiHeader proto.InternalMessageInfo

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

//轮询请求
type PollReq struct {
	Header               *ReqHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	NotifyId             int32      `protobuf:"varint,2,opt,name=notifyId,proto3" json:"notifyId,omitempty"`
	Body                 string     `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *PollReq) Reset()         { *m = PollReq{} }
func (m *PollReq) String() string { return proto.CompactTextString(m) }
func (*PollReq) ProtoMessage()    {}
func (*PollReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_6398613e36d6c2ce, []int{4}
}

func (m *PollReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PollReq.Unmarshal(m, b)
}
func (m *PollReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PollReq.Marshal(b, m, deterministic)
}
func (m *PollReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PollReq.Merge(m, src)
}
func (m *PollReq) XXX_Size() int {
	return xxx_messageInfo_PollReq.Size(m)
}
func (m *PollReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PollReq.DiscardUnknown(m)
}

var xxx_messageInfo_PollReq proto.InternalMessageInfo

func (m *PollReq) GetHeader() *ReqHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *PollReq) GetNotifyId() int32 {
	if m != nil {
		return m.NotifyId
	}
	return 0
}

func (m *PollReq) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func init() {
	proto.RegisterType((*ReqHeader)(nil), "protocol.ReqHeader")
	proto.RegisterType((*AckHeader)(nil), "protocol.AckHeader")
	proto.RegisterType((*NotiHeader)(nil), "protocol.NotiHeader")
	proto.RegisterType((*CallAckHeader)(nil), "protocol.CallAckHeader")
	proto.RegisterType((*PollReq)(nil), "protocol.PollReq")
}

func init() { proto.RegisterFile("header.proto", fileDescriptor_6398613e36d6c2ce) }

var fileDescriptor_6398613e36d6c2ce = []byte{
	// 202 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xc9, 0x48, 0x4d, 0x4c,
	0x49, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00, 0x53, 0xc9, 0xf9, 0x39, 0x4a,
	0xc6, 0x5c, 0x9c, 0x41, 0xa9, 0x85, 0x1e, 0x60, 0x49, 0x21, 0x01, 0x2e, 0xe6, 0xd2, 0xcc, 0x14,
	0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x10, 0x53, 0x48, 0x84, 0x8b, 0xb5, 0x24, 0x3f, 0x3b,
	0x35, 0x4f, 0x82, 0x09, 0x2c, 0x06, 0xe1, 0x28, 0x19, 0x72, 0x71, 0x3a, 0x26, 0x67, 0x43, 0x35,
	0x09, 0x71, 0xb1, 0x24, 0xe7, 0xa7, 0xa4, 0x82, 0x75, 0xb1, 0x06, 0x81, 0xd9, 0x20, 0x83, 0x72,
	0x8b, 0xd3, 0xa1, 0x9a, 0x40, 0x4c, 0x25, 0x1e, 0x2e, 0x2e, 0xbf, 0xfc, 0x92, 0x4c, 0x88, 0x1e,
	0x25, 0x53, 0x2e, 0x5e, 0xe7, 0xc4, 0x9c, 0x1c, 0x52, 0x0d, 0x49, 0xe3, 0x62, 0x0f, 0xc8, 0xcf,
	0xc9, 0x09, 0x4a, 0x2d, 0x14, 0xd2, 0xe6, 0x62, 0x83, 0xf8, 0x08, 0xac, 0x85, 0xdb, 0x48, 0x58,
	0x0f, 0xe6, 0x25, 0x3d, 0xb8, 0x7f, 0x82, 0xa0, 0x4a, 0x84, 0xa4, 0xb8, 0x38, 0xf2, 0xf2, 0x4b,
	0x32, 0xd3, 0x2a, 0x3d, 0x53, 0xc0, 0xc6, 0xb1, 0x06, 0xc1, 0xf9, 0x20, 0x9b, 0x93, 0xf2, 0x53,
	0x2a, 0x25, 0x98, 0xc1, 0xd6, 0x80, 0xd9, 0x49, 0x6c, 0x60, 0xb3, 0x8c, 0x01, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x7d, 0xc6, 0x6d, 0xe4, 0x35, 0x01, 0x00, 0x00,
}
