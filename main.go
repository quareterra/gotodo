package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var message string
var isLinethroughed bool

func main() {
	message = "initial message"

	router := gin.Default()
	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*")

	router.GET("/", getIndex)
	router.POST("/set", setMessage)
	router.POST("/switch", switchLinethrough)

	router.Run(":8080")
}

func setMessage(context *gin.Context) {
	message = context.PostForm("message")
	context.Redirect(http.StatusFound, "/")
}

func getIndex(context *gin.Context) {
	context.HTML(http.StatusOK, "index.tmpl", gin.H{
		"message":     		 	message,
		"isLinethroughed": 	isLinethroughed,
	})
}

func switchLinethrough(context *gin.Context) {
	isLinethroughed = !isLinethroughed
	context.Redirect(http.StatusFound, "/")
}
