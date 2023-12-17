package config

import (
	"os"

	"github.com/joho/godotenv"
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

type GithubConfig struct {
	ClientID     string
	ClientSecret string
}

func NewGithubConfig() *GithubConfig {
	cfg := &GithubConfig{
		ClientID:     LookUpEnv("GITHUB_CLIENT_ID", ""),
		ClientSecret: LookUpEnv("GITHUB_CLIENT_SECRET", ""),
	}
	return cfg
}

func NewDBConfig() *DBConfig {
	// cfg := &DBConfig{
	// 	DBHost: LookUpEnv("MYSQL_HOST", "127.0.0.1"),
	// 	DBName: LookUpEnv("MYSQL_DATABASE", "main"),
	// 	DBUser: LookUpEnv("MYSQL_USER", "admin"),
	// 	DBPass: LookUpEnv("MYSQL_PASSWORD", "kumayama0211"),
	// 	DBPort: LookUpEnv("MYSQL_PORT", "3306"),
	// }

	godotenv.Load()

	cfg := &DBConfig{
		DBHost: LookUpEnv("MYSQL_HOST", "127.0.0.1"),
		DBName: LookUpEnv("MYSQL_DATABASE", "db"), //main
		DBUser: LookUpEnv("MYSQL_USER", "admin"),
		DBPass: LookUpEnv("MYSQL_PASSWORD", "kumayama0211"),

		DBPort: LookUpEnv("MYSQL_PORT", "3306"),
	}
	return cfg
}

func NewRedisConfig() *RedisConfig {
	godotenv.Load()
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
