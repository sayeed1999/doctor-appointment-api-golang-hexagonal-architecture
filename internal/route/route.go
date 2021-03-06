package route

import (
	"github.com/gin-gonic/gin"
	"github.com/sayeed1999/doctor-appointment-api-golang-hexagonal-architecture/internal/handlers"
)

func InitializeRoutes(router *gin.Engine) {

	// defining routes for doctors
	router.GET("/doctors", handlers.DoctorHandler.GetDoctors)
	router.GET("/doctors/:id", handlers.DoctorHandler.GetDoctorById)
	router.POST("/doctors/:id/make-appointment", handlers.DoctorHandler.RequestAppointmentToDoctor)
	router.GET("/doctors/fromDate/:fromDate/toDate/:toDate", handlers.DoctorHandler.GetDoctorsByAvailability) // format should be: /doctors/fromDate/2022-03-01/toDate/2022-04-30

	// defining routes for accounts
	router.POST("/account/register", handlers.AccountHandler.Register)
	router.POST("/account/login", handlers.AccountHandler.Login)
	router.GET("/account/user", handlers.AccountHandler.GetAuthenticatedUser)
}
