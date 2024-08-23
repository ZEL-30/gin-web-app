package router

import (
	"github.com/ZEL-30/gin-web-app/handler"
	"github.com/ZEL-30/gin-web-app/infrastructure/auth"
	"github.com/ZEL-30/gin-web-app/infrastructure/config"
	"github.com/ZEL-30/gin-web-app/infrastructure/repository"
	"github.com/ZEL-30/gin-web-app/middleware"
	"github.com/ZEL-30/gin-web-app/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	// 配置静态文件服务
	router.Static("/static", config.App.StaticPath)

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
	userRepo := repository.NewUserRepo(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService, authService)
	// 使用自定义的 JWT 中间件，并传入认证服务实例
	api.Use(middleware.NewJWTMiddleware(authService).JWT())
	{
		api.GET("/users/:id", userHandler.Get)
		api.GET("/users", userHandler.List)
		api.POST("/users", userHandler.Add)
		api.PUT("/users/:id", userHandler.Update)
		api.DELETE("/users/:id", userHandler.Delete)
	}
}
