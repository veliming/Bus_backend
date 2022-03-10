package db

import (
	"Bus/models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *Session) CreateFeedback(fb *models.Feedback) (primitive.ObjectID, error) {
	res, err := s.DB.Collection("feedbacks").InsertOne(context.Background(), fb)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return res.InsertedID.(primitive.ObjectID), nil
}

func (s *Session) GetFeedbackByType(t string) ([]*models.Feedback, error) {
	var fbs []*models.Feedback
	cur, err := s.DB.Collection("feedbacks").Find(
		context.Background(), bson.M{"type": t})
	if err != nil {
		return nil, err
	}
	for cur.Next(context.Background()) {
		fb := &models.Feedback{}
		err = cur.Decode(fb)
		if err != nil {
			return nil, err
		}
		fbs = append(fbs, fb)
	}
	return fbs, nil
}

func (s *Session) GetFeedbackByUserID(id primitive.ObjectID) ([]*models.Feedback, error) {
	var fbs []*models.Feedback
	cur, err := s.DB.Collection("feedbacks").Find(
		context.Background(), bson.M{"user_id": id})
	if err != nil {
		return nil, err
	}
	for cur.Next(context.Background()) {
		fb := &models.Feedback{}
		err = cur.Decode(fb)
		if err != nil {
			return nil, err
		}
		fbs = append(fbs, fb)
	}
	return fbs, nil
}

func (s *Session) UpdateFeedbackByID(id primitive.ObjectID, newDoc *bson.M) (int64, error) {
	res, err := s.DB.Collection("feedbacks").UpdateOne(
		context.Background(),
		bson.M{"_id": id},
		newDoc)
	if err != nil {
		return 0, err
	}
	return res.ModifiedCount, nil
}
