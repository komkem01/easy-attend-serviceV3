package redis

// import (
// 	"context"

// 	"github.com/easy-attend-serviceV3/internal/provider"
// )

// type RedisModule struct {
// 	Svc *RedisService
// }

// var _ provider.Close = (*RedisModule)(nil)

// // If Option is not defined, define it here or import the correct type.
// // For now, let's define a placeholder Option type in this file for compilation.
// // Remove this and use the correct import if Option exists elsewhere.

// type Option struct {
// 	// Add fields as needed
// }

// func New(appEnv string, opts map[string]*Option) *RedisModule {
// 	svc := newService(appEnv, opts)
// 	return &RedisModule{
// 		Svc: svc,
// 	}
// }

// func (db *RedisModule) Close(ctx context.Context) error {
// 	return db.Svc.close(ctx)
// }
