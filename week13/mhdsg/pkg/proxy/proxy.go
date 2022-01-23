package proxy

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

// ProviderSet is facade providers.
//var ProviderSet = wire.NewSet(NewRedisHelper)

func newRedis(urlStr string) (*redis.Client, error) {
	url, err := redis.ParseURL(urlStr)
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()
	rdb := redis.NewClient(url)
	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}

	return rdb, nil
}
