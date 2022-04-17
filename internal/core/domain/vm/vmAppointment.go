package vm

import "time"

type AppointmentVM struct {
	DoctorId          int       `json:"doctorId"`
	DateOfAppointment time.Time `json:"dateOfAppointment"`
	PatientName       string    `json:"patientName"`
	PatientEmail      string    `json:"patientEmail"`
	PatientPhone      string    `json:"patientPhone"`
}
