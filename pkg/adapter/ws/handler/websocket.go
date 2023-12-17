package handler

import (
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/yamato0211/plesio-server/pkg/web/ws"
)

type WebSocketHandler struct {
	upgrader websocket.Upgrader
	hub      *ws.Hub
}

func NewWebSocketHandler(hub *ws.Hub) *WebSocketHandler {
	return &WebSocketHandler{
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
			HandshakeTimeout: time.Duration(60 * time.Second),
		},
		hub: hub,
	}
}

func (h *WebSocketHandler) Handle() echo.HandlerFunc {
	return func(c echo.Context) error {
		conn, err := h.upgrader.Upgrade(c.Response(), c.Request(), nil)
		if err != nil {
			c.Logger().Errorf("failed to upgrade: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "failed to upgrade")
		}
		conn.SetReadDeadline(time.Now().Add(10 * time.Minute))
		conn.SetWriteDeadline(time.Now().Add(10 * time.Minute))
		client := ws.NewClient(conn)
		go client.ReadLoop(h.hub.BroadcastCh, h.hub.UnRegisterCh)
		go client.WriteLoop()
		h.hub.RegisterCh <- client
		return nil
	}
}
