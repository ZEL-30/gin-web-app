package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/ZEL-30/gin-web-app/config"
	"github.com/ZEL-30/gin-web-app/infrastructure/repository"
	"github.com/ZEL-30/gin-web-app/middleware"
	"github.com/ZEL-30/gin-web-app/router"
)

func main() {
	environment := flag.String("e", "default", "")

	// 初始化配置
	config.Init(*environment)

	// 初始化路由
	r := router.NewRouter()

	// 初始化中间件
	middleware.Init(r)

	// 初始化数据库
	db := repository.InitDB()

	// 注册路由
	router.Register(r, db)

	// 初始化数据
	if config.GetBool("database.init_data") {
		repository.InitData(db)
	}

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", config.GetInt("server.port")),
		Handler:        r,
		ReadTimeout:    time.Duration(60) * time.Second,
		WriteTimeout:   time.Duration(60) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// 启动服务器
	err := s.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
