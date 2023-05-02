package service

import (
	"GoZhixue/model"
)

type TokenRequired struct {
	UUID string `json:"uuid" binding:"required,len=36"`
}

func (token TokenRequired) CheckAuth() *model.Session {
	var session model.Session
	if err := model.DB.Where("UUID = ?").First(&session); err != nil {
		return nil
	} else {
		return &session
	}
}
