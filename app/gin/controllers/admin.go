package controllers

import (
	domainMaker "local.packages/game-information/lib/domain/maker"

	"github.com/gin-gonic/gin"
)

// CreateMakerVideoAll 全てのメーカー動画更新
func CreateMakerVideoAll(c *gin.Context) {

	if !checkAdmin(c) {
		outNotFound(c)
		return
	}

	domainMakerInstance := domainMaker.GetInstance()
	makers := domainMakerInstance.GetMakerList()
	if len(*makers) == 0 {
		outjson(c, 200, gin.H{})
		return
	}
	for _, item := range *makers {
		if !domainMakerInstance.UpdateVideoList(item.ID) {
			outError(c)
			return
		}
	}
	outjson(c, 200, gin.H{})
}

// CreateMakerVideo 指定のメーカー動画更新
func CreateMakerVideo(c *gin.Context) {

	if !checkAdmin(c) {
		outNotFound(c)
		return
	}

	path := c.Param("path")
	domainMakerInstance := domainMaker.GetInstance()
	maker := domainMakerInstance.GetMaker(path)
	if maker.ID == 0 {
		outNotFound(c)
		return
	}

	if !domainMakerInstance.UpdateVideoList(maker.ID) {
		outError(c)
		return
	}
	outjson(c, 200, gin.H{})
}
