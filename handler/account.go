package handler

import (
	"go-chat/logic"
	"go-chat/protocol"
	"go-lib/log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AccountHandler struct {
	logic *logic.AccountLogic
}

func (h *AccountHandler) String() string {
	return "handler.AccountHandler"
}

func (h *AccountHandler) HandleAll(g *gin.Engine) error {
	g.POST("/signUp", h.SignUp)
	return nil
}

func (h *AccountHandler) SignUp(c *gin.Context) {
	var req protocol.SignUpReq
	var ack protocol.SignUpAck
	if err := c.ShouldBindJSON(&req); err != nil {
		ack.Header.Code = 400
		ack.Header.Msg = "parse json error"
		c.JSON(http.StatusBadRequest, ack)
		log.Errorf("parse json error: req:%v, error:%v", req, err)
		return
	}
	if err := h.logic.SignUp(&req, &ack); err != nil {
		if ack.Header.Code == 0 {
			ack.Header.Code = 400
		}
		if ack.Header.Msg == "" {
			ack.Header.Msg = "unexpect error"
		}
		log.Errorf("req:%v, error:%v", req, err)
	}
	c.JSON(http.StatusOK, ack)
}

func (h *AccountHandler) SignIn(c *gin.Context) {
	var req protocol.SignInReq
	var ack protocol.SignInAck
	if err := c.ShouldBindJSON(&req); err != nil {
		ack.Header.Code = 400
		ack.Header.Msg = "parse json error"
		c.JSON(http.StatusBadRequest, ack)
		log.Errorf("parse json error: req:%v, error:%v", req, err)
		return
	}
	if err := h.logic.SignIn(&req, &ack); err != nil {
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

func (h *AccountHandler) SignOut(c *gin.Context) {
	var req protocol.SignOutReq
	var ack protocol.SignOutAck
	if err := c.ShouldBindJSON(&req); err != nil {
		ack.Header.Code = 400
		ack.Header.Msg = "parse json error"
		c.JSON(http.StatusBadRequest, ack)
		log.Errorf("parse json error: req:%v, error:%v", req, err)
		return
	}
	if err := h.logic.SignOut(&req, &ack); err != nil {
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

func (h *AccountHandler) Delete(c *gin.Context) {
	var req protocol.DeleteReq
	var ack protocol.DeleteAck
	if err := c.ShouldBindJSON(&req); err != nil {
		ack.Header.Code = 400
		ack.Header.Msg = "parse json error"
		c.JSON(http.StatusBadRequest, ack)
		log.Errorf("parse json error: req:%v, error:%v", req, err)
		return
	}
	if err := h.logic.Delete(&req, &ack); err != nil {
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

func (h *AccountHandler) ChangePassword(c *gin.Context) {
	var req protocol.ChangePasswordReq
	var ack protocol.ChangePasswordAck
	if err := c.ShouldBindJSON(&req); err != nil {
		ack.Header.Code = 400
		ack.Header.Msg = "parse json error"
		c.JSON(http.StatusBadRequest, ack)
		log.Errorf("parse json error: req:%v, error:%v", req, err)
		return
	}
	if err := h.logic.ChangePassword(&req, &ack); err != nil {
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

func (h *AccountHandler) ResetPassword(c *gin.Context) {
	var req protocol.ResetPasswordReq
	var ack protocol.ResetPasswordAck
	if err := c.ShouldBindJSON(&req); err != nil {
		ack.Header.Code = 400
		ack.Header.Msg = "parse json error"
		c.JSON(http.StatusBadRequest, ack)
		log.Errorf("parse json error: req:%v, error:%v", req, err)
		return
	}
	if err := h.logic.ResetPassword(&req, &ack); err != nil {
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
