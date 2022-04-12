package dao

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	RecordID uint   `json:"record_id" binding:"required"`
	Content  string `json:"content" binding:"required"`
}
