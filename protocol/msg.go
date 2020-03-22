package protocol

var ProtocolMap = make(map[string][]interface{}, 1)

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
	)

	ProtocolMap["protocol"] = protocol
}
