package server

func (s *httpServer) setRoutes() {
	s.admin.POST("/user", s.handler.adminCreateUser)
	s.admin.GET("/user/:id", s.handler.adminGetUser)
	s.admin.PATCH("/user", s.handler.adminUpdateUser)
	s.admin.DELETE("/user/:id", s.handler.adminDeleteUser)

	s.user.GET("/token/refresh/:id", s.handler.refreshToken)

	s.public.POST("/login", s.handler.login)
}
