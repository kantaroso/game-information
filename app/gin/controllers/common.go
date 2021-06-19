package controllers

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gin-gonic/gin"
)

func outjson(c *gin.Context, code int, obj interface{}) {
	c.Header("access-control-allow-origin", "*")
	c.JSON(code, obj)
}

func outjsonFromText(c *gin.Context, code int, body string) {
	c.Header("access-control-allow-origin", "*")
	c.Header("Content-Type", "application/json")
	c.String(code, body)
}

func makeJson(path string, body []byte) {
	err := ioutil.WriteFile(makeTargetJsonFilePath(path), body, 0666)
	if err != nil {
		fmt.Println(err)
	}
}

func getJsonFromFile(path string) string {
	bytes, err := ioutil.ReadFile(makeTargetJsonFilePath(path))
	if err != nil {
		fmt.Println(err)
		return "{}"
	}
	return string(bytes)
}

func makeTargetJsonFilePath(path string) string {
	return "assets/controllers" + path + ".json"
}

func isProduction() bool {
	if os.Getenv("IS_PROD") == "true" {
		return true
	}
	return false
}

func renderProduction(c *gin.Context) bool {
	if !isProduction() {
		return false
	}
	outjsonFromText(c, 200, getJsonFromFile(c.Request.URL.Path))
	return true
}
