package logic

import "chat/protocol"

type Reqer interface {
	GetHeader() *protocol.ReqHeader
}

type Acker interface {
	GetHeader() *protocol.AckHeader
}

type Notifier interface {
	GetHeader() *protocol.NotifyHeader
}
