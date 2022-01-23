package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	conf "mhdsg/app/server/internal/conf"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewCoreMDB, NewCoreRedis, NewServerRepository)

// Data .
type Data struct {
	// TODO wrapped database client
	db  *mongo.Database
	rdb *redis.Client
}

//NewData .
func NewData(conf *conf.AppSettings, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	mdb, _ := NewCoreMDB(conf)
	rdb, _ := NewCoreRedis(conf)
	return &Data{
		db:  mdb,
		rdb: rdb,
	}, cleanup, nil
}

func NewCoreMDB(conf *conf.AppSettings) (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.NewClient(options.Client().ApplyURI(conf.Mongo.Server.ConnectionString))
	if err != nil {
		panic(err)
	}

	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}
	return client.Database(conf.Mongo.Server.Database), err
}

func NewCoreRedis(conf *conf.AppSettings) (*redis.Client, error) {
	url, err := redis.ParseURL(conf.Redis.Core)
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
