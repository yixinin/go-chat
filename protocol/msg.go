package protocol

var ProtocolMap = make(map[string][]interface{}, 1)

func init() {
	var protocol = make([]interface{}, 0, 100)
	protocol = append(protocol,
		(*EchoReq)(nil),
		(*EchoAck)(nil),
		// (*LoginAck)(nil),
		// (*LogoutReq)(nil),
		// (*LogoutAck)(nil),
		// (*GetGameRoomTypeListReq)(nil),
		// (*GetGameRoomTypeListAck)(nil),
	)

	ProtocolMap["protocol"] = protocol
}
