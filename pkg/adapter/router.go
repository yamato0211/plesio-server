package adapter

import (
	http "github.com/yamato0211/plesio-server/pkg/adapter/http/handler"
)

type MasterHandler struct {
	User *http.UserHandler
	Item *http.ItemHandler
}

func NewMasterHandler(user *http.UserHandler, item *http.ItemHandler) *MasterHandler {
	return &MasterHandler{
		User: user,
		Item: item,
	}
}
