package controllers

import (
	"encoding/json"

	domainMaker "local.packages/game-information/lib/domain/maker"

	"github.com/gin-gonic/gin"
)

// GetMakerList メーカー一覧
func GetMakerList(c *gin.Context) {
	outjson(c, 200, wrapperGetMakerList())
}

func wrapperGetMakerList() interface{} {
	domainMakerInstance := domainMaker.GetInstance()
	makers := domainMakerInstance.GetMakerList()
	if len(*makers) == 0 {
		return gin.H{}
	}
	var results []map[string]string
	var result map[string]string
	for _, v := range *makers {
		result = map[string]string{"code": v.Code, "name": v.Name}
		results = append(results, result)
	}
	return results
}

func MakeMakerListJson() {
	path := "/maker/list"
	body, _ := json.Marshal(wrapperGetMakerList())
	makeJson(path, body)
}

// GetMakerInfo トップページの処理
func GetMakerInfo(c *gin.Context) {
	path := c.Param("path")
	domainMakerInstance := domainMaker.GetInstance()
	maker := domainMakerInstance.GetMaker(path)
	if maker.ID == 0 {
		outjson(c, 404, gin.H{})
		return
	}
	detail := domainMakerInstance.GetDetail(maker.ID)
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
	domainMakerInstance := domainMaker.GetInstance()
	maker := domainMakerInstance.GetMaker(path)
	if maker.ID == 0 {
		outjson(c, 404, gin.H{})
		return
	}
	detail := domainMakerInstance.GetDetail(maker.ID)
	if detail.YoutubeChannelID == "" {
		outjson(c, 404, gin.H{})
		return
	}
	videos := domainMakerInstance.GetVideoList(maker.ID)
	if len(*videos) == 0 {
		outjson(c, 404, gin.H{})
		return
	}

	var response []map[string]string
	for _, item := range *videos {
		response = append(response, map[string]string{"ID": item.VideoID})
	}
	outjson(c, 200, response)
}
