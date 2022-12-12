package logic

import (
	"context"

	"like-rpc/internal/svc"
	"like-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DoLikeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDoLikeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DoLikeLogic {
	return &DoLikeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  点赞/取消赞
func (l *DoLikeLogic) DoLike(in *pb.DoLikeReq) (*pb.DoLikeResp, error) {
	// todo: add your logic here and delete this line

	return &pb.DoLikeResp{}, nil
}
