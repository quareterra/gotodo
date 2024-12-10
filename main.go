package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Element struct {
	Name			string
	IsCrossed	bool
}

var elementList []Element

func main() {
	router := gin.Default()
	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*")

	router.GET("/", getIndex)
	router.POST("/addElement", addElement)
	router.POST("/removeElement", removeElement)
	router.POST("/markElement", markElement)

	router.Run(":8080")
}

func addElement(context *gin.Context) {
	elementList = append(elementList, Element{context.PostForm("element"), false})
	context.Redirect(http.StatusFound, "/")
}

func removeElement(context *gin.Context) {
	index, _ := strconv.Atoi(context.PostForm("index"))
	elementList = append(elementList[:index], elementList[index+1:]...)
	context.Redirect(http.StatusFound, "/")
}

func markElement(context *gin.Context) {
	index, _ := strconv.Atoi(context.PostForm("index"))
	elementList[index].IsCrossed = !elementList[index].IsCrossed
	context.Redirect(http.StatusFound, "/")
}

func getIndex(context *gin.Context) {
	context.HTML(http.StatusOK, "index.tmpl", gin.H{
		"elementList": elementList,
	})
}
