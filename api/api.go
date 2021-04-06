package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/michaelt0520/nfc-card/handlers"
)

// Server : struct
type Server struct {
	g             *gin.Engine
	authHandler   *handlers.AuthHandler
	userHandler   *handlers.UserHandler
	socialHandler *handlers.SocialInformationHandler
}

// NewServer ...
func NewServer(
	g *gin.Engine,
	authHandler *handlers.AuthHandler,
	userHandler *handlers.UserHandler,
	socialHandler *handlers.SocialInformationHandler) *Server {
	return &Server{
		g:             g,
		authHandler:   authHandler,
		userHandler:   userHandler,
		socialHandler: socialHandler,
	}
}

// DefaultWelcome ...
func (s *Server) DefaultWelcome(c *gin.Context) {
	c.JSON(http.StatusOK, "It's working!")
}
