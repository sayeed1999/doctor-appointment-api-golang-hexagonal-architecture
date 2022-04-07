package repository

type IRepository[T any] interface {
	Create(model T) error

	// Update(id int, model interface{}) error

	// Delete(id int) error

	FindById(receiver *T, id int) error

	FindFirst(receiver *T, where string, args ...interface{}) error

	FindAll(models *[]T, where string, args ...interface{}) error

	ExecuteRawSqlAndScan(receiver interface{}, query string, args ...interface{}) error
}
