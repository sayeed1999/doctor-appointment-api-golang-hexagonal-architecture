package repository

import (
	"log"

	"github.com/sayeed1999/doctor-appointment-api-golang-hexagonal-architecture/domain"
	"github.com/sayeed1999/doctor-appointment-api-golang-hexagonal-architecture/domain/vm"
	"gorm.io/gorm"
)

func KeepAutoMigrationUpAndRunning(db *gorm.DB) {
	// the entities from vm package are only used to represent stored proc result sets, these tables are empty
	err := db.AutoMigrate(
		&domain.Doctor{},
		&domain.Appointment{},
		&domain.AppointmentDetails{},
		&vm.SearchAvailability{},
		&domain.User{},
	)
	if err != nil {
		log.Fatal(err)
	}
}
