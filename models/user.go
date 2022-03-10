package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id"`
	Nickname string             `bson:"nickname"`
	Openid   string             `bson:"openid"`
	SchoolID string             `bson:"schoolid"`
	Regtime  time.Time          `bson:"regtime"`
	Role     int                `bson:"role"` // -1 -> 未同意协议 0 -> 已同意协议 1 -> 管理员
}

func (u *User) New() *User {
	return &User{
		ID:       primitive.NewObjectID(),
		Nickname: "",
		Openid:   "",
		SchoolID: "",
		Regtime:  time.Now(),
		Role:     -1,
	}
}
