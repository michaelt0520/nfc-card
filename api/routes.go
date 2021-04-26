package api

import "github.com/michaelt0520/nfc-card/middlewares"

// InitRoutes ...
func (s *Server) InitRoutes() {
	s.g.GET("/", s.DefaultWelcome)

	s.g.Static("/public", "./public")

	apiGroup := s.g.Group("/api")
	{
		apiV1 := apiGroup.Group("/v1")
		{
			apiV1.POST("/signin", s.authHandler.Signin)
			apiV1.POST("/signup", s.authHandler.Signup)
			apiV1.DELETE("/signout", middlewares.AllAuth(), s.authHandler.Signout)

			appGroup := apiV1.Group("/app")
			{
				appGroup.GET("/card/:code", s.appHandler.ShowCard)
				appGroup.POST("/contact", s.appHandler.CreateContact)
			}

			uploadGroup := apiV1.Group("/upload")
			{
				uploadGroup.POST("/avatar", middlewares.AllAuth(), s.uploadHandler.Avatar)
				uploadGroup.POST("/logo", middlewares.CompanyAuth(), s.uploadHandler.Logo)
			}

			userGroup := apiV1.Group("/user")
			userGroup.Use(middlewares.AllAuth())
			{
				userGroup.GET("/", s.userHandler.Show)
				userGroup.PUT("/", s.userHandler.Update)

				infoGroup := userGroup.Group("/informations")
				{
					infoGroup.POST("/", s.infoHandler.Create)
					infoGroup.PUT("/:id", s.infoHandler.Update)
					infoGroup.DELETE("/:id", s.infoHandler.Destroy)
				}
			}

			companyGroup := apiV1.Group(("/companies/:id"))
			companyGroup.Use(middlewares.CompanyAuth())
			{
				companyGroup.GET("/users", s.compUserHandler.Index)
				companyGroup.POST("/users", s.compUserHandler.Create)
				companyGroup.DELETE("/users/:username", s.compUserHandler.Destroy)
				companyGroup.GET("/cards", s.compCardHandler.Index)
				companyGroup.PUT("/cards/:code", s.compCardHandler.Update)
				companyGroup.PUT("/", s.compHandler.Update)
			}

			adminGroup := apiV1.Group("/admin")
			adminGroup.Use(middlewares.AdminAuth())
			{
				userGroup := adminGroup.Group("/users")
				{
					userGroup.GET("/", s.adminUserHandler.Index)
					userGroup.GET("/:username", s.adminUserHandler.Show)
					userGroup.POST("/", s.adminUserHandler.Create)
					userGroup.PUT("/:username", s.adminUserHandler.Update)
					userGroup.DELETE("/:username", s.adminUserHandler.Destroy)

					infoGroup := userGroup.Group("/:username/informations")
					{
						infoGroup.GET("/", s.adminInfoHandler.Index)
						infoGroup.POST("/", s.adminInfoHandler.Create)
						infoGroup.PUT("/:id", s.adminInfoHandler.Update)
						infoGroup.DELETE("/:id", s.adminInfoHandler.Destroy)
					}
				}

				cardGroup := adminGroup.Group("/cards")
				{
					cardGroup.GET("/", s.adminCardHandler.Index)
					cardGroup.GET("/:code", s.adminCardHandler.Show)
					cardGroup.POST("/", s.adminCardHandler.Create)
					cardGroup.PUT("/:code", s.adminCardHandler.Update)
					cardGroup.DELETE("/:code", s.adminCardHandler.Destroy)
				}

				companyGroup := adminGroup.Group("/companies")
				{
					companyGroup.GET("/", s.adminCompanyHandler.Index)
					companyGroup.GET("/:id", s.adminCompanyHandler.Show)
					companyGroup.POST("/", s.adminCompanyHandler.Create)
					companyGroup.PUT("/:id", s.adminCompanyHandler.Update)
					companyGroup.DELETE("/:id", s.adminCompanyHandler.Destroy)
				}
			}
		}
	}
}
