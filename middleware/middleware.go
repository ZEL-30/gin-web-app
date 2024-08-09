package middleware

import (
	infra "github.com/ZEL-30/gin-web-app/infrastructure"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Init(router *gin.Engine) {

	// 使用自定义的 JSON 格式错误报告中间件
	router.Use(JSONAppErrorReporter())

	// 设置日志记录级别为 Debug
	logrus.SetLevel(logrus.DebugLevel)
	// 使用自定义的日志记录中间件，将日志输出到文件
	router.Use(LoggerToFile(logrus.StandardLogger()))
	// 使用 CORS 跨域中间件
	router.Use(CORS())
	// 根据配置设置 gin 的运行模式
	gin.SetMode(infra.AppConfig.RunMode)

}
