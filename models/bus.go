package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Position struct {
	Latitude  float64 `bson:"latitude" json:"la"`
	Longitude float64 `bson:"longitude" json:"lo"`
}

type Bus struct {
	ID         primitive.ObjectID `bson:"_id" json:"id"`
	Type       string             `bson:"type" json:"type"`
	Status     int                `bson:"status" json:"status"` // 0 -> 未上线
	Position   Position           `bson:"position" json:"position"`
	Speed 		float64				`bson:"speed" json:"s"`
	UpdateTime time.Time          `bson:"updatetime" json:"update_time"`
}

func (b *Bus) New() *Bus {
	return &Bus{
		ID:         primitive.NewObjectID(),
		Type:       "",
		Status:     1,
		Position:   Position{},
		Speed:		0.0,
		UpdateTime: time.Now(),
	}
}
