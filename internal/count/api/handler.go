package api

import (
	"errors"
	"net/http"

	"github.com/ValeryBMSTU/web-10/pkg/vars"
	"github.com/labstack/echo/v4"
)

func (srv *Server) GetCount(e echo.Context) error {
	msg, err := srv.us.FetchCount()
	if err != nil {
		return e.String(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusOK, "count = "+msg)
}

func (srv *Server) PostCount(e echo.Context) error {
	input := struct {
		Msg *int `json:"count"`
	}{}
	err := e.Bind(&input)
	if err != nil {
		return e.String(http.StatusInternalServerError, err.Error())
	}
	if input.Msg == nil {
		return e.String(http.StatusBadRequest, "count is empty")
	}
	err = srv.us.SetCount(*input.Msg)
	if err != nil {
		if errors.Is(err, vars.ErrAlreadyExist) {
			return e.String(http.StatusConflict, err.Error())
		}
		return e.String(http.StatusInternalServerError, err.Error())
	}
	return e.String(http.StatusCreated, "OK")
}
