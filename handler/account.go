package handler

// import (
// 	"chat/handler/middleware"
// 	"chat/logic"
// 	"chat/protocol"

// 	"github.com/gin-gonic/gin"
// )

// type AccountHandler struct {
// 	logic  *logic.AccountLogic
// 	Handle HttpHandler
// }

// func NewAccountHandler() *AccountHandler {
// 	return &AccountHandler{
// 		logic: &logic.AccountLogic{},
// 	}
// }

// func (h *AccountHandler) String() string {
// 	return "handler.AccountHandler"
// }

// func (h *AccountHandler) HandleAll(g *gin.Engine) error {
// 	var group = g.Group("/account")

// 	// group.POST("/signUp")
// 	group.POST("/signUp", h.SignUp)
// 	group.POST("/signIn", middleware.Auth, h.SignIn)
// 	group.POST("/signOut", middleware.Auth, h.SignOut)
// 	group.POST("/delete", middleware.Auth, h.Delete)
// 	group.POST("/changePassword", middleware.Auth, h.ChangePassword)
// 	group.POST("/resetPassword", h.ResetPassword)
// 	return nil
// }

// func (h *AccountHandler) SignUp(c *gin.Context) {
// 	h.Handle(c, &protocol.SignUpReq{}, h.logic.SignUp)
// }

// func (h *AccountHandler) SignIn(c *gin.Context) {
// 	var req protocol.SignInReq
// 	h.Handle(c, &req, h.logic.SignIn)
// }

// func (h *AccountHandler) SignOut(c *gin.Context) {
// 	var req protocol.SignOutReq
// 	h.Handle(c, &req, h.logic.SignOut)
// }

// func (h *AccountHandler) Delete(c *gin.Context) {
// 	var req protocol.DeleteReq
// 	h.Handle(c, &req, h.logic.Delete)
// }

// func (h *AccountHandler) ChangePassword(c *gin.Context) {
// 	var req protocol.ChangePasswordReq
// 	h.Handle(c, &req, h.logic.ChangePassword)
// }

// func (h *AccountHandler) ResetPassword(c *gin.Context) {
// 	var req protocol.ResetPasswordReq
// 	h.Handle(c, &req, h.logic.ResetPassword)
// }
