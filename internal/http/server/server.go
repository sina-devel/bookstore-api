package server

import (
	"fmt"
	"github.com/kianooshaz/bookstore-api/internal/contract"
	"github.com/labstack/echo/v4"
)

type (
	httpServer struct {
		handler *handler
	}
)

func NewHttpServer(h *handler) contract.HttpServer {
	return &httpServer{
		handler: h,
	}
}

func (s *httpServer) Start(port uint) error {
	e := echo.New()

	s.setRoutes(e)

	if port == 0 {
		port = 8083
	}

	return e.Start(fmt.Sprintf(":%d", port))
}
