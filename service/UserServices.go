package service

import (
	"GoZhixue/model"
	"GoZhixue/serializer"
	"github.com/google/uuid"
	"time"
)

type UserRegisterService struct {
	Username string `form:"username" json:"username" binding:"required,min=2,max=30"`
	Password string `form:"password" json:"password" binding:"required,len=32"`
	RealName string `form:"real_name" json:"real_name" binding:"required"`
	SchoolID uint   `form:"school_id" json:"school_id" binding:"required"`
	Type     uint   `form:"type" json:"type" binding:"required"`
}

func (service *UserRegisterService) Valid() *serializer.Response {
	var count int64
	if model.DB.Model(&model.User{}).Where("username = ?", service.Username).Count(&count); count > 0 {
		return &serializer.Response{
			Code: 40001,
			Msg:  "用户名已存在",
		}
	}

	count = 0
	if model.DB.Model(&model.School{}).Where("id = ?", service.SchoolID).Count(&count); count == 0 {
		return &serializer.Response{
			Code: 40002,
			Msg:  "指定的学校不存在",
		}
	}

	return nil
}

func (service UserRegisterService) Register() serializer.Response {
	if err := service.Valid(); err != nil {
		return *err
	}

	user := model.User{
		Username: service.Username,
		Password: service.Password,
		RealName: service.RealName,
		SchoolID: service.SchoolID,
		Type:     service.Type,
	}

	if err := model.DB.Create(&user).Error; err != nil {
		return serializer.Response{
			Code: 50001,
			Msg:  "数据库错误",
		}
	}

	return serializer.Response{
		Code: 20001,
		Data: model.User{
			Username: user.Username,
			RealName: user.RealName,
			School:   user.School,
		},
		Msg: "注册成功",
	}
}

type UserLoginService struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (service UserLoginService) Login() serializer.Response {
	var count int64
	var user model.User
	if model.DB.Model(&model.User{}).Where("username = ?", service.Username).Count(&count); count == 0 {
		return serializer.Response{
			Code: 40004,
			Msg:  "账户不存在",
		}
	} else {
		model.DB.Where("username = ?", service.Username).First(&user)
		if service.Password != user.Password {
			return serializer.Response{
				Code: 40005,
				Msg:  "密码错误",
			}
		} else {
			// 清理过期Session
			sessions := [...]model.Session{}
			model.DB.Where("user_id = ?", user.ID).Find(&sessions)
			delta, _ := time.ParseDuration("289h")
			for _, session := range sessions {
				if time.Now().Sub(session.CreatedAt) > delta {
					model.DB.Delete(&session)
				}
			}

			count = 0
			session := model.Session{UserID: user.ID, UUID: uuid.NewString()}
			if model.DB.Model(&model.Session{}).Where("user_id = ?", user.ID).Count(&count); count > 3 {
				return serializer.Response{
					Code: 40006,
					Msg:  "登录设备过多",
				}
			}
			model.DB.Create(&session)
			return serializer.Response{
				Code: 20002,
				Msg:  "登陆成功",
				Data: session,
			}
		}
	}
}

type UserLogoutService struct {
	Session string `json:"session" binding:"required,len=36"`
}

func (service UserLogoutService) Logout() serializer.Response {
	var count int64
	var session model.Session
	if model.DB.Model(&model.Session{}).Where("UUID = ?", service.Session).Count(&count); count == 0 {
		return serializer.Response{
			Code: 40004,
			Msg:  "未找到会话",
		}
	}
	model.DB.Where("UUID = ?", service.Session).First(&session)
	model.DB.Delete(&session)
	return serializer.Response{
		Code: 20000,
		Msg:  "成功登出",
	}
}
