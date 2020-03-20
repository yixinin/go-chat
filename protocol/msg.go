package protocol

var ProtocolMap = make(map[string][]interface{}, 1)

func init() {
	var protocol = make([]interface{}, 0, 100)
	protocol = append(protocol,
		(*EchoReq)(nil),
		(*EchoAck)(nil),
		new(SignUpReq),
		new(SignUpAck),
		new(SignInReq),
		new(SignInAck),
		new(SendMessageReq),
		new(SendMessageAck),
		new(MessageNotify),
		new(AddContactReq),
		new(AddContactAck),
	)

	ProtocolMap["protocol"] = protocol
}
