package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	chunkandembed "github.com/Naadborole/TextRAGApi/ChunkAndEmbed"
	generatechat "github.com/Naadborole/TextRAGApi/GenerateChat"
	"github.com/Naadborole/TextRAGApi/models"
	"github.com/google/uuid"

	"github.com/Naadborole/TextRAGApi/database"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	server := gin.Default()
	server.GET("/index/:id", getIndex)
	server.GET("/indexList", getIndexList)
	server.POST("/index", postIndex)
	server.POST("/upload", postUpload)
	server.POST("/embedAndStore", postEmbed)
	server.POST("/queryDoc", postQueryDoc)
	server.POST("/queryChat", postQueryChat)
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
	if err != nil || !status {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, index)
}

func postUpload(context *gin.Context) {
	file, _ := context.FormFile("file")
	log.Println(file.Filename)

	// Upload the file to specific dst.
	context.SaveUploadedFile(file, "./Data")

	context.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}

func postEmbed(c *gin.Context) {
	var reqBody struct {
		Text string `json:"text"`
	}
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	Texts := make([]string, 1)
	Texts[0] = reqBody.Text
	metadatas := make([]map[string]any, 1)
	metadatas[0] = make(map[string]any)
	metadatas[0]["Source"] = uuid.New().String()
	idString := metadatas[0]["Source"].(string)
	docs, err := chunkandembed.SplitText(Texts, metadatas)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	chunkandembed.Store.AddDocuments(context.Background(), docs)
	c.JSON(http.StatusAccepted, gin.H{"ID": idString})
}

func postQueryDoc(c *gin.Context) {
	var reqBody struct {
		Text string `json:"text"`
	}
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	docs, err := chunkandembed.Store.SimilaritySearch(context.Background(), reqBody.Text, 3)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusFound, docs)
}

func postQueryChat(c *gin.Context) {
	var reqBody struct {
		Text string `json:"text"`
	}
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, generatechat.GetResponse(reqBody.Text))
}
