package repository

import (
	"reflect"

	"github.com/sayeed1999/doctor-appointment-api-golang-hexagonal-architecture/domain"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var Repository *repository

// defining struct
type repository struct {
	db *gorm.DB
}

// new sqlserver client
func newSqlServerClient(connStr string) (*gorm.DB, error) {

	db, err := gorm.Open(sqlserver.Open(connStr), &gorm.Config{})

	if err == nil {
		KeepAutoMigrationUpAndRunning(db)
	}
	return db, err
}

// This repository uses sql server
func (r *repository) InitializeRepository(connStr string) (*repository, error) {

	db, err := newSqlServerClient(connStr)

	if err != nil {
		return nil, err
	}

	Repository = &repository{
		db: db,
	}
	return Repository, nil
}

func (r *repository) Create(model interface{}) error {
	// We use type switches in golang to identify which entity it is, because gorm is failing ..
	T := reflect.TypeOf(model)
	switch T {
	case reflect.TypeOf(domain.User{}):
		user := model.(domain.User)
		return Repository.db.Create(&user).Error
	case reflect.TypeOf(domain.Doctor{}):
		doctor := model.(domain.Doctor)
		return Repository.db.Create(&doctor).Error
	}
	return nil
}

func (r *repository) FindById(receiver interface{}, id int) error {
	return Repository.db.First(receiver, id).Error
}

func (r *repository) FindAll(models interface{}, where string, args ...interface{}) error {
	return Repository.db.Where(where, args).Find(models).Error
}

func (r *repository) ExecuteRawSqlAndScan(receiver interface{}, query string, args ...interface{}) error {
	tx := Repository.db.Raw(query, args...).Scan(receiver)
	return tx.Error
}
