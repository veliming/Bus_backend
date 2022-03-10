package api

import (
	"Bus/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)

type FeedbackDB interface {
	CreateFeedback(fb *models.Feedback) (primitive.ObjectID, error)
	UpdateFeedbackByID(id primitive.ObjectID, newDoc *bson.M) (int64, error)
	GetFeedbackByType(t string) ([]*models.Feedback, error)
	GetFeedbackByUserID(id primitive.ObjectID) ([]*models.Feedback, error)
}

type FeedbackAPI struct {
	DB FeedbackDB
}

func (a *FeedbackAPI) CreateFB(ctx *gin.Context) {
	type reqmsg struct {
		Text string `json:"text"`
		Type string `json:"type"`
	}
	var req reqmsg
	err := ctx.ShouldBindJSON(&req)
	if len(req.Text) == 0 || len(req.Type) == 0 || err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
	}
	newFB := &models.Feedback{
		ID:               primitive.NewObjectID(),
		Text:             req.Text,
		Status:           0,
		Type:             req.Type,
		CreateTime:       time.Now(),
		LastModifiedTime: time.Now(),
	}
	_, err = a.DB.CreateFeedback(newFB)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "数据库操作失败"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"error": "0"})
}

func (a *FeedbackAPI) CheckFB(ctx *gin.Context) {
	type reqmsg struct {
		ID string `json:"id"`
	}
	var req reqmsg
	err := ctx.ShouldBindJSON(&req)
	if len(req.ID) == 0 || err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
	}
	id, _ := primitive.ObjectIDFromHex(req.ID)
	count, err := a.DB.UpdateFeedbackByID(id,
		&bson.M{"$set": bson.M{"status": 1, "lastmodifiedtime": time.Now()}})
	if err != nil || count != 1 {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "数据库操作失败"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"error": "0"})
}
