package handler

import (
	"go-chat/logic"
	"go-chat/protocol"
	"go-lib/log"
	"net/http"

	"go-lib/registry"

	"github.com/gin-gonic/gin"
)

type MessageHandler struct {
	logic *logic.ChatLogic
}

func NewMessageHandler(watcher registry.Watcher, notifyFuncs ...logic.NotifyFunc) *MessageHandler {

	return &MessageHandler{
		logic: logic.NewChatLogic(watcher, notifyFuncs...),
	}
}

func (h *MessageHandler) String() string {
	return "handler.MessageHandler"
}

func (h *MessageHandler) HandleAll(g *gin.Engine) error {
	g.POST("/sendMessage", h.SendMessage)
	g.POST("/realTime", h.RealTime)
	g.POST("/pollnotify", h.Pollnotify)
	return nil
}

func (h *MessageHandler) SendMessage(c *gin.Context) {
	var req protocol.SendMessageReq
	var ack protocol.SendMessageAck
	if err := c.ShouldBindJSON(&req); err != nil {
		ack.Header.Code = 400
		ack.Header.Msg = "parse json error"
		c.JSON(http.StatusBadRequest, ack)
		log.Errorf("parse json error: req:%v, error:%v", req, err)
		return
	}
	if err := h.logic.SendMessage(&req, &ack); err != nil {
		log.Errorf("req:%v, error:%v", req, err)
		if ack.Header.Code == 0 {
			ack.Header.Code = 400
		}
		if ack.Header.Msg == "" {
			ack.Header.Msg = "unexpect error"
		}
	}
	c.JSON(http.StatusOK, ack)
}

//请求语音/视频
func (h *MessageHandler) RealTime(c *gin.Context) {
	var req protocol.RealTimeReq
	var ack protocol.RealTimeAck
	if err := c.ShouldBindJSON(&req); err != nil {
		ack.Header.Code = 400
		ack.Header.Msg = "parse json error"
		c.JSON(http.StatusBadRequest, ack)
		log.Errorf("parse json error: req:%v, error:%v", req, err)
		return
	}
	if err := h.logic.RealTime(&req, &ack); err != nil {
		log.Errorf("req:%v, error:%v", req, err)
		if ack.Header.Code == 0 {
			ack.Header.Code = 400
		}
		if ack.Header.Msg == "" {
			ack.Header.Msg = "unexpect error"
		}
	}
	c.JSON(http.StatusOK, ack)
}

func (h *MessageHandler) CancelRealTime(c *gin.Context) {
	var req protocol.CancelRealTimeReq
	var ack protocol.CancelRealTimeAck
	if err := c.ShouldBindJSON(&req); err != nil {
		ack.Header.Code = 400
		ack.Header.Msg = "parse json error"
		c.JSON(http.StatusBadRequest, ack)
		log.Errorf("parse json error: req:%v, error:%v", req, err)
		return
	}
	if err := h.logic.CancelRealTime(&req, &ack); err != nil {
		log.Errorf("req:%v, error:%v", req, err)
		if ack.Header.Code == 0 {
			ack.Header.Code = 400
		}
		if ack.Header.Msg == "" {
			ack.Header.Msg = "unexpect error"
		}
	}
	c.JSON(http.StatusOK, ack)
}

func (h *MessageHandler) Pollnotify(c *gin.Context) {
	var req protocol.PollNotifyReq
	var ack protocol.PollNotifyAck
	if err := c.ShouldBindJSON(&req); err != nil {
		ack.Header.Code = 400
		ack.Header.Msg = "parse json error"
		c.JSON(http.StatusBadRequest, ack)
		log.Errorf("parse json error: req:%v, error:%v", req, err)
		return
	}
	if err := h.logic.PollNotify(&req, &ack); err != nil {
		log.Errorf("req:%v, error:%v", req, err)
		if ack.Header.Code == 0 {
			ack.Header.Code = 400
		}
		if ack.Header.Msg == "" {
			ack.Header.Msg = "unexpect error"
		}
	}
	c.JSON(http.StatusOK, ack)
}
