package controllers

import (
	"encoding/json"
	"sort"

	domainMaker "game-information/lib/domain/maker"

	"github.com/gin-gonic/gin"
)

// -----------------------------------------------
// GetMakerList メーカー一覧
func GetMakerList(c *gin.Context) {
	if renderProduction(c) {
		return
	}
	outjson(c, 200, wrapperGetMakerList())
}

func wrapperGetMakerList() []map[string]string {
	domainMakerInstance := domainMaker.GetInstance()
	makers := domainMakerInstance.GetMakerList()
	if len(*makers) == 0 {
		return []map[string]string{}
	}
	var results []map[string]string
	var result map[string]string
	sort.SliceStable(*makers, func(i, j int) bool { return (*makers)[i].Name < (*makers)[j].Name })
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

// -----------------------------------------------
// GetMakerInfo トップページの処理
func GetMakerInfo(c *gin.Context) {
	if renderProduction(c) {
		return
	}
	path := c.Param("path")
	result := wrapperGetMakerInfo(path)
	if _, exist := result["code"]; !exist {
		outjson(c, 404, gin.H{})
		return
	}
	outjson(c, 200, result)
}

func wrapperGetMakerInfo(path string) map[string]string {
	domainMakerInstance := domainMaker.GetInstance()
	maker := domainMakerInstance.GetMaker(path)
	if maker.ID == 0 {
		return map[string]string{}
	}
	detail := domainMakerInstance.GetDetail(maker.ID)
	return map[string]string{
		"code":         maker.Code,
		"name":         maker.Name,
		"ohp":          detail.OHP,
		"twitter_name": detail.TwitterName,
	}
}

func MakeMakerInfoJson() {
	domainMakerInstance := domainMaker.GetInstance()
	makers := domainMakerInstance.GetMakerList()
	if len(*makers) == 0 {
		return
	}
	path := ""
	for _, v := range *makers {
		path = "/maker/detail/" + v.Code
		body, _ := json.Marshal(wrapperGetMakerInfo(v.Code))
		makeJson(path, body)
	}
}

// -----------------------------------------------
// GetMakerVideos トップページの処理
func GetMakerVideos(c *gin.Context) {
	if renderProduction(c) {
		return
	}
	path := c.Param("path")
	result := wrapperGetMakerVideos(path)
	if len(result) == 0 {
		outjson(c, 404, gin.H{})
		return
	}
	outjson(c, 200, result)
}

func wrapperGetMakerVideos(path string) []map[string]string {
	domainMakerInstance := domainMaker.GetInstance()
	maker := domainMakerInstance.GetMaker(path)
	if maker.ID == 0 {
		return []map[string]string{}
	}
	detail := domainMakerInstance.GetDetail(maker.ID)
	if detail.YoutubeChannelID == "" {
		return []map[string]string{}
	}
	videos := domainMakerInstance.GetVideoList(maker.ID)
	if len(*videos) == 0 {
		return []map[string]string{}
	}

	var response []map[string]string
	for _, item := range *videos {
		response = append(response, map[string]string{"id": item.VideoID, "title": item.Title})
	}
	return response
}

func MakeMakerVideosJson() {
	domainMakerInstance := domainMaker.GetInstance()
	makers := domainMakerInstance.GetMakerList()
	if len(*makers) == 0 {
		return
	}
	path := ""
	for _, v := range *makers {
		path = "/maker/videos/" + v.Code
		body, _ := json.Marshal(wrapperGetMakerVideos(v.Code))
		makeJson(path, body)
	}
}
