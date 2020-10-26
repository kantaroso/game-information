package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kantaroso/game-information/lib/youtube"
)

// GetMakerInfo トップページの処理
func GetMakerInfo(c *gin.Context) {
	movies := youtube.GetMovieList()
	outjson(c, 200, gin.H{
		"movies": movies,
	})
}
