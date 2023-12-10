package service

import "context"

type RedisService interface {
	Publish(ctx context.Context, msg []byte) error
	Subscribe(ctx context.Context) error
}
