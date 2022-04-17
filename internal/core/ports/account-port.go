package ports

import "github.com/sayeed1999/doctor-appointment-api-golang-hexagonal-architecture/internal/core/domain"

type AccountService interface {
	Register(user domain.User) (domain.User, int, string)
	Login(email string, password string) (string, int, string)
	GetUserById(id int) (domain.User, int, string)
}

type AccountRepository interface {
	BaseRepository[domain.User]
}
