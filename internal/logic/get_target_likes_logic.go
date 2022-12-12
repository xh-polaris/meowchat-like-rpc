package logic

import (
	"context"

	"like-rpc/internal/svc"
	"like-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTargetLikesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTargetLikesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTargetLikesLogic {
	return &GetTargetLikesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  获取目标点赞数
func (l *GetTargetLikesLogic) GetTargetLikes(in *pb.GetTargetLikesReq) (*pb.GetTargetLikesResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetTargetLikesResp{}, nil
}
