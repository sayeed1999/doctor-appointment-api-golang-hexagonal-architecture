package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sayeed1999/doctor-appointment-api-golang-hexagonal-architecture/config"
	"github.com/sayeed1999/doctor-appointment-api-golang-hexagonal-architecture/internal/core/domain"
	"github.com/sayeed1999/doctor-appointment-api-golang-hexagonal-architecture/internal/core/service"
	"github.com/sayeed1999/doctor-appointment-api-golang-hexagonal-architecture/internal/handlers"
	"github.com/sayeed1999/doctor-appointment-api-golang-hexagonal-architecture/internal/mailing"
	"github.com/sayeed1999/doctor-appointment-api-golang-hexagonal-architecture/internal/middleware"
	"github.com/sayeed1999/doctor-appointment-api-golang-hexagonal-architecture/internal/redisconfig"
	"github.com/sayeed1999/doctor-appointment-api-golang-hexagonal-architecture/internal/repository"
	"github.com/sayeed1999/doctor-appointment-api-golang-hexagonal-architecture/internal/route"
)

func main() {

	if err := run(); err != nil {
		log.Fatal(err)
	}

	// Creates a gin router with default middleware: logger and recovery (crash-free) middleware
	router := gin.Default()

	// Initialize middlewares for this app
	middleware.InitializeMiddlewares(router)

	// Initialize routes for this app
	route.InitializeRoutes(router)

	// By default it serves on :8080 unless a PORT environment variable was defined.
	log.Fatal(router.Run()) // router.Run(":3000") for a hard coded port
}

func run() error {
	conf, err := config.NewConfig("./config/config.yaml")
	if err != nil {
		// log.Fatal(err)
		return err
	}

	// mail config ..
	mailChan := make(chan domain.MailData, 2) // i make a buffered channel of capacity: 2 for 2 mails at once, doctor & patient
	mailing.ListenForMail(mailChan, conf.Mail.SmtpHost, conf.Mail.SmtpPort, conf.Mail.SenderEmail, conf.Mail.Password)

	// redis cache config ..
	redisClient := redisconfig.InitializeRedisServer(conf.Redis.Address, conf.Redis.Password, conf.Redis.DB)
	context := context.Background()

	// Initializing database ..
	db, err := repository.InitializeDB(conf.Database.ConnectionString)
	if err != nil {
		return err
	}

	// Initializing repos ..
	docRepo := repository.DoctorRepo.Initialize(db)
	accRepo := repository.AccountRepo.Initialize(db)

	// Initializing services ...
	baseSrv := service.InitializeBaseService(redisClient, &context, mailChan, conf)
	docSrv := service.InitializeDoctorService(baseSrv, docRepo)
	accSrv := service.InitializeAccountService(baseSrv, accRepo)

	// Initializing handlers ...
	handlers.Base.Initialize(conf)
	handlers.DoctorHandler.Initialize(docSrv)
	handlers.AccountHandler.Initialize(accSrv)

	return nil
}
