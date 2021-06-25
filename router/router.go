package router

import (
	"github.com/gin-gonic/gin"
	"github.com/huahuayu/go-gin-app/handler/param"
	"github.com/huahuayu/go-gin-app/handler/user"
	"github.com/huahuayu/go-gin-app/middleware"
)

func Init(r *gin.Engine) {
	r.GET("/health", HealthGET)

	userGroup := r.Group("/user")
	userGroup.POST("/register", user.Register)
	userGroup.POST("/login", user.Login)
	userGroup.POST("/updatePassword", middleware.AuthMiddleware(), user.UpdatePassword)
	userGroup.POST("/updateUsername", middleware.AuthMiddleware(), user.UpdateUsername)
	userGroup.GET("/info", middleware.AuthMiddleware(), user.Info)
	userGroup.GET("/logout", middleware.AuthMiddleware(), user.Logout)
	userGroup.GET("/delete", middleware.AuthMiddleware(), user.Delete)

	paramGroup := r.Group("/param", middleware.AuthMiddleware())
	paramGroup.POST("/addParam", param.AddParam)
	paramGroup.POST("/getParamByType", param.GetParamByType)
	paramGroup.POST("/getAllParam", param.GetAllParam)
	paramGroup.POST("/updateParam", param.UpdateParam)
	paramGroup.POST("/deleteParam", param.DeleteParam)
}
