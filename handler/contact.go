package handler

import (
	"chat/logic"

	"github.com/gin-gonic/gin"
)

type ContactHandler struct {
	logic  *logic.ChatLogic
	Handle HttpHandler
}

func (h *ContactHandler) AddContact(c *gin.Context) {

}

func (h *ContactHandler) DeleteContact(c *gin.Context) {

}

func (h *ContactHandler) UpdateContact(c *gin.Context) {

}
