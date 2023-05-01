package model

import "gorm.io/gorm"

const (
	Admin = iota
	Teacher
	Student
)

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password" gorm:"size=32"`
	RealName string `json:"real_name" gorm:"not null"`
	SchoolID uint
	School   School `gorm:"foreignKey:SchoolID"`
	Type     uint   `json:"type"`
}
