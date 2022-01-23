package proxy

import (
	"context"
	"encoding/json"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"time"
)

type IRedisHelper interface {
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) bool
	Get(ctx context.Context, key string, dest interface{}) bool
}

type redisHelper struct {
	rdb *redis.Client
	log *log.Helper
}

func NewRedisHelper(urlStr string, l log.Logger) IRedisHelper {
	client, _ := newRedis(urlStr)
	return &redisHelper{rdb: client, log: log.NewHelper(log.With(l, "module", "pkg/facade/redis"))}
}

func (h *redisHelper) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) bool {
	p, err := json.Marshal(value)
	if err != nil {
		h.log.Errorf("Marshal error before set, key:%s value:%s", key, value)
		return false
	}
	err = h.rdb.Set(ctx, key, p, expiration).Err()
	if err != nil {
		h.log.Errorf("set error, key:%s value:%s", key, value)
		return false
	}
	return true
}

func (h *redisHelper) Get(ctx context.Context, key string, dest interface{}) bool {
	cmd := h.rdb.Get(ctx, key)
	if cmd.Err() == redis.Nil {
		return false
	} else if cmd.Err() != nil {
		h.log.Errorf("Get error, key:%s", key)
		return false
	} else {
		result, _ := cmd.Bytes()
		json.Unmarshal(result, dest)
		return true
	}
}
