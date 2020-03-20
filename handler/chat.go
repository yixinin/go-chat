package handler

// import (
// 	"chat/handler/middleware"
// 	"chat/logic"
// 	"chat/protocol"

// 	"go-lib/registry"

// 	"github.com/gin-gonic/gin"
// )

// type ChatHandler struct {
// 	logic  *logic.ChatLogic
// 	Handle HttpHandler
// }

// func NewChatHandler(regist registry.Registry, notifyFuncs ...logic.NotifyFunc) *ChatHandler {

// 	return &ChatHandler{
// 		logic: logic.NewChatLogic(regist, notifyFuncs...),
// 	}
// }

// func (h *ChatHandler) String() string {
// 	return "handler.ChatHandler"
// }

// func (h *ChatHandler) HandleAll(g *gin.Engine) error {
// 	var group = g.Group("/chat")
// 	group.Use(middleware.Auth)
// 	group.POST("/send", h.SendMessage)
// 	group.POST("/realTime", h.RealTime)
// 	group.POST("/cancelRealTime", h.CancelRealTime)
// 	group.POST("/pollRealTimeNotify", h.PollRealTimeNotify)
// 	group.POST("/pollMessage", h.PollMessage)
// 	return nil
// }

// func (h *ChatHandler) SendMessage(c *gin.Context) {
// 	var req protocol.SendMessageReq
// 	h.Handle(c, &req, h.logic.SendMessage)
// }

// //请求语音/视频
// func (h *ChatHandler) RealTime(c *gin.Context) {
// 	var req protocol.RealTimeReq
// 	h.Handle(c, &req, h.logic.RealTime)
// }

// func (h *ChatHandler) CancelRealTime(c *gin.Context) {
// 	var req protocol.CancelRealTimeReq
// 	h.Handle(c, &req, h.logic.CancelRealTime)
// }

// func (h *ChatHandler) PollRealTimeNotify(c *gin.Context) {
// 	var req protocol.PollReq
// 	h.Handle(c, &req, h.logic.Poll)
// }

// func (h *ChatHandler) PollMessage(c *gin.Context) {
// 	var req protocol.PollMessageReq
// 	h.Handle(c, &req, h.logic.PollMessage)
// }
