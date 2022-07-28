package controllers

import (
	"github.com/gin-gonic/gin"
)

func Debug304(c *gin.Context) {
	outjson(c, 304, gin.H{})
}
func Debug403(c *gin.Context) {
	outjson(c, 403, gin.H{})
}
func Debug500(c *gin.Context) {
	outjson(c, 500, gin.H{})
}
