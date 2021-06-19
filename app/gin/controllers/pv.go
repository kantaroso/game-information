package controllers

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	domainAccesslog "local.packages/game-information/lib/domain/accesslog"
)

// Pv トップページの処理
func Pv(c *gin.Context) {
	if renderProduction(c) {
		return
	}
	domainAccesslogInstance := domainAccesslog.GetInstance()
	domainAccesslogInstance.Register(c.Request)
	accessCount := domainAccesslogInstance.GetAccessCount()
	outjson(c, 200, gin.H{
		"pv": accessCount,
	})
}

// Pv 静的ページ
func MakePvJson() {
	body, _ := json.Marshal(gin.H{"pv": 0})
	makeJson("/common/pv", body)
}
