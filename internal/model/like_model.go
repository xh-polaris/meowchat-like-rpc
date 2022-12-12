package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/monc"
)

const LikeCollectionName = "like"

var _ LikeModel = (*customLikeModel)(nil)

type (
	// LikeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customLikeModel.
	LikeModel interface {
		likeModel
	}

	customLikeModel struct {
		*defaultLikeModel
	}
)

// NewLikeModel returns a model for the mongo.
func NewLikeModel(url, db, collection string, c cache.CacheConf) LikeModel {
	conn := monc.MustNewModel(url, db, collection, c)
	return &customLikeModel{
		defaultLikeModel: newDefaultLikeModel(conn),
	}
}
