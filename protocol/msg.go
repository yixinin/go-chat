package protocol

type Reqer interface {
	GetHeader() *ReqHeader
}

type Acker interface {
	GetHeader() *AckHeader
}

type Notifier interface {
	GetHeader() *NotifyHeader
}

var ProtocolMap = make(map[string][]interface{}, 1)

var ReqAcks = make(map[string]map[Reqer]Acker)
var Notifies = make(map[string][]Notifier)

func init() {

	var protocol = make([]interface{}, 0, 100)
	protocol = append(protocol,
		(*EchoReq)(nil),
		(*EchoAck)(nil),

		//账户
		new(SignUpReq),
		new(SignUpAck),
		new(SignInReq),
		new(SignInAck),

		//消息
		new(SendMessageReq),
		new(SendMessageAck),
		new(MessageNotify),
		new(RealTimeAck),
		new(RealTimeReq),
		new(RealTimeNotify),
		new(GetMessageUserReq),
		new(GetMessageUserAck),
		new(GetMessageReq),
		new(GetMessageAck),

		//联系人
		new(AddContactReq),
		new(AddContactAck),
		new(GetContactListReq),
		new(GetContactListAck),
	)

	ProtocolMap["protocol"] = protocol
}

func BuildRas() {
	var ra = map[Reqer]Acker{
		new(EchoReq): new(EchoAck),
	}
	ReqAcks["protocol"] = ra
}

func BuildNs() {
	var n = []Notifier{}

	Notifies["protocol"] = n
}
