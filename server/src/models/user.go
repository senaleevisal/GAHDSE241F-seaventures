package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string `json:"userName"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
