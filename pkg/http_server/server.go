package httpserver

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	Config Config
	Router *echo.Echo
}

func New(cfg Config) *Server {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: cfg.Cors.AllowOrigins,
	}))

	return &Server{
		Config: cfg,
		Router: e,
	}
}

// register custom handler
func (s Server) RegisterHandler(route string, handler echo.HandlerFunc) {
	s.Router.GET(route, handler)
}

// start server
func (s Server) Start() error {
	return s.Router.Start(fmt.Sprintf(":%d", s.Config.Port))
}

func (s Server) Stop(ctx context.Context) error {
	return s.Router.Shutdown(ctx)
}
