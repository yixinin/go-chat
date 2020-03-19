package handler

import (
	"chat/handler/middleware"
	"chat/logic"
	"chat/protocol"

	"go-lib/registry"

	"github.com/gin-gonic/gin"
)

type ChatHandler struct {
	logic  *logic.ChatLogic
	Handle HttpHandler
}

func NewChatHandler(regist registry.Registry, notifyFuncs ...logic.NotifyFunc) *ChatHandler {

	return &ChatHandler{
		logic: logic.NewChatLogic(regist, notifyFuncs...),
	}
}

func (h *ChatHandler) String() string {
	return "handler.ChatHandler"
}

func (h *ChatHandler) HandleAll(g *gin.Engine) error {
	var group = g.Group("/chat")
	group.Use(middleware.Auth)
	group.POST("/send", h.SendMessage)
	group.POST("/realTime", h.RealTime)
	group.POST("/cancelRealTime", h.CancelRealTime)
	group.POST("/pollRealTimeNotify", h.PollRealTimeNotify)
	group.POST("/pollMessage", h.PollMessage)
	return nil
}

func (h *ChatHandler) SendMessage(c *gin.Context) {
	var req protocol.SendMessageReq
	var ack protocol.SendMessageAck
	h.Handle(c, &req, &ack, h.logic.SendMessage)
}

//请求语音/视频
func (h *ChatHandler) RealTime(c *gin.Context) {
	var req protocol.RealTimeReq
	var ack protocol.RealTimeAck
	h.Handle(c, &req, &ack, h.logic.RealTime)
}

func (h *ChatHandler) CancelRealTime(c *gin.Context) {
	var req protocol.CancelRealTimeReq
	var ack protocol.CancelRealTimeAck
	h.Handle(c, &req, &ack, h.logic.CancelRealTime)
}

func (h *ChatHandler) PollRealTimeNotify(c *gin.Context) {
	var req protocol.PollReq
	var ack protocol.PollAck
	h.Handle(c, &req, &ack, h.logic.Poll)
}

func (h *ChatHandler) PollMessage(c *gin.Context) {
	var req protocol.PollMessageReq
	var ack protocol.PollMessageAck
	h.Handle(c, &req, &ack, h.logic.PollMessage)
}
