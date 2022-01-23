package service

import (
	"context"
	pb "mhdsg/api/server/service/v1"
	"mhdsg/app/server/internal/biz"
	"mhdsg/pkg/res"
)

type ServerService struct {
	pb.UnimplementedServerServer
	bizS *biz.ServerUseCase
}

func NewServerService(bizS *biz.ServerUseCase) *ServerService {
	return &ServerService{
		bizS: bizS,
	}
}

func (s *ServerService) GetServer(ctx context.Context, req *pb.GetServerRequestViewModel) (*res.ResultObj, error) {
	//serverContext, _ := transport.FromServerContext(ctx)
	//req.Platform = serverContext.RequestHeader().Get("taozhiqiang")
	info, err := s.bizS.GetUserServerInfo(ctx, req)
	if err != nil {
		return nil, err
	}

	return info, nil
}
