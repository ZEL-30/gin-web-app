package repository

import (
	"fmt"
	"log"

	"github.com/ZEL-30/gin-web-app/config"
	"github.com/ZEL-30/gin-web-app/entity"
	"github.com/ZEL-30/gin-web-app/util"

	"github.com/google/uuid"
	logger "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func InitDB() *gorm.DB {
	var (
		dbDriver, dbName, user, password, host string
		port                                   int
		db                                     *gorm.DB
		err                                    error
	)

	// 加载数据库配置
	dbDriver = config.GetString("database.driver")
	dbName = config.GetString("database.name")
	user = config.GetString("database.username")
	password = config.GetString("database.password")
	host = config.GetString("database.host")
	port = config.GetInt("database.port")

	switch dbDriver {
	case "mysql":
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbName)
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
		})

	case "postgres":
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

	// 定义回调函数，在创建记录前生成 UUID
	var createCallback = func(db *gorm.DB) {
		idField := db.Statement.Schema.LookUpField("id")
		if idField != nil {
			_ = idField.Set(db.Statement.Context, db.Statement.ReflectValue, uuid.New().String())
		}
	}

	// 注册 create 回调函数，在执行 gorm:create 操作之前执行名为 uuid 的回调函数
	err = db.Callback().Create().Before("gorm:create").Register("uuid", createCallback)
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

	return db
}

// migrate 函数用于更新数据库结构
func migrate(db *gorm.DB) {
	// 自动迁移数据库模式以匹配 User 结构体
	err := db.AutoMigrate(
		&entity.User{},
	)

	// 记录日志，如果迁移过程中出错则记录日志
	if err != nil {
		logger.Fatal("migration failed.")
	}
}

// InitData 初始化数据
func InitData(db *gorm.DB) {
	// 创建管理员
	admin := &entity.User{
		Name:     "zel",
		Password: util.EncodeMD5("!Qw2!Qw2"),
		Email:    "1362848545@qq.com",
	}
	// 插入管理员信息
	db.Create(admin)
}
