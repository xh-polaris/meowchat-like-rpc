package logic

import (
	"context"
	"github.com/xh-polaris/meowchat-like-rpc/errorx"
	"github.com/xh-polaris/meowchat-like-rpc/internal/svc"
	"github.com/xh-polaris/meowchat-like-rpc/pb"

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

// GetTargetLikes 获取目标点赞数
func (l *GetTargetLikesLogic) GetTargetLikes(in *pb.GetTargetLikesReq) (*pb.GetTargetLikesResp, error) {
	likeModel := l.svcCtx.LikeModel
	count, err := likeModel.GetTargetLikes(l.ctx, in.TargetId, in.Type)
	if err != nil {
		return &pb.GetTargetLikesResp{}, errorx.ErrDataBase
	} else {
		return &pb.GetTargetLikesResp{Count: count}, nil
	}
}
