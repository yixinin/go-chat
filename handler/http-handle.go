package handler

import (
	"chat/cache"
	"chat/handler/iface"
	"chat/logic"
	"chat/protocol"
	"go-lib/utils"
	"io/ioutil"
	"net/http"
	"reflect"

	log "github.com/sirupsen/logrus"

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
	defer func() {
		if err := recover(); err != nil {
			log.Errorf("http handler recovered, err:%v", err)
		}
	}()
	var buf, err = ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unexpect error"))
		return
	}
	if len(buf) > 0 {
		var msgid = utils.GetMsgId(buf[2:4])
		msg, _, err := codec.DecodeMessage(int(msgid), buf[4:])
		if err != nil {
			log.Error(err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Unexpect decode error"))
			return
		}
		//读取cookie
		uid, token, ok := h.Auth(r, msg)
		if msg != nil {
			log.Infof("recv http msg:%+v", msg)
			var sender = iface.NewHttpSender(w)
			if ok && uid > 0 {
				h.logic.acceptFunc(iface.NewSessoin(sender, uid, token))
			}
			h.logic.handleMessage(sender, msg)
		}
	}
}

func (h *Http) Auth(r *http.Request, msg interface{}) (int64, string, bool) {
	c, err := r.Cookie("token")
	if err != nil {
		log.Error(err)
	}
	reqer, ok := msg.(logic.Reqer)
	if !ok {
		return 0, "", false
	}
	var header = reqer.GetHeader()
	if header == nil {
		header = &protocol.ReqHeader{}
		var v = reflect.Indirect(reflect.ValueOf(msg))
		v.FieldByName("Header").Set(reflect.ValueOf(header))
	}
	if c != nil && c.Value != "" && header.Token == "" {
		header.Token = c.Value
	}
	if header.Token != "" {
		uid, err := cache.GetToken(header.Token)
		if err != nil {
			log.Error(err)
			return 0, header.Token, false
		}
		header.Uid = uid
	}

	return header.Uid, header.Token, header.Uid > 0
}

func (h *Http) String() string {
	return "http"
}
