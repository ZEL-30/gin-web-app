package handler

import (
	"net/http"
	"time"

	"github.com/ZEL-30/gin-web-app/domain"
	"github.com/gin-gonic/gin"
)

type authHandler struct {
	authService domain.AuthInterface
}

func NewAuthHandler(authService domain.AuthInterface) authHandler {
	return authHandler{
		authService,
	}
}

func (ah *authHandler) GetAuth(c *gin.Context) {
	type authInfo struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	auth := authInfo{}
	if err := c.ShouldBind(&auth); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	err := ah.authService.Auth(auth.Username, auth.Password)
	if err != nil {
		_ = c.Error(err)
		return
	}

	token, err := ah.authService.GenerateToken(auth.Username, auth.Password)
	if err != nil {
		_ = c.AbortWithError(401, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":  200,
		"msg":   "Logon suceeded.",
		"token": token,
	})
}

// might not useful
func (ah *authHandler) CheckAuth(c *gin.Context) {
	token := ah.authService.ExtractToken(c)
	var rtnVal bool
	if token == "" {
		rtnVal = false
	} else {
		claims, err := ah.authService.ParseToken(token)
		if err != nil {
			rtnVal = false
		} else if time.Now().Unix() > claims.ExpiresAt {
			rtnVal = false
		} else {
			rtnVal = true
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"auth": rtnVal,
	})
}
