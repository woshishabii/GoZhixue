package model

import "gorm.io/gorm"

type School struct {
	gorm.Model
	Name string `json:"name" gorm:"name"`
}

type Class struct {
	gorm.Model
	Name     string `json:"name"`
	SchoolID uint
	School   School `gorm:"foreighKey:SchoolID"`
	ExamID   uint
}
