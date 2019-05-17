package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetUint ...URIに含まれるクエリストリングをuint型に変換する
func GetUint(c *gin.Context, key string) (uint, error) {
	i, err := strconv.Atoi(c.Param(key))
	return uint(i), err
}
