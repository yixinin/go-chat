package iface

type Session struct {
	sender Sender
	Uid    int64
	Token  string
}

func (s *Session) Send(v interface{}) {
	if s.sender != nil {
		s.sender.Send(v)
	}
}

func (s *Session) Close() {
	if s.sender != nil {
		s.sender.Close()
	}
}

func (s *Session) ID() int64 {
	if s.sender != nil {
		return s.sender.ID()
	}
	return 0
}

func NewSessoin(sender Sender, uid int64, token string) *Session {
	return &Session{
		sender: sender,
		Uid:    uid,
		Token:  token,
	}
}
