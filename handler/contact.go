package handler

// import (
// 	"chat/handler/middleware"
// 	"chat/logic"
// 	"chat/protocol"

// 	"github.com/gin-gonic/gin"
// )

// type ContactHandler struct {
// 	logic  *logic.ContactLogic
// 	Handle HttpHandler
// }

// func (h *ContactHandler) HandleAll(g *gin.Engine) {
// 	var group = g.Group("/contact")

// 	group.Use(middleware.Auth)
// 	group.POST("/searchUser", h.SearchUser)
// 	group.POST("/addContact", h.AddContact)
// 	group.POST("/deleteContact", h.DeleteContact)
// 	group.POST("/updateContact", h.UpdateContact)
// }

// func (h *ContactHandler) SearchUser(c *gin.Context) {
// 	h.Handle(c, &protocol.SearchUserReq{}, h.logic.SearchUser)
// }

// func (h *ContactHandler) AddContact(c *gin.Context) {
// 	h.Handle(c, &protocol.AddContactReq{}, h.logic.AddContact)
// }

// func (h *ContactHandler) DeleteContact(c *gin.Context) {
// 	h.Handle(c, &protocol.DeleteContactReq{}, h.logic.DeleteContact)
// }

// func (h *ContactHandler) UpdateContact(c *gin.Context) {
// 	h.Handle(c, &protocol.UpdateContactReq{}, h.logic.UpdateContact)
// }
