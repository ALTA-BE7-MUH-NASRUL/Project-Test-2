package entities

import (
	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	UserID uint
	Title  string `json:"title" form:"title"`
	Status string `json:"status" form:"status"`
	Task   []Task `gorm:"foreignKey:ProjectID;references:ID"`
}
