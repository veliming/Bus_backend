package api

import (
	"Bus/models"
	"Bus/pkg/setting"
	"Bus/pkg/util"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)

type UserDB interface {
	GetUserByID(id primitive.ObjectID) *models.User
	GetUserByOpenid(openid string) *models.User
	CreateUser(user *models.User) (primitive.ObjectID, error)
	UpdateUserByID(id primitive.ObjectID, newDoc *bson.M) (int64, error)
}

type UserAPI struct {
	DB UserDB
}

func (a *UserAPI) Code2session(ctx *gin.Context) {
	type result struct {
		ErrorCode  int    `json:"errcode"`
		ErrorMsg   string `json:"errmsg,omitempty"`
		SessionKey string `json:"session_key,omitempty"`
		ExpiresIn  int    `json:"expires_in,omitempty"`
		Openid     string `json:"openid,omitempty"`
	}
	type reqmsg struct {
		Code string `json:"code"`
	}
	var req reqmsg
	err := ctx.ShouldBindJSON(&req)
	if len(req.Code) == 0 || err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": "未获取到code"})
	}
	baseUrl := "https://api.weixin.qq.com/sns/jscode2session?"
	opts := fmt.Sprintf("appid=%s&secret=%s&js_code=%s&grant_type=authorization_code",
		setting.AppID, setting.AppSecret, req.Code)
	res, err := http.Get(baseUrl + opts)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": "请求微信服务器失败"})
		return
	}
	defer res.Body.Close()
	r := result{}
	err = json.NewDecoder(res.Body).Decode(r)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": "微信服务器返回不可预期值"})
		return
	}
	if r.ErrorCode != 0 {
		ctx.JSON(http.StatusOK, gin.H{"error": "微信验证失败"})
		return
	}
	user := a.DB.GetUserByOpenid(r.Openid)
	if user == nil {
		newUser := &models.User{
			ID:       primitive.NewObjectID(),
			Nickname: "",
			Openid:   r.Openid,
			SchoolID: "",
			Regtime:  time.Now(),
			Role:     -1,
		}
		res, err := a.DB.CreateUser(newUser)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{"error": "创建新用户失败"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"error": "0", "user_id": res.Hex()})
		return
	}

	token, err := util.GenerateToken(user.ID, user.Role)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": "创建新用户失败"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"error": "0", "user_id": user.ID.Hex()})
	ctx.Header("Authorization", token)
}

func (a *UserAPI) CreateUser(ctx *gin.Context) {
	newUser := &models.User{
		ID:       primitive.NewObjectID(),
		Nickname: "",
		Openid:   "test",
		SchoolID: "",
		Regtime:  time.Now(),
		Role:     1,
	}
	res, err := a.DB.CreateUser(newUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "数据库操作失败"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"error": "0", "user_id": res.Hex()})
	return
}

func (a *UserAPI) GetUserCode(ctx *gin.Context) {
	type reqmsg struct {
		UserID string `json:"user_id"`
	}
	var req reqmsg
	err := ctx.ShouldBindJSON(&req)
	if len(req.UserID) == 0 || err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
	}
	id, _ := primitive.ObjectIDFromHex(req.UserID)
	user := a.DB.GetUserByID(id)
	token, _ := util.GenerateToken(user.ID, user.Role)
	ctx.Header("Authorization", token)
}

func (a *UserAPI) AccceptAgreement(ctx *gin.Context) {
	claims, check := util.CheckRole(ctx, -1)
	if !check {
		ctx.JSON(http.StatusOK, gin.H{"error": "不可重复操作"})
		return
	}
	count, err := a.DB.UpdateUserByID(claims.ID, &bson.M{"set": bson.M{"role": 0}})
	if err != nil || count != 1 {
		ctx.JSON(http.StatusOK, gin.H{"error": "数据库操作失败"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"error": "0"})
	return
}
