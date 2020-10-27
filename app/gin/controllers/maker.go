package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kantaroso/game-information/lib/youtube"
)

// GetMakerInfo トップページの処理
func GetMakerInfo(c *gin.Context) {
	//path = c.Param("path")
	//パスからid取得予定
	id := 1
	movies := youtube.GetMovieList(id)
	outjson(c, 200, gin.H{
		"movies": movies,
	})
}
