package repository

import (
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

// defining struct
type repository[T any] struct {
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
func InitializeDB(connStr string) (*gorm.DB, error) {

	db, err := newSqlServerClient(connStr)

	if err != nil {
		return nil, err
	}

	return db, nil
}

func (r *repository[T]) Create(model T) error {
	// We use type switches in golang to identify which entity it is, because gorm is failing ..
	// T := reflect.TypeOf(model)
	// switch T {
	// case reflect.TypeOf(domain.User{}):
	// 	user := model.(domain.User)
	// 	return r.db.Create(&user).Error
	// case reflect.TypeOf(domain.Doctor{}):
	// 	doctor := model.(domain.Doctor)
	// 	return r.db.Create(&doctor).Error
	// }
	// return nil
	return r.db.Create(&model).Error
}

func (r *repository[T]) FindById(receiver *T, id int) error {
	return r.db.First(receiver, id).Error
}

func (r *repository[T]) FindFirst(receiver *T, where string, args ...interface{}) error {
	return r.db.Where(where, args...).Limit(1).Find(receiver).Error
}

func (r *repository[T]) FindAll(models *[]T, where string, args ...interface{}) error {
	return r.db.Where(where, args...).Find(models).Error
}

func (r *repository[T]) ExecuteRawSqlAndScan(receiver interface{}, query string, args ...interface{}) error {
	tx := r.db.Raw(query, args...).Scan(receiver)
	return tx.Error
}
