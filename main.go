package main

import (
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/michaelt0520/nfc-card/api"
	"github.com/michaelt0520/nfc-card/config"
	"github.com/michaelt0520/nfc-card/db"
	"github.com/michaelt0520/nfc-card/db/seeds"
	"github.com/michaelt0520/nfc-card/handlers"
	"github.com/michaelt0520/nfc-card/logger"
	"github.com/michaelt0520/nfc-card/models"
	"github.com/michaelt0520/nfc-card/repositories"
	"github.com/michaelt0520/nfc-card/services"
)

var log *zap.Logger

func main() {
	// get config
	conf := config.GetConfig()

	// init log
	log = logger.InitLogger(conf.Env)

  // connect db mysql
	if err := db.ConnectDatabase(conf); err != nil {
		log.Error("failed to connect db", zap.Error(err))
		return
	}

  // connect mongo
  if err := db.ConnectMongoClient(); err != nil {
		log.Error("failed to connect db", zap.Error(err))
		return
	}

	// config CORS
	r := gin.New()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://*", "https://*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc:  func(origin string) bool { return true },
		AllowMethods:     []string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		MaxAge:           12 * time.Hour,
	}))

	// Migration
	db.Database().AutoMigrate(
		&models.User{},
		&models.Information{},
		&models.Card{},
		&models.Company{},
		&models.Contact{},
	)

	// Seed
	seeds.Load(log)

	// init Repository
	userRepo := repositories.NewUserRepository()
	infoRepo := repositories.NewInformationRepository()
	cardRepo := repositories.NewCardRepository()
	compRepo := repositories.NewCompanyRepository()
	contactRepo := repositories.NewContactRepository()

	// init services
	userSrv := services.NewUserService(userRepo)
	infoSrv := services.NewInformationService(infoRepo)
	compSrv := services.NewCompanyService(compRepo, userRepo)
	cardSrv := services.NewCardService(cardRepo)
	contactSrv := services.NewContactService(contactRepo)

	// init Handler
	appHandler := handlers.NewAppHandler(cardSrv, contactSrv, userSrv)
	authHandler := handlers.NewAuthHandler(userSrv)
	uploadHandler := handlers.NewUploadHandler(userSrv, compSrv)
	userHandler := handlers.NewUserHandler(userSrv)
	infoHandler := handlers.NewInformationHandler(userSrv, infoSrv)
	cardHandler := handlers.NewCardHandler(userSrv, cardSrv)
	compUserHandler := handlers.NewCompanyUserHandler(compSrv, userSrv)
	compCardHandler := handlers.NewCompanyCardHandler(compSrv, cardSrv)
	compHandler := handlers.NewCompanyHandler(compSrv)
	adminUserHandler := handlers.NewAdminUserHandler(userSrv)
	adminInfoHandler := handlers.NewAdminInformationHandler(userSrv, infoSrv)
	adminCardHandler := handlers.NewAdminCardHandler(cardSrv)
	adminCompHandler := handlers.NewAdminCompanyHandler(compSrv)

	// init Server
	svr := api.NewServer(
		r,
		appHandler,
		uploadHandler,
		authHandler,
		userHandler,
		infoHandler,
		cardHandler,
		compUserHandler,
		compCardHandler,
		compHandler,
		adminUserHandler,
		adminInfoHandler,
		adminCardHandler,
		adminCompHandler,
	)
	svr.InitRoutes()

	// run Server
	if err := r.Run(fmt.Sprintf(":%d", conf.Port)); err != nil {
		log.Fatal("router.Run", zap.Error(err))
	}
}
