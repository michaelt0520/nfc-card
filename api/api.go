package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/michaelt0520/nfc-card/handlers"
)

// Server : struct
type Server struct {
	g           *gin.Engine
	authHandler *handlers.AuthHandler
	userHandler *handlers.UserHandler
	infoHandler *handlers.InformationHandler
	cardHandler *handlers.CardHandler
}

// NewServer ...
func NewServer(
	g *gin.Engine,
	authHandler *handlers.AuthHandler,
	userHandler *handlers.UserHandler,
	infoHandler *handlers.InformationHandler,
	cardHandler *handlers.CardHandler) *Server {
	return &Server{
		g:           g,
		authHandler: authHandler,
		userHandler: userHandler,
		infoHandler: infoHandler,
		cardHandler: cardHandler,
	}
}

// DefaultWelcome ...
func (s *Server) DefaultWelcome(c *gin.Context) {
	c.JSON(http.StatusOK, "It's working!")
}
