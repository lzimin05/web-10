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
	api.server.GET("/api/user", api.GetHelloUser)
	api.server.POST("/api/user", api.PostUser)
	api.address = fmt.Sprintf("%s:%d", ip, port)
	return &api
}

func (api *Server) Run() {
	api.server.Logger.Fatal(api.server.Start(api.address))
}
