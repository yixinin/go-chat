package handler

import (
	"chat/cache"
	"chat/logic"
	"go-lib/log"
	"go-lib/utils"
	"io/ioutil"
	"net/http"

	"github.com/davyxu/cellnet/codec"
)

type Http struct {
	logic *Logic
}

func NewHttp(l *Logic) *Http {
	return &Http{
		logic: l,
	}
}

func (h *Http) Name() string {
	return "http"
}

func (h *Http) Handle(w http.ResponseWriter, r *http.Request) {
	var buf, err = ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unexpect error"))
		return
	}
	if len(buf) > 0 {
		var msgid = utils.BytesToUint16(buf[:2])
		msg, _, err := codec.DecodeMessage(int(msgid), buf[2:])
		if err != nil {
			log.Error(err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Unexpect decode error"))
			return
		}
		//读取cookie
		h.Auth(r, msg)
		if msg != nil {
			var sender = NewHttpSender(w)
			h.logic.handleMessage(sender, msg)
		}
	}
}

func (h *Http) Auth(r *http.Request, msg interface{}) bool {
	c, err := r.Cookie("token")
	if err != nil {
		log.Error(err)
		return false
	}
	var reqer, ok = msg.(logic.Reqer)
	if !ok {
		return false
	}
	var header = reqer.GetHeader()
	header.Token = c.Value
	uid, err := cache.GetToken(header.Token)
	if err != nil {
		log.Error(err)
		return false
	}
	header.Uid = uid
	return uid > 0
}

func (h *Http) String() string {
	return "http"
}
