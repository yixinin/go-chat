package middleware

import (
	"chat/cache"

	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	var v, ok = c.Get("uid")
	if !ok {
		c.String(401, "access denied")
		return
	}
	if uid, ok := v.(int64); !ok || uid <= 0 {
		c.String(401, "access denied")
		return
	}
	c.Next()
}

func SetUid(c *gin.Context) {
	var token, err = c.Cookie("token")
	if err != nil {
		c.Next()
		return
	}
	uid, err := cache.GetToken(token)
	if err != nil {
		c.Next()
		return
	}
	c.Set("uid", uid)
	c.Next()
}
