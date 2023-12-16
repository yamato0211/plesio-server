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

type RedisConfig struct {
	RedisEndpoint string
}

func NewDBConfig() *DBConfig {
	cfg := &DBConfig{
		DBHost: LookUpEnv("MYSQL_HOST", "db"),
		DBName: LookUpEnv("MYSQL_DATABASE", "main"),
		DBUser: LookUpEnv("MYSQL_USER", "user"),
		DBPass: LookUpEnv("MYSQL_PASSWORD", "password"),
		DBPort: LookUpEnv("MYSQL_PORT", "3306"),
	}
	return cfg
}

func NewRedisConfig() *RedisConfig {
	cfg := &RedisConfig{
		RedisEndpoint: LookUpEnv("REDIS_ENDPOINT", "redis:6379"),
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
