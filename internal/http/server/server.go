package server

import (
	"fmt"
	"github.com/kianooshaz/bookstore-api/internal/contract"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

type (
	httpServer struct {
		handler *handler
		public  *echo.Group
		admin   *echo.Group
		seller  *echo.Group
		user    *echo.Group
	}
)

var (
	e = echo.New()
)

func NewHttpServer(h *handler) contract.HttpServer {

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPatch, http.MethodPut, http.MethodDelete},
	}))

	e.Use(middleware.Gzip())
	e.Use(middleware.RequestID())

	public := e.Group("")
	admin := e.Group("/admin")
	seller := e.Group("/seller")
	user := e.Group("/user")

	return &httpServer{
		handler: h,
		public:  public,
		admin:   admin,
		seller:  seller,
		user:    user,
	}
}

func (s *httpServer) Start(port int) error {
	s.setRoutes()

	if port == 0 {
		port = 8083
	}

	return e.Start(fmt.Sprintf(":%d", port))
}
