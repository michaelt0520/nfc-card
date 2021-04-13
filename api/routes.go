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
				cardGroup.GET("/:code", s.cardHandler.Show)
				cardGroup.POST("/", s.cardHandler.Create, middlewares.Auth())
				cardGroup.PUT("/:code", s.cardHandler.Update, middlewares.Auth())
				cardGroup.DELETE("/:code", s.cardHandler.Create, middlewares.Auth())
			}

			infoGroup := apiV1.Group("/informations")
			infoGroup.Use(middlewares.Auth())
			{
				infoGroup.POST("/", s.infoHandler.Create)
				infoGroup.PUT("/:id", s.infoHandler.Update)
				infoGroup.DELETE("/:id", s.infoHandler.Destroy)
			}

			userGroup := apiV1.Group("/users")
			{
				userGroup.PUT("/:username", s.userHandler.Update)
			}

			companyGroup := apiV1.Group("companies")
			companyGroup.Use(middlewares.Auth())
			{
				companyGroup.GET("/", s.compHandler.Index)
				companyGroup.GET("/:id", s.compHandler.Show)
				companyGroup.POST("/", s.compHandler.Create)
				companyGroup.PUT("/:id", s.compHandler.Update)
				companyGroup.DELETE("/:id", s.compHandler.Create)
			}
		}
	}
}
