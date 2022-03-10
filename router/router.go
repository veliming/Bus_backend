package router

import (
	"Bus/db"
	"Bus/middleware"
	"Bus/pkg/setting"
	"Bus/router/api"
	"github.com/gin-gonic/gin"
)

func InitRouter(db *db.Session) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	r.Use(middleware.Cors())

	gin.SetMode(setting.RunMode)

	userHandler := api.UserAPI{DB: db}
	busHandler := api.BusAPI{DB: db}
	feedbackHandler := api.FeedbackAPI{DB: db}

	user := r.Group("/user")
	{
		userWithJWT := user.Group("/")
	//	userWithJWT.Use(middleware.JWT())
	//	{
			userWithJWT.POST("/session", userHandler.Code2session)
			userWithJWT.GET("/agreement", userHandler.AccceptAgreement)
	//	}
	//	user.POST("/getusercode", userHandler.GetUserCode)
	//	user.GET("/getuserid", userHandler.CreateUser)

	}
	bus := r.Group("/bus")
	{
		busWithAuth := bus.Group("/")
		//busWithAuth.Use(middleware.BusAUTH())
		//{
			busWithAuth.POST("/position", busHandler.PositionUpload)
		//}

		busWithJWT := bus.Group("/")
		//busWithJWT.Use(middleware.JWT())
		//{
			busWithJWT.POST("/logon", busHandler.LogonBus)
			busWithJWT.POST("/logout", busHandler.LogoutBus)
			busWithJWT.PUT("/", busHandler.OfflineBus)

			busWithJWT.GET("/all", busHandler.GetAllBus)
		//}
	}
	feedback := r.Group("/feedback")
	feedback.Use(middleware.JWT())
	{
		feedback.POST("/", feedbackHandler.CreateFB)
		feedback.PUT("/", feedbackHandler.CheckFB)
	}
	return r
}
