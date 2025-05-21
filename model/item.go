package model

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	ItemName    string `gorm:"not null" json:"itemName"`
	Status      int    `json:"status"`
	ChecklistId uint   `json:"checklist_id"`
	Checklist   Checklist
}
