package server

import "github.com/labstack/echo/v4"

func (s *httpServer) setRoutes(e *echo.Echo) {
	e.POST("/user", s.handler.createUser)
}
