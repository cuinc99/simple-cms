package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title      string `gorm:"type:varchar(300)" json:"title"`
	Content    string `gorm:"type:text" json:"content"`
	CategoryId uint   `json:"category_id"`
}
