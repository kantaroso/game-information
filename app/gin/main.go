package main

import (

	// ginを利用する
	// http://sekitaka-1214.hatenablog.com/entry/2016/08/11/153816
	"github.com/gin-gonic/gin"

	// controller の 処理はモジュール化
	// https://qiita.com/tkj06/items/a5f79417935100045650
	"local.packages/game-information/controllers"
)

func main() {

	router := gin.Default()

	router.GET("/common/pv", func(c *gin.Context) { controllers.Pv(c) })
	router.GET("/maker/list", func(c *gin.Context) { controllers.GetMakerList(c) })
	router.GET("/maker/detail/:path", func(c *gin.Context) { controllers.GetMakerInfo(c) })
	router.GET("/maker/videos/:path", func(c *gin.Context) { controllers.GetMakerVideos(c) })

	router.Run(":8090")
}
