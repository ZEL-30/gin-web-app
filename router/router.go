package router

import (
	"github.com/ZEL-30/gin-web-app/config"
	"github.com/ZEL-30/gin-web-app/handler"
	"github.com/ZEL-30/gin-web-app/infrastructure/auth"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	// 配置静态文件服务
	router.Static("/static", config.GetString("server.static_path"))

	return router
}

// 注册路由
func Register(router *gin.Engine, db *gorm.DB) {
	api := router.Group("/api")

	// 注册认证路由
	authService := auth.NewAuthService(db)
	authHandler := handler.NewAuthHandler(authService)
	api.POST("/auth", authHandler.GetAuth)

	// 注册用户路由
	registerUserRoutes(api, db, authService)

}
