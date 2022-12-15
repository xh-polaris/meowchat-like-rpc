package logic

import (
	"context"
	"github.com/xh-polaris/meowchat-like-rpc/internal/svc"
	"github.com/xh-polaris/meowchat-like-rpc/pb"
	"github.com/zeromicro/go-zero/core/stores/monc"

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

// GetUserLike 获取用户是否点赞
func (l *GetUserLikeLogic) GetUserLike(in *pb.GetUserLikedReq) (*pb.GetUserLikedResp, error) {
	likeModel := l.svcCtx.LikeModel
	err := likeModel.GetUserLike(l.ctx, in.UserId, in.TargetId, in.Type)
	switch err {
	case nil:
		return &pb.GetUserLikedResp{Liked: true}, nil
	case monc.ErrNotFound:
		return &pb.GetUserLikedResp{Liked: false}, nil
	}
	return &pb.GetUserLikedResp{}, nil
}
