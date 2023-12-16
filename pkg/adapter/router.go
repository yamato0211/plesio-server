package adapter

import (
	http "github.com/yamato0211/plesio-server/pkg/adapter/http/handler"
	ws "github.com/yamato0211/plesio-server/pkg/adapter/ws/handler"
)

type MasterHandler struct {
	Ws    *ws.WebSocketHandler
	User  *http.UserHandler
	Redis *http.RedisHandler
}

func NewMasterHandler(ws *ws.WebSocketHandler, user *http.UserHandler, redis *http.RedisHandler) *MasterHandler {
	return &MasterHandler{
		Ws:    ws,
		User:  user,
		Redis: redis,
	}
}
