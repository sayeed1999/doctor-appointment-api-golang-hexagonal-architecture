package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sayeed1999/doctor-appointment-api-golang-hexagonal-architecture/domain/vm"
	"github.com/sayeed1999/doctor-appointment-api-golang-hexagonal-architecture/helpers"
	"github.com/sayeed1999/doctor-appointment-api-golang-hexagonal-architecture/pkg"
	"github.com/sayeed1999/doctor-appointment-api-golang-hexagonal-architecture/service"
)

var DoctorHandler *doctorHandler

type doctorHandler struct {
	*base
}

func (h *doctorHandler) Initialize() {
	DoctorHandler = &doctorHandler{
		base: Base,
	}
}

// Gets all doctors

func (h *doctorHandler) GetDoctors(c *gin.Context) {
	doctors, code, text := service.DoctorService.GetAllDoctors()
	helpers.Response(c, code, text, doctors)
}

// Gets the doctor by it's primary key

func (h *doctorHandler) GetDoctorById(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		helpers.Response(c, 400, "could not parse primary key ID of doctor from route parameter", nil)
		return
	}
	doctor, code, text := service.DoctorService.GetDoctorById(int(id))
	helpers.Response(c, code, text, doctor)
}

// Get doctors available within a specific date range

func (h *doctorHandler) GetDoctorsByAvailability(c *gin.Context) {
	params := c.Params
	fromDateInString := params.ByName("fromDate")
	toDateInString := params.ByName("toDate")

	fromDate, err := pkg.GetUtcDateFromYearMonthDayFormattedString(fromDateInString)
	toDate, err := pkg.GetUtcDateFromYearMonthDayFormattedString(toDateInString)

	if err != nil {
		helpers.Response(c, 404, "invalid date format cannot be parsed", nil)
		return
	}

	availabilities, code, text := service.DoctorService.GetDoctorsByAvailability(fromDate, toDate)
	helpers.Response(c, code, text, availabilities)
}

// This POST method requests an appointment to a doctor

func (h *doctorHandler) RequestAppointmentToDoctor(c *gin.Context) {
	headerContentTtype := c.Request.Header.Get("Content-Type")

	if headerContentTtype != "application/json" {
		helpers.Response(c, http.StatusUnsupportedMediaType, "Content Type is not application/json", nil)
		return
	}

	var apnt vm.AppointmentVM
	var unmarshalErr *json.UnmarshalTypeError

	// decoder := json.NewDecoder(c.Request.Body)
	// decoder.DisallowUnknownFields()
	// err := decoder.Decode(&apnt)

	if err := pkg.DecodeJSON(c.Request.Body, &apnt); err != nil {
		fmt.Println(err.Error()) // badly needed to understand where error occurred while parsing the form data
		if errors.As(err, &unmarshalErr) {
			helpers.Response(c, 400, "Bad Request. Wrong Type provided for field "+unmarshalErr.Field, nil)
		} else {
			helpers.Response(c, 400, "Bad Request "+err.Error(), nil)
		}
		return
	}

	code, text := service.DoctorService.RequestAppointmentToDoctor(apnt)
	// json.NewEncoder(w).Encode(message)
	helpers.Response(c, code, text, nil)
}
