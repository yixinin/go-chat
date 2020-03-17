package middleware

import (
	"chat/cache"

	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	var token, err = c.Cookie("token")
	if err != nil {
		return
	}
	uid, err := cache.GetToken(token)
	if err != nil {
		return
	}
	c.Set("uid", uid)
}
