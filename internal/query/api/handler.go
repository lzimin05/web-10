package api

import (
	"errors"
	"net/http"

	"github.com/ValeryBMSTU/web-10/pkg/vars"
	"github.com/labstack/echo/v4"
)

func (srv *Server) GetHelloUser(e echo.Context) error {
	msg, err := srv.us.FetchHelloUserMessage()
	if err != nil {
		return e.String(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusOK, msg)
}

func (srv *Server) PostUser(e echo.Context) error {
	input := e.FormValue("name")
	err := srv.us.SetHelloUserMessage(input)
	if err != nil {
		if errors.Is(err, vars.ErrAlreadyExist) {
			return e.String(http.StatusConflict, err.Error())
		}
		return e.String(http.StatusInternalServerError, err.Error())
	}
	return e.String(http.StatusCreated, "OK")
}
