package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/monc"
	"go.mongodb.org/mongo-driver/bson"
)

const LikeCollectionName = "like"

var _ LikeModel = (*CustomLikeModel)(nil)

type (
	// LikeModel is an interface to be customized, add more methods here,
	// and implement the added methods in CustomLikeModel.
	LikeModel interface {
		likeModel
		GetUserLike(ctx context.Context, userId string, targetId string, targetType int32) error
		GetTargetLikes(ctx context.Context, targetId string, targetType int32) (int64, error)
		GetId(ctx context.Context, userId string, targetId string, targetType int32) (string, error)
	}

	CustomLikeModel struct {
		*defaultLikeModel
	}
)

func (m *CustomLikeModel) GetId(ctx context.Context, userId string, targetId string, targetType int32) (id string, err error) {
	like := Like{}
	err = m.conn.FindOneNoCache(ctx, &like, bson.M{"userId": userId, "targetId": targetId, "targetType": targetType})
	id = like.ID.Hex()
	return
}

func (m *CustomLikeModel) GetTargetLikes(ctx context.Context, targetId string, targetType int32) (count int64, err error) {
	count, err = m.conn.CountDocuments(ctx, bson.M{"targetId": targetId, "targetType": targetType})
	if err != nil {
		return 0, err
	} else {
		return count, nil
	}
}

func (m *CustomLikeModel) GetUserLike(ctx context.Context, userId string, targetId string, targetType int32) (err error) {
	like := Like{}
	err = m.conn.FindOneNoCache(ctx, &like, bson.M{"userId": userId, "targetId": targetId, "targetType": targetType})
	return
}

// NewLikeModel returns a model for the mongo.
func NewLikeModel(url, db, collection string, c cache.CacheConf) LikeModel {
	conn := monc.MustNewModel(url, db, collection, c)
	return &CustomLikeModel{
		defaultLikeModel: newDefaultLikeModel(conn),
	}
}
