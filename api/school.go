package api

import (
	"GoZhixue/model"
	"GoZhixue/serializer"
	"GoZhixue/service"

	"github.com/gin-gonic/gin"
)

func GetSchools(c *gin.Context) {
	schools := []model.School{}
	err := model.DB.Find(&schools)
	if err.Error != nil {
		c.JSON(200, serializer.DBErr("数据库错误", err.Error))
	} else {
		c.JSON(200, serializer.Response{
			Code: 20000,
			Msg:  "成功",
			Data: schools,
		})
	}
}

func GetClasses(c *gin.Context) {
	var service service.GetClassesService
	if err := c.ShouldBindJSON(&service); err == nil {
		res := service.Get()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
