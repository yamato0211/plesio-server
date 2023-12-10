//go:build wireinject
// +build wireinject

package injection

import (
	"github.com/google/wire"
	"github.com/yamato0211/plesio-server/pkg/infra/mysql"
	"github.com/yamato0211/plesio-server/pkg/usecase"
	"github.com/yamato0211/plesio-server/pkg/utils/config"
)

func InitializeUserUsecase() usecase.IUserUsecase {
	wire.Build(
		config.NewDBConfig,
		mysql.NewMySQLConnector,
		mysql.NewUserRepository,
		usecase.NewUserUsecase,
	)
	return &usecase.UserUsecase{}
}
