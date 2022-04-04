package entities

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	UserID    uint   `json:"UserID"`
	ProjectID uint   `json:"ProjcetID"`
	List      string `json:"list" form:"list"`
	Status    string `json:"status" form:"status"`
}
