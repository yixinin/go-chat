// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chat.proto

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

type MessageType int32

const (
	MessageType_None    MessageType = 0
	MessageType_Text    MessageType = 1
	MessageType_Pictrue MessageType = 2
	MessageType_Audio   MessageType = 3
	MessageType_Video   MessageType = 4
	MessageType_Link    MessageType = 5
)

var MessageType_name = map[int32]string{
	0: "None",
	1: "Text",
	2: "Pictrue",
	3: "Audio",
	4: "Video",
	5: "Link",
}

var MessageType_value = map[string]int32{
	"None":    0,
	"Text":    1,
	"Pictrue": 2,
	"Audio":   3,
	"Video":   4,
	"Link":    5,
}

func (x MessageType) String() string {
	return proto.EnumName(MessageType_name, int32(x))
}

func (MessageType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{0}
}

type SendMessageReq struct {
	Header               *ReqHeader  `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	MessageType          MessageType `protobuf:"varint,2,opt,name=messageType,proto3,enum=protocol.MessageType" json:"messageType,omitempty"`
	TextMessage          string      `protobuf:"bytes,3,opt,name=textMessage,proto3" json:"textMessage,omitempty"`
	Memtions             []string    `protobuf:"bytes,4,rep,name=memtions,proto3" json:"memtions,omitempty"`
	ContactId            string      `protobuf:"bytes,5,opt,name=contactId,proto3" json:"contactId,omitempty"`
	Username             string      `protobuf:"bytes,6,opt,name=username,proto3" json:"username,omitempty"`
	GroupId              string      `protobuf:"bytes,7,opt,name=groupId,proto3" json:"groupId,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *SendMessageReq) Reset()         { *m = SendMessageReq{} }
func (m *SendMessageReq) String() string { return proto.CompactTextString(m) }
func (*SendMessageReq) ProtoMessage()    {}
func (*SendMessageReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{0}
}

func (m *SendMessageReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendMessageReq.Unmarshal(m, b)
}
func (m *SendMessageReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendMessageReq.Marshal(b, m, deterministic)
}
func (m *SendMessageReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendMessageReq.Merge(m, src)
}
func (m *SendMessageReq) XXX_Size() int {
	return xxx_messageInfo_SendMessageReq.Size(m)
}
func (m *SendMessageReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SendMessageReq.DiscardUnknown(m)
}

var xxx_messageInfo_SendMessageReq proto.InternalMessageInfo

func (m *SendMessageReq) GetHeader() *ReqHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SendMessageReq) GetMessageType() MessageType {
	if m != nil {
		return m.MessageType
	}
	return MessageType_None
}

func (m *SendMessageReq) GetTextMessage() string {
	if m != nil {
		return m.TextMessage
	}
	return ""
}

func (m *SendMessageReq) GetMemtions() []string {
	if m != nil {
		return m.Memtions
	}
	return nil
}

func (m *SendMessageReq) GetContactId() string {
	if m != nil {
		return m.ContactId
	}
	return ""
}

func (m *SendMessageReq) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *SendMessageReq) GetGroupId() string {
	if m != nil {
		return m.GroupId
	}
	return ""
}

type SendMessageAck struct {
	Header               *AckHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *SendMessageAck) Reset()         { *m = SendMessageAck{} }
func (m *SendMessageAck) String() string { return proto.CompactTextString(m) }
func (*SendMessageAck) ProtoMessage()    {}
func (*SendMessageAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{1}
}

func (m *SendMessageAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendMessageAck.Unmarshal(m, b)
}
func (m *SendMessageAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendMessageAck.Marshal(b, m, deterministic)
}
func (m *SendMessageAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendMessageAck.Merge(m, src)
}
func (m *SendMessageAck) XXX_Size() int {
	return xxx_messageInfo_SendMessageAck.Size(m)
}
func (m *SendMessageAck) XXX_DiscardUnknown() {
	xxx_messageInfo_SendMessageAck.DiscardUnknown(m)
}

var xxx_messageInfo_SendMessageAck proto.InternalMessageInfo

func (m *SendMessageAck) GetHeader() *AckHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type MessageNotify struct {
	Header               *NotifyHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	MessageType          MessageType   `protobuf:"varint,2,opt,name=messageType,proto3,enum=protocol.MessageType" json:"messageType,omitempty"`
	TextMessage          string        `protobuf:"bytes,3,opt,name=textMessage,proto3" json:"textMessage,omitempty"`
	Memtions             []string      `protobuf:"bytes,4,rep,name=memtions,proto3" json:"memtions,omitempty"`
	ContactId            string        `protobuf:"bytes,5,opt,name=contactId,proto3" json:"contactId,omitempty"`
	Username             string        `protobuf:"bytes,6,opt,name=username,proto3" json:"username,omitempty"`
	GroupId              string        `protobuf:"bytes,7,opt,name=groupId,proto3" json:"groupId,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *MessageNotify) Reset()         { *m = MessageNotify{} }
func (m *MessageNotify) String() string { return proto.CompactTextString(m) }
func (*MessageNotify) ProtoMessage()    {}
func (*MessageNotify) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{2}
}

func (m *MessageNotify) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageNotify.Unmarshal(m, b)
}
func (m *MessageNotify) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageNotify.Marshal(b, m, deterministic)
}
func (m *MessageNotify) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageNotify.Merge(m, src)
}
func (m *MessageNotify) XXX_Size() int {
	return xxx_messageInfo_MessageNotify.Size(m)
}
func (m *MessageNotify) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageNotify.DiscardUnknown(m)
}

var xxx_messageInfo_MessageNotify proto.InternalMessageInfo

func (m *MessageNotify) GetHeader() *NotifyHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *MessageNotify) GetMessageType() MessageType {
	if m != nil {
		return m.MessageType
	}
	return MessageType_None
}

func (m *MessageNotify) GetTextMessage() string {
	if m != nil {
		return m.TextMessage
	}
	return ""
}

func (m *MessageNotify) GetMemtions() []string {
	if m != nil {
		return m.Memtions
	}
	return nil
}

func (m *MessageNotify) GetContactId() string {
	if m != nil {
		return m.ContactId
	}
	return ""
}

func (m *MessageNotify) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *MessageNotify) GetGroupId() string {
	if m != nil {
		return m.GroupId
	}
	return ""
}

type RealTimeReq struct {
	Header               *ReqHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	ContactId            string     `protobuf:"bytes,2,opt,name=contactId,proto3" json:"contactId,omitempty"`
	GroupId              string     `protobuf:"bytes,3,opt,name=groupId,proto3" json:"groupId,omitempty"`
	Protocol             string     `protobuf:"bytes,4,opt,name=protocol,proto3" json:"protocol,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *RealTimeReq) Reset()         { *m = RealTimeReq{} }
func (m *RealTimeReq) String() string { return proto.CompactTextString(m) }
func (*RealTimeReq) ProtoMessage()    {}
func (*RealTimeReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{3}
}

func (m *RealTimeReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RealTimeReq.Unmarshal(m, b)
}
func (m *RealTimeReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RealTimeReq.Marshal(b, m, deterministic)
}
func (m *RealTimeReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RealTimeReq.Merge(m, src)
}
func (m *RealTimeReq) XXX_Size() int {
	return xxx_messageInfo_RealTimeReq.Size(m)
}
func (m *RealTimeReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RealTimeReq.DiscardUnknown(m)
}

var xxx_messageInfo_RealTimeReq proto.InternalMessageInfo

func (m *RealTimeReq) GetHeader() *ReqHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *RealTimeReq) GetContactId() string {
	if m != nil {
		return m.ContactId
	}
	return ""
}

func (m *RealTimeReq) GetGroupId() string {
	if m != nil {
		return m.GroupId
	}
	return ""
}

func (m *RealTimeReq) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

type RealTimeAck struct {
	Header               *AckHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	TcpAddr              string     `protobuf:"bytes,2,opt,name=tcpAddr,proto3" json:"tcpAddr,omitempty"`
	Token                string     `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	RoomId               int32      `protobuf:"varint,4,opt,name=roomId,proto3" json:"roomId,omitempty"`
	WsAddr               string     `protobuf:"bytes,5,opt,name=wsAddr,proto3" json:"wsAddr,omitempty"`
	HttpAddr             string     `protobuf:"bytes,6,opt,name=httpAddr,proto3" json:"httpAddr,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *RealTimeAck) Reset()         { *m = RealTimeAck{} }
func (m *RealTimeAck) String() string { return proto.CompactTextString(m) }
func (*RealTimeAck) ProtoMessage()    {}
func (*RealTimeAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{4}
}

func (m *RealTimeAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RealTimeAck.Unmarshal(m, b)
}
func (m *RealTimeAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RealTimeAck.Marshal(b, m, deterministic)
}
func (m *RealTimeAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RealTimeAck.Merge(m, src)
}
func (m *RealTimeAck) XXX_Size() int {
	return xxx_messageInfo_RealTimeAck.Size(m)
}
func (m *RealTimeAck) XXX_DiscardUnknown() {
	xxx_messageInfo_RealTimeAck.DiscardUnknown(m)
}

var xxx_messageInfo_RealTimeAck proto.InternalMessageInfo

func (m *RealTimeAck) GetHeader() *AckHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *RealTimeAck) GetTcpAddr() string {
	if m != nil {
		return m.TcpAddr
	}
	return ""
}

func (m *RealTimeAck) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *RealTimeAck) GetRoomId() int32 {
	if m != nil {
		return m.RoomId
	}
	return 0
}

func (m *RealTimeAck) GetWsAddr() string {
	if m != nil {
		return m.WsAddr
	}
	return ""
}

func (m *RealTimeAck) GetHttpAddr() string {
	if m != nil {
		return m.HttpAddr
	}
	return ""
}

type SimpleUser struct {
	Nickname             string   `protobuf:"bytes,1,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Avatart              string   `protobuf:"bytes,2,opt,name=avatart,proto3" json:"avatart,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SimpleUser) Reset()         { *m = SimpleUser{} }
func (m *SimpleUser) String() string { return proto.CompactTextString(m) }
func (*SimpleUser) ProtoMessage()    {}
func (*SimpleUser) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{5}
}

func (m *SimpleUser) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SimpleUser.Unmarshal(m, b)
}
func (m *SimpleUser) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SimpleUser.Marshal(b, m, deterministic)
}
func (m *SimpleUser) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SimpleUser.Merge(m, src)
}
func (m *SimpleUser) XXX_Size() int {
	return xxx_messageInfo_SimpleUser.Size(m)
}
func (m *SimpleUser) XXX_DiscardUnknown() {
	xxx_messageInfo_SimpleUser.DiscardUnknown(m)
}

var xxx_messageInfo_SimpleUser proto.InternalMessageInfo

func (m *SimpleUser) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *SimpleUser) GetAvatart() string {
	if m != nil {
		return m.Avatart
	}
	return ""
}

type RealTimeInfo struct {
	Token                string        `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	RoomId               int32         `protobuf:"varint,4,opt,name=roomId,proto3" json:"roomId,omitempty"`
	Uid                  int64         `protobuf:"varint,5,opt,name=uid,proto3" json:"uid,omitempty"`
	GroupId              string        `protobuf:"bytes,6,opt,name=groupId,proto3" json:"groupId,omitempty"`
	Users                []*SimpleUser `protobuf:"bytes,7,rep,name=users,proto3" json:"users,omitempty"`
	HttpAddr             string        `protobuf:"bytes,8,opt,name=httpAddr,proto3" json:"httpAddr,omitempty"`
	WsAddr               string        `protobuf:"bytes,1,opt,name=wsAddr,proto3" json:"wsAddr,omitempty"`
	TcpAddr              string        `protobuf:"bytes,2,opt,name=tcpAddr,proto3" json:"tcpAddr,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *RealTimeInfo) Reset()         { *m = RealTimeInfo{} }
func (m *RealTimeInfo) String() string { return proto.CompactTextString(m) }
func (*RealTimeInfo) ProtoMessage()    {}
func (*RealTimeInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{6}
}

func (m *RealTimeInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RealTimeInfo.Unmarshal(m, b)
}
func (m *RealTimeInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RealTimeInfo.Marshal(b, m, deterministic)
}
func (m *RealTimeInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RealTimeInfo.Merge(m, src)
}
func (m *RealTimeInfo) XXX_Size() int {
	return xxx_messageInfo_RealTimeInfo.Size(m)
}
func (m *RealTimeInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RealTimeInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RealTimeInfo proto.InternalMessageInfo

func (m *RealTimeInfo) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *RealTimeInfo) GetRoomId() int32 {
	if m != nil {
		return m.RoomId
	}
	return 0
}

func (m *RealTimeInfo) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *RealTimeInfo) GetGroupId() string {
	if m != nil {
		return m.GroupId
	}
	return ""
}

func (m *RealTimeInfo) GetUsers() []*SimpleUser {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *RealTimeInfo) GetHttpAddr() string {
	if m != nil {
		return m.HttpAddr
	}
	return ""
}

func (m *RealTimeInfo) GetWsAddr() string {
	if m != nil {
		return m.WsAddr
	}
	return ""
}

func (m *RealTimeInfo) GetTcpAddr() string {
	if m != nil {
		return m.TcpAddr
	}
	return ""
}

type CancelRealTimeReq struct {
	Header               *ReqHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *CancelRealTimeReq) Reset()         { *m = CancelRealTimeReq{} }
func (m *CancelRealTimeReq) String() string { return proto.CompactTextString(m) }
func (*CancelRealTimeReq) ProtoMessage()    {}
func (*CancelRealTimeReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{7}
}

func (m *CancelRealTimeReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CancelRealTimeReq.Unmarshal(m, b)
}
func (m *CancelRealTimeReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CancelRealTimeReq.Marshal(b, m, deterministic)
}
func (m *CancelRealTimeReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CancelRealTimeReq.Merge(m, src)
}
func (m *CancelRealTimeReq) XXX_Size() int {
	return xxx_messageInfo_CancelRealTimeReq.Size(m)
}
func (m *CancelRealTimeReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CancelRealTimeReq.DiscardUnknown(m)
}

var xxx_messageInfo_CancelRealTimeReq proto.InternalMessageInfo

func (m *CancelRealTimeReq) GetHeader() *ReqHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type CancelRealTimeAck struct {
	Header               *AckHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *CancelRealTimeAck) Reset()         { *m = CancelRealTimeAck{} }
func (m *CancelRealTimeAck) String() string { return proto.CompactTextString(m) }
func (*CancelRealTimeAck) ProtoMessage()    {}
func (*CancelRealTimeAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{8}
}

func (m *CancelRealTimeAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CancelRealTimeAck.Unmarshal(m, b)
}
func (m *CancelRealTimeAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CancelRealTimeAck.Marshal(b, m, deterministic)
}
func (m *CancelRealTimeAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CancelRealTimeAck.Merge(m, src)
}
func (m *CancelRealTimeAck) XXX_Size() int {
	return xxx_messageInfo_CancelRealTimeAck.Size(m)
}
func (m *CancelRealTimeAck) XXX_DiscardUnknown() {
	xxx_messageInfo_CancelRealTimeAck.DiscardUnknown(m)
}

var xxx_messageInfo_CancelRealTimeAck proto.InternalMessageInfo

func (m *CancelRealTimeAck) GetHeader() *AckHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type RealTimeNotify struct {
	Header               *NotifyHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	IsConnect            bool          `protobuf:"varint,2,opt,name=IsConnect,proto3" json:"IsConnect,omitempty"`
	RealTimeInfo         *RealTimeInfo `protobuf:"bytes,3,opt,name=realTimeInfo,proto3" json:"realTimeInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *RealTimeNotify) Reset()         { *m = RealTimeNotify{} }
func (m *RealTimeNotify) String() string { return proto.CompactTextString(m) }
func (*RealTimeNotify) ProtoMessage()    {}
func (*RealTimeNotify) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{9}
}

func (m *RealTimeNotify) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RealTimeNotify.Unmarshal(m, b)
}
func (m *RealTimeNotify) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RealTimeNotify.Marshal(b, m, deterministic)
}
func (m *RealTimeNotify) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RealTimeNotify.Merge(m, src)
}
func (m *RealTimeNotify) XXX_Size() int {
	return xxx_messageInfo_RealTimeNotify.Size(m)
}
func (m *RealTimeNotify) XXX_DiscardUnknown() {
	xxx_messageInfo_RealTimeNotify.DiscardUnknown(m)
}

var xxx_messageInfo_RealTimeNotify proto.InternalMessageInfo

func (m *RealTimeNotify) GetHeader() *NotifyHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *RealTimeNotify) GetIsConnect() bool {
	if m != nil {
		return m.IsConnect
	}
	return false
}

func (m *RealTimeNotify) GetRealTimeInfo() *RealTimeInfo {
	if m != nil {
		return m.RealTimeInfo
	}
	return nil
}

type PollMessageReq struct {
	Header               *ReqHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *PollMessageReq) Reset()         { *m = PollMessageReq{} }
func (m *PollMessageReq) String() string { return proto.CompactTextString(m) }
func (*PollMessageReq) ProtoMessage()    {}
func (*PollMessageReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{10}
}

func (m *PollMessageReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PollMessageReq.Unmarshal(m, b)
}
func (m *PollMessageReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PollMessageReq.Marshal(b, m, deterministic)
}
func (m *PollMessageReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PollMessageReq.Merge(m, src)
}
func (m *PollMessageReq) XXX_Size() int {
	return xxx_messageInfo_PollMessageReq.Size(m)
}
func (m *PollMessageReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PollMessageReq.DiscardUnknown(m)
}

var xxx_messageInfo_PollMessageReq proto.InternalMessageInfo

func (m *PollMessageReq) GetHeader() *ReqHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type PollMessageAck struct {
	Header               *AckHeader                 `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Total                int32                      `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	Data                 []*PollMessageAck_DataItem `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *PollMessageAck) Reset()         { *m = PollMessageAck{} }
func (m *PollMessageAck) String() string { return proto.CompactTextString(m) }
func (*PollMessageAck) ProtoMessage()    {}
func (*PollMessageAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{11}
}

func (m *PollMessageAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PollMessageAck.Unmarshal(m, b)
}
func (m *PollMessageAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PollMessageAck.Marshal(b, m, deterministic)
}
func (m *PollMessageAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PollMessageAck.Merge(m, src)
}
func (m *PollMessageAck) XXX_Size() int {
	return xxx_messageInfo_PollMessageAck.Size(m)
}
func (m *PollMessageAck) XXX_DiscardUnknown() {
	xxx_messageInfo_PollMessageAck.DiscardUnknown(m)
}

var xxx_messageInfo_PollMessageAck proto.InternalMessageInfo

func (m *PollMessageAck) GetHeader() *AckHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *PollMessageAck) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *PollMessageAck) GetData() []*PollMessageAck_DataItem {
	if m != nil {
		return m.Data
	}
	return nil
}

type PollMessageAck_DataItem struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	FromUid              int64    `protobuf:"varint,2,opt,name=fromUid,proto3" json:"fromUid,omitempty"`
	ToUid                int64    `protobuf:"varint,3,opt,name=toUid,proto3" json:"toUid,omitempty"`
	GroupId              int64    `protobuf:"varint,4,opt,name=groupId,proto3" json:"groupId,omitempty"`
	MessageType          int32    `protobuf:"varint,5,opt,name=messageType,proto3" json:"messageType,omitempty"`
	MediaUrl             string   `protobuf:"bytes,6,opt,name=mediaUrl,proto3" json:"mediaUrl,omitempty"`
	CreateTime           int64    `protobuf:"varint,7,opt,name=createTime,proto3" json:"createTime,omitempty"`
	Memtions             []string `protobuf:"bytes,16,rep,name=memtions,proto3" json:"memtions,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PollMessageAck_DataItem) Reset()         { *m = PollMessageAck_DataItem{} }
func (m *PollMessageAck_DataItem) String() string { return proto.CompactTextString(m) }
func (*PollMessageAck_DataItem) ProtoMessage()    {}
func (*PollMessageAck_DataItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{11, 0}
}

func (m *PollMessageAck_DataItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PollMessageAck_DataItem.Unmarshal(m, b)
}
func (m *PollMessageAck_DataItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PollMessageAck_DataItem.Marshal(b, m, deterministic)
}
func (m *PollMessageAck_DataItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PollMessageAck_DataItem.Merge(m, src)
}
func (m *PollMessageAck_DataItem) XXX_Size() int {
	return xxx_messageInfo_PollMessageAck_DataItem.Size(m)
}
func (m *PollMessageAck_DataItem) XXX_DiscardUnknown() {
	xxx_messageInfo_PollMessageAck_DataItem.DiscardUnknown(m)
}

var xxx_messageInfo_PollMessageAck_DataItem proto.InternalMessageInfo

func (m *PollMessageAck_DataItem) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *PollMessageAck_DataItem) GetFromUid() int64 {
	if m != nil {
		return m.FromUid
	}
	return 0
}

func (m *PollMessageAck_DataItem) GetToUid() int64 {
	if m != nil {
		return m.ToUid
	}
	return 0
}

func (m *PollMessageAck_DataItem) GetGroupId() int64 {
	if m != nil {
		return m.GroupId
	}
	return 0
}

func (m *PollMessageAck_DataItem) GetMessageType() int32 {
	if m != nil {
		return m.MessageType
	}
	return 0
}

func (m *PollMessageAck_DataItem) GetMediaUrl() string {
	if m != nil {
		return m.MediaUrl
	}
	return ""
}

func (m *PollMessageAck_DataItem) GetCreateTime() int64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *PollMessageAck_DataItem) GetMemtions() []string {
	if m != nil {
		return m.Memtions
	}
	return nil
}

type PollReq struct {
	Header               *ReqHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *PollReq) Reset()         { *m = PollReq{} }
func (m *PollReq) String() string { return proto.CompactTextString(m) }
func (*PollReq) ProtoMessage()    {}
func (*PollReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{12}
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

//根据轮询结果请求相应消息数据
type PollAck struct {
	Header               *AckHeader         `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Msgs                 []*PollAck_Message `protobuf:"bytes,2,rep,name=msgs,proto3" json:"msgs,omitempty"`
	RealTime             *PollAck_RealTime  `protobuf:"bytes,3,opt,name=realTime,proto3" json:"realTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *PollAck) Reset()         { *m = PollAck{} }
func (m *PollAck) String() string { return proto.CompactTextString(m) }
func (*PollAck) ProtoMessage()    {}
func (*PollAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{13}
}

func (m *PollAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PollAck.Unmarshal(m, b)
}
func (m *PollAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PollAck.Marshal(b, m, deterministic)
}
func (m *PollAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PollAck.Merge(m, src)
}
func (m *PollAck) XXX_Size() int {
	return xxx_messageInfo_PollAck.Size(m)
}
func (m *PollAck) XXX_DiscardUnknown() {
	xxx_messageInfo_PollAck.DiscardUnknown(m)
}

var xxx_messageInfo_PollAck proto.InternalMessageInfo

func (m *PollAck) GetHeader() *AckHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *PollAck) GetMsgs() []*PollAck_Message {
	if m != nil {
		return m.Msgs
	}
	return nil
}

func (m *PollAck) GetRealTime() *PollAck_RealTime {
	if m != nil {
		return m.RealTime
	}
	return nil
}

type PollAck_Message struct {
	Count                int32    `protobuf:"varint,1,opt,name=Count,proto3" json:"Count,omitempty"`
	MessageKind          int32    `protobuf:"varint,2,opt,name=messageKind,proto3" json:"messageKind,omitempty"`
	GroupId              string   `protobuf:"bytes,3,opt,name=groupId,proto3" json:"groupId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PollAck_Message) Reset()         { *m = PollAck_Message{} }
func (m *PollAck_Message) String() string { return proto.CompactTextString(m) }
func (*PollAck_Message) ProtoMessage()    {}
func (*PollAck_Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{13, 0}
}

func (m *PollAck_Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PollAck_Message.Unmarshal(m, b)
}
func (m *PollAck_Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PollAck_Message.Marshal(b, m, deterministic)
}
func (m *PollAck_Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PollAck_Message.Merge(m, src)
}
func (m *PollAck_Message) XXX_Size() int {
	return xxx_messageInfo_PollAck_Message.Size(m)
}
func (m *PollAck_Message) XXX_DiscardUnknown() {
	xxx_messageInfo_PollAck_Message.DiscardUnknown(m)
}

var xxx_messageInfo_PollAck_Message proto.InternalMessageInfo

func (m *PollAck_Message) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *PollAck_Message) GetMessageKind() int32 {
	if m != nil {
		return m.MessageKind
	}
	return 0
}

func (m *PollAck_Message) GetGroupId() string {
	if m != nil {
		return m.GroupId
	}
	return ""
}

type PollAck_RealTime struct {
	IsConnect            bool          `protobuf:"varint,2,opt,name=IsConnect,proto3" json:"IsConnect,omitempty"`
	RealTimeInfo         *RealTimeInfo `protobuf:"bytes,3,opt,name=realTimeInfo,proto3" json:"realTimeInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *PollAck_RealTime) Reset()         { *m = PollAck_RealTime{} }
func (m *PollAck_RealTime) String() string { return proto.CompactTextString(m) }
func (*PollAck_RealTime) ProtoMessage()    {}
func (*PollAck_RealTime) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{13, 1}
}

func (m *PollAck_RealTime) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PollAck_RealTime.Unmarshal(m, b)
}
func (m *PollAck_RealTime) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PollAck_RealTime.Marshal(b, m, deterministic)
}
func (m *PollAck_RealTime) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PollAck_RealTime.Merge(m, src)
}
func (m *PollAck_RealTime) XXX_Size() int {
	return xxx_messageInfo_PollAck_RealTime.Size(m)
}
func (m *PollAck_RealTime) XXX_DiscardUnknown() {
	xxx_messageInfo_PollAck_RealTime.DiscardUnknown(m)
}

var xxx_messageInfo_PollAck_RealTime proto.InternalMessageInfo

func (m *PollAck_RealTime) GetIsConnect() bool {
	if m != nil {
		return m.IsConnect
	}
	return false
}

func (m *PollAck_RealTime) GetRealTimeInfo() *RealTimeInfo {
	if m != nil {
		return m.RealTimeInfo
	}
	return nil
}

func init() {
	proto.RegisterEnum("protocol.MessageType", MessageType_name, MessageType_value)
	proto.RegisterType((*SendMessageReq)(nil), "protocol.SendMessageReq")
	proto.RegisterType((*SendMessageAck)(nil), "protocol.SendMessageAck")
	proto.RegisterType((*MessageNotify)(nil), "protocol.MessageNotify")
	proto.RegisterType((*RealTimeReq)(nil), "protocol.RealTimeReq")
	proto.RegisterType((*RealTimeAck)(nil), "protocol.RealTimeAck")
	proto.RegisterType((*SimpleUser)(nil), "protocol.SimpleUser")
	proto.RegisterType((*RealTimeInfo)(nil), "protocol.RealTimeInfo")
	proto.RegisterType((*CancelRealTimeReq)(nil), "protocol.CancelRealTimeReq")
	proto.RegisterType((*CancelRealTimeAck)(nil), "protocol.CancelRealTimeAck")
	proto.RegisterType((*RealTimeNotify)(nil), "protocol.RealTimeNotify")
	proto.RegisterType((*PollMessageReq)(nil), "protocol.PollMessageReq")
	proto.RegisterType((*PollMessageAck)(nil), "protocol.PollMessageAck")
	proto.RegisterType((*PollMessageAck_DataItem)(nil), "protocol.PollMessageAck.DataItem")
	proto.RegisterType((*PollReq)(nil), "protocol.PollReq")
	proto.RegisterType((*PollAck)(nil), "protocol.PollAck")
	proto.RegisterType((*PollAck_Message)(nil), "protocol.PollAck.Message")
	proto.RegisterType((*PollAck_RealTime)(nil), "protocol.PollAck.RealTime")
}

func init() { proto.RegisterFile("chat.proto", fileDescriptor_8c585a45e2093e54) }

var fileDescriptor_8c585a45e2093e54 = []byte{
	// 783 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x55, 0xcd, 0x6e, 0xdb, 0x38,
	0x10, 0x5e, 0xfd, 0xd9, 0xf2, 0x38, 0x1b, 0x68, 0xb9, 0xd9, 0x40, 0x6b, 0x04, 0x0b, 0xaf, 0x4e,
	0x46, 0x16, 0xeb, 0x43, 0x8a, 0xa6, 0x40, 0x81, 0x02, 0x75, 0xd3, 0x43, 0x8d, 0x36, 0x41, 0xc0,
	0x24, 0xbd, 0xf4, 0xc4, 0x4a, 0x4c, 0x22, 0x58, 0x12, 0x1d, 0x89, 0x6e, 0x93, 0x37, 0x28, 0x7a,
	0xe8, 0x03, 0xf4, 0x41, 0xfa, 0x3c, 0x45, 0x5e, 0xa0, 0xaf, 0x50, 0x90, 0x14, 0x2d, 0x31, 0x41,
	0x0a, 0x38, 0x40, 0x2f, 0x3d, 0x89, 0xdf, 0x70, 0x86, 0x33, 0x9c, 0xef, 0xe3, 0x08, 0x20, 0x3e,
	0x27, 0x7c, 0x3c, 0x2f, 0x19, 0x67, 0xc8, 0x97, 0x9f, 0x98, 0x65, 0x83, 0xb5, 0x73, 0x4a, 0x12,
	0x5a, 0x2a, 0x7b, 0xf4, 0xc1, 0x86, 0xf5, 0x23, 0x5a, 0x24, 0xfb, 0xb4, 0xaa, 0xc8, 0x19, 0xc5,
	0xf4, 0x02, 0xfd, 0x07, 0x1d, 0xe5, 0x12, 0x5a, 0x43, 0x6b, 0xd4, 0xdf, 0xf9, 0x73, 0xac, 0x63,
	0xc7, 0x98, 0x5e, 0xbc, 0x90, 0x5b, 0xb8, 0x76, 0x41, 0x8f, 0xa0, 0x9f, 0xab, 0xd0, 0xe3, 0xab,
	0x39, 0x0d, 0xed, 0xa1, 0x35, 0x5a, 0xdf, 0xf9, 0xab, 0x89, 0xd8, 0x6f, 0x36, 0x71, 0xdb, 0x13,
	0x0d, 0xa1, 0xcf, 0xe9, 0x25, 0xaf, 0xf7, 0x43, 0x67, 0x68, 0x8d, 0x7a, 0xb8, 0x6d, 0x42, 0x03,
	0xf0, 0x73, 0x9a, 0xf3, 0x94, 0x15, 0x55, 0xe8, 0x0e, 0x9d, 0x51, 0x0f, 0x2f, 0x31, 0xda, 0x82,
	0x5e, 0xcc, 0x0a, 0x4e, 0x62, 0x3e, 0x4d, 0x42, 0x4f, 0xc6, 0x36, 0x06, 0x11, 0xb9, 0xa8, 0x68,
	0x59, 0x90, 0x9c, 0x86, 0x1d, 0xb9, 0xb9, 0xc4, 0x28, 0x84, 0xee, 0x59, 0xc9, 0x16, 0xf3, 0x69,
	0x12, 0x76, 0xe5, 0x96, 0x86, 0xd1, 0x13, 0xa3, 0x13, 0x93, 0x78, 0xf6, 0xa3, 0x4e, 0x4c, 0xe2,
	0x99, 0xd9, 0x89, 0xe8, 0xa3, 0x0d, 0xbf, 0xd7, 0xb1, 0x07, 0x8c, 0xa7, 0xa7, 0x57, 0x68, 0x7c,
	0x23, 0x7c, 0xb3, 0x09, 0x57, 0x1e, 0xbf, 0x6e, 0x2f, 0x3f, 0x59, 0xd0, 0xc7, 0x94, 0x64, 0xc7,
	0x69, 0xbe, 0xba, 0xa6, 0x8c, 0x82, 0xec, 0x9b, 0x05, 0xb5, 0x92, 0x3a, 0x46, 0x52, 0x51, 0xaa,
	0x3e, 0x35, 0x74, 0x55, 0xa9, 0x1a, 0x47, 0x5f, 0x5a, 0x05, 0xad, 0x4a, 0xad, 0x48, 0xc9, 0xe3,
	0xf9, 0x24, 0x49, 0xca, 0xba, 0x1c, 0x0d, 0xd1, 0x06, 0x78, 0x9c, 0xcd, 0x68, 0x51, 0x97, 0xa2,
	0x00, 0xda, 0x84, 0x4e, 0xc9, 0x58, 0x3e, 0x4d, 0x64, 0x19, 0x1e, 0xae, 0x91, 0xb0, 0xbf, 0xaf,
	0xe4, 0x31, 0xaa, 0xcd, 0x35, 0x12, 0x85, 0x9f, 0x73, 0xae, 0x12, 0xd4, 0x3d, 0xd6, 0x38, 0x7a,
	0x06, 0x70, 0x94, 0xe6, 0xf3, 0x8c, 0x9e, 0x54, 0x54, 0x7a, 0x16, 0x69, 0x3c, 0x93, 0x6c, 0x58,
	0xca, 0x53, 0x63, 0x51, 0x25, 0x79, 0x47, 0x38, 0x29, 0xb9, 0xae, 0xb2, 0x86, 0xd1, 0xb5, 0x05,
	0x6b, 0xfa, 0xf2, 0xd3, 0xe2, 0x94, 0xad, 0x58, 0x76, 0x00, 0xce, 0x22, 0x55, 0xd2, 0x70, 0xb0,
	0x58, 0xb6, 0x39, 0xe8, 0x98, 0x1c, 0x6c, 0x83, 0x27, 0xe4, 0x51, 0x85, 0xdd, 0xa1, 0x33, 0xea,
	0xef, 0x6c, 0x34, 0x6d, 0x6d, 0x6e, 0x81, 0x95, 0x8b, 0x71, 0x6d, 0xdf, 0xbc, 0x76, 0xab, 0x55,
	0x96, 0xd1, 0xaa, 0x3b, 0xa9, 0x88, 0x9e, 0xc2, 0x1f, 0x7b, 0xa4, 0x88, 0x69, 0x76, 0x5f, 0xdd,
	0xdd, 0x3e, 0x61, 0xe5, 0x19, 0xf0, 0xd9, 0x82, 0x75, 0x1d, 0x7c, 0xcf, 0x21, 0xb0, 0x05, 0xbd,
	0x69, 0xb5, 0xc7, 0x8a, 0x82, 0xc6, 0x8a, 0x47, 0x1f, 0x37, 0x06, 0xf4, 0x18, 0xd6, 0xca, 0x16,
	0x91, 0x92, 0x3f, 0xe3, 0xcc, 0x36, 0xcd, 0xd8, 0xf0, 0x15, 0xf3, 0xed, 0x90, 0x65, 0xd9, 0x3d,
	0x27, 0x7d, 0xf4, 0xcd, 0x36, 0xe2, 0x57, 0x7e, 0x44, 0x52, 0x73, 0x9c, 0x64, 0xf2, 0x52, 0x1e,
	0x56, 0x00, 0x3d, 0x04, 0x37, 0x21, 0x9c, 0x84, 0x8e, 0x94, 0xcb, 0xbf, 0xcd, 0x01, 0x66, 0xaa,
	0xf1, 0x73, 0xc2, 0xc9, 0x94, 0xd3, 0x1c, 0x4b, 0xf7, 0xc1, 0xb5, 0x05, 0xbe, 0x36, 0x21, 0x04,
	0xae, 0x98, 0x75, 0xb5, 0x52, 0xe4, 0x5a, 0xe8, 0xe4, 0xb4, 0x64, 0xf9, 0x49, 0xaa, 0x26, 0x88,
	0x83, 0x35, 0x54, 0x75, 0x08, 0xbb, 0x23, 0xed, 0x0a, 0xb4, 0x15, 0xed, 0x2a, 0x7f, 0xad, 0xe8,
	0xa1, 0x39, 0x95, 0x3d, 0x59, 0xbd, 0x31, 0x7e, 0xe5, 0x70, 0x4d, 0x52, 0x72, 0x52, 0x66, 0xfa,
	0xf9, 0x6a, 0x8c, 0xfe, 0x01, 0x88, 0x4b, 0x4a, 0x38, 0x15, 0x34, 0xc8, 0x29, 0xe9, 0xe0, 0x96,
	0xc5, 0x18, 0xcc, 0x81, 0x39, 0x98, 0xa3, 0x5d, 0xe8, 0x8a, 0x2e, 0xac, 0xcc, 0xd4, 0x57, 0x5b,
	0x05, 0xae, 0x4c, 0xd1, 0xff, 0xe0, 0xe6, 0xd5, 0x59, 0x15, 0xda, 0x92, 0x8c, 0xbf, 0x4d, 0x32,
	0x04, 0x0b, 0x5a, 0x3b, 0xd2, 0x0d, 0xed, 0x82, 0xaf, 0x05, 0x56, 0x0b, 0x71, 0x70, 0x3b, 0x64,
	0xf9, 0x1a, 0x97, 0xbe, 0x83, 0x37, 0xd0, 0xd5, 0xff, 0xa5, 0x0d, 0xf0, 0xf6, 0xd8, 0xa2, 0x50,
	0xdc, 0x79, 0x58, 0x81, 0x56, 0xcb, 0x5f, 0xa6, 0x45, 0x52, 0x0b, 0xa6, 0x6d, 0xba, 0xfb, 0x27,
	0x30, 0x48, 0xc0, 0xd7, 0x29, 0x7f, 0xde, 0x5b, 0xda, 0x3e, 0x80, 0x7e, 0xeb, 0x6f, 0x8c, 0x7c,
	0x70, 0x0f, 0x58, 0x41, 0x83, 0xdf, 0xc4, 0xea, 0x98, 0x5e, 0xf2, 0xc0, 0x42, 0x7d, 0xe8, 0x1e,
	0xa6, 0x31, 0x2f, 0x17, 0x34, 0xb0, 0x51, 0x0f, 0xbc, 0xc9, 0x22, 0x49, 0x59, 0xe0, 0x88, 0xe5,
	0xeb, 0x34, 0xa1, 0x2c, 0x70, 0x85, 0xf3, 0xab, 0xb4, 0x98, 0x05, 0xde, 0xdb, 0x8e, 0x4c, 0xfa,
	0xe0, 0x7b, 0x00, 0x00, 0x00, 0xff, 0xff, 0x01, 0x84, 0x4b, 0xb0, 0xb3, 0x09, 0x00, 0x00,
}
