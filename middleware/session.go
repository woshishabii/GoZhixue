/**
 * @Author: woshishabii
 * @Description:
 * @File: session
 * @Version: 0.0.1
 * @Date: 4/15/2023 4:10 PM
 */

package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// Session 初始化session
func Session(secret string) gin.HandlerFunc {
	store := cookie.NewStore([]byte(secret))
	// Also set Secure: true id using SSL, you should though
	store.Options(sessions.Options{HttpOnly: true, MaxAge: 7 * 86400, Path: "/"})
	return sessions.Sessions("gin-session", store)
}
