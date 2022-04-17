package repository

import (
	"github.com/sayeed1999/doctor-appointment-api-golang-hexagonal-architecture/internal/core/domain"
	"gorm.io/gorm"
)

var AccountRepo *AccountRepository

type AccountRepository struct {
	*repository[domain.User]
}

func (r *AccountRepository) Initialize(db *gorm.DB) *AccountRepository {
	AccountRepo = &AccountRepository{
		repository: &repository[domain.User]{db: db},
	}
	return AccountRepo
}
