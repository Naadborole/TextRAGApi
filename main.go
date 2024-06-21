package main

import (
	"fmt"
	"net/http"

	"github.com/Naadborole/TextRAGApi/database"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	fmt.Println(database.GetIndexList())
	server := gin.Default()
	server.GET("/index", getIndex)
	server.Run()
}

func getIndex(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Request received"})
}
