package domain

import (
	"gorm.io/gorm"
)

type User struct {
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Phone    int    `json:"phone"`
	Password string `json:"-"`
	Cost     int    `json:"-"` // Cost is the cost by which we encrypted Password via bcrypt
	gorm.Model
}
