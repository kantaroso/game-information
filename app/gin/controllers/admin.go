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

	makers := domainMaker.GetMakerList()
	if len(*makers) == 0 {
		outjson(c, 200, gin.H{})
		return
	}
	for _, item := range *makers {
		if !domainMaker.UpdateVideoList(item.ID) {
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
	maker := domainMaker.GetMaker(path)
	if maker.ID == 0 {
		outNotFound(c)
		return
	}

	if !domainMaker.UpdateVideoList(maker.ID) {
		outError(c)
		return
	}
	outjson(c, 200, gin.H{})
}
