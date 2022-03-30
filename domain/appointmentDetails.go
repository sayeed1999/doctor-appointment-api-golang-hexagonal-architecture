package domain

import (
	"gorm.io/gorm"
)

type AppointmentDetails struct {
	AppointmentId int         `json:"appointmentId"`
	PatientName   string      `json:"patientName"`
	PatientEmail  string      `json:"patientEmail"`
	PatientPhone  string      `json:"patientPhone"`
	Appointment   Appointment `json:"appointment" gorm:"foreignkey:AppointmentId"`
	gorm.Model
}
