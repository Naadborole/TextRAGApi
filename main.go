package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	chunkandembed "github.com/Naadborole/TextRAGApi/ChunkAndEmbed"
	generatechat "github.com/Naadborole/TextRAGApi/GenerateChat"
	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.POST("/upload", postUpload)
	server.POST("/embedAndStore", postEmbed)
	server.POST("/queryDoc", postQueryDoc)
	server.POST("/queryChat", postQueryChat)
	server.Run()
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
		Text string `json:"query"`
	}
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, generatechat.GetResponse(reqBody.Text))
	//_ = generatechat.GetResponse(reqBody.Text)
	//mem, err := generatechat.RAGChain.GetMemory().LoadMemoryVariables(context.Background(), make(map[string]any))
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	//}
	//c.JSON(http.StatusAccepted, mem)
}
