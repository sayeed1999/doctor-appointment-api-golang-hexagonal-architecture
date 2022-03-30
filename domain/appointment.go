package domain

import (
	"time"

	"gorm.io/gorm"
)

// One thing to keep in mind that a doc can make max 03 appointments each day.
// count the now of rows in appointments table where dateOfAppointment is that date and doctor is that doctor.
// if count == 3, cannot appoint on this day :( the doctor is full under pressure!

type Appointment struct {
	DoctorId          int       `json:"doctorId"`
	DateOfAppointment time.Time `json:"dateOfAppointment"`
	Doctor            Doctor    `json:"doctor" gorm:"foreignkey:DoctorId"`
	gorm.Model
}
