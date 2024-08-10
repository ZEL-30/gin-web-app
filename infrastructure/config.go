package infrastructure

type app struct {
	AppName  string
	RunMode  string
	HTTPPort int
	LogPath  string
}

var AppConfig = &app{
	AppName:  "gin-web-app",
	RunMode:  "debug",
	HTTPPort: 8888,
	LogPath:  "./log",
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

var DatabaseConfig = &database{
	DBType:     MySQL,
	DBUser:     "root",
	DBPassword: "!Qw2!Qw2",
	DBHost:     "124.71.108.242",
	DBPort:     3306,
	DBName:     "example",
}
