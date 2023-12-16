package adapter

import (
	http "github.com/yamato0211/plesio-server/pkg/adapter/http/handler"
)

type MasterHandler struct {
	User *http.UserHandler
}

func NewMasterHandler(user *http.UserHandler) *MasterHandler {
	return &MasterHandler{
		User: user,
	}
}
