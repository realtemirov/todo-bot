package redis

import (
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/realtemirov/projects/tgbot/config"
)

const (
	expireTime = 60 * 60 * 24
)

type RedisCache struct {
	rds *redis.Client
}

func NewRedis(cnf config.Config) (*RedisCache, error) {

	rds := redis.NewClient(&redis.Options{
		Addr:        fmt.Sprintf("%s:%s", cnf.Redis_HOST, cnf.Redis_PORT),
		Password:    cnf.Redis_PASS,
		DB:          cnf.Redis_DB,
		PoolTimeout: 10,
		PoolSize:    100,
	})

	pong := rds.Ping(rds.Context())
	if pong.Err() != nil {
		return nil, pong.Err()
	}
	return &RedisCache{rds: rds}, nil
}

func (r *RedisCache) Set(key string, value string) error {
	err := r.rds.Set(r.rds.Context(), key, value, time.Duration(time.Second*expireTime))
	return err.Err()
}

func (r *RedisCache) Get(key string) (string, error) {
	return r.rds.Get(r.rds.Context(), key).Result()
}

func (r *RedisCache) Del(key string) error {
	return r.rds.Del(r.rds.Context(), key).Err()
}

func (r *RedisCache) Close() error {
	return r.rds.Close()
}

func (r *RedisCache) Contains(key string) bool {
	return r.rds.Exists(r.rds.Context(), key).Val() == 1
}
