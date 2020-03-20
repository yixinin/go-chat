package handler

import (
	"go-lib/log"
	"go-lib/utils"
	"net/http"

	"github.com/davyxu/cellnet/codec"
)

type Sender interface {
	Send(interface{})
	Close()
	ID() int64
}

type HttpSender struct {
	w http.ResponseWriter
}

func NewHttpSender(w http.ResponseWriter) Sender {
	return &HttpSender{
		w: w,
	}
}

func (s *HttpSender) Send(msg interface{}) {
	var buf, _, err = codec.EncodeMessage(msg, nil)
	if err != nil {
		log.Error(err)
		return
	}
	s.w.WriteHeader(http.StatusOK)
	n, err := s.w.Write(buf)
	if err != nil {
		log.Error(err)
		return
	}
	if n != len(buf) {
		log.Warnf("http writer not sent complete, sent:%d, expect:%d", n, len(buf))
	}
}

func (s *HttpSender) Close() {
	return
}

func (s *HttpSender) ID() int64 {
	return utils.RandInt64(1024)
}
