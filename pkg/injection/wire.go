//go:build wireinject
// +build wireinject

package injection

import (
	"github.com/google/wire"
	"github.com/yamato0211/plesio-server/pkg/adapter"
	http_handler "github.com/yamato0211/plesio-server/pkg/adapter/http/handler"
	ws_handler "github.com/yamato0211/plesio-server/pkg/adapter/ws/handler"
	"github.com/yamato0211/plesio-server/pkg/infra/mysql"
	"github.com/yamato0211/plesio-server/pkg/infra/redis"
	"github.com/yamato0211/plesio-server/pkg/usecase"
	"github.com/yamato0211/plesio-server/pkg/utils/config"
	"github.com/yamato0211/plesio-server/pkg/web/ws"
)

func InitializeMasterHandler() *adapter.MasterHandler {
	wire.Build(
		config.NewDBConfig,
		config.NewRedisConfig,
		mysql.NewMySQLConnector,
		mysql.NewUserRepository,
		redis.NewRedisConnector,
		usecase.NewUserUsecase,
		usecase.NewRedisUsecase,
		http_handler.NewUserHandler,
		http_handler.NewRedisHandler,
		ws.NewHub,
		ws_handler.NewWebSocketHandler,
		adapter.NewMasterHandler,
	)
	return &adapter.MasterHandler{}
}
