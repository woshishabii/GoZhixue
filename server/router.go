/**
 * @Author: woshishabii
 * @Description:
 * @File: router
 * @Version: 0.0.1
 * @Date: 4/15/2023 4:21 PM
 */

package server

import (
	"GoZhixue/api"
	"GoZhixue/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	// Middleware
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.Cors())
	// r.Use(middleware.CurrentUser())

	v1 := r.Group("/api/v1")
	{
		v1.GET("ping", api.Ping)
		v1.POST("ping", api.Ping)

		user := v1.Group("user")
		{
			user.POST("register", api.UserRegister)
			user.POST("login", api.UserLogin)
			user.POST("logout", api.UserLogout)
		}

		school := v1.Group("school")
		{
			school.GET("list", api.GetSchools)
		}

		class := v1.Group("class")
		{
			class.POST("list", api.GetClasses)
		}
	}

	return r
}
