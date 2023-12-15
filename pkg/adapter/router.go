package adapter

import (
	http "github.com/yamato0211/plesio-server/pkg/adapter/http/handler"
	ws "github.com/yamato0211/plesio-server/pkg/adapter/ws/handler"
)

type MasterHandler struct {
	Ws   *ws.WebSocketHandler
	User *http.UserHandler
}

func NewMasterHandler(ws *ws.WebSocketHandler, user *http.UserHandler) *MasterHandler {
	return &MasterHandler{
		Ws:   ws,
		User: user,
	}
}
