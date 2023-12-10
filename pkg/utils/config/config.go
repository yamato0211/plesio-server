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
		DBHost: LookUpEnv("POSTGRES_USER_HOST", "db"),
		DBName: LookUpEnv("POSTGRES_DATABASE", "main"),
		DBUser: LookUpEnv("POSTGRES_USER", "user"),
		DBPass: LookUpEnv("POSTGRES_PASSWORD", "password"),
		DBPort: LookUpEnv("POSTGRES_USER_PORT", "3306"),
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
