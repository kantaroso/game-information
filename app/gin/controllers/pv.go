package controllers

import (
	"github.com/gin-gonic/gin"
	domainAccesslog "local.packages/game-information/lib/domain/accesslog"
)

// Pv トップページの処理
func Pv(c *gin.Context) {
	domainAccesslogInstance := domainAccesslog.GetInstance()
	domainAccesslogInstance.Register(c.Request)
	accessCount := domainAccesslogInstance.GetAccessCount()
	outjson(c, 200, gin.H{
		"pv": accessCount,
	})
}
