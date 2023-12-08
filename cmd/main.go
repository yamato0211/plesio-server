package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/yamato0211/plesio-server/pkg/adapter/handler"
	"github.com/yamato0211/plesio-server/pkg/web/ws"
)

func main() {
	e := echo.New()

	hub := ws.NewHub()
	handler := handler.NewWebSocketHandler(hub)

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello")
	})

	e.GET("/ws", handler.Handle)
	e.Logger.Fatal(e.Start(":8000"))
}
