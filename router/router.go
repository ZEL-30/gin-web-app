package router

import (
	"github.com/ZEL-30/gin-web-app/controller"
	"github.com/ZEL-30/gin-web-app/infrastructure"
	"github.com/ZEL-30/gin-web-app/infrastructure/repository"
	"github.com/ZEL-30/gin-web-app/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	return router
}

// 注册路由
func Register(router *gin.Engine, db *gorm.DB) {

	authService := infrastructure.NewAuthService(db)
	userRepo := repository.NewUserRepo(db)
	userService := service.NewUserService(userRepo)

	authController := controller.NewAuthContorller(authService)
	userController := controller.NewUserHandler(userService, authService)

	// 设置认证路由，POST 请求到 /auth 路径
	router.POST("/auth", authController.GetAuth)

	user := router.Group("/user")

	user.POST("/register", userController.Register)
	// // 使用自定义的 JWT 中间件，并传入认证服务实例
	// user.Use(middleware.NewJWTMiddleware(authService).JWT())
	// {
	// 	// 设置图书相关的路由
	// 	user.GET("/user/:id", userController.Get)
	// 	user.GET("/user", userController.GetAll)
	// 	user.PUT("/user/:id", userController.Update)
	// 	user.DELETE("/user/:id", userController.Delete)
	// }
}
