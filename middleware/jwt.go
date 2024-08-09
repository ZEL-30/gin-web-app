package middleware

import (
	"time"

	"github.com/ZEL-30/gin-web-app/domain"
	rep "github.com/ZEL-30/gin-web-app/representation"
	"github.com/gin-gonic/gin"
)

type jwtMiddleware struct {
	authService domain.AuthInterface
}

func NewJWTMiddleware(authService domain.AuthInterface) *jwtMiddleware {
	return &jwtMiddleware{
		authService,
	}
}

func (s *jwtMiddleware) JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		authErr := rep.AppError{
			Code:    401,
			Message: "You are not authorized.",
		}
		var appErr rep.AppError

		token := s.authService.ExtractToken(c)
		if token == "" {
			appErr = authErr
		} else {
			claims, err := s.authService.ParseToken(token)
			if err != nil || time.Now().Unix() > claims.ExpiresAt {
				appErr = authErr
			}
		}

		if appErr.Code > 0 {
			c.JSON(appErr.Code, appErr)
			c.Abort()
			return
		}

		c.Next()
	}
}
