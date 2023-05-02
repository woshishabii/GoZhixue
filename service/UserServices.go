package service

import (
	"GoZhixue/model"
	"GoZhixue/serializer"
	"time"

	"github.com/google/uuid"
)

type UserRegisterService struct {
	Username string `form:"username" json:"username" binding:"required,min=2,max=30"`
	Password string `form:"password" json:"password" binding:"required,len=32"`
	RealName string `form:"real_name" json:"real_name" binding:"required"`
	ClassID  uint   `form:"class_id" json:"class_id" binding:"required"`
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
	if model.DB.Model(&model.Class{}).Where("id = ?", service.ClassID).Count(&count); count == 0 {
		return &serializer.Response{
			Code: 40002,
			Msg:  "指定的学校/学校不存在",
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
		ClassID:  service.ClassID,
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
		Data: user,
		Msg:  "注册成功",
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
			model.DB.Where("userid = ?", user.ID).Find(&sessions)
			delta, _ := time.ParseDuration("289h")
			for _, session := range sessions {
				if time.Since(session.CreatedAt) > delta {
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
	TokenRequired
}

func (service UserLogoutService) Logout() serializer.Response {
	if session := service.CheckAuth(); session == nil {
		return serializer.Response{
			Code: 40003,
			Msg:  "未找到会话",
		}
	} else {
		model.DB.Where("UUID = ?", service.UUID).First(&session)
		model.DB.Delete(&session)
		return serializer.Response{
			Code: 20000,
			Msg:  "成功登出",
		}
	}
}
