package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var counter uint

func main() {
	counter = 0

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/", getIndex)

	router.Run(":8080")
}

func getIndex(context *gin.Context) {
	context.HTML(http.StatusOK, "index.tmpl", gin.H{
		"counter": counter,
	})
	counter++
}
