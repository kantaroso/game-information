package controllers

import (
	"local.packages/game-information/lib/domain/accesslog"

	"github.com/gin-gonic/gin"
)

// Pv トップページの処理
func Pv(c *gin.Context) {
	accesslog.Register(c.Request)
	accessCount := accesslog.GetAccessCount()
	outjson(c, 200, gin.H{
		"pv": accessCount,
	})
}
