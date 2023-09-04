package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name  string `gorm:"type:varchar(300)" json:"name"`
	Posts []Post `json:"posts"`
}
