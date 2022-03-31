package repository

type IRepository interface {
	Create(model interface{}) error

	// Update(id int, model interface{}) error

	// Delete(id int) error

	FindById(receiver interface{}, id int) error

	FindFirst(receiver interface{}, where string, args ...interface{}) error

	FindAll(models interface{}, where string, args ...interface{}) error

	ExecuteRawSqlAndScan(receiver interface{}, query string, args ...interface{}) error
}
