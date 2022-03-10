package db

import (
	"Bus/models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *Session) CreateUser(user *models.User) (primitive.ObjectID, error) {
	res, err := s.DB.Collection("users").InsertOne(context.Background(), user)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return res.InsertedID.(primitive.ObjectID), nil
}

func (s *Session) DeleteUserByID(id primitive.ObjectID) (int64, error) {
	res, err := s.DB.Collection("users").DeleteOne(context.Background(), bson.D{{Key: "_id", Value: id}})
	if err != nil {
		return 0, err
	}
	return res.DeletedCount, err
}

func (s *Session) GetUserByName(name string) *models.User {
	var user *models.User
	err := s.DB.Collection("users").
		FindOne(context.Background(), bson.D{{Key: "name", Value: name}}).
		Decode(&user)
	if err != nil {
		return nil
	}
	return user
}

func (s *Session) GetUserByID(id primitive.ObjectID) *models.User {
	var user *models.User
	err := s.DB.Collection("users").FindOne(
		context.Background(), bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return nil
	}
	return user
}

func (s *Session) GetUserByIDs(ids []primitive.ObjectID) []*models.User {
	var users []*models.User
	cursor, err := s.DB.Collection("users").
		Find(context.Background(), bson.D{{
			Key: "_id",
			Value: bson.D{{
				Key:   "$in",
				Value: ids,
			}},
		}})
	if err != nil {
		return nil
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		user := &models.User{}
		if err := cursor.Decode(user); err != nil {
			return nil
		}
		users = append(users, user)
	}

	return users
}

func (s *Session) GetUserByOpenid(openid string) *models.User {
	var user *models.User
	err := s.DB.Collection("users").
		FindOne(context.Background(), bson.M{"openid": openid}).
		Decode(&user)
	if err != nil {
		return nil
	}
	return user
}

func (s *Session) UpdateUserByID(id primitive.ObjectID, newDoc *bson.M) (int64, error) {
	res, err := s.DB.Collection("users").UpdateOne(context.Background(),
		bson.M{"_id": id}, newDoc)
	// s.DB.Collection("users").FindOneAndUpdate(context.Background(), bson.M{"_id": id}, newDoc)
	if err != nil {
		return 0, err
	}
	return res.ModifiedCount, nil
}
