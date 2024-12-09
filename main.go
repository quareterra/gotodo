package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var counter uint

func main() {
	counter = 0

	router := gin.Default()
	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*")

	router.GET("/", getIndex)

	router.Run(":80")
}

func getIndex(context *gin.Context) {
	var underline string
	if counter%2 == 0 {
		underline = "underline"
	}
	context.HTML(http.StatusOK, "index.tmpl", gin.H{
		"counter":   counter,
		"underline": underline,
	})
	counter++
}
