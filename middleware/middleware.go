package middleware

import (
	"github.com/ZEL-30/gin-web-app/infrastructure/config"
	"github.com/ZEL-30/gin-web-app/infrastructure/logger"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Init(router *gin.Engine) {

	// 使用自定义的 JSON 格式错误报告中间件
	router.Use(JSONAppErrorReporter())

	// 设置日志记录级别为 Debug
	logrus.SetLevel(logrus.DebugLevel)
	router.Use(logger.LoggerToFile(logrus.StandardLogger()))

	// 使用 CORS 跨域中间件
	router.Use(CORS())

	// 根据配置设置 gin 的运行模式
	gin.SetMode(config.App.RunMode)
}
