package config

type app struct {
	AppName    string
	RunMode    string
	HTTPPort   int
	StaticPath string
	LogPath    string
	UploadPath string
}

var App = &app{
	AppName:    "gin-web-app",
	RunMode:    "debug",
	HTTPPort:   8888,
	StaticPath: "D:/temp",
	LogPath:    "./log",
	UploadPath: "images",
}

type database struct {
	DBType     DatabaseType
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     int
	DBName     string
}

// 定义数据库类型的枚举
type DatabaseType int

const (
	MySQL DatabaseType = iota
	PostgreSQL
	SQLite
	MongoDB
)

var Database = &database{
	DBType:     MySQL,
	DBUser:     "root",
	DBPassword: "!Qw2!Qw2",
	DBHost:     "124.71.108.242",
	DBPort:     3306,
	DBName:     "example",
}
