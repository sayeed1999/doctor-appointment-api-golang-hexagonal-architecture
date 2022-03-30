package domain

import (
	"gorm.io/gorm"
)

type User struct {
	Fullname   string `json:"fullname"`
	Email      string `json:"email"`
	Phone      int    `json:"phone"`
	Password   string `json:"-"`
	gorm.Model `json:"model"`
}
