package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var elementList []string

func main() {
	router := gin.Default()
	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*")

	router.GET("/", getIndex)
	router.POST("/addElement", addElement)

	router.Run(":8080")
}

func addElement(context *gin.Context) {
	elementList = append(elementList, context.PostForm("element"))
	context.Redirect(http.StatusFound, "/")
}

func getIndex(context *gin.Context) {
	context.HTML(http.StatusOK, "index.tmpl", gin.H{
		"elementList": elementList,
	})
}
