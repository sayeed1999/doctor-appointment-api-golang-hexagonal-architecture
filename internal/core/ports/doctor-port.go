package ports

import (
	"time"

	"github.com/sayeed1999/doctor-appointment-api-golang-hexagonal-architecture/internal/core/domain"
	"github.com/sayeed1999/doctor-appointment-api-golang-hexagonal-architecture/internal/core/domain/vm"
)

type DoctorService interface {
	GetAllDoctors() ([]domain.Doctor, int, string)
	GetDoctorById(id int) (domain.Doctor, int, string)
	GetDoctorsByAvailability(fromDate time.Time, toDate time.Time) ([]vm.SearchAvailability, int, string)
	RequestAppointmentToDoctor(apnt vm.AppointmentVM) (int, string)
}

type DoctorRepository interface {
}
