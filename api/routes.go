package api

// InitRoutes ...
func (s *Server) InitRoutes() {
	s.g.GET("/", s.DefaultWelcome)

	apiGroup := s.g.Group("/api")
	{
		apiV1 := apiGroup.Group("/v1")
		{
			userGroup := apiV1.Group("/users")
			{
				userGroup.GET("/:username", s.userHandler.Find)
			}

			socialGroup := apiV1.Group("/social")
			{
				socialGroup.POST("/", s.socialHandler.Create)
				socialGroup.DELETE("/:id", s.socialHandler.Destroy)
			}
		}
	}
}
