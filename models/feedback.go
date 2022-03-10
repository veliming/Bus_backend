package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Feedback struct {
	ID               primitive.ObjectID `bson:"_id"`
	Text             string             `bson:"text"`
	Status           int                `bson:"status"`
	Type             string             `bson:"type"`
	UserID           primitive.ObjectID `bson:"user_id"`
	CreateTime       time.Time          `bson:"createtime"`
	LastModifiedTime time.Time          `bson:"lastmodifiedtime"`
}
