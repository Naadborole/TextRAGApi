package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	server := gin.Default()
	server.GET("/index", getIndex)
	server.Run()
}

func getIndex(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Request received"})
}
