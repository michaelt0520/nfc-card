package api

import "github.com/michaelt0520/nfc-card/middlewares"

// InitRoutes ...
func (s *Server) InitRoutes() {
	s.g.GET("/", s.DefaultWelcome)

	s.g.Static("/public", "./static")

	apiGroup := s.g.Group("/api")
	{
		apiV1 := apiGroup.Group("/v1")
		{
			apiV1.POST("/signin", s.authHandler.Signin)
			apiV1.POST("/signup", s.authHandler.Signup)
			apiV1.DELETE("/signout", s.authHandler.Signout, middlewares.Auth())

			cardGroup := apiV1.Group("/cards")
			{
				cardGroup.GET("/:code", s.cardHandler.Find)
			}

			socialGroup := apiV1.Group("/information")
			socialGroup.Use(middlewares.Auth())
			{
				socialGroup.POST("/", s.infoHandler.Create)
				socialGroup.DELETE("/:id", s.infoHandler.Destroy)
			}
		}
	}
}
