package logic

import (
	"context"

	"like-rpc/internal/svc"
	"like-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLikeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLikeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLikeLogic {
	return &GetUserLikeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  获取用户是否点赞
func (l *GetUserLikeLogic) GetUserLike(in *pb.GetUserLikedReq) (*pb.GetUserLikedResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetUserLikedResp{}, nil
}
