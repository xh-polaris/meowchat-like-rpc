package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Like struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserId       string             `bson:"userId,omitempty" json:"userId,omitempty"`
	TargetId     string             `bson:"targetId,omitempty" json:"targetId,omitempty"`
	Type         int                `bson:"targetType,omitempty" json:"targetType,omitempty"`
	Liked        bool               `bson:"liked,omitempty" json:"liked,omitempty"`
	AssociatedId string             `bson:"associatedId,omitempty" json:"associatedId,omitempty"`
	UpdateAt     time.Time          `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt     time.Time          `bson:"createAt,omitempty" json:"createAt,omitempty"`
}
