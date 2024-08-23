package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

var config *viper.Viper

var configTemplate = `server:
  app_name: gin-web-app
  port: 8080
  mode: debug
  static_path: /temp
  log_path: ./logs

database:
  driver: mysql
  host: localhost
  port: 3306
  name: example
  username: root
  password: 123456
  init_data: false
`

// Init is an exported method that takes the environment starts the viper
// (external lib) and returns the configuration struct.
func Init(env string) {

	checkConfig("config/default.yaml")

	var err error
	config = viper.New()
	config.SetConfigType("yaml")
	config.SetConfigName("default")
	config.AddConfigPath("config/")
	err = config.ReadInConfig()
	if err != nil {
		log.Fatal("error on parsing default configuration file")
	}

	envConfig := viper.New()
	envConfig.SetConfigType("yaml")
	envConfig.AddConfigPath("config/")
	envConfig.SetConfigName(env)
	err = envConfig.ReadInConfig()
	if err != nil {
		log.Fatal("error on parsing env configuration file")
	}

	config.MergeConfigMap(envConfig.AllSettings())
}

// 检查配置文件是否存在，不存在则创建
func checkConfig(configPath string) {

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		// 创建配置文件
		f, err := os.Create(configPath)
		if err != nil {
			log.Fatal("error on creating configuration file")
		}
		defer f.Close()
		f.WriteString(configTemplate)
	}

}

func relativePath(basedir string, path *string) {
	p := *path
	if len(p) > 0 && p[0] != '/' {
		*path = filepath.Join(basedir, p)
	}
}

func GetString(key string) string {
	return config.GetString(key)
}

func GetBool(key string) bool {
	return config.GetBool(key)
}

func GetInt(key string) int {
	return config.GetInt(key)
}

func Set(key string, value interface{}) {
	config.Set(key, value)
}
