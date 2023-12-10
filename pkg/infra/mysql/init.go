package mysql

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/yamato0211/plesio-server/pkg/utils/config"
)

const driverName = "mysql"

func NewMySQLConnector(cfg *config.DBConfig) *sqlx.DB {
	dsn := mysqlConnDSN(cfg)
	db, err := sqlx.Open(driverName, dsn)
	if err != nil {
		fmt.Println(err.Error())
		log.Fatal(err)
	}
	return db
}

func mysqlConnDSN(cfg *config.DBConfig) string {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local",
		cfg.DBUser,
		cfg.DBPass,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)
	return dsn
}
