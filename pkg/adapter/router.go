package adapter

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	http_handler "github.com/yamato0211/plesio-server/pkg/adapter/http/handler"
	ws_handler "github.com/yamato0211/plesio-server/pkg/adapter/ws/handler"
	"github.com/yamato0211/plesio-server/pkg/injection"
	websocket "github.com/yamato0211/plesio-server/pkg/web/ws"
)

func InitRouter() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	// DI
	userUsecase := injection.InitializeUserUsecase()

	ws := e.Group("/ws")
	{
		hub := websocket.NewHub()
		handler := ws_handler.NewWebSocketHandler(hub)
		ws.GET("/", handler.Handle)
	}
	user := e.Group("/users")
	{
		handler := http_handler.NewUserHandler(userUsecase)
		user.GET("/:id", handler.GetUser())
	}
	return e
}
