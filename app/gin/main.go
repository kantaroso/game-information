package main

import (

	// ginを利用する
	// http://sekitaka-1214.hatenablog.com/entry/2016/08/11/153816
	"github.com/gin-gonic/gin"
	// mysql driver
	_ "github.com/go-sql-driver/mysql"

	// controller の 処理はモジュール化
	// https://qiita.com/tkj06/items/a5f79417935100045650
	"github.com/kantaroso/game-information/controllers"
)

func main() {

	router := gin.Default()
	// API
	router.GET("/pv", func(c *gin.Context) { controllers.Pv(c) })
	router.GET("/maker/list", func(c *gin.Context) { controllers.GetMakerInfo(c) })

	router.Run(":8090")
}
