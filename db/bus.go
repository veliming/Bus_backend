package db

import (
	"Bus/models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *Session) CreateBus(bus *models.Bus) (primitive.ObjectID, error) {
	res, err := s.DB.Collection("buses").InsertOne(context.Background(), bus)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return res.InsertedID.(primitive.ObjectID), nil
}

func (s *Session) DeleteBusByID(id primitive.ObjectID) (int64, error) {
	res, err := s.DB.Collection("buses").DeleteOne(context.Background(), bson.D{{Key: "_id", Value: id}})
	if err != nil {
		return 0, err
	}
	return res.DeletedCount, err
}

func (s *Session) GetBuses() ([]*models.Bus, error) {
	var buses []*models.Bus
	cursor, err := s.DB.Collection("buses").Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		bus := &models.Bus{}
		if err := cursor.Decode(bus); err != nil {
			return nil, err
		}
		buses = append(buses, bus)
	}

	return buses, nil
}

func (s *Session) GetBusByIDs(ids []primitive.ObjectID) ([]*models.Bus, error) {
	var buses []*models.Bus
	cursor, err := s.DB.Collection("buses").
		Find(context.Background(), bson.D{{
			Key: "_id",
			Value: bson.D{{
				Key:   "$in",
				Value: ids,
			}},
		}})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		bus := &models.Bus{}
		if err := cursor.Decode(bus); err != nil {
			return nil, err
		}
		buses = append(buses, bus)
	}

	return buses, nil
}

func (s *Session) UpdateBusByID(id primitive.ObjectID, newDoc *bson.M) (int64, error) {
	res, err := s.DB.Collection("buses").UpdateOne(context.Background(),
		bson.M{"_id": id}, newDoc)
	if err != nil {
		return 0, err
	}
	return res.ModifiedCount, nil
}
