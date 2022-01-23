package biz

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	pb "mhdsg/api/server/service/v1"
	"mhdsg/app/server/internal/conf"
	"mhdsg/app/server/internal/entity"
	"mhdsg/pkg/constant"
	"mhdsg/pkg/proxy"
	"mhdsg/pkg/res"
	"time"
)

type IServerRepository interface {
	GetServerByIdentity(ctx context.Context, identity int32) *entity.Sg_server_config
	GetServerUserByUid(ctx context.Context, uid int64) *entity.Sg_server_user_relation
}

type ServerUseCase struct {
	rep IServerRepository
	log *log.Helper
	rh  proxy.IRedisHelper
}

func NewServerUseCase(rep IServerRepository, logger log.Logger, conf *conf.AppSettings) *ServerUseCase {
	return &ServerUseCase{
		rep: rep,
		log: log.NewHelper(log.With(logger, "module", "usecase/server")),
		rh:  proxy.NewRedisHelper(conf.Redis.Core, logger),
	}
}

func (s *ServerUseCase) GetUserServerInfo(ctx context.Context, vm *pb.GetServerRequestViewModel) (*res.ResultObj, error) {
	//先从缓存读取
	key := fmt.Sprintf(constant.CacheKey_User_Server_Relation, vm.Userid)
	cache := &pb.UserServerResponseData{}
	if ok := s.rh.Get(ctx, key, cache); ok {
		return res.Success(cache), nil
	}

	server := s.rep.GetServerUserByUid(ctx, vm.Userid)
	if server == nil {
		return nil, errors.New(500, "NO_USER_SERVER_INFO", "无法获取用户服务器信息")
	}

	config := s.rep.GetServerByIdentity(ctx, server.Server_identity)
	if config == nil {
		return nil, errors.New(500, "NO_SERVER_INFO", "查无服务器信息")
	}

	rooms := []*pb.ChatRoomMapDetailData{}
	for _, item := range config.Rongyun.Chatroom_id_map {
		rooms = append(rooms, &pb.ChatRoomMapDetailData{
			BizType: item.K,
			Roomid:  item.V,
		})
	}
	d := &pb.UserServerResponseData{
		Userid:               vm.Userid,
		Identity:             config.Identity,
		Name:                 config.Name,
		PublicDomainGame:     config.Public_domain_game,
		PublicDomainPlatform: config.Public_domain_game,
		ChatroomIdMap:        rooms,
	}
	s.rh.Set(ctx, key, d, (1440 * time.Minute))
	return res.Success(d), nil
}
