package domain

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// 定义 Claims 结构体，用于在 JWT 中存储用户信息
type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

// 定义 AuthInterface 接口，包含了认证所需的五个方法
type AuthInterface interface {
	// Auth 方法用来认证用户，接收用户名和密码作为参数，返回错误信息
	Auth(user string, password string) error
	// GenerateToken 方法用来根据用户名和密码生成 JWT 令牌，返回令牌 string 和错误信息
	GenerateToken(username, password string) (string, error)
	// ParseToken 方法用来解析 JWT 令牌，接收令牌 string 作为参数，返回 Claims 指针和错误信息
	ParseToken(token string) (*Claims, error)
	// GetUserFromToken 方法用来从 JWT 令牌中提取用户名，接收令牌 string 作为参数，返回用户名 string 和错误信息
	GetUserFromToken(token string) (string, error)
	// ExtractToken 方法用来从 gin.Context 中提取 JWT 令牌，接收 gin.Context 指针作为参数，返回令牌 string
	ExtractToken(c *gin.Context) string
}
