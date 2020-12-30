package router

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/huahuayu/go-gin-app/common/config"
	"github.com/huahuayu/go-gin-app/handler/param"
	"github.com/huahuayu/go-gin-app/handler/user"
	"github.com/huahuayu/go-gin-app/middleware"
	"os"
)

func Init(r *gin.Engine) {
	r.LoadHTMLGlob(config.App.Server.HtmlPath + string(os.PathSeparator) + "*.html")
	r.Use(static.Serve("/", static.LocalFile(config.App.Server.HtmlPath, false)))
	r.GET("/health", HealthGET)
	r.GET("/", Index)
	r.NoRoute(NotFoundError)
	r.GET("/500", InternalServerError)

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
