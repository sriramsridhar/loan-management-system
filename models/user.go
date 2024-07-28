package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique" json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"` // customer, agent, admin
	Approved bool   `json:"approved"`
}
