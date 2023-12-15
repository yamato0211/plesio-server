package config

import (
	"os"
)

type DBConfig struct {
	DBHost string
	DBName string
	DBUser string
	DBPass string
	DBPort string
}

func NewDBConfig() *DBConfig {
	cfg := &DBConfig{
		DBHost: LookUpEnv("MYSQL_HOST", "127.0.0.1"),
		DBName: LookUpEnv("MYSQL_DATABASE", "main"),
		DBUser: LookUpEnv("MYSQL_USER", "admin"),
		DBPass: LookUpEnv("MYSQL_PASSWORD", "kumayama0211"),
		DBPort: LookUpEnv("MYSQL_PORT", "3306"),
	}
	return cfg
}

func LookUpEnv(key string, fallback string) string {
	if e, ok := os.LookupEnv(key); ok {
		return e
	} else {
		return fallback
	}
}
