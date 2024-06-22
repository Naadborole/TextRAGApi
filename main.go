package main

import (
	"github.com/Naadborole/TextRAGApi/models"
	"net/http"

	"github.com/Naadborole/TextRAGApi/database"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	server := gin.Default()
	server.GET("/index/:id", getIndex)
	server.GET("/indexList", getIndexList)
	server.POST("/index", postIndex)
	server.Run()
}

func getIndex(context *gin.Context) {
	id := context.Param("id")
	index, err := database.GetIndex(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Index not found"})
		return
	}
	context.JSON(http.StatusFound, index)
}

func getIndexList(context *gin.Context) {
	context.JSON(http.StatusOK, database.GetIndexList())
}

func postIndex(context *gin.Context) {
	var index models.Index
	if err := context.ShouldBindJSON(&index); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	status, err := database.AddIndex(index)
	if err != nil || status == false {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, index)
}
