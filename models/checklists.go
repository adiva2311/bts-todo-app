package models

import "gorm.io/gorm"

type Checklist struct {
	gorm.Model
	Title  string `json:"title"`
	UserID uint   `json:"user_id"`
	Items  []Item `json:"items"`
}

func (Checklist) TableName() string {
	return "checklist"
}
