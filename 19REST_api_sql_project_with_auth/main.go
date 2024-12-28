package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.LoadHTMLGlob("templates/*")

	server.GET("/", home)

	server.GET("/events", getEvents)

	server.Run(":8080") // localhost:8080
}

func home(context *gin.Context) {
	context.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "Welcome To My First API Project",
	})
}

func getEvents(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Welcome API", "status": true})
}
