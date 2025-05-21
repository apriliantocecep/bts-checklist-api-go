package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"not null" json:"email"`
	Username string `gorm:"not null" json:"username"`
	Password string `gorm:"not null" json:"password,omitempty"`
}
