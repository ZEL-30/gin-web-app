package utils

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func ParamToInt(c *gin.Context, param string) (int, error) {
	paramStr := c.Param(param)
	paramInt, err := strconv.Atoi(paramStr)
	if err != nil {
		return 0, err
	}

	return paramInt, nil
}

func ParamToSlice(c *gin.Context, param string) ([]int, error) {
	paramStr := c.Param(param)
	paramStrings := strings.Split(paramStr, ",")
	paramSlice := make([]int, 0, len(paramStrings)) // 优化内存分配

	for _, paramStr := range paramStrings {
		paramInt, err := strconv.Atoi(paramStr)
		if err != nil {
			return nil, err
		}

		paramSlice = append(paramSlice, paramInt)
	}

	return paramSlice, nil
}
