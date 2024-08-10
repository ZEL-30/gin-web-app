package router

import (
	"github.com/ZEL-30/gin-web-app/controller"
	"github.com/ZEL-30/gin-web-app/infrastructure"
	"github.com/ZEL-30/gin-web-app/infrastructure/repository"
	"github.com/ZEL-30/gin-web-app/middleware"
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
	authController := controller.NewAuthContorller(authService)
	router.POST("/auth", authController.GetAuth)

	userRepo := repository.NewUserRepo(db)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService, authService)

	// 使用自定义的 JWT 中间件，并传入认证服务实例
	user := router.Group("/user")
	user.Use(middleware.NewJWTMiddleware(authService).JWT())
	{
		user.POST("/add", userController.Add)
		// user.GET("/user/:id", userController.Get)
		user.GET("/list", userController.GetAll)
		user.PUT("/update/:id", userController.Update)
		user.DELETE("/delete/:id", userController.Delete)
	}
}
