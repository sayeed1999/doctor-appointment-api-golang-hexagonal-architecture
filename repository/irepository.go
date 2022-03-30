package repository

type IRepository interface {

	// Create(model interface{}) (interface{}, int, string)

	// Update(id int, model interface{}) (interface{}, int, string)

	// Delete(id int) (interface{}, int, string)

	FindById(receiver interface{}, id int) error

	FindAll(models interface{}, where string, args ...interface{}) error

	ExecuteRawSqlAndScan(receiver interface{}, query string, args ...interface{}) error
}
