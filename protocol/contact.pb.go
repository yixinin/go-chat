// Code generated by protoc-gen-go. DO NOT EDIT.
// source: contact.proto

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

type SearchUserReq struct {
	Header               *ReqHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Key                  string     `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *SearchUserReq) Reset()         { *m = SearchUserReq{} }
func (m *SearchUserReq) String() string { return proto.CompactTextString(m) }
func (*SearchUserReq) ProtoMessage()    {}
func (*SearchUserReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5036fff2565fb15, []int{0}
}

func (m *SearchUserReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchUserReq.Unmarshal(m, b)
}
func (m *SearchUserReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchUserReq.Marshal(b, m, deterministic)
}
func (m *SearchUserReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchUserReq.Merge(m, src)
}
func (m *SearchUserReq) XXX_Size() int {
	return xxx_messageInfo_SearchUserReq.Size(m)
}
func (m *SearchUserReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchUserReq.DiscardUnknown(m)
}

var xxx_messageInfo_SearchUserReq proto.InternalMessageInfo

func (m *SearchUserReq) GetHeader() *ReqHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SearchUserReq) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type SearchUserAck struct {
	Header               *AckHeader                `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	PageSize             string                    `protobuf:"bytes,2,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	Data                 []*SearchUserAck_DataItem `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *SearchUserAck) Reset()         { *m = SearchUserAck{} }
func (m *SearchUserAck) String() string { return proto.CompactTextString(m) }
func (*SearchUserAck) ProtoMessage()    {}
func (*SearchUserAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5036fff2565fb15, []int{1}
}

func (m *SearchUserAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchUserAck.Unmarshal(m, b)
}
func (m *SearchUserAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchUserAck.Marshal(b, m, deterministic)
}
func (m *SearchUserAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchUserAck.Merge(m, src)
}
func (m *SearchUserAck) XXX_Size() int {
	return xxx_messageInfo_SearchUserAck.Size(m)
}
func (m *SearchUserAck) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchUserAck.DiscardUnknown(m)
}

var xxx_messageInfo_SearchUserAck proto.InternalMessageInfo

func (m *SearchUserAck) GetHeader() *AckHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SearchUserAck) GetPageSize() string {
	if m != nil {
		return m.PageSize
	}
	return ""
}

func (m *SearchUserAck) GetData() []*SearchUserAck_DataItem {
	if m != nil {
		return m.Data
	}
	return nil
}

type SearchUserAck_DataItem struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Nickname             string   `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Avatar               string   `protobuf:"bytes,3,opt,name=avatar,proto3" json:"avatar,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchUserAck_DataItem) Reset()         { *m = SearchUserAck_DataItem{} }
func (m *SearchUserAck_DataItem) String() string { return proto.CompactTextString(m) }
func (*SearchUserAck_DataItem) ProtoMessage()    {}
func (*SearchUserAck_DataItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5036fff2565fb15, []int{1, 0}
}

func (m *SearchUserAck_DataItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchUserAck_DataItem.Unmarshal(m, b)
}
func (m *SearchUserAck_DataItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchUserAck_DataItem.Marshal(b, m, deterministic)
}
func (m *SearchUserAck_DataItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchUserAck_DataItem.Merge(m, src)
}
func (m *SearchUserAck_DataItem) XXX_Size() int {
	return xxx_messageInfo_SearchUserAck_DataItem.Size(m)
}
func (m *SearchUserAck_DataItem) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchUserAck_DataItem.DiscardUnknown(m)
}

var xxx_messageInfo_SearchUserAck_DataItem proto.InternalMessageInfo

func (m *SearchUserAck_DataItem) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *SearchUserAck_DataItem) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *SearchUserAck_DataItem) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

type AddContactReq struct {
	Header               *ReqHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Username             string     `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	ContactType          int32      `protobuf:"varint,3,opt,name=contactType,proto3" json:"contactType,omitempty"`
	SetRemarks           string     `protobuf:"bytes,4,opt,name=setRemarks,proto3" json:"setRemarks,omitempty"`
	Msg                  string     `protobuf:"bytes,5,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *AddContactReq) Reset()         { *m = AddContactReq{} }
func (m *AddContactReq) String() string { return proto.CompactTextString(m) }
func (*AddContactReq) ProtoMessage()    {}
func (*AddContactReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5036fff2565fb15, []int{2}
}

func (m *AddContactReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddContactReq.Unmarshal(m, b)
}
func (m *AddContactReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddContactReq.Marshal(b, m, deterministic)
}
func (m *AddContactReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddContactReq.Merge(m, src)
}
func (m *AddContactReq) XXX_Size() int {
	return xxx_messageInfo_AddContactReq.Size(m)
}
func (m *AddContactReq) XXX_DiscardUnknown() {
	xxx_messageInfo_AddContactReq.DiscardUnknown(m)
}

var xxx_messageInfo_AddContactReq proto.InternalMessageInfo

func (m *AddContactReq) GetHeader() *ReqHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *AddContactReq) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *AddContactReq) GetContactType() int32 {
	if m != nil {
		return m.ContactType
	}
	return 0
}

func (m *AddContactReq) GetSetRemarks() string {
	if m != nil {
		return m.SetRemarks
	}
	return ""
}

func (m *AddContactReq) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type AddContactAck struct {
	Header               *AckHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *AddContactAck) Reset()         { *m = AddContactAck{} }
func (m *AddContactAck) String() string { return proto.CompactTextString(m) }
func (*AddContactAck) ProtoMessage()    {}
func (*AddContactAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5036fff2565fb15, []int{3}
}

func (m *AddContactAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddContactAck.Unmarshal(m, b)
}
func (m *AddContactAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddContactAck.Marshal(b, m, deterministic)
}
func (m *AddContactAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddContactAck.Merge(m, src)
}
func (m *AddContactAck) XXX_Size() int {
	return xxx_messageInfo_AddContactAck.Size(m)
}
func (m *AddContactAck) XXX_DiscardUnknown() {
	xxx_messageInfo_AddContactAck.DiscardUnknown(m)
}

var xxx_messageInfo_AddContactAck proto.InternalMessageInfo

func (m *AddContactAck) GetHeader() *AckHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type AddContactNotify struct {
	Header               *ReqHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	ContactId            string     `protobuf:"bytes,2,opt,name=contactId,proto3" json:"contactId,omitempty"`
	Nickname             string     `protobuf:"bytes,3,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Avatar               string     `protobuf:"bytes,4,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Msg                  string     `protobuf:"bytes,5,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *AddContactNotify) Reset()         { *m = AddContactNotify{} }
func (m *AddContactNotify) String() string { return proto.CompactTextString(m) }
func (*AddContactNotify) ProtoMessage()    {}
func (*AddContactNotify) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5036fff2565fb15, []int{4}
}

func (m *AddContactNotify) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddContactNotify.Unmarshal(m, b)
}
func (m *AddContactNotify) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddContactNotify.Marshal(b, m, deterministic)
}
func (m *AddContactNotify) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddContactNotify.Merge(m, src)
}
func (m *AddContactNotify) XXX_Size() int {
	return xxx_messageInfo_AddContactNotify.Size(m)
}
func (m *AddContactNotify) XXX_DiscardUnknown() {
	xxx_messageInfo_AddContactNotify.DiscardUnknown(m)
}

var xxx_messageInfo_AddContactNotify proto.InternalMessageInfo

func (m *AddContactNotify) GetHeader() *ReqHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *AddContactNotify) GetContactId() string {
	if m != nil {
		return m.ContactId
	}
	return ""
}

func (m *AddContactNotify) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *AddContactNotify) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

func (m *AddContactNotify) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type DeleteContactReq struct {
	ContactId            string   `protobuf:"bytes,1,opt,name=contactId,proto3" json:"contactId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteContactReq) Reset()         { *m = DeleteContactReq{} }
func (m *DeleteContactReq) String() string { return proto.CompactTextString(m) }
func (*DeleteContactReq) ProtoMessage()    {}
func (*DeleteContactReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5036fff2565fb15, []int{5}
}

func (m *DeleteContactReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteContactReq.Unmarshal(m, b)
}
func (m *DeleteContactReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteContactReq.Marshal(b, m, deterministic)
}
func (m *DeleteContactReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteContactReq.Merge(m, src)
}
func (m *DeleteContactReq) XXX_Size() int {
	return xxx_messageInfo_DeleteContactReq.Size(m)
}
func (m *DeleteContactReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteContactReq.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteContactReq proto.InternalMessageInfo

func (m *DeleteContactReq) GetContactId() string {
	if m != nil {
		return m.ContactId
	}
	return ""
}

type DeleteContactAck struct {
	Header               *AckHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *DeleteContactAck) Reset()         { *m = DeleteContactAck{} }
func (m *DeleteContactAck) String() string { return proto.CompactTextString(m) }
func (*DeleteContactAck) ProtoMessage()    {}
func (*DeleteContactAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5036fff2565fb15, []int{6}
}

func (m *DeleteContactAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteContactAck.Unmarshal(m, b)
}
func (m *DeleteContactAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteContactAck.Marshal(b, m, deterministic)
}
func (m *DeleteContactAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteContactAck.Merge(m, src)
}
func (m *DeleteContactAck) XXX_Size() int {
	return xxx_messageInfo_DeleteContactAck.Size(m)
}
func (m *DeleteContactAck) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteContactAck.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteContactAck proto.InternalMessageInfo

func (m *DeleteContactAck) GetHeader() *AckHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type UpdateContactReq struct {
	Header               *ReqHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	SetRemarks           string     `protobuf:"bytes,2,opt,name=setRemarks,proto3" json:"setRemarks,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *UpdateContactReq) Reset()         { *m = UpdateContactReq{} }
func (m *UpdateContactReq) String() string { return proto.CompactTextString(m) }
func (*UpdateContactReq) ProtoMessage()    {}
func (*UpdateContactReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5036fff2565fb15, []int{7}
}

func (m *UpdateContactReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateContactReq.Unmarshal(m, b)
}
func (m *UpdateContactReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateContactReq.Marshal(b, m, deterministic)
}
func (m *UpdateContactReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateContactReq.Merge(m, src)
}
func (m *UpdateContactReq) XXX_Size() int {
	return xxx_messageInfo_UpdateContactReq.Size(m)
}
func (m *UpdateContactReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateContactReq.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateContactReq proto.InternalMessageInfo

func (m *UpdateContactReq) GetHeader() *ReqHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *UpdateContactReq) GetSetRemarks() string {
	if m != nil {
		return m.SetRemarks
	}
	return ""
}

type UpdateContactAck struct {
	Header               *AckHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *UpdateContactAck) Reset()         { *m = UpdateContactAck{} }
func (m *UpdateContactAck) String() string { return proto.CompactTextString(m) }
func (*UpdateContactAck) ProtoMessage()    {}
func (*UpdateContactAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5036fff2565fb15, []int{8}
}

func (m *UpdateContactAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateContactAck.Unmarshal(m, b)
}
func (m *UpdateContactAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateContactAck.Marshal(b, m, deterministic)
}
func (m *UpdateContactAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateContactAck.Merge(m, src)
}
func (m *UpdateContactAck) XXX_Size() int {
	return xxx_messageInfo_UpdateContactAck.Size(m)
}
func (m *UpdateContactAck) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateContactAck.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateContactAck proto.InternalMessageInfo

func (m *UpdateContactAck) GetHeader() *AckHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func init() {
	proto.RegisterType((*SearchUserReq)(nil), "protocol.SearchUserReq")
	proto.RegisterType((*SearchUserAck)(nil), "protocol.SearchUserAck")
	proto.RegisterType((*SearchUserAck_DataItem)(nil), "protocol.SearchUserAck.DataItem")
	proto.RegisterType((*AddContactReq)(nil), "protocol.AddContactReq")
	proto.RegisterType((*AddContactAck)(nil), "protocol.AddContactAck")
	proto.RegisterType((*AddContactNotify)(nil), "protocol.AddContactNotify")
	proto.RegisterType((*DeleteContactReq)(nil), "protocol.DeleteContactReq")
	proto.RegisterType((*DeleteContactAck)(nil), "protocol.DeleteContactAck")
	proto.RegisterType((*UpdateContactReq)(nil), "protocol.UpdateContactReq")
	proto.RegisterType((*UpdateContactAck)(nil), "protocol.UpdateContactAck")
}

func init() { proto.RegisterFile("contact.proto", fileDescriptor_a5036fff2565fb15) }

var fileDescriptor_a5036fff2565fb15 = []byte{
	// 381 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x51, 0x41, 0x6b, 0xf2, 0x40,
	0x10, 0x25, 0x46, 0x45, 0xc7, 0x4f, 0x08, 0xf9, 0xe0, 0x23, 0xc8, 0x47, 0x09, 0x39, 0x09, 0x85,
	0x50, 0x6c, 0x8f, 0x85, 0x22, 0xf5, 0x50, 0x2f, 0x1e, 0xd6, 0x7a, 0xe9, 0xa5, 0x4c, 0x93, 0xa9,
	0x4a, 0x8c, 0x89, 0x9b, 0x6d, 0xc1, 0xfe, 0x9d, 0x5e, 0xfb, 0xcb, 0xfa, 0x2b, 0x4a, 0x36, 0xab,
	0x6e, 0x42, 0x5b, 0xc8, 0x29, 0x99, 0x79, 0x3b, 0xef, 0xbd, 0x99, 0x07, 0xfd, 0x20, 0xd9, 0x0a,
	0x0c, 0x84, 0x9f, 0xf2, 0x44, 0x24, 0x76, 0x47, 0x7e, 0x82, 0x64, 0x33, 0xf8, 0xb3, 0x22, 0x0c,
	0x89, 0x17, 0x7d, 0x6f, 0x06, 0xfd, 0x39, 0x21, 0x0f, 0x56, 0x8b, 0x8c, 0x38, 0xa3, 0x9d, 0x7d,
	0x0e, 0xed, 0xe2, 0x81, 0x63, 0xb8, 0xc6, 0xb0, 0x37, 0xfa, 0xeb, 0x1f, 0x26, 0x7d, 0x46, 0xbb,
	0x3b, 0x09, 0x31, 0xf5, 0xc4, 0xb6, 0xc0, 0x8c, 0x68, 0xef, 0x34, 0x5c, 0x63, 0xd8, 0x65, 0xf9,
	0xaf, 0xf7, 0x69, 0xe8, 0x84, 0xe3, 0x20, 0xfa, 0x8d, 0x70, 0x1c, 0x44, 0x15, 0xc2, 0x01, 0x74,
	0x52, 0x5c, 0xd2, 0x7c, 0xfd, 0x46, 0x8a, 0xf5, 0x58, 0xdb, 0x57, 0xd0, 0x0c, 0x51, 0xa0, 0x63,
	0xba, 0xe6, 0xb0, 0x37, 0x72, 0x4f, 0x34, 0x25, 0x3d, 0x7f, 0x82, 0x02, 0xa7, 0x82, 0x62, 0x26,
	0x5f, 0x0f, 0x1e, 0xa0, 0x73, 0xe8, 0xe4, 0xec, 0x2f, 0x19, 0xf1, 0x2d, 0xc6, 0x24, 0xcd, 0x74,
	0xd9, 0xb1, 0xce, 0xb1, 0xed, 0x3a, 0x88, 0x24, 0xa6, 0x94, 0x0f, 0xb5, 0xfd, 0x0f, 0xda, 0xf8,
	0x8a, 0x02, 0xb9, 0x63, 0x4a, 0x44, 0x55, 0xde, 0x87, 0x01, 0xfd, 0x71, 0x18, 0xde, 0x16, 0x97,
	0xae, 0x7d, 0x3d, 0xdd, 0x4e, 0xa3, 0x62, 0xc7, 0x85, 0x9e, 0x0a, 0xf0, 0x7e, 0x9f, 0x92, 0xd4,
	0x6d, 0x31, 0xbd, 0x65, 0x9f, 0x01, 0x64, 0x24, 0x18, 0xc5, 0xc8, 0xa3, 0xcc, 0x69, 0xca, 0x79,
	0xad, 0x93, 0x67, 0x13, 0x67, 0x4b, 0xa7, 0x55, 0x64, 0x13, 0x67, 0x4b, 0xef, 0x5a, 0x77, 0x5b,
	0x37, 0x1a, 0xef, 0xdd, 0x00, 0xeb, 0x34, 0x3e, 0x4b, 0xc4, 0xfa, 0x79, 0x5f, 0x6f, 0xdf, 0xff,
	0xd0, 0x55, 0x0b, 0x4c, 0x43, 0xb5, 0xf0, 0xa9, 0x51, 0x0a, 0xc0, 0xfc, 0x31, 0x80, 0xa6, 0x1e,
	0xc0, 0x37, 0x3b, 0x5e, 0x80, 0x35, 0xa1, 0x0d, 0x09, 0xd2, 0x42, 0x29, 0xe9, 0x1a, 0x15, 0x5d,
	0xef, 0xa6, 0x32, 0x51, 0xfb, 0x30, 0x8f, 0x60, 0x2d, 0xd2, 0x10, 0x4b, 0x92, 0xb5, 0xee, 0x52,
	0x4e, 0xb2, 0x51, 0x4d, 0x32, 0x77, 0x58, 0x12, 0xa8, 0xeb, 0xf0, 0xa9, 0x2d, 0xb1, 0xcb, 0xaf,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x74, 0x9b, 0x92, 0xab, 0x14, 0x04, 0x00, 0x00,
}