package controllers

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	domainAccesslog "local.packages/game-information/lib/domain/accesslog"
)

const pathCommonPV = "/common/pv"

// Pv トップページの処理
func Pv(c *gin.Context) {
	if isProduction() {
		body := getJsonFromFile(pathCommonPV)
		outjsonFromText(c, 200, body)
		return
	}
	domainAccesslogInstance := domainAccesslog.GetInstance()
	domainAccesslogInstance.Register(c.Request)
	accessCount := domainAccesslogInstance.GetAccessCount()
	outjson(c, 200, gin.H{
		"pv": accessCount,
	})
}

func MakePvJson() {
	body, _ := json.Marshal(gin.H{"pv": 0})
	makeJson(pathCommonPV, body)
}
