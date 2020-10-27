package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kantaroso/game-information/lib/accesslog"
)

// Pv トップページの処理
func Pv(c *gin.Context) {
	accesslog.Register(c.Request)
	accessCount := accesslog.GetAccessCount()
	outjson(c, 200, gin.H{
		"pv": accessCount,
	})
}
