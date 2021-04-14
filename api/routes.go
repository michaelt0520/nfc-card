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
			apiV1.DELETE("/signout", middlewares.AllAuth(), s.authHandler.Signout)

			cardGroup := apiV1.Group("/cards")
			{
				cardGroup.GET("/", middlewares.AdminAuth(), s.cardHandler.Index)
				cardGroup.GET("/:code", s.cardHandler.Show)
				cardGroup.POST("/", middlewares.AdminAuth(), s.cardHandler.Create)
				cardGroup.PUT("/:code", middlewares.AdminAuth(), s.cardHandler.Update)
				cardGroup.DELETE("/:code", middlewares.AdminAuth(), s.cardHandler.Destroy)
			}

			infoGroup := apiV1.Group("/informations")
			infoGroup.Use(middlewares.AllAuth())
			{
				infoGroup.POST("/", s.infoHandler.Create)
				infoGroup.PUT("/:id", s.infoHandler.Update)
				infoGroup.DELETE("/:id", s.infoHandler.Destroy)
			}

			userGroup := apiV1.Group("/users")
			{
				userGroup.GET("/", middlewares.AdminAuth(), s.userHandler.Index)
				userGroup.PUT("/:username", middlewares.MemberAuth(), s.userHandler.Update)
			}

			companyGroup := apiV1.Group("companies")
			companyGroup.Use(middlewares.AdminAuth())
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
