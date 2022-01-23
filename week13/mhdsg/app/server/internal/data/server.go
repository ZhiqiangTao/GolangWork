package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"mhdsg/app/server/internal/biz"
	model "mhdsg/app/server/internal/entity"
)

var _ biz.IServerRepository = (*serverRepository)(nil)

type serverRepository struct {
	data                     *Data
	col_server_config        *mongo.Collection
	col_server_user_relation *mongo.Collection
	log                      *log.Helper
}

func NewServerRepository(data *Data, l log.Logger) biz.IServerRepository {
	return &serverRepository{
		data:                     data,
		col_server_config:        data.db.Collection("sg_server_config"),
		col_server_user_relation: data.db.Collection("sg_server_user_relation"),
		log:                      log.NewHelper(log.With(l, "module", "usecase/server")),
	}
}

func (s *serverRepository) GetServerByIdentity(ctx context.Context, identity int32) *model.Sg_server_config {
	res := &model.Sg_server_config{}
	err := s.col_server_config.FindOne(ctx, bson.D{{"identity", identity}}).Decode(&res)
	if err != nil {
		return nil
	}
	return res
}

func (s *serverRepository) GetServerUserByUid(ctx context.Context, uid int64) *model.Sg_server_user_relation {
	e := &model.Sg_server_user_relation{}
	err := s.col_server_user_relation.FindOne(ctx, bson.D{{"userid", uid}}).Decode(&e)
	if err != nil {
		return nil
	}
	return e
}
