//go:build wireinject
// +build wireinject

package injection

import (
	"github.com/google/wire"
	"github.com/yamato0211/plesio-server/pkg/adapter"
	http_handler "github.com/yamato0211/plesio-server/pkg/adapter/http/handler"
	"github.com/yamato0211/plesio-server/pkg/infra/mysql"
	"github.com/yamato0211/plesio-server/pkg/infra/redis"
	"github.com/yamato0211/plesio-server/pkg/usecase"
	"github.com/yamato0211/plesio-server/pkg/utils/config"
	"github.com/yamato0211/plesio-server/pkg/web/ws"
)

func InitializeMasterHandler() *adapter.MasterHandler {
	wire.Build(
		config.NewDBConfig,
		mysql.NewMySQLConnector,
		mysql.NewUserRepository,
		mysql.NewItemRepository,
		mysql.NewUsersItemsRepository,
		usecase.NewUserUsecase,
		usecase.NewItemUsecase,
		usecase.NewUsersItemsUseCase,
		http_handler.NewUserHandler,
		http_handler.NewItemHandler,
		adapter.NewMasterHandler,
	)
	return &adapter.MasterHandler{}
}

func InitializeWebSocketHub() *ws.Hub {
	wire.Build(
		config.NewRedisConfig,
		redis.NewRedisRepository,
		ws.NewHub,
	)
	return &ws.Hub{}
}
