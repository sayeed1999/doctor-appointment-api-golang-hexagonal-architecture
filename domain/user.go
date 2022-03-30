package domain

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Phone      string    `json:"phone"`
	DOB        time.Time `json:"dOB"`
	Gender     string    `json:"gender"`
	Password   string    `json:"-"`
	gorm.Model           // Base
}
