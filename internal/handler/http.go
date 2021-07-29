package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mohsanabbas/loggi/internal/service"
)

type HttpHandler interface {
	Print(w echo.Context) error
}

type httpHandler struct {
	service service.Service
}

func NewHttpHandler(s service.Service) HttpHandler {
	return &httpHandler{
		service: s,
	}
}

func (h *httpHandler) Print(c echo.Context) error {
	response, err := h.service.PrintInService("HI from http")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, response)

}
