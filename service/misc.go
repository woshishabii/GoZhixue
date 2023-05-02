package service

import (
	"GoZhixue/model"
	"GoZhixue/serializer"
)

type GetClassesService struct {
	SchoolID uint `json:"school_id" binding:"required"`
}

func (service GetClassesService) Get() serializer.Response {
	classes := []model.Class{}
	if err := model.DB.Model(&model.Class{}).Where("school_id = ?", service.SchoolID).Find(&classes); err.Error != nil {
		return serializer.DBErr("数据库错误", err.Error)
	} else {
		return serializer.Response{
			Code: 20000,
			Msg:  "成功",
			Data: classes,
		}
	}
}
