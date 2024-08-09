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
	DBPassword: "123456",
	DBHost:     "127.0.0.1",
	DBPort:     3306,
	DBName:     "test_orm",
}

type ldapConfig struct {
	Url      string
	User     string
	Password string
	DC       string //domain component
}

var LDAPConfig = &ldapConfig{
	Url:      "ldap://127.0.0.1:10389",
	User:     "uid=admin,ou=system",
	Password: "secret",
	DC:       "dc=chuyang,dc=org",
}
