package domain

import (
	"gorm.io/gorm"
)

type MailData struct {
	To      string
	From    string
	Subject string
	Content string //template.HTML
	gorm.Model
}
