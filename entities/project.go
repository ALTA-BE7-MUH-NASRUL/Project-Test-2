package entities

import (
	"time"

	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	UserID   uint
	Title    string    `json:"title" form:"title"`
	Deadline time.Time `json:"deadline" form:"deadline"`
	Status   string    `json:"status" form:"status"`
	Task     []Task    `gorm:"foreignKey:ProjectID;references:ID"`
}
