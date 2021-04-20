package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/michaelt0520/nfc-card/handlers"
)

// Server : struct
type Server struct {
	g                   *gin.Engine
	uploadHandler       *handlers.UploadHandler
	authHandler         *handlers.AuthHandler
	userHandler         *handlers.UserHandler
	infoHandler         *handlers.InformationHandler
	compUserHandler     *handlers.CompanyUserHandler
	compCardHandler     *handlers.CompanyCardHandler
	compHandler         *handlers.CompanyHandler
	adminUserHandler    *handlers.AdminUserHandler
	adminInfoHandler    *handlers.AdminInformationHandler
	adminCardHandler    *handlers.AdminCardHandler
	adminCompanyHandler *handlers.AdminCompanyHandler
}

// NewServer ...
func NewServer(
	g *gin.Engine,
	uploadHandler *handlers.UploadHandler,
	authHandler *handlers.AuthHandler,
	userHandler *handlers.UserHandler,
	infoHandler *handlers.InformationHandler,
	compUserHandler *handlers.CompanyUserHandler,
	compCardHandler *handlers.CompanyCardHandler,
	compHandler *handlers.CompanyHandler,
	adminUserHandler *handlers.AdminUserHandler,
	adminInfoHandler *handlers.AdminInformationHandler,
	adminCardHandler *handlers.AdminCardHandler,
	adminCompanyHandler *handlers.AdminCompanyHandler) *Server {

	return &Server{
		g:                   g,
		uploadHandler:       uploadHandler,
		authHandler:         authHandler,
		userHandler:         userHandler,
		infoHandler:         infoHandler,
		compUserHandler:     compUserHandler,
		compCardHandler:     compCardHandler,
		compHandler:         compHandler,
		adminUserHandler:    adminUserHandler,
		adminInfoHandler:    adminInfoHandler,
		adminCardHandler:    adminCardHandler,
		adminCompanyHandler: adminCompanyHandler,
	}
}

// DefaultWelcome ...
func (s *Server) DefaultWelcome(c *gin.Context) {
	c.JSON(http.StatusOK, "It's working!")
}
