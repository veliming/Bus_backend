package api

import (
	"Bus/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
	"time"
)

type BusDB interface {
	GetBuses() ([]*models.Bus, error)
	CreateBus(bus *models.Bus) (primitive.ObjectID, error)
	DeleteBusByID(id primitive.ObjectID) (int64, error)
	UpdateBusByID(id primitive.ObjectID, newDoc *bson.M) (int64, error)
}

type BusAPI struct {
	DB BusDB
}

func (a *BusAPI) LogonBus(ctx *gin.Context) {
	//_, check := util.CheckRole(ctx, 1)
	//if !check {
	//	ctx.JSON(http.StatusForbidden, gin.H{"error": "无权操作"})
	//	return
	//}
	type reqmsg struct {
		Type   string `json:"type"`
		Status int    `json:"status"`
	}
	var req reqmsg
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
		log.Print(err)
		return
	}
	obj := &models.Bus{
		ID:     primitive.NewObjectID(),
		Type:   req.Type,
		Status: req.Status,
		Position: models.Position{
			Latitude:  0,
			Longitude: 0,
		},
		Speed: 0.0,
		UpdateTime: time.Now(),
	}
	objID, err := a.DB.CreateBus(obj)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "数据库操作失败"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"error": "0", "id": objID.Hex()})
}

func (a *BusAPI) LogoutBus(ctx *gin.Context) {
	//_, check := util.CheckRole(ctx, 1)
	//if !check {
	//	ctx.JSON(http.StatusForbidden, gin.H{"error": "无权操作"})
	//	return
	//}
	type reqmsg struct {
		ID string `json:"id"`
	}
	var req reqmsg
	err := ctx.ShouldBindJSON(&req)
	if req.ID == "" || err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
		return
	}
	ObjID, err := primitive.ObjectIDFromHex(req.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ObjectID错误"})
		return
	}
	deleteCount, err := a.DB.DeleteBusByID(ObjID)
	if deleteCount != 1 || err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "数据库操作失败"})
	}
	ctx.JSON(http.StatusOK, gin.H{"error": "0"})
}

func (a *BusAPI) PositionUpload(ctx *gin.Context) {
	type reqmsg struct {
		ID       string          `json:"i"`
		Position models.Position `json:"p"`
		Speed    float64		`json:"s"`
	}
	var req reqmsg
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
		return
	}
	objID, err := primitive.ObjectIDFromHex(req.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ObjectID错误"})
		return
	}
	temp := &bson.M{
		"$set": bson.M{"position": req.Position, "speed":req.Speed,"updatetime": time.Now()},
	}
	count, err := a.DB.UpdateBusByID(objID, temp)
	if err != nil || count != 1 {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "数据库操作失败"})
	}
	ctx.JSON(http.StatusOK, gin.H{"error": "0"})
}

func (a *BusAPI) GetAllBus(ctx *gin.Context) {
	var buses []*models.Bus
	buses, err := a.DB.GetBuses()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "获取数据失败"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"error": "0", "bused": buses})
}

func (a *BusAPI) OfflineBus(ctx *gin.Context) {
	type reqmsg struct {
		ID     string `json:"id"`
		Status int    `json:"status"`
	}
	var req reqmsg
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
		return
	}
	objID, err := primitive.ObjectIDFromHex(req.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ObjectID错误"})
		return
	}
	count, err := a.DB.UpdateBusByID(objID, &bson.M{"$set": bson.M{"status": req.Status}})
	if err != nil || count != 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "数据库操作失败"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"error": "0"})
}
