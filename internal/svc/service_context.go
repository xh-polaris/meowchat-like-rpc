package svc

import (
	"github.com/xh-polaris/meowchat-like-rpc/internal/config"
	"github.com/xh-polaris/meowchat-like-rpc/internal/model"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type ServiceContext struct {
	Config config.Config
	model.LikeModel
	*redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		LikeModel: model.NewLikeModel(c.Mongo.URL, c.Mongo.DB, model.LikeCollectionName, c.CacheConf),
		Redis:     c.Redis.NewRedis(),
	}
}
