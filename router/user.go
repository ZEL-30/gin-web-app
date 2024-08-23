package router

import (
	"github.com/ZEL-30/gin-web-app/domain"
	"github.com/ZEL-30/gin-web-app/handler"
	"github.com/ZEL-30/gin-web-app/infrastructure/repository"
	"github.com/ZEL-30/gin-web-app/middleware"
	"github.com/ZEL-30/gin-web-app/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func registerUserRoutes(api *gin.RouterGroup, db *gorm.DB, authService domain.AuthInterface) {
	// 用户路由设置
	userRepo := repository.NewUserRepo(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService, authService)

	api.Use(middleware.NewJWTMiddleware(authService).JWT())
	{
		api.GET("/users/:id", userHandler.Get)
		api.GET("/users", userHandler.List)
		api.POST("/users", userHandler.Add)
		api.PUT("/users/:id", userHandler.Update)
		api.DELETE("/users/:id", userHandler.Delete)
	}
}
