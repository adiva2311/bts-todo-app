package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	Name        string `json:"name"`
	Status      bool   `json:"status"`
	ChecklistID uint   `json:"checklist_id"`
}

func (Item) TableName() string {
	return "items"
}
