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

	router.GET("/health", func(c *gin.Context) { controllers.Health(c) })
	router.GET("/common/pv", func(c *gin.Context) { controllers.Pv(c) })
	router.GET("/maker/list", func(c *gin.Context) { controllers.GetMakerList(c) })
	router.GET("/maker/detail/:path", func(c *gin.Context) { controllers.GetMakerInfo(c) })
	router.GET("/maker/videos/:path", func(c *gin.Context) { controllers.GetMakerVideos(c) })
	// テスト用
	router.GET("/debug304", func(c *gin.Context) { controllers.Debug304(c) })
	router.GET("/debug403", func(c *gin.Context) { controllers.Debug403(c) })
	router.GET("/debug500", func(c *gin.Context) { controllers.Debug500(c) })
	router.GET("/common/pv_nocache", func(c *gin.Context) { controllers.Pv(c) })
	router.GET("/common/pv_cachepurge", func(c *gin.Context) { controllers.Pv(c) })

	router.Run(":8080")
}
