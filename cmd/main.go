package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"

	"github.com/mohsanabbas/loggi/pkg/application"
)

var (
	e = echo.New()
)

func main() {
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 1 << 10,
		LogLevel:  log.ERROR,
	}))
	e.Use(middleware.Timeout())
	e.Use(middleware.CORS())
	application.StartApp(e)

	// start the server
	go func() {
		fmt.Println("Starting server on port 9090")
		if err := e.Start(":9090"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	sig := <-quit
	fmt.Printf("Got signal:%v", sig)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
