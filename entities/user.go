package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string    `json:"name" form:"name"`
	Email    string    `json:"email" form:"email"`
	Password string    `json:"password" form:"password"`
	Project  []Project `gorm:"foreignKey:UserID;references:ID"`
	Task     []Task    `gorm:"foreignKey:UserID;references:ID"`
}
