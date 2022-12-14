package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"like-rpc/errorx"
	"like-rpc/internal/model"
	"like-rpc/internal/svc"
	"like-rpc/pb"
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

// DoLike 点赞/取消赞
func (l *DoLikeLogic) DoLike(in *pb.DoLikeReq) (*pb.DoLikeResp, error) {
	// 判断是否点过赞
	getUserLikeLogic := NewGetUserLikeLogic(l.ctx, l.svcCtx)
	data := &pb.GetUserLikedReq{
		UserId:   in.UserId,
		TargetId: in.TargetId,
		Type:     in.Type,
	}
	response, _ := getUserLikeLogic.GetUserLike(data)
	switch response.Liked {
	case false:
		// 插入数据
		likeModel := l.svcCtx.LikeModel
		like := &model.Like{
			UserId:       in.UserId,
			TargetId:     in.TargetId,
			TargetType:   int(in.Type),
			AssociatedId: in.AssociatedId,
		}
		err := likeModel.Insert(l.ctx, like)
		if err == nil {
			return &pb.DoLikeResp{}, nil
		} else {
			return &pb.DoLikeResp{}, errorx.ErrDataBase
		}
	case true:
		likeModel := l.svcCtx.LikeModel
		ID, err := likeModel.GetId(l.ctx, in.UserId, in.TargetId, in.Type)
		if err != nil {
			return &pb.DoLikeResp{}, errorx.ErrDataBase
		}
		err = likeModel.Delete(l.ctx, ID)
		if err == nil {
			return &pb.DoLikeResp{}, nil
		} else {
			return &pb.DoLikeResp{}, errorx.ErrDataBase
		}
	default:
		return &pb.DoLikeResp{}, nil
	}
}
