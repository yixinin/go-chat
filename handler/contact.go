package handler

import (
	"chat/logic"
	"chat/protocol"

	"github.com/gin-gonic/gin"
)

type ContactHandler struct {
	logic  *logic.ContactLogic
	Handle HttpHandler
}

func (h *ContactHandler) SearchUser(c *gin.Context) {
	h.Handle(c, &protocol.SearchUserReq{}, &protocol.SearchUserAck{}, h.logic.SearchUser)
}

func (h *ContactHandler) AddContact(c *gin.Context) {
	h.Handle(c, &protocol.AddContactReq{}, &protocol.AddContactAck{}, h.logic.AddContact)
}

func (h *ContactHandler) DeleteContact(c *gin.Context) {
	h.Handle(c, &protocol.DeleteContactReq{}, &protocol.DeleteContactAck{}, h.logic.DeleteContact)
}

func (h *ContactHandler) UpdateContact(c *gin.Context) {
	h.Handle(c, &protocol.UpdateContactReq{}, &protocol.UpdateContactAck{}, h.logic.UpdateContact)
}
