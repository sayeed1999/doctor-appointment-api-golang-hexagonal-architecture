package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sayeed1999/doctor-appointment-api-golang-hexagonal-architecture/config"
	"github.com/sayeed1999/doctor-appointment-api-golang-hexagonal-architecture/domain"
	"github.com/sayeed1999/doctor-appointment-api-golang-hexagonal-architecture/handlers"
	"github.com/sayeed1999/doctor-appointment-api-golang-hexagonal-architecture/mailing"
	"github.com/sayeed1999/doctor-appointment-api-golang-hexagonal-architecture/middleware"
	"github.com/sayeed1999/doctor-appointment-api-golang-hexagonal-architecture/redisconfig"
	"github.com/sayeed1999/doctor-appointment-api-golang-hexagonal-architecture/repository"
	"github.com/sayeed1999/doctor-appointment-api-golang-hexagonal-architecture/route"
	"github.com/sayeed1999/doctor-appointment-api-golang-hexagonal-architecture/service"
)

func main() {
	conf, err := config.NewConfig("./config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	mailChan := make(chan domain.MailData, 2) // i make a buffered channel of capacity: 2 for 2 mails at once, doctor & patient
	mailing.ListenForMail(mailChan, conf.Mail.SmtpHost, conf.Mail.SmtpPort, conf.Mail.SenderEmail, conf.Mail.Password)

	repo, _ := repository.Repository.InitializeRepository(conf.Database.ConnectionString)
	redisClient := redisconfig.InitializeRedisServer(conf.Redis.Address, conf.Redis.Password, conf.Redis.DB)
	context := context.Background()

	// Initializing services ...
	service.Base.Initialize(repo, redisClient, &context, mailChan, conf)
	service.DoctorService.Initialize()
	service.AccountService.Initialize()

	// Initializing handlers ...
	handlers.Base.Initialize()
	handlers.DoctorHandler.Initialize()
	handlers.AccountHandler.Initialize()

	// Creates a gin router with default middleware: logger and recovery (crash-free) middleware
	router := gin.Default()

	// Initialize middlewares for this app
	middleware.InitializeMiddlewares(router)

	// Initialize routes for this app
	route.InitializeRoutes(router)

	// By default it serves on :8080 unless a PORT environment variable was defined.
	log.Fatal(router.Run()) // router.Run(":3000") for a hard coded port
}
