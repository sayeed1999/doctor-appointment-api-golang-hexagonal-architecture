package repository

import (
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
		defer KeepAutoMigrationUpAndRunning(db)
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
