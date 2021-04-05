package api

import (
	"net/http"

	"github.com/michaelt0520/nfc-card/handlers"
	"github.com/gin-gonic/gin"
)

// Server : struct
type Server struct {
	g       			*gin.Engine
	userHandler   *handlers.UserHandler
	socialHandler *handlers.SocialInformationHandler
}

// NewServer ...
func NewServer(
	g 						*gin.Engine,
	userHandler   *handlers.UserHandler
	socialHandler *handlers.SocialInformationHandler) *Server {
	return &Server{
		g:       			 g,
		userHandler:   userHandler,
		socialHandler: socialHandler,
	}
}

// DefaultWelcome ...
func (s *Server) DefaultWelcome(c *gin.Context) {
	c.JSON(http.StatusOK, "It's working!")
}
