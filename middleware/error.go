package middleware

import (
	"net/http"

	rep "github.com/ZEL-30/gin-web-app/representation"
	"github.com/gin-gonic/gin"
)

func JSONAppErrorReporter() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		detectedErrors := c.Errors.ByType(gin.ErrorTypeAny)

		if len(detectedErrors) > 0 {
			err := detectedErrors[0].Err
			var parsedError *rep.AppError
			switch e := err.(type) {
			case *rep.AppError:
				parsedError = e
			default:
				parsedError = &rep.AppError{
					Code:    http.StatusInternalServerError,
					Message: "Internal Server Error",
				}
			}
			c.AbortWithStatusJSON(parsedError.Code, parsedError)
			return
		}
	}
}
