# gin-web-app

 - [项目介绍](#项目介绍)
 - [项目目录结构](#项目目录结构)
 - [项目技术栈](#项目技术栈)
 - [项目运行](#项目运行)
 - [项目构建](#项目构建) 

 ## 项目介绍
 

gin-web-app 是一个基于 gin 框架的 web 应用的企业级项目。

此示例包含：  
1. 企业级应用的分层架构
2. gin + 最新版gorm的示例
3. 四个middleware  
  1.1 cors  
  1.2 authentication  
  1.3 log  
  1.4 error message  
4. ldap 登录功能  
6. 基于sqlite的自动化测试 
7. 初始化/客户化脚本

### 一些关于技术的选择  
#### 1. 分层
  * DDD:  
    - controller  
    - service  
    - infrastructure
  * The clean architecture  https://zhuanlan.zhihu.com/p/64343082
#### 2. 数据库主键
* 自增优点：
  - 存储空间小
  - 查询效率高
* uuid优点：
  - 分布式友好
  - 避免暴露业务规模
* snowflake
### 3. 返回值struct vs. pointer  
https://www.ardanlabs.com/bookshop/2014/12/using-pointers-in-go.html
#### 4. 尽量避免在运行时创建实例
#### 5. ldap
apache directory studio的使用：
https://www.bilibili.com/video/BV1kh411h7yB?from=search&seid=4698533963557832110
#### 6. 运行在docker中的postgre sql
#### 7. 命名规范  
包的名称用单数  
https://rakyll.org/style-packages/  
https://github.com/golang-standards/project-layout/issues/7
### 8. 自动化测试


## 项目目录结构

```
gin-web-app
├── cmd
│   └── main.go  // 项目入口
├── config
│   ├── config.go  // 配置文件
│   ├── config_test.go  // 配置文件测试
│   └── config.yaml  // 配置文件模板
├── controllers
│   ├── auth.go  // 认证控制器
│   ├── user.go  // 用户控制器
│   └── health.go  // 健康检查控制器
├── docs  // 项目文档
├── infrastructure
│   ├── database
│   │   ├── db.go  // 数据库连接
│   │   ├── db_test.go  // 数据库连接测试
│   │   └── migrate  // 数据库迁移脚本
│   ├── ldap
│   │   ├── ldap.go  // ldap连接
│   │   └── ldap_test.go  // ldap连接测试
│   └── logger
│       ├── logger.go  // 日志
│       └── logger_test.go  // 日志测试
├── middleware
│   ├── authentication.go  // 认证中间件
│   ├── cors.go  // cors中间件
│   ├── error_message.go  // 错误信息中间件
│   └── log.go  // 日志中间件
├── models
│   ├── auth.go  // 认证模型
│   ├── user.go  // 用户模型
│   └── health.go  // 健康检查模型
├── routes
│   ├── auth.go  // 认证路由
│   ├── user.go  // 用户路由
│   └── health.go  // 健康检查路由
├── services
│   ├── auth.go  // 认证服务
│   ├── user.go  // 用户服务
│   └── health.go  // 健康检查服务
├── tests
│   ├── integration  // 集成测试
│   ├── unit  // 单元测试
│   └── utils  // 工具类
├── utils
│   ├── common.go  // 公共方法
│   ├── response.go  // 响应方法
│   └── utils_test.go  // 工具类测试
├── go.mod  // go mod文件
├── go.sum  // go mod文件
├── LICENSE  // 许可证文件
├── README.md  // 项目说明文件
└── main.go  // 项目入口文件
```


## 项目技术栈

- gin
- gorm
- sqlite3
- ldap
- jwt
- cors


## 项目运行

```
# 克隆项目
git clone https://github.com/ZEL-30/gin-web-app.git

# 进入项目目录
cd gin-web-app

# 安装依赖
go mod tidy


# 运行项目
go run cmd/main.go

# 或者编译项目
go build -o gin-web-app cmd/main.go

# 运行编译后的项目
./gin-web-app
```

## 项目构建

```
# 编译项目
go build -o gin-web-app cmd/main.go

# 运行编译后的项目
./gin-web-app
```

## 项目测试

```
# 单元测试
go test ./... -v -cover


# 集成测试
go test ./tests/integration -v -cover

```

## 项目部署

# 编译项目
```
go env -w GOOS=linux GOARCH=amd64
go build -o gin-web-app cmd/main.go
```
# 打包成服务

```
vim /etc/systemd/system/gin-web-app.service

[Unit]
Description=gin-web-app
After=network.target


[Service]
Type=simple
User=root  

WorkingDirectory=/path/to/gin-web-app
ExecStart=/path/to/gin-web-app/gin-web-app

Restart=always
RestartSec=10s


[Install]
WantedBy=multi-user.target
```

# 启动服务
```
systemctl start gin-web-app.service
systemctl enable gin-web-app.service
```

# 项目日志

```
# 项目日志
config/log.yaml
```

# 项目ldap配置

```
# 项目ldap配置
config/ldap.yaml
```

# 项目jwt配置

```
# 项目jwt配置
config/jwt.yaml
```

# 项目cors配置

```
# 项目cors配置
config/cors.yaml
```

# 项目数据库配置

```
# 项目数据库配置
config/database.yaml
```

## 项目迁移

```
# 备份数据库
cp database/db.db database/db.bak

# 迁移数据库
go run cmd/migrate/main.go

```

## 项目初始化

```
# 运行初始化脚本
go run cmd/init/main.go

```