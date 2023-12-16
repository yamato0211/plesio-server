package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/yamato0211/plesio-server/pkg/injection"
)

func main() {
	e := echo.New()
	e.Use(middleware.Recover())
	// e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	// DI
	mh := injection.InitializeMasterHandler()

	// Health Check
	e.GET("/", func(c echo.Context) error {
		fmt.Fscanf(os.Stderr, "Hello, World!\n")
		return c.String(http.StatusOK, "Hello, World!")
	})

	// Routing
	api := e.Group("/api/v1")
	{
		ws := api.Group("/ws")
		{
			ws.GET("/", mh.Ws.Handle())
		}
		user := api.Group("/users")
		{
			user.GET("/:id", mh.User.GetUser())
		}
		redis := api.Group("/redis")
		{
			redis.GET("/:key", mh.Redis.Ping())
		}
	}

	go func() {
		if err := e.Start(":8080"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
