package controllers

import (
	domainMaker "local.packages/game-information/lib/domain/maker"

	"github.com/gin-gonic/gin"
)

// GetMakerList メーカー一覧
func GetMakerList(c *gin.Context) {
	makers := domainMaker.GetMakerList()
	if len(*makers) == 0 {
		outjson(c, 200, gin.H{})
		return
	}
	var results []map[string]string
	var result map[string]string
	for _, v := range *makers {
		result = map[string]string{"code": v.Code, "name": v.Name}
		results = append(results, result)
	}
	outjson(c, 200, results)
}

// GetMakerInfo トップページの処理
func GetMakerInfo(c *gin.Context) {
	path := c.Param("path")
	maker := domainMaker.GetMaker(path)
	if maker.ID == 0 {
		outNotFound(c)
		return
	}
	detail := domainMaker.GetDetail(maker.ID)
	outjson(c, 200, gin.H{
		"code":         maker.Code,
		"name":         maker.Name,
		"ohp":          detail.OHP,
		"twitter_name": detail.TwitterName,
	})
}

// GetMakerVideos トップページの処理
func GetMakerVideos(c *gin.Context) {
	path := c.Param("path")
	maker := domainMaker.GetMaker(path)
	if maker.ID == 0 {
		outNotFound(c)
		return
	}
	detail := domainMaker.GetDetail(maker.ID)
	if detail.YoutubeChannelID == "" {
		outNotFound(c)
		return
	}
	videos := domainMaker.GetVideoList(maker.ID)
	if len(*videos) == 0 {
		outNotFound(c)
		return
	}

	var response []map[string]string
	for _, item := range *videos {
		response = append(response, map[string]string{"ID": item.VideoID})
	}
	outjson(c, 200, response)
}
