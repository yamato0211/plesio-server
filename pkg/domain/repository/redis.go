package repository

import "github.com/labstack/echo/v4"

type RedisRepository interface {
	// Publish(ctx context.Context, msg []byte) error
	// Subscribe(ctx context.Context) error
	Ping(ctx echo.Context) error
}
