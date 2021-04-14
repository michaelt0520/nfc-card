package main

import (
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/michaelt0520/nfc-card/api"
	"github.com/michaelt0520/nfc-card/config"
	"github.com/michaelt0520/nfc-card/handlers"
	"github.com/michaelt0520/nfc-card/logger"
	"github.com/michaelt0520/nfc-card/models"
	"github.com/michaelt0520/nfc-card/repositories"
	"github.com/michaelt0520/nfc-card/seeds"
)

var log *zap.Logger

func main() {
	// get config
	conf := config.GetConfig()

	// init log
	log = logger.InitLogger(conf.Env)

	// init repositories
	if err := repositories.InitRepository(conf); err != nil {
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
	repositories.GetDB().AutoMigrate(
		&models.User{},
		&models.Information{},
		&models.Card{},
		&models.Company{},
	)

	// Seed
	seeds.Load(log)

	// init Repository
	userRepo := repositories.NewUserRepository()
	infoRepo := repositories.NewInformationRepository()
	cardRepo := repositories.NewCardRepository()
	compRepo := repositories.NewCompanyRepository()

	// init Handler
	authHandler := handlers.NewAuthHandler(userRepo)
	userHandler := handlers.NewUserHandler(userRepo)
	socialHandler := handlers.NewInformationHandler(infoRepo)
	cardHandler := handlers.NewCardHandler(cardRepo)
	compHanlder := handlers.NewCompanyHandler(compRepo)

	// init Server
	svr := api.NewServer(r, authHandler, userHandler, socialHandler, cardHandler, compHanlder)
	svr.InitRoutes()

	// run Server
	if err := r.Run(fmt.Sprintf(":%d", conf.Port)); err != nil {
		log.Fatal("router.Run", zap.Error(err))
	}
}
