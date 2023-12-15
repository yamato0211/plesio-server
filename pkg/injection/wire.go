//go:build wireinject
// +build wireinject

package injection

import (
	"github.com/google/wire"
	"github.com/yamato0211/plesio-server/pkg/adapter"
	http_handler "github.com/yamato0211/plesio-server/pkg/adapter/http/handler"
	ws_handler "github.com/yamato0211/plesio-server/pkg/adapter/ws/handler"
	"github.com/yamato0211/plesio-server/pkg/infra/mysql"
	"github.com/yamato0211/plesio-server/pkg/usecase"
	"github.com/yamato0211/plesio-server/pkg/utils/config"
	"github.com/yamato0211/plesio-server/pkg/web/ws"
)

func InitializeMasterHandler() *adapter.MasterHandler {
	wire.Build(
		config.NewDBConfig,
		mysql.NewMySQLConnector,
		mysql.NewUserRepository,
		usecase.NewUserUsecase,
		http_handler.NewUserHandler,
		ws.NewHub,
		ws_handler.NewWebSocketHandler,
		adapter.NewMasterHandler,
	)
	return &adapter.MasterHandler{}
}
