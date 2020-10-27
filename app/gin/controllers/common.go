package controllers

import (
	"github.com/gin-gonic/gin"
)

// outjson json出力
func outjson(c *gin.Context, code int, obj interface{}) {
	c.Header("access-control-allow-origin", "*")
	c.JSON(code, obj)
}
