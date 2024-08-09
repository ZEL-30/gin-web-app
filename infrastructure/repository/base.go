package repository

import (
	"fmt"
	"log"

	"github.com/ZEL-30/gin-web-app/entity"
	infra "github.com/ZEL-30/gin-web-app/infrastructure"

	"github.com/google/uuid"
	logger "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func NewDB() *gorm.DB {
	var (
		dbName, user, password, host string
		port                         int
		dbType                       infra.DatabaseType
		db                           *gorm.DB
		err                          error
	)

	// 加载数据库配置
	dbType = infra.DatabaseConfig.DBType
	dbName = infra.DatabaseConfig.DBName
	user = infra.DatabaseConfig.DBUser
	password = infra.DatabaseConfig.DBPassword
	host = infra.DatabaseConfig.DBHost
	port = infra.DatabaseConfig.DBPort

	switch dbType {
	case infra.MySQL:
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbName)
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
		})

	case infra.PostgreSQL:
		db, err = gorm.Open(postgres.New(postgres.Config{
			DSN:                  fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", host, user, password, dbName, port),
			PreferSimpleProtocol: true,
		}), &gorm.Config{NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		}})
	}

	if err != nil {
		log.Fatal("failed to open database")
	}

	return db
}

func InitDB(db *gorm.DB) {

	// 定义回调函数，在创建记录前生成 UUID
	var createCallback = func(db *gorm.DB) {
		idField := db.Statement.Schema.LookUpField("id")
		if idField != nil {
			_ = idField.Set(db.Statement.Context, db.Statement.ReflectValue, uuid.New().String())
		}
	}

	// 注册 create 回调函数，在执行 gorm:create 操作之前执行名为 uuid 的回调函数
	err := db.Callback().Create().Before("gorm:create").Register("uuid", createCallback)
	if err != nil {
		// 记录日志，如果注册回调出现错误则记录到日志中
		logger.Fatal("failed to register uuid hook")
	}

	// 获取底层的 sqlDB 实例
	sqlDB, _ := db.DB()
	// 设置连接的最大空闲时间为 10 秒
	sqlDB.SetConnMaxIdleTime(10)
	// 设置最大的并发连接数为 100
	sqlDB.SetMaxOpenConns(100)
	// 调用 migrate 函数更新数据库结构
	migrate(db)
}

func migrate(db *gorm.DB) {
	// 自动迁移数据库模式以匹配 Book 结构体
	err := db.AutoMigrate(&entity.User{})
	if err != nil {
		// 记录日志，如果迁移过程中出错则记录日志
		logger.Fatal("migration failed.")
	}
}
