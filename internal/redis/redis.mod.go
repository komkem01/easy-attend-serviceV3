package redis

import (
	"context"

	"github.com/easy-attend-serviceV3/internal/config/dto"
	"github.com/easy-attend-serviceV3/internal/provider"
)

type RedisModule struct {
	Svc *RedisService
}

var _ provider.Close = (*RedisModule)(nil)

func New(appEnv string, opts map[string]*dto.Option) *RedisModule {
	svc := newService(appEnv, opts)
	return &RedisModule{
		Svc: svc,
	}
}

func (db *RedisModule) Close(ctx context.Context) error {
	return db.Svc.close(ctx)
}
