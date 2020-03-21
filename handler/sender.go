package handler

import (
	"encoding/binary"
	"go-lib/utils"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/davyxu/cellnet/codec"
)

const (
	HeaderSize = 4
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
	var msgData, meta, err = codec.EncodeMessage(msg, nil)
	// var body = make([]byte, len(buf)+2)
	pktData := make([]byte, HeaderSize+len(msgData))

	// 写入消息长度做验证
	binary.LittleEndian.PutUint16(pktData, uint16(HeaderSize+len(msgData)))

	// Type
	binary.LittleEndian.PutUint16(pktData[2:], uint16(meta.ID))

	// Value
	copy(pktData[HeaderSize:], msgData)
	s.w.WriteHeader(http.StatusOK)
	n, err := s.w.Write(pktData)
	if err != nil {
		log.Error(err)
		return
	}
	if n != len(pktData) {
		log.Warnf("http writer not sent complete, sent:%d, expect:%d", n, len(pktData))
	}
}

func (s *HttpSender) Close() {
	return
}

func (s *HttpSender) ID() int64 {
	return utils.RandInt64(1024)
}
