package api

import (
	"GoZhixue/service"
	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context) {
	var service service.UserRegisterService
	if err := c.ShouldBindJSON(&service); err == nil {
		res := service.Register()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

func UserLogin(c *gin.Context) {
	var service service.UserLoginService
	if err := c.ShouldBindJSON(&service); err == nil {
		res := service.Login()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

func UserLogout(c *gin.Context) {
	var service service.UserLogoutService
	if err := c.ShouldBindJSON(&service); err == nil {
		res := service.Logout()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
