package application

import (
	"github.com/labstack/echo/v4"
	"github.com/mohsanabbas/loggi/internal/handler"
	"github.com/mohsanabbas/loggi/internal/repository"
	"github.com/mohsanabbas/loggi/internal/service"
)

func StartApp(e *echo.Echo) {
	handle := handler.NewHttpHandler(service.NewService(repository.NewRepository()))

	// Routes
	e.GET("/", handle.Print)

}
