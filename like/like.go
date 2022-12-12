// Code generated by goctl. DO NOT EDIT!
// Source: like.proto

package like

import (
	"context"

	"like-rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	DoLikeReq          = pb.DoLikeReq
	DoLikeResp         = pb.DoLikeResp
	GetTargetLikesReq  = pb.GetTargetLikesReq
	GetTargetLikesResp = pb.GetTargetLikesResp
	GetUserLikedReq    = pb.GetUserLikedReq
	GetUserLikedResp   = pb.GetUserLikedResp

	Like interface {
		//  点赞/取消赞
		DoLike(ctx context.Context, in *DoLikeReq, opts ...grpc.CallOption) (*DoLikeResp, error)
		//  获取用户是否点赞
		GetUserLike(ctx context.Context, in *GetUserLikedReq, opts ...grpc.CallOption) (*GetUserLikedResp, error)
		//  获取目标点赞数
		GetTargetLikes(ctx context.Context, in *GetTargetLikesReq, opts ...grpc.CallOption) (*GetTargetLikesResp, error)
	}

	defaultLike struct {
		cli zrpc.Client
	}
)

func NewLike(cli zrpc.Client) Like {
	return &defaultLike{
		cli: cli,
	}
}

//  点赞/取消赞
func (m *defaultLike) DoLike(ctx context.Context, in *DoLikeReq, opts ...grpc.CallOption) (*DoLikeResp, error) {
	client := pb.NewLikeClient(m.cli.Conn())
	return client.DoLike(ctx, in, opts...)
}

//  获取用户是否点赞
func (m *defaultLike) GetUserLike(ctx context.Context, in *GetUserLikedReq, opts ...grpc.CallOption) (*GetUserLikedResp, error) {
	client := pb.NewLikeClient(m.cli.Conn())
	return client.GetUserLike(ctx, in, opts...)
}

//  获取目标点赞数
func (m *defaultLike) GetTargetLikes(ctx context.Context, in *GetTargetLikesReq, opts ...grpc.CallOption) (*GetTargetLikesResp, error) {
	client := pb.NewLikeClient(m.cli.Conn())
	return client.GetTargetLikes(ctx, in, opts...)
}