package api

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

type Server struct {
	maxSize int

	server  *echo.Echo
	address string

	us Usecase
}

func NewServer(ip string, port int, maxSize int, us Usecase) *Server {
	api := Server{
		maxSize: maxSize,
		us:      us,
	}

	api.server = echo.New()
	//handlers
	api.server.GET("/count", api.GetCount)
	api.server.POST("/count", api.PostCount)
	api.address = fmt.Sprintf("%s:%d", ip, port)
	return &api
}

func (api *Server) Run() {
	api.server.Logger.Fatal(api.server.Start(api.address))
}
