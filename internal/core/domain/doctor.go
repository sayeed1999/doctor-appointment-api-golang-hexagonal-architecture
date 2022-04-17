package domain

import "gorm.io/gorm"

type Doctor struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	gorm.Model
}
