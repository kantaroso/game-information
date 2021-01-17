package controllers

import (
	"github.com/gin-gonic/gin"
)

func outjson(c *gin.Context, code int, obj interface{}) {
	c.Header("access-control-allow-origin", "*")
	c.JSON(code, obj)
}

func outNotFound(c *gin.Context) {
	outjson(c, 404, gin.H{})
}

func outError(c *gin.Context) {
	outjson(c, 500, gin.H{})
}

func checkAdmin(c *gin.Context) bool {
	if c.ClientIP() == "127.0.0.1" {
		return true
	}
	return false
}
