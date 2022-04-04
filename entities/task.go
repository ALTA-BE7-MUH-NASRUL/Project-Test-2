package entities

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	UserID    uint      `json:"UserID"`
	ProjectID uint      `json:"ProjcetID"`
	List      string    `json:"list" form:"list"`
	Deadline  time.Time `json:"deadline" form:"deadline"`
	Status    string    `json:"status" form:"status"`
}
