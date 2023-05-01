package model

import "gorm.io/gorm"

type School struct {
	gorm.Model
	Name string `json:"name" gorm:"name"`
}
