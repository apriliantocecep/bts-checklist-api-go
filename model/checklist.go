package model

import "gorm.io/gorm"

type Checklist struct {
	gorm.Model
	Name   string `gorm:"not null" json:"name"`
	UserId uint   `json:"user_id"`
	User   User
}
